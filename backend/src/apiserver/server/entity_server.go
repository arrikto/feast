package server

import (
	"context"
	"net/http"

	api "github.com/feast-dev/feast/backend/api/go_client"
	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/model"
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	util "github.com/feast-dev/feast/backend/src/utils"
	authorizationv1 "k8s.io/api/authorization/v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EntityServerOptions struct{}

type EntityServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *EntityServerOptions
}

func (s *EntityServer) CreateEntity(ctx context.Context, request *api.CreateEntityRequest) (*api.Entity, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Entity.Name,
		Namespace: request.Entity.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	entity, err := s.resourceManager.CreateEntity(request.Entity)
	if err != nil {
		return nil, util.Wrap(err, "Create entity failed")
	}

	return ToApiEntity(entity), nil
}

func (s *EntityServer) GetEntity(ctx context.Context, request *api.GetEntityRequest) (*api.Entity, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	entity, err := s.resourceManager.GetEntity(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get entity failed")
	}

	return ToApiEntity(entity), nil
}

func (s *EntityServer) UpdateEntity(ctx context.Context, request *api.UpdateEntityRequest) (*api.Entity, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Entity.Name,
		Namespace: request.Entity.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	entity, err := s.resourceManager.UpdateEntity(request.Entity)
	if err != nil {
		return nil, util.Wrap(err, "Update entity failed")
	}

	return ToApiEntity(entity), nil
}

func (s *EntityServer) DeleteEntity(ctx context.Context, request *api.DeleteEntityRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteEntity(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete entity failed")
	}

	return &empty.Empty{}, nil
}

func (s *EntityServer) ListEntities(ctx context.Context, request *api.ListEntitiesRequest) (*api.ListEntitiesResponse, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbList,
	}

	var listDenied = false
	err := s.haveAccess(ctx, resourceAttributes)
	if err, ok := err.(*util.UserError); ok && err.ExternalStatusCode() == codes.PermissionDenied {
		listDenied = true
	} else if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	entities, err := s.resourceManager.ListEntities(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List entities failed")
	}

	if listDenied {
		var allowedEntities []*model.Entity
		for _, entity := range entities {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      entity.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedEntities = append(allowedEntities, entity)
			}
		}
		entities = allowedEntities
	}
	apiEntities := ToApiEntities(entities)

	return &api.ListEntitiesResponse{Entities: apiEntities}, nil
}

func (s *EntityServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeEntities

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewEntityServer(resourceManager *resource.ResourceManager, options *EntityServerOptions) *EntityServer {
	return &EntityServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
