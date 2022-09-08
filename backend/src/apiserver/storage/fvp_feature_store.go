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

	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var fvpfeatureColumns = []string{
	"fvp_features.id",
	"fvp_features.fvpid",
	"fvp_features.name",
	"fvp_features.value_type",
	"fvp_features.tags",
}

type FvpFeatureStoreInterface interface {
	GetFeature(name string, fvpid string) (*model.FvpFeature, error)
	CreateFeature(tx *sql.Tx, feature *model.FvpFeature, fvpid string) (*model.FvpFeature, error)
	UpdateFeature(tx *sql.Tx, feature *model.FvpFeature, fvpid string) (*model.FvpFeature, error)
	DeleteFeature(tx *sql.Tx, name string, fvpid string) error
	DeleteFeatures(tx *sql.Tx, fvpid string) error
	ListFeatures(fvpid string) ([]*model.FvpFeature, error)
	SafeUpdateFeatures(tx *sql.Tx, features []*model.FvpFeature, fvpid string) error
}

type FvpFeatureStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *FvpFeatureStore) GetFeature(name string, fvpid string) (*model.FvpFeature, error) {
	sql, args, err := sq.
		Select(fvpfeatureColumns...).
		From("fvp_features").
		Where(sq.And{sq.Eq{"fvp_features.name": name}, sq.Eq{"fvp_features.fvid": fvpid}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get feature: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get feature: %v", err.Error())
	}
	defer r.Close()

	features, err := s.scanRows(r)
	if err != nil || len(features) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get feature: %v", err.Error())
	}
	if len(features) == 0 {
		return nil, util.NewResourceNotFoundError("Feature", fmt.Sprint(name))
	}

	return features[0], nil
}

func (s *FvpFeatureStore) CreateFeature(tx *sql.Tx, feature *model.FvpFeature, fvpid string) (*model.FvpFeature, error) {
	newFeature := *feature

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create feature id.")
	}
	newFeature.Id = id.String()
	newFeature.FVPId = fvpid

	sql, args, err := sq.
		Insert("fvp_features").
		SetMap(
			sq.Eq{
				"id":         newFeature.Id,
				"fvpid":      newFeature.FVPId,
				"name":       newFeature.Name,
				"value_type": newFeature.ValueType,
				"tags":       newFeature.Tags,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert feature: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to store feature.")
	}

	return &newFeature, nil
}

func (s *FvpFeatureStore) UpdateFeature(tx *sql.Tx, feature *model.FvpFeature, fvpid string) (*model.FvpFeature, error) {
	updatedFeature := *feature
	updatedFeature.FVPId = fvpid

	sql, args, err := sq.
		Update("fvp_features").
		SetMap(
			sq.Eq{
				"value_type": updatedFeature.ValueType,
				"tags":       updatedFeature.Tags,
			}).
		Where(sq.And{sq.Eq{"fvp_features.name": updatedFeature.Name}, sq.Eq{"fvp_features.fvpid": updatedFeature.FVPId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update feature: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update feature: %s", err.Error())
	}

	return &updatedFeature, nil
}

func (s *FvpFeatureStore) DeleteFeature(tx *sql.Tx, name string, fvpid string) error {
	sql, args, err := sq.
		Delete("fvp_features").
		Where(sq.And{sq.Eq{"fvp_features.name": name}, sq.Eq{"fvp_features.fvpid": fvpid}}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete feature: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete feature: %v", err.Error())
	}

	return nil
}

func (s *FvpFeatureStore) DeleteFeatures(tx *sql.Tx, fvpid string) error {
	sql, args, err := sq.
		Delete("fvp_features").
		Where(sq.Eq{"fvp_features.fvpid": fvpid}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete features: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete features: %v", err.Error())
	}

	return nil
}

func (s *FvpFeatureStore) ListFeatures(fvpid string) ([]*model.FvpFeature, error) {
	sql, args, err := sq.
		Select(fvpfeatureColumns...).
		From("fvp_features").
		Where(sq.Eq{"fvp_features.fvpid": fvpid}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list features: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list features: %v", err.Error())
	}
	defer r.Close()

	features, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list features: %v", err.Error())
	}

	return features, nil
}

func (s *FvpFeatureStore) SafeUpdateFeatures(tx *sql.Tx, features []*model.FvpFeature, fvpid string) error {
	prevFeatures, err := s.ListFeatures(fvpid)
	if err != nil {
		return err
	}

	featuresToBeDeleted := make(map[string]int)
	// Assume all features will be deleted
	for _, f := range prevFeatures {
		featuresToBeDeleted[f.Name] = 0
	}

	for _, f := range features {
		if _, ok := featuresToBeDeleted[f.Name]; ok {
			// Remove feature from featuresToBeDeleted and update it
			delete(featuresToBeDeleted, f.Name)
			_, err := s.UpdateFeature(tx, f, fvpid)
			if err != nil {
				return err
			}
		} else {
			_, err := s.CreateFeature(tx, f, fvpid)
			if err != nil {
				return err
			}
		}
	}

	// Delete remaining features of prevFeatures
	for name := range featuresToBeDeleted {
		err := s.DeleteFeature(tx, name, fvpid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *FvpFeatureStore) scanRows(rows *sql.Rows) ([]*model.FvpFeature, error) {
	var features []*model.FvpFeature
	for rows.Next() {
		var id, name, fvpid string
		var value_type sql.NullInt64
		var tags []byte

		if err := rows.Scan(
			&id,
			&fvpid,
			&name,
			&value_type,
			&tags); err != nil {
			return nil, err
		}

		features = append(features, &model.FvpFeature{
			Id:        id,
			FVPId:     fvpid,
			Name:      name,
			ValueType: value_type.Int64,
			Tags:      tags,
		})
	}

	return features, nil
}

func NewFvpFeatureStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *FvpFeatureStore {
	return &FvpFeatureStore{db: db, time: time, uuid: uuid}
}
