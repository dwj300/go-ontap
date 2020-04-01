# SoftwareReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Action** | **string** | User triggered action to apply to the install operation | [optional] [default to null]
**ElapsedDuration** | **int32** | Elapsed time during the upgrade or validation operation | [optional] [default to null]
**EstimatedDuration** | **int32** | Estimated time remaining until completion of the upgrade or validation operation. | [optional] [default to null]
**Metrocluster** | [***SoftwareReferenceMetrocluster**](software_reference_metrocluster.md) |  | [optional] [default to null]
**Nodes** | [**[]SoftwareNodeReference**](software_node_reference.md) | List of nodes and active versions. | [optional] [default to null]
**PendingVersion** | **string** | Version being installed on the system. | [optional] [default to null]
**State** | **string** | Operational state of the upgrade | [optional] [default to null]
**StatusDetails** | [**[]SoftwareStatusDetailsReference**](software_status_details_reference.md) | Display status details. | [optional] [default to null]
**UpdateDetails** | [**[]SoftwareUpdateDetailsReference**](software_update_details_reference.md) | Display update procress details. | [optional] [default to null]
**ValidationResults** | [**[]SoftwareValidationReference**](software_validation_reference.md) | List of validation warnings, errors, and advice. | [optional] [default to null]
**Version** | **string** | Version of ONTAP installed and currently active on the system. During PATCH, using the &#39;validate_only&#39; parameter on the request executes pre-checks, but does not perform the full installation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


