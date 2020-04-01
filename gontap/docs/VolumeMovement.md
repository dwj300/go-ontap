# VolumeMovement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CutoverWindow** | **int32** | Time window in seconds for cutover. The allowed range is between 30 to 300 seconds. | [optional] [default to null]
**DestinationAggregate** | [***DiskAggregates**](disk_aggregates.md) |  | [optional] [default to null]
**PercentComplete** | **string** | Completion percentage | [optional] [default to null]
**State** | **string** | State of volume move operation. PATCH the state to \&quot;aborted\&quot; to abort the move operation. PATCH the state to \&quot;cutover\&quot; to trigger cutover. PATCH the state to \&quot;paused\&quot; to pause the volume move operation in progress. PATCH the state to \&quot;replicating\&quot; to resume the paused volume move operation. PATCH the state to \&quot;cutover-wait\&quot; to go into cutover manually. Change of state is only supported if volume movement is in progress. | [optional] [default to null]
**TieringPolicy** | **string** | Tiering policy for FabricPool | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


