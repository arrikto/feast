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

package common

import (
	"github.com/feast-dev/feast/backend/src/apiserver/model"
)

const (
	Project             model.ResourceType = "Project"
	Entity              model.ResourceType = "Entity"
	DataSource          model.ResourceType = "DataSource"
	FeatureService      model.ResourceType = "FeatureService"
	FeatureView         model.ResourceType = "FeatureView"
	InfraObject         model.ResourceType = "InfraObject"
	OnDemandFeatureView model.ResourceType = "OnDemandFeatureView"
	RequestFeatureView  model.ResourceType = "RequestFeatureView"
	SavedDataset        model.ResourceType = "SavedDataset"
)

const (
	RbacKubeflowGroup   = "kubeflow.org"
	RbacFeaturesGroup   = "features.kubeflow.org"
	RbacFeaturesVersion = "v1beta1"

	RbacResourceTypeProjects             = "projects"
	RbacResourceTypeEntities             = "entities"
	RbacResourceTypeDataSources          = "data_sources"
	RbacResourceTypeFeatureServices      = "feature_services"
	RbacResourceTypeFeatureViews         = "feature_views"
	RbacResourceTypeInfraObjects         = "infra_objects"
	RbacResourceTypeOnDemandFeatureViews = "on_demand_feature_views"
	RbacResourceTypeRequestFeatureViews  = "request_feature_views"
	RbacResourceTypeSavedDatasets        = "saved_datasets"

	RbacResourceVerbUpdate = "update"
	RbacResourceVerbCreate = "create"
	RbacResourceVerbDelete = "delete"
	RbacResourceVerbGet    = "get"
	RbacResourceVerbList   = "list"
)

const (
	GoogleIAPUserIdentityHeader    string = "x-goog-authenticated-user-email"
	GoogleIAPUserIdentityPrefix    string = "accounts.google.com:"
	AuthorizationBearerTokenHeader string = "Authorization"
	AuthorizationBearerTokenPrefix string = "Bearer "
)

const DefaultTokenReviewAudience string = "features.kubeflow.org"
