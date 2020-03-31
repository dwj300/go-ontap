# FpolicyEventFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseWithModification** | **bool** | Filter the client request for close with modification. | [optional] 
**CloseWithRead** | **bool** | Filter the client request for close with read. | [optional] 
**CloseWithoutModification** | **bool** | Filter the client request for close without modification. | [optional] 
**ExcludeDirectory** | **bool** | Filter the client requests for directory operations. When this filter is specified directory operations are not monitored. | [optional] 
**FirstRead** | **bool** | Filter the client requests for the first-read. | [optional] 
**FirstWrite** | **bool** | Filter the client requests for the first-write. | [optional] 
**MonitorAds** | **bool** | Filter the client request for alternate data stream. | [optional] 
**OfflineBit** | **bool** | Filter the client request for offline bit set. FPolicy server receives notification only when offline files are accessed. | [optional] 
**OpenWithDeleteIntent** | **bool** | Filter the client request for open with delete intent. | [optional] 
**OpenWithWriteIntent** | **bool** | Filter the client request for open with write intent. | [optional] 
**SetattrWithAccessTimeChange** | **bool** | Filter the client setattr requests for changing the access time of a file or directory. | [optional] 
**SetattrWithAllocationSizeChange** | **bool** | Filter the client setattr requests for changing the allocation size of a file. | [optional] 
**SetattrWithCreationTimeChange** | **bool** | Filter the client setattr requests for changing the creation time of a file or directory. | [optional] 
**SetattrWithDaclChange** | **bool** | Filter the client setattr requests for changing dacl on a file or directory. | [optional] 
**SetattrWithGroupChange** | **bool** | Filter the client setattr requests for changing group of a file or directory. | [optional] 
**SetattrWithModeChange** | **bool** | Filter the client setattr requests for changing the mode bits on a file or directory. | [optional] 
**SetattrWithModifyTimeChange** | **bool** | Filter the client setattr requests for changing the modification time of a file or directory. | [optional] 
**SetattrWithOwnerChange** | **bool** | Filter the client setattr requests for changing owner of a file or directory. | [optional] 
**SetattrWithSaclChange** | **bool** | Filter the client setattr requests for changing sacl on a file or directory. | [optional] 
**SetattrWithSizeChange** | **bool** | Filter the client setattr requests for changing the size of a file. | [optional] 
**WriteWithSizeChange** | **bool** | Filter the client request for write with size change. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


