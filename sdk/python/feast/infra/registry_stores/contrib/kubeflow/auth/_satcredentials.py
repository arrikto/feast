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

import os
from datetime import datetime, timedelta

import frs_api
import jwt

from ._consts import KF_FEATURES_SA_TOKEN_ENV
from ._tokencredentialsbase import TokenCredentialsBase, read_token_from_file


class ServiceAccountTokenCredentials(TokenCredentialsBase):

    DEFAULT_AUDIENCES = ["istio-ingressgateway.istio-system.svc.cluster.local"]
    DEFAULT_LEEWAY = timedelta(minutes=1)

    def __init__(
        self,
        long_lived_token=None,
        long_lived_token_path=None,
        kubernetes_url=None,
        def_audiences=DEFAULT_AUDIENCES,
        def_expiration_seconds=3600,
    ):
        super().__init__()

        try:
            long_lived_token = long_lived_token or read_token_from_file(
                long_lived_token_path or os.getenv(KF_FEATURES_SA_TOKEN_ENV)
            )
        except Exception:
            pass
        if not long_lived_token:
            raise ValueError(
                "Please provide a long_lived_token either"
                " directly, or as a file pointed by the"
                " long_lived_token_path argument or the '%s'"
                " environment variable." % KF_FEATURES_SA_TOKEN_ENV
            )
        claims = jwt.decode(long_lived_token, options={"verify_signature": False})
        self.sa_name = claims["kubernetes.io/serviceaccount/service-account.name"]
        self.sa_namespace = claims["kubernetes.io/serviceaccount/namespace"]
        self.sa_secret = claims["kubernetes.io/serviceaccount/secret.name"]
        self.short_lived_tokens = {}
        self.def_audiences = def_audiences
        self.def_expiration_seconds = def_expiration_seconds

        if not kubernetes_url:
            raise ValueError("Please provide a kubernetes_url")
        config = frs_api.Configuration(
            host=kubernetes_url,
            api_key={"authorization": long_lived_token},
            api_key_prefix={"authorization": "Bearer"},
        )
        self.api_client = frs_api.ApiClient(configuration=config)

    def get_token(self, audiences=None, expiration_seconds=None):
        audiences = audiences or self.def_audiences
        expiration_seconds = expiration_seconds or self.def_expiration_seconds
        # Check if we already have a short-lived token and if it will be valid
        # for at least some more time (DEFAULT_LEEWAY).
        short_lived_token = self.short_lived_tokens.get(str(audiences), "")
        if short_lived_token:
            claims = jwt.decode(short_lived_token, options={"verify_signature": False})
            expiration_time = datetime.utcfromtimestamp(claims.get("exp"))
            if expiration_time > datetime.now() + self.DEFAULT_LEEWAY:
                return short_lived_token

        # Need to get a new token
        short_lived_token = self._generate_sa_token(audiences, expiration_seconds)
        self.short_lived_tokens[str(audiences)] = short_lived_token
        return short_lived_token

    def _generate_sa_token(self, audiences, expiration_seconds):
        """Generate a token with specific audience for service account."""
        # XXX: TokenRequest is not possible using the official or the
        # rok_kubernetes library
        api_client = self.api_client
        url_path = "/api/v1/namespaces/{namespace}/serviceaccounts/{name}" "/token"
        path_params = {"name": self.sa_name, "namespace": self.sa_namespace}
        query_params = []
        header_params = {
            "Accept": api_client.select_header_accept(
                [
                    "application/json",
                    "application/yaml",
                    "application/vnd.kubernetes.protobuf",
                ]
            )
        }
        auth_settings = ["Bearer"]
        body_params = {
            "spec": {
                "audiences": audiences,
                "expirationSeconds": expiration_seconds,
                "boundObjectRef": {
                    "apiVersion": "v1",
                    "kind": "Secret",
                    "name": self.sa_secret,
                },
            }
        }

        # The following API call will return a dict
        resp = api_client.call_api(
            url_path,
            "POST",
            path_params,
            query_params,
            header_params,
            body=body_params,
            auth_settings=auth_settings,
            response_type="object",
            _return_http_data_only=True,
        )
        token = resp.get("status", {}).get("token")
        return token

    def refresh_api_key_hook(self, config: frs_api.Configuration):
        """Refresh the api key.
        This is a helper function for registering token refresh with swagger
        generated clients.
        Args:
            config (frs_api.Configuration): The configuration object
                that the client uses.
        """
        config.api_key["authorization"] = self.get_token()
