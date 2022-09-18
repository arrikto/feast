# Feast x Kubeflow Demo

This demo focuses on how to use Feast in a Kubeflow environment. A feature
store like Feast can help you solve typical ML problems such as [training-serving
skew](https://developers.google.com/machine-learning/guides/rules-of-ml#training-serving_skew)
or [data leakage](https://machinelearningmastery.com/data-leakage-machine-learning/).
Using the Feast SDK, you can retrieve historical features to train your 
model in a [point-in-time correct way](https://docs.feast.dev/getting-started/concepts/point-in-time-joins)
and avoid the data leakage problem. In addition, you can ensure that your
model gets consistent data for training and serving and avert any drop in
its performance.

Integrating Feast in a multi-user environment like Kubeflow allows you to
share features in a standardized way. As a result, it helps you to reduce
the time and effort you need to create new features and make them available
in production systems. The best way to understand the value of this integration
is to get your hands dirty with the following demo. Once you complete it,
imagine how different the process would be without a feature store.

Outline:
- Part A: Set up your cluster and deploy all the required Feast components.
- Part B: Set up your notebook servers.
- Part C: Run the actual demo.

**WIP**: In future MiniKF releases, you will not need to execute
parts A and B to use Feast.

## Part A: Cluster Setup Instructions

This first part aims to set up the Kubernetes cluster and deploy all the
required Feast components (Registry, Offline Store, Online Store). Until
now, the Feast Registry was a single file. Feast stored all its metadata
and feature definitions in a single [protobuf](https://developers.google.com/protocol-buffers)
file and dumped it in an S3 or GCS bucket. We turned the Feast Registry
into a client-server system that:
- provides a REST API to interact with
- stores all metadata and feature definitions in an SQL database
- enforces access control using Kubernetes RBAC

### Step 0: Preparation
To start, connect to your Kubernetes cluster (MiniKF) and ensure that everything is up and running.

Then, clone the Feast repo and navigate to the `cluster-setup` directory.
```
git clone https://github.com/arrikto/prv-feast
cd prv-feast
git checkout -b feature-kubeflow origin/feature-kubeflow
cd examples/kubeflow-demo/cluster-setup
```

### Step 1.A: Deploy Registry - MySQL Database
Set up a MySQL database under `kubeflow` namespace (StatefulSet).
The API server uses it to save the Feast registry objects.

The database has two users:
- user `root` - passw `root`
- user `feastadmin` - passw `feastadmin`

User `feastadmin` has permissions only for `feast` database created with
`MYSQL_DATABASE` env variable.

We choose the MySQL native password hashing method for authentication.
```
kubectl apply -f mysql.yaml
```

Verify that MySQL is up and running.
```
kubectl get pod feast-mysql-sts-0 -n kubeflow
```
```
NAME                READY   STATUS    RESTARTS   AGE
feast-mysql-sts-0   2/2     Running   0          1m
```

### Step 1.B: Deploy Registry - Security
Set up Istio authorization policies.

Allow access to MySQL only from `feast-registry` service account.
Only the API server uses this service account.

Allow access to the API server from inside the Kubernetes cluster.
After the API server receives a request, it performs authentication (TokenReview)
and authorization (SubjectAccessReview).
```
kubectl apply -f istio-policies.yaml
```

Set up the `feast-registry` service account.
```
kubectl apply -f feast-apiserver-sa.yaml
```

The API server performs TokenReviews (authentication) and SubjectAccessReviews
(authorization). Thus, it needs permissions to perform `create` on `tokenreviews`
and `subjectaccessreviews` resources. Τhese are cluster scoped, and we need
to create a ClusterRole and a corresponding ClusterRoleBinding.
```
k apply -f feast-apiserver-role.yaml
k apply -f feast-apiserver-rolebinding.yaml
```

### Step 1.C: Deploy Registry - API Server
Set up the API server under `kubeflow` namespace (Deployment).
The API server handles the Feast client requests and interacts with MySQL.

The ConfigMap contains the database credentials used by the API server to
interact with MySQL.
```
kubectl apply -f feast-registry-config.yaml
```

The API server uses the previously created `feast-registry` service account.
```
kubectl apply -f feast-registry.yaml
```

Verify that the API server is up and running.
```
kubectl get pods -n kubeflow | grep feast-registry-d*
```
```
feast-registry-d-74675b8654-cwmp2    2/2    Running    0    1m
```

### Step 1.D: Deploy Registry - Authentication & Authorization support
By default, a namespace owner can perform every action on every resource
belonging to the namespace. The same logic applies to Feast registry resources.
We create two new ClusterRoles, `kubeflow-feature-edit` and `kubeflow-feature-view`,
that we aggregate to `kubeflow-edit` and `kubeflow-view` roles accordingly.

Users bound to `kubeflow-edit` can `create`, `update` and `delete` every
Feast registry resource, whereas users tied to `kubeflow-view` can `get`
and `list` every Feast registry resource.
```
kubectl apply -f features-cluster-roles.yaml
```

Users must provide a service account token in every request they make to
the API server. This token has a specific target audience `features.kubeflow.org`
and is mounted to every new Pod using a Service Account Token Volume Projection.

After creating new users, apply the `access-features` PodDefault for every
namespace you need to make requests to the API server.

Example:
```
kubectl apply -f features-poddefault.yaml -n kubeflow-user
```

Every time a Pod starts, it mounts the token targeting the Feast API server.

### Step 2: Deploy Offline and Online Store
In general, feature stores are implemented as a dual database system. An
**Offline Store** and an **Online Store**. The Offline Store is usually
a scale-out database and contains historical time-series feature data (data
\+ timestamp), whereas the Online Store is a fast key-value store and contains
only the latest feature data.

Feast supports BigQuery, Redshift, Snowflake, and PostgreSQL as an Offline
Store. We use PostgreSQL for our demo since it is vendor-independent.

Feast supports Datastore, DynamoDB, Redis, and PostgreSQL as an Online Store.
We use Redis for our Demo since it is the most [performant](https://feast.dev/blog/feast-benchmarks/).

We set up both the Offline and the Online Store under the `default` namespace.
All users have direct access (the Feast SDK client interacts directly with
the underlying data source) and we perform access control by specifying
username and password.

### Step 2.A: Deploy Offline Store - PostgreSQL Database

Admin Credentials: user `postgres` - passw `postgres`

The admin who is responsible for the database creates three users:
- user `alice` - passw `alice`
- user `bob` - passw `bob`
- user `charlie` - passw `charlie`

User `alice` has all permissions only for `driver_data` database.
User `charlie` has all permissions only for `driver_stream_data` database.

User `bob` can only get data from `driver_data` and `driver_stream_data`
databases. He has permission to connect to the databases and use the `driver_data`
and `driver_stream_data` schemas accordingly.

Note: Schema is a collection of logical structures of data. In PostgreSQL,
Schema is a named collection of tables, views, functions, constraints, indexes,
sequences, etc. [Source](https://hasura.io/learn/database/postgresql/core-concepts/1-postgresql-schema/)

```
kubectl apply -f postgresql-config.yaml
kubectl apply -f postgresql.yaml
```

Verify that PostgreSQL is up and running.
```
kubectl get pods feast-postgresql-sts-0
```
```
NAME                     READY   STATUS    RESTARTS   AGE
feast-postgresql-sts-0   1/1     Running   0          1m
```

We fill the database with data. We will use them to create new features
and train a model. 
1. Create a new Notebook Server and run `pip install sqlalchemy psycopg2-binary`
2. Run the `populate_offline_store.py` script

You can use the default `user` to run the script inside a Notebook Server.

This script creates two tables:
- `driver_daily_stats` in `driver_data` db under `driver_data` schema owned
by `alice`
- `driver_locations` in `driver_stream_data` db under `driver_stream_data`
schema owned by `charlie`.

#### Table: driver_daily_stats
| event_timestamp  | driver_id | profit  | acc_rate |
|------------------|-----------|---------|----------|
| 2021-03-17 19:31 | 1001      | 88.9188 | 0.049057 |
| 2021-03-18 19:31 | 1001      | 97.9273 | 0.212630 |
|                  |    ...    |   ...   |    ...   |
| 2021-03-19 19:31 | 1020      | 73.8860 | 0.857422 |

#### Table: driver_locations
| event_timestamp  | driver_id | lat       | lon       |
|------------------|-----------|-----------|-----------|
| 2021-03-17 19:31 | 1001      | 37.979161 | 23.783086 |
| 2021-03-18 19:41 | 1001      | 37.979163 | 23.783089 |
|                  |  ...      |    ...    |    ...    |
| 2021-03-23 19:51 | 1001      | 35.491932 | 23.783099 |

### Step 2.B: Deploy Online Store - Redis

Admin Credentials: user `default` - passw `redis`

The admin who is responsible for the database creates three users:
- user `alice` - passw `alice`
- user `bob` - passw `bob`
- user `charlie` - passw `charlie`

```
kubectl apply -f redis-config.yaml
kubectl apply -f redis.yaml
```

Verify that Redis is up and running.
```
kubectl get pods feast-redis-sts-0
```
```
NAME                READY   STATUS    RESTARTS   AGE
feast-redis-sts-0   1/1     Running   0          1m
```

Feast stores the latest feature data in a specific format inside Redis
([source](https://github.com/feast-dev/feast/blob/master/docs/specs/online_store_format.md#redis-online-store-format)).
The users have permission to read and write keys belonging only to their
namespace. We achieve this by specifying appropriate key patterns based on
a user's namespace. We use these patterns on [ACL](https://redis.io/docs/manual/security/acl/) 
(Access Control Lists).

Note: A Feast project has a 1:1 relationship to a Kubernetes namespace.

### Step 3: Deploy Kafka Broker
The following is a minimum setup of a Kafka broker for the scope of the
current demo.

```
kubectl apply -f zookeeper.yaml
kubectl apply -f kafka.yaml
```

Verify that Kafka is up and running.
```
kubectl get pods -n default | grep zookeeper-d*
kubectl get pods -n default | grep kafka-broker-d*
```
```
zookeeper-d-6bb7bf4c68-lkd7x       1/1    Running    3    1m
kafka-broker-d-65548766b9-fr8hc    1/1    Running    3    1m
```

Then, create a new topic called `driver-locations`.
```
k exec kafka-broker-d-65548766b9-fr8hc -- kafka-topics --bootstrap-server localhost:9092 --create --topic driver-locations
```

Note: Adjust the Pod name (kafka-broker-d-*) based on the above Verify section

---
## Part B: Notebook Setup Instructions

When you create a new Notebook Server, you must select the `Allow access to Kubeflow Features`
option in the Configuration. This option instructs Kubernetes to mount a
service account token targeting `features.kubeflow.org` audience. The Feast
client uses this token in every request to the API server.

Then, hit `Launch`, wait for the Notebook to start and connect to it.
Finally, install Feast and other required packages.

```
git clone https://github.com/arrikto/prv-feast.git
cd prv-feast
git checkout -b feature-kubeflow origin/feature-kubeflow
pip install --upgrade pip
pip install backend/api/python_http_client/ ".[postgres]" sqlalchemy
```

**WIP**: In future MiniKF releases, we will install Feast on every
new Notebook Server by default.

---
## Part C: Demo

![](files/imgs/demo-diagram.png)
### Story
This demo shows how data engineers can easily share features with data scientists.
Then, data scientists can use these features to train a model and make real-time
predictions.

Alice and Charlie are data engineers in a big transportation company. Alice
is responsible for batch sources and collects daily data for drivers, whereas
Charlie is responsible for stream sources and gathers the current location
of every single driver.

Bob is an ambitious data scientist who wants to create a model that shows
how appropriate a driver is to complete a ride. The model gets the following
three features as input:
- the acceptance rate of the previous day
- yesterday's profit
- current zone of a driver

It returns a score as output. The higher the score, the better a driver
to execute a ride.

Note: The acceptance rate is the number of ride requests a driver has accepted
divided by the total number of ride requests they have received.

### Step 0.A: Create Users
Create three new Dex users:
- user `alice` - passw `alice`
- user `bob` - passw `bob`
- user `charlie` - passw `charlie`

1\. To start, connect to your Kubernetes cluster and ensure that everything
is up and running.

2\. Navigate under the deployments folder.

MiniKF:
```
cd /data/deployments
```

3\. Edit `kubeflow/manifests/common/dex/overlays/deploy/patches/config-map.yaml`
and add the following users to the `staticPasswords` field.
```
- email: alice
  hash: $2y$12$Id0M1S83xP7IZmhFq/8WJOl0rXUeacn4xtz.Oxg4AwLU0Y5UznZ6i
  username: alice
  userID: 6a3f8618-26ab-467f-b561-5df08f7e8300
- email: bob
  hash: $2y$12$JEphTIS2GCXevcwDxo5kJeIsRz9W.VHqpnXO6jrcEjabrZQAfS77e
  username: bob
  userID: 3c190f8a-9b63-4d66-90c9-3da3e490b964
- email: charlie
  hash: $2y$12$2r/eTPdeQAuLGybG9Hy0M.eNnwyW.vbWKxjjnQjQz5TANC5o2inPG
  username: charlie
  userID: b7f7984e-5f6e-4ba7-a13b-8cbfbedf4446
```

4\. Commit the changes.
```
git commit -am "Add static users {alice, bob, charlie} in Dex"
```

5\. Apply the changes.
```
rok-deploy --apply kubeflow/manifests/common/dex/overlays/deploy
```

6\. Restart Dex and Authservice.
```
k delete pods -n auth -l app=dex
k delete pod authservice-0 -n istio-system
```

7\. Verify that both Dex and Authservice are up and running.
```
k get pods -n auth
k get pod authservice-0 -n istio-system
```

```
NAME                   READY   STATUS    RESTARTS   AGE
dex-64b8987bdc-pt56v   2/2     Running   1          1m

NAME            READY   STATUS    RESTARTS   AGE
authservice-0   1/1     Running   0          1m
```

8\. Log in to Kubeflow using the new users to check that everything works.

9\. Navigate under the `cluster-setup` folder of this repository and run
the following commands.

```
kubectl apply -f features-poddefault.yaml -n kubeflow-alice
kubectl apply -f features-poddefault.yaml -n kubeflow-bob
kubectl apply -f features-poddefault.yaml -n kubeflow-charlie
```
### Step 0.B: Create Notebook Servers
For each user, connect to Kubeflow and create a new Notebook following the
instructions of [Part B](#part-b-notebook-setup-instructions).

### Step 1: Alice
Connect to Alice's Notebook Server before continuing.

Alice inspects the database and finds the `driver_daily_stats` table. She
creates three definitions specifying the data source, the entity and, the
feature view (the available features).

Navigate to Alice's folder and inspect the `driver_daily_stats.py` file.
```
cd ~/prv-feast/examples/kubeflow-demo/files/kubeflow-alice
```

Then, Alice wants to push these definitions to the Feast Registry. Thus,
she creates a `feature_store.yaml` file with all the configuration and her
credentials and runs `feast apply`.
```
feast apply
```

Finally, she executes the materialization process. This process moves the
latest feature data from the Offline Store to the Online Store.
```
CURRENT_TIME=$(date -u +"%Y-%m-%dT%H:%M:%S") && feast materialize-incremental $CURRENT_TIME
```

At this point, the features are only visible to Alice, but she wants to
share them with Bob. To do so, she must share:
1. the definitions
2. the historical features of the Offline Store
3. the latest features of the Online Store

To share the definitions, she creates a new Role and binds it to Bob. The
`view-driver-daily-stats` Role allows users to view the `driver_daily_stats_fv`
and `driver` definitions. Inspect the `share.yaml` file to find out more.
```
kubectl apply -f share.yaml
```

To share the actual feature data, she instructs the admin of PostgreSQL
and Redis to give appropriate permissions.

The admin connects to the Kubernetes cluster (MiniKF) and runs the
following commands:
```
kubectl exec -it feast-postgresql-sts-0 -- psql postgresql://postgres:postgres@localhost:5432/driver_data
```

```
GRANT CREATE ON SCHEMA driver_data TO bob;
GRANT SELECT ON TABLE driver_data.driver_daily_stats TO bob;
exit
```

```
kubectl exec feast-redis-sts-0 -it -- redis-cli --user default -a redis
```

```
ACL SETUSER bob (-@all +@read ~*driver*kubeflow-alice)
exit
```

Notes:

In PostgreSQL, we need to allow `CREATE` on the `driver_data` schema. When
we execute the Feast `get_historical_features()` function, the Feast client
sends the `entity_df` DataFrame to the database. Thus, we need to create
a new table to store it temporarily. Then, the database executes the join
and sends the resulting DataFrame back to the client.

In Redis, we can restrict access based on key patterns. Feast stores the
latest feature data in Redis based on entity keys. The key that Feast stores
in Redis looks like the following:

`xxxdriverxxx1001xxxkubeflow-alice`

Thus, we can only restrict access based on entities. If a driver appears
in multiple feature views, we can allow access to all of them or none.

### Step 2: Charlie
Connect to Charlie's Notebook Server before continuing.

Charlie sets up a streaming system that receives the latest locations of
drivers, pushes them to the Online Store, and logs them to the Offline Store.
He creates a few definitions specifying the data sources, the entity, and
the feature view. He also defines an on-demand feature view that encodes
the latitude and longitude of a driver into a number that describes the
current zone of a driver.

Navigate to Charlie's folder and inspect the `driver_locations.py` file.
```
cd ~/prv-feast/examples/kubeflow-demo/files/kubeflow-charlie
```

Then, Charlie wants to push these definitions to the Feast Registry. Thus,
he creates a `feature_store.yaml` file with all the configuration and his
credentials and runs `feast apply`.
```
feast apply
```

Finally, he starts the streaming system. He opens a new terminal and runs:
```
pip install kafka-python
cd ~/prv-feast/examples/kubeflow-demo/files/stream
python3 consume_stream.py
```

He opens another terminal and runs:
```
cd ~/prv-feast/examples/kubeflow-demo/files/stream
python3 produce_stream.py
```

At this point, the features are only visible to Charlie, but he wants to
share them with Bob. He follows the same procedure as before.

To share the definitions, he creates a new Role and binds it to Bob. The
`view-driver-locations` Role allows users to view the `driver_locations_fv`,
`driver_zones_odfv`, and `driver` definitions. Inspect the `share.yaml`
file to find out more.

Go back to the initial terminal and run:
```
kubectl apply -f share.yaml
```

To share the actual feature data, he instructs the admin of PostgreSQL
and Redis to give appropriate permissions.

The admin performs the following actions inside the Kubernetes cluster:
```
kubectl exec -it feast-postgresql-sts-0 -- psql postgresql://postgres:postgres@localhost:5432/driver_stream_data
```

```
GRANT CREATE ON SCHEMA driver_stream_data TO bob;
GRANT SELECT ON TABLE driver_stream_data.driver_locations TO bob;
exit
```

```
kubectl exec feast-redis-sts-0 -it -- redis-cli --user default -a redis
```

```
ACL SETUSER bob (-@all +@read ~*driver*kubeflow-charlie)
exit
```

### Step 3: Bob
Connect to Bob's Notebook Server before continuing.

Bob wants to discover the available features. He creates a new notebook
`discover_features.ipynb` and finds the features Alice and Charlie shared
with him.
```
cd ~/prv-feast/examples/kubeflow-demo/files/kubeflow-bob
```

Then, he uses those features to train a model. Run `train.ipynb` step by
step to find out more.
