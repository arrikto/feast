package storage

import (
	"database/sql"
	"fmt"

	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var featureColumns = []string{
	"features.id",
	"features.fvid",
	"features.name",
	"features.value_type",
	"features.tags",
}

type FeatureStoreInterface interface {
	GetFeature(name string, fvid string) (*model.Feature, error)
	CreateFeature(tx *sql.Tx, feature *model.Feature, fvid string) (*model.Feature, error)
	UpdateFeature(tx *sql.Tx, feature *model.Feature, fvid string) (*model.Feature, error)
	DeleteFeature(tx *sql.Tx, name string, fvid string) error
	ListFeatures(fvid string) ([]*model.Feature, error)
	SafeUpdateFeatures(tx *sql.Tx, features []*model.Feature, fvid string) error
}

type FeatureStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *FeatureStore) GetFeature(name string, fvid string) (*model.Feature, error) {
	sql, args, err := sq.
		Select(featureColumns...).
		From("features").
		Where(sq.And{sq.Eq{"features.name": name}, sq.Eq{"features.fvid": fvid}}).
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

func (s *FeatureStore) CreateFeature(tx *sql.Tx, feature *model.Feature, fvid string) (*model.Feature, error) {
	newFeature := *feature

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create feature id.")
	}
	newFeature.Id = id.String()
	newFeature.FVId = fvid

	sql, args, err := sq.
		Insert("features").
		SetMap(
			sq.Eq{
				"id":         newFeature.Id,
				"fvid":       newFeature.FVId,
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
		return nil, util.NewInternalServerError(err, "Failed to add feature to features table: %v", err.Error())
	}

	return &newFeature, nil
}

func (s *FeatureStore) UpdateFeature(tx *sql.Tx, feature *model.Feature, fvid string) (*model.Feature, error) {
	updatedFeature := *feature
	updatedFeature.FVId = fvid

	sql, args, err := sq.
		Update("features").
		SetMap(
			sq.Eq{
				"value_type": updatedFeature.ValueType,
				"tags":       updatedFeature.Tags,
			}).
		Where(sq.And{sq.Eq{"features.name": updatedFeature.Name}, sq.Eq{"features.fvid": updatedFeature.FVId}}).
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

func (s *FeatureStore) DeleteFeature(tx *sql.Tx, name string, fvid string) error {
	sql, args, err := sq.
		Delete("features").
		Where(sq.And{sq.Eq{"features.name": name}, sq.Eq{"features.fvid": fvid}}).
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

func (s *FeatureStore) ListFeatures(fvid string) ([]*model.Feature, error) {
	sql, args, err := sq.
		Select(featureColumns...).
		From("features").
		Where(sq.Eq{"features.fvid": fvid}).
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

func (s *FeatureStore) SafeUpdateFeatures(tx *sql.Tx, features []*model.Feature, fvid string) error {
	prevFeatures, err := s.ListFeatures(fvid)
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
			updFeature, err := s.UpdateFeature(tx, f, fvid)
			if err != nil {
				return err
			}
			f.Id = updFeature.Id
			f.FVId = updFeature.FVId
		} else {
			newFeature, err := s.CreateFeature(tx, f, fvid)
			if err != nil {
				return err
			}
			f.Id = newFeature.Id
			f.FVId = newFeature.FVId
		}
	}

	// Delete remaining features of prevFeatures
	for name := range featuresToBeDeleted {
		err := s.DeleteFeature(tx, name, fvid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *FeatureStore) scanRows(rows *sql.Rows) ([]*model.Feature, error) {
	var features []*model.Feature
	for rows.Next() {
		var id, name, fvid string
		var value_type sql.NullInt64
		var tags []byte

		if err := rows.Scan(
			&id,
			&fvid,
			&name,
			&value_type,
			&tags); err != nil {
			return nil, err
		}

		features = append(features, &model.Feature{
			Id:        id,
			FVId:      fvid,
			Name:      name,
			ValueType: value_type.Int64,
			Tags:      tags,
		})
	}

	return features, nil
}

func NewFeatureStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *FeatureStore {
	return &FeatureStore{db: db, time: time, uuid: uuid}
}
