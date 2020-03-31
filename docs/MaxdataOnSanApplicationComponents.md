# MaxdataOnSanApplicationComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application component. Required in the POST body and optional in the PATCH body | [optional] 
**FileSystem** | **string** | Defines the kind of file system that will be installed on this application component. Optional in the POST body | [optional] [default to FILE_SYSTEM_M1FS]
**HostManagementUrl** | **string** | The host management URL for this application component | [optional] [readonly] 
**HostName** | **string** | FQDN of the L2 host that contains the hot tier of this application component. Required in the POST body | [optional] 
**IgroupName** | **string** | The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup. Required in the POST body and optional in the PATCH body | [optional] 
**LunCount** | **int32** | The number of LUNs in the application component. Required in the POST body | [optional] 
**Metadata** | [**[]MaxdataOnSanApplicationComponentsMetadata**](maxdata_on_san_application_components_metadata.md) |  | [optional] 
**ProtectionType** | [**MaxdataOnSanApplicationComponentsProtectionType**](maxdata_on_san_application_components_protection_type.md) |  | [optional] 
**StorageService** | [**MaxdataOnSanApplicationComponentsStorageService**](maxdata_on_san_application_components_storage_service.md) |  | [optional] 
**TotalSize** | **int32** | The total size of the application component, split across the member LUNs. Usage: {&amp;lt;integer&amp;gt;[KB|MB|GB|TB|PB]} Required in the POST body | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


