# VolumeSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **int32** | The available space, in bytes. | [optional] [readonly] 
**BlockStorageInactiveUserData** | **int32** | The size that is physically used in the block storage of the volume and has a cold temperature. In bytes. This parameter is only supported if the volume is in an aggregate that is either attached to a cloud store or could be attached to a cloud store. | [optional] [readonly] 
**CapacityTierFootprint** | **int32** | The space used by capacity tier for this volume in the aggregate, in bytes. | [optional] [readonly] 
**Footprint** | **int32** | Data and metadata used for this volume in the aggregate, in bytes. | [optional] [readonly] 
**LogicalSpace** | [**VolumeSpaceLogicalSpace**](volume_space_logical_space.md) |  | [optional] 
**Metadata** | **int32** | The space used by the total metadata in the volume, in bytes. | [optional] [readonly] 
**OverProvisioned** | **int32** | The amount of space not available for this volume in the aggregate, in bytes. | [optional] [readonly] 
**Size** | **int32** | Total provisioned size. The default size is equal to the minimum size of 20MB, in bytes. | [optional] 
**Snapshot** | [**VolumeSpaceSnapshot**](volume_space_snapshot.md) |  | [optional] 
**Used** | **int32** | The virtual space used (includes volume reserves) before storage efficiency, in bytes. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


