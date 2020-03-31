# \SecurityApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountCollectionGet**](SecurityApi.md#AccountCollectionGet) | **Get** /security/accounts | 
[**AccountCreate**](SecurityApi.md#AccountCreate) | **Post** /security/accounts | 
[**AccountDelete**](SecurityApi.md#AccountDelete) | **Delete** /security/accounts/{owner.uuid}/{name} | 
[**AccountGet**](SecurityApi.md#AccountGet) | **Get** /security/accounts/{owner.uuid}/{name} | 
[**AccountModify**](SecurityApi.md#AccountModify) | **Patch** /security/accounts/{owner.uuid}/{name} | 
[**AccountPasswordModify**](SecurityApi.md#AccountPasswordModify) | **Post** /security/authentication/password | 
[**AuditLogForwardingGet**](SecurityApi.md#AuditLogForwardingGet) | **Get** /security/audit/destinations | 
[**ClusterLdapCreate**](SecurityApi.md#ClusterLdapCreate) | **Post** /security/authentication/cluster/ldap | 
[**ClusterLdapDelete**](SecurityApi.md#ClusterLdapDelete) | **Delete** /security/authentication/cluster/ldap | 
[**ClusterLdapGet**](SecurityApi.md#ClusterLdapGet) | **Get** /security/authentication/cluster/ldap | 
[**ClusterLdapModify**](SecurityApi.md#ClusterLdapModify) | **Patch** /security/authentication/cluster/ldap | 
[**ClusterNisCreate**](SecurityApi.md#ClusterNisCreate) | **Post** /security/authentication/cluster/nis | 
[**ClusterNisDelete**](SecurityApi.md#ClusterNisDelete) | **Delete** /security/authentication/cluster/nis | 
[**ClusterNisGet**](SecurityApi.md#ClusterNisGet) | **Get** /security/authentication/cluster/nis | 
[**ClusterNisModify**](SecurityApi.md#ClusterNisModify) | **Patch** /security/authentication/cluster/nis | 
[**LoginMessagesCollectionGet**](SecurityApi.md#LoginMessagesCollectionGet) | **Get** /security/login/messages | 
[**LoginMessagesGet**](SecurityApi.md#LoginMessagesGet) | **Get** /security/login/messages/{uuid} | 
[**LoginMessagesModify**](SecurityApi.md#LoginMessagesModify) | **Patch** /security/login/messages/{uuid} | 
[**RoleCollectionGet**](SecurityApi.md#RoleCollectionGet) | **Get** /security/roles | 
[**RoleCreate**](SecurityApi.md#RoleCreate) | **Post** /security/roles | 
[**RoleDelete**](SecurityApi.md#RoleDelete) | **Delete** /security/roles/{owner.uuid}/{name} | 
[**RoleGet**](SecurityApi.md#RoleGet) | **Get** /security/roles/{owner.uuid}/{name} | 
[**RolePrivilegeCollectionGet**](SecurityApi.md#RolePrivilegeCollectionGet) | **Get** /security/roles/{owner.uuid}/{name}/privileges | 
[**RolePrivilegeCreate**](SecurityApi.md#RolePrivilegeCreate) | **Post** /security/roles/{owner.uuid}/{name}/privileges | 
[**RolePrivilegeDelete**](SecurityApi.md#RolePrivilegeDelete) | **Delete** /security/roles/{owner.uuid}/{name}/privileges/{path} | 
[**RolePrivilegeGet**](SecurityApi.md#RolePrivilegeGet) | **Get** /security/roles/{owner.uuid}/{name}/privileges/{path} | 
[**RolePrivilegeModify**](SecurityApi.md#RolePrivilegeModify) | **Patch** /security/roles/{owner.uuid}/{name}/privileges/{path} | 
[**SecurityAuditGet**](SecurityApi.md#SecurityAuditGet) | **Get** /security/audit | 
[**SecurityAuditLogCollectionGet**](SecurityApi.md#SecurityAuditLogCollectionGet) | **Get** /security/audit/messages | 
[**SecurityAuditModify**](SecurityApi.md#SecurityAuditModify) | **Patch** /security/audit | 
[**SecurityCertificateCollectionGet**](SecurityApi.md#SecurityCertificateCollectionGet) | **Get** /security/certificates | 
[**SecurityCertificateCreate**](SecurityApi.md#SecurityCertificateCreate) | **Post** /security/certificates | 
[**SecurityCertificateDelete**](SecurityApi.md#SecurityCertificateDelete) | **Delete** /security/certificates/{uuid} | 
[**SecurityCertificateGet**](SecurityApi.md#SecurityCertificateGet) | **Get** /security/certificates/{uuid} | 
[**SecurityCertificateSign**](SecurityApi.md#SecurityCertificateSign) | **Post** /security/certificates/{ca.uuid}/sign | 
[**SecurityKeyManagerCollectionGet**](SecurityApi.md#SecurityKeyManagerCollectionGet) | **Get** /security/key-managers | 
[**SecurityKeyManagerDelete**](SecurityApi.md#SecurityKeyManagerDelete) | **Delete** /security/key-managers/{uuid} | 
[**SecurityKeyManagerEnable**](SecurityApi.md#SecurityKeyManagerEnable) | **Post** /security/key-managers | 
[**SecurityKeyManagerGet**](SecurityApi.md#SecurityKeyManagerGet) | **Get** /security/key-managers/{uuid} | 
[**SecurityKeyManagerKeyServersAdd**](SecurityApi.md#SecurityKeyManagerKeyServersAdd) | **Post** /security/key-managers/{uuid}/key-servers | 
[**SecurityKeyManagerKeyServersCollectionGet**](SecurityApi.md#SecurityKeyManagerKeyServersCollectionGet) | **Get** /security/key-managers/{uuid}/key-servers | 
[**SecurityKeyManagerKeyServersDelete**](SecurityApi.md#SecurityKeyManagerKeyServersDelete) | **Delete** /security/key-managers/{uuid}/key-servers/{server} | 
[**SecurityKeyManagerKeyServersGet**](SecurityApi.md#SecurityKeyManagerKeyServersGet) | **Get** /security/key-managers/{uuid}/key-servers/{server} | 
[**SecurityKeyManagerKeyServersModify**](SecurityApi.md#SecurityKeyManagerKeyServersModify) | **Patch** /security/key-managers/{uuid}/key-servers/{server} | 
[**SecurityKeyManagerModify**](SecurityApi.md#SecurityKeyManagerModify) | **Patch** /security/key-managers/{uuid} | 
[**SecurityLogForwardingCreate**](SecurityApi.md#SecurityLogForwardingCreate) | **Post** /security/audit/destinations | 
[**SecurityLogForwardingDelete**](SecurityApi.md#SecurityLogForwardingDelete) | **Delete** /security/audit/destinations/{address}/{port} | 
[**SecurityLogForwardingGet**](SecurityApi.md#SecurityLogForwardingGet) | **Get** /security/audit/destinations/{address}/{port} | 
[**SecurityLogForwardingModify**](SecurityApi.md#SecurityLogForwardingModify) | **Patch** /security/audit/destinations/{address}/{port} | 
[**SecuritySamlSpCreate**](SecurityApi.md#SecuritySamlSpCreate) | **Post** /security/authentication/cluster/saml-sp | 
[**SecuritySamlSpDelete**](SecurityApi.md#SecuritySamlSpDelete) | **Delete** /security/authentication/cluster/saml-sp | 
[**SecuritySamlSpGet**](SecurityApi.md#SecuritySamlSpGet) | **Get** /security/authentication/cluster/saml-sp | 
[**SecuritySamlSpModify**](SecurityApi.md#SecuritySamlSpModify) | **Patch** /security/authentication/cluster/saml-sp | 



## AccountCollectionGet

> AccountResponse AccountCollectionGet(ctx, optional)



Retrieves a list of user accounts in the cluster. ### Related ONTAP commands * `security login show` ### Learn more * [`DOC /security/accounts`](#docs-security-security_accounts) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**AccountResponse**](account_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountCreate

> AccountCreate(ctx, optional)



Creates a new user account. ### Required parameters * `name` - Account name to be created. * `applications` - Array of one or more application tuples (of application and authentication methods). ### Optional parameters * `owner.name` or `owner.uuid`  - Name or UUID of the SVM for an SVM-scoped user account. If not supplied, a cluster-scoped user account is created. * `role` - RBAC role for the user account. Defaulted to `admin` for cluster user account and to `vsadmin` for SVM-scoped account. * `password` - Password for the user account (if the authentication method is opted as password for one or more of applications). * `second_authentication_method` - Needed for MFA and only supported for ssh application. Defaults to `none` if not supplied. * `comment` - Comment for the user account (e.g purpose of this account). * `locked` - Locks the account after creation. Defaults to `false` if not supplied. ### Related ONTAP commands * `security login create` ### Learn more * [`DOC /security/accounts`](#docs-security-security_accounts) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Account**](Account.md)| Details for the user account to be created. | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountDelete

> AccountDelete(ctx, ownerUuid, name)



Deletes a user account. ### Required parameters * `name` - Account name to be deleted. * `owner.uuid`  - UUID of the SVM housing the user account to be deleted. ### Related ONTAP commands * `security login delete` ### Learn more * [`DOC /security/accounts/{owner.uuid}/{name}`](#docs-security-security_accounts_{owner.uuid}_{name}) * [`DOC /security/accounts`](#docs-security-security_accounts) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the account. | 
**name** | **string**| User account name | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGet

> AccountResponse AccountGet(ctx, ownerUuid, name)



Retrieves a specific user account. ### Related ONTAP commands * `security login show` ### Learn more * [`DOC /security/accounts/{owner.uuid}/{name}`](#docs-security-security_accounts_{owner.uuid}_{name}) * [`DOC /security/accounts`](#docs-security-security_accounts) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the account. | 
**name** | **string**| User account name | 

### Return type

[**AccountResponse**](account_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountModify

> AccountModify(ctx, ownerUuid, name, info)



Updates a user account. Locks or unlocks a user account and/or updates the role, applications, and/or password for the user account. ### Required parameters * `name` - Account name to be updated. * `owner.uuid`  - UUID of the SVM housing the user account to be updated. ### Optional parameters * `applications` - Array of one or more tuples (of application and authentication methods). * `role` - RBAC role for the user account. * `password` - Password for the user account (if the authentication method is opted as password for one or more of applications). * `second_authentication_method` - Needed for MFA and only supported for ssh application. Defaults to `none` if not supplied. * `comment` - Comment for the user account (e.g purpose of this account). * `locked` - Set to true/false to lock/unlock the account. ### Related ONTAP commands * `security login create` * `security login modify` * `security login password` * `security login lock` * `security login unlock` ### Learn more * [`DOC /security/accounts/{owner.uuid}/{name}`](#docs-security-security_accounts_{owner.uuid}_{name}) * [`DOC /security/accounts`](#docs-security-security_accounts) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the account. | 
**name** | **string**| User account name | 
**info** | [**Account**](Account.md)| User account details | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountPasswordModify

> AccountPasswordModify(ctx, info)



Updates the password for a user account. ### Required parameters * `name` - User account name. * `password` - New password for the user account. ### Optional parameters * `owner.name` or `owner.uuid` - Name or UUID of the SVM for an SVM-scoped user account. ### Related ONTAP commands * `security login password` ### Learn more * [`DOC /security/authentication/password`](#docs-security-security_authentication_password) * [`DOC /security/accounts`](#docs-security-security_accounts) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**AccountPassword**](AccountPassword.md)| A new password for the user account. | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogForwardingGet

> SecurityAuditLogForwardResponse AuditLogForwardingGet(ctx, optional)



Defines remote syslog/splunk server for sending audit information ### Learn more * [`DOC /security/audit/destinations`](#docs-security-security_audit_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditLogForwardingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditLogForwardingGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **port** | **optional.Int32**| Filter by port | 
 **address** | **optional.String**| Filter by address | 
 **protocol** | **optional.String**| Filter by protocol | 
 **verifyServer** | **optional.Bool**| Filter by verify_server | 
 **facility** | **optional.String**| Filter by facility | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]

### Return type

[**SecurityAuditLogForwardResponse**](security_audit_log_forward_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterLdapCreate

> LdapServiceResponse ClusterLdapCreate(ctx, optional)



A cluster can have only one LDAP configuration. IPv6 must be enabled if IPv6 family addresses are specified. The following parameters are optional: - schema - port - min_bind_level - bind_password - base_scope - use_start_tls - session_security Configuring more than one LDAP server is recommended to avoid a single point of failure. Both FQDNs and IP addresses are supported for the 'servers' field. The LDAP servers are validated as part of this operation. LDAP validation fails in the following scenarios:<br/> 1. The server does not have LDAP installed. 2. The server is invalid. 3. The server is unreachable.<br/>  ### Learn more * [`DOC /security/authentication/cluster/ldap`](#docs-security-security_authentication_cluster_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterLdapCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterLdapCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ClusterLdap**](ClusterLdap.md)| Info specification | 

### Return type

[**LdapServiceResponse**](ldap_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterLdapDelete

> ClusterLdapDelete(ctx, )



The DELETE operation removes the LDAP configuration of the cluster.  ### Learn more * [`DOC /security/authentication/cluster/ldap`](#docs-security-security_authentication_cluster_ldap)

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterLdapGet

> ClusterLdap ClusterLdapGet(ctx, optional)



Retrieves the cluster LDAP configuration.  ### Learn more * [`DOC /security/authentication/cluster/ldap`](#docs-security-security_authentication_cluster_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterLdapGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterLdapGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ClusterLdap**](cluster_ldap.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterLdapModify

> ClusterLdapModify(ctx, optional)



Both mandatory and optional parameters of the LDAP configuration can be updated. IPv6 must be enabled if IPv6 family addresses are specified. Configuring more than one LDAP server is recommended to avoid a single point of failure. Both FQDNs and IP addresses are supported for the 'servers' field. The LDAP servers are validated as part of this operation. LDAP validation fails in the following scenarios:<br/> 1. The server does not have LDAP installed. 2. The server is invalid. 3. The server is unreachable<br/>  ### Learn more * [`DOC /security/authentication/cluster/ldap`](#docs-security-security_authentication_cluster_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterLdapModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterLdapModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ClusterLdap**](ClusterLdap.md)| Info specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterNisCreate

> ClusterNisServiceResponse ClusterNisCreate(ctx, optional)



The cluster can have one NIS server configuration. Specify the NIS domain and NIS servers as input. Domain name and servers fields cannot be empty. Both FQDNs and IP addresses are supported for the 'servers' field. IPv6 must be enabled if IPv6 family addresses are specified in the 'servers' field. A maximum of ten NIS servers are supported.  ### Learn more * [`DOC /security/authentication/cluster/nis`](#docs-security-security_authentication_cluster_nis)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterNisCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterNisCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ClusterNisService**](ClusterNisService.md)| Info specification | 

### Return type

[**ClusterNisServiceResponse**](cluster_nis_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterNisDelete

> ClusterNisDelete(ctx, )



The DELETE operation removes the NIS configuration of the cluster. NIS can be removed as a source from ns-switch if NIS is not used for lookups.  ### Learn more * [`DOC /security/authentication/cluster/nis`](#docs-security-security_authentication_cluster_nis)

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterNisGet

> ClusterNisService ClusterNisGet(ctx, optional)



Retrieves the NIS configuration of the cluster. Both NIS domain and servers are displayed by default. The 'bound servers' field indicates the successfully bound NIS servers.  ### Learn more * [`DOC /security/authentication/cluster/nis`](#docs-security-security_authentication_cluster_nis)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterNisGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterNisGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ClusterNisService**](cluster_nis_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClusterNisModify

> ClusterNisModify(ctx, optional)



Both NIS domain and servers can be modified. Domains and servers cannot be empty. Both FQDNs and IP addresses are supported for the 'servers' field. If the domain is modified, NIS servers must also be specified. IPv6 must be enabled if IPv6 family addresses are specified for the 'servers' field.<br/>  ### Learn more * [`DOC /security/authentication/cluster/nis`](#docs-security-security_authentication_cluster_nis)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterNisModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterNisModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ClusterNisService**](ClusterNisService.md)| Info specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginMessagesCollectionGet

> LoginMessagesResponse LoginMessagesCollectionGet(ctx, optional)



Retrieves the login banner and messages of the day (MOTD) configured in the cluster and in specific SVMs.  ### Learn more * [`DOC /security/login/messages`](#docs-security-security_login_messages)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginMessagesCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginMessagesCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **banner** | **optional.String**| Filter by banner | 
 **uuid** | **optional.String**| Filter by uuid | 
 **showClusterMessage** | **optional.Bool**| Filter by show_cluster_message | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **message** | **optional.String**| Filter by message | 
 **scope** | **optional.String**| Filter by scope | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**LoginMessagesResponse**](login_messages_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginMessagesGet

> LoginMessages LoginMessagesGet(ctx, uuid, optional)



Retrieves the login messages configuration by UUID. ### Learn more * [`DOC /security/login/messages`](#docs-security-security_login_messages)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Login messages configuration UUID | 
 **optional** | ***LoginMessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginMessagesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**LoginMessages**](login_messages.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginMessagesModify

> LoginMessagesModify(ctx, uuid, optional)



Updates the login messages configuration. There are no required fields. An empty body will make no modifications.  ### Learn more * [`DOC /security/login/messages`](#docs-security-security_login_messages)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Login messages configuration UUID | 
 **optional** | ***LoginMessagesModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginMessagesModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of LoginMessages**](LoginMessages.md)| Info specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleCollectionGet

> RoleResponse RoleCollectionGet(ctx, optional)



Retrieves a list of roles configured in the cluster. ### Related ONTAP commands * `security login rest-role show` ### Learn more * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RoleCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**RoleResponse**](role_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleCreate

> RoleCreate(ctx, info)



Creates a new cluster-scoped role or an SVM-scoped role. For an SVM-scoped role, specify either the SVM name as the owner.name or SVM UUID as the owner.uuid in the request body along with other parameters for the role. The owner.uuid or owner.name are not required to be specified for a cluster-scoped role. ### Required parameters * `name` - Name of the role to be created. * `privileges` - Array of privilege tuples. Each tuple consists of a REST API path and its desired access level. ### Optional parameters * `owner.name` or `owner.uuid`  - Name or UUID of the SVM for an SVM-scoped role. ### Related ONTAP commands * `security login rest-role create` ### Learn more * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**Role**](Role.md)| Role specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleDelete

> RoleDelete(ctx, ownerUuid, name)



Delete the specified role ### Required parameters * `name` - Name of the role to be deleted. * `owner.uuid` - UUID of the SVM housing the role. ### Related ONTAP commands * `security login rest-role delete` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}`](#docs-security-security_roles_{owner.uuid}_{name}) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name to be deleted. | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleGet

> RoleResponse RoleGet(ctx, ownerUuid, name)



Retrieves the details of the specified role. ### Related ONTAP commands * `security login rest-role show` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}`](#docs-security-security_roles_{owner.uuid}_{name}) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name | 

### Return type

[**RoleResponse**](role_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePrivilegeCollectionGet

> RolePrivilegeResponse RolePrivilegeCollectionGet(ctx, ownerUuid, name)



Retrieves privilege details of the specified role. ### Related ONTAP commands * `security login rest-role show` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}/privileges`](#docs-security-security_roles_{owner.uuid}_{name}_privileges) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name | 

### Return type

[**RolePrivilegeResponse**](role_privilege_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePrivilegeCreate

> RolePrivilegeCreate(ctx, ownerUuid, name, info)



Add a privilege tuple (of REST URI and its access level) to an existing role. ### Required parameters * `owner.uuid` - UUID of the SVM that houses this role. * `name` - Name of the role to be updated. * `path` - REST URI path (example: \"/api/storage/volumes\"). * `access` - Desired access level for the REST URI path (one of \"all\", \"readonly\" or \"none\"). ### Optional parameters none ### Related ONTAP commands * `security login rest-role create` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}/privileges`](#docs-security-security_roles_{owner.uuid}_{name}_privileges) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name | 
**info** | [**RolePrivilege**](RolePrivilege.md)| Role privilege specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePrivilegeDelete

> RolePrivilegeDelete(ctx, ownerUuid, name, path)



Delete a privilege tuple (of REST URI and its access level) from the role. ### Required parameters * `owner.uuid` - UUID of the SVM which houses this role. * `name` - Name of the role to be updated. * `path` - Constituent REST API path to be deleted from this role. ### Related ONTAP commands * `security login rest-role delete` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}/privileges/{path}`](#docs-security-security_roles_{owner.uuid}_{name}_privileges_{path}) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name | 
**path** | **string**| REST API path | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePrivilegeGet

> RolePrivilege RolePrivilegeGet(ctx, ownerUuid, name, path)



Retrieves the privilege level for a REST API path for the specified role. ### Related ONTAP commands * `security login rest-role show` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}/privileges/{path}`](#docs-security-security_roles_{owner.uuid}_{name}_privileges_{path}) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name | 
**path** | **string**| REST API path | 

### Return type

[**RolePrivilege**](role_privilege.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePrivilegeModify

> RolePrivilegeModify(ctx, ownerUuid, name, path, access)



Updates the privilege level for a REST API path. ### Required parameters * `owner.uuid` - UUID of the SVM that houses this role. * `name` - Name of the role to be updated. * `path` - Constituent REST API path whose access level has to be updated. * `access` - Access level for the path (one of \"all\", \"readonly\", or \"none\") ### Related ONTAP commands * `security login rest-role modify` ### Learn more * [`DOC /security/roles/{owner.uuid}/{name}/privileges/{path}`](#docs-security-security_roles_{owner.uuid}_{name}_privileges_{path}) * [`DOC /security/roles`](#docs-security-security_roles) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUuid** | **string**| Owner UUID of the role. | 
**name** | **string**| Role name | 
**path** | **string**| REST API path | 
**access** | [**RolePrivilege**](RolePrivilege.md)|  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityAuditGet

> SecurityAudit SecurityAuditGet(ctx, optional)



Retrieves administrative audit settings for GET operations. ### Learn more * [`DOC /security/audit`](#docs-security-security_audit)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityAuditGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityAuditGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SecurityAudit**](security_audit.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityAuditLogCollectionGet

> SecurityAuditLogResponse SecurityAuditLogCollectionGet(ctx, optional)



Retrieves the administrative audit log viewer. ### Learn more * [`DOC /security/audit/messages`](#docs-security-security_audit_messages)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityAuditLogCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityAuditLogCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commandId** | **optional.String**| Filter by command_id | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **sessionId** | **optional.String**| Filter by session_id | 
 **scope** | **optional.String**| Filter by scope | 
 **timestamp** | **optional.String**| Filter by timestamp | 
 **user** | **optional.String**| Filter by user | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **location** | **optional.String**| Filter by location | 
 **message** | **optional.String**| Filter by message | 
 **index** | **optional.Int32**| Filter by index | 
 **input** | **optional.String**| Filter by input | 
 **application** | **optional.String**| Filter by application | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SecurityAuditLogResponse**](security_audit_log_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityAuditModify

> SecurityAudit SecurityAuditModify(ctx, optional)



Updates administrative audit settings for GET operations. All of the fields are optional. An empty body will make no changes.  ### Learn more * [`DOC /security/audit`](#docs-security-security_audit)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityAuditModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityAuditModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SecurityAudit**](SecurityAudit.md)| Info specification | 

### Return type

[**SecurityAudit**](security_audit.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCertificateCollectionGet

> SecurityCertificateResponse SecurityCertificateCollectionGet(ctx, optional)



Retrieves security certificates. ### Related ONTAP commands * `security certificate show`  ### Learn more * [`DOC /security/certificates`](#docs-security-security_certificates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityCertificateCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityCertificateCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Filter by scope | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **commonName** | **optional.String**| Filter by common_name | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **ca** | **optional.String**| Filter by ca | 
 **type_** | **optional.String**| Filter by certificate type | 
 **keySize** | **optional.String**| Filter by key_size | 
 **expiryTime** | **optional.String**| Filter by expiry_time | 
 **hashFunction** | **optional.String**| Filter by hash_function | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SecurityCertificateResponse**](security_certificate_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCertificateCreate

> SecurityCertificateResponse SecurityCertificateCreate(ctx, optional)



Creates or installs a certificate. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create or install the certificate. * `common_name` - Common name of the certificate. Required when creating a certificate. * `type` - Type of certificate. * `public_certificate` - Public key certificate in PEM format. Required when installing a certificate. * `private_key` - Private key certificate in PEM format. Required when installing a CA-signed certificate. ### Recommended optional properties * `expiry_time` - Certificate expiration time. Specifying an expiration time is recommended when creating a certificate. * `key_size` - Key size of the certificate in bits. Specifying a strong key size is recommended when creating a certificate. ### Default property values If not specified in POST, the following default property values are assigned: * `key_size` - _2048_ * `expiry_time` - _P365DT_ * `hash_function` - _sha256_ ### Related ONTAP commands * `security certificate create` * `security certificate install`  ### Learn more * [`DOC /security/certificates`](#docs-security-security_certificates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityCertificateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityCertificateCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SecurityCertificate**](SecurityCertificate.md)| Info specification | 

### Return type

[**SecurityCertificateResponse**](security_certificate_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCertificateDelete

> SecurityCertificateDelete(ctx, uuid)



Deletes a security certificate. ### Related ONTAP commands * `security certificate delete`  ### Learn more * [`DOC /security/certificates`](#docs-security-security_certificates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Certificate UUID | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCertificateGet

> SecurityCertificate SecurityCertificateGet(ctx, uuid, optional)



Retrieves security certificates. ### Related ONTAP commands * `security certificate show`  ### Learn more * [`DOC /security/certificates`](#docs-security-security_certificates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Certificate UUID | 
 **optional** | ***SecurityCertificateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityCertificateGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SecurityCertificate**](security_certificate.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCertificateSign

> SecurityCertificateSignResponse SecurityCertificateSign(ctx, caUuid, optional)



Signs a certificate. ### Required properties * `signing_request` - Certificate signing request to be signed by the given certificate authority. ### Recommended optional properties * `expiry_time` - Certificate expiration time. Specifying an expiration time for a signed certificate is recommended. * `hash_function` - Hashing function. Specifying a strong hashing function is recommended when signing a certificate. ### Default property values If not specified in POST, the following default property values are assigned: * `expiry_time` - _P365DT_ * `hash_function` - _sha256_ ### Related ONTAP commands * `security certificate sign` This API is used to sign a certificate request using a pre-existing self-signed root certificate. The self-signed root certificate acts as a certificate authority within its scope and maintains the records of its signed certificates. <br/> The root certificate can be created for a given SVM or for the cluster using [`POST security/certificates`].<br/> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caUuid** | **string**| UUID of the existing certificate authority certificate | 
 **optional** | ***SecurityCertificateSignOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityCertificateSignOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of SecurityCertificateSign**](SecurityCertificateSign.md)| Certificate sign info specification | 

### Return type

[**SecurityCertificateSignResponse**](security_certificate_sign_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerCollectionGet

> SecurityKeyManagerResponse SecurityKeyManagerCollectionGet(ctx, optional)



Retrieves key managers. ### Related ONTAP commands * `security key-manager show-keystore` * `security key-manager external show`  ### Learn more * [`DOC /security/key-managers`](#docs-security-security_key-managers)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityKeyManagerCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalServerCaCertificatesUuid** | **optional.String**| Filter by external.server_ca_certificates.uuid | 
 **externalServersServer** | **optional.String**| Filter by external.servers.server | 
 **externalServersTimeout** | **optional.Int32**| Filter by external.servers.timeout | 
 **externalServersUsername** | **optional.String**| Filter by external.servers.username | 
 **externalClientCertificateUuid** | **optional.String**| Filter by external.client_certificate.uuid | 
 **uuid** | **optional.String**| Filter by uuid | 
 **onboardEnabled** | **optional.Bool**| Filter by onboard.enabled | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **scope** | **optional.String**| Filter by scope | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SecurityKeyManagerResponse**](security_key_manager_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerDelete

> SecurityKeyManagerDelete(ctx, uuid)



Deletes a key manager. ### Related ONTAP commands * `security key-manager external disable` * `security key-manager onboard disable`  ### Learn more * [`DOC /security/key-managers`](#docs-security-security_key-managers)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Key manager UUID | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerEnable

> SecurityKeyManagerResponse SecurityKeyManagerEnable(ctx, optional)



Creates a key manager. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create a key manager. * `external.client_certificate` - Client certificate. Required only when creating an external key manager. * `external.server_ca_certificates` - Server CA certificates. Required only when creating an external key manager. * `external.servers.server` - Key servers. Required only when creating an external key manager. * `onboard.passphrase` - Cluster-wide passphrase. Required only when creating an onboard key manager. ### Related ONTAP commands * `security key-manager external enable` * `security key-manager onboard enable`  ### Learn more * [`DOC /security/key-managers`](#docs-security-security_key-managers)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityKeyManagerEnableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerEnableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SecurityKeyManager**](SecurityKeyManager.md)| Info specification | 

### Return type

[**SecurityKeyManagerResponse**](security_key_manager_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerGet

> SecurityKeyManager SecurityKeyManagerGet(ctx, uuid, optional)



Retrieves key managers. ### Related ONTAP commands * `security key-manager show-keystore` * `security key-manager external show`  ### Learn more * [`DOC /security/key-managers`](#docs-security-security_key-managers)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Key manager UUID | 
 **optional** | ***SecurityKeyManagerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SecurityKeyManager**](security_key_manager.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerKeyServersAdd

> KeyServerResponse SecurityKeyManagerKeyServersAdd(ctx, uuid, optional)



Adds key servers to a configured external key manager. ### Required properties * `uuid` - UUID of the external key manager. * `server` - Key server name. ### Related ONTAP commands * `security key-manager external add-servers` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| External key manager UUID | 
 **optional** | ***SecurityKeyManagerKeyServersAddOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerKeyServersAddOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of KeyServer**](KeyServer.md)| Info specification | 

### Return type

[**KeyServerResponse**](key_server_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerKeyServersCollectionGet

> KeyServerResponse SecurityKeyManagerKeyServersCollectionGet(ctx, uuid, optional)



Retrieves key servers. ### Related ONTAP commands * `security key-manager external show` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| External key manager UUID | 
 **optional** | ***SecurityKeyManagerKeyServersCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerKeyServersCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Filter by username | 
 **timeout** | **optional.Int32**| Filter by timeout | 
 **server** | **optional.String**| Filter by server | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**KeyServerResponse**](key_server_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerKeyServersDelete

> SecurityKeyManagerKeyServersDelete(ctx, uuid, server)



Deletes a key server. ### Related ONTAP commands * `security key-manager external remove-servers` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| External key manager UUID | 
**server** | **string**| Key server configured in the external key manager | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerKeyServersGet

> KeyServer SecurityKeyManagerKeyServersGet(ctx, uuid, server, optional)



Retrieves key servers configured in an external key manager. ### Related ONTAP commands * `security key-manager external show` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| External key manager UUID | 
**server** | **string**| Key server configured in the key manager | 
 **optional** | ***SecurityKeyManagerKeyServersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerKeyServersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**KeyServer**](key_server.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerKeyServersModify

> SecurityKeyManagerKeyServersModify(ctx, uuid, server, optional)



Updates a key server. ### Related ONTAP commands * `security key-manager external modify-server` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| External key manager UUID | 
**server** | **string**| Key server configured in the external key manager | 
 **optional** | ***SecurityKeyManagerKeyServersModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerKeyServersModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of KeyServer**](KeyServer.md)| Key server info | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityKeyManagerModify

> SecurityKeyManagerModify(ctx, uuid, optional)



Updates a key manager. ### Related ONTAP commands * `security key-manager external modify` * `security key-manager onboard update-passphrase`  ### Learn more * [`DOC /security/key-managers`](#docs-security-security_key-managers)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Key manager UUID | 
 **optional** | ***SecurityKeyManagerModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityKeyManagerModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of SecurityKeyManager**](SecurityKeyManager.md)| Key manager info | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityLogForwardingCreate

> SecurityAuditLogForwardResponse SecurityLogForwardingCreate(ctx, info, optional)



Configures remote syslog/splunk server information. ### Required properties All of the following fields are required for creating a remote syslog/splunk destination * `address` ### Optional properties All of the following fields are optional for creating a remote syslog/splunk destination * `port` * `protocol` * `facility` * `verify_server` (Can only be \"true\" when protocol is \"tcp_encrypted\")  ### Learn more * [`DOC /security/audit/destinations`](#docs-security-security_audit_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**SecurityAuditLogForward**](SecurityAuditLogForward.md)| Remote syslog/splunk server information | 
 **optional** | ***SecurityLogForwardingCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityLogForwardingCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Skip the Connectivity Test | [default to false]

### Return type

[**SecurityAuditLogForwardResponse**](security_audit_log_forward_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityLogForwardingDelete

> SecurityLogForwardingDelete(ctx, address, port)



Deletes remote syslog/splunk server information. ### Learn more * [`DOC /security/audit/destinations`](#docs-security-security_audit_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| IP address of remote syslog/splunk server. | 
**port** | **int32**| Port number of remote syslog/splunk server. | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityLogForwardingGet

> SecurityAuditLogForward SecurityLogForwardingGet(ctx, address, port, optional)



Defines remote syslog/splunk server for sending audit information. ### Learn more * [`DOC /security/audit/destinations`](#docs-security-security_audit_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| IP address of remote syslog/splunk server | 
**port** | **int32**| Port number of remote syslog/splunk server | 
 **optional** | ***SecurityLogForwardingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityLogForwardingGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SecurityAuditLogForward**](security_audit_log_forward.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityLogForwardingModify

> SecurityAuditLogForward SecurityLogForwardingModify(ctx, address, port, optional)



Updates remote syslog/splunk server information. ### Learn more * [`DOC /security/audit/destinations`](#docs-security-security_audit_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| IP address of remote syslog/splunk server. | 
**port** | **int32**| Port number of remote syslog/splunk server. | 
 **optional** | ***SecurityLogForwardingModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityLogForwardingModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of SecurityAuditLogForward**](SecurityAuditLogForward.md)| Modify remote syslog/splunk server information. | 

### Return type

[**SecurityAuditLogForward**](security_audit_log_forward.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySamlSpCreate

> JobLinkResponse SecuritySamlSpCreate(ctx, optional)



Creates a SAML service provider configuration. Note that \"common_name\" is mutually exclusive with \"serial_number\" and \"ca\" in the POST. SAML will initially be disabled, requiring a patch to set \"enabled\" to \"true\", so that the user has time to complete the setup of the IdP. ### Required properties * `idp_uri` ### Optional properties * `certificate` * `enabled` * `host`  ### Learn more * [`DOC /security/authentication/cluster/saml-sp`](#docs-security-security_authentication_cluster_saml-sp)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritySamlSpCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecuritySamlSpCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyMetadataServer** | **optional.Bool**| Verify IdP metadata server identity. | [default to true]
 **info** | [**optional.Interface of SecuritySamlSp**](SecuritySamlSp.md)| Info specification | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySamlSpDelete

> SecuritySamlSpDelete(ctx, )



Deletes a SAML service provider configuration. ### Learn more * [`DOC /security/authentication/cluster/saml-sp`](#docs-security-security_authentication_cluster_saml-sp)

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySamlSpGet

> SecuritySamlSp SecuritySamlSpGet(ctx, optional)



Retrieves a SAML service provider configuration. ### Learn more * [`DOC /security/authentication/cluster/saml-sp`](#docs-security-security_authentication_cluster_saml-sp)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritySamlSpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecuritySamlSpGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SecuritySamlSp**](security_saml_sp.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySamlSpModify

> SecuritySamlSpModify(ctx, optional)



Updates a SAML service provider configuration. ### Learn more * [`DOC /security/authentication/cluster/saml-sp`](#docs-security-security_authentication_cluster_saml-sp)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritySamlSpModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecuritySamlSpModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SecuritySamlSp**](SecuritySamlSp.md)| Info specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

