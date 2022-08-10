package storage

import (
	"database/sql"
	"fmt"

	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var odfeatureColumns = []string{
	"on_demand_features.id",
	"on_demand_features.odfvid",
	"on_demand_features.name",
	"on_demand_features.value_type",
	"on_demand_features.tags",
}

type OnDemandFeatureStoreInterface interface {
	GetFeature(name string, odfvid string) (*model.OnDemandFeature, error)
	CreateFeature(tx *sql.Tx, feature *model.OnDemandFeature, odfvid string) (*model.OnDemandFeature, error)
	UpdateFeature(tx *sql.Tx, feature *model.OnDemandFeature, odfvid string) (*model.OnDemandFeature, error)
	DeleteFeature(tx *sql.Tx, name string, odfvid string) error
	ListFeatures(odfvid string) ([]*model.OnDemandFeature, error)
	SafeUpdateFeatures(tx *sql.Tx, features []*model.OnDemandFeature, odfvid string) error
}

type OnDemandFeatureStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *OnDemandFeatureStore) GetFeature(name string, odfvid string) (*model.OnDemandFeature, error) {
	sql, args, err := sq.
		Select(odfeatureColumns...).
		From("on_demand_features").
		Where(sq.And{sq.Eq{"on_demand_features.name": name}, sq.Eq{"on_demand_features.fvid": odfvid}}).
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

func (s *OnDemandFeatureStore) CreateFeature(tx *sql.Tx, feature *model.OnDemandFeature, odfvid string) (*model.OnDemandFeature, error) {
	newFeature := *feature

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create feature id.")
	}
	newFeature.Id = id.String()
	newFeature.ODFVId = odfvid

	sql, args, err := sq.
		Insert("on_demand_features").
		SetMap(
			sq.Eq{
				"id":         newFeature.Id,
				"odfvid":     newFeature.ODFVId,
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

func (s *OnDemandFeatureStore) UpdateFeature(tx *sql.Tx, feature *model.OnDemandFeature, odfvid string) (*model.OnDemandFeature, error) {
	updatedFeature := *feature
	updatedFeature.ODFVId = odfvid

	sql, args, err := sq.
		Update("on_demand_features").
		SetMap(
			sq.Eq{
				"value_type": updatedFeature.ValueType,
				"tags":       updatedFeature.Tags,
			}).
		Where(sq.And{sq.Eq{"on_demand_features.name": updatedFeature.Name}, sq.Eq{"on_demand_features.odfvid": updatedFeature.ODFVId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update feature: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update the feature: %s", err.Error())
	}

	return &updatedFeature, nil
}

func (s *OnDemandFeatureStore) DeleteFeature(tx *sql.Tx, name string, odfvid string) error {
	sql, args, err := sq.
		Delete("on_demand_features").
		Where(sq.And{sq.Eq{"on_demand_features.name": name}, sq.Eq{"on_demand_features.odfvid": odfvid}}).
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

func (s *OnDemandFeatureStore) ListFeatures(odfvid string) ([]*model.OnDemandFeature, error) {
	sql, args, err := sq.
		Select(odfeatureColumns...).
		From("on_demand_features").
		Where(sq.Eq{"on_demand_features.odfvid": odfvid}).
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

func (s *OnDemandFeatureStore) SafeUpdateFeatures(tx *sql.Tx, features []*model.OnDemandFeature, odfvid string) error {
	prevFeatures, err := s.ListFeatures(odfvid)
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
			_, err := s.UpdateFeature(tx, f, odfvid)
			if err != nil {
				return err
			}
		} else {
			_, err := s.CreateFeature(tx, f, odfvid)
			if err != nil {
				return err
			}
		}
	}

	// Delete remaining features of prevFeatures
	for name := range featuresToBeDeleted {
		err := s.DeleteFeature(tx, name, odfvid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *OnDemandFeatureStore) scanRows(rows *sql.Rows) ([]*model.OnDemandFeature, error) {
	var features []*model.OnDemandFeature
	for rows.Next() {
		var id, name, odfvid string
		var value_type sql.NullInt64
		var tags []byte

		if err := rows.Scan(
			&id,
			&odfvid,
			&name,
			&value_type,
			&tags); err != nil {
			return nil, err
		}

		features = append(features, &model.OnDemandFeature{
			Id:        id,
			ODFVId:    odfvid,
			Name:      name,
			ValueType: value_type.Int64,
			Tags:      tags,
		})
	}

	return features, nil
}

func NewOnDemandFeatureStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *OnDemandFeatureStore {
	return &OnDemandFeatureStore{db: db, time: time, uuid: uuid}
}
