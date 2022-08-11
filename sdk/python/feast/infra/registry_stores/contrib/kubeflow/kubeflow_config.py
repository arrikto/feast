from pydantic import StrictStr
from pydantic.typing import Optional

from feast.repo_config import FeastConfigBaseModel


class KubeflowRegistryConfig(FeastConfigBaseModel):
    host: StrictStr = "feast-registry-svc.kubeflow.svc.cluster.local"
    port: int = 8888
    project: StrictStr = ""
    readMode: Optional[bool]
