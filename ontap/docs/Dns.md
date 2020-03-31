# Dns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Domains** | **[]string** | A list of DNS domains. Domain names have the following requirements: * The name must contain only the following characters: A through Z,   a through z, 0 through 9, \&quot;.\&quot;, \&quot;-\&quot; or \&quot;_\&quot;. * The first character of each label, delimited by \&quot;.\&quot;, must be one   of the following characters: A through Z or a through z or 0   through 9. * The last character of each label, delimited by \&quot;.\&quot;, must be one of   the following characters: A through Z, a through z, or 0 through 9. * The top level domain must contain only the following characters: A   through Z, a through z. * The system reserves the following names:\&quot;all\&quot;, \&quot;local\&quot;, and \&quot;localhost\&quot;.  | [optional] 
**Servers** | **[]string** | The list of IP addresses of the DNS servers. Addresses can be either IPv4 or IPv6 addresses.  | [optional] 
**Svm** | [**DnsSvm**](dns_svm.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


