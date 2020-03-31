# KeyServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Password** | **string** | Password credentials for connecting with the key server. This is not audited. | [optional] [default to ]
**Records** | [**[]KeyServerRecords**](key_server_records.md) | An array of key servers specified to add multiple key servers to a key manager in a single API call. Valid in POST only and not valid if &#x60;server&#x60; is provided.  | [optional] 
**Server** | **string** | External key server for key management. If no port is provided, a default port of 5696 is used. Not valid in POST if &#x60;records&#x60; is provided. | [optional] 
**Timeout** | **int32** | I/O timeout in seconds for communicating with the key server. | [optional] 
**Username** | **string** | KMIP username credentials for connecting with the key server. | [optional] [default to ]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


