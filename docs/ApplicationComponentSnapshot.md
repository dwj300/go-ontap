# ApplicationComponentSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Application** | [**ApplicationComponentSnapshotApplication**](application_component_snapshot_application.md) |  | [optional] 
**Comment** | **string** | Comment. Valid in POST | [optional] 
**Component** | [**ApplicationComponentSnapshotComponent**](application_component_snapshot_component.md) |  | [optional] 
**ConsistencyType** | **string** | Consistency Type. This is for categorization only. A snapshot should not be set to application consistent unless the host application is quiesced for the snapshot. Valid in POST | [optional] 
**CreateTime** | **string** | Creation Time | [optional] [readonly] 
**IsPartial** | **bool** | A partial snapshot means that not all volumes in an application component were included in the snapshot. | [optional] [readonly] 
**Name** | **string** | Snapshot Name. Valid in POST | [optional] 
**Svm** | [**ApplicationComponentSnapshotSvm**](application_component_snapshot_svm.md) |  | [optional] 
**Uuid** | **string** | Snapshot UUID. Valid in URL | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


