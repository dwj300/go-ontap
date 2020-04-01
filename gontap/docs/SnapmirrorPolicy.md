# SnapmirrorPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Comment** | **string** | Comment associated with the policy. | [optional] [default to null]
**IdentityPreservation** | **string** | Specifies which configuration of the source SVM is replicated to destination. This field is applicable only for SVM data protection and async policies. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NetworkCompressionEnabled** | **bool** | Specifies whether network compression is enabled for transfers. This is applicable only to async policies. | [optional] [default to null]
**Retention** | [**[]SnapmirrorPolicyRule**](snapmirror_policy_rule.md) | Policy on Snapshot copy retention. This is applicable only to async policies. | [optional] [default to null]
**Scope** | **string** | Set to \&quot;svm\&quot; for policies owned by an SVM, otherwise set to \&quot;cluster\&quot;. | [optional] [default to null]
**Svm** | [***FlexcacheRelationshipSvm**](flexcache_relationship_svm.md) |  | [optional] [default to null]
**SyncCommonSnapshotSchedule** | [***SnapmirrorPolicySyncCommonSnapshotSchedule**](snapmirror_policy_sync_common_snapshot_schedule.md) |  | [optional] [default to null]
**SyncType** | **string** |  | [optional] [default to null]
**Throttle** | **int32** | Throttle in KB/s. Default to unlimited. | [optional] [default to null]
**TransferSchedule** | [***SnapmirrorPolicyTransferSchedule**](snapmirror_policy_transfer_schedule.md) |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Uuid** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


