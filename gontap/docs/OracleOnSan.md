# OracleOnSan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLog** | [***OracleOnNfsArchiveLog**](oracle_on_nfs_archive_log.md) |  | [optional] [default to null]
**Db** | [***OracleOnNfsDb**](oracle_on_nfs_db.md) |  | [optional] [default to null]
**IgroupName** | **string** | The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup. Required in the POST body and optional in the PATCH body | [optional] [default to null]
**NewIgroups** | [**[]OracleOnSanNewIgroups**](oracle_on_san_new_igroups.md) | The list of initiator groups to create. Optional in the POST or PATCH body | [optional] [default to null]
**OraHome** | [***OracleOnNfsOraHome**](oracle_on_nfs_ora_home.md) |  | [optional] [default to null]
**OsType** | **string** | The name of the host OS running the application. Required in the POST body | [optional] [default to null]
**ProtectionType** | [***MongoDbOnSanProtectionType**](mongo_db_on_san_protection_type.md) |  | [optional] [default to null]
**RedoLog** | [***OracleOnNfsRedoLog**](oracle_on_nfs_redo_log.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


