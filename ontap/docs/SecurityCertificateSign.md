# SecurityCertificateSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTime** | **string** | Certificate expiration time. The allowed expiration time range is between 1 day to 10 years. | [optional] [default to 365days]
**HashFunction** | **string** | Hashing function | [optional] [default to HASH_FUNCTION_SHA256]
**SigningRequest** | **string** | Certificate signing request to be signed by the given certificate authority. Request should be in X509 PEM format. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


