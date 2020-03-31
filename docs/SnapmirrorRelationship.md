# SnapmirrorRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**SnapmirrorEndpoint**](snapmirror_endpoint.md) |  | [optional] 
**ExportedSnapshot** | **string** | Snapshot copy exported to clients on destination. | [optional] [readonly] 
**Healthy** | **bool** | Is the relationship healthy? | [optional] [readonly] 
**LagTime** | **string** | Time since the exported Snapshot copy was created. | [optional] [readonly] 
**Policy** | [**SnapmirrorRelationshipPolicy**](snapmirror_relationship_policy.md) |  | [optional] 
**Preserve** | **bool** | Set to true on resync to preserve Snapshot copies on the destination that are newer than the latest common Snapshot copy. This field is applicable only for relationships with FlexGroup or FlexVol endpoints and when the PATCH state is \&quot;snapmirrored\&quot;. | [optional] [default to false]
**QuickResync** | **bool** | Set to true to reduce resync time by not preserving storage efficiency. This field is applicable only for relationships with FlexVol endpoints and when the PATCH state is \&quot;snapmirrored\&quot;. | [optional] [default to false]
**RecoverAfterBreak** | **bool** | Set to true to recover from a failed SnapMirror break operation on a FlexGroup relationship. This restores all destination FlexGroup constituents to the latest Snapshot copy, and any writes to the read-write constituents are lost. This field is applicable only for SnapMirror relationships with FlexGroup endpoints and when the PATCH state is \&quot;broken_off\&quot;. | [optional] [default to false]
**Restore** | **bool** | Set to true to create a relationship for restore. To trigger restore-transfer, use transfers POST on the restore relationship. | [optional] [default to false]
**RestoreToSnapshot** | **string** | Specifies the Snapshot copy to restore to on the destination after a break operation. This field is applicable only for SnapMirror relationships with FlexVol endpoints and when the PATCH state is \&quot;broken_off\&quot;. | [optional] 
**Source** | [**SnapmirrorEndpoint**](snapmirror_endpoint.md) |  | [optional] 
**State** | **string** | State of the relationship. To initialize the relationship, PATCH the state to \&quot;snapmirrored\&quot;. To break the relationship, PATCH the state to \&quot;broken_off\&quot;. To resync the broken relationship, PATCH the state to \&quot;snapmirrored\&quot; for relationships with async policy type or \&quot;in_sync\&quot; for relationships with sync policy type. To pause the relationship, suspending further transfers, PATCH the state to \&quot;paused\&quot;. To resume transfers for a paused relationship, PATCH the state to \&quot;snapmirrored\&quot; or \&quot;in_sync\&quot;. The entries \&quot;in_sync\&quot;, \&quot;out_of_sync\&quot;, and \&quot;synchronizing\&quot; are only applicable to relationships of the sync policy type. A PATCH call on the state change only triggers the transition to the specified state. You must poll on the \&quot;state\&quot;, \&quot;healthy\&quot; and \&quot;unhealthy_reason\&quot; fields using GET to determine if the transition is successful. | [optional] 
**Transfer** | [**SnapmirrorRelationshipTransfer**](snapmirror_relationship_transfer.md) |  | [optional] 
**UnhealthyReason** | [**[]SnapmirrorError**](snapmirror_error.md) | Reason the relationship is not healthy. It is a concatenation of up to four levels of error messages. | [optional] [readonly] 
**Uuid** | **string** |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


