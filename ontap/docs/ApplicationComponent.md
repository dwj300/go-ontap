# ApplicationComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Application** | [**ApplicationComponentApplication**](application_component_application.md) |  | [optional] 
**BackingStorage** | [**ApplicationBackingStorage**](application_backing_storage.md) |  | [optional] 
**CifsAccess** | [**[]ApplicationCifsProperties**](application_cifs_properties.md) |  | [optional] 
**FileSystem** | **string** | Defines the type of file system that will be installed on this application component. | [optional] [readonly] [default to FILE_SYSTEM_M1FS]
**HostManagementUrl** | **string** | Host management URL | [optional] [readonly] 
**HostName** | **string** | L2 Host FQDN | [optional] [readonly] 
**Name** | **string** | Application component name | [optional] [readonly] 
**NfsAccess** | [**[]ApplicationNfsProperties**](application_nfs_properties.md) |  | [optional] 
**ProtectionGroups** | [**[]ApplicationProtectionGroups**](application_protection_groups.md) |  | [optional] 
**SanAccess** | [**[]ApplicationSanAccess**](application_san_access.md) |  | [optional] 
**StorageService** | [**ApplicationComponentStorageService**](application_component_storage_service.md) |  | [optional] 
**Svm** | [**ApplicationComponentSvm**](application_component_svm.md) |  | [optional] 
**Uuid** | **string** | The application component UUID. Valid in URL. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


