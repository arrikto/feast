# Copyright 2022 Arrikto Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import base64
import enum
from datetime import datetime, timedelta, timezone

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
    ApiRequestFeatureView,
    ApiSavedDataset,
)
from google.protobuf.duration_pb2 import Duration
from google.protobuf.json_format import Parse

from feast import ValueType
from feast.protos.feast.core.DataSource_pb2 import DataSource as DataSourceProto
from feast.protos.feast.core.Entity_pb2 import Entity as EntityProto
from feast.protos.feast.core.Entity_pb2 import EntityMeta as EntityMetaProto
from feast.protos.feast.core.Entity_pb2 import EntitySpecV2 as EntitySpecProto
from feast.protos.feast.core.Feature_pb2 import FeatureSpecV2 as FeatureProto
from feast.protos.feast.core.FeatureService_pb2 import (
    FeatureService as FeatureServiceProto,
)
from feast.protos.feast.core.FeatureService_pb2 import (
    FeatureServiceMeta as FeatureServiceMetaProto,
)
from feast.protos.feast.core.FeatureService_pb2 import (
    FeatureServiceSpec as FeatureServiceSpecProto,
)
from feast.protos.feast.core.FeatureView_pb2 import FeatureView as FeatureViewProto
from feast.protos.feast.core.FeatureView_pb2 import (
    FeatureViewMeta as FeatureViewMetaProto,
)
from feast.protos.feast.core.FeatureView_pb2 import (
    FeatureViewSpec as FeatureViewSpecProto,
)
from feast.protos.feast.core.FeatureView_pb2 import MaterializationInterval as MIProto
from feast.protos.feast.core.FeatureViewProjection_pb2 import (
    FeatureViewProjection as FVPProto,
)
from feast.protos.feast.core.InfraObject_pb2 import InfraObject as InfraObjectProto
from feast.protos.feast.core.OnDemandFeatureView_pb2 import (
    OnDemandFeatureView as OnDemandFVProto,
)
from feast.protos.feast.core.OnDemandFeatureView_pb2 import (
    OnDemandFeatureViewMeta as OnDemandFVMetaProto,
)
from feast.protos.feast.core.OnDemandFeatureView_pb2 import (
    OnDemandFeatureViewSpec as OnDemandFVSpecProto,
)
from feast.protos.feast.core.OnDemandFeatureView_pb2 import (
    OnDemandSource as OnDemandSourceProto,
)
from feast.protos.feast.core.OnDemandFeatureView_pb2 import (
    UserDefinedFunction as UDFProto,
)
from feast.protos.feast.core.RequestFeatureView_pb2 import (
    RequestFeatureView as RequestFVProto,
)
from feast.protos.feast.core.RequestFeatureView_pb2 import (
    RequestFeatureViewSpec as RequestFVSpecProto,
)
from feast.protos.feast.core.SavedDataset_pb2 import SavedDataset as SavedDatasetProto
from feast.protos.feast.core.SavedDataset_pb2 import (
    SavedDatasetMeta as SavedDatasetMetaProto,
)
from feast.protos.feast.core.SavedDataset_pb2 import (
    SavedDatasetSpec as SavedDatasetSpecProto,
)


class SourceType(enum.Enum):
    INVALID = DataSourceProto.INVALID
    BATCH_FILE = DataSourceProto.BATCH_FILE
    BATCH_SNOWFLAKE = DataSourceProto.BATCH_SNOWFLAKE
    BATCH_BIGQUERY = DataSourceProto.BATCH_BIGQUERY
    BATCH_REDSHIFT = DataSourceProto.BATCH_REDSHIFT
    STREAM_KAFKA = DataSourceProto.STREAM_KAFKA
    STREAM_KINESIS = DataSourceProto.STREAM_KINESIS
    CUSTOM_SOURCE = DataSourceProto.CUSTOM_SOURCE
    REQUEST_SOURCE = DataSourceProto.REQUEST_SOURCE
    PUSH_SOURCE = DataSourceProto.PUSH_SOURCE
    BATCH_TRINO = DataSourceProto.BATCH_TRINO
    BATCH_SPARK = DataSourceProto.BATCH_SPARK


