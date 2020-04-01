# Flexcache

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Aggregates** | [**[]DiskAggregates**](disk_aggregates.md) |  | [optional] [default to null]
**ConstituentsPerAggregate** | **int32** | Number of FlexCache constituents per aggregate when the &#39;aggregates&#39; field is mentioned. | [optional] [default to null]
**Name** | **string** | FlexCache name | [optional] [default to null]
**Origins** | [**[]FlexcacheRelationship**](flexcache_relationship.md) |  | [optional] [default to null]
**Path** | **string** | The fully-qualified path in the owning SVM&#39;s namespace at which the FlexCache is mounted. The path is case insensitive and must be unique within a SVM&#39;s namespace. Path must begin with &#39;/&#39; and must not end with &#39;/&#39;. Only one FlexCache be mounted at any given junction path. | [optional] [default to null]
**Size** | **int32** | Physical size of the FlexCache. The recommended size for a FlexCache is 10% of the origin volume. The minimum FlexCache constituent size is 1GB. | [optional] [default to null]
**Svm** | [***FlexcacheSvm**](flexcache_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | FlexCache UUID. Unique identifier for the FlexCache. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


