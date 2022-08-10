package storage

import (
	"database/sql"

	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var fvpColumns = []string{
	"feature_view_projections.id",
	"feature_view_projections.fsid",
	"feature_view_projections.feature_view_name",
	"feature_view_projections.feature_view_name_alias",
	"feature_view_projections.join_key_map",
}

type FVPStoreInterface interface {
	CreateFVP(tx *sql.Tx, fvp *model.FeatureViewProjection, fsid string) (*model.FeatureViewProjection, error)
	DeleteFVPs(tx *sql.Tx, fsid string) error
	ListFVPs(fsid string) ([]*model.FeatureViewProjection, error)
}

type FVPStore struct {
	db           *DB
	featureStore *FvpFeatureStore
	time         util.TimeInterface
	uuid         util.UUIDGeneratorInterface
}

func (s *FVPStore) CreateFVP(tx *sql.Tx, fvp *model.FeatureViewProjection, fsid string) (*model.FeatureViewProjection, error) {
	newFVP := *fvp

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create feature view projection id.")
	}
	newFVP.Id = id.String()
	newFVP.FSId = fsid

	sql, args, err := sq.
		Insert("feature_view_projections").
		SetMap(
			sq.Eq{
				"id":                      newFVP.Id,
				"fsid":                    newFVP.FSId,
				"feature_view_name":       newFVP.FVName,
				"feature_view_name_alias": newFVP.FVNameAlias,
				"join_key_map":            newFVP.JoinKeyMap,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert feature view projection: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to add feature view projection to feature_view_projections table: %v",
			err.Error())
	}

	for _, feature := range newFVP.FVPFeatures {
		newFeature, err := s.featureStore.CreateFeature(tx, feature, newFVP.Id)
		if err != nil {
			return nil, util.NewInternalServerError(err, "Failed to store feature %v for feature view projection %v ", feature.Name, fvp.FVName)
		}
		feature.Id = newFeature.Id
		feature.FVPId = newFeature.FVPId
	}

	return &newFVP, nil
}

func (s *FVPStore) DeleteFVPs(tx *sql.Tx, fsid string) error {
	fvps, err := s.ListFVPs(fsid)
	if err != nil {
		return err
	}

	for _, fvp := range fvps {
		err := s.featureStore.DeleteFeatures(tx, fvp.Id)
		if err != nil {
			return util.NewInternalServerError(err, "Failed to delete features of feature view projection: %v", err.Error())
		}
	}

	sql, args, err := sq.
		Delete("feature_view_projections").
		Where(sq.Eq{"feature_view_projections.fsid": fsid}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete feature view projections: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete feature view projections: %v", err.Error())
	}

	return nil
}

func (s *FVPStore) ListFVPs(fsid string) ([]*model.FeatureViewProjection, error) {
	sql, args, err := sq.
		Select(fvpColumns...).
		From("feature_view_projections").
		Where(sq.Eq{"feature_view_projections.fsid": fsid}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list feature view projections: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list feature view projections: %v", err.Error())
	}
	defer r.Close()

	fvps, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list feature view projections: %v", err.Error())
	}

	for _, fvp := range fvps {
		features, err := s.featureStore.ListFeatures(fvp.Id)
		if err != nil {
			return nil, util.NewInternalServerError(err, "Failed to get features: %v", err.Error())
		}
		fvp.FVPFeatures = append(fvp.FVPFeatures, features...)
	}

	return fvps, nil
}

func (s *FVPStore) scanRows(rows *sql.Rows) ([]*model.FeatureViewProjection, error) {
	var fvps []*model.FeatureViewProjection
	for rows.Next() {
		var id, fsid string
		var feature_view_name, feature_view_name_alias sql.NullString
		var join_key_map []byte

		if err := rows.Scan(
			&id,
			&fsid,
			&feature_view_name,
			&feature_view_name_alias,
			&join_key_map); err != nil {
			return nil, err
		}

		fvps = append(fvps, &model.FeatureViewProjection{
			Id:          id,
			FSId:        fsid,
			FVName:      feature_view_name.String,
			FVNameAlias: feature_view_name_alias.String,
			JoinKeyMap:  join_key_map,
		})
	}

	return fvps, nil
}

func NewFVPStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *FVPStore {
	featureStore := NewFvpFeatureStore(db, time, uuid)

	return &FVPStore{db: db, featureStore: featureStore, time: time, uuid: uuid}
}
