# ApplicationStatisticsSpace

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **int32** | The available amount of space left in the application component. Note that this field has limited meaning for SAN applications. Space may be considered used from ONTAP&#39;s perspective while the host filesystem still considers it available | [optional] [default to null]
**LogicalUsed** | **int32** | The amount of space that would currently be used if no space saving features were enabled. For example, if compression were the only space saving feature enabled, this field would represent the uncompressed amount of space used | [optional] [default to null]
**Provisioned** | **int32** | The originally requested amount of space that was provisioned for the application component | [optional] [default to null]
**ReservedUnused** | **int32** | The amount of space reserved for system features such as snapshots that has not yet been used | [optional] [default to null]
**Savings** | **int32** | The amount of space saved by all enabled space saving features | [optional] [default to null]
**Used** | **int32** | The amount of space that is currently being used by the application component. Note that this includes any space reserved by the system for features such as snapshots | [optional] [default to null]
**UsedExcludingReserves** | **int32** | The amount of space that is currently being used, excluding any space that is reserved by the system for features such as snapshots | [optional] [default to null]
**UsedPercent** | **int32** | The percentage of the originally provisioned space that is currently being used by the application component | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


