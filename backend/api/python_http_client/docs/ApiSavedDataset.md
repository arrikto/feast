# ApiSavedDataset

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the dataset. Must be unique since it&#39;s possible to overwrite dataset by name. | [optional] 
**project** | **str** | Name of Feast project that this dataset belongs to. | [optional] 
**features** | **list[str]** | List of feature references with format \&quot;&lt;view name&gt;:&lt;feature name&gt;\&quot;. | [optional] 
**join_keys** | **list[str]** | Entity columns + request columns from all feature views used during retrieval. | [optional] 
**full_feature_names** | **bool** | Whether full feature names are used in stored data. | [optional] 
**storage** | **str** | Storage location of the saved dataset. Protobuf object transformed to a JSON string. | [optional] 
**feature_service_name** | **str** | Optional and only populated if generated from a feature service fetch. | [optional] 
**tags** | **dict(str, str)** | User defined metadata. | [optional] 
**created_timestamp** | **datetime** | Creation time of the saved dataset. | [optional] 
**last_updated_timestamp** | **datetime** | Last update time of the saved dataset. | [optional] 
**min_event_timestamp** | **datetime** | Min timestamp in the dataset (needed for retrieval). | [optional] 
**max_event_timestamp** | **datetime** | Max timestamp in the dataset (needed for retrieval). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


