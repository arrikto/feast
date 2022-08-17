# Configure Feature Stores Help Guide

## Existing Configuration

Currently, users must provide a `path` to configure the registry. They can
also provide a `registry_store_type` and a `cache_ttl_seconds` parameter.

To use the Feast Registry Server we don't need the `path` parameter and
we replace it with `host` and `port` parameters. However, we will set
`registry_store_type` to `KubeflowRegistryStore` and keep using `cache_ttl_seconds` as before.

## Extra Configuration

The Kubeflow configuration has four fields: `host`, `port`, `project` and
`readMode`. However, `host`, `port` and `readMode` have default values,
`feast-registry-svc.kubeflow.svc.cluster.local`, `8888` and `False`
respectively. If a user doesn't provide the `project` parameter, the Feast
client will try to find it in the repo config file, `feature_store.yaml`,
and if not, it will fail.

The purpose of having a `readMode` parameter is to specify whether a user
has access to modify the project's metadata and infrastructure. If a user
doesn't have access, then set `readMode` to `True`.

## Examples

### YAML file

`feature_store.yaml`

```
project: kubeflow-alice
provider: local
registry:
  registry_store_type: KubeflowRegistryStore
  path: ""
```

### Python SDK

```
from feast import FeatureStore
from feast.repo_config import RepoConfig, RegistryConfig

registry_config_alice = RegistryConfig(
    registry_store_type="KubeflowRegistryStore",
    path="",
    project="kubeflow-alice",
    readMode=True
)

repo_config_alice = RepoConfig(
    project="kubeflow-alice",
    registry=registry_config_alice,
    provider="local"
)

fs_alice = FeatureStore(config=repo_config_alice, repo_path=None)
```
