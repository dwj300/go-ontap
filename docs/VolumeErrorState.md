# VolumeErrorState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasBadBlocks** | **bool** | Indicates whether the volume has any corrupt data blocks. If the damaged data block is accessed, an IO error, such as EIO for NFS or STATUS_FILE_CORRUPT for CIFS, is returned. | [optional] [readonly] 
**IsInconsistent** | **bool** | Indicates whether the file system has any inconsistencies.&lt;br&gt;true &amp;dash; File system is inconsistent.&lt;br&gt;false &amp;dash; File system in not inconsistent. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


