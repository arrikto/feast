# coding: utf-8

"""
    Feast Registry API

    This file contains REST API specification for Feast Registry. The file is autogenerated from the swagger definition.  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from frs_api.configuration import Configuration


class ApiSavedDataset(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'name': 'str',
        'project': 'str',
        'features': 'list[str]',
        'join_keys': 'list[str]',
        'full_feature_names': 'bool',
        'storage': 'str',
        'feature_service_name': 'str',
        'tags': 'dict(str, str)',
        'created_timestamp': 'datetime',
        'last_updated_timestamp': 'datetime',
        'min_event_timestamp': 'datetime',
        'max_event_timestamp': 'datetime'
    }

    attribute_map = {
        'name': 'name',
        'project': 'project',
        'features': 'features',
        'join_keys': 'join_keys',
        'full_feature_names': 'full_feature_names',
        'storage': 'storage',
        'feature_service_name': 'feature_service_name',
        'tags': 'tags',
        'created_timestamp': 'created_timestamp',
        'last_updated_timestamp': 'last_updated_timestamp',
        'min_event_timestamp': 'min_event_timestamp',
        'max_event_timestamp': 'max_event_timestamp'
    }

    def __init__(self, name=None, project=None, features=None, join_keys=None, full_feature_names=None, storage=None, feature_service_name=None, tags=None, created_timestamp=None, last_updated_timestamp=None, min_event_timestamp=None, max_event_timestamp=None, local_vars_configuration=None):  # noqa: E501
        """ApiSavedDataset - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._name = None
        self._project = None
        self._features = None
        self._join_keys = None
        self._full_feature_names = None
        self._storage = None
        self._feature_service_name = None
        self._tags = None
        self._created_timestamp = None
        self._last_updated_timestamp = None
        self._min_event_timestamp = None
        self._max_event_timestamp = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if project is not None:
            self.project = project
        if features is not None:
            self.features = features
        if join_keys is not None:
            self.join_keys = join_keys
        if full_feature_names is not None:
            self.full_feature_names = full_feature_names
        if storage is not None:
            self.storage = storage
        if feature_service_name is not None:
            self.feature_service_name = feature_service_name
        if tags is not None:
            self.tags = tags
        if created_timestamp is not None:
            self.created_timestamp = created_timestamp
        if last_updated_timestamp is not None:
            self.last_updated_timestamp = last_updated_timestamp
        if min_event_timestamp is not None:
            self.min_event_timestamp = min_event_timestamp
        if max_event_timestamp is not None:
            self.max_event_timestamp = max_event_timestamp

    @property
    def name(self):
        """Gets the name of this ApiSavedDataset.  # noqa: E501

        Name of the dataset. Must be unique since it's possible to overwrite dataset by name.  # noqa: E501

        :return: The name of this ApiSavedDataset.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this ApiSavedDataset.

        Name of the dataset. Must be unique since it's possible to overwrite dataset by name.  # noqa: E501

        :param name: The name of this ApiSavedDataset.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def project(self):
        """Gets the project of this ApiSavedDataset.  # noqa: E501

        Name of Feast project that this dataset belongs to.  # noqa: E501

        :return: The project of this ApiSavedDataset.  # noqa: E501
        :rtype: str
        """
        return self._project

    @project.setter
    def project(self, project):
        """Sets the project of this ApiSavedDataset.

        Name of Feast project that this dataset belongs to.  # noqa: E501

        :param project: The project of this ApiSavedDataset.  # noqa: E501
        :type: str
        """

        self._project = project

    @property
    def features(self):
        """Gets the features of this ApiSavedDataset.  # noqa: E501

        List of feature references with format \"<view name>:<feature name>\".  # noqa: E501

        :return: The features of this ApiSavedDataset.  # noqa: E501
        :rtype: list[str]
        """
        return self._features

    @features.setter
    def features(self, features):
        """Sets the features of this ApiSavedDataset.

        List of feature references with format \"<view name>:<feature name>\".  # noqa: E501

        :param features: The features of this ApiSavedDataset.  # noqa: E501
        :type: list[str]
        """

        self._features = features

    @property
    def join_keys(self):
        """Gets the join_keys of this ApiSavedDataset.  # noqa: E501

        Entity columns + request columns from all feature views used during retrieval.  # noqa: E501

        :return: The join_keys of this ApiSavedDataset.  # noqa: E501
        :rtype: list[str]
        """
        return self._join_keys

    @join_keys.setter
    def join_keys(self, join_keys):
        """Sets the join_keys of this ApiSavedDataset.

        Entity columns + request columns from all feature views used during retrieval.  # noqa: E501

        :param join_keys: The join_keys of this ApiSavedDataset.  # noqa: E501
        :type: list[str]
        """

        self._join_keys = join_keys

    @property
    def full_feature_names(self):
        """Gets the full_feature_names of this ApiSavedDataset.  # noqa: E501

        Whether full feature names are used in stored data.  # noqa: E501

        :return: The full_feature_names of this ApiSavedDataset.  # noqa: E501
        :rtype: bool
        """
        return self._full_feature_names

    @full_feature_names.setter
    def full_feature_names(self, full_feature_names):
        """Sets the full_feature_names of this ApiSavedDataset.

        Whether full feature names are used in stored data.  # noqa: E501

        :param full_feature_names: The full_feature_names of this ApiSavedDataset.  # noqa: E501
        :type: bool
        """

        self._full_feature_names = full_feature_names

    @property
    def storage(self):
        """Gets the storage of this ApiSavedDataset.  # noqa: E501

        Storage location of the saved dataset. Protobuf object transformed to a JSON string.  # noqa: E501

        :return: The storage of this ApiSavedDataset.  # noqa: E501
        :rtype: str
        """
        return self._storage

    @storage.setter
    def storage(self, storage):
        """Sets the storage of this ApiSavedDataset.

        Storage location of the saved dataset. Protobuf object transformed to a JSON string.  # noqa: E501

        :param storage: The storage of this ApiSavedDataset.  # noqa: E501
        :type: str
        """

        self._storage = storage

    @property
    def feature_service_name(self):
        """Gets the feature_service_name of this ApiSavedDataset.  # noqa: E501

        Optional and only populated if generated from a feature service fetch.  # noqa: E501

        :return: The feature_service_name of this ApiSavedDataset.  # noqa: E501
        :rtype: str
        """
        return self._feature_service_name

    @feature_service_name.setter
    def feature_service_name(self, feature_service_name):
        """Sets the feature_service_name of this ApiSavedDataset.

        Optional and only populated if generated from a feature service fetch.  # noqa: E501

        :param feature_service_name: The feature_service_name of this ApiSavedDataset.  # noqa: E501
        :type: str
        """

        self._feature_service_name = feature_service_name

    @property
    def tags(self):
        """Gets the tags of this ApiSavedDataset.  # noqa: E501

        User defined metadata.  # noqa: E501

        :return: The tags of this ApiSavedDataset.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._tags

    @tags.setter
    def tags(self, tags):
        """Sets the tags of this ApiSavedDataset.

        User defined metadata.  # noqa: E501

        :param tags: The tags of this ApiSavedDataset.  # noqa: E501
        :type: dict(str, str)
        """

        self._tags = tags

    @property
    def created_timestamp(self):
        """Gets the created_timestamp of this ApiSavedDataset.  # noqa: E501

        Creation time of the saved dataset.  # noqa: E501

        :return: The created_timestamp of this ApiSavedDataset.  # noqa: E501
        :rtype: datetime
        """
        return self._created_timestamp

    @created_timestamp.setter
    def created_timestamp(self, created_timestamp):
        """Sets the created_timestamp of this ApiSavedDataset.

        Creation time of the saved dataset.  # noqa: E501

        :param created_timestamp: The created_timestamp of this ApiSavedDataset.  # noqa: E501
        :type: datetime
        """

        self._created_timestamp = created_timestamp

    @property
    def last_updated_timestamp(self):
        """Gets the last_updated_timestamp of this ApiSavedDataset.  # noqa: E501

        Last update time of the saved dataset.  # noqa: E501

        :return: The last_updated_timestamp of this ApiSavedDataset.  # noqa: E501
        :rtype: datetime
        """
        return self._last_updated_timestamp

    @last_updated_timestamp.setter
    def last_updated_timestamp(self, last_updated_timestamp):
        """Sets the last_updated_timestamp of this ApiSavedDataset.

        Last update time of the saved dataset.  # noqa: E501

        :param last_updated_timestamp: The last_updated_timestamp of this ApiSavedDataset.  # noqa: E501
        :type: datetime
        """

        self._last_updated_timestamp = last_updated_timestamp

    @property
    def min_event_timestamp(self):
        """Gets the min_event_timestamp of this ApiSavedDataset.  # noqa: E501

        Min timestamp in the dataset (needed for retrieval).  # noqa: E501

        :return: The min_event_timestamp of this ApiSavedDataset.  # noqa: E501
        :rtype: datetime
        """
        return self._min_event_timestamp

    @min_event_timestamp.setter
    def min_event_timestamp(self, min_event_timestamp):
        """Sets the min_event_timestamp of this ApiSavedDataset.

        Min timestamp in the dataset (needed for retrieval).  # noqa: E501

        :param min_event_timestamp: The min_event_timestamp of this ApiSavedDataset.  # noqa: E501
        :type: datetime
        """

        self._min_event_timestamp = min_event_timestamp

    @property
    def max_event_timestamp(self):
        """Gets the max_event_timestamp of this ApiSavedDataset.  # noqa: E501

        Max timestamp in the dataset (needed for retrieval).  # noqa: E501

        :return: The max_event_timestamp of this ApiSavedDataset.  # noqa: E501
        :rtype: datetime
        """
        return self._max_event_timestamp

    @max_event_timestamp.setter
    def max_event_timestamp(self, max_event_timestamp):
        """Sets the max_event_timestamp of this ApiSavedDataset.

        Max timestamp in the dataset (needed for retrieval).  # noqa: E501

        :param max_event_timestamp: The max_event_timestamp of this ApiSavedDataset.  # noqa: E501
        :type: datetime
        """

        self._max_event_timestamp = max_event_timestamp

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ApiSavedDataset):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ApiSavedDataset):
            return True

        return self.to_dict() != other.to_dict()