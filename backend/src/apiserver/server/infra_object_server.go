package server

import (
	"context"
	"net/http"

	api "github.com/feast-dev/feast/backend/api/go_client"
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	util "github.com/feast-dev/feast/backend/src/utils"
)

type InfraObjectServerOptions struct{}

type InfraObjectServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *InfraObjectServerOptions
}

func (s *InfraObjectServer) UpdateInfraObjects(ctx context.Context, request *api.UpdateInfraObjectsRequest) (*api.UpdateInfraObjectsResponse, error) {
	infra_objects, err := s.resourceManager.UpdateInfraObjects(request.InfraObjects.Objects, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update infra objects failed")
	}
	apiInfraObjects := ToApiInfraObjects(infra_objects)

	return &api.UpdateInfraObjectsResponse{InfraObjects: apiInfraObjects}, nil

}

func (s *InfraObjectServer) ListInfraObjects(ctx context.Context, request *api.ListInfraObjectsRequest) (*api.ListInfraObjectsResponse, error) {
	infra_objects, err := s.resourceManager.ListInfraObjects(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List infra objects failed")
	}
	apiInfraObjects := ToApiInfraObjects(infra_objects)

	return &api.ListInfraObjectsResponse{InfraObjects: apiInfraObjects}, nil
}

func NewInfraObjectServer(resourceManager *resource.ResourceManager, options *InfraObjectServerOptions) *InfraObjectServer {
	return &InfraObjectServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
