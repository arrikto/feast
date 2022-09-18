import json
import pandas as pd
import sqlalchemy

from kafka import KafkaConsumer

from feast import FeatureStore
from feast.repo_config import RepoConfig, RegistryConfig
from feast.infra.offline_stores.contrib.postgres_offline_store.postgres import PostgreSQLOfflineStoreConfig
from feast.infra.online_stores.redis import RedisOnlineStoreConfig

print("Configuring feature store")
offline_store_config = PostgreSQLOfflineStoreConfig(
    host="postgresql-offline-store.default.svc.cluster.local",
    database="driver_stream_data",
    db_schema="driver_stream_data",
    user="charlie",
    password="charlie"
)

online_store_config = RedisOnlineStoreConfig(
    connection_string="redis-online-store.default.svc.cluster.local:6379,username=charlie,password=charlie,db=0"
)

registry_config = RegistryConfig(
    registry_store_type="KubeflowRegistryStore",
    path="",
    project="kubeflow-charlie",
)

repo_config = RepoConfig(
    project="kubeflow-charlie",
    registry=registry_config,
    provider="local",
    offline_store=offline_store_config,
    online_store=online_store_config
)

fs = FeatureStore(config=repo_config, repo_path=None)

def get_sqlalchemy_engine(config: PostgreSQLOfflineStoreConfig):
    url = f"postgresql+psycopg2://{config.user}:{config.password}@{config.host}:{config.port}/{config.database}"
    print("Connecting to", config.db_schema, "schema using:", url)
    return sqlalchemy.create_engine(url, client_encoding='utf8', connect_args={'options': '-c search_path={}'.format(config.db_schema)})

con = get_sqlalchemy_engine(offline_store_config)

servers = "kafka-broker.default.svc.cluster.local:29092"

consumer = KafkaConsumer('driver-locations',
                        bootstrap_servers=[servers]
                        )

print("Start consuming")
for message in consumer:
    msg = json.loads(message.value)
    print(msg)
    
    df = pd.DataFrame.from_dict([msg])
    
    print("Pushing to offline store")
    df.to_sql(
        name="driver_locations",
        con=con,
        if_exists="append",
        dtype={
            "event_timestamp": sqlalchemy.TIMESTAMP,
            "driver_id": sqlalchemy.INT,
            "lat": sqlalchemy.FLOAT,
            "lon": sqlalchemy.FLOAT
        }
    )
    
    print("Pushing to online store")
    fs.write_to_online_store(
        feature_view_name="driver_locations_fv",
        df=df
    )
