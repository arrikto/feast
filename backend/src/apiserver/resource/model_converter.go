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

package resource

import (
	"encoding/json"

	api "github.com/feast-dev/feast/backend/api/go_client"
	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"
)

func (r *ResourceManager) ToModelDataSource(ds *api.DataSource, id string, projectId string) (*model.DataSource, error) {
	encodedTags, err := model.MapStrEncode(ds.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding data source's tags")
	}

	encodedFMs, err := model.MapStrEncode(ds.FieldMapping)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding data source's field mapping")
	}

	encodedBS, err := json.Marshal(ds.BatchSource)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding data source's batch source")
	}

	encodedOptions, err := json.Marshal(ds.Options)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding data source's options")
	}

	return &model.DataSource{
		Id:                  id,
		ProjectId:           projectId,
		Name:                ds.Name,
		Description:         ds.Description,
		Tags:                encodedTags,
		Owner:               ds.Owner,
		Type:                int64(ds.Type),
		FieldMapping:        encodedFMs,
		TimestampField:      ds.TimestampField,
		DatePartitionCol:    ds.DatePartitionColumn,
		CreatedTimestampCol: ds.CreatedTimestampColumn,
		ClassType:           ds.DataSourceClassType,
		BatchSource:         encodedBS,
		Options:             encodedOptions,
		ProjectName:         ds.Project,
	}, nil
}

func (r *ResourceManager) ToModelEntity(entity *api.Entity, id string, projectId string) (*model.Entity, error) {
	encodedTags, err := model.MapStrEncode(entity.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding entity's tags")
	}

	return &model.Entity{
		Id:                   id,
		ProjectId:            projectId,
		Name:                 entity.Name,
		ValueType:            int64(entity.ValueType),
		Description:          entity.Description,
		JoinKey:              entity.JoinKey,
		Tags:                 encodedTags,
		Owner:                entity.Owner,
		CreatedTimestamp:     entity.CreatedTimestamp.AsTime(),
		LastUpdatedTimestamp: entity.LastUpdatedTimestamp.AsTime(),
		ProjectName:          entity.Project,
	}, nil
}

func (r *ResourceManager) ToModelFeatureService(fs *api.FeatureService, id string, projectId string) (*model.FeatureService, error) {
	encodedTags, err := model.MapStrEncode(fs.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature service's tags")
	}

	encodedLC, err := json.Marshal(fs.LoggingConfig)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature service's logging config")
	}

	fsModel := &model.FeatureService{
		Id:                   id,
		ProjectId:            projectId,
		Name:                 fs.Name,
		Tags:                 encodedTags,
		Description:          fs.Description,
		Owner:                fs.Owner,
		LoggingConfig:        encodedLC,
		CreatedTimestamp:     fs.CreatedTimestamp.AsTime(),
		LastUpdatedTimestamp: fs.LastUpdatedTimestamp.AsTime(),
		ProjectName:          fs.Project,
	}

	for _, feature := range fs.Features {
		fvpModel, err := r.ToModelFVP(feature, "", fsModel.Id)
		if err != nil {
			return nil, err
		}
		fsModel.FeatureViewProjections = append(fsModel.FeatureViewProjections, fvpModel)
	}

	return fsModel, nil
}

func (r *ResourceManager) ToModelFVP(fvp *api.FeatureViewProjection, id string, fsid string) (*model.FeatureViewProjection, error) {
	encodedJKMap, err := model.MapStrEncode(fvp.JoinKeyMap)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature view projection's join key map")
	}

	fvpModel := &model.FeatureViewProjection{
		Id:          id,
		FSId:        fsid,
		FVName:      fvp.FeatureViewName,
		FVNameAlias: fvp.FeatureViewNameAlias,
		JoinKeyMap:  encodedJKMap,
	}

	for _, feature := range fvp.FeatureColumns {
		fModel, err := r.ToModelFVPFeature(feature, "", fvpModel.Id)
		if err != nil {
			return nil, err
		}
		fvpModel.FVPFeatures = append(fvpModel.FVPFeatures, fModel)
	}

	return fvpModel, nil
}

func (r *ResourceManager) ToModelFVPFeature(f *api.Feature, id string, fvpid string) (*model.FvpFeature, error) {
	encodedTags, err := model.MapStrEncode(f.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature's tags")
	}

	return &model.FvpFeature{
		Id:        id,
		FVPId:     fvpid,
		Name:      f.Name,
		ValueType: int64(f.ValueType),
		Tags:      encodedTags,
	}, nil
}

func (r *ResourceManager) ToModelFeatureView(fv *api.FeatureView, id string, projectId string) (*model.FeatureView, error) {
	encodedTags, err := model.MapStrEncode(fv.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature view's tags")
	}

	encodedEntities, err := json.Marshal(fv.Entities)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature view's entities")
	}

	encodedBS, err := json.Marshal(fv.BatchSource)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature view's batch source")
	}

	encodedSS, err := json.Marshal(fv.StreamSource)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature view's stream source")
	}

	fvModel := &model.FeatureView{
		Id:                   id,
		ProjectId:            projectId,
		Name:                 fv.Name,
		Entities:             encodedEntities,
		Description:          fv.Description,
		Tags:                 encodedTags,
		Owner:                fv.Owner,
		Ttl:                  fv.Ttl.AsDuration(),
		BatchSource:          encodedBS,
		StreamSource:         encodedSS,
		Online:               fv.Online,
		CreatedTimestamp:     fv.CreatedTimestamp.AsTime(),
		LastUpdatedTimestamp: fv.LastUpdatedTimestamp.AsTime(),
		ProjectName:          fv.Project,
	}

	for _, feature := range fv.Features {
		fModel, err := r.ToModelFeature(feature, "", fvModel.Id)
		if err != nil {
			return nil, err
		}
		fvModel.Features = append(fvModel.Features, fModel)
	}

	for _, mi := range fv.MaterializationIntervals {
		miModel, err := r.ToModelMI(mi, fvModel.Id)
		if err != nil {
			return nil, err
		}
		fvModel.MIs = append(fvModel.MIs, miModel)
	}

	return fvModel, nil
}

