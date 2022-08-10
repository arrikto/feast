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

type RequestFeatureViewServerOptions struct{}

type RequestFeatureViewServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *RequestFeatureViewServerOptions
}

func (s *RequestFeatureViewServer) CreateRequestFeatureView(ctx context.Context, request *api.CreateRequestFeatureViewRequest) (*api.RequestFeatureView, error) {
	request_feature_view, err := s.resourceManager.CreateRequestFeatureView(request.RequestFeatureView, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create request feature view failed")
	}

	return ToApiRequestFeatureView(request_feature_view), nil
}

func (s *RequestFeatureViewServer) GetRequestFeatureView(ctx context.Context, request *api.GetRequestFeatureViewRequest) (*api.RequestFeatureView, error) {
	request_feature_view, err := s.resourceManager.GetRequestFeatureView(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get request feature view failed")
	}

	return ToApiRequestFeatureView(request_feature_view), nil
}

func (s *RequestFeatureViewServer) UpdateRequestFeatureView(ctx context.Context, request *api.UpdateRequestFeatureViewRequest) (*api.RequestFeatureView, error) {
	request_feature_view, err := s.resourceManager.UpdateRequestFeatureView(request.RequestFeatureView, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update request feature view failed")
	}

	return ToApiRequestFeatureView(request_feature_view), nil
}

func (s *RequestFeatureViewServer) DeleteRequestFeatureView(ctx context.Context, request *api.DeleteRequestFeatureViewRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteRequestFeatureView(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete request feature view failed")
	}

	return &empty.Empty{}, nil
}

func (s *RequestFeatureViewServer) ListRequestFeatureViews(ctx context.Context, request *api.ListRequestFeatureViewsRequest) (*api.ListRequestFeatureViewsResponse, error) {
	request_feature_views, err := s.resourceManager.ListRequestFeatureViews(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List request feature views failed")
	}
	apiRequestFeatureViews := ToApiRequestFeatureViews(request_feature_views)

	return &api.ListRequestFeatureViewsResponse{RequestFeatureViews: apiRequestFeatureViews}, nil
}

func NewRequestFeatureViewServer(resourceManager *resource.ResourceManager, options *RequestFeatureViewServerOptions) *RequestFeatureViewServer {
	return &RequestFeatureViewServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
