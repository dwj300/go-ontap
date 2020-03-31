# AutosupportMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | Destination for the AutoSupport | [optional] [readonly] 
**Error** | [**AutosupportMessageError**](autosupport_message_error.md) |  | [optional] 
**GeneratedOn** | [**time.Time**](time.Time.md) | Date and Time of AutoSupport generation in ISO-8601 format | [optional] [readonly] 
**Index** | **int32** | Sequence number of the AutoSupport | [optional] [readonly] 
**Message** | **string** | Message included in the AutoSupport subject | [optional] 
**Node** | [**InlineResponse201Node**](inline_response_201_node.md) |  | [optional] 
**State** | **string** | State of AutoSupport delivery | [optional] [readonly] 
**Subject** | **string** | Subject line for the AutoSupport | [optional] [readonly] 
**Type** | **string** | Type of AutoSupport collection to issue | [optional] [default to TYPE_ALL]
**Uri** | **string** | Alternate destination for the AutoSupport | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


