# VolumeQosPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**MaxThroughputIops** | **int32** | Specifies the maximum throughput in IOPS, 0 means none. This is mutually exclusive with name and UUID during POST and PATCH. Default is 15000 on AFF platforms and 10000 on all other platforms. | [optional] 
**MaxThroughputMbps** | **int32** | Specifies the maximum throughput in Megabytes per sec, 0 means none. This is mutually exclusive with name, UUID and \&quot;min_throughput_iops\&quot; during POST and PATCH. | [optional] 
**MinThroughputIops** | **int32** | Specifies the minimum throughput in IOPS, 0 means none. Setting \&quot;min_throughput\&quot; is supported on AFF platforms only, unless FabricPool tiering policies are set. This is mutually exclusive with name, UUID and\&quot; max_throughput_mbps\&quot; during POST and PATCH. | [optional] 
**Name** | **string** | The QoS policy group name. This is mutually exclusive with UUID and other QoS attributes during POST and PATCH. | [optional] 
**Uuid** | **string** | The QoS policy group UUID. This is mutually exclusive with name and other QoS attributes during POST and PATCH. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


