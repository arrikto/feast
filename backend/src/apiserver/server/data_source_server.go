// Copyright 2022 Arrikto Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

type DataSourceServerOptions struct{}

type DataSourceServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *DataSourceServerOptions
}

func (s *DataSourceServer) CreateDataSource(ctx context.Context, request *api.CreateDataSourceRequest) (*api.DataSource, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.DataSource.Name,
		Namespace: request.DataSource.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	data_source, err := s.resourceManager.CreateDataSource(request.DataSource)
	if err != nil {
		return nil, util.Wrap(err, "Create data source failed")
	}

	return ToApiDataSource(data_source), nil
}

func (s *DataSourceServer) GetDataSource(ctx context.Context, request *api.GetDataSourceRequest) (*api.DataSource, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	data_source, err := s.resourceManager.GetDataSource(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get data source failed")
	}

	return ToApiDataSource(data_source), nil
}

func (s *DataSourceServer) UpdateDataSource(ctx context.Context, request *api.UpdateDataSourceRequest) (*api.DataSource, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.DataSource.Name,
		Namespace: request.DataSource.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	data_source, err := s.resourceManager.UpdateDataSource(request.DataSource)
	if err != nil {
		return nil, util.Wrap(err, "Update data source failed")
	}

	return ToApiDataSource(data_source), nil
}

func (s *DataSourceServer) DeleteDataSource(ctx context.Context, request *api.DeleteDataSourceRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteDataSource(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete data source failed")
	}

	return &empty.Empty{}, nil
}

func (s *DataSourceServer) ListDataSources(ctx context.Context, request *api.ListDataSourcesRequest) (*api.ListDataSourcesResponse, error) {
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

	data_sources, err := s.resourceManager.ListDataSources(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List data sources failed")
	}

	if listDenied {
		var allowedDataSources []*model.DataSource
		for _, ds := range data_sources {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      ds.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedDataSources = append(allowedDataSources, ds)
			}
		}
		data_sources = allowedDataSources
	}
	apiDataSources := ToApiDataSources(data_sources)

	return &api.ListDataSourcesResponse{DataSources: apiDataSources}, nil
}

func (s *DataSourceServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeDataSources

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewDataSourceServer(resourceManager *resource.ResourceManager, options *DataSourceServerOptions) *DataSourceServer {
	return &DataSourceServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
