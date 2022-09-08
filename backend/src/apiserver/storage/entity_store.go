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
var entityColumns = []string{
	"entities.id",
	"entities.project_id",
	"entities.name",
	"entities.value_type",
	"entities.description",
	"entities.join_key",
	"entities.tags",
	"entities.owner",
	"entities.created_timestamp",
	"entities.last_updated_timestamp",
}

type EntityStoreInterface interface {
	GetEntity(name string, projectId string) (*model.Entity, error)
	CreateEntity(*model.Entity) (*model.Entity, error)
	UpdateEntity(*model.Entity) (*model.Entity, error)
	DeleteEntity(name string, projectId string) error
	ListEntities(projectId string) ([]*model.Entity, error)
}

type EntityStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *EntityStore) GetEntity(name string, projectId string) (*model.Entity, error) {
	sql, args, err := sq.
		Select(entityColumns...).
		From("entities").
		Where(sq.And{sq.Eq{"entities.name": name}, sq.Eq{"entities.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get entity: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get entity: %v", err.Error())
	}
	defer r.Close()

	entities, err := s.scanRows(r)
	if err != nil || len(entities) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get entity: %v", err.Error())
	}
	if len(entities) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.Entity), fmt.Sprint(name))
	}

	return entities[0], nil
}

func (s *EntityStore) CreateEntity(e *model.Entity) (*model.Entity, error) {
	newEntity := *e

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create entity id.")
	}
	newEntity.Id = id.String()

	sql, args, err := sq.
		Insert("entities").
		SetMap(
			sq.Eq{
				"id":                     newEntity.Id,
				"project_id":             newEntity.ProjectId,
				"name":                   newEntity.Name,
				"value_type":             newEntity.ValueType,
				"description":            newEntity.Description,
				"join_key":               newEntity.JoinKey,
				"tags":                   newEntity.Tags,
				"owner":                  newEntity.Owner,
				"created_timestamp":      newEntity.CreatedTimestamp,
				"last_updated_timestamp": newEntity.LastUpdatedTimestamp,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert entity: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new entity. The name %v already exists. Please specify a new name.", e.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add entity to entities table: %v",
			err.Error())
	}

	return &newEntity, nil
}

func (s *EntityStore) UpdateEntity(e *model.Entity) (*model.Entity, error) {
	updatedEntity := *e

	sql, args, err := sq.
		Update("entities").
		SetMap(
			sq.Eq{
				"value_type":             updatedEntity.ValueType,
				"description":            updatedEntity.Description,
				"join_key":               updatedEntity.JoinKey,
				"tags":                   updatedEntity.Tags,
				"owner":                  updatedEntity.Owner,
				"created_timestamp":      updatedEntity.CreatedTimestamp,
				"last_updated_timestamp": updatedEntity.LastUpdatedTimestamp,
			}).
		Where(sq.And{sq.Eq{"entities.name": updatedEntity.Name}, sq.Eq{"entities.project_id": updatedEntity.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update entity: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update entity: %s", err.Error())
	}

	return &updatedEntity, nil
}

func (s *EntityStore) DeleteEntity(name string, projectId string) error {
	sql, args, err := sq.
		Delete("entities").
		Where(sq.And{sq.Eq{"entities.name": name}, sq.Eq{"entities.project_id": projectId}}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete entity: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete entity: %v", err.Error())
	}

	return nil
}

func (s *EntityStore) ListEntities(projectId string) ([]*model.Entity, error) {
	sql, args, err := sq.
		Select(entityColumns...).
		From("entities").
		Where(sq.Eq{"entities.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list entities: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list entities: %v", err.Error())
	}
	defer r.Close()

	entities, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list entities: %v", err.Error())
	}

	return entities, nil
}

func (s *EntityStore) scanRows(rows *sql.Rows) ([]*model.Entity, error) {
	var entities []*model.Entity
	for rows.Next() {
		var id, name, project_id string
		var description, join_key, owner sql.NullString
		var value_type sql.NullInt64
		var tags []byte
		var created_timestamp, last_updated_timestamp sql.NullTime

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&value_type,
			&description,
			&join_key,
			&tags,
			&owner,
			&created_timestamp,
			&last_updated_timestamp); err != nil {
			return nil, err
		}

		entities = append(entities, &model.Entity{
			Id:                   id,
			ProjectId:            project_id,
			Name:                 name,
			ValueType:            value_type.Int64,
			Description:          description.String,
			JoinKey:              join_key.String,
			Tags:                 tags,
			Owner:                owner.String,
			CreatedTimestamp:     created_timestamp.Time,
			LastUpdatedTimestamp: last_updated_timestamp.Time,
		})
	}

	return entities, nil
}

func NewEntityStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *EntityStore {
	return &EntityStore{db: db, time: time, uuid: uuid}
}
