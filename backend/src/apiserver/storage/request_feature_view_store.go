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
var requestFVColumns = []string{
	"request_feature_views.id",
	"request_feature_views.project_id",
	"request_feature_views.name",
	"request_feature_views.data_source",
	"request_feature_views.description",
	"request_feature_views.tags",
	"request_feature_views.owner",
}

type RequestFeatureViewStoreInterface interface {
	GetRequestFeatureView(name string, projectId string) (*model.RequestFeatureView, error)
	CreateRequestFeatureView(*model.RequestFeatureView) (*model.RequestFeatureView, error)
	UpdateRequestFeatureView(*model.RequestFeatureView) (*model.RequestFeatureView, error)
	DeleteRequestFeatureView(name string, projectId string) error
	ListRequestFeatureViews(projectId string) ([]*model.RequestFeatureView, error)
}

type RequestFeatureViewStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *RequestFeatureViewStore) GetRequestFeatureView(name string, projectId string) (*model.RequestFeatureView, error) {
	sql, args, err := sq.
		Select(requestFVColumns...).
		From("request_feature_views").
		Where(sq.And{sq.Eq{"request_feature_views.name": name}, sq.Eq{"request_feature_views.project_id": projectId}}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get request feature view: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get request feature view: %v", err.Error())
	}
	defer r.Close()

	request_feature_views, err := s.scanRows(r)
	if err != nil || len(request_feature_views) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get request feature view: %v", err.Error())
	}
	if len(request_feature_views) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.RequestFeatureView), fmt.Sprint(name))
	}

	return request_feature_views[0], nil
}

func (s *RequestFeatureViewStore) CreateRequestFeatureView(rfv *model.RequestFeatureView) (*model.RequestFeatureView, error) {
	newRequestFeatureView := *rfv

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create request feature view id.")
	}
	newRequestFeatureView.Id = id.String()

	sql, args, err := sq.
		Insert("request_feature_views").
		SetMap(
			sq.Eq{
				"id":          newRequestFeatureView.Id,
				"project_id":  newRequestFeatureView.ProjectId,
				"name":        newRequestFeatureView.Name,
				"data_source": newRequestFeatureView.DataSource,
				"description": newRequestFeatureView.Description,
				"tags":        newRequestFeatureView.Tags,
				"owner":       newRequestFeatureView.Owner,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert request feature view: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new request feature view. The name %v already exists. Please specify a new name.", rfv.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add request feature view to request_feature_views table: %v",
			err.Error())
	}

	return &newRequestFeatureView, nil
}

func (s *RequestFeatureViewStore) UpdateRequestFeatureView(rfv *model.RequestFeatureView) (*model.RequestFeatureView, error) {
	updatedRequestFeatureView := *rfv

	sql, args, err := sq.
		Update("request_feature_views").
		SetMap(
			sq.Eq{
				"data_source": updatedRequestFeatureView.DataSource,
				"description": updatedRequestFeatureView.Description,
				"tags":        updatedRequestFeatureView.Tags,
				"owner":       updatedRequestFeatureView.Owner,
			}).
		Where(sq.And{sq.Eq{"request_feature_views.name": updatedRequestFeatureView.Name}, sq.Eq{"request_feature_views.project_id": updatedRequestFeatureView.ProjectId}}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update request feature view: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update request feature view: %s", err.Error())
	}

	return &updatedRequestFeatureView, nil
}

func (s *RequestFeatureViewStore) DeleteRequestFeatureView(name string, projectId string) error {
	sql, args, err := sq.
		Delete("request_feature_views").
		Where(sq.And{sq.Eq{"request_feature_views.name": name}, sq.Eq{"request_feature_views.project_id": projectId}}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete request feature view: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete request feature view: %v", err.Error())
	}

	return nil
}

func (s *RequestFeatureViewStore) ListRequestFeatureViews(projectId string) ([]*model.RequestFeatureView, error) {
	sql, args, err := sq.
		Select(requestFVColumns...).
		From("request_feature_views").
		Where(sq.Eq{"request_feature_views.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list request feature views: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list request feature views: %v", err.Error())
	}
	defer r.Close()

	request_feature_views, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list request feature views: %v", err.Error())
	}

	return request_feature_views, nil
}

func (s *RequestFeatureViewStore) scanRows(rows *sql.Rows) ([]*model.RequestFeatureView, error) {
	var request_feature_views []*model.RequestFeatureView
	for rows.Next() {
		var id, name, project_id string
		var description, owner sql.NullString
		var tags, data_source []byte

		if err := rows.Scan(
			&id,
			&project_id,
			&name,
			&data_source,
			&description,
			&tags,
			&owner); err != nil {
			return nil, err
		}

		request_feature_views = append(request_feature_views, &model.RequestFeatureView{
			Id:          id,
			ProjectId:   project_id,
			Name:        name,
			DataSource:  data_source,
			Description: description.String,
			Tags:        tags,
			Owner:       owner.String,
		})
	}

	return request_feature_views, nil
}

func NewRequestFeatureViewStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *RequestFeatureViewStore {
	return &RequestFeatureViewStore{db: db, time: time, uuid: uuid}
}
