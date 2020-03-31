# ApplicationStatisticsComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Component UUID | [optional] [readonly] 
**Name** | **string** | Component Name | [optional] [readonly] 
**Iops** | [**ApplicationStatisticsIops**](application_statistics_iops.md) |  | [optional] 
**Latency** | [**ApplicationStatisticsLatency**](application_statistics_latency.md) |  | [optional] 
**SharedStoragePool** | **bool** | An application component is considered to use a shared storage pool if storage elements for for other components reside on the same aggregate as storage elements for this component | [optional] [readonly] 
**Snapshot** | [**ApplicationStatisticsSnapshot**](application_statistics_snapshot.md) |  | [optional] 
**Space** | [**ApplicationStatisticsSpace**](application_statistics_space.md) |  | [optional] 
**StatisticsIncomplete** | **bool** | If not all storage elements of the application component are currently available, the returned statistics might only include data from those elements that were available | [optional] [readonly] 
**StorageService** | [**ApplicationStatisticsStorageService**](application_statistics_storage_service.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


