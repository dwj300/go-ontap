# LdapService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**AdDomain** | **string** | This parameter specifies the name of the Active Directory domain used to discover LDAP servers for use by this client. This is mutually exclusive with &#x60;servers&#x60; during POST and PATCH.  | [optional] [default to null]
**BaseDn** | **string** | Specifies the default base DN for all searches. | [optional] [default to null]
**BaseScope** | **string** | Specifies the default search scope for LDAP queries: * base - search the named entry only * onelevel - search all entries immediately below the DN * subtree - search the named DN entry and the entire subtree below the DN  | [optional] [default to null]
**BindDn** | **string** | Specifies the user that binds to the LDAP servers. | [optional] [default to null]
**BindPassword** | **string** | Specifies the bind password for the LDAP servers. | [optional] [default to null]
**MinBindLevel** | **string** | The minimum bind authentication level. Possible values are: * anonymous - anonymous bind * simple - simple bind * sasl - Simple Authentication and Security Layer (SASL) bind  | [optional] [default to null]
**Port** | **int32** | The port used to connect to the LDAP Servers. | [optional] [default to null]
**PreferredAdServers** | **[]string** |  | [optional] [default to null]
**Schema** | **string** | The name of the schema template used by the SVM. * AD-IDMU - Active Directory Identity Management for UNIX * AD-SFU - Active Directory Services for UNIX * MS-AD-BIS - Active Directory Identity Management for UNIX * RFC-2307 - Schema based on RFC 2307 * Custom schema  | [optional] [default to null]
**Servers** | **[]string** |  | [optional] [default to null]
**SessionSecurity** | **string** | Specifies the level of security to be used for LDAP communications: * none - no signing or sealing * sign - sign LDAP traffic * seal - seal and sign LDAP traffic  | [optional] [default to null]
**Svm** | [***DnsSvm**](dns_svm.md) |  | [optional] [default to null]
**UseStartTls** | **bool** | Specifies whether or not to use Start TLS over LDAP connections.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


