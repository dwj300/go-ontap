# SvmLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomain** | **string** | This parameter specifies the name of the Active Directory domain used to discover LDAP servers for use by this client. This is mutually exclusive with &#x60;servers&#x60; during POST.  | [optional] 
**BaseDn** | **string** | Specifies the default base DN for all searches. | [optional] 
**BindDn** | **string** | Specifies the user that binds to the LDAP servers. SVM API supports anonymous binding. For Simple and SASL LDAP binding, use the LDAP API endpoint. | [optional] 
**Enabled** | **bool** | Enable LDAP? Setting to true creates a configuration if not already created. | [optional] 
**Servers** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


