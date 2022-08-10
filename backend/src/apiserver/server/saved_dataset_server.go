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

type SavedDatasetServerOptions struct{}

type SavedDatasetServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *SavedDatasetServerOptions
}

func (s *SavedDatasetServer) CreateSavedDataset(ctx context.Context, request *api.CreateSavedDatasetRequest) (*api.SavedDataset, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.SavedDataset.Name,
		Namespace: request.SavedDataset.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	saved_dataset, err := s.resourceManager.CreateSavedDataset(request.SavedDataset)
	if err != nil {
		return nil, util.Wrap(err, "Create saved dataset failed")
	}

	return ToApiSavedDataset(saved_dataset), nil
}

func (s *SavedDatasetServer) GetSavedDataset(ctx context.Context, request *api.GetSavedDatasetRequest) (*api.SavedDataset, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	saved_dataset, err := s.resourceManager.GetSavedDataset(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get saved dataset failed")
	}

	return ToApiSavedDataset(saved_dataset), nil
}

func (s *SavedDatasetServer) UpdateSavedDataset(ctx context.Context, request *api.UpdateSavedDatasetRequest) (*api.SavedDataset, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.SavedDataset.Name,
		Namespace: request.SavedDataset.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	saved_dataset, err := s.resourceManager.UpdateSavedDataset(request.SavedDataset)
	if err != nil {
		return nil, util.Wrap(err, "Update saved dataset failed")
	}

	return ToApiSavedDataset(saved_dataset), nil
}

func (s *SavedDatasetServer) DeleteSavedDataset(ctx context.Context, request *api.DeleteSavedDatasetRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteSavedDataset(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete saved dataset failed")
	}

	return &empty.Empty{}, nil
}

func (s *SavedDatasetServer) ListSavedDatasets(ctx context.Context, request *api.ListSavedDatasetsRequest) (*api.ListSavedDatasetsResponse, error) {
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

	saved_datasets, err := s.resourceManager.ListSavedDatasets(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List saved datasets failed")
	}

	if listDenied {
		var allowedSavedDatasets []*model.SavedDataset
		for _, sd := range saved_datasets {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      sd.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedSavedDatasets = append(allowedSavedDatasets, sd)
			}
		}
		saved_datasets = allowedSavedDatasets
	}
	apiSavedDatasets := ToApiSavedDatasets(saved_datasets)

	return &api.ListSavedDatasetsResponse{SavedDatasets: apiSavedDatasets}, nil
}

func (s *SavedDatasetServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeSavedDatasets

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewSavedDatasetServer(resourceManager *resource.ResourceManager, options *SavedDatasetServerOptions) *SavedDatasetServer {
	return &SavedDatasetServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
