# SecurityAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Application** | **string** | This identifies the \&quot;application\&quot; by which the request was processed.  | [optional] [readonly] 
**CommandId** | **string** | This is the command ID for this request. Each command received on a CLI session is assigned a command ID. This enables you to correlate a request and response.  | [optional] [readonly] 
**Index** | **int32** | Internal index for accessing records with same time/node. This is a 64 bit unsigned value. | [optional] [readonly] 
**Input** | **string** | The request. | [optional] [readonly] 
**Location** | **string** | This identifies the location of the remote user. This is an IP address or \&quot;console\&quot;. | [optional] [readonly] 
**Message** | **string** | This is an optional field that might contain \&quot;error\&quot; or \&quot;additional information\&quot; about the status of a command. | [optional] [readonly] 
**Node** | [**SecurityAuditLogNode**](security_audit_log_node.md) |  | [optional] 
**Scope** | **string** | Set to \&quot;svm\&quot; when the request is on a data SVM; otherwise set to \&quot;cluster\&quot;. | [optional] 
**SessionId** | **string** | This is the session ID on which the request is received. Each SSH session is assigned a session ID. Each http/ontapi/snmp request is assigned a unique session ID.  | [optional] [readonly] 
**State** | **string** | State of of this request. | [optional] [readonly] 
**Svm** | [**SecurityAuditLogSvm**](security_audit_log_svm.md) |  | [optional] 
**Timestamp** | **string** | Log entry timestamp. Valid in URL | [optional] [readonly] 
**User** | **string** | Username of the remote user. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


