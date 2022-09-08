// Copyright 2022 Arrikto Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"database/sql"
	"fmt"

	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var featureServiceColumns = []string{
	"feature_services.id",
	"feature_services.project_id",
	"feature_services.name",
	"feature_services.tags",
	"feature_services.description",
	"feature_services.owner",
	"feature_services.logging_config",
	"feature_services.created_timestamp",
	"feature_services.last_updated_timestamp",
}

type FeatureServiceStoreInterface interface {
	GetFeatureService(name string, projectId string) (*model.FeatureService, error)
	CreateFeatureService(*model.FeatureService) (*model.FeatureService, error)
	UpdateFeatureService(*model.FeatureService) (*model.FeatureService, error)
	DeleteFeatureService(name string, projectId string) error
	ListFeatureServices(projectId string) ([]*model.FeatureService, error)
}

type FeatureServiceStore struct {
	db       *DB
	fvpStore *FVPStore
	time     util.TimeInterface
	uuid     util.UUIDGeneratorInterface
}

func (s *FeatureServiceStore) GetFeatureService(name string, projectId string) (*model.FeatureService, error) {
	sql, args, err := sq.
		Select(featureServiceColumns...).
		From("feature_services").
		Where(sq.And{sq.Eq{"feature_services.name": name}, sq.Eq{"feature_services.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get feature service: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get feature service: %v", err.Error())
	}
	defer r.Close()

	feature_services, err := s.scanRows(r)
	if err != nil || len(feature_services) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get feature service: %v", err.Error())
	}
	if len(feature_services) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.FeatureService), fmt.Sprint(name))
	}

	fvps, err := s.fvpStore.ListFVPs(feature_services[0].Id)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get feature view projections: %v", err.Error())
	}
	feature_services[0].FeatureViewProjections = append(feature_services[0].FeatureViewProjections, fvps...)

	return feature_services[0], nil
}

func (s *FeatureServiceStore) CreateFeatureService(fs *model.FeatureService) (*model.FeatureService, error) {
	newFeatureService := *fs

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create feature service id.")
	}
	newFeatureService.Id = id.String()

	sql, args, err := sq.
		Insert("feature_services").
		SetMap(
			sq.Eq{
				"id":                     newFeatureService.Id,
				"project_id":             newFeatureService.ProjectId,
				"name":                   newFeatureService.Name,
				"tags":                   newFeatureService.Tags,
				"description":            newFeatureService.Description,
				"owner":                  newFeatureService.Owner,
				"logging_config":         newFeatureService.LoggingConfig,
				"created_timestamp":      newFeatureService.CreatedTimestamp,
				"last_updated_timestamp": newFeatureService.LastUpdatedTimestamp,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert feature service: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create a new transaction to create feature service.")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new feature service. The name %v already exists. Please specify a new name.", fs.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add feature service to feature_services table: %v",
			err.Error())
	}

	for _, fvp := range newFeatureService.FeatureViewProjections {
		newFvp, err := s.fvpStore.CreateFVP(tx, fvp, newFeatureService.Id)
		if err != nil {
			tx.Rollback()
			return nil, util.NewInternalServerError(err, "Failed to store feature view projection %v for feature service %v ", fvp.FVName, fs.Name)
		}
		fvp.Id = newFvp.Id
		fvp.FSId = newFvp.FSId
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to store feature service %v", fs.Name)
	}

	return &newFeatureService, nil
}

func (s *FeatureServiceStore) UpdateFeatureService(fs *model.FeatureService) (*model.FeatureService, error) {
	updatedFeatureService := *fs

	sql, args, err := sq.
		Update("feature_services").
		SetMap(
			sq.Eq{
				"tags":                   updatedFeatureService.Tags,
				"description":            updatedFeatureService.Description,
				"owner":                  updatedFeatureService.Owner,
				"logging_config":         updatedFeatureService.LoggingConfig,
				"created_timestamp":      updatedFeatureService.CreatedTimestamp,
				"last_updated_timestamp": updatedFeatureService.LastUpdatedTimestamp,
			}).
		Where(sq.And{sq.Eq{"feature_services.name": updatedFeatureService.Name}, sq.Eq{"feature_services.project_id": updatedFeatureService.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update feature service: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create a new transaction to update feature service.")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update feature service: %s", err.Error())
	}

	err = s.fvpStore.DeleteFVPs(tx, updatedFeatureService.Id)
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update feature view projections of feature service: %s", err.Error())
	}

	for _, fvp := range updatedFeatureService.FeatureViewProjections {
		updFvp, err := s.fvpStore.CreateFVP(tx, fvp, updatedFeatureService.Id)
		if err != nil {
			tx.Rollback()
			return nil, util.NewInternalServerError(err, "Failed to update feature view projection %v to table for feature service %v ", fvp.FVName, fs.Name)
		}
		fvp.Id = updFvp.Id
		fvp.FSId = updFvp.FSId
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update feature service %v", fs.Name)
	}

	return &updatedFeatureService, nil
}

func (s *FeatureServiceStore) DeleteFeatureService(name string, projectId string) error {
	fs, err := s.GetFeatureService(name, projectId)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete feature service: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create a new transaction to delete feature service.")
	}

	err = s.fvpStore.DeleteFVPs(tx, fs.Id)
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete feature view projections of feature service: %v", err.Error())
	}

	sql, args, err := sq.
		Delete("feature_services").
		Where(sq.And{sq.Eq{"feature_services.name": name}, sq.Eq{"feature_services.project_id": projectId}}).
		ToSql()
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to create query to delete feature service: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete feature service: %v", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete feature service: %v", err.Error())
	}

	return nil
}

func (s *FeatureServiceStore) ListFeatureServices(projectId string) ([]*model.FeatureService, error) {
	sql, args, err := sq.
		Select(featureServiceColumns...).
		From("feature_services").
		Where(sq.Eq{"feature_services.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list feature services: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list feature services: %v", err.Error())
	}
	defer r.Close()

	feature_services, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list feature services: %v", err.Error())
	}

	for _, fs := range feature_services {
		fvps, err := s.fvpStore.ListFVPs(fs.Id)
		if err != nil {
			return nil, util.NewInternalServerError(err, "Failed to get feature view projections: %v", err.Error())
		}
		fs.FeatureViewProjections = append(fs.FeatureViewProjections, fvps...)
	}

	return feature_services, nil
}

func (s *FeatureServiceStore) scanRows(rows *sql.Rows) ([]*model.FeatureService, error) {
	var feature_services []*model.FeatureService
	for rows.Next() {
		var id, name, project_id string
		var description, owner sql.NullString
		var tags, logging_config []byte
		var created_timestamp, last_updated_timestamp sql.NullTime

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&tags,
			&description,
			&owner,
			&logging_config,
			&created_timestamp,
			&last_updated_timestamp); err != nil {
			return nil, err
		}

		feature_services = append(feature_services, &model.FeatureService{
			Id:                   id,
			ProjectId:            project_id,
			Name:                 name,
			Tags:                 tags,
			Description:          description.String,
			Owner:                owner.String,
			LoggingConfig:        logging_config,
			CreatedTimestamp:     created_timestamp.Time,
			LastUpdatedTimestamp: last_updated_timestamp.Time,
		})
	}

	return feature_services, nil
}

func NewFeatureServiceStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *FeatureServiceStore {
	fvpStore := NewFVPStore(db, time, uuid)
	return &FeatureServiceStore{db: db, fvpStore: fvpStore, time: time, uuid: uuid}
}
