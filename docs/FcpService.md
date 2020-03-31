# FcpService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Enabled** | **bool** | The administrative state of the FC Protocol service. The FC Protocol service can be disabled to block all FC Protocol connectivity to the SVM.&lt;br/&gt; This is optional in POST and PATCH. The default setting is _true_ (enabled) in POST.  | [optional] [default to true]
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Target** | [**FcpServiceTarget**](fcp_service_target.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


