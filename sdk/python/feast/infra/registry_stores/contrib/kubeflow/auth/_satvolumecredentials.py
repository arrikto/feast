import logging
import os

import frs_api

from ._consts import KF_FEATURES_SA_TOKEN_ENV, KF_FEATURES_SA_TOKEN_PATH
from ._tokencredentialsbase import TokenCredentialsBase, read_token_from_file


class ServiceAccountTokenVolumeCredentials(TokenCredentialsBase):
    """Audience-bound ServiceAccountToken in the local filesystem.
    This is a credentials interface for audience-bound ServiceAccountTokens
    found in the local filesystem, that get refreshed by the kubelet.
    The constructor of the class expects a filesystem path.
    If not provided, it uses the path stored in the environment variable
    defined in ``auth.KF_FEATURES_SA_TOKEN_ENV``.
    If the environment variable is also empty, it falls back to the path
    specified in ``auth.KF_FEATURES_SA_TOKEN_PATH``.
    This method of authentication is meant for use inside a Kubernetes cluster.
    Relevant documentation:
    https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#service-account-token-volume-projection
    """

    def __init__(self, path=None):
        self._token_path = (
            path or os.getenv(KF_FEATURES_SA_TOKEN_ENV) or KF_FEATURES_SA_TOKEN_PATH
        )

    def _get_token(self):
        token = None
        try:
            token = read_token_from_file(self._token_path)
        except OSError as e:
            logging.error(
                "Failed to read a token from file '%s' (%s).", self._token_path, str(e)
            )
            raise
        return token

    def refresh_api_key_hook(self, config: frs_api.Configuration):
        """Refresh the api key.
        This is a helper function for registering token refresh with swagger
        generated clients.
        Args:
            config (frs_api.Configuration): The configuration object
                that the client uses.
        """
        config.api_key["authorization"] = self._get_token()
