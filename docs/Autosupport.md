# Autosupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactSupport** | **bool** | Specifies whether to send the AutoSupport messages to vendor support. | [optional] 
**Enabled** | **bool** | Specifies whether the AutoSupport daemon is enabled.  When this setting is disabled, delivery of all AutoSupport messages is turned off. | [optional] 
**From** | **string** | The e-mail address from which the AutoSupport messages are sent. To generate node-specific &#39;from&#39; addresses, enable &#39;-node-specific-from&#39; parameter via ONTAP CLI. | [optional] 
**IsMinimal** | **bool** | Specifies whether the system information is collected in compliant form, to remove private data or in complete form, to enhance diagnostics. | [optional] 
**Issues** | [**[]AutosupportIssues**](autosupport_issues.md) | A list of nodes in the cluster with connectivity issues to HTTP/SMTP/AOD AutoSupport destinations along with the corresponding error descriptions and corrective actions. | [optional] [readonly] 
**MailHosts** | **[]string** | The names of the mail servers used to deliver AutoSupport messages via SMTP. | [optional] 
**PartnerAddresses** | **[]string** | The list of partner addresses. | [optional] 
**ProxyUrl** | **string** | Proxy server for AutoSupport message delivery via HTTP/S. Optionally specify a username/password for authentication with the proxy server. | [optional] 
**To** | **[]string** | The e-mail addresses to which the AutoSupport messages are sent. | [optional] 
**Transport** | **string** | The name of the transport protocol used to deliver AutoSupport messages. | [optional] [default to TRANSPORT_HTTPS]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


