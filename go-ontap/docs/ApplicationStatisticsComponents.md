# ApplicationStatisticsComponents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Component UUID | [optional] [default to null]
**Name** | **string** | Component Name | [optional] [default to null]
**Iops** | [***ApplicationStatisticsIops**](application_statistics_iops.md) |  | [optional] [default to null]
**Latency** | [***ApplicationStatisticsLatency**](application_statistics_latency.md) |  | [optional] [default to null]
**SharedStoragePool** | **bool** | An application component is considered to use a shared storage pool if storage elements for for other components reside on the same aggregate as storage elements for this component | [optional] [default to null]
**Snapshot** | [***ApplicationStatisticsSnapshot**](application_statistics_snapshot.md) |  | [optional] [default to null]
**Space** | [***ApplicationStatisticsSpace**](application_statistics_space.md) |  | [optional] [default to null]
**StatisticsIncomplete** | **bool** | If not all storage elements of the application component are currently available, the returned statistics might only include data from those elements that were available | [optional] [default to null]
**StorageService** | [***ApplicationStatisticsStorageService**](application_statistics_storage_service.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


