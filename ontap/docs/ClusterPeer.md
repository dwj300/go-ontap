# ClusterPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**ClusterPeerLinks**](cluster_peer__links.md) |  | [optional] 
**Authentication** | [**ClusterPeerAuthentication**](cluster_peer_authentication.md) |  | [optional] 
**Encryption** | [**ClusterPeerEncryption**](cluster_peer_encryption.md) |  | [optional] 
**InitialAllowedSvms** | [**[]AuditSvm**](audit_svm.md) | The local SVMs allowed to peer with the peer cluster&#39;s SVMs. This list can be modified until the remote cluster accepts this cluster peering relationship. | [optional] 
**Ipspace** | [**ClusterPeerIpspace**](cluster_peer_ipspace.md) |  | [optional] 
**LocalNetwork** | [**ClusterPeerLocalNetwork**](cluster_peer_local_network.md) |  | [optional] 
**Name** | **string** | Optional name for the cluster peer relationship. By default it is the name of the remote cluster. | [optional] 
**Remote** | [**ClusterPeerRemote**](cluster_peer_remote.md) |  | [optional] 
**Status** | [**ClusterPeerStatus**](cluster_peer_status.md) |  | [optional] 
**Uuid** | **string** | UUID of the cluster peer relationship. For anonymous cluster peer offers, the UUID will change when the remote cluster accepts the relationship. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


