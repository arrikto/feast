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
var dataSourceColumns = []string{
	"data_sources.id",
	"data_sources.project_id",
	"data_sources.name",
	"data_sources.description",
	"data_sources.tags",
	"data_sources.owner",
	"data_sources.type",
	"data_sources.field_mapping",
	"data_sources.timestamp_field",
	"data_sources.date_partition_column",
	"data_sources.created_timestamp_column",
	"data_sources.class_type",
	"data_sources.batch_source",
	"data_sources.options",
}

type DataSourceStoreInterface interface {
	GetDataSource(name string, projectId string) (*model.DataSource, error)
	CreateDataSource(*model.DataSource) (*model.DataSource, error)
	UpdateDataSource(*model.DataSource) (*model.DataSource, error)
	DeleteDataSource(name string, projectId string) error
	ListDataSources(projectId string) ([]*model.DataSource, error)
}

type DataSourceStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *DataSourceStore) GetDataSource(name string, projectId string) (*model.DataSource, error) {
	sql, args, err := sq.
		Select(dataSourceColumns...).
		From("data_sources").
		Where(sq.And{sq.Eq{"data_sources.name": name}, sq.Eq{"data_sources.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get data source: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get data source: %v", err.Error())
	}
	defer r.Close()

	data_sources, err := s.scanRows(r)
	if err != nil || len(data_sources) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get data source: %v", err.Error())
	}
	if len(data_sources) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.DataSource), fmt.Sprint(name))
	}

	return data_sources[0], nil
}

func (s *DataSourceStore) CreateDataSource(ds *model.DataSource) (*model.DataSource, error) {
	newDataSource := *ds

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create data source id.")
	}
	newDataSource.Id = id.String()

	sql, args, err := sq.
		Insert("data_sources").
		SetMap(
			sq.Eq{
				"id":                       newDataSource.Id,
				"project_id":               newDataSource.ProjectId,
				"name":                     newDataSource.Name,
				"description":              newDataSource.Description,
				"tags":                     newDataSource.Tags,
				"owner":                    newDataSource.Owner,
				"type":                     newDataSource.Type,
				"field_mapping":            newDataSource.FieldMapping,
				"timestamp_field":          newDataSource.TimestampField,
				"date_partition_column":    newDataSource.DatePartitionCol,
				"created_timestamp_column": newDataSource.CreatedTimestampCol,
				"class_type":               newDataSource.ClassType,
				"batch_source":             newDataSource.BatchSource,
				"options":                  newDataSource.Options,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert data source: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new data source. The name %v already exists. Please specify a new name.", ds.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add data source to data_sources table: %v",
			err.Error())
	}

	return &newDataSource, nil
}

func (s *DataSourceStore) UpdateDataSource(ds *model.DataSource) (*model.DataSource, error) {
	updatedDataSource := *ds

	sql, args, err := sq.
		Update("data_sources").
		SetMap(
			sq.Eq{
				"description":              updatedDataSource.Description,
				"tags":                     updatedDataSource.Tags,
				"owner":                    updatedDataSource.Owner,
				"type":                     updatedDataSource.Type,
				"field_mapping":            updatedDataSource.FieldMapping,
				"timestamp_field":          updatedDataSource.TimestampField,
				"date_partition_column":    updatedDataSource.DatePartitionCol,
				"created_timestamp_column": updatedDataSource.CreatedTimestampCol,
				"class_type":               updatedDataSource.ClassType,
				"batch_source":             updatedDataSource.BatchSource,
				"options":                  updatedDataSource.Options,
			}).
		Where(sq.And{sq.Eq{"data_sources.name": updatedDataSource.Name}, sq.Eq{"data_sources.project_id": updatedDataSource.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update data source: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update data source: %s", err.Error())
	}

	return &updatedDataSource, nil
}

func (s *DataSourceStore) DeleteDataSource(name string, projectId string) error {
	sql, args, err := sq.
		Delete("data_sources").
		Where(sq.And{sq.Eq{"data_sources.name": name}, sq.Eq{"data_sources.project_id": projectId}}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete data source: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete data source: %v", err.Error())
	}

	return nil
}

func (s *DataSourceStore) ListDataSources(projectId string) ([]*model.DataSource, error) {
	sql, args, err := sq.
		Select(dataSourceColumns...).
		From("data_sources").
		Where(sq.Eq{"data_sources.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list data sources: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list data sources: %v", err.Error())
	}
	defer r.Close()

	data_sources, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list data sources: %v", err.Error())
	}

	return data_sources, nil
}

func (s *DataSourceStore) scanRows(rows *sql.Rows) ([]*model.DataSource, error) {
	var data_sources []*model.DataSource
	for rows.Next() {
		var id, name, project_id string
		var description, owner, timestamp_field, date_partition_col sql.NullString
		var created_timestamp_col, class_type sql.NullString
		var type_ sql.NullInt64
		var tags, field_mapping, batch_source, options []byte

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&description,
			&tags,
			&owner,
			&type_,
			&field_mapping,
			&timestamp_field,
			&date_partition_col,
			&created_timestamp_col,
			&class_type,
			&batch_source,
			&options); err != nil {
			return nil, err
		}

		data_sources = append(data_sources, &model.DataSource{
			Id:                  id,
			ProjectId:           project_id,
			Name:                name,
			Description:         description.String,
			Tags:                tags,
			Owner:               owner.String,
			Type:                type_.Int64,
			FieldMapping:        field_mapping,
			TimestampField:      timestamp_field.String,
			DatePartitionCol:    date_partition_col.String,
			CreatedTimestampCol: created_timestamp_col.String,
			ClassType:           class_type.String,
			BatchSource:         batch_source,
			Options:             options,
		})
	}

	return data_sources, nil
}

func NewDataSourceStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *DataSourceStore {
	return &DataSourceStore{db: db, time: time, uuid: uuid}
}
