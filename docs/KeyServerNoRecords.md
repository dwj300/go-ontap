# KeyServerNoRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Password** | **string** | Password credentials for connecting with the key server. This is not audited. | [optional] [default to ]
**Server** | **string** | External key server for key management. If no port is provided, a default port of 5696 is used. Not valid in POST if &#x60;records&#x60; is provided. | [optional] 
**Timeout** | **int32** | I/O timeout in seconds for communicating with the key server. | [optional] 
**Username** | **string** |  | [optional] [default to ]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


