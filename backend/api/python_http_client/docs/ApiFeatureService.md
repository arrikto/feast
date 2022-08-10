# ApiFeatureService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the feature service. Must be unique. Not updated. | [optional] 
**project** | **str** | Name of Feast project that this feature service belongs to. | [optional] 
**features** | [**list[ApiFeatureViewProjection]**](ApiFeatureViewProjection.md) | Represents a projection that&#39;s to be applied on top of the FeatureView. Contains data such as the features to use from a FeatureView. | [optional] 
**tags** | **dict(str, str)** | User defined metadata. | [optional] 
**description** | **str** | Description of the feature service. | [optional] 
**owner** | **str** | Owner of the feature service. | [optional] 
**logging_config** | **str** | (optional) If provided logging will be enabled for this feature service. Protobuf object transformed to a JSON string. | [optional] 
**created_timestamp** | **datetime** | Creation time of the feature service. | [optional] 
**last_updated_timestamp** | **datetime** | Last update time of the feature service. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