func (r *ResourceManager) ToModelFeature(f *api.Feature, id string, fvid string) (*model.Feature, error) {
	encodedTags, err := model.MapStrEncode(f.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature's tags")
	}

	return &model.Feature{
		Id:        id,
		FVId:      fvid,
		Name:      f.Name,
		ValueType: int64(f.ValueType),
		Tags:      encodedTags,
	}, nil
}

func (r *ResourceManager) ToModelMI(mi *api.MaterializationInterval, fvid string) (*model.MaterializationInterval, error) {
	return &model.MaterializationInterval{
		FVId:      fvid,
		StartTime: mi.StartTime.AsTime(),
		EndTime:   mi.EndTime.AsTime(),
	}, nil
}

func (r *ResourceManager) ToModelInfraObject(io *api.InfraObject, id string, projectId string, projectName string) (*model.InfraObject, error) {
	encodedObject, err := json.Marshal(io.InfraObject)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding infra object's object")
	}

	return &model.InfraObject{
		Id:          id,
		ProjectId:   projectId,
		ClassType:   io.InfraObjectClassType,
		Object:      encodedObject,
		ProjectName: projectName,
	}, nil
}

func (r *ResourceManager) ToModelOnDemandFeatureView(odfv *api.OnDemandFeatureView, id string, projectId string) (*model.OnDemandFeatureView, error) {
	encodedTags, err := model.MapStrEncode(odfv.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding on demand feature view's tags")
	}

	encodedSources, err := model.MapStrEncode(odfv.Sources)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding on demand feature view's sources")
	}

	odfvModel := &model.OnDemandFeatureView{
		Id:                   id,
		ProjectId:            projectId,
		Name:                 odfv.Name,
		Sources:              encodedSources,
		UdfName:              odfv.UserDefinedFunction.Name,
		UdfBody:              odfv.UserDefinedFunction.Body,
		Description:          odfv.Description,
		Tags:                 encodedTags,
		Owner:                odfv.Owner,
		CreatedTimestamp:     odfv.CreatedTimestamp.AsTime(),
		LastUpdatedTimestamp: odfv.LastUpdatedTimestamp.AsTime(),
		ProjectName:          odfv.Project,
	}

	for _, feature := range odfv.Features {
		odfModel, err := r.ToModelOnDemandFeature(feature, "", odfvModel.Id)
		if err != nil {
			return nil, err
		}
		odfvModel.Features = append(odfvModel.Features, odfModel)
	}

	return odfvModel, nil
}

func (r *ResourceManager) ToModelOnDemandFeature(f *api.Feature, id string, odfvid string) (*model.OnDemandFeature, error) {
	encodedTags, err := model.MapStrEncode(f.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding feature's tags")
	}

	return &model.OnDemandFeature{
		Id:        id,
		ODFVId:    odfvid,
		Name:      f.Name,
		ValueType: int64(f.ValueType),
		Tags:      encodedTags,
	}, nil
}

func (r *ResourceManager) ToModelProject(project *api.Project, id string) (*model.Project, error) {
	return &model.Project{
		Id:                    id,
		Name:                  project.Name,
		RegistrySchemaVersion: project.RegistrySchemaVersion,
		VersionId:             project.VersionId,
		LastUpdated:           project.LastUpdated.AsTime(),
	}, nil
}

func (r *ResourceManager) ToModelRequestFeatureView(rfv *api.RequestFeatureView, id string, projectId string) (*model.RequestFeatureView, error) {
	encodedTags, err := model.MapStrEncode(rfv.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding request feature view's tags")
	}

	encodedDS, err := json.Marshal(rfv.RequestDataSource)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding request feature view's data source")
	}

	return &model.RequestFeatureView{
		Id:          id,
		ProjectId:   projectId,
		Name:        rfv.Name,
		DataSource:  encodedDS,
		Description: rfv.Description,
		Tags:        encodedTags,
		Owner:       rfv.Owner,
		ProjectName: rfv.Project,
	}, nil
}

func (r *ResourceManager) ToModelSavedDataset(sd *api.SavedDataset, id string, projectId string) (*model.SavedDataset, error) {
	encodedTags, err := model.MapStrEncode(sd.Tags)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding saved dataset's tags")
	}

	encodedFeatures, err := json.Marshal(sd.Features)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding saved dataset's features")
	}

	encodedJoinKeys, err := json.Marshal(sd.JoinKeys)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding saved dataset's join keys")
	}

	encodedStorage, err := json.Marshal(sd.Storage)
	if err != nil {
		return nil, util.Wrap(err, "Error encoding saved dataset's storage")
	}

	return &model.SavedDataset{
		Id:                   id,
		ProjectId:            projectId,
		Name:                 sd.Name,
		Features:             encodedFeatures,
		JoinKeys:             encodedJoinKeys,
		FullFeatureNames:     sd.FullFeatureNames,
		Storage:              encodedStorage,
		FeatureServiceName:   sd.FeatureServiceName,
		Tags:                 encodedTags,
		CreatedTimestamp:     sd.CreatedTimestamp.AsTime(),
		LastUpdatedTimestamp: sd.LastUpdatedTimestamp.AsTime(),
		MinEventTimestamp:    sd.MinEventTimestamp.AsTime(),
		MaxEventTimestamp:    sd.MaxEventTimestamp.AsTime(),
		ProjectName:          sd.Project,
	}, nil
}
