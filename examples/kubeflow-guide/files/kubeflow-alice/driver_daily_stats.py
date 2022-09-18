from datetime import timedelta
from feast import Entity, FeatureView, Field
from feast.infra.offline_stores.contrib.postgres_offline_store.postgres_source import (
    PostgreSQLSource,
)
from feast.types import Float32, Int64

driver_daily_stats_source = PostgreSQLSource(
    name="driver_daily_stats_source",
    query="SELECT * FROM driver_daily_stats",
    timestamp_field="event_timestamp"
)

driver = Entity(
    name="driver",
    value_type=Int64,
    description="",
    join_keys=["driver_id"],
    tags={},
    owner="alice"
)

driver_daily_stats_fv = FeatureView(
    name="driver_daily_stats_fv",
    entities=["driver"],
    description="",
    tags={},
    owner="alice",
    ttl=timedelta(days=7),
    source=driver_daily_stats_source,
    online=True,
    schema=[
        Field(name="acc_rate", dtype=Float32, tags={}),
        Field(name="profit",dtype=Float32, tags={})
    ]
)