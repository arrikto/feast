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

type OnDemandFeatureViewServerOptions struct{}

type OnDemandFeatureViewServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *OnDemandFeatureViewServerOptions
}

func (s *OnDemandFeatureViewServer) CreateOnDemandFeatureView(ctx context.Context, request *api.CreateOnDemandFeatureViewRequest) (*api.OnDemandFeatureView, error) {
	on_demand_feature_view, err := s.resourceManager.CreateOnDemandFeatureView(request.OnDemandFeatureView, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create on demand feature view failed")
	}

	return ToApiOnDemandFeatureView(on_demand_feature_view), nil
}

func (s *OnDemandFeatureViewServer) GetOnDemandFeatureView(ctx context.Context, request *api.GetOnDemandFeatureViewRequest) (*api.OnDemandFeatureView, error) {
	on_demand_feature_view, err := s.resourceManager.GetOnDemandFeatureView(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get on demand feature view failed")
	}

	return ToApiOnDemandFeatureView(on_demand_feature_view), nil
}

func (s *OnDemandFeatureViewServer) UpdateOnDemandFeatureView(ctx context.Context, request *api.UpdateOnDemandFeatureViewRequest) (*api.OnDemandFeatureView, error) {
	on_demand_feature_view, err := s.resourceManager.UpdateOnDemandFeatureView(request.OnDemandFeatureView, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update on demand feature view failed")
	}

	return ToApiOnDemandFeatureView(on_demand_feature_view), nil
}

func (s *OnDemandFeatureViewServer) DeleteOnDemandFeatureView(ctx context.Context, request *api.DeleteOnDemandFeatureViewRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteOnDemandFeatureView(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete on demand feature view failed")
	}

	return &empty.Empty{}, nil
}

func (s *OnDemandFeatureViewServer) ListOnDemandFeatureViews(ctx context.Context, request *api.ListOnDemandFeatureViewsRequest) (*api.ListOnDemandFeatureViewsResponse, error) {
	on_demand_feature_views, err := s.resourceManager.ListOnDemandFeatureViews(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List on demand feature views failed")
	}
	apiOnDemandFeatureViews := ToApiOnDemandFeatureViews(on_demand_feature_views)

	return &api.ListOnDemandFeatureViewsResponse{OnDemandFeatureViews: apiOnDemandFeatureViews}, nil
}

func NewOnDemandFeatureViewServer(resourceManager *resource.ResourceManager, options *OnDemandFeatureViewServerOptions) *OnDemandFeatureViewServer {
	return &OnDemandFeatureViewServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
