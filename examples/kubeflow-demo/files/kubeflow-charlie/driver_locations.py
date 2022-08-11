import pandas as pd

from feast import Entity, Field, FeatureView, KafkaSource
from feast.data_format import AvroFormat
from feast.infra.offline_stores.contrib.postgres_offline_store.postgres_source import (
    PostgreSQLSource,
)
from feast.on_demand_feature_view import on_demand_feature_view
from feast.types import Float32, Int64

driver_locations_batch_source = PostgreSQLSource(
    name="driver_locations_batch_source",
    query="SELECT * FROM driver_locations",
    timestamp_field="event_timestamp"
)

driver_locations_stream_source = KafkaSource(
    name="driver_locations_stream",
    bootstrap_servers="kafka-broker.default.svc.cluster.local:29092",
    topic="driver-locations",
    timestamp_field="event_timestamp",
    batch_source=driver_locations_batch_source,
    message_format=AvroFormat(
        schema_json="""{
        "name": "driver_location",
        "type": "record",
        "namespace": "none",
        "fields": [
        {
          "name": "driver_id",
          "type": "int"
        },
        {
          "name": "event_timestamp",
          "type": "string"
        },
        {
          "name": "lat",
          "type": "float"
        },
        {
          "name": "lon",
          "type": "float"
        }
        ]
        }"""
    ),
    owner="charlie"
)

driver = Entity(
    name="driver",
    value_type=Int64,
    description="",
    join_keys=["driver_id"],
    tags={},
    owner="charlie"
)

driver_locations_fv = FeatureView(
    name="driver_locations_fv",
    entities=["driver"],
    schema=[
        Field(name="lat",dtype=Float32, tags={}),
        Field(name="lon", dtype=Float32, tags={})
    ],
    source=driver_locations_stream_source
)

@on_demand_feature_view(
   sources=[
       driver_locations_fv
   ],
   schema=[
     Field(name='zone', dtype=Int64)
   ]
)
def driver_zones_odfv(features_df: pd.DataFrame) -> pd.DataFrame:
    def encode(lat, lon):
        if lat == None or lon == None:
            return -1
        if lat > 37.97900 and lon < 23.78600:
            return 1
        elif lat > 37.97900 and lon >= 23.78600:
            return 2
        elif lat <= 37.97900 and lon >= 23.78600:
            return 3
        else:
            return 4

    df = pd.DataFrame()
    df['zone'] = features_df.apply(lambda x: encode(x.lat, x.lon), axis=1)
    return df
