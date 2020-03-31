# PerformanceMetricRawReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IopsRaw** | [**ClusterStatisticsIopsRaw**](cluster_statistics_iops_raw.md) |  | [optional] 
**LatencyRaw** | [**ClusterStatisticsLatencyRaw**](cluster_statistics_latency_raw.md) |  | [optional] 
**Status** | **string** | Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, \&quot;ok\&quot; on success, or \&quot;error\&quot; on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with \&quot;backfilled_data\&quot;. \&quot;Inconsistent_delta_time\&quot; is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. \&quot;Negative_delta\&quot; is returned when an expected monotonically increasing value has decreased in value. \&quot;Inconsistent_old_data\&quot; is returned when one or more nodes does not have the latest data. | [optional] [readonly] 
**ThroughputRaw** | [**ClusterStatisticsThroughputRaw**](cluster_statistics_throughput_raw.md) |  | [optional] 
**Timestamp** | [**time.Time**](time.Time.md) | The timestamp of the performance data. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


