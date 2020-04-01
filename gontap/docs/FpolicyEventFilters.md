# FpolicyEventFilters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseWithModification** | **bool** | Filter the client request for close with modification. | [optional] [default to null]
**CloseWithRead** | **bool** | Filter the client request for close with read. | [optional] [default to null]
**CloseWithoutModification** | **bool** | Filter the client request for close without modification. | [optional] [default to null]
**ExcludeDirectory** | **bool** | Filter the client requests for directory operations. When this filter is specified directory operations are not monitored. | [optional] [default to null]
**FirstRead** | **bool** | Filter the client requests for the first-read. | [optional] [default to null]
**FirstWrite** | **bool** | Filter the client requests for the first-write. | [optional] [default to null]
**MonitorAds** | **bool** | Filter the client request for alternate data stream. | [optional] [default to null]
**OfflineBit** | **bool** | Filter the client request for offline bit set. FPolicy server receives notification only when offline files are accessed. | [optional] [default to null]
**OpenWithDeleteIntent** | **bool** | Filter the client request for open with delete intent. | [optional] [default to null]
**OpenWithWriteIntent** | **bool** | Filter the client request for open with write intent. | [optional] [default to null]
**SetattrWithAccessTimeChange** | **bool** | Filter the client setattr requests for changing the access time of a file or directory. | [optional] [default to null]
**SetattrWithAllocationSizeChange** | **bool** | Filter the client setattr requests for changing the allocation size of a file. | [optional] [default to null]
**SetattrWithCreationTimeChange** | **bool** | Filter the client setattr requests for changing the creation time of a file or directory. | [optional] [default to null]
**SetattrWithDaclChange** | **bool** | Filter the client setattr requests for changing dacl on a file or directory. | [optional] [default to null]
**SetattrWithGroupChange** | **bool** | Filter the client setattr requests for changing group of a file or directory. | [optional] [default to null]
**SetattrWithModeChange** | **bool** | Filter the client setattr requests for changing the mode bits on a file or directory. | [optional] [default to null]
**SetattrWithModifyTimeChange** | **bool** | Filter the client setattr requests for changing the modification time of a file or directory. | [optional] [default to null]
**SetattrWithOwnerChange** | **bool** | Filter the client setattr requests for changing owner of a file or directory. | [optional] [default to null]
**SetattrWithSaclChange** | **bool** | Filter the client setattr requests for changing sacl on a file or directory. | [optional] [default to null]
**SetattrWithSizeChange** | **bool** | Filter the client setattr requests for changing the size of a file. | [optional] [default to null]
**WriteWithSizeChange** | **bool** | Filter the client request for write with size change. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


