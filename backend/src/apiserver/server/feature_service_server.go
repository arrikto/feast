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

type FeatureServiceServerOptions struct{}

type FeatureServiceServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *FeatureServiceServerOptions
}

func (s *FeatureServiceServer) CreateFeatureService(ctx context.Context, request *api.CreateFeatureServiceRequest) (*api.FeatureService, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.FeatureService.Name,
		Namespace: request.FeatureService.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	feature_service, err := s.resourceManager.CreateFeatureService(request.FeatureService)
	if err != nil {
		return nil, util.Wrap(err, "Create feature service failed")
	}

	return ToApiFeatureService(feature_service), nil
}

func (s *FeatureServiceServer) GetFeatureService(ctx context.Context, request *api.GetFeatureServiceRequest) (*api.FeatureService, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	feature_service, err := s.resourceManager.GetFeatureService(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get feature service failed")
	}

	return ToApiFeatureService(feature_service), nil
}

func (s *FeatureServiceServer) UpdateFeatureService(ctx context.Context, request *api.UpdateFeatureServiceRequest) (*api.FeatureService, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.FeatureService.Name,
		Namespace: request.FeatureService.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	feature_service, err := s.resourceManager.UpdateFeatureService(request.FeatureService)
	if err != nil {
		return nil, util.Wrap(err, "Update feature service failed")
	}

	return ToApiFeatureService(feature_service), nil
}

func (s *FeatureServiceServer) DeleteFeatureService(ctx context.Context, request *api.DeleteFeatureServiceRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteFeatureService(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete feature service failed")
	}

	return &empty.Empty{}, nil
}

func (s *FeatureServiceServer) ListFeatureServices(ctx context.Context, request *api.ListFeatureServicesRequest) (*api.ListFeatureServicesResponse, error) {
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

	feature_services, err := s.resourceManager.ListFeatureServices(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List feature services failed")
	}

	if listDenied {
		var allowedFeatureServices []*model.FeatureService
		for _, fs := range feature_services {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      fs.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedFeatureServices = append(allowedFeatureServices, fs)
			}
		}
		feature_services = allowedFeatureServices
	}
	apiFeatureServices := ToApiFeatureServices(feature_services)

	return &api.ListFeatureServicesResponse{FeatureServices: apiFeatureServices}, nil
}

func (s *FeatureServiceServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeFeatureServices

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewFeatureServiceServer(resourceManager *resource.ResourceManager, options *FeatureServiceServerOptions) *FeatureServiceServer {
	return &FeatureServiceServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
