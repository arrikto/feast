# frs_api.DataSourceServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**data_source_service_create_data_source**](DataSourceServiceApi.md#data_source_service_create_data_source) | **POST** /CreateDataSource | 
[**data_source_service_delete_data_source**](DataSourceServiceApi.md#data_source_service_delete_data_source) | **DELETE** /DeleteDataSource | 
[**data_source_service_get_data_source**](DataSourceServiceApi.md#data_source_service_get_data_source) | **GET** /GetDataSource | 
[**data_source_service_list_data_sources**](DataSourceServiceApi.md#data_source_service_list_data_sources) | **GET** /ListDataSources | 
[**data_source_service_update_data_source**](DataSourceServiceApi.md#data_source_service_update_data_source) | **POST** /UpdateDataSource | 


# **data_source_service_create_data_source**
> ApiDataSource data_source_service_create_data_source(body)



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
    api_instance = frs_api.DataSourceServiceApi(api_client)
    body = frs_api.ApiDataSource() # ApiDataSource | 

    try:
        api_response = api_instance.data_source_service_create_data_source(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DataSourceServiceApi->data_source_service_create_data_source: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiDataSource**](ApiDataSource.md)|  | 

### Return type

[**ApiDataSource**](ApiDataSource.md)

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

# **data_source_service_delete_data_source**
> object data_source_service_delete_data_source(name=name, project=project)



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
    api_instance = frs_api.DataSourceServiceApi(api_client)
    name = 'name_example' # str |  (optional)
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.data_source_service_delete_data_source(name=name, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DataSourceServiceApi->data_source_service_delete_data_source: %s\n" % e)
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

# **data_source_service_get_data_source**
> ApiDataSource data_source_service_get_data_source(name=name, project=project)



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
    api_instance = frs_api.DataSourceServiceApi(api_client)
    name = 'name_example' # str |  (optional)
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.data_source_service_get_data_source(name=name, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DataSourceServiceApi->data_source_service_get_data_source: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**|  | [optional] 
 **project** | **str**|  | [optional] 

### Return type

[**ApiDataSource**](ApiDataSource.md)

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

# **data_source_service_list_data_sources**
> ApiListDataSourcesResponse data_source_service_list_data_sources(project=project)



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
    api_instance = frs_api.DataSourceServiceApi(api_client)
    project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.data_source_service_list_data_sources(project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DataSourceServiceApi->data_source_service_list_data_sources: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **str**|  | [optional] 

### Return type

[**ApiListDataSourcesResponse**](ApiListDataSourcesResponse.md)

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

# **data_source_service_update_data_source**
> ApiDataSource data_source_service_update_data_source(body)



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
    api_instance = frs_api.DataSourceServiceApi(api_client)
    body = frs_api.ApiDataSource() # ApiDataSource | 

    try:
        api_response = api_instance.data_source_service_update_data_source(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DataSourceServiceApi->data_source_service_update_data_source: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiDataSource**](ApiDataSource.md)|  | 

### Return type

[**ApiDataSource**](ApiDataSource.md)

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

