# SvmPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Applications** | [**[]SvmPeerApplications**](svm_peer_applications.md) | A list of applications for an SVM peer relation. | [optional] 
**Name** | **string** | A peer SVM alias name to avoid a name conflict on the local cluster. | [optional] 
**Peer** | [**SvmPeerPeer**](svm_peer_peer.md) |  | [optional] 
**State** | **string** | SVM peering state. To accept a pending SVM peer request, PATCH the state to \&quot;peered\&quot;. To reject a pending SVM peer request, PATCH the state to \&quot;rejected\&quot;. To suspend a peered SVM peer relation, PATCH the state to \&quot;suspended\&quot;. To resume a suspended SVM peer relation, PATCH the state to \&quot;peered\&quot;. The states \&quot;initiated\&quot;, \&quot;pending\&quot;, and \&quot;initializing\&quot; are system-generated and cannot be used for PATCH. | [optional] 
**Svm** | [**SvmPeerSvm**](svm_peer_svm.md) |  | [optional] 
**Uuid** | **string** | SVM peer relation UUID | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


