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

type DataSourceServerOptions struct{}

type DataSourceServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *DataSourceServerOptions
}

func (s *DataSourceServer) CreateDataSource(ctx context.Context, request *api.CreateDataSourceRequest) (*api.DataSource, error) {
	data_source, err := s.resourceManager.CreateDataSource(request.DataSource, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Create data source failed")
	}

	return ToApiDataSource(data_source), nil
}

func (s *DataSourceServer) GetDataSource(ctx context.Context, request *api.GetDataSourceRequest) (*api.DataSource, error) {
	data_source, err := s.resourceManager.GetDataSource(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Get data source failed")
	}

	return ToApiDataSource(data_source), nil
}

func (s *DataSourceServer) UpdateDataSource(ctx context.Context, request *api.UpdateDataSourceRequest) (*api.DataSource, error) {
	data_source, err := s.resourceManager.UpdateDataSource(request.DataSource, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Update data source failed")
	}

	return ToApiDataSource(data_source), nil
}

func (s *DataSourceServer) DeleteDataSource(ctx context.Context, request *api.DeleteDataSourceRequest) (*emptypb.Empty, error) {
	err := s.resourceManager.DeleteDataSource(request.Name, request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "Delete data source failed")
	}

	return &empty.Empty{}, nil
}

func (s *DataSourceServer) ListDataSources(ctx context.Context, request *api.ListDataSourcesRequest) (*api.ListDataSourcesResponse, error) {
	data_sources, err := s.resourceManager.ListDataSources(request.Project, request.Namespace)
	if err != nil {
		return nil, util.Wrap(err, "List data sources failed")
	}
	apiDataSources := ToApiDataSources(data_sources)

	return &api.ListDataSourcesResponse{DataSources: apiDataSources}, nil
}

func NewDataSourceServer(resourceManager *resource.ResourceManager, options *DataSourceServerOptions) *DataSourceServer {
	return &DataSourceServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
