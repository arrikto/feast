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
var odfeatureViewColumns = []string{
	"on_demand_feature_views.id",
	"on_demand_feature_views.project_id",
	"on_demand_feature_views.name",
	"on_demand_feature_views.sources",
	"on_demand_feature_views.udf_name",
	"on_demand_feature_views.udf_body",
	"on_demand_feature_views.description",
	"on_demand_feature_views.tags",
	"on_demand_feature_views.owner",
	"on_demand_feature_views.created_timestamp",
	"on_demand_feature_views.last_updated_timestamp",
}

type OnDemandFeatureViewStoreInterface interface {
	GetOnDemandFeatureView(name string, projectId string) (*model.OnDemandFeatureView, error)
	CreateOnDemandFeatureView(*model.OnDemandFeatureView) (*model.OnDemandFeatureView, error)
	UpdateOnDemandFeatureView(*model.OnDemandFeatureView) (*model.OnDemandFeatureView, error)
	DeleteOnDemandFeatureView(name string, projectId string) error
	ListOnDemandFeatureViews(projectId string) ([]*model.OnDemandFeatureView, error)
}

type OnDemandFeatureViewStore struct {
	db           *DB
	featureStore *OnDemandFeatureStore
	time         util.TimeInterface
	uuid         util.UUIDGeneratorInterface
}

func (s *OnDemandFeatureViewStore) GetOnDemandFeatureView(name string, projectId string) (*model.OnDemandFeatureView, error) {
	sql, args, err := sq.
		Select(odfeatureViewColumns...).
		From("on_demand_feature_views").
		Where(sq.And{sq.Eq{"on_demand_feature_views.name": name}, sq.Eq{"on_demand_feature_views.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get on demand feature view: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get on demand feature view: %v", err.Error())
	}
	defer r.Close()

	od_feature_views, err := s.scanRows(r)
	if err != nil || len(od_feature_views) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get on demand feature view: %v", err.Error())
	}
	if len(od_feature_views) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.OnDemandFeatureView), fmt.Sprint(name))
	}

	features, err := s.featureStore.ListFeatures(od_feature_views[0].Id)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get features: %v", err.Error())
	}
	od_feature_views[0].Features = append(od_feature_views[0].Features, features...)

	return od_feature_views[0], nil
}

func (s *OnDemandFeatureViewStore) CreateOnDemandFeatureView(odfv *model.OnDemandFeatureView) (*model.OnDemandFeatureView, error) {
	newOnDemandFeatureView := *odfv

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create on demand feature view id.")
	}
	newOnDemandFeatureView.Id = id.String()

	sql, args, err := sq.
		Insert("on_demand_feature_views").
		SetMap(
			sq.Eq{
				"id":                     newOnDemandFeatureView.Id,
				"project_id":             newOnDemandFeatureView.ProjectId,
				"name":                   newOnDemandFeatureView.Name,
				"sources":                newOnDemandFeatureView.Sources,
				"udf_name":               newOnDemandFeatureView.UdfName,
				"udf_body":               newOnDemandFeatureView.UdfBody,
				"description":            newOnDemandFeatureView.Description,
				"tags":                   newOnDemandFeatureView.Tags,
				"owner":                  newOnDemandFeatureView.Owner,
				"created_timestamp":      newOnDemandFeatureView.CreatedTimestamp,
				"last_updated_timestamp": newOnDemandFeatureView.LastUpdatedTimestamp,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert on demand feature view: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create a new transaction to create on demand feature view.")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new on demand feature view. The name %v already exists. Please specify a new name.", odfv.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add on demand feature view to on_deand_feature_views table: %v",
			err.Error())
	}

	for _, feature := range newOnDemandFeatureView.Features {
		_, err = s.featureStore.CreateFeature(tx, feature, newOnDemandFeatureView.Id)
		if err != nil {
			tx.Rollback()
			return nil, util.NewInternalServerError(err, "Failed to store feature %v for on demand feature view %v ", feature.Name, odfv.Name)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to store on demand feature view %v", odfv.Name)
	}

	return &newOnDemandFeatureView, nil
}

