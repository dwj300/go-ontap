# IscsiService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Enabled** | **bool** | The administrative state of the iSCSI service. The iSCSI service can be disabled to block all iSCSI connectivity to the SVM.&lt;br/&gt; Optional in POST and PATCH. The default setting is _true_ (enabled) in POST.  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**Target** | [***IscsiServiceTarget**](iscsi_service_target.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


