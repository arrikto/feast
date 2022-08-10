# ApiFeatureView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the feature view. Must be unique. Not updated. | [optional] 
**project** | **str** | Name of Feast project that this feature view belongs to. | [optional] 
**entities** | **list[str]** | List names of entities to associate with the features defined in this feature view. Not updatable. | [optional] 
**features** | [**list[ApiFeature]**](ApiFeature.md) | List of specifications for each field defined as part of this feature view. | [optional] 
**description** | **str** | Description of the feature view. | [optional] 
**tags** | **dict(str, str)** | User defined metadata. | [optional] 
**owner** | **str** | Owner of the feature view. | [optional] 
**ttl** | **str** | Features in this feature view can only be retrieved from online serving younger than ttl. Ttl is measured as the duration of time between the feature&#39;s event timestamp and when the feature is retrieved. Feature values outside ttl will be returned as unset values and indicated to end user. | [optional] 
**batch_source** | **str** | Batch/Offline DataSource where this view can retrieve offline feature data. Protobuf object transformed to a JSON string. | [optional] 
**stream_source** | **str** | Streaming DataSource from where this view can consume \&quot;online\&quot; feature data. Protobuf object transformed to a JSON string. | [optional] 
**online** | **bool** | Whether these features should be served online or not. | [optional] 
**created_timestamp** | **datetime** | Creation time of the feature view. | [optional] 
**last_updated_timestamp** | **datetime** | Last update time of the feature view. | [optional] 
**materialization_intervals** | [**list[ApiMaterializationInterval]**](ApiMaterializationInterval.md) | List of pairs (start_time, end_time) for which this feature view has been materialized. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


