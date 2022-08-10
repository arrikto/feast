# coding: utf-8

"""
    Feast Registry API

    This file contains REST API specification for Feast Registry. The file is autogenerated from the swagger definition.  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import frs_api
from frs_api.api.feature_view_service_api import FeatureViewServiceApi  # noqa: E501
from frs_api.rest import ApiException


class TestFeatureViewServiceApi(unittest.TestCase):
    """FeatureViewServiceApi unit test stubs"""

    def setUp(self):
        self.api = frs_api.api.feature_view_service_api.FeatureViewServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_feature_view_service_create_feature_view(self):
        """Test case for feature_view_service_create_feature_view

        """
        pass

    def test_feature_view_service_delete_feature_view(self):
        """Test case for feature_view_service_delete_feature_view

        """
        pass

    def test_feature_view_service_get_feature_view(self):
        """Test case for feature_view_service_get_feature_view

        """
        pass

    def test_feature_view_service_list_feature_views(self):
        """Test case for feature_view_service_list_feature_views

        """
        pass

    def test_feature_view_service_report_mi(self):
        """Test case for feature_view_service_report_mi

        """
        pass

    def test_feature_view_service_update_feature_view(self):
        """Test case for feature_view_service_update_feature_view

        """
        pass


if __name__ == '__main__':
    unittest.main()
