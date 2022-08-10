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


class ApiFeatureView(object):
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
        'entities': 'list[str]',
        'features': 'list[ApiFeature]',
        'description': 'str',
        'tags': 'dict(str, str)',
        'owner': 'str',
        'ttl': 'str',
        'batch_source': 'str',
        'stream_source': 'str',
        'online': 'bool',
        'created_timestamp': 'datetime',
        'last_updated_timestamp': 'datetime',
        'materialization_intervals': 'list[ApiMaterializationInterval]'
    }

    attribute_map = {
        'name': 'name',
        'project': 'project',
        'entities': 'entities',
        'features': 'features',
        'description': 'description',
        'tags': 'tags',
        'owner': 'owner',
        'ttl': 'ttl',
        'batch_source': 'batch_source',
        'stream_source': 'stream_source',
        'online': 'online',
        'created_timestamp': 'created_timestamp',
        'last_updated_timestamp': 'last_updated_timestamp',
        'materialization_intervals': 'materialization_intervals'
    }

    def __init__(self, name=None, project=None, entities=None, features=None, description=None, tags=None, owner=None, ttl=None, batch_source=None, stream_source=None, online=None, created_timestamp=None, last_updated_timestamp=None, materialization_intervals=None, local_vars_configuration=None):  # noqa: E501
        """ApiFeatureView - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._name = None
        self._project = None
        self._entities = None
        self._features = None
        self._description = None
        self._tags = None
        self._owner = None
        self._ttl = None
        self._batch_source = None
        self._stream_source = None
        self._online = None
        self._created_timestamp = None
        self._last_updated_timestamp = None
        self._materialization_intervals = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if project is not None:
            self.project = project
        if entities is not None:
            self.entities = entities
        if features is not None:
            self.features = features
        if description is not None:
            self.description = description
        if tags is not None:
            self.tags = tags
        if owner is not None:
            self.owner = owner
        if ttl is not None:
            self.ttl = ttl
        if batch_source is not None:
            self.batch_source = batch_source
        if stream_source is not None:
            self.stream_source = stream_source
        if online is not None:
            self.online = online
        if created_timestamp is not None:
            self.created_timestamp = created_timestamp
        if last_updated_timestamp is not None:
            self.last_updated_timestamp = last_updated_timestamp
        if materialization_intervals is not None:
            self.materialization_intervals = materialization_intervals

    @property
    def name(self):
        """Gets the name of this ApiFeatureView.  # noqa: E501

        Name of the feature view. Must be unique. Not updated.  # noqa: E501

        :return: The name of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this ApiFeatureView.

        Name of the feature view. Must be unique. Not updated.  # noqa: E501

        :param name: The name of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def project(self):
        """Gets the project of this ApiFeatureView.  # noqa: E501

        Name of Feast project that this feature view belongs to.  # noqa: E501

        :return: The project of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._project

    @project.setter
    def project(self, project):
        """Sets the project of this ApiFeatureView.

        Name of Feast project that this feature view belongs to.  # noqa: E501

        :param project: The project of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._project = project

    @property
    def entities(self):
        """Gets the entities of this ApiFeatureView.  # noqa: E501

        List names of entities to associate with the features defined in this feature view. Not updatable.  # noqa: E501

        :return: The entities of this ApiFeatureView.  # noqa: E501
        :rtype: list[str]
        """
        return self._entities

    @entities.setter
    def entities(self, entities):
        """Sets the entities of this ApiFeatureView.

        List names of entities to associate with the features defined in this feature view. Not updatable.  # noqa: E501

        :param entities: The entities of this ApiFeatureView.  # noqa: E501
        :type: list[str]
        """

        self._entities = entities

    @property
    def features(self):
        """Gets the features of this ApiFeatureView.  # noqa: E501

        List of specifications for each field defined as part of this feature view.  # noqa: E501

        :return: The features of this ApiFeatureView.  # noqa: E501
        :rtype: list[ApiFeature]
        """
        return self._features

    @features.setter
    def features(self, features):
        """Sets the features of this ApiFeatureView.

        List of specifications for each field defined as part of this feature view.  # noqa: E501

        :param features: The features of this ApiFeatureView.  # noqa: E501
        :type: list[ApiFeature]
        """

        self._features = features

    @property
    def description(self):
        """Gets the description of this ApiFeatureView.  # noqa: E501

        Description of the feature view.  # noqa: E501

        :return: The description of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._description

    @description.setter
    def description(self, description):
        """Sets the description of this ApiFeatureView.

        Description of the feature view.  # noqa: E501

        :param description: The description of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._description = description

    @property
    def tags(self):
        """Gets the tags of this ApiFeatureView.  # noqa: E501

        User defined metadata.  # noqa: E501

        :return: The tags of this ApiFeatureView.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._tags

    @tags.setter
    def tags(self, tags):
        """Sets the tags of this ApiFeatureView.

        User defined metadata.  # noqa: E501

        :param tags: The tags of this ApiFeatureView.  # noqa: E501
        :type: dict(str, str)
        """

        self._tags = tags

    @property
    def owner(self):
        """Gets the owner of this ApiFeatureView.  # noqa: E501

        Owner of the feature view.  # noqa: E501

        :return: The owner of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._owner

    @owner.setter
    def owner(self, owner):
        """Sets the owner of this ApiFeatureView.

        Owner of the feature view.  # noqa: E501

        :param owner: The owner of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._owner = owner

    @property
    def ttl(self):
        """Gets the ttl of this ApiFeatureView.  # noqa: E501

        Features in this feature view can only be retrieved from online serving younger than ttl. Ttl is measured as the duration of time between the feature's event timestamp and when the feature is retrieved. Feature values outside ttl will be returned as unset values and indicated to end user.  # noqa: E501

        :return: The ttl of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._ttl

    @ttl.setter
    def ttl(self, ttl):
        """Sets the ttl of this ApiFeatureView.

        Features in this feature view can only be retrieved from online serving younger than ttl. Ttl is measured as the duration of time between the feature's event timestamp and when the feature is retrieved. Feature values outside ttl will be returned as unset values and indicated to end user.  # noqa: E501

        :param ttl: The ttl of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._ttl = ttl

    @property
    def batch_source(self):
        """Gets the batch_source of this ApiFeatureView.  # noqa: E501

        Batch/Offline DataSource where this view can retrieve offline feature data. Protobuf object transformed to a JSON string.  # noqa: E501

        :return: The batch_source of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._batch_source

    @batch_source.setter
    def batch_source(self, batch_source):
        """Sets the batch_source of this ApiFeatureView.

        Batch/Offline DataSource where this view can retrieve offline feature data. Protobuf object transformed to a JSON string.  # noqa: E501

        :param batch_source: The batch_source of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._batch_source = batch_source

    @property
    def stream_source(self):
        """Gets the stream_source of this ApiFeatureView.  # noqa: E501

        Streaming DataSource from where this view can consume \"online\" feature data. Protobuf object transformed to a JSON string.  # noqa: E501

        :return: The stream_source of this ApiFeatureView.  # noqa: E501
        :rtype: str
        """
        return self._stream_source

    @stream_source.setter
    def stream_source(self, stream_source):
        """Sets the stream_source of this ApiFeatureView.

        Streaming DataSource from where this view can consume \"online\" feature data. Protobuf object transformed to a JSON string.  # noqa: E501

        :param stream_source: The stream_source of this ApiFeatureView.  # noqa: E501
        :type: str
        """

        self._stream_source = stream_source

    @property
    def online(self):
        """Gets the online of this ApiFeatureView.  # noqa: E501

        Whether these features should be served online or not.  # noqa: E501

        :return: The online of this ApiFeatureView.  # noqa: E501
        :rtype: bool
        """
        return self._online

    @online.setter
    def online(self, online):
        """Sets the online of this ApiFeatureView.

        Whether these features should be served online or not.  # noqa: E501

        :param online: The online of this ApiFeatureView.  # noqa: E501
        :type: bool
        """

        self._online = online

    @property
    def created_timestamp(self):
        """Gets the created_timestamp of this ApiFeatureView.  # noqa: E501

        Creation time of the feature view.  # noqa: E501

        :return: The created_timestamp of this ApiFeatureView.  # noqa: E501
        :rtype: datetime
        """
        return self._created_timestamp

    @created_timestamp.setter
    def created_timestamp(self, created_timestamp):
        """Sets the created_timestamp of this ApiFeatureView.

        Creation time of the feature view.  # noqa: E501

        :param created_timestamp: The created_timestamp of this ApiFeatureView.  # noqa: E501
        :type: datetime
        """

        self._created_timestamp = created_timestamp

    @property
    def last_updated_timestamp(self):
        """Gets the last_updated_timestamp of this ApiFeatureView.  # noqa: E501

        Last update time of the feature view.  # noqa: E501

        :return: The last_updated_timestamp of this ApiFeatureView.  # noqa: E501
        :rtype: datetime
        """
        return self._last_updated_timestamp

    @last_updated_timestamp.setter
    def last_updated_timestamp(self, last_updated_timestamp):
        """Sets the last_updated_timestamp of this ApiFeatureView.

        Last update time of the feature view.  # noqa: E501

        :param last_updated_timestamp: The last_updated_timestamp of this ApiFeatureView.  # noqa: E501
        :type: datetime
        """

        self._last_updated_timestamp = last_updated_timestamp

    @property
    def materialization_intervals(self):
        """Gets the materialization_intervals of this ApiFeatureView.  # noqa: E501

        List of pairs (start_time, end_time) for which this feature view has been materialized.  # noqa: E501

        :return: The materialization_intervals of this ApiFeatureView.  # noqa: E501
        :rtype: list[ApiMaterializationInterval]
        """
        return self._materialization_intervals

    @materialization_intervals.setter
    def materialization_intervals(self, materialization_intervals):
        """Sets the materialization_intervals of this ApiFeatureView.

        List of pairs (start_time, end_time) for which this feature view has been materialized.  # noqa: E501

        :param materialization_intervals: The materialization_intervals of this ApiFeatureView.  # noqa: E501
        :type: list[ApiMaterializationInterval]
        """

        self._materialization_intervals = materialization_intervals

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
        if not isinstance(other, ApiFeatureView):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ApiFeatureView):
            return True

        return self.to_dict() != other.to_dict()
