# ClusterPeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ClusterPeerLinks**](cluster_peer__links.md) |  | [optional] [default to null]
**Authentication** | [***ClusterPeerAuthentication**](cluster_peer_authentication.md) |  | [optional] [default to null]
**Encryption** | [***ClusterPeerEncryption**](cluster_peer_encryption.md) |  | [optional] [default to null]
**InitialAllowedSvms** | [**[]AuditSvm**](audit_svm.md) | The local SVMs allowed to peer with the peer cluster&#39;s SVMs. This list can be modified until the remote cluster accepts this cluster peering relationship. | [optional] [default to null]
**Ipspace** | [***ClusterPeerIpspace**](cluster_peer_ipspace.md) |  | [optional] [default to null]
**LocalNetwork** | [***ClusterPeerLocalNetwork**](cluster_peer_local_network.md) |  | [optional] [default to null]
**Name** | **string** | Optional name for the cluster peer relationship. By default it is the name of the remote cluster. | [optional] [default to null]
**Remote** | [***ClusterPeerRemote**](cluster_peer_remote.md) |  | [optional] [default to null]
**Status** | [***ClusterPeerStatus**](cluster_peer_status.md) |  | [optional] [default to null]
**Uuid** | **string** | UUID of the cluster peer relationship. For anonymous cluster peer offers, the UUID will change when the remote cluster accepts the relationship. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


