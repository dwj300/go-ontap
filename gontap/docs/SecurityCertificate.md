# SecurityCertificate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Ca** | **string** | Certificate authority | [optional] [default to null]
**CommonName** | **string** | FQDN or custom common name. Provide on POST when creating a self-signed certificate. | [optional] [default to null]
**ExpiryTime** | **string** | Certificate expiration time. Can be provided on POST if creating self-signed certificate. The expiration time range is between 1 day to 10 years. | [optional] [default to null]
**HashFunction** | **string** | Hashing function. Can be provided on POST when creating a self-signed certificate. Hash functions md5 and sha1 are not allowed on POST. | [optional] [default to null]
**IntermediateCertificates** | **[]string** | Chain of intermediate Certificates in PEM format. Only valid in POST when installing a certificate. | [optional] [default to null]
**KeySize** | **int32** | Key size of requested Certificate in bits. One of 512, 1024, 1536, 2048, 3072. Can be provided on POST if creating self-signed certificate. Key size of 512 is not allowed on POST. | [optional] [default to null]
**PrivateKey** | **string** | Private key Certificate in PEM format. Only valid for create when installing a CA-signed certificate. This is not audited. | [optional] [default to null]
**PublicCertificate** | **string** | Public key Certificate in PEM format. If this is not provided in POST, a self-signed certificate is created. | [optional] [default to null]
**Scope** | [***NetworkScopeReadonly**](network_scope_readonly.md) |  | [optional] [default to null]
**SerialNumber** | **string** | Serial number of certificate. | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**Type_** | **string** | Type of Certificate. The following types are supported: * client - a certificate and its private key used by an SSL client in ONTAP. * server - a certificate and its private key used by an SSL server in ONTAP. * client_ca - a Certificate Authority certificate used by an SSL server in ONTAP to verify an SSL client certificate. * server_ca - a Certificate Authority certificate used by an SSL client in ONTAP to verify an SSL server certificate. * root_ca - a self-signed certificate used by ONTAP to sign other certificates by acting as a Certificate Authority.  | [optional] [default to null]
**Uuid** | **string** | Unique ID that identifies a certificate. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


