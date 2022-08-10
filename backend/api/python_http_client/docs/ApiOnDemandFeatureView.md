# ApiOnDemandFeatureView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the on demand feature view. Must be unique. Not updated. | [optional] 
**project** | **str** | Name of Feast project that this on demand feature view belongs to. | [optional] 
**features** | [**list[ApiFeature]**](ApiFeature.md) | List of features specifications for each feature defined with this on demand feature view. | [optional] 
**sources** | **dict(str, str)** | Map of sources for this on demand feature view. Sources are transformed from Protobuf objects to JSON strings. | [optional] 
**user_defined_function** | [**ApiUserDefinedFunction**](ApiUserDefinedFunction.md) |  | [optional] 
**description** | **str** | Description of the on demand feature view. | [optional] 
**tags** | **dict(str, str)** | User defined metadata. | [optional] 
**owner** | **str** | Owner of the on demand feature view. | [optional] 
**created_timestamp** | **datetime** | Creation time of the on demand feature view. | [optional] 
**last_updated_timestamp** | **datetime** | Last update time of the on demand feature view. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


