# CloudStore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Availability** | **string** | Availability of the object store. | [optional] [default to null]
**MirrorDegraded** | **bool** | This field identifies if the mirror cloud store is in sync with the primary cloud store of a FabricPool. | [optional] [default to null]
**Primary** | **bool** | This field indicates whether the cloud store is the primary cloud store of a cloud mirrored composite aggregate. | [optional] [default to null]
**Target** | [***CloudStoreTarget**](cloud_store_target.md) |  | [optional] [default to null]
**UnreclaimedSpaceThreshold** | **int32** | Usage threshold for reclaiming unused space in the cloud store. Valid values are 0 to 99. The default value depends on the provider type. This can be specified in PATCH but not POST. | [optional] [default to null]
**Used** | **int32** | The amount of object space used. Calculated every 5 minutes and cached. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


