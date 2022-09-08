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
	"context"
	"errors"

	api "github.com/feast-dev/feast/backend/api/go_client"
	frsauth "github.com/feast-dev/feast/backend/src/apiserver/auth"
	"github.com/feast-dev/feast/backend/src/apiserver/client"
	"github.com/feast-dev/feast/backend/src/apiserver/model"
	"github.com/feast-dev/feast/backend/src/apiserver/storage"

	util "github.com/feast-dev/feast/backend/src/utils"
	authorizationv1 "k8s.io/api/authorization/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

type ClientManagerInterface interface {
	Authenticators() []frsauth.Authenticator
	DataSourceStore() storage.DataSourceStoreInterface
	EntityStore() storage.EntityStoreInterface
	FeatureViewStore() storage.FeatureViewStoreInterface
	FeatureServiceStore() storage.FeatureServiceStoreInterface
	InfraObjectStore() storage.InfraObjectStoreInterface
	MiStore() storage.MIStoreInterface
	OnDemandFeatureViewStore() storage.OnDemandFeatureViewStoreInterface
	ProjectStore() storage.ProjectStoreInterface
	RequestFeatureViewStore() storage.RequestFeatureViewStoreInterface
	SavedDatasetStore() storage.SavedDatasetStoreInterface
	SubjectAccessReviewClient() client.SubjectAccessReviewInterface
	Time() util.TimeInterface
	TokenReviewClient() client.TokenReviewInterface
	UUID() util.UUIDGeneratorInterface
}

type ResourceManager struct {
	authenticators            []frsauth.Authenticator
	dataSourceStore           storage.DataSourceStoreInterface
	entityStore               storage.EntityStoreInterface
	featureViewStore          storage.FeatureViewStoreInterface
	featureServiceStore       storage.FeatureServiceStoreInterface
	infraObjectStore          storage.InfraObjectStoreInterface
	miStore                   storage.MIStoreInterface
	odFeatureViewStore        storage.OnDemandFeatureViewStoreInterface
	projectStore              storage.ProjectStoreInterface
	requestFeatureViewStore   storage.RequestFeatureViewStoreInterface
	savedDatasetStore         storage.SavedDatasetStoreInterface
	subjectAccessReviewClient client.SubjectAccessReviewInterface
	time                      util.TimeInterface
	tokenReviewClient         client.TokenReviewInterface
	uuid                      util.UUIDGeneratorInterface
}

func NewResourceManager(clientManager ClientManagerInterface) *ResourceManager {
	return &ResourceManager{
		authenticators:            clientManager.Authenticators(),
		dataSourceStore:           clientManager.DataSourceStore(),
		entityStore:               clientManager.EntityStore(),
		featureViewStore:          clientManager.FeatureViewStore(),
		featureServiceStore:       clientManager.FeatureServiceStore(),
		infraObjectStore:          clientManager.InfraObjectStore(),
		miStore:                   clientManager.MiStore(),
		odFeatureViewStore:        clientManager.OnDemandFeatureViewStore(),
		projectStore:              clientManager.ProjectStore(),
		requestFeatureViewStore:   clientManager.RequestFeatureViewStore(),
		savedDatasetStore:         clientManager.SavedDatasetStore(),
		subjectAccessReviewClient: clientManager.SubjectAccessReviewClient(),
		time:                      clientManager.Time(),
		tokenReviewClient:         clientManager.TokenReviewClient(),
		uuid:                      clientManager.UUID(),
	}
}

