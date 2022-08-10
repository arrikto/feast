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
