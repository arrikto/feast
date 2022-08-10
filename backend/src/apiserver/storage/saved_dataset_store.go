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
var savedDatasetColumns = []string{
	"saved_datasets.id",
	"saved_datasets.project_id",
	"saved_datasets.name",
	"saved_datasets.features",
	"saved_datasets.join_keys",
	"saved_datasets.full_feature_names",
	"saved_datasets.storage",
	"saved_datasets.feature_service_name",
	"saved_datasets.tags",
	"saved_datasets.created_timestamp",
	"saved_datasets.last_updated_timestamp",
	"saved_datasets.min_event_timestamp",
	"saved_datasets.max_event_timestamp",
}

type SavedDatasetStoreInterface interface {
	GetSavedDataset(name string, projectId string) (*model.SavedDataset, error)
	CreateSavedDataset(*model.SavedDataset) (*model.SavedDataset, error)
	UpdateSavedDataset(*model.SavedDataset) (*model.SavedDataset, error)
	DeleteSavedDataset(name string, projectId string) error
	ListSavedDatasets(projectId string) ([]*model.SavedDataset, error)
}

type SavedDatasetStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *SavedDatasetStore) GetSavedDataset(name string, projectId string) (*model.SavedDataset, error) {
	sql, args, err := sq.
		Select(savedDatasetColumns...).
		From("saved_datasets").
		Where(sq.And{sq.Eq{"saved_datasets.name": name}, sq.Eq{"saved_datasets.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get saved dataset: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get saved dataset: %v", err.Error())
	}
	defer r.Close()

	saved_datasets, err := s.scanRows(r)
	if err != nil || len(saved_datasets) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get saved dataset: %v", err.Error())
	}
	if len(saved_datasets) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.SavedDataset), fmt.Sprint(name))
	}

	return saved_datasets[0], nil
}

func (s *SavedDatasetStore) CreateSavedDataset(sd *model.SavedDataset) (*model.SavedDataset, error) {
	newSavedDataset := *sd

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create saved dataset id.")
	}
	newSavedDataset.Id = id.String()

	sql, args, err := sq.
		Insert("saved_datasets").
		SetMap(
			sq.Eq{
				"id":                     newSavedDataset.Id,
				"project_id":             newSavedDataset.ProjectId,
				"name":                   newSavedDataset.Name,
				"features":               newSavedDataset.Features,
				"join_keys":              newSavedDataset.JoinKeys,
				"full_feature_names":     newSavedDataset.FullFeatureNames,
				"storage":                newSavedDataset.Storage,
				"feature_service_name":   newSavedDataset.FeatureServiceName,
				"tags":                   newSavedDataset.Tags,
				"created_timestamp":      newSavedDataset.CreatedTimestamp,
				"last_updated_timestamp": newSavedDataset.LastUpdatedTimestamp,
				"min_event_timestamp":    newSavedDataset.MinEventTimestamp,
				"max_event_timestamp":    newSavedDataset.MaxEventTimestamp,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert saved dataset: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new saved dataset. The name %v already exists. Please specify a new name.", sd.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add saved dataset to saved_datasets table: %v",
			err.Error())
	}

	return &newSavedDataset, nil
}

func (s *SavedDatasetStore) UpdateSavedDataset(sd *model.SavedDataset) (*model.SavedDataset, error) {
	updatedSavedDataset := *sd

	sql, args, err := sq.
		Update("saved_datasets").
		SetMap(
			sq.Eq{
				"features":               updatedSavedDataset.Features,
				"join_keys":              updatedSavedDataset.JoinKeys,
				"full_feature_names":     updatedSavedDataset.FullFeatureNames,
				"storage":                updatedSavedDataset.Storage,
				"feature_service_name":   updatedSavedDataset.FeatureServiceName,
				"tags":                   updatedSavedDataset.Tags,
				"created_timestamp":      updatedSavedDataset.CreatedTimestamp,
				"last_updated_timestamp": updatedSavedDataset.LastUpdatedTimestamp,
				"min_event_timestamp":    updatedSavedDataset.MinEventTimestamp,
				"max_event_timestamp":    updatedSavedDataset.MaxEventTimestamp,
			}).
		Where(sq.And{sq.Eq{"saved_datasets.name": updatedSavedDataset.Name}, sq.Eq{"saved_datasets.project_id": updatedSavedDataset.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update saved dataset: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update saved dataset: %s", err.Error())
	}

	return &updatedSavedDataset, nil
}

func (s *SavedDatasetStore) DeleteSavedDataset(name string, projectId string) error {
	sql, args, err := sq.
		Delete("saved_datasets").
		Where(sq.And{sq.Eq{"saved_datasets.name": name}, sq.Eq{"saved_datasets.project_id": projectId}}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete saved dataset: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete saved dataset: %v", err.Error())
	}

	return nil
}

func (s *SavedDatasetStore) ListSavedDatasets(projectId string) ([]*model.SavedDataset, error) {
	sql, args, err := sq.
		Select(savedDatasetColumns...).
		From("saved_datasets").
		Where(sq.Eq{"saved_datasets.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list saved datasets: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list saved datasets: %v", err.Error())
	}
	defer r.Close()

	saved_datasets, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list saved datasets: %v", err.Error())
	}

	return saved_datasets, nil
}

func (s *SavedDatasetStore) scanRows(rows *sql.Rows) ([]*model.SavedDataset, error) {
	var saved_datasets []*model.SavedDataset
	for rows.Next() {
		var id, name, project_id string
		var feature_service_name sql.NullString
		var full_feature_names sql.NullBool
		var tags, features, join_keys, storage []byte
		var created_timestamp, last_updated_timestamp, min_event_timestamp, max_event_timestamp time.Time

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&features,
			&join_keys,
			&full_feature_names,
			&storage,
			&feature_service_name,
			&tags,
			&created_timestamp,
			&last_updated_timestamp,
			&min_event_timestamp,
			&max_event_timestamp); err != nil {
			return nil, err
		}

		saved_datasets = append(saved_datasets, &model.SavedDataset{
			Id:                   id,
			ProjectId:            project_id,
			Name:                 name,
			Features:             features,
			JoinKeys:             join_keys,
			FullFeatureNames:     full_feature_names.Bool,
			Storage:              storage,
			FeatureServiceName:   feature_service_name.String,
			Tags:                 tags,
			CreatedTimestamp:     created_timestamp,
			LastUpdatedTimestamp: last_updated_timestamp,
			MinEventTimestamp:    min_event_timestamp,
			MaxEventTimestamp:    max_event_timestamp,
		})
	}

	return saved_datasets, nil
}

func NewSavedDatasetStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *SavedDatasetStore {
	return &SavedDatasetStore{db: db, time: time, uuid: uuid}
}
