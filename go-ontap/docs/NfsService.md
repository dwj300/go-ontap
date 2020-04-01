# NfsService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Enabled** | **bool** | Specifies if the NFS service is administratively enabled.  | [optional] [default to null]
**Protocol** | [***NfsServiceProtocol**](nfs_service_protocol.md) |  | [optional] [default to null]
**State** | **string** | Specifies the state of the NFS service on the SVM. The following values are supported:           * online - NFS server is ready to accept client requests.           * offline - NFS server is not ready to accept client requests.  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**Transport** | [***NfsServiceTransport**](nfs_service_transport.md) |  | [optional] [default to null]
**VstorageEnabled** | **bool** | Specifies whether VMware vstorage feature is enabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