func (r *ResourceManager) CreateDataSource(apiDataSource *api.DataSource) (*model.DataSource, error) {
	project, err := r.projectStore.GetProject(apiDataSource.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	data_source, err := r.ToModelDataSource(apiDataSource, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert data source model")
	}

	return r.dataSourceStore.CreateDataSource(data_source)
}

func (r *ResourceManager) GetDataSource(name string, project string) (*model.DataSource, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	dsModel, err := r.dataSourceStore.GetDataSource(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	dsModel.ProjectName = project

	return dsModel, nil
}

func (r *ResourceManager) UpdateDataSource(apiDataSource *api.DataSource) (*model.DataSource, error) {
	project, err := r.projectStore.GetProject(apiDataSource.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	data_source, err := r.ToModelDataSource(apiDataSource, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert data source model")
	}

	_, err = r.dataSourceStore.GetDataSource(data_source.Name, data_source.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update data source failed")
	}

	return r.dataSourceStore.UpdateDataSource(data_source)
}

func (r *ResourceManager) DeleteDataSource(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.dataSourceStore.GetDataSource(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete data source failed")
	}

	return r.dataSourceStore.DeleteDataSource(name, projectModel.Id)
}

func (r *ResourceManager) ListDataSources(project string) ([]*model.DataSource, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	dsModels, err := r.dataSourceStore.ListDataSources(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, dsModel := range dsModels {
		dsModel.ProjectName = project
	}

	return dsModels, nil
}

func (r *ResourceManager) CreateEntity(apiEntity *api.Entity) (*model.Entity, error) {
	project, err := r.projectStore.GetProject(apiEntity.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	entity, err := r.ToModelEntity(apiEntity, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert entity model")
	}

	return r.entityStore.CreateEntity(entity)
}

func (r *ResourceManager) GetEntity(name string, project string) (*model.Entity, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	entityModel, err := r.entityStore.GetEntity(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	entityModel.ProjectName = project

	return entityModel, nil
}

func (r *ResourceManager) UpdateEntity(apiEntity *api.Entity) (*model.Entity, error) {
	project, err := r.projectStore.GetProject(apiEntity.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	entity, err := r.ToModelEntity(apiEntity, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert entity model")
	}

	_, err = r.entityStore.GetEntity(entity.Name, entity.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update entity failed")
	}

	return r.entityStore.UpdateEntity(entity)
}

func (r *ResourceManager) DeleteEntity(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.entityStore.GetEntity(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete entity failed")
	}

	return r.entityStore.DeleteEntity(name, projectModel.Id)
}

func (r *ResourceManager) ListEntities(project string) ([]*model.Entity, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	entityModels, err := r.entityStore.ListEntities(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, entityModel := range entityModels {
		entityModel.ProjectName = project
	}

	return entityModels, nil
}

func (r *ResourceManager) CreateFeatureService(apiFeatureService *api.FeatureService) (*model.FeatureService, error) {
	project, err := r.projectStore.GetProject(apiFeatureService.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	feature_service, err := r.ToModelFeatureService(apiFeatureService, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert feature service model")
	}

	return r.featureServiceStore.CreateFeatureService(feature_service)
}

func (r *ResourceManager) GetFeatureService(name string, project string) (*model.FeatureService, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	featureServiceModel, err := r.featureServiceStore.GetFeatureService(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	featureServiceModel.ProjectName = project

	return featureServiceModel, nil
}

func (r *ResourceManager) UpdateFeatureService(apiFeatureService *api.FeatureService) (*model.FeatureService, error) {
	project, err := r.projectStore.GetProject(apiFeatureService.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	feature_service, err := r.ToModelFeatureService(apiFeatureService, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert feature service model")
	}

	oldFs, err := r.featureServiceStore.GetFeatureService(feature_service.Name, feature_service.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update feature service failed")
	}
	feature_service.Id = oldFs.Id

	return r.featureServiceStore.UpdateFeatureService(feature_service)
}

func (r *ResourceManager) DeleteFeatureService(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.featureServiceStore.GetFeatureService(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete feature service failed")
	}

	return r.featureServiceStore.DeleteFeatureService(name, projectModel.Id)
}

func (r *ResourceManager) ListFeatureServices(project string) ([]*model.FeatureService, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	featureServiceModels, err := r.featureServiceStore.ListFeatureServices(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, featureServiceModel := range featureServiceModels {
		featureServiceModel.ProjectName = project
	}

	return featureServiceModels, nil
}

func (r *ResourceManager) CreateFeatureView(apiFeatureView *api.FeatureView) (*model.FeatureView, error) {
	project, err := r.projectStore.GetProject(apiFeatureView.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	fv, err := r.ToModelFeatureView(apiFeatureView, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert feature view model")
	}

	return r.featureViewStore.CreateFeatureView(fv)
}

func (r *ResourceManager) GetFeatureView(name string, project string) (*model.FeatureView, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	fvModel, err := r.featureViewStore.GetFeatureView(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	fvModel.ProjectName = project

	return fvModel, nil
}

func (r *ResourceManager) UpdateFeatureView(apiFeatureView *api.FeatureView) (*model.FeatureView, error) {
	project, err := r.projectStore.GetProject(apiFeatureView.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	fv, err := r.ToModelFeatureView(apiFeatureView, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert feature view model")
	}

	oldFv, err := r.featureViewStore.GetFeatureView(fv.Name, fv.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update feature view failed")
	}
	fv.Id = oldFv.Id

	return r.featureViewStore.UpdateFeatureView(fv)
}

func (r *ResourceManager) DeleteFeatureView(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.featureViewStore.GetFeatureView(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete feature view failed")
	}

	return r.featureViewStore.DeleteFeatureView(name, projectModel.Id)
}

func (r *ResourceManager) ListFeatureViews(project string) ([]*model.FeatureView, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	fvModels, err := r.featureViewStore.ListFeatureViews(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, fvModel := range fvModels {
		fvModel.ProjectName = project
	}

	return fvModels, nil
}

func (r *ResourceManager) UpdateInfraObjects(apiInfraObjects []*api.InfraObject, project string) ([]*model.InfraObject, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	var infraObjectModels []*model.InfraObject
	for _, apiInfraObject := range apiInfraObjects {
		infraObjectModel, err := r.ToModelInfraObject(apiInfraObject, "", projectModel.Id, projectModel.Name)
		if err != nil {
			return nil, util.Wrap(err, "Failed to convert infra object model")
		}
		infraObjectModel.ProjectName = project
		infraObjectModels = append(infraObjectModels, infraObjectModel)
	}

	err = r.infraObjectStore.DeleteInfraObjects(projectModel.Id)
	if err != nil {
		return nil, util.Wrap(err, "Update infra objects failed")
	}

	for _, infraObjectModel := range infraObjectModels {
		_, err := r.infraObjectStore.CreateInfraObject(infraObjectModel)
		if err != nil {
			return nil, util.Wrap(err, "Update infra objects failed")
		}
	}
	return infraObjectModels, nil
}

func (r *ResourceManager) ListInfraObjects(project string) ([]*model.InfraObject, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	infraObjectModels, err := r.infraObjectStore.ListInfraObjects(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, infraObjectModel := range infraObjectModels {
		infraObjectModel.ProjectName = project
	}

	return infraObjectModels, nil
}

func (r *ResourceManager) ReportMI(apiMi *api.MaterializationInterval, name string, project string) (*model.MaterializationInterval, error) {
	mi, err := r.ToModelMI(apiMi, "")
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert materialization interval model")
	}

	fv, err := r.GetFeatureView(name, project)
	if err != nil {
		return nil, util.Wrap(err, "Report materialization interval failed")
	}

	return r.miStore.CreateMINoTx(mi, fv.Id)
}

func (r *ResourceManager) CreateOnDemandFeatureView(apiODFV *api.OnDemandFeatureView) (*model.OnDemandFeatureView, error) {
	project, err := r.projectStore.GetProject(apiODFV.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	on_demand_fv, err := r.ToModelOnDemandFeatureView(apiODFV, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert on demand feature view model")
	}

	return r.odFeatureViewStore.CreateOnDemandFeatureView(on_demand_fv)
}

func (r *ResourceManager) GetOnDemandFeatureView(name string, project string) (*model.OnDemandFeatureView, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	odfvModel, err := r.odFeatureViewStore.GetOnDemandFeatureView(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	odfvModel.ProjectName = project

	return odfvModel, nil
}

func (r *ResourceManager) UpdateOnDemandFeatureView(apiODFV *api.OnDemandFeatureView) (*model.OnDemandFeatureView, error) {
	project, err := r.projectStore.GetProject(apiODFV.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	on_demand_fv, err := r.ToModelOnDemandFeatureView(apiODFV, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert on demand feature view model")
	}

	oldODFV, err := r.odFeatureViewStore.GetOnDemandFeatureView(on_demand_fv.Name, on_demand_fv.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update on demand feature view failed")
	}
	on_demand_fv.Id = oldODFV.Id

	return r.odFeatureViewStore.UpdateOnDemandFeatureView(on_demand_fv)
}

func (r *ResourceManager) DeleteOnDemandFeatureView(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.odFeatureViewStore.GetOnDemandFeatureView(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete on demand feature view failed")
	}

	return r.odFeatureViewStore.DeleteOnDemandFeatureView(name, projectModel.Id)
}

func (r *ResourceManager) ListOnDemandFeatureViews(project string) ([]*model.OnDemandFeatureView, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	odfvModels, err := r.odFeatureViewStore.ListOnDemandFeatureViews(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, odfvModel := range odfvModels {
		odfvModel.ProjectName = project
	}

	return odfvModels, nil
}

func (r *ResourceManager) CreateProject(apiProject *api.Project) (*model.Project, error) {
	project, err := r.ToModelProject(apiProject, "")
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert project model")
	}

	return r.projectStore.CreateProject(project)
}

func (r *ResourceManager) GetProject(project string) (*model.Project, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, err
	}

	return projectModel, nil
}

func (r *ResourceManager) UpdateProject(apiProject *api.Project) (*model.Project, error) {
	project, err := r.ToModelProject(apiProject, "")
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert project model")
	}

	_, err = r.projectStore.GetProject(project.Name)
	if err != nil {
		return nil, util.Wrap(err, "Update project failed")
	}

	return r.projectStore.UpdateProject(project)
}

func (r *ResourceManager) DeleteProject(project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Delete project failed")
	}

	entities, err := r.entityStore.ListEntities(projectModel.Id)
	if err != nil {
		return err
	}
	data_sources, err := r.dataSourceStore.ListDataSources(projectModel.Id)
	if err != nil {
		return err
	}
	feature_services, err := r.featureServiceStore.ListFeatureServices(projectModel.Id)
	if err != nil {
		return err
	}
	feature_views, err := r.featureViewStore.ListFeatureViews(projectModel.Id)
	if err != nil {
		return err
	}
	on_demand_feature_views, err := r.odFeatureViewStore.ListOnDemandFeatureViews(projectModel.Id)
	if err != nil {
		return err
	}
	request_feature_views, err := r.requestFeatureViewStore.ListRequestFeatureViews(projectModel.Id)
	if err != nil {
		return err
	}
	saved_datasets, err := r.savedDatasetStore.ListSavedDatasets(projectModel.Id)
	if err != nil {
		return err
	}

	for _, e := range entities {
		err = r.entityStore.DeleteEntity(e.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}
	for _, ds := range data_sources {
		err = r.dataSourceStore.DeleteDataSource(ds.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}
	for _, fs := range feature_services {
		err = r.featureServiceStore.DeleteFeatureService(fs.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}
	for _, fv := range feature_views {
		err = r.featureViewStore.DeleteFeatureView(fv.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}
	err = r.infraObjectStore.DeleteInfraObjects(projectModel.Id)
	if err != nil {
		return err
	}
	for _, odfv := range on_demand_feature_views {
		err = r.odFeatureViewStore.DeleteOnDemandFeatureView(odfv.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}
	for _, rfv := range request_feature_views {
		err = r.requestFeatureViewStore.DeleteRequestFeatureView(rfv.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}
	for _, sd := range saved_datasets {
		err = r.savedDatasetStore.DeleteSavedDataset(sd.Name, projectModel.Id)
		if err != nil {
			return err
		}
	}

	return r.projectStore.DeleteProject(project)
}

func (r *ResourceManager) CreateRequestFeatureView(apiRequestFV *api.RequestFeatureView) (*model.RequestFeatureView, error) {
	project, err := r.projectStore.GetProject(apiRequestFV.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	request_fv, err := r.ToModelRequestFeatureView(apiRequestFV, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert request feature view model")
	}

	return r.requestFeatureViewStore.CreateRequestFeatureView(request_fv)
}

func (r *ResourceManager) GetRequestFeatureView(name string, project string) (*model.RequestFeatureView, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	rfvModel, err := r.requestFeatureViewStore.GetRequestFeatureView(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	rfvModel.ProjectName = project

	return rfvModel, nil
}

func (r *ResourceManager) UpdateRequestFeatureView(apiRequestFV *api.RequestFeatureView) (*model.RequestFeatureView, error) {
	project, err := r.projectStore.GetProject(apiRequestFV.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	request_fv, err := r.ToModelRequestFeatureView(apiRequestFV, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert request feature view model")
	}

	_, err = r.requestFeatureViewStore.GetRequestFeatureView(request_fv.Name, request_fv.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update request feature view failed")
	}

	return r.requestFeatureViewStore.UpdateRequestFeatureView(request_fv)
}

func (r *ResourceManager) DeleteRequestFeatureView(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.requestFeatureViewStore.GetRequestFeatureView(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete request feature view failed")
	}

	return r.requestFeatureViewStore.DeleteRequestFeatureView(name, projectModel.Id)
}

func (r *ResourceManager) ListRequestFeatureViews(project string) ([]*model.RequestFeatureView, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	rfvModels, err := r.requestFeatureViewStore.ListRequestFeatureViews(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, rfvModel := range rfvModels {
		rfvModel.ProjectName = project
	}

	return rfvModels, nil
}

func (r *ResourceManager) CreateSavedDataset(apiSavedDataset *api.SavedDataset) (*model.SavedDataset, error) {
	project, err := r.projectStore.GetProject(apiSavedDataset.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	saved_dataset, err := r.ToModelSavedDataset(apiSavedDataset, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert saved dataset model")
	}

	return r.savedDatasetStore.CreateSavedDataset(saved_dataset)
}

func (r *ResourceManager) GetSavedDataset(name string, project string) (*model.SavedDataset, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	savedDatasetModel, err := r.savedDatasetStore.GetSavedDataset(name, projectModel.Id)
	if err != nil {
		return nil, err
	}
	savedDatasetModel.ProjectName = project

	return savedDatasetModel, nil
}

func (r *ResourceManager) UpdateSavedDataset(apiSavedDataset *api.SavedDataset) (*model.SavedDataset, error) {
	project, err := r.projectStore.GetProject(apiSavedDataset.Project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	savedDataset, err := r.ToModelSavedDataset(apiSavedDataset, "", project.Id)
	if err != nil {
		return nil, util.Wrap(err, "Failed to convert saved dataset model")
	}

	_, err = r.savedDatasetStore.GetSavedDataset(savedDataset.Name, savedDataset.ProjectId)
	if err != nil {
		return nil, util.Wrap(err, "Update saved dataset failed")
	}

	return r.savedDatasetStore.UpdateSavedDataset(savedDataset)
}

func (r *ResourceManager) DeleteSavedDataset(name string, project string) error {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return util.Wrap(err, "Failed to find project")
	}

	_, err = r.savedDatasetStore.GetSavedDataset(name, projectModel.Id)
	if err != nil {
		return util.Wrap(err, "Delete saved dataset failed")
	}

	return r.savedDatasetStore.DeleteSavedDataset(name, projectModel.Id)
}

func (r *ResourceManager) ListSavedDatasets(project string) ([]*model.SavedDataset, error) {
	projectModel, err := r.projectStore.GetProject(project)
	if err != nil {
		return nil, util.Wrap(err, "Failed to find project")
	}

	savedDatasetModels, err := r.savedDatasetStore.ListSavedDatasets(projectModel.Id)
	if err != nil {
		return nil, err
	}

	for _, savedDatasetModel := range savedDatasetModels {
		savedDatasetModel.ProjectName = project
	}

	return savedDatasetModels, nil
}

func (r *ResourceManager) AuthenticateRequest(ctx context.Context) (string, []string, error) {
	if ctx == nil {
		return "", make([]string, 0), util.NewUnauthenticatedError(errors.New("request error: context is nil"), "Request error: context is nil.")
	}

	// If the request header contains the user identity, requests are authorized
	// based on the namespace field in the request.
	var errlist []error
	for _, auth := range r.authenticators {
		userIdentity, userGroups, err := auth.GetUserIdentity(ctx)
		if err == nil {
			return userIdentity, userGroups, nil
		}
		errlist = append(errlist, err)
	}

	return "", make([]string, 0), utilerrors.NewAggregate(errlist)
}

func (r *ResourceManager) IsRequestAuthorized(ctx context.Context, userIdentity string, userGroups []string, resourceAttributes *authorizationv1.ResourceAttributes) error {
	result, err := r.subjectAccessReviewClient.Create(
		ctx,
		&authorizationv1.SubjectAccessReview{
			Spec: authorizationv1.SubjectAccessReviewSpec{
				ResourceAttributes: resourceAttributes,
				User:               userIdentity,
				Groups:             userGroups,
			},
		},
		v1.CreateOptions{},
	)
	if err != nil {
		return util.NewInternalServerError(
			err,
			"Failed to create SubjectAccessReview for user '%s' (request: %+v)",
			userIdentity,
			resourceAttributes,
		)
	}
	if !result.Status.Allowed {
		return util.NewPermissionDeniedError(
			errors.New("unauthorized access"),
			"User '%s' is not authorized with reason: %s (request: %+v)",
			userIdentity,
			result.Status.Reason,
			resourceAttributes,
		)
	}

	return nil
}
