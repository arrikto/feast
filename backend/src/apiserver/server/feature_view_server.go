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

type FeatureViewServerOptions struct{}

type FeatureViewServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *FeatureViewServerOptions
}

func (s *FeatureViewServer) CreateFeatureView(ctx context.Context, request *api.CreateFeatureViewRequest) (*api.FeatureView, error) {
	feature_view, err := s.resourceManager.CreateFeatureView(request.FeatureView, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create feature view failed")
	}

	return ToApiFeatureView(feature_view), nil
}

func (s *FeatureViewServer) GetFeatureView(ctx context.Context, request *api.GetFeatureViewRequest) (*api.FeatureView, error) {
	feature_view, err := s.resourceManager.GetFeatureView(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get feature view failed")
	}

	return ToApiFeatureView(feature_view), nil
}

func (s *FeatureViewServer) UpdateFeatureView(ctx context.Context, request *api.UpdateFeatureViewRequest) (*api.FeatureView, error) {
	feature_view, err := s.resourceManager.UpdateFeatureView(request.FeatureView, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update feature view failed")
	}

	return ToApiFeatureView(feature_view), nil
}

func (s *FeatureViewServer) DeleteFeatureView(ctx context.Context, request *api.DeleteFeatureViewRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteFeatureView(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete feature view failed")
	}

	return &empty.Empty{}, nil
}

func (s *FeatureViewServer) ListFeatureViews(ctx context.Context, request *api.ListFeatureViewsRequest) (*api.ListFeatureViewsResponse, error) {
	feature_views, err := s.resourceManager.ListFeatureViews(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List feature views failed")
	}
	apiFeatureViews := ToApiFeatureViews(feature_views)

	return &api.ListFeatureViewsResponse{FeatureViews: apiFeatureViews}, nil
}

func (s *FeatureViewServer) ReportMI(ctx context.Context, request *api.ReportMIRequest) (*api.MaterializationInterval, error) {
	mi, err := s.resourceManager.ReportMI(request.MaterializationInterval, request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Report materialization interval failed")
	}

	return ToApiMI(mi), nil
}

func NewFeatureViewServer(resourceManager *resource.ResourceManager, options *FeatureViewServerOptions) *FeatureViewServer {
	return &FeatureViewServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
