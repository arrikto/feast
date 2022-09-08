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
	"time"

	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var featureViewColumns = []string{
	"feature_views.id",
	"feature_views.project_id",
	"feature_views.name",
	"feature_views.entities",
	"feature_views.description",
	"feature_views.tags",
	"feature_views.owner",
	"feature_views.ttl",
	"feature_views.batch_source",
	"feature_views.stream_source",
	"feature_views.online",
	"feature_views.created_timestamp",
	"feature_views.last_updated_timestamp",
}

type FeatureViewStoreInterface interface {
	GetFeatureView(name string, projectId string) (*model.FeatureView, error)
	CreateFeatureView(*model.FeatureView) (*model.FeatureView, error)
	UpdateFeatureView(*model.FeatureView) (*model.FeatureView, error)
	DeleteFeatureView(name string, projectId string) error
	ListFeatureViews(projectId string) ([]*model.FeatureView, error)
}

type FeatureViewStore struct {
	db           *DB
	featureStore *FeatureStore
	miStore      *MIStore
	time         util.TimeInterface
	uuid         util.UUIDGeneratorInterface
}

func (s *FeatureViewStore) GetFeatureView(name string, projectId string) (*model.FeatureView, error) {
	sql, args, err := sq.
		Select(featureViewColumns...).
		From("feature_views").
		Where(sq.And{sq.Eq{"feature_views.name": name}, sq.Eq{"feature_views.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get feature view: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get feature view: %v", err.Error())
	}
	defer r.Close()

	feature_views, err := s.scanRows(r)
	if err != nil || len(feature_views) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get feature view: %v", err.Error())
	}
	if len(feature_views) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.FeatureView), fmt.Sprint(name))
	}

	features, err := s.featureStore.ListFeatures(feature_views[0].Id)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get features: %v", err.Error())
	}
	feature_views[0].Features = append(feature_views[0].Features, features...)

	mis, err := s.miStore.ListMIs(feature_views[0].Id)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get materialization intervals: %v", err.Error())
	}
	feature_views[0].MIs = append(feature_views[0].MIs, mis...)

	return feature_views[0], nil
}

func (s *FeatureViewStore) CreateFeatureView(fv *model.FeatureView) (*model.FeatureView, error) {
	// WARNING: Doesn't create materialization intervals belonging to the feature view
	// There is a separate endpoint /ReportMI for this reason
	newFeatureView := *fv

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create feature view id.")
	}
	newFeatureView.Id = id.String()

	sql, args, err := sq.
		Insert("feature_views").
		SetMap(
			sq.Eq{
				"id":                     newFeatureView.Id,
				"project_id":             newFeatureView.ProjectId,
				"name":                   newFeatureView.Name,
				"entities":               newFeatureView.Entities,
				"description":            newFeatureView.Description,
				"tags":                   newFeatureView.Tags,
				"owner":                  newFeatureView.Owner,
				"ttl":                    newFeatureView.Ttl,
				"batch_source":           newFeatureView.BatchSource,
				"stream_source":          newFeatureView.StreamSource,
				"online":                 newFeatureView.Online,
				"created_timestamp":      newFeatureView.CreatedTimestamp,
				"last_updated_timestamp": newFeatureView.LastUpdatedTimestamp,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert feature view: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create a new transaction to create feature view.")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new feature view. The name %v already exists. Please specify a new name.", fv.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add feature view to feature_views table: %v",
			err.Error())
	}

	for _, feature := range newFeatureView.Features {
		newFeature, err := s.featureStore.CreateFeature(tx, feature, newFeatureView.Id)
		if err != nil {
			tx.Rollback()
			return nil, util.NewInternalServerError(err, "Failed to store feature %v for feature view %v ", feature.Name, fv.Name)
		}
		feature.Id = newFeature.Id
		feature.FVId = newFeature.FVId
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to store feature view %v", fv.Name)
	}

	return &newFeatureView, nil
}

