# SqlOnSan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | [**SqlOnSanDb**](sql_on_san_db.md) |  | [optional] 
**IgroupName** | **string** | The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup. Required in the POST body and optional in the PATCH body | [optional] 
**Log** | [**SqlOnSanLog**](sql_on_san_log.md) |  | [optional] 
**NewIgroups** | [**[]SqlOnSanNewIgroups**](sql_on_san_new_igroups.md) | The list of initiator groups to create. Optional in the POST or PATCH body | [optional] 
**OsType** | **string** | The name of the host OS running the application. Optional in the POST body | [optional] [default to OS_TYPE_WINDOWS_2008]
**ProtectionType** | [**MongoDbOnSanProtectionType**](mongo_db_on_san_protection_type.md) |  | [optional] 
**ServerCoresCount** | **int32** | The number of server cores for the db. Optional in the POST body | [optional] 
**TempDb** | [**SqlOnSanTempDb**](sql_on_san_temp_db.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


