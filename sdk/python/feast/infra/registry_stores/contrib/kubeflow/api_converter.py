from datetime import datetime, timezone
from typing import Tuple

from frs_api.models import (
    ApiDataSource,
    ApiEntity,
    ApiFeature,
    ApiFeatureService,
    ApiFeatureView,
    ApiFeatureViewProjection,
    ApiInfraObject,
    ApiMaterializationInterval,
    ApiOnDemandFeatureView,
    ApiProject,
    ApiRequestFeatureView,
    ApiSavedDataset,
    ApiUserDefinedFunction,
)
from google.protobuf.json_format import MessageToDict, MessageToJson

from feast.protos.feast.core.DataSource_pb2 import DataSource as DataSourceProto
from feast.protos.feast.core.Entity_pb2 import Entity as EntityProto
from feast.protos.feast.core.Feature_pb2 import FeatureSpecV2 as FeatureProto
from feast.protos.feast.core.FeatureService_pb2 import (
    FeatureService as FeatureServiceProto,
)
from feast.protos.feast.core.FeatureView_pb2 import FeatureView as FeatureViewProto
from feast.protos.feast.core.FeatureViewProjection_pb2 import (
    FeatureViewProjection as FeatureViewProjectionProto,
)
from feast.protos.feast.core.InfraObject_pb2 import Infra as InfraProto
from feast.protos.feast.core.OnDemandFeatureView_pb2 import (
    OnDemandFeatureView as OnDemandFeatureViewProto,
)
from feast.protos.feast.core.Registry_pb2 import Registry as RegistryProto
from feast.protos.feast.core.RequestFeatureView_pb2 import (
    RequestFeatureView as RequestFeatureViewProto,
)
from feast.protos.feast.core.SavedDataset_pb2 import SavedDataset as SavedDatasetProto


def entity_to_api(proto: EntityProto) -> ApiEntity:
    return ApiEntity(
        name=proto.spec.name,
        project=proto.spec.project,
        value_type=proto.spec.value_type,
        description=proto.spec.description,
        join_key=proto.spec.join_key,
        tags=dict(proto.spec.tags),
        owner=proto.spec.owner,
        created_timestamp=proto.meta.created_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        last_updated_timestamp=proto.meta.last_updated_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
    )


def data_source_to_api(proto: DataSourceProto) -> ApiDataSource:
    if proto.HasField("file_options"):
        options = MessageToJson(proto.file_options)
    elif proto.HasField("bigquery_options"):
        options = MessageToJson(proto.bigquery_options)
    elif proto.HasField("kafka_options"):
        options = MessageToJson(proto.kafka_options)
    elif proto.HasField("kinesis_options"):
        options = MessageToJson(proto.kinesis_options)
    elif proto.HasField("redshift_options"):
        options = MessageToJson(proto.redshift_options)
    elif proto.HasField("request_data_options"):
        options = MessageToJson(proto.request_data_options)
    elif proto.HasField("custom_options"):
        options = MessageToJson(proto.custom_options)
    elif proto.HasField("snowflake_options"):
        options = MessageToJson(proto.snowflake_options)
    elif proto.HasField("push_options"):
        options = MessageToJson(proto.push_options)
    elif proto.HasField("spark_options"):
        options = MessageToJson(proto.spark_options)
    elif proto.HasField("trino_options"):
        options = MessageToJson(proto.trino_options)
    else:
        raise ValueError("This type of options is not supported")

    return ApiDataSource(
        name=proto.name,
        project=proto.project,
        description=proto.description,
        tags=dict(proto.tags),
        owner=proto.owner,
        type=proto.type,
        field_mapping=dict(proto.field_mapping),
        timestamp_field=proto.timestamp_field,
        date_partition_column=proto.date_partition_column,
        created_timestamp_column=proto.created_timestamp_column,
        data_source_class_type=proto.data_source_class_type,
        batch_source=MessageToJson(proto.batch_source),
        options=options,
    )


def feature_service_to_api(proto: FeatureServiceProto) -> ApiFeatureService:
    features = []
    for feature in proto.spec.features:
        features.append(fvp_to_api(feature))

    return ApiFeatureService(
        name=proto.spec.name,
        project=proto.spec.project,
        features=features,
        tags=dict(proto.spec.tags),
        description=proto.spec.description,
        owner=proto.spec.owner,
        logging_config=MessageToJson(proto.spec.logging_config),
        created_timestamp=proto.meta.created_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        last_updated_timestamp=proto.meta.last_updated_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
    )


def fvp_to_api(proto: FeatureViewProjectionProto) -> ApiFeatureViewProjection:
    feature_columns = []
    for feature in proto.feature_columns:
        feature_columns.append(feature_to_api(feature))

    return ApiFeatureViewProjection(
        feature_view_name=proto.feature_view_name,
        feature_view_name_alias=proto.feature_view_name_alias,
        feature_columns=feature_columns,
        join_key_map=dict(proto.join_key_map),
    )


