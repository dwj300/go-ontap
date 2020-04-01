# AggregateSpaceBlockStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **int32** | Space available in bytes | [optional] [default to null]
**FullThresholdPercent** | **int32** | The aggregate used percentage at which &#39;monitor.volume.full&#39; EMS is generated. | [optional] [default to null]
**InactiveUserData** | **int32** | The size that is physically used in the block storage and has a cold temperature, in bytes. This property is only supported if the aggregate is either attached to a cloud store or can be attached to a cloud store. This is an advanced property; there is an added cost to retrieving its value. The field is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &lt;i&gt;fields&lt;/i&gt; query parameter containing either block_storage.inactive_user_data or **.  | [optional] [default to null]
**Size** | **int32** | Total usable space in bytes, not including WAFL reserve and aggregate Snapshot copy reserve. | [optional] [default to null]
**Used** | **int32** | Space used or reserved in bytes. Includes volume guarantees and aggregate metadata. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


