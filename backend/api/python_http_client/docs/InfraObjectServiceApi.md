# frs_api.InfraObjectServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**infra_object_service_list_infra_objects**](InfraObjectServiceApi.md#infra_object_service_list_infra_objects) | **GET** /ListInfraObjects | 
[**infra_object_service_update_infra_objects**](InfraObjectServiceApi.md#infra_object_service_update_infra_objects) | **POST** /UpdateInfraObjects | 


# **infra_object_service_list_infra_objects**
> ApiListInfraObjectsResponse infra_object_service_list_infra_objects(project=project)



### Example

* Api Key Authentication (Bearer):
```python
from __future__ import print_function
import time
import frs_api
from frs_api.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = frs_api.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: Bearer
configuration = frs_api.Configuration(
    host = "http://localhost",
    api_key = {
        'authorization': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['authorization'] = 'Bearer'

# Enter a context with an instance of the API client
with frs_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = frs_api.InfraObjectServiceApi(api_client)
    project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.infra_object_service_list_infra_objects(project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling InfraObjectServiceApi->infra_object_service_list_infra_objects: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **str**|  | [optional] 

### Return type

[**ApiListInfraObjectsResponse**](ApiListInfraObjectsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **infra_object_service_update_infra_objects**
> ApiUpdateInfraObjectsResponse infra_object_service_update_infra_objects(body, project=project)



### Example

* Api Key Authentication (Bearer):
```python
from __future__ import print_function
import time
import frs_api
from frs_api.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = frs_api.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: Bearer
configuration = frs_api.Configuration(
    host = "http://localhost",
    api_key = {
        'authorization': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['authorization'] = 'Bearer'

# Enter a context with an instance of the API client
with frs_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = frs_api.InfraObjectServiceApi(api_client)
    body = frs_api.ApiInfraObjects() # ApiInfraObjects | 
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.infra_object_service_update_infra_objects(body, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling InfraObjectServiceApi->infra_object_service_update_infra_objects: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiInfraObjects**](ApiInfraObjects.md)|  | 
 **project** | **str**|  | [optional] 

### Return type

[**ApiUpdateInfraObjectsResponse**](ApiUpdateInfraObjectsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