def feature_view_to_api(proto: FeatureViewProto) -> ApiFeatureView:
    features = []
    for feature in proto.spec.features:
        features.append(feature_to_api(feature))

    mis = []
    for mi in proto.meta.materialization_intervals:
        mis.append(
            mi_to_api(
                (
                    mi.start_time.ToDatetime().replace(tzinfo=timezone.utc),
                    mi.end_time.ToDatetime().replace(tzinfo=timezone.utc),
                )
            )
        )

    return ApiFeatureView(
        name=proto.spec.name,
        project=proto.spec.project,
        entities=list(proto.spec.entities),
        features=features,
        description=proto.spec.description,
        tags=dict(proto.spec.tags),
        owner=proto.spec.owner,
        ttl=str(proto.spec.ttl.ToSeconds()) + "s",
        batch_source=MessageToJson(proto.spec.batch_source),
        stream_source=MessageToJson(proto.spec.stream_source),
        online=proto.spec.online,
        created_timestamp=proto.meta.created_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        last_updated_timestamp=proto.meta.last_updated_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        materialization_intervals=mis,
    )


def feature_to_api(proto: FeatureProto) -> ApiFeature:
    return ApiFeature(
        name=proto.name,
        value_type=proto.value_type,
        tags=dict(proto.tags) if hasattr(proto, "tags") else dict(),
    )


def mi_to_api(mi: Tuple[datetime, datetime]) -> ApiMaterializationInterval:
    return ApiMaterializationInterval(start_time=mi[0], end_time=mi[1])


def infra_to_api(proto: InfraProto) -> ApiInfraObject:
    infra_objects = []

    for infra_object_proto in proto.infra_objects:
        infra_type = infra_object_proto.WhichOneof("infra_object")
        if infra_type == "dynamodb_table":
            infra_object = MessageToJson(infra_object_proto.dynamodb_table)
        elif infra_type == "datastore_table":
            infra_object = MessageToJson(infra_object_proto.datastore_table)
        elif infra_type == "sqlite_table":
            infra_object = MessageToJson(infra_object_proto.sqlite_table)
        elif infra_type == "custom_infra":
            infra_object = MessageToJson(infra_object_proto.custom_infra)
        else:
            raise ValueError("This type of infra object is not supported")

        infra_objects.append(
            ApiInfraObject(
                infra_object_class_type=infra_object_proto.infra_object_class_type,
                infra_object=infra_object,
            )
        )

    return infra_objects


def odfv_to_api(proto: OnDemandFeatureViewProto) -> ApiOnDemandFeatureView:
    features = []
    for feature in proto.spec.features:
        features.append(feature_to_api(feature))

    sources = dict()
    for name, source in proto.spec.sources.items():
        sources[name] = MessageToJson(source)

    return ApiOnDemandFeatureView(
        name=proto.spec.name,
        project=proto.spec.project,
        features=features,
        sources=sources,
        user_defined_function=ApiUserDefinedFunction(
            name=proto.spec.user_defined_function.name,
            body=MessageToDict(proto.spec.user_defined_function)["body"],
        ),
        description=proto.spec.description,
        tags=dict(proto.spec.tags),
        owner=proto.spec.owner,
        created_timestamp=proto.meta.created_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        last_updated_timestamp=proto.meta.last_updated_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
    )


def project_to_api(proto: RegistryProto, project_name: str) -> ApiProject:
    return ApiProject(
        name=project_name,
        registry_schema_version=proto.registry_schema_version,
        version_id=proto.version_id,
        last_updated=proto.last_updated.ToDatetime().replace(tzinfo=timezone.utc),
    )


def rfv_to_api(proto: RequestFeatureViewProto) -> ApiRequestFeatureView:
    return ApiRequestFeatureView(
        name=proto.spec.name,
        project=proto.spec.project,
        request_data_source=MessageToJson(proto.spec.request_data_source),
        description=proto.spec.description,
        tags=dict(proto.spec.tags),
        owner=proto.spec.owner,
    )


def saved_dataset_to_api(proto: SavedDatasetProto) -> ApiSavedDataset:
    return ApiSavedDataset(
        name=proto.spec.name,
        project=proto.spec.project,
        features=list(proto.spec.features),
        join_keys=list(proto.spec.join_keys),
        full_feature_names=proto.spec.full_feature_names,
        storage=MessageToJson(proto.spec.storage),
        feature_service_name=proto.spec.feature_service_name,
        tags=dict(proto.spec.tags),
        created_timestamp=proto.meta.created_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        last_updated_timestamp=proto.meta.last_updated_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        min_event_timestamp=proto.meta.min_event_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
        max_event_timestamp=proto.meta.max_event_timestamp.ToDatetime().replace(
            tzinfo=timezone.utc
        ),
    )
