# Flexcache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Aggregates** | [**[]DiskAggregates**](disk_aggregates.md) |  | [optional] 
**ConstituentsPerAggregate** | **int32** | Number of FlexCache constituents per aggregate when the &#39;aggregates&#39; field is mentioned. | [optional] 
**Name** | **string** | FlexCache name | [optional] 
**Origins** | [**[]FlexcacheRelationship**](flexcache_relationship.md) |  | [optional] 
**Path** | **string** | The fully-qualified path in the owning SVM&#39;s namespace at which the FlexCache is mounted. The path is case insensitive and must be unique within a SVM&#39;s namespace. Path must begin with &#39;/&#39; and must not end with &#39;/&#39;. Only one FlexCache be mounted at any given junction path. | [optional] 
**Size** | **int32** | Physical size of the FlexCache. The recommended size for a FlexCache is 10% of the origin volume. The minimum FlexCache constituent size is 1GB. | [optional] 
**Svm** | [**FlexcacheSvm**](flexcache_svm.md) |  | [optional] 
**Uuid** | **string** | FlexCache UUID. Unique identifier for the FlexCache. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


