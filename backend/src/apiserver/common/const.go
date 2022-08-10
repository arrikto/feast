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
