# AggregateSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorage** | [**AggregateSpaceBlockStorage**](aggregate_space_block_storage.md) |  | [optional] 
**CloudStorage** | [**AggregateSpaceCloudStorage**](aggregate_space_cloud_storage.md) |  | [optional] 
**Efficiency** | [**AggregateSpaceEfficiency**](aggregate_space_efficiency.md) |  | [optional] 
**EfficiencyWithoutSnapshots** | [**AggregateSpaceEfficiencyWithoutSnapshots**](aggregate_space_efficiency_without_snapshots.md) |  | [optional] 
**Footprint** | **int32** | A summation of volume footprints (including volume guarantees), in bytes. This includes all of the volume footprints in the block_storage tier and the cloud_storage tier. This is an advanced property; there is an added cost to retrieving its value. The field is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &lt;i&gt;fields&lt;/i&gt; query parameter containing either footprint or **.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


