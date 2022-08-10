package storage

import (
	"database/sql"

	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var infraObjectColumns = []string{
	"infra_objects.id",
	"infra_objects.project_id",
	"infra_objects.class_type",
	"infra_objects.object",
}

type InfraObjectStoreInterface interface {
	CreateInfraObject(*model.InfraObject) (*model.InfraObject, error)
	DeleteInfraObjects(projectId string) error
	ListInfraObjects(projectId string) ([]*model.InfraObject, error)
}

type InfraObjectStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *InfraObjectStore) CreateInfraObject(io *model.InfraObject) (*model.InfraObject, error) {
	newInfraObject := *io

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create infra object id.")
	}
	newInfraObject.Id = id.String()

	sql, args, err := sq.
		Insert("infra_objects").
		SetMap(
			sq.Eq{
				"id":         newInfraObject.Id,
				"project_id": newInfraObject.ProjectId,
				"class_type": newInfraObject.ClassType,
				"object":     newInfraObject.Object,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert infra object: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to add infra object to infra_objects table: %v",
			err.Error())
	}

	return &newInfraObject, nil
}

func (s *InfraObjectStore) DeleteInfraObjects(projectId string) error {
	sql, args, err := sq.
		Delete("infra_objects").
		Where(sq.Eq{"infra_objects.project_id": projectId}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete infra objects: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete infra objects: %v", err.Error())
	}

	return nil
}

func (s *InfraObjectStore) ListInfraObjects(projectId string) ([]*model.InfraObject, error) {
	sql, args, err := sq.
		Select(infraObjectColumns...).
		From("infra_objects").
		Where(sq.Eq{"infra_objects.project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list infra objects: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list infra objects: %v", err.Error())
	}
	defer r.Close()

	infra_objects, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list infra objects: %v", err.Error())
	}

	return infra_objects, nil
}

func (s *InfraObjectStore) scanRows(rows *sql.Rows) ([]*model.InfraObject, error) {
	var infra_objects []*model.InfraObject
	for rows.Next() {
		var id, project_id string
		var class_type sql.NullString
		var object []byte

		if err := rows.Scan(
			&id,
			&project_id,
			&class_type,
			&object); err != nil {
			return nil, err
		}

		infra_objects = append(infra_objects, &model.InfraObject{
			Id:        id,
			ProjectId: project_id,
			ClassType: class_type.String,
			Object:    object,
		})
	}

	return infra_objects, nil
}

func NewInfraObjectStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *InfraObjectStore {
	return &InfraObjectStore{db: db, time: time, uuid: uuid}
}
