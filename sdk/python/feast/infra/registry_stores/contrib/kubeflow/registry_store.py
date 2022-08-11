import logging
import queue
from datetime import datetime, timezone
from pathlib import Path

import frs_api

from feast.protos.feast.core.Registry_pb2 import Registry as RegistryProto
from feast.registry_store import RegistryStore
from feast.repo_config import load_repo_config
from feast.usage import log_exceptions_and_usage

from .api_converter import (
    data_source_to_api,
    entity_to_api,
    feature_service_to_api,
    feature_view_to_api,
    infra_to_api,
    mi_to_api,
    odfv_to_api,
    project_to_api,
    rfv_to_api,
    saved_dataset_to_api,
)
from .auth import ServiceAccountTokenVolumeCredentials
from .kubeflow_config import KubeflowRegistryConfig
from .proto_converter import (
    data_source_to_proto,
    entity_to_proto,
    feature_service_to_proto,
    feature_view_to_proto,
    infra_to_proto,
    odfv_to_proto,
    rfv_to_proto,
    saved_dataset_to_proto,
)


class KubeflowRegistryStore(RegistryStore):
    def __init__(self, registry_config: KubeflowRegistryConfig, repo_path: Path):
        kf_registry_config = KubeflowRegistryConfig(
            host=getattr(
                registry_config, "host", "feast-registry-svc.kubeflow.svc.cluster.local"
            ),
            port=getattr(registry_config, "port", 8888),
            project=getattr(registry_config, "project", ""),
            readMode=getattr(registry_config, "readMode", False),
        )

        if kf_registry_config.project != "":
            self.project = kf_registry_config.project
        else:
            repo_config = load_repo_config(repo_path)
            self.project = repo_config.project

        if get_pod_namespace() != self.project:
            print(
                """Warning! You are rying to access a different namespace, make sure you have the right permissions."""
            )
        self.readMode = kf_registry_config.readMode

        config = frs_api.configuration.Configuration()
        config.host = f"{kf_registry_config.host}:{kf_registry_config.port}"

        config.api_key["authorization"] = "token"
        config.api_key_prefix["authorization"] = "Bearer"
        config.refresh_api_key_hook = (
            ServiceAccountTokenVolumeCredentials().refresh_api_key_hook
        )

        api_client = frs_api.api_client.ApiClient(configuration=config)

        self._data_source_api = frs_api.api.data_source_service_api.DataSourceServiceApi(
            api_client
        )
        self._entity_api = frs_api.api.entity_service_api.EntityServiceApi(api_client)
        self._feature_service_api = frs_api.api.feature_service_service_api.FeatureServiceServiceApi(
            api_client
        )
        self._feature_view_api = frs_api.api.feature_view_service_api.FeatureViewServiceApi(
            api_client
        )
        self._infra_object_api = frs_api.api.infra_object_service_api.InfraObjectServiceApi(
            api_client
        )
        self._on_demand_fv_api = frs_api.api.on_demand_feature_view_service_api.OnDemandFeatureViewServiceApi(
            api_client
        )
        self._project_api = frs_api.api.project_service_api.ProjectServiceApi(
            api_client
        )
        self._request_fv_api = frs_api.api.request_feature_view_service_api.RequestFeatureViewServiceApi(
            api_client
        )
        self._saved_dataset_api = frs_api.api.saved_dataset_service_api.SavedDatasetServiceApi(
            api_client
        )

    @log_exceptions_and_usage(registry="kubeflow")
    def get_registry_proto(self):
        try:
            if not self.readMode:
                proj = self._project_api.project_service_get_project(
                    project=self.project
                )
        except frs_api.exceptions.ApiException:
            raise FileNotFoundError(
                f'Project named "{self.project}" not found. Have you run "feast apply"?'
            )

        if not self.readMode:
            registry_proto = RegistryProto(
                registry_schema_version=proj.registry_schema_version,
                version_id=proj.version_id,
            )
            registry_proto.last_updated.FromDatetime(
                getattr(proj, "last_updated", datetime(1970, 1, 1, tzinfo=timezone.utc))
            )

            self._get_infra(registry_proto)
        else:
            registry_proto = RegistryProto()

        self._get_entities(registry_proto)
        self._get_data_sources(registry_proto)
        self._get_feature_services(registry_proto)
        self._get_feature_views(registry_proto)
        self._get_on_demand_fvs(registry_proto)
        self._get_request_fvs(registry_proto)
        self._get_saved_datasets(registry_proto)

        return registry_proto

    @log_exceptions_and_usage(registry="kubeflow")
    def update_registry_proto(
        self, registry_proto: RegistryProto, pending_ops: queue.Queue = queue.Queue()
    ):
        while not pending_ops.empty():
            op = pending_ops.get()
            if op["op"] == "CreateProject":
                project = project_to_api(op["proto"], self.project)
                self._project_api.project_service_create_project(body=project)
            elif op["op"].endswith("Entity"):
                self._update_entities(op)
            elif op["op"].endswith("DataSource"):
                self._update_data_sources(op)
            elif op["op"].endswith("FeatureService"):
                self._update_feature_services(op)
            elif op["op"] == "UpdateInfra":
                infra_object = infra_to_api(op["proto"])
                self._infra_object_api.infra_object_service_update_infra_objects(
                    body={"objects": infra_object}, project=self.project
                )
            elif op["op"].endswith("OnDemandFeatureView"):
                self._update_on_demand_fvs(op)
            elif op["op"].endswith("RequestFeatureView"):
                self._update_request_fvs(op)
            elif op["op"].endswith("FeatureView"):
                self._update_feature_views(op)
            elif op["op"].endswith("SavedDataset"):
                self._update_saved_datasets(op)
            elif op["op"] == "ReportMI":
                mi = mi_to_api(op["mi"])
                self._feature_view_api.feature_view_service_report_mi(
                    body=mi, name=op["name"], project=self.project
                )
            else:
                raise ValueError(f"Not supported op: {op['op']}")

        # Update last_updated of project
        project = frs_api.models.ApiProject(
            name=self.project,
            registry_schema_version=registry_proto.registry_schema_version,
            version_id=registry_proto.version_id,
            last_updated=datetime.now().astimezone(timezone.utc),
        )
        self._project_api.project_service_update_project(body=project)

    @log_exceptions_and_usage(registry="kubeflow")
    def teardown(self):
        self._project_api.project_service_delete_project(project=self.project)

    # Get methods
    def _get_entities(self, registry_proto: RegistryProto):
        res = self._entity_api.entity_service_list_entities(project=self.project)

        if not res.entities:
            return

        for e in res.entities:
            entity_proto = entity_to_proto(e)
            registry_proto.entities.append(entity_proto)

    def _get_data_sources(self, registry_proto: RegistryProto):
        res = self._data_source_api.data_source_service_list_data_sources(
            project=self.project
        )

        if not res.data_sources:
            return

        for ds in res.data_sources:
            data_source_proto = data_source_to_proto(ds)
            registry_proto.data_sources.append(data_source_proto)

    def _get_feature_services(self, registry_proto: RegistryProto):
        res = self._feature_service_api.feature_service_service_list_feature_services(
            project=self.project
        )

        if not res.feature_services:
            return

        for fs in res.feature_services:
            fs_proto = feature_service_to_proto(fs)
            registry_proto.feature_services.append(fs_proto)

    def _get_feature_views(self, registry_proto: RegistryProto):
        res = self._feature_view_api.feature_view_service_list_feature_views(
            project=self.project
        )

        if not res.feature_views:
            return

        for fv in res.feature_views:
            fv_proto = feature_view_to_proto(fv)
            registry_proto.feature_views.append(fv_proto)

    def _get_infra(self, registry_proto: RegistryProto):
        res = self._infra_object_api.infra_object_service_list_infra_objects(
            project=self.project
        )

        if not res.infra_objects:
            return

        for infra in res.infra_objects:
            infra_proto = infra_to_proto(infra)
            registry_proto.infra.infra_objects.append(infra_proto)

    def _get_on_demand_fvs(self, registry_proto: RegistryProto):
        res = self._on_demand_fv_api.on_demand_feature_view_service_list_on_demand_feature_views(
            project=self.project
        )

        if not res.on_demand_feature_views:
            return

        for odfv in res.on_demand_feature_views:
            odfv_proto = odfv_to_proto(odfv)
            registry_proto.on_demand_feature_views.append(odfv_proto)

    def _get_request_fvs(self, registry_proto: RegistryProto):
        res = self._request_fv_api.request_feature_view_service_list_request_feature_views(
            project=self.project
        )

        if not res.request_feature_views:
            return

        for rfv in res.request_feature_views:
            rfv_proto = rfv_to_proto(rfv)
            registry_proto.request_feature_views.append(rfv_proto)

    def _get_saved_datasets(self, registry_proto: RegistryProto):
        res = self._saved_dataset_api.saved_dataset_service_list_saved_datasets(
            project=self.project
        )

        if not res.saved_datasets:
            return

        for sd in res.saved_datasets:
            sd_proto = saved_dataset_to_proto(sd)
            registry_proto.saved_datasets.append(sd_proto)

    # Update methods
    def _update_entities(self, op):
        if op["op"] == "CreateEntity":
            entity = entity_to_api(op["proto"])
            self._entity_api.entity_service_create_entity(body=entity)
        elif op["op"] == "UpdateEntity":
            entity = entity_to_api(op["proto"])
            self._entity_api.entity_service_update_entity(body=entity)
        elif op["op"] == "DeleteEntity":
            self._entity_api.entity_service_delete_entity(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")

    def _update_data_sources(self, op):
        if op["op"] == "CreateDataSource":
            data_source = data_source_to_api(op["proto"])
            self._data_source_api.data_source_service_create_data_source(
                body=data_source
            )
        elif op["op"] == "UpdateDataSource":
            data_source = data_source_to_api(op["proto"])
            self._data_source_api.data_source_service_update_data_source(
                body=data_source
            )
        elif op["op"] == "DeleteDataSource":
            self._data_source_api.data_source_service_delete_data_source(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")

    def _update_feature_services(self, op):
        if op["op"] == "CreateFeatureService":
            feature_service = feature_service_to_api(op["proto"])
            self._feature_service_api.feature_service_service_create_feature_service(
                body=feature_service
            )
        elif op["op"] == "UpdateFeatureService":
            feature_service = feature_service_to_api(op["proto"])
            self._feature_service_api.feature_service_service_update_feature_service(
                body=feature_service
            )
        elif op["op"] == "DeleteFeatureService":
            self._feature_service_api.feature_service_service_delete_feature_service(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")

    def _update_feature_views(self, op):
        if op["op"] == "CreateFeatureView":
            feature_view = feature_view_to_api(op["proto"])
            self._feature_view_api.feature_view_service_create_feature_view(
                body=feature_view
            )
        elif op["op"] == "UpdateFeatureView":
            feature_view = feature_view_to_api(op["proto"])
            self._feature_view_api.feature_view_service_update_feature_view(
                body=feature_view
            )
        elif op["op"] == "DeleteFeatureView":
            self._feature_view_api.feature_view_service_delete_feature_view(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")

    def _update_on_demand_fvs(self, op):
        if op["op"] == "CreateOnDemandFeatureView":
            on_demand_fv = odfv_to_api(op["proto"])
            self._on_demand_fv_api.on_demand_feature_view_service_create_on_demand_feature_view(
                body=on_demand_fv
            )
        elif op["op"] == "UpdateOnDemandFeatureView":
            on_demand_fv = odfv_to_api(op["proto"])
            self._on_demand_fv_api.on_demand_feature_view_service_update_on_demand_feature_view(
                body=on_demand_fv
            )
        elif op["op"] == "DeleteOnDemandFeatureView":
            self._on_demand_fv_api.on_demand_feature_view_service_delete_on_demand_feature_view(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")

    def _update_request_fvs(self, op):
        if op["op"] == "CreateRequestFeatureView":
            request_fv = rfv_to_api(op["proto"])
            self._request_fv_api.request_feature_view_service_create_request_feature_view(
                body=request_fv
            )
        elif op["op"] == "UpdateRequestFeatureView":
            request_fv = rfv_to_api(op["proto"])
            self._request_fv_api.request_feature_view_service_update_request_feature_view(
                body=request_fv
            )
        elif op["op"] == "DeleteRequestFeatureView":
            self._request_fv_api.request_feature_view_service_delete_request_feature_view(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")

    def _update_saved_datasets(self, op):
        if op["op"] == "CreateSavedDataset":
            saved_dataset = saved_dataset_to_api(op["proto"])
            self._saved_dataset_api.saved_dataset_service_create_saved_dataset(
                body=saved_dataset
            )
        elif op["op"] == "UpdateSavedDataset":
            saved_dataset = saved_dataset_to_api(op["proto"])
            self._saved_dataset_api.saved_dataset_service_update_saved_dataset(
                body=saved_dataset
            )
        elif op["op"] == "DeleteSavedDataset":
            self._saved_dataset_api.saved_dataset_service_delete_saved_dataset(
                name=op["name"], project=self.project
            )
        else:
            raise ValueError(f"Not supported op: {op['op']}")


NAMESPACE_PATH = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"


def get_pod_namespace() -> str:
    namespace = None
    try:
        with open(NAMESPACE_PATH, "r") as f:
            namespace = f.read().strip()
    except OSError as e:
        logging.error(
            "Failed to read a token from file '%s' (%s).", NAMESPACE_PATH, str(e)
        )
        raise
    return namespace
