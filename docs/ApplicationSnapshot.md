# ApplicationSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Application** | [**ApplicationComponentApplication**](application_component_application.md) |  | [optional] 
**Comment** | **string** | Comment. Valid in POST. | [optional] 
**Components** | [**[]ApplicationSnapshotComponents**](application_snapshot_components.md) |  | [optional] [readonly] 
**ConsistencyType** | **string** | Consistency type. This is for categorization purposes only. A Snapshot copy should not be set to &#39;application consistent&#39; unless the host application is quiesced for the Snapshot copy. Valid in POST. | [optional] 
**CreateTime** | **string** | Creation time | [optional] [readonly] 
**IsPartial** | **bool** | A partial Snapshot copy means that not all volumes in an application component were included in the Snapshot copy. | [optional] [readonly] 
**Name** | **string** | The Snapshot copy name. Valid in POST. | [optional] 
**Svm** | [**ApplicationComponentSvm**](application_component_svm.md) |  | [optional] 
**Uuid** | **string** | The Snapshot copy UUID. Valid in URL. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


