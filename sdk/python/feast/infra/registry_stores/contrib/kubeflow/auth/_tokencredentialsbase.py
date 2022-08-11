import abc

from kubernetes.client import configuration


class TokenCredentialsBase(abc.ABC):
    @abc.abstractmethod
    def refresh_api_key_hook(self, config: configuration.Configuration):
        """Refresh the api key.
        This is a helper function for registering token refresh with swagger
        generated clients.
        All classes that inherit from TokenCredentialsBase must implement this
        method to refresh the credentials.
        Args:
            config (kubernetes.client.configuration.Configuration):
                The configuration object that the client uses.
                The Configuration object of the kubernetes client's is the same
                with frs_api.configuration.Configuration.
        """
        raise NotImplementedError()


def read_token_from_file(path=None):
    """Read a token found in some file."""
    token = None
    with open(path, "r") as f:
        token = f.read().strip()
    return token
