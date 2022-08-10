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

type FeatureServiceServerOptions struct{}

type FeatureServiceServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *FeatureServiceServerOptions
}

func (s *FeatureServiceServer) CreateFeatureService(ctx context.Context, request *api.CreateFeatureServiceRequest) (*api.FeatureService, error) {
	feature_service, err := s.resourceManager.CreateFeatureService(request.FeatureService, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create feature service failed")
	}

	return ToApiFeatureService(feature_service), nil
}

func (s *FeatureServiceServer) GetFeatureService(ctx context.Context, request *api.GetFeatureServiceRequest) (*api.FeatureService, error) {
	feature_service, err := s.resourceManager.GetFeatureService(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get feature service failed")
	}

	return ToApiFeatureService(feature_service), nil
}

func (s *FeatureServiceServer) UpdateFeatureService(ctx context.Context, request *api.UpdateFeatureServiceRequest) (*api.FeatureService, error) {
	feature_service, err := s.resourceManager.UpdateFeatureService(request.FeatureService, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update feature service failed")
	}

	return ToApiFeatureService(feature_service), nil
}

func (s *FeatureServiceServer) DeleteFeatureService(ctx context.Context, request *api.DeleteFeatureServiceRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteFeatureService(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete feature service failed")
	}

	return &empty.Empty{}, nil
}

func (s *FeatureServiceServer) ListFeatureServices(ctx context.Context, request *api.ListFeatureServicesRequest) (*api.ListFeatureServicesResponse, error) {
	feature_services, err := s.resourceManager.ListFeatureServices(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List feature services failed")
	}
	apiFeatureServices := ToApiFeatureServices(feature_services)

	return &api.ListFeatureServicesResponse{FeatureServices: apiFeatureServices}, nil
}

func NewFeatureServiceServer(resourceManager *resource.ResourceManager, options *FeatureServiceServerOptions) *FeatureServiceServer {
	return &FeatureServiceServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
