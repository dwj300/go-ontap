# Qtree

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**ExportPolicy** | [***QtreeExportPolicy**](qtree_export_policy.md) |  | [optional] [default to null]
**Id** | **int32** | The identifier for the qtree, unique within the qtree&#39;s volume.  | [optional] [default to null]
**Name** | **string** | The name of the qtree. Required in POST; optional in PATCH. | [optional] [default to null]
**Path** | **string** | Client visible path to the qtree. This field is not available if the volume does not have a junction-path configured. Not valid in POST or PATCH. | [optional] [default to null]
**SecurityStyle** | [***SecurityStyle**](security_style.md) |  | [optional] [default to null]
**Svm** | [***QtreeSvm**](qtree_svm.md) |  | [optional] [default to null]
**UnixPermissions** | **int32** | The UNIX permissions for the qtree. Valid in POST or PATCH. | [optional] [default to null]
**Volume** | [***QtreeVolume**](qtree_volume.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


