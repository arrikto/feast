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
var projectColumns = []string{
	"projects.id",
	"projects.name",
	"projects.registry_schema_version",
	"projects.version_id",
	"projects.last_updated",
}

type ProjectStoreInterface interface {
	GetProject(project string) (*model.Project, error)
	CreateProject(*model.Project) (*model.Project, error)
	UpdateProject(*model.Project) (*model.Project, error)
	DeleteProject(project string) error
}

type ProjectStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *ProjectStore) GetProject(project string) (*model.Project, error) {
	sql, args, err := sq.
		Select(projectColumns...).
		From("projects").
		Where(sq.Eq{"projects.name": project}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get project: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get project: %v", err.Error())
	}
	defer r.Close()

	projects, err := s.scanRows(r)
	if err != nil || len(projects) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get project: %v", err.Error())
	}
	if len(projects) == 0 {
		return nil, util.NewResourceNotFoundError(string(common.Project), fmt.Sprint(project))
	}

	return projects[0], nil
}

func (s *ProjectStore) CreateProject(p *model.Project) (*model.Project, error) {
	newProject := *p

	id, err := s.uuid.NewRandom()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create project id.")
	}
	newProject.Id = id.String()

	sql, args, err := sq.
		Insert("projects").
		SetMap(
			sq.Eq{
				"id":                      newProject.Id,
				"name":                    newProject.Name,
				"registry_schema_version": newProject.RegistrySchemaVersion,
				"version_id":              newProject.VersionId,
				"last_updated":            newProject.LastUpdated,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert project: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewAlreadyExistError(
				"Failed to create a new project. The project %v already exists. Please specify a new project name.", p.Name)
		}
		return nil, util.NewInternalServerError(err, "Failed to add project to projects table: %v",
			err.Error())
	}

	return &newProject, nil
}

func (s *ProjectStore) UpdateProject(p *model.Project) (*model.Project, error) {
	updatedProject := *p

	sql, args, err := sq.
		Update("projects").
		SetMap(
			sq.Eq{
				"registry_schema_version": updatedProject.RegistrySchemaVersion,
				"version_id":              updatedProject.VersionId,
				"last_updated":            updatedProject.LastUpdated,
			}).
		Where(sq.Eq{"projects.name": updatedProject.Name}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to update project: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to update project: %s", err.Error())
	}

	return &updatedProject, nil
}

func (s *ProjectStore) DeleteProject(project string) error {
	sql, args, err := sq.
		Delete("projects").
		Where(sq.Eq{"projects.name": project}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete project: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete project: %v", err.Error())
	}

	return nil
}

func (s *ProjectStore) scanRows(rows *sql.Rows) ([]*model.Project, error) {
	var projects []*model.Project
	for rows.Next() {
		var id, name string
		var registry_schema_version, version_id sql.NullString
		var last_updated sql.NullTime

		if err := rows.Scan(
			&id,
			&name,
			&registry_schema_version,
			&version_id,
			&last_updated); err != nil {
			return nil, err
		}

		projects = append(projects, &model.Project{
			Id:                    id,
			Name:                  name,
			RegistrySchemaVersion: registry_schema_version.String,
			VersionId:             version_id.String,
			LastUpdated:           last_updated.Time,
		})
	}

	return projects, nil
}

func NewProjectStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *ProjectStore {
	return &ProjectStore{db: db, time: time, uuid: uuid}
}
