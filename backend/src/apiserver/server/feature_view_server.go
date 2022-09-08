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

type FeatureViewServerOptions struct{}

type FeatureViewServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *FeatureViewServerOptions
}

func (s *FeatureViewServer) CreateFeatureView(ctx context.Context, request *api.CreateFeatureViewRequest) (*api.FeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.FeatureView.Name,
		Namespace: request.FeatureView.Project,
		Verb:      common.RbacResourceVerbCreate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	feature_view, err := s.resourceManager.CreateFeatureView(request.FeatureView)
	if err != nil {
		return nil, util.Wrap(err, "Create feature view failed")
	}

	return ToApiFeatureView(feature_view), nil
}

func (s *FeatureViewServer) GetFeatureView(ctx context.Context, request *api.GetFeatureViewRequest) (*api.FeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbGet,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	feature_view, err := s.resourceManager.GetFeatureView(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Get feature view failed")
	}

	return ToApiFeatureView(feature_view), nil
}

func (s *FeatureViewServer) UpdateFeatureView(ctx context.Context, request *api.UpdateFeatureViewRequest) (*api.FeatureView, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.FeatureView.Name,
		Namespace: request.FeatureView.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	feature_view, err := s.resourceManager.UpdateFeatureView(request.FeatureView)
	if err != nil {
		return nil, util.Wrap(err, "Update feature view failed")
	}

	return ToApiFeatureView(feature_view), nil
}

func (s *FeatureViewServer) DeleteFeatureView(ctx context.Context, request *api.DeleteFeatureViewRequest) (*emptypb.Empty, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbDelete,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	err = s.resourceManager.DeleteFeatureView(request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Delete feature view failed")
	}

	return &empty.Empty{}, nil
}

func (s *FeatureViewServer) ListFeatureViews(ctx context.Context, request *api.ListFeatureViewsRequest) (*api.ListFeatureViewsResponse, error) {
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

	feature_views, err := s.resourceManager.ListFeatureViews(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List feature views failed")
	}

	if listDenied {
		var allowedFeatureViews []*model.FeatureView
		for _, fv := range feature_views {
			resourceAttributes = &authorizationv1.ResourceAttributes{
				Name:      fv.Name,
				Namespace: request.Project,
				Verb:      common.RbacResourceVerbList,
			}
			err = s.haveAccess(ctx, resourceAttributes)
			if err != nil {
				continue
			} else {
				allowedFeatureViews = append(allowedFeatureViews, fv)
			}
		}
		feature_views = allowedFeatureViews
	}
	apiFeatureViews := ToApiFeatureViews(feature_views)

	return &api.ListFeatureViewsResponse{FeatureViews: apiFeatureViews}, nil
}

func (s *FeatureViewServer) ReportMI(ctx context.Context, request *api.ReportMIRequest) (*api.MaterializationInterval, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Name:      request.Name,
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	mi, err := s.resourceManager.ReportMI(request.MaterializationInterval, request.Name, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Report materialization interval failed")
	}

	return ToApiMI(mi), nil
}

func (s *FeatureViewServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeFeatureViews

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewFeatureViewServer(resourceManager *resource.ResourceManager, options *FeatureViewServerOptions) *FeatureViewServer {
	return &FeatureViewServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
