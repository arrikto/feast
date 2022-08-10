# ApiEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the entity. | [optional] 
**project** | **str** | Name of Feast project that this entity belongs to. | [optional] 
**value_type** | [**ValueTypeEnum**](ValueTypeEnum.md) |  | [optional] 
**description** | **str** | Description of the entity. | [optional] 
**join_key** | **str** | Join key for the entity (i.e. name of the column the entity maps to). | [optional] 
**tags** | **dict(str, str)** | User defined metadata. | [optional] 
**owner** | **str** | Owner of the entity. | [optional] 
**created_timestamp** | **datetime** | Creation time of the entity. | [optional] 
**last_updated_timestamp** | **datetime** | Last update time of the entity. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


