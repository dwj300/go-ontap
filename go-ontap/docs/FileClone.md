# FileClone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autodelete** | **bool** | Mark file clone for auto deletion. | [optional] [default to null]
**DestinationPath** | **string** | Relative path of the clone/destination file in the volume. | [optional] [default to null]
**Range_** | **[]string** | List of block ranges for sub-file cloning in the format \&quot;source-file-block-number:destination-file-block-number:block-count\&quot; | [optional] [default to null]
**SourcePath** | **string** | Relative path of the source file in the volume. | [optional] [default to null]
**Volume** | [***CifsShareVolume**](cifs_share_volume.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