func (s *OnDemandFeatureViewStore) UpdateOnDemandFeatureView(odfv *model.OnDemandFeatureView) (*model.OnDemandFeatureView, error) {
	updatedOnDemandFeatureView := *odfv

	sql, args, err := sq.
		Update("on_demand_feature_views").
		SetMap(
			sq.Eq{
				"sources":                updatedOnDemandFeatureView.Sources,
				"udf_name":               updatedOnDemandFeatureView.UdfName,
				"udf_body":               updatedOnDemandFeatureView.UdfBody,
				"description":            updatedOnDemandFeatureView.Description,
				"tags":                   updatedOnDemandFeatureView.Tags,
				"owner":                  updatedOnDemandFeatureView.Owner,
				"created_timestamp":      updatedOnDemandFeatureView.CreatedTimestamp,
				"last_updated_timestamp": updatedOnDemandFeatureView.LastUpdatedTimestamp,
			}).
		Where(sq.And{sq.Eq{"on_demand_feature_views.name": updatedOnDemandFeatureView.Name}, sq.Eq{"on_demand_feature_views.project_id": updatedOnDemandFeatureView.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update on demand feature view: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create a new transaction to update on demand feature view.")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update on demand feature view: %s", err.Error())
	}

	err = s.featureStore.SafeUpdateFeatures(tx, odfv.Features, odfv.Id)
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update features of on demand feature view: %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, util.NewInternalServerError(err, "Failed to update on demand feature view %v", odfv.Name)
	}

	return &updatedOnDemandFeatureView, nil
}

func (s *OnDemandFeatureViewStore) DeleteOnDemandFeatureView(name string, projectId string) error {
	fv, err := s.GetOnDemandFeatureView(name, projectId)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete on demand feature view: %v", err.Error())
	}

	tx, err := s.db.Begin()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create a new transaction to delete on demand feature view.")
	}

	for _, feature := range fv.Features {
		err := s.featureStore.DeleteFeature(tx, feature.Name, feature.ODFVId)
		if err != nil {
			tx.Rollback()
			return util.NewInternalServerError(err, "Failed to delete features of feature view: %v", err.Error())
		}
	}

	sql, args, err := sq.
		Delete("on_demand_feature_views").
		Where(sq.And{sq.Eq{"on_demand_feature_views.name": name}, sq.Eq{"on_demand_feature_views.project_id": projectId}}).
		ToSql()
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to create query to delete on demand feature view: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete on demand feature view: %v", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return util.NewInternalServerError(err, "Failed to delete on demand feature view: %v", err.Error())
	}

	return nil
}

func (s *OnDemandFeatureViewStore) ListOnDemandFeatureViews(projectId string) ([]*model.OnDemandFeatureView, error) {
	sql, args, err := sq.
		Select(odfeatureViewColumns...).
		From("on_demand_feature_views").
		Where(sq.Eq{"on_demand_feature_views.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list on demand feature views: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list on demand feature views: %v", err.Error())
	}
	defer r.Close()

	od_feature_views, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list on demand feature views: %v", err.Error())
	}

	for _, fv := range od_feature_views {
		features, err := s.featureStore.ListFeatures(fv.Id)
		if err != nil {
			return nil, util.NewInternalServerError(err, "Failed to get features: %v", err.Error())
		}
		fv.Features = append(fv.Features, features...)
	}

	return od_feature_views, nil
}

func (s *OnDemandFeatureViewStore) scanRows(rows *sql.Rows) ([]*model.OnDemandFeatureView, error) {
	var od_feature_views []*model.OnDemandFeatureView
	for rows.Next() {
		var id, name, project_id string
		var description, owner, udf_name sql.NullString
		var sources, tags, udf_body []byte
		var created_timestamp, last_updated_timestamp sql.NullTime

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&sources,
			&udf_name,
			&udf_body,
			&description,
			&tags,
			&owner,
			&created_timestamp,
			&last_updated_timestamp); err != nil {
			return nil, err
		}

		od_feature_views = append(od_feature_views, &model.OnDemandFeatureView{
			Id:                   id,
			ProjectId:            project_id,
			Name:                 name,
			Sources:              sources,
			UdfName:              udf_name.String,
			UdfBody:              udf_body,
			Description:          description.String,
			Tags:                 tags,
			Owner:                owner.String,
			CreatedTimestamp:     created_timestamp.Time,
			LastUpdatedTimestamp: last_updated_timestamp.Time,
		})
	}

	return od_feature_views, nil
}

func NewOnDemandFeatureViewStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *OnDemandFeatureViewStore {
	featureStore := NewOnDemandFeatureStore(db, time, uuid)
	return &OnDemandFeatureViewStore{db: db, featureStore: featureStore, time: time, uuid: uuid}
}