def data_source_to_proto(api: ApiDataSource) -> DataSourceProto:
    data_source_proto = DataSourceProto(
        name=api.name,
        project=api.project,
        description=api.description,
        tags=api.tags,
        owner=api.owner,
        type=SourceType[api.type].value,
        field_mapping=api.field_mapping,
        timestamp_field=api.timestamp_field,
        date_partition_column=api.date_partition_column,
        created_timestamp_column=api.created_timestamp_column,
        data_source_class_type=api.data_source_class_type,
    )

    Parse(api.batch_source, data_source_proto.batch_source)

    source_type = getattr(api, "type", "INVALID")
    if source_type == "BATCH_FILE":
        Parse(api.options, data_source_proto.file_options)
    elif source_type == "BATCH_SNOWFLAKE":
        Parse(api.options, data_source_proto.snowflake_options)
    elif source_type == "BATCH_BIGQUERY":
        Parse(api.options, data_source_proto.bigquery_options)
    elif source_type == "BATCH_REDSHIFT":
        Parse(api.options, data_source_proto.redshift_options)
    elif source_type == "STREAM_KAFKA":
        Parse(api.options, data_source_proto.kafka_options)
    elif source_type == "STREAM_KINESIS":
        Parse(api.options, data_source_proto.kinesis_options)
    elif source_type == "CUSTOM_SOURCE":
        Parse(api.options, data_source_proto.custom_options)
    elif source_type == "REQUEST_SOURCE":
        Parse(api.options, data_source_proto.request_data_options)
    elif source_type == "PUSH_SOURCE":
        Parse(api.options, data_source_proto.push_options)
    elif source_type == "BATCH_TRINO":
        Parse(api.options, data_source_proto.trino_options)
    elif source_type == "BATCH_SPARK":
        Parse(api.options, data_source_proto.spark_options)
    else:
        raise ValueError("Invalid data source")

    return data_source_proto