func (s *FeatureViewStore) UpdateFeatureView(fv *model.FeatureView) (*model.FeatureView, error) {
	// WARNING: Doesn't update materialization intervals belonging to the feature view
	// There is a separate endpoint /ReportMI for this reason
	updatedFeatureView := *fv

	sql, args, err := sq.
		Update("feature_views").
		SetMap(
			sq.Eq{
				"entities":               updatedFeatureView.Entities,
				"description":            updatedFeatureView.Description,
				"tags":                   updatedFeatureView.Tags,
				"owner":                  updatedFeatureView.Owner,
				"ttl":                    updatedFeatureView.Ttl,
				"batch_source":           updatedFeatureView.BatchSource,
				"stream_source":          updatedFeatureView.StreamSource,
				"online":                 updatedFeatureView.Online,
				"created_timestamp":      updatedFeatureView.CreatedTimestamp,
				"last_updated_timestamp": updatedFeatureView.LastUpdatedTimestamp,
			}).
		Where(sq.And{sq.Eq{"feature_views.name": updatedFeatureView.Name}, sq.Eq{"feature_views.project_id": updatedFeatureView.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update feature view: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create a new transaction to update feature view.")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update feature view: %s", err.Error())
	}

	err = s.featureStore.SafeUpdateFeatures(tx, fv.Features, fv.Id)
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update features of feature view: %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update feature view %v", fv.Name)
	}

	return &updatedFeatureView, nil
}

func (s *FeatureViewStore) DeleteFeatureView(name string, projectId string) error {
	fv, err := s.GetFeatureView(name, projectId)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete feature view: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create a new transaction to delete feature view.")
	}

	for _, feature := range fv.Features {
		err := s.featureStore.DeleteFeature(tx, feature.Name, feature.FVId)
		if err != nil {
			tx.Rollback()
			return util.NewInternalServerError(err, "Failed to delete features of feature view: %v", err.Error())
		}
	}

	err = s.miStore.DeleteMIs(tx, fv.Id)
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete materialization intervals of feature view: %v", err.Error())
	}

	sql, args, err := sq.
		Delete("feature_views").
		Where(sq.And{sq.Eq{"feature_views.name": name}, sq.Eq{"feature_views.project_id": projectId}}).
		ToSql()
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to create query to delete feature view: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete feature view: %v", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete feature view: %v", err.Error())
	}

	return nil
}

func (s *FeatureViewStore) ListFeatureViews(projectId string) ([]*model.FeatureView, error) {
	sql, args, err := sq.
		Select(featureViewColumns...).
		From("feature_views").
		Where(sq.Eq{"feature_views.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list feature views: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list feature views: %v", err.Error())
	}
	defer r.Close()

	feature_views, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list feature views: %v", err.Error())
	}

	for _, fv := range feature_views {
		features, err := s.featureStore.ListFeatures(fv.Id)
		if err != nil {
			return nil, util.NewInternalServerError(err, "Failed to get features: %v", err.Error())
		}
		fv.Features = append(fv.Features, features...)

		mis, err := s.miStore.ListMIs(fv.Id)
		if err != nil {
			return nil, util.NewInternalServerError(err, "Failed to get materialization intervals: %v", err.Error())
		}
		fv.MIs = append(fv.MIs, mis...)
	}

	return feature_views, nil
}

func (s *FeatureViewStore) scanRows(rows *sql.Rows) ([]*model.FeatureView, error) {
	var feature_views []*model.FeatureView
	for rows.Next() {
		var id, name, project_id string
		var description, owner sql.NullString
		var entities, tags, batch_source, stream_source []byte
		var ttl sql.NullInt64
		var online sql.NullBool
		var created_timestamp, last_updated_timestamp sql.NullTime

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&entities,
			&description,
			&tags,
			&owner,
			&ttl,
			&batch_source,
			&stream_source,
			&online,
			&created_timestamp,
			&last_updated_timestamp); err != nil {
			return nil, err
		}

		feature_views = append(feature_views, &model.FeatureView{
			Id:                   id,
			ProjectId:            project_id,
			Name:                 name,
			Entities:             entities,
			Description:          description.String,
			Tags:                 tags,
			Owner:                owner.String,
			Ttl:                  time.Duration(ttl.Int64),
			BatchSource:          batch_source,
			StreamSource:         stream_source,
			Online:               online.Bool,
			CreatedTimestamp:     created_timestamp.Time,
			LastUpdatedTimestamp: last_updated_timestamp.Time,
		})
	}

	return feature_views, nil
}

func NewFeatureViewStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *FeatureViewStore {
	featureStore := NewFeatureStore(db, time, uuid)
	miStore := NewMIStore(db, time, uuid)

	return &FeatureViewStore{db: db, featureStore: featureStore, miStore: miStore, time: time, uuid: uuid}
}
