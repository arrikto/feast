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

type RequestFeatureViewServerOptions struct{}

type RequestFeatureViewServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *RequestFeatureViewServerOptions
}

func (s *RequestFeatureViewServer) CreateRequestFeatureView(ctx context.Context, request *api.CreateRequestFeatureViewRequest) (*api.RequestFeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.RequestFeatureView.Name,
		Namespace: request.RequestFeatureView.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	request_feature_view, err := s.resourceManager.CreateRequestFeatureView(request.RequestFeatureView)
	if err != nil {
		return nil, util.Wrap(err, "Create request feature view failed")
	}

	return ToApiRequestFeatureView(request_feature_view), nil
}

func (s *RequestFeatureViewServer) GetRequestFeatureView(ctx context.Context, request *api.GetRequestFeatureViewRequest) (*api.RequestFeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	request_feature_view, err := s.resourceManager.GetRequestFeatureView(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get request feature view failed")
	}

	return ToApiRequestFeatureView(request_feature_view), nil
}

func (s *RequestFeatureViewServer) UpdateRequestFeatureView(ctx context.Context, request *api.UpdateRequestFeatureViewRequest) (*api.RequestFeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.RequestFeatureView.Name,
		Namespace: request.RequestFeatureView.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	request_feature_view, err := s.resourceManager.UpdateRequestFeatureView(request.RequestFeatureView)
	if err != nil {
		return nil, util.Wrap(err, "Update request feature view failed")
	}

	return ToApiRequestFeatureView(request_feature_view), nil
}

func (s *RequestFeatureViewServer) DeleteRequestFeatureView(ctx context.Context, request *api.DeleteRequestFeatureViewRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteRequestFeatureView(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete request feature view failed")
	}

	return &empty.Empty{}, nil
}

func (s *RequestFeatureViewServer) ListRequestFeatureViews(ctx context.Context, request *api.ListRequestFeatureViewsRequest) (*api.ListRequestFeatureViewsResponse, error) {
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

	request_feature_views, err := s.resourceManager.ListRequestFeatureViews(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List request feature views failed")
	}

	if listDenied {
		var allowedRequestFeatureViews []*model.RequestFeatureView
		for _, rfv := range request_feature_views {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      rfv.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedRequestFeatureViews = append(allowedRequestFeatureViews, rfv)
			}
		}
		request_feature_views = allowedRequestFeatureViews
	}
	apiRequestFeatureViews := ToApiRequestFeatureViews(request_feature_views)

	return &api.ListRequestFeatureViewsResponse{RequestFeatureViews: apiRequestFeatureViews}, nil
}

func (s *RequestFeatureViewServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeRequestFeatureViews

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewRequestFeatureViewServer(resourceManager *resource.ResourceManager, options *RequestFeatureViewServerOptions) *RequestFeatureViewServer {
	return &RequestFeatureViewServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
