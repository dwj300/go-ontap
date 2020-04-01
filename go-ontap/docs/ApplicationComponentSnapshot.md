# ApplicationComponentSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Application** | [***ApplicationComponentSnapshotApplication**](application_component_snapshot_application.md) |  | [optional] [default to null]
**Comment** | **string** | Comment. Valid in POST | [optional] [default to null]
**Component** | [***ApplicationComponentSnapshotComponent**](application_component_snapshot_component.md) |  | [optional] [default to null]
**ConsistencyType** | **string** | Consistency Type. This is for categorization only. A snapshot should not be set to application consistent unless the host application is quiesced for the snapshot. Valid in POST | [optional] [default to null]
**CreateTime** | **string** | Creation Time | [optional] [default to null]
**IsPartial** | **bool** | A partial snapshot means that not all volumes in an application component were included in the snapshot. | [optional] [default to null]
**Name** | **string** | Snapshot Name. Valid in POST | [optional] [default to null]
**Svm** | [***ApplicationComponentSnapshotSvm**](application_component_snapshot_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | Snapshot UUID. Valid in URL | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


