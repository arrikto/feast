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

	"github.com/feast-dev/feast/backend/src/apiserver/common"
	"github.com/feast-dev/feast/backend/src/apiserver/resource"
	util "github.com/feast-dev/feast/backend/src/utils"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	authorizationv1 "k8s.io/api/authorization/v1"
)

// isAuthorized verifies whether the user identity, which is contained in the context object,
// can perform some action (verb) on a resource (resourceType/resourceName) living in the
// target namespace. If the returned error is nil, the authorization passes. Otherwise,
// authorization fails with a non-nil error.
func isAuthorized(resourceManager *resource.ResourceManager, ctx context.Context, resourceAttributes *authorizationv1.ResourceAttributes) error {
	if !common.IsMultiUserMode() {
		// Skip authz if not multi-user mode.
		return nil
	}
	if common.IsMultiUserSharedReadMode() &&
		(resourceAttributes.Verb == common.RbacResourceVerbGet ||
			resourceAttributes.Verb == common.RbacResourceVerbList) {
		glog.Infof("Multi-user shared read mode is enabled. Request allowed: %+v", resourceAttributes)
		return nil
	}

	glog.Info("Getting user identity and groups...")
	userIdentity, userGroups, err := resourceManager.AuthenticateRequest(ctx)
	if err != nil {
		return err
	}

	if len(userIdentity) == 0 {
		return util.NewUnauthenticatedError(errors.New("Request header error: user identity is empty."), "Request header error: user identity is empty.")
	}

	glog.Infof("User: %s, ResourceAttributes: %+v", userIdentity, resourceAttributes)
	glog.Info("Authorizing request...")
	err = resourceManager.IsRequestAuthorized(ctx, userIdentity, userGroups, resourceAttributes)
	if err != nil {
		glog.Info(err.Error())
		return err
	}

	glog.Infof("Authorized user '%s': %+v", userIdentity, resourceAttributes)
	return nil
}