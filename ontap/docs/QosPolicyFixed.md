# QosPolicyFixed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityShared** | **bool** | Specifies whether the capacities are shared across all objects that use this QoS policy-group. Default is false. | [optional] [default to false]
**MaxThroughputIops** | **int32** | Maximum throughput defined by this policy.  It is specified in terms of IOPS. 0 means no maximum throughput is enforced. | [optional] 
**MaxThroughputMbps** | **int32** | Maximum throughput defined by this policy.  It is specified in terms of Mbps. 0 means no maximum throughput is enforced. | [optional] 
**MinThroughputIops** | **int32** | Minimum throughput defined by this policy.  It is specified in terms of IOPS. 0 means no minimum throughput is enforced. These floors are not guaranteed on non-AFF platforms or when FabricPool tiering policies are set. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


