# QosPolicyReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**MaxThroughputIops** | **int32** | Specifies the maximum throughput in IOPS, 0 means none. This is mutually exclusive with name and UUID during POST and PATCH. Default is 15000 on AFF platforms and 10000 on all other platforms. | [optional] [default to null]
**MaxThroughputMbps** | **int32** | Specifies the maximum throughput in Megabytes per sec, 0 means none. This is mutually exclusive with name, UUID and \&quot;min_throughput_iops\&quot; during POST and PATCH. | [optional] [default to null]
**MinThroughputIops** | **int32** | Specifies the minimum throughput in IOPS, 0 means none. Setting \&quot;min_throughput\&quot; is supported on AFF platforms only, unless FabricPool tiering policies are set. This is mutually exclusive with name, UUID and\&quot; max_throughput_mbps\&quot; during POST and PATCH. | [optional] [default to null]
**Name** | **string** | The QoS policy group name. This is mutually exclusive with UUID and other QoS attributes during POST and PATCH. | [optional] [default to null]
**Uuid** | **string** | The QoS policy group UUID. This is mutually exclusive with name and other QoS attributes during POST and PATCH. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


