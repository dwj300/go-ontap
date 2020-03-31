# VdiOnSan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desktops** | [**VdiOnNasDesktops**](vdi_on_nas_desktops.md) |  | [optional] 
**Hypervisor** | **string** | The name of the hypervisor hosting the application. Required in the POST body | [optional] 
**IgroupName** | **string** | The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup. Required in the POST body and optional in the PATCH body | [optional] 
**NewIgroups** | [**[]VdiOnSanNewIgroups**](vdi_on_san_new_igroups.md) | The list of initiator groups to create. Optional in the POST or PATCH body | [optional] 
**ProtectionType** | [**MongoDbOnSanProtectionType**](mongo_db_on_san_protection_type.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


