package server

import (
	"context"
	"net/http"

	api "github.com/feast-dev/feast/backend/api/go_client"
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	util "github.com/feast-dev/feast/backend/src/utils"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProjectServerOptions struct{}

type ProjectServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *ProjectServerOptions
}

func (s *ProjectServer) CreateProject(ctx context.Context, request *api.CreateProjectRequest) (*api.Project, error) {
	project, err := s.resourceManager.CreateProject(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create project failed")
	}

	return ToApiProject(project), nil
}

func (s *ProjectServer) GetProject(ctx context.Context, request *api.GetProjectRequest) (*api.Project, error) {
	project, err := s.resourceManager.GetProject(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get project failed.")
	}

	return ToApiProject(project), nil
}

func (s *ProjectServer) UpdateProject(ctx context.Context, request *api.UpdateProjectRequest) (*api.Project, error) {
	project, err := s.resourceManager.UpdateProject(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update project failed")
	}

	return ToApiProject(project), nil
}

func (s *ProjectServer) DeleteProject(ctx context.Context, request *api.DeleteProjectRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteProject(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete project failed")
	}

	return &empty.Empty{}, nil
}

func NewProjectServer(resourceManager *resource.ResourceManager, options *ProjectServerOptions) *ProjectServer {
	return &ProjectServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
