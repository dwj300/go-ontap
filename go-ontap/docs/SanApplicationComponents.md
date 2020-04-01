# SanApplicationComponents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application component. Optional in the POST or PATCH body | [optional] [default to null]
**IgroupName** | **string** | The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup. Optional in the POST or PATCH body | [optional] [default to null]
**LunCount** | **int32** | The number of LUNs in the application component. Optional in the POST body | [optional] [default to null]
**StorageService** | [***NasStorageService**](nas_storage_service.md) |  | [optional] [default to null]
**TotalSize** | **int32** | The total size of the application component, split across the member LUNs. Usage: {&amp;lt;integer&amp;gt;[KB|MB|GB|TB|PB]} Optional in the POST or PATCH body | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


