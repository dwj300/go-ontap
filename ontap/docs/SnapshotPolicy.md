# SnapshotPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Comment** | **string** | A comment associated with the Snapshot copy policy. | [optional] 
**Copies** | [**[]SnapshotPolicyCopies**](snapshot_policy_copies.md) |  | [optional] 
**Enabled** | **bool** | Is the Snapshot copy policy enabled? | [optional] 
**Name** | **string** | Name of the Snapshot copy policy. | [optional] 
**Scope** | **string** | Set to \&quot;svm\&quot; when the request is on a data SVM, otherwise set to \&quot;cluster\&quot;. | [optional] [readonly] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Uuid** | **string** |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


