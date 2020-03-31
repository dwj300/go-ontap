# ApplicationStatisticsSpace1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **int32** | The available amount of space left in the application. Note that this field has limited meaning for SAN applications. Space may be considered used from ONTAP&#39;s perspective while the host filesystem still considers it available | [optional] [readonly] 
**LogicalUsed** | **int32** | The amount of space that would currently be used if no space saving features were enabled. For example, if compression were the only space saving feature enabled, this field would represent the uncompressed amount of space used | [optional] [readonly] 
**Provisioned** | **int32** | The originally requested amount of space that was provisioned for the application | [optional] [readonly] 
**ReservedUnused** | **int32** | The amount of space reserved for system features such as snapshots that has not yet been used | [optional] [readonly] 
**Savings** | **int32** | The amount of space saved by all enabled space saving features | [optional] [readonly] 
**Used** | **int32** | The amount of space that is currently being used by the application. Note that this includes any space reserved by the system for features such as snapshots | [optional] [readonly] 
**UsedExcludingReserves** | **int32** | The amount of space that is currently being used, excluding any space that is reserved by the system for features such as snapshots | [optional] [readonly] 
**UsedPercent** | **int32** | The percentage of the originally provisioned space that is currently being used by the application | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


