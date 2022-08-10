package server

import (
	"context"
	"net/http"

	api "github.com/feast-dev/feast/backend/api/go_client"
	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	util "github.com/feast-dev/feast/backend/src/utils"
	authorizationv1 "k8s.io/api/authorization/v1"

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
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project.Name,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	project, err := s.resourceManager.CreateProject(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Create project failed")
	}

	return ToApiProject(project), nil
}

func (s *ProjectServer) GetProject(ctx context.Context, request *api.GetProjectRequest) (*api.Project, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	project, err := s.resourceManager.GetProject(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get project failed.")
	}

	return ToApiProject(project), nil
}

func (s *ProjectServer) UpdateProject(ctx context.Context, request *api.UpdateProjectRequest) (*api.Project, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project.Name,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	project, err := s.resourceManager.UpdateProject(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Update project failed")
	}

	return ToApiProject(project), nil
}

func (s *ProjectServer) DeleteProject(ctx context.Context, request *api.DeleteProjectRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteProject(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete project failed")
	}

	return &empty.Empty{}, nil
}

func (s *ProjectServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeProjects

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewProjectServer(resourceManager *resource.ResourceManager, options *ProjectServerOptions) *ProjectServer {
	return &ProjectServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
