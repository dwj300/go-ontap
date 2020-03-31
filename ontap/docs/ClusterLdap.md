# ClusterLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**BaseDn** | **string** | Specifies the default base DN for all searches. | [optional] 
**BaseScope** | **string** | Specifies the default search scope for LDAP queries: * base - search the named entry only * onelevel - search all entries immediately below the DN * subtree - search the named DN entry and the entire subtree below the DN  | [optional] [default to BASE_SCOPE_SUBTREE]
**BindDn** | **string** | Specifies the user that binds to the LDAP servers. | [optional] 
**BindPassword** | **string** | Specifies the bind password for the LDAP servers. | [optional] 
**MinBindLevel** | **string** | The minimum bind authentication level. Possible values are: * anonymous - anonymous bind * simple - simple bind * sasl - Simple Authentication and Security Layer (SASL) bind  | [optional] [default to MIN_BIND_LEVEL_SIMPLE]
**Port** | **int32** | The port used to connect to the LDAP Servers. | [optional] 
**Schema** | **string** | The name of the schema template used by the SVM. * AD-IDMU - Active Directory Identity Management for UNIX * AD-SFU - Active Directory Services for UNIX * MS-AD-BIS - Active Directory Identity Management for UNIX * RFC-2307 - Schema based on RFC 2307 * Custom schema  | [optional] [default to RFC-2307]
**Servers** | **[]string** |  | [optional] 
**SessionSecurity** | **string** | Specifies the level of security to be used for LDAP communications: * none - no signing or sealing * sign - sign LDAP traffic * seal - seal and sign LDAP traffic  | [optional] [default to SESSION_SECURITY_NONE]
**UseStartTls** | **bool** | Specifies whether or not to use Start TLS over LDAP connections.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


