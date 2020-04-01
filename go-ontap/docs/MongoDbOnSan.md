# MongoDbOnSan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | [***MongoDbOnSanDataset**](mongo_db_on_san_dataset.md) |  | [optional] [default to null]
**NewIgroups** | [**[]MongoDbOnSanNewIgroups**](mongo_db_on_san_new_igroups.md) | The list of initiator groups to create. Optional in the POST or PATCH body | [optional] [default to null]
**OsType** | **string** | The name of the host OS running the application. Optional in the POST body | [optional] [default to null]
**PrimaryIgroupName** | **string** | The initiator group for the primary. Required in the POST body and optional in the PATCH body | [optional] [default to null]
**ProtectionType** | [***MongoDbOnSanProtectionType**](mongo_db_on_san_protection_type.md) |  | [optional] [default to null]
**SecondaryIgroups** | [**[]MongoDbOnSanSecondaryIgroups**](mongo_db_on_san_secondary_igroups.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


