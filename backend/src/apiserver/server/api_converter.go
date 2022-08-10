package server

import (
	"encoding/json"

	"github.com/feast-dev/feast/backend/src/apiserver/model"

	api "github.com/feast-dev/feast/backend/api/go_client"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ToApiDataSource(ds *model.DataSource) *api.DataSource {
	decodedTags, _ := model.MapStrDecode(ds.Tags)
	decodedFMs, _ := model.MapStrDecode(ds.FieldMapping)

	var decodedBS, decodedOptions string
	err := json.Unmarshal(ds.BatchSource, &decodedBS)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(ds.Options, &decodedOptions)
	if err != nil {
		return nil
	}

	return &api.DataSource{
		Name:                   ds.Name,
		Project:                ds.ProjectName,
		Tags:                   decodedTags,
		Owner:                  ds.Owner,
		Type:                   api.DataSource_SourceType(ds.Type),
		FieldMapping:           decodedFMs,
		TimestampField:         ds.TimestampField,
		DatePartitionColumn:    ds.DatePartitionCol,
		CreatedTimestampColumn: ds.CreatedTimestampCol,
		DataSourceClassType:    ds.ClassType,
		BatchSource:            decodedBS,
		Options:                decodedOptions,
	}
}

func ToApiDataSources(data_sources []*model.DataSource) []*api.DataSource {
	apiDataSources := make([]*api.DataSource, 0)
	for _, ds := range data_sources {
		apiDataSources = append(apiDataSources, ToApiDataSource(ds))
	}

	return apiDataSources
}

func ToApiEntity(entity *model.Entity) *api.Entity {
	decodedTags, _ := model.MapStrDecode(entity.Tags)

	return &api.Entity{
		Name:                 entity.Name,
		Project:              entity.ProjectName,
		ValueType:            api.ValueType_Enum(entity.ValueType),
		Description:          entity.Description,
		JoinKey:              entity.JoinKey,
		Tags:                 decodedTags,
		Owner:                entity.Owner,
		CreatedTimestamp:     &timestamp.Timestamp{Seconds: entity.CreatedTimestamp.Unix()},
		LastUpdatedTimestamp: &timestamp.Timestamp{Seconds: entity.LastUpdatedTimestamp.Unix()},
	}
}

func ToApiEntities(entities []*model.Entity) []*api.Entity {
	apiEntities := make([]*api.Entity, 0)
	for _, entity := range entities {
		apiEntities = append(apiEntities, ToApiEntity(entity))
	}

	return apiEntities
}

func ToApiFeatureService(fs *model.FeatureService) *api.FeatureService {
	decodedTags, _ := model.MapStrDecode(fs.Tags)

	var decodedLC string
	err := json.Unmarshal(fs.LoggingConfig, &decodedLC)
	if err != nil {
		return nil
	}

	return &api.FeatureService{
		Name:                 fs.Name,
		Project:              fs.ProjectName,
		Features:             ToApiFVPs(fs.FeatureViewProjections),
		Tags:                 decodedTags,
		Description:          fs.Description,
		Owner:                fs.Owner,
		LoggingConfig:        decodedLC,
		CreatedTimestamp:     &timestamp.Timestamp{Seconds: fs.CreatedTimestamp.Unix()},
		LastUpdatedTimestamp: &timestamp.Timestamp{Seconds: fs.LastUpdatedTimestamp.Unix()},
	}
}

func ToApiFeatureServices(fss []*model.FeatureService) []*api.FeatureService {
	apiFeatureServices := make([]*api.FeatureService, 0)
	for _, fs := range fss {
		apiFeatureServices = append(apiFeatureServices, ToApiFeatureService(fs))
	}

	return apiFeatureServices
}

func ToApiFVP(fvp *model.FeatureViewProjection) *api.FeatureViewProjection {
	decodedJKM, _ := model.MapStrDecode(fvp.JoinKeyMap)

	return &api.FeatureViewProjection{
		FeatureViewName:      fvp.FVName,
		FeatureViewNameAlias: fvp.FVNameAlias,
		FeatureColumns:       ToApiFVPFeatures(fvp.FVPFeatures),
		JoinKeyMap:           decodedJKM,
	}
}

func ToApiFVPs(fvps []*model.FeatureViewProjection) []*api.FeatureViewProjection {
	apiFVPs := make([]*api.FeatureViewProjection, 0)
	for _, fvp := range fvps {
		apiFVPs = append(apiFVPs, ToApiFVP(fvp))
	}

	return apiFVPs
}

func ToApiFVPFeature(f *model.FvpFeature) *api.Feature {
	decodedTags, _ := model.MapStrDecode(f.Tags)

	return &api.Feature{
		Name:      f.Name,
		ValueType: api.ValueType_Enum(f.ValueType),
		Tags:      decodedTags,
	}
}

func ToApiFVPFeatures(fs []*model.FvpFeature) []*api.Feature {
	apiFeatures := make([]*api.Feature, 0)
	for _, f := range fs {
		apiFeatures = append(apiFeatures, ToApiFVPFeature(f))
	}

	return apiFeatures
}

func ToApiFeatureView(fv *model.FeatureView) *api.FeatureView {
	decodedTags, _ := model.MapStrDecode(fv.Tags)

	var decodedEntities []string
	var decodedBS, decodedSS string
	err := json.Unmarshal(fv.Entities, &decodedEntities)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(fv.BatchSource, &decodedBS)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(fv.StreamSource, &decodedSS)
	if err != nil {
		return nil
	}

	return &api.FeatureView{
		Name:                     fv.Name,
		Project:                  fv.ProjectName,
		Entities:                 decodedEntities,
		Features:                 ToApiFeatures(fv.Features),
		Description:              fv.Description,
		Tags:                     decodedTags,
		Owner:                    fv.Owner,
		Ttl:                      &duration.Duration{Seconds: int64(fv.Ttl.Seconds())},
		BatchSource:              decodedBS,
		StreamSource:             decodedSS,
		Online:                   fv.Online,
		CreatedTimestamp:         &timestamp.Timestamp{Seconds: fv.CreatedTimestamp.Unix()},
		LastUpdatedTimestamp:     &timestamp.Timestamp{Seconds: fv.LastUpdatedTimestamp.Unix()},
		MaterializationIntervals: ToApiMIs(fv.MIs),
	}
}

func ToApiFeatureViews(fvs []*model.FeatureView) []*api.FeatureView {
	apiFeatureViews := make([]*api.FeatureView, 0)
	for _, fv := range fvs {
		apiFeatureViews = append(apiFeatureViews, ToApiFeatureView(fv))
	}

	return apiFeatureViews
}

func ToApiFeature(f *model.Feature) *api.Feature {
	decodedTags, _ := model.MapStrDecode(f.Tags)

	return &api.Feature{
		Name:      f.Name,
		ValueType: api.ValueType_Enum(f.ValueType),
		Tags:      decodedTags,
	}
}

func ToApiFeatures(fs []*model.Feature) []*api.Feature {
	apiFeatures := make([]*api.Feature, 0)
	for _, f := range fs {
		apiFeatures = append(apiFeatures, ToApiFeature(f))
	}

	return apiFeatures
}

func ToApiMI(mi *model.MaterializationInterval) *api.MaterializationInterval {
	return &api.MaterializationInterval{
		StartTime: &timestamp.Timestamp{Seconds: mi.StartTime.Unix()},
		EndTime:   &timestamp.Timestamp{Seconds: mi.EndTime.Unix()},
	}
}

func ToApiMIs(mis []*model.MaterializationInterval) []*api.MaterializationInterval {
	apiMIs := make([]*api.MaterializationInterval, 0)
	for _, mi := range mis {
		apiMIs = append(apiMIs, ToApiMI(mi))
	}

	return apiMIs
}

func ToApiInfraObject(io *model.InfraObject) *api.InfraObject {
	var decodedObject string
	err := json.Unmarshal(io.Object, &decodedObject)
	if err != nil {
		return nil
	}

	return &api.InfraObject{
		InfraObjectClassType: io.ClassType,
		InfraObject:          decodedObject,
	}
}

func ToApiInfraObjects(infra_objects []*model.InfraObject) []*api.InfraObject {
	apiInfraObjects := make([]*api.InfraObject, 0)
	for _, io := range infra_objects {
		apiInfraObjects = append(apiInfraObjects, ToApiInfraObject(io))
	}

	return apiInfraObjects
}

func ToApiOnDemandFeatureView(odfv *model.OnDemandFeatureView) *api.OnDemandFeatureView {
	decodedTags, _ := model.MapStrDecode(odfv.Tags)

	decodedSources, _ := model.MapStrDecode(odfv.Sources)

	return &api.OnDemandFeatureView{
		Name:     odfv.Name,
		Project:  odfv.ProjectName,
		Features: ToApiODFeatures(odfv.Features),
		Sources:  decodedSources,
		UserDefinedFunction: &api.UserDefinedFunction{
			Name: odfv.UdfName,
			Body: odfv.UdfBody,
		},
		Description:          odfv.Description,
		Tags:                 decodedTags,
		Owner:                odfv.Owner,
		CreatedTimestamp:     &timestamp.Timestamp{Seconds: odfv.CreatedTimestamp.Unix()},
		LastUpdatedTimestamp: &timestamp.Timestamp{Seconds: odfv.LastUpdatedTimestamp.Unix()},
	}
}

func ToApiOnDemandFeatureViews(odfvs []*model.OnDemandFeatureView) []*api.OnDemandFeatureView {
	apiOnDemandFeatureViews := make([]*api.OnDemandFeatureView, 0)
	for _, odfv := range odfvs {
		apiOnDemandFeatureViews = append(apiOnDemandFeatureViews, ToApiOnDemandFeatureView(odfv))
	}

	return apiOnDemandFeatureViews
}

func ToApiODFeature(f *model.OnDemandFeature) *api.Feature {
	decodedTags, _ := model.MapStrDecode(f.Tags)

	return &api.Feature{
		Name:      f.Name,
		ValueType: api.ValueType_Enum(f.ValueType),
		Tags:      decodedTags,
	}
}

func ToApiODFeatures(fs []*model.OnDemandFeature) []*api.Feature {
	apiFeatures := make([]*api.Feature, 0)
	for _, f := range fs {
		apiFeatures = append(apiFeatures, ToApiODFeature(f))
	}

	return apiFeatures
}

func ToApiProject(project *model.Project) *api.Project {
	return &api.Project{
		Name:                  project.Name,
		RegistrySchemaVersion: project.RegistrySchemaVersion,
		VersionId:             project.VersionId,
		LastUpdated:           &timestamp.Timestamp{Seconds: project.LastUpdated.Unix()},
	}
}

func ToApiRequestFeatureView(rfv *model.RequestFeatureView) *api.RequestFeatureView {
	decodedTags, _ := model.MapStrDecode(rfv.Tags)

	var decodedDS string
	err := json.Unmarshal(rfv.DataSource, &decodedDS)
	if err != nil {
		return nil
	}

	return &api.RequestFeatureView{
		Name:              rfv.Name,
		Project:           rfv.ProjectName,
		RequestDataSource: decodedDS,
		Description:       rfv.Description,
		Tags:              decodedTags,
		Owner:             rfv.Owner,
	}
}

func ToApiRequestFeatureViews(rfvs []*model.RequestFeatureView) []*api.RequestFeatureView {
	apiRequestFeatureViews := make([]*api.RequestFeatureView, 0)
	for _, rfv := range rfvs {
		apiRequestFeatureViews = append(apiRequestFeatureViews, ToApiRequestFeatureView(rfv))
	}

	return apiRequestFeatureViews
}

func ToApiSavedDataset(sd *model.SavedDataset) *api.SavedDataset {
	decodedTags, _ := model.MapStrDecode(sd.Tags)

	var decodedFeatures, decodedJoinKeys []string
	var decodedStorage string
	err := json.Unmarshal(sd.Features, &decodedFeatures)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(sd.JoinKeys, &decodedJoinKeys)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(sd.Storage, &decodedStorage)
	if err != nil {
		return nil
	}

	return &api.SavedDataset{
		Name:                 sd.Name,
		Project:              sd.ProjectName,
		Features:             decodedFeatures,
		JoinKeys:             decodedJoinKeys,
		FullFeatureNames:     sd.FullFeatureNames,
		Storage:              decodedStorage,
		FeatureServiceName:   sd.FeatureServiceName,
		Tags:                 decodedTags,
		CreatedTimestamp:     &timestamp.Timestamp{Seconds: sd.CreatedTimestamp.Unix()},
		LastUpdatedTimestamp: &timestamp.Timestamp{Seconds: sd.LastUpdatedTimestamp.Unix()},
		MinEventTimestamp:    &timestamp.Timestamp{Seconds: sd.MinEventTimestamp.Unix()},
		MaxEventTimestamp:    &timestamp.Timestamp{Seconds: sd.MaxEventTimestamp.Unix()},
	}
}

func ToApiSavedDatasets(saved_datasets []*model.SavedDataset) []*api.SavedDataset {
	apiSavedDatasets := make([]*api.SavedDataset, 0)
	for _, sd := range saved_datasets {
		apiSavedDatasets = append(apiSavedDatasets, ToApiSavedDataset(sd))
	}

	return apiSavedDatasets
}
