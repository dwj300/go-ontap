# AggregateCloudStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachEligible** | **bool** | Aggregate is eligible for a cloud store to be attached. | [optional] [default to null]
**Stores** | [**[]CloudStorageTier**](cloud_storage_tier.md) | Configuration information for each cloud storage portion of the aggregate. | [optional] [default to null]
**TieringFullnessThreshold** | **int32** | The percentage of space in the performance tier that must be used before data is tiered out to the cloud store. Only valid for PATCH operations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


