# FpolicyEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Specifies the name to assign to the external server configuration. | [optional] 
**Port** | **int32** | Port number of the FPolicy server application. | [optional] 
**PrimaryServers** | **[]string** |  | [optional] 
**SecondaryServers** | **[]string** |  | [optional] 
**Type** | **string** | The notification mode determines what ONTAP does after sending notifications to FPolicy servers.   The possible values are:     * synchronous  - After sending a notification, wait for a response from the FPolicy server.     * asynchronous - After sending a notification, file request processing continues.  | [optional] [default to TYPE_SYNCHRONOUS]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


