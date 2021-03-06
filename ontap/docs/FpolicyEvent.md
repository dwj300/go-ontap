# FpolicyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileOperations** | [**FpolicyEventFileOperations**](fpolicy_event_file_operations.md) |  | [optional] 
**Filters** | [**FpolicyEventFilters**](fpolicy_event_filters.md) |  | [optional] 
**Name** | **string** | Specifies the name of the FPolicy event. | [optional] 
**Protocol** | **string** | Protocol for which event is created. If you specify protocol, then you must also specify a valid value for the file operation parameters.   The value of this parameter must be one of the following:     * cifs  - for the CIFS protocol.     * nfsv3 - for the NFSv3 protocol.     * nfsv4 - for the NFSv4 protocol.  | [optional] 
**VolumeMonitoring** | **bool** | Specifies whether volume operation monitoring is required. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


