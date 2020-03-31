# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Code** | **int32** | If the state indicates \&quot;failure\&quot;, this is the final error code. | [optional] [readonly] 
**Description** | **string** | The description of the job to help identify it independent of the UUID. | [optional] [readonly] 
**EndTime** | [**time.Time**](time.Time.md) | The time the job ended. | [optional] [readonly] 
**Message** | **string** | A message corresponding to the state of the job providing additional details about the current state. | [optional] [readonly] 
**StartTime** | [**time.Time**](time.Time.md) | The time the job started. | [optional] [readonly] 
**State** | **string** | The state of the job. | [optional] [readonly] 
**Uuid** | **string** |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


