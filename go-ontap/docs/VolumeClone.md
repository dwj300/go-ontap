# VolumeClone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFlexclone** | **bool** | Specifies if this volume is a normal FlexVol or FlexClone. This field needs to be set when creating a FlexClone. Valid in POST. | [optional] [default to null]
**ParentSnapshot** | [***SnapshotReference**](snapshot_reference.md) |  | [optional] [default to null]
**ParentSvm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**ParentVolume** | [***CifsShareVolume**](cifs_share_volume.md) |  | [optional] [default to null]
**SplitCompletePercent** | **int32** | Percentage of FlexClone blocks split from its parent volume. | [optional] [default to null]
**SplitEstimate** | **int32** | Space required by the containing-aggregate to split the FlexClone volume. | [optional] [default to null]
**SplitInitiated** | **bool** | This field is set when split is executed on any FlexClone, that is when the FlexClone volume is split from its parent FlexVol. This field needs to be set for splitting a FlexClone form FlexVol. Valid in PATCH. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


