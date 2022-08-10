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

type SavedDatasetServerOptions struct{}

type SavedDatasetServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *SavedDatasetServerOptions
}

func (s *SavedDatasetServer) CreateSavedDataset(ctx context.Context, request *api.CreateSavedDatasetRequest) (*api.SavedDataset, error) {
	saved_dataset, err := s.resourceManager.CreateSavedDataset(request.SavedDataset, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create saved dataset failed")
	}

	return ToApiSavedDataset(saved_dataset), nil
}

func (s *SavedDatasetServer) GetSavedDataset(ctx context.Context, request *api.GetSavedDatasetRequest) (*api.SavedDataset, error) {
	saved_dataset, err := s.resourceManager.GetSavedDataset(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get saved dataset failed")
	}

	return ToApiSavedDataset(saved_dataset), nil
}

func (s *SavedDatasetServer) UpdateSavedDataset(ctx context.Context, request *api.UpdateSavedDatasetRequest) (*api.SavedDataset, error) {
	saved_dataset, err := s.resourceManager.UpdateSavedDataset(request.SavedDataset, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update saved dataset failed")
	}

	return ToApiSavedDataset(saved_dataset), nil
}

func (s *SavedDatasetServer) DeleteSavedDataset(ctx context.Context, request *api.DeleteSavedDatasetRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteSavedDataset(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete saved dataset failed")
	}

	return &empty.Empty{}, nil
}

func (s *SavedDatasetServer) ListSavedDatasets(ctx context.Context, request *api.ListSavedDatasetsRequest) (*api.ListSavedDatasetsResponse, error) {
	saved_datasets, err := s.resourceManager.ListSavedDatasets(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List saved datasets failed")
	}
	apiSavedDatasets := ToApiSavedDatasets(saved_datasets)

	return &api.ListSavedDatasetsResponse{SavedDatasets: apiSavedDatasets}, nil
}

func NewSavedDatasetServer(resourceManager *resource.ResourceManager, options *SavedDatasetServerOptions) *SavedDatasetServer {
	return &SavedDatasetServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
