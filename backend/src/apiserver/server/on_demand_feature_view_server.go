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

type OnDemandFeatureViewServerOptions struct{}

type OnDemandFeatureViewServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *OnDemandFeatureViewServerOptions
}

func (s *OnDemandFeatureViewServer) CreateOnDemandFeatureView(ctx context.Context, request *api.CreateOnDemandFeatureViewRequest) (*api.OnDemandFeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.OnDemandFeatureView.Name,
		Namespace: request.OnDemandFeatureView.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	on_demand_feature_view, err := s.resourceManager.CreateOnDemandFeatureView(request.OnDemandFeatureView)
	if err != nil {
		return nil, util.Wrap(err, "Create on demand feature view failed")
	}

	return ToApiOnDemandFeatureView(on_demand_feature_view), nil
}

func (s *OnDemandFeatureViewServer) GetOnDemandFeatureView(ctx context.Context, request *api.GetOnDemandFeatureViewRequest) (*api.OnDemandFeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	on_demand_feature_view, err := s.resourceManager.GetOnDemandFeatureView(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get on demand feature view failed")
	}

	return ToApiOnDemandFeatureView(on_demand_feature_view), nil
}

func (s *OnDemandFeatureViewServer) UpdateOnDemandFeatureView(ctx context.Context, request *api.UpdateOnDemandFeatureViewRequest) (*api.OnDemandFeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.OnDemandFeatureView.Name,
		Namespace: request.OnDemandFeatureView.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	on_demand_feature_view, err := s.resourceManager.UpdateOnDemandFeatureView(request.OnDemandFeatureView)
	if err != nil {
		return nil, util.Wrap(err, "Update on demand feature view failed")
	}

	return ToApiOnDemandFeatureView(on_demand_feature_view), nil
}

func (s *OnDemandFeatureViewServer) DeleteOnDemandFeatureView(ctx context.Context, request *api.DeleteOnDemandFeatureViewRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteOnDemandFeatureView(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete on demand feature view failed")
	}

	return &empty.Empty{}, nil
}

func (s *OnDemandFeatureViewServer) ListOnDemandFeatureViews(ctx context.Context, request *api.ListOnDemandFeatureViewsRequest) (*api.ListOnDemandFeatureViewsResponse, error) {
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

	on_demand_feature_views, err := s.resourceManager.ListOnDemandFeatureViews(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List on demand feature views failed")
	}

	if listDenied {
		var allowedOnDemandFeatureViews []*model.OnDemandFeatureView
		for _, odfv := range on_demand_feature_views {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      odfv.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedOnDemandFeatureViews = append(allowedOnDemandFeatureViews, odfv)
			}
		}
		on_demand_feature_views = allowedOnDemandFeatureViews
	}
	apiOnDemandFeatureViews := ToApiOnDemandFeatureViews(on_demand_feature_views)

	return &api.ListOnDemandFeatureViewsResponse{OnDemandFeatureViews: apiOnDemandFeatureViews}, nil
}

func (s *OnDemandFeatureViewServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeOnDemandFeatureViews

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewOnDemandFeatureViewServer(resourceManager *resource.ResourceManager, options *OnDemandFeatureViewServerOptions) *OnDemandFeatureViewServer {
	return &OnDemandFeatureViewServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
