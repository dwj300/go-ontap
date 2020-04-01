# LunMovementProgress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | **int32** | The amount of time, in seconds, that has elapsed since the start of the LUN movement.  | [optional] [default to null]
**Failure** | [***LunMovementProgressFailure**](lun_movement_progress_failure.md) |  | [optional] [default to null]
**PercentComplete** | **int32** | The percentage complete of the LUN movement.  | [optional] [default to null]
**State** | **string** | The state of the LUN movement.&lt;br/&gt; Valid in PATCH when an LUN movement is active. Set to _paused_ to pause a LUN movement. Set to _replicating_ to resume a paused LUN movement.  | [optional] [default to null]
**VolumeSnapshotBlocked** | **bool** | This property reports if volume Snapshot copies are blocked by the LUN movement. This property can be polled to identify when volume Snapshot copies can be resumed after beginning a LUN movement.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


