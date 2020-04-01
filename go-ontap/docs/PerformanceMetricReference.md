# PerformanceMetricReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Duration** | **string** | The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:  | [optional] [default to null]
**Iops** | [***ClusterMetricIops**](cluster_metric_iops.md) |  | [optional] [default to null]
**Latency** | [***ClusterMetricLatency**](cluster_metric_latency.md) |  | [optional] [default to null]
**Status** | **string** | Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, \&quot;ok\&quot; on success, or \&quot;error\&quot; on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with \&quot;backfilled_data\&quot;. \&quot;Inconsistent_ delta_time\&quot; is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. \&quot;Negative_delta\&quot; is returned when an expected monotonically increasing value has decreased in value. \&quot;Inconsistent_old_data\&quot; is returned when one or more nodes does not have the latest data. | [optional] [default to null]
**Throughput** | [***ClusterMetricThroughput**](cluster_metric_throughput.md) |  | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The timestamp of the performance data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


