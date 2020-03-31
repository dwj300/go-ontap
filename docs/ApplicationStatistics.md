# ApplicationStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | [**[]ApplicationStatisticsComponents**](application_statistics_components.md) |  | [optional] [readonly] 
**Iops** | [**ApplicationStatisticsIops1**](application_statistics_iops_1.md) |  | [optional] 
**Latency** | [**ApplicationStatisticsLatency1**](application_statistics_latency_1.md) |  | [optional] 
**SharedStoragePool** | **bool** | An application is considered to use a shared storage pool if storage elements for multiple components reside on the same aggregate | [optional] [readonly] 
**Snapshot** | [**ApplicationStatisticsSnapshot**](application_statistics_snapshot.md) |  | [optional] 
**Space** | [**ApplicationStatisticsSpace1**](application_statistics_space_1.md) |  | [optional] 
**StatisticsIncomplete** | **bool** | If not all storage elements of the application are currently available, the returned statistics might only include data from those elements that were available | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


