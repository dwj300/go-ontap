# Autosupport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactSupport** | **bool** | Specifies whether to send the AutoSupport messages to vendor support. | [optional] [default to null]
**Enabled** | **bool** | Specifies whether the AutoSupport daemon is enabled.  When this setting is disabled, delivery of all AutoSupport messages is turned off. | [optional] [default to null]
**From** | **string** | The e-mail address from which the AutoSupport messages are sent. To generate node-specific &#39;from&#39; addresses, enable &#39;-node-specific-from&#39; parameter via ONTAP CLI. | [optional] [default to null]
**IsMinimal** | **bool** | Specifies whether the system information is collected in compliant form, to remove private data or in complete form, to enhance diagnostics. | [optional] [default to null]
**Issues** | [**[]AutosupportIssues**](autosupport_issues.md) | A list of nodes in the cluster with connectivity issues to HTTP/SMTP/AOD AutoSupport destinations along with the corresponding error descriptions and corrective actions. | [optional] [default to null]
**MailHosts** | **[]string** | The names of the mail servers used to deliver AutoSupport messages via SMTP. | [optional] [default to null]
**PartnerAddresses** | **[]string** | The list of partner addresses. | [optional] [default to null]
**ProxyUrl** | **string** | Proxy server for AutoSupport message delivery via HTTP/S. Optionally specify a username/password for authentication with the proxy server. | [optional] [default to null]
**To** | **[]string** | The e-mail addresses to which the AutoSupport messages are sent. | [optional] [default to null]
**Transport** | **string** | The name of the transport protocol used to deliver AutoSupport messages. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


