# coding: utf-8

"""
    Feast Registry API

    This file contains REST API specification for Feast Registry. The file is autogenerated from the swagger definition.  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import frs_api
from frs_api.models.api_request_feature_view import ApiRequestFeatureView  # noqa: E501
from frs_api.rest import ApiException

class TestApiRequestFeatureView(unittest.TestCase):
    """ApiRequestFeatureView unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ApiRequestFeatureView
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = frs_api.models.api_request_feature_view.ApiRequestFeatureView()  # noqa: E501
        if include_optional :
            return ApiRequestFeatureView(
                name = '0', 
                project = '0', 
                request_data_source = '0', 
                description = '0', 
                tags = {
                    'key' : '0'
                    }, 
                owner = '0'
            )
        else :
            return ApiRequestFeatureView(
        )

    def testApiRequestFeatureView(self):
        """Test ApiRequestFeatureView"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()