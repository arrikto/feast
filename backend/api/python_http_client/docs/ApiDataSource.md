# ApiDataSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name of data source within the project. | [optional] 
**project** | **str** | Name of Feast project that this data source belongs to. | [optional] 
**description** | **str** | Description of the data source. | [optional] 
**tags** | **dict(str, str)** | User defined metadata. | [optional] 
**owner** | **str** | Owner of the data source. | [optional] 
**type** | [**DataSourceSourceType**](DataSourceSourceType.md) |  | [optional] 
**field_mapping** | **dict(str, str)** | Defines mapping between fields in the sourced data and fields in parent FeatureView. | [optional] 
**timestamp_field** | **str** | Event timestamp column name. | [optional] 
**date_partition_column** | **str** | Partition column (useful for file sources). | [optional] 
**created_timestamp_column** | **str** | Creation timestamp column name. | [optional] 
**data_source_class_type** | **str** | This is an internal field that represents the Python class of the data source object a proto object represents. This should be set by Feast, and not by users. The field is used primarily by custom data sources and is mandatory for them to set. Feast may set it for first party sources as well. | [optional] 
**batch_source** | **str** | Optional batch source for streaming sources for historical features and materialization. Protobuf object transformed to a JSON string. | [optional] 
**options** | **str** | DataSource options. Protobuf object transformed to a JSON string. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


