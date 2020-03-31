# SoftwareReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Action** | **string** | User triggered action to apply to the install operation | [optional] 
**ElapsedDuration** | **int32** | Elapsed time during the upgrade or validation operation | [optional] [readonly] 
**EstimatedDuration** | **int32** | Estimated time remaining until completion of the upgrade or validation operation. | [optional] [readonly] 
**Metrocluster** | [**SoftwareReferenceMetrocluster**](software_reference_metrocluster.md) |  | [optional] 
**Nodes** | [**[]SoftwareNodeReference**](software_node_reference.md) | List of nodes and active versions. | [optional] [readonly] 
**PendingVersion** | **string** | Version being installed on the system. | [optional] [readonly] 
**State** | **string** | Operational state of the upgrade | [optional] [readonly] 
**StatusDetails** | [**[]SoftwareStatusDetailsReference**](software_status_details_reference.md) | Display status details. | [optional] [readonly] 
**UpdateDetails** | [**[]SoftwareUpdateDetailsReference**](software_update_details_reference.md) | Display update procress details. | [optional] [readonly] 
**ValidationResults** | [**[]SoftwareValidationReference**](software_validation_reference.md) | List of validation warnings, errors, and advice. | [optional] [readonly] 
**Version** | **string** | Version of ONTAP installed and currently active on the system. During PATCH, using the &#39;validate_only&#39; parameter on the request executes pre-checks, but does not perform the full installation. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


