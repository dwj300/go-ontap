# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Comment** | **string** | A comment associated with the Snapshot copy. This is an optional attribute for POST or PATCH. | [optional] 
**CreateTime** | [**time.Time**](time.Time.md) | Creation time of the Snapshot copy. It is the volume access time when the Snapshot copy was created. | [optional] [readonly] 
**ExpiryTime** | [**time.Time**](time.Time.md) | The expiry time for the Snapshot copy. This is an optional attribute for POST or PATCH. Snapshot copies with an expiry time set are not allowed to be deleted until the retention time is reached. | [optional] 
**Name** | **string** | Snapshot copy. Valid in POST or PATCH. | [optional] 
**SnaplockExpiryTime** | [**time.Time**](time.Time.md) | SnapLock expiry time for the Snapshot copy, if the Snapshot copy is taken on a SnapLock volume. A Snapshot copy is not allowed to be deleted or renamed until the SnapLock ComplianceClock time goes beyond this retention time. | [optional] [readonly] 
**State** | **string** | State of the Snapshot copy. There are cases where some Snapshot copies are not complete. In the \&quot;partial\&quot; state, the Snapshot copy is consistent but exists only on the subset of the constituents that existed prior to the FlexGroup&#39;s expansion. Partial Snapshot copies cannot be used for a Snapshot copy restore operation. A Snapshot copy is in an \&quot;invalid\&quot; state when it is present in some FlexGroup constituents but not in others. At all other times, a Snapshot copy is valid. | [optional] [readonly] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Uuid** | **string** | The UUID of the Snapshot copy in the volume that uniquely identifies the Snapshot copy in that volume. | [optional] [readonly] 
**Volume** | [**CifsShareVolume**](cifs_share_volume.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


