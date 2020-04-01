# SecurityKeyManagerExternal

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCertificate** | [***SecurityKeyManagerExternalClientCertificate**](security_key_manager_external_client_certificate.md) |  | [optional] [default to null]
**ServerCaCertificates** | [**[]SecurityKeyManagerExternalServerCaCertificates**](security_key_manager_external_server_ca_certificates.md) | The UUIDs of the server CA certificates already installed in the cluster or SVM. The array of certificates are common for all the keyservers per SVM. | [optional] [default to null]
**Servers** | [**[]KeyServerReadcreate**](key_server_readcreate.md) | The set of external key servers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


