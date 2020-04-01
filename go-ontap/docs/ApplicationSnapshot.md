# ApplicationSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Application** | [***ApplicationComponentApplication**](application_component_application.md) |  | [optional] [default to null]
**Comment** | **string** | Comment. Valid in POST. | [optional] [default to null]
**Components** | [**[]ApplicationSnapshotComponents**](application_snapshot_components.md) |  | [optional] [default to null]
**ConsistencyType** | **string** | Consistency type. This is for categorization purposes only. A Snapshot copy should not be set to &#39;application consistent&#39; unless the host application is quiesced for the Snapshot copy. Valid in POST. | [optional] [default to null]
**CreateTime** | **string** | Creation time | [optional] [default to null]
**IsPartial** | **bool** | A partial Snapshot copy means that not all volumes in an application component were included in the Snapshot copy. | [optional] [default to null]
**Name** | **string** | The Snapshot copy name. Valid in POST. | [optional] [default to null]
**Svm** | [***ApplicationComponentSvm**](application_component_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | The Snapshot copy UUID. Valid in URL. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


