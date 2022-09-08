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
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	util "github.com/feast-dev/feast/backend/src/utils"
	authorizationv1 "k8s.io/api/authorization/v1"
)

type InfraObjectServerOptions struct{}

type InfraObjectServer struct {
	resourceManager *resource.ResourceManager
	httpClient      *http.Client
	options         *InfraObjectServerOptions
}

func (s *InfraObjectServer) UpdateInfraObjects(ctx context.Context, request *api.UpdateInfraObjectsRequest) (*api.UpdateInfraObjectsResponse, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbUpdate,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	infra_objects, err := s.resourceManager.UpdateInfraObjects(request.InfraObjects.Objects, request.Project)
	if err != nil {
		return nil, util.Wrap(err, "Update infra objects failed")
	}
	apiInfraObjects := ToApiInfraObjects(infra_objects)

	return &api.UpdateInfraObjectsResponse{InfraObjects: apiInfraObjects}, nil

}

func (s *InfraObjectServer) ListInfraObjects(ctx context.Context, request *api.ListInfraObjectsRequest) (*api.ListInfraObjectsResponse, error) {
	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: request.Project,
		Verb:      common.RbacResourceVerbList,
	}

	err := s.haveAccess(ctx, resourceAttributes)
	if err != nil {
		return nil, util.Wrap(err, "Failed to authorize the request")
	}

	infra_objects, err := s.resourceManager.ListInfraObjects(request.Project)
	if err != nil {
		return nil, util.Wrap(err, "List infra objects failed")
	}
	apiInfraObjects := ToApiInfraObjects(infra_objects)

	return &api.ListInfraObjectsResponse{InfraObjects: apiInfraObjects}, nil
}

func (s *InfraObjectServer) haveAccess(ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authorization if not multi-user mode.
		return nil
	}
	if resourceAttributes.Namespace == "" {
		return nil
	}

	resourceAttributes.Group = common.RbacFeaturesGroup
	resourceAttributes.Version = common.RbacFeaturesVersion
	resourceAttributes.Resource = common.RbacResourceTypeInfraObjects

	err := isAuthorized(s.resourceManager, ctx, resourceAttributes)
	if err != nil {
		return util.Wrap(err, "Failed to authorize with API resource references")
	}

	return nil
}

func NewInfraObjectServer(resourceManager *resource.ResourceManager, options *InfraObjectServerOptions) *InfraObjectServer {
	return &InfraObjectServer{resourceManager: resourceManager, httpClient: http.DefaultClient, options: options}
}
