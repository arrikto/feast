# frs_api.RequestFeatureViewServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**request_feature_view_service_create_request_feature_view**](RequestFeatureViewServiceApi.md#request_feature_view_service_create_request_feature_view) | **POST** /CreateRequestFeatureView | 
[**request_feature_view_service_delete_request_feature_view**](RequestFeatureViewServiceApi.md#request_feature_view_service_delete_request_feature_view) | **DELETE** /DeleteRequestFeatureView | 
[**request_feature_view_service_get_request_feature_view**](RequestFeatureViewServiceApi.md#request_feature_view_service_get_request_feature_view) | **GET** /GetRequestFeatureView | 
[**request_feature_view_service_list_request_feature_views**](RequestFeatureViewServiceApi.md#request_feature_view_service_list_request_feature_views) | **GET** /ListRequestFeatureViews | 
[**request_feature_view_service_update_request_feature_view**](RequestFeatureViewServiceApi.md#request_feature_view_service_update_request_feature_view) | **POST** /UpdateRequestFeatureView | 


# **request_feature_view_service_create_request_feature_view**
> ApiRequestFeatureView request_feature_view_service_create_request_feature_view(body)



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
    api_instance = frs_api.RequestFeatureViewServiceApi(api_client)
    body = frs_api.ApiRequestFeatureView() # ApiRequestFeatureView | 

    try:
        api_response = api_instance.request_feature_view_service_create_request_feature_view(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling RequestFeatureViewServiceApi->request_feature_view_service_create_request_feature_view: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRequestFeatureView**](ApiRequestFeatureView.md)|  | 

### Return type

[**ApiRequestFeatureView**](ApiRequestFeatureView.md)

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

# **request_feature_view_service_delete_request_feature_view**
> object request_feature_view_service_delete_request_feature_view(name=name, project=project)



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
    api_instance = frs_api.RequestFeatureViewServiceApi(api_client)
    name = 'name_example' # str |  (optional)
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.request_feature_view_service_delete_request_feature_view(name=name, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling RequestFeatureViewServiceApi->request_feature_view_service_delete_request_feature_view: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**|  | [optional] 
 **project** | **str**|  | [optional] 

### Return type

**object**

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

# **request_feature_view_service_get_request_feature_view**
> ApiRequestFeatureView request_feature_view_service_get_request_feature_view(name=name, project=project)



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
    api_instance = frs_api.RequestFeatureViewServiceApi(api_client)
    name = 'name_example' # str |  (optional)
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.request_feature_view_service_get_request_feature_view(name=name, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling RequestFeatureViewServiceApi->request_feature_view_service_get_request_feature_view: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**|  | [optional] 
 **project** | **str**|  | [optional] 

### Return type

[**ApiRequestFeatureView**](ApiRequestFeatureView.md)

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

# **request_feature_view_service_list_request_feature_views**
> ApiListRequestFeatureViewsResponse request_feature_view_service_list_request_feature_views(project=project)



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
    api_instance = frs_api.RequestFeatureViewServiceApi(api_client)
    project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.request_feature_view_service_list_request_feature_views(project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling RequestFeatureViewServiceApi->request_feature_view_service_list_request_feature_views: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **str**|  | [optional] 

### Return type

[**ApiListRequestFeatureViewsResponse**](ApiListRequestFeatureViewsResponse.md)

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

# **request_feature_view_service_update_request_feature_view**
> ApiRequestFeatureView request_feature_view_service_update_request_feature_view(body)



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
    api_instance = frs_api.RequestFeatureViewServiceApi(api_client)
    body = frs_api.ApiRequestFeatureView() # ApiRequestFeatureView | 

    try:
        api_response = api_instance.request_feature_view_service_update_request_feature_view(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling RequestFeatureViewServiceApi->request_feature_view_service_update_request_feature_view: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRequestFeatureView**](ApiRequestFeatureView.md)|  | 

### Return type

[**ApiRequestFeatureView**](ApiRequestFeatureView.md)

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

