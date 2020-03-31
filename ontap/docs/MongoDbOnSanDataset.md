# MongoDbOnSanDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementCount** | **int32** | The number of storage elements (LUNs for SAN) of the database to maintain.  Must be an even number between 2 and 16.  Odd numbers will be rounded up to the next even number within range. Optional in the POST body | [optional] 
**ReplicationFactor** | **int32** | The number of data bearing members of the replicaset, including 1 primary and at least 1 secondary. Optional in the POST body | [optional] 
**Size** | **int32** | The size of the database. Usage: {&amp;lt;integer&amp;gt;[KB|MB|GB|TB|PB]} Required in the POST body and optional in the PATCH body | [optional] 
**StorageService** | [**MongoDbOnSanDatasetStorageService**](mongo_db_on_san_dataset_storage_service.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


