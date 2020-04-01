# MaxdataOnSanApplicationComponents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application component. Required in the POST body and optional in the PATCH body | [optional] [default to null]
**FileSystem** | **string** | Defines the kind of file system that will be installed on this application component. Optional in the POST body | [optional] [default to null]
**HostManagementUrl** | **string** | The host management URL for this application component | [optional] [default to null]
**HostName** | **string** | FQDN of the L2 host that contains the hot tier of this application component. Required in the POST body | [optional] [default to null]
**IgroupName** | **string** | The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup. Required in the POST body and optional in the PATCH body | [optional] [default to null]
**LunCount** | **int32** | The number of LUNs in the application component. Required in the POST body | [optional] [default to null]
**Metadata** | [**[]MaxdataOnSanApplicationComponentsMetadata**](maxdata_on_san_application_components_metadata.md) |  | [optional] [default to null]
**ProtectionType** | [***MaxdataOnSanApplicationComponentsProtectionType**](maxdata_on_san_application_components_protection_type.md) |  | [optional] [default to null]
**StorageService** | [***MaxdataOnSanApplicationComponentsStorageService**](maxdata_on_san_application_components_storage_service.md) |  | [optional] [default to null]
**TotalSize** | **int32** | The total size of the application component, split across the member LUNs. Usage: {&amp;lt;integer&amp;gt;[KB|MB|GB|TB|PB]} Required in the POST body | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


