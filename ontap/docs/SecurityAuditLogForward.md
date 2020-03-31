# SecurityAuditLogForward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Destination syslog|splunk host to forward audit records to. This can be an IP address (IPv4|IPv6) or a hostname. | [optional] 
**Facility** | **string** | This is the standard Syslog Facility value that is used when sending audit records to a remote server. | [optional] 
**Port** | **int32** | Destination Port. The default port depends on the protocol chosen: For un-encrypted destinations the default port is 514. For encrypted destinations the default port is 6514.  | [optional] 
**Protocol** | **string** | Log forwarding protocol | [optional] [default to PROTOCOL_UDP_UNENCRYPTED]
**VerifyServer** | **bool** | This is only applicable when the protocol is tcp_encrypted. This controls whether the remote server&#39;s certificate is validated. Setting \&quot;verify_server\&quot; to \&quot;true\&quot; will enforce validation of remote server&#39;s certificate. Setting \&quot;verify_server\&quot; to \&quot;false\&quot; will not enforce validation of remote server&#39;s certificate. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


