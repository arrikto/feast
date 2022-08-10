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
from frs_api.models.api_saved_dataset import ApiSavedDataset  # noqa: E501
from frs_api.rest import ApiException

class TestApiSavedDataset(unittest.TestCase):
    """ApiSavedDataset unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ApiSavedDataset
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = frs_api.models.api_saved_dataset.ApiSavedDataset()  # noqa: E501
        if include_optional :
            return ApiSavedDataset(
                name = '0', 
                project = '0', 
                features = [
                    '0'
                    ], 
                join_keys = [
                    '0'
                    ], 
                full_feature_names = True, 
                storage = '0', 
                feature_service_name = '0', 
                tags = {
                    'key' : '0'
                    }, 
                created_timestamp = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                last_updated_timestamp = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                min_event_timestamp = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                max_event_timestamp = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else :
            return ApiSavedDataset(
        )

    def testApiSavedDataset(self):
        """Test ApiSavedDataset"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
