# SvmPeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Applications** | [**[]SvmPeerApplications**](svm_peer_applications.md) | A list of applications for an SVM peer relation. | [optional] [default to null]
**Name** | **string** | A peer SVM alias name to avoid a name conflict on the local cluster. | [optional] [default to null]
**Peer** | [***SvmPeerPeer**](svm_peer_peer.md) |  | [optional] [default to null]
**State** | **string** | SVM peering state. To accept a pending SVM peer request, PATCH the state to \&quot;peered\&quot;. To reject a pending SVM peer request, PATCH the state to \&quot;rejected\&quot;. To suspend a peered SVM peer relation, PATCH the state to \&quot;suspended\&quot;. To resume a suspended SVM peer relation, PATCH the state to \&quot;peered\&quot;. The states \&quot;initiated\&quot;, \&quot;pending\&quot;, and \&quot;initializing\&quot; are system-generated and cannot be used for PATCH. | [optional] [default to null]
**Svm** | [***SvmPeerSvm**](svm_peer_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | SVM peer relation UUID | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