def entity_to_proto(api: ApiEntity) -> EntityProto:
    meta = EntityMetaProto()
    meta.created_timestamp.FromDatetime(
        getattr(api, "created_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )
    meta.last_updated_timestamp.FromDatetime(
        getattr(
            api, "last_updated_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc)
        )
    )

    spec = EntitySpecProto(
        name=api.name,
        project=api.project,
        value_type=ValueType[api.value_type].value,
        description=api.description,
        join_key=api.join_key,
        tags=api.tags,
        owner=api.owner,
    )

    return EntityProto(spec=spec, meta=meta)


def feature_service_to_proto(api: ApiFeatureService) -> FeatureServiceProto:
    meta = FeatureServiceMetaProto()
    meta.created_timestamp.FromDatetime(
        getattr(api, "created_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )
    meta.last_updated_timestamp.FromDatetime(
        getattr(
            api, "last_updated_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc)
        )
    )

    features = []
    for feature in api.features:
        features.append(fvp_to_proto(feature))

    spec = FeatureServiceSpecProto(
        name=api.name,
        project=api.project,
        features=features,
        tags=api.tags,
        description=api.description,
        owner=api.owner,
    )

    Parse(api.logging_config, spec.logging_config)

    return FeatureServiceProto(spec=spec, meta=meta)


def fvp_to_proto(api: ApiFeatureViewProjection) -> FVPProto:
    feature_columns = []
    for feature in api.feature_columns:
        feature_columns.append(feature_to_proto(feature))

    return FVPProto(
        feature_view_name=api.feature_view_name,
        feature_view_name_alias=api.feature_view_name_alias,
        feature_columns=feature_columns,
        join_key_map=api.join_key_map,
    )


def feature_view_to_proto(api: ApiFeatureView) -> FeatureViewProto:
    meta = FeatureViewMetaProto()
    meta.created_timestamp.FromDatetime(
        getattr(api, "created_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )
    meta.last_updated_timestamp.FromDatetime(
        getattr(
            api, "last_updated_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc)
        )
    )

    if api.materialization_intervals:
        for mi in api.materialization_intervals:
            meta.materialization_intervals.append(mi_to_proto(mi))

    features = []
    for feature in api.features:
        features.append(feature_to_proto(feature))

    ttl = Duration()
    ttl.FromTimedelta(timedelta(seconds=float(api.ttl[:-1])))

    spec = FeatureViewSpecProto(
        name=api.name,
        project=api.project,
        entities=api.entities,
        features=features,
        description=api.description,
        tags=api.tags,
        owner=api.owner,
        ttl=ttl,
        online=api.online,
    )

    Parse(api.batch_source, spec.batch_source)
    Parse(api.stream_source, spec.stream_source)

    return FeatureViewProto(spec=spec, meta=meta)


def feature_to_proto(api: ApiFeature) -> FeatureProto:
    return FeatureProto(
        name=api.name, value_type=ValueType[api.value_type].value, tags=api.tags
    )


def mi_to_proto(api: ApiMaterializationInterval) -> MIProto:
    mi_proto = MIProto()
    mi_proto.start_time.FromDatetime(api.start_time)
    mi_proto.end_time.FromDatetime(api.end_time)

    return mi_proto


def infra_to_proto(api: ApiInfraObject) -> InfraObjectProto:
    infra_proto = InfraObjectProto(infra_object_class_type=api.infra_object_class_type)

    if (
        infra_proto.infra_object_class_type
        == "feast.infra.online_stores.datastore.DatastoreTable"
    ):
        Parse(api.infra_object, infra_proto.datastore_table)
    elif (
        infra_proto.infra_object_class_type
        == "feast.infra.online_stores.dynamodb.DynamoDBTable"
    ):
        Parse(api.infra_object, infra_proto.dynamodb_table)
    elif (
        infra_proto.infra_object_class_type
        == "feast.infra.online_stores.sqlite.SqliteTable"
    ):
        Parse(api.infra_object, infra_proto.sqlite_table)
    else:
        Parse(api.infra_object, infra_proto.custom_infra)

    return infra_proto


def odfv_to_proto(api: ApiOnDemandFeatureView) -> OnDemandFVProto:
    meta = OnDemandFVMetaProto()
    meta.created_timestamp.FromDatetime(
        getattr(api, "created_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )
    meta.last_updated_timestamp.FromDatetime(
        getattr(
            api, "last_updated_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc)
        )
    )

    features = []
    for feature in api.features:
        features.append(feature_to_proto(feature))

    sources = dict()
    for name, source in api.sources.items():
        on_demand_source_proto = OnDemandSourceProto()
        Parse(source, on_demand_source_proto)
        sources[name] = on_demand_source_proto

    base64_body_bytes = api.user_defined_function.body.encode("ascii")

    spec = OnDemandFVSpecProto(
        name=api.name,
        project=api.project,
        features=features,
        sources=sources,
        user_defined_function=UDFProto(
            name=api.user_defined_function.name,
            body=base64.b64decode(base64_body_bytes),
        ),
        description=api.description,
        tags=api.tags,
        owner=api.owner,
    )

    return OnDemandFVProto(spec=spec, meta=meta)


def rfv_to_proto(api: ApiRequestFeatureView) -> RequestFVProto:
    spec = RequestFVSpecProto(
        name=api.name,
        project=api.project,
        description=api.description,
        tags=api.tags,
        owner=api.owner,
    )

    Parse(api.request_data_source, spec.request_data_source)

    return RequestFVProto(spec=spec)


def saved_dataset_to_proto(api: ApiSavedDataset) -> SavedDatasetProto:
    meta = SavedDatasetMetaProto()
    meta.created_timestamp.FromDatetime(
        getattr(api, "created_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )
    meta.last_updated_timestamp.FromDatetime(
        getattr(
            api, "last_updated_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc)
        )
    )
    meta.min_event_timestamp.FromDatetime(
        getattr(api, "min_event_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )
    meta.max_event_timestamp.FromDatetime(
        getattr(api, "max_event_timestamp", datetime(1970, 1, 1, tzinfo=timezone.utc))
    )

    spec = SavedDatasetSpecProto(
        name=api.name,
        project=api.project,
        features=api.features,
        join_keys=api.join_keys,
        full_feature_names=api.full_feature_names,
        feature_service_name=api.feature_service_name,
        tags=api.tags,
    )

    Parse(api.storage, spec.storage)

    return SavedDatasetProto(spec=spec, meta=meta)
