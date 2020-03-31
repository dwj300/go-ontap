# SnapmirrorPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Comment** | **string** | Comment associated with the policy. | [optional] 
**IdentityPreservation** | **string** | Specifies which configuration of the source SVM is replicated to destination. This field is applicable only for SVM data protection and async policies. | [optional] 
**Name** | **string** |  | [optional] 
**NetworkCompressionEnabled** | **bool** | Specifies whether network compression is enabled for transfers. This is applicable only to async policies. | [optional] [default to false]
**Retention** | [**[]SnapmirrorPolicyRule**](snapmirror_policy_rule.md) | Policy on Snapshot copy retention. This is applicable only to async policies. | [optional] 
**Scope** | **string** | Set to \&quot;svm\&quot; for policies owned by an SVM, otherwise set to \&quot;cluster\&quot;. | [optional] [readonly] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**SyncCommonSnapshotSchedule** | [**SnapmirrorPolicySyncCommonSnapshotSchedule**](snapmirror_policy_sync_common_snapshot_schedule.md) |  | [optional] 
**SyncType** | **string** |  | [optional] 
**Throttle** | **int32** | Throttle in KB/s. Default to unlimited. | [optional] 
**TransferSchedule** | [**SnapmirrorPolicyTransferSchedule**](snapmirror_policy_transfer_schedule.md) |  | [optional] 
**Type** | **string** |  | [optional] [default to TYPE_ASYNC]
**Uuid** | **string** |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


