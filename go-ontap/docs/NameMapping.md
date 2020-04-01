# NameMapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**ClientMatch** | **string** | Client workstation IP Address which is matched when searching for the pattern.   You can specify the value in any of the following formats: * As an IPv4 address with a subnet mask expressed as a number of bits; for instance, 10.1.12.0/24 * As an IPv6 address with a subnet mask expressed as a number of bits; for instance, fd20:8b1e:b255:4071::/64 * As an IPv4 address with a network mask; for instance, 10.1.16.0/255.255.255.0 * As a hostname  | [optional] [default to null]
**Direction** | **string** | Direction in which the name mapping is applied. The possible values are:   * krb_unix  - Kerberos principal name to UNIX user name   * win_unix  - Windows user name to UNIX user name   * unix_win  - UNIX user name to Windows user name mapping  | [optional] [default to null]
**Index** | **int32** | Position in the list of name mappings. | [optional] [default to null]
**Pattern** | **string** | Pattern used to match the name while searching for a name that can be used as a replacement. The pattern is a UNIX-style regular expression. Regular expressions are case-insensitive when mapping from Windows to UNIX, and they are case-sensitive for mappings from Kerberos to UNIX and UNIX to Windows. | [optional] [default to null]
**Replacement** | **string** | The name that is used as a replacement, if the pattern associated with this entry matches. | [optional] [default to null]
**Svm** | [***DnsSvm**](dns_svm.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


