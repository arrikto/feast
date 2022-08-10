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

type EntityServerOptions struct{}

type EntityServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *EntityServerOptions
}

func (s *EntityServer) CreateEntity(ctx context.Context, request *api.CreateEntityRequest) (*api.Entity, error) {
	entity, err := s.resourceManager.CreateEntity(request.Entity, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create entity failed")
	}

	return ToApiEntity(entity), nil
}

func (s *EntityServer) GetEntity(ctx context.Context, request *api.GetEntityRequest) (*api.Entity, error) {
	entity, err := s.resourceManager.GetEntity(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get entity failed")
	}

	return ToApiEntity(entity), nil
}

func (s *EntityServer) UpdateEntity(ctx context.Context, request *api.UpdateEntityRequest) (*api.Entity, error) {
	entity, err := s.resourceManager.UpdateEntity(request.Entity, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update entity failed")
	}

	return ToApiEntity(entity), nil
}

func (s *EntityServer) DeleteEntity(ctx context.Context, request *api.DeleteEntityRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteEntity(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete entity failed")
	}

	return &empty.Empty{}, nil
}

func (s *EntityServer) ListEntities(ctx context.Context, request *api.ListEntitiesRequest) (*api.ListEntitiesResponse, error) {
	entities, err := s.resourceManager.ListEntities(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List entities failed")
	}
	apiEntities := ToApiEntities(entities)

	return &api.ListEntitiesResponse{Entities: apiEntities}, nil
}

func NewEntityServer(resourceManager *resource.ResourceManager, options *EntityServerOptions) *EntityServer {
	return &EntityServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
