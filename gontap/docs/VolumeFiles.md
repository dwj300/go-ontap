# VolumeFiles

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maximum** | **int32** | The maximum number of files (inodes) for user-visible data allowed on the volume. This value can be increased or decreased. Increasing the maximum number of files does not immediately cause additional disk space to be used to track files. Instead, as more files are created on the volume, the system dynamically increases the number of disk blocks that are used to track files. The space assigned to track files is never freed, and this value cannot be decreased below the current number of files that can be tracked within the assigned space for the volume. Valid in PATCH. | [optional] [default to null]
**Used** | **int32** | Number of files (inodes) used for user-visible data permitted on the volume. This field is valid only when the volume is online. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


