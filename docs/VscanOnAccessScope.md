# VscanOnAccessScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeExtensions** | **[]string** | List of file extensions for which scanning is not performed. | [optional] 
**ExcludePaths** | **[]string** | List of file paths for which scanning must not be performed. | [optional] 
**IncludeExtensions** | **[]string** | List of file extensions to be scanned. | [optional] 
**MaxFileSize** | **int32** | Maximum file size, in bytes, allowed for scanning. | [optional] 
**OnlyExecuteAccess** | **bool** | Scan only files opened with execute-access. | [optional] [default to false]
**ScanReadonlyVolumes** | **bool** | Specifies whether or not read-only volume can be scanned. | [optional] [default to false]
**ScanWithoutExtension** | **bool** | Specifies whether or not files without any extension can be scanned. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


