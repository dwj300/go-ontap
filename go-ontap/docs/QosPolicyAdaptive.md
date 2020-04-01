# QosPolicyAdaptive

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteMinIops** | **int32** | Specifies the absolute minimum IOPS that is used as an override when the expected_iops is less than this value. These floors are not guaranteed on non-AFF platforms or when FabricPool tiering policies are set. | [optional] [default to null]
**ExpectedIops** | **int32** | Expected IOPS. Specifies the minimum expected IOPS per TB allocated based on the storage object allocated size. These floors are not guaranteed on non-AFF platforms or when FabricPool tiering policies are set. | [optional] [default to null]
**PeakIops** | **int32** | Peak IOPS. Specifies the maximum possible IOPS per TB allocated based on the storage object allocated size or the storage object used size. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


