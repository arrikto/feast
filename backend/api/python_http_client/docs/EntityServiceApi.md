# frs_api.EntityServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**entity_service_create_entity**](EntityServiceApi.md#entity_service_create_entity) | **POST** /CreateEntity | 
[**entity_service_delete_entity**](EntityServiceApi.md#entity_service_delete_entity) | **DELETE** /DeleteEntity | 
[**entity_service_get_entity**](EntityServiceApi.md#entity_service_get_entity) | **GET** /GetEntity | 
[**entity_service_list_entities**](EntityServiceApi.md#entity_service_list_entities) | **GET** /ListEntities | 
[**entity_service_update_entity**](EntityServiceApi.md#entity_service_update_entity) | **POST** /UpdateEntity | 


# **entity_service_create_entity**
> ApiEntity entity_service_create_entity(body)



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
    api_instance = frs_api.EntityServiceApi(api_client)
    body = frs_api.ApiEntity() # ApiEntity | 

    try:
        api_response = api_instance.entity_service_create_entity(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling EntityServiceApi->entity_service_create_entity: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiEntity**](ApiEntity.md)|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

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

# **entity_service_delete_entity**
> object entity_service_delete_entity(name=name, project=project)



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
    api_instance = frs_api.EntityServiceApi(api_client)
    name = 'name_example' # str |  (optional)
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.entity_service_delete_entity(name=name, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling EntityServiceApi->entity_service_delete_entity: %s\n" % e)
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

# **entity_service_get_entity**
> ApiEntity entity_service_get_entity(name=name, project=project)



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
    api_instance = frs_api.EntityServiceApi(api_client)
    name = 'name_example' # str |  (optional)
project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.entity_service_get_entity(name=name, project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling EntityServiceApi->entity_service_get_entity: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**|  | [optional] 
 **project** | **str**|  | [optional] 

### Return type

[**ApiEntity**](ApiEntity.md)

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

# **entity_service_list_entities**
> ApiListEntitiesResponse entity_service_list_entities(project=project)



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
    api_instance = frs_api.EntityServiceApi(api_client)
    project = 'project_example' # str |  (optional)

    try:
        api_response = api_instance.entity_service_list_entities(project=project)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling EntityServiceApi->entity_service_list_entities: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **str**|  | [optional] 

### Return type

[**ApiListEntitiesResponse**](ApiListEntitiesResponse.md)

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

# **entity_service_update_entity**
> ApiEntity entity_service_update_entity(body)



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
    api_instance = frs_api.EntityServiceApi(api_client)
    body = frs_api.ApiEntity() # ApiEntity | 

    try:
        api_response = api_instance.entity_service_update_entity(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling EntityServiceApi->entity_service_update_entity: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiEntity**](ApiEntity.md)|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

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

