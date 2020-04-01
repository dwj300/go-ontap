# AggregateBlockStorageHybridCache

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskCount** | **int32** | Number of disks used in the cache tier of the aggregate. Only provided when hybrid_cache.enabled is &#39;true&#39;. | [optional] [default to null]
**Enabled** | **bool** | Aggregate uses HDDs with SSDs as a cache | [optional] [default to null]
**RaidType** | **string** | RAID type for SSD cache of the aggregate. Only provided when hybrid_cache.enabled is &#39;true&#39;. | [optional] [default to null]
**Size** | **int32** | Total usable space in bytes of SSD cache. Only provided when hybrid_cache.enabled is &#39;true&#39;. | [optional] [default to null]
**Used** | **int32** | Space used in bytes of SSD cache. Only provided when hybrid_cache.enabled is &#39;true&#39;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


