# \NameServicesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsCollectionGet**](NameServicesApi.md#DnsCollectionGet) | **Get** /name-services/dns | 
[**DnsCreate**](NameServicesApi.md#DnsCreate) | **Post** /name-services/dns | 
[**DnsDelete**](NameServicesApi.md#DnsDelete) | **Delete** /name-services/dns/{svm.uuid} | 
[**DnsGet**](NameServicesApi.md#DnsGet) | **Get** /name-services/dns/{svm.uuid} | 
[**DnsModify**](NameServicesApi.md#DnsModify) | **Patch** /name-services/dns/{svm.uuid} | 
[**LdapCollectionGet**](NameServicesApi.md#LdapCollectionGet) | **Get** /name-services/ldap | 
[**LdapCreate**](NameServicesApi.md#LdapCreate) | **Post** /name-services/ldap | 
[**LdapDelete**](NameServicesApi.md#LdapDelete) | **Delete** /name-services/ldap/{svm.uuid} | 
[**LdapGet**](NameServicesApi.md#LdapGet) | **Get** /name-services/ldap/{svm.uuid} | 
[**LdapModify**](NameServicesApi.md#LdapModify) | **Patch** /name-services/ldap/{svm.uuid} | 
[**NameMappingCollectionGet**](NameServicesApi.md#NameMappingCollectionGet) | **Get** /name-services/name-mappings | 
[**NameMappingCreate**](NameServicesApi.md#NameMappingCreate) | **Post** /name-services/name-mappings | 
[**NameMappingDelete**](NameServicesApi.md#NameMappingDelete) | **Delete** /name-services/name-mappings/{svm.uuid}/{direction}/{index} | 
[**NameMappingModify**](NameServicesApi.md#NameMappingModify) | **Patch** /name-services/name-mappings/{svm.uuid}/{direction}/{index} | 
[**NameMappingPositionGet**](NameServicesApi.md#NameMappingPositionGet) | **Get** /name-services/name-mappings/{svm.uuid}/{direction}/{index} | 
[**NisCollectionGet**](NameServicesApi.md#NisCollectionGet) | **Get** /name-services/nis | 
[**NisCreate**](NameServicesApi.md#NisCreate) | **Post** /name-services/nis | 
[**NisDelete**](NameServicesApi.md#NisDelete) | **Delete** /name-services/nis/{svm.uuid} | 
[**NisGet**](NameServicesApi.md#NisGet) | **Get** /name-services/nis/{svm.uuid} | 
[**NisModify**](NameServicesApi.md#NisModify) | **Patch** /name-services/nis/{svm.uuid} | 



## DnsCollectionGet

> DnsResponse DnsCollectionGet(ctx, optional)



Retrieves the DNS configurations of all data SVMs. DNS configuration for the cluster is retrieved and managed via [`/api/cluster`](#docs-cluster-cluster). ### Related ONTAP commands * `vserver services name-service dns show` * `vserver services name-service dns check` ### Learn more * [`DOC /name-services/dns`](#docs-name-services-name-services_dns) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DnsCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domains** | **optional.String**| Filter by domains | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **servers** | **optional.String**| Filter by servers | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**DnsResponse**](dns_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsCreate

> DnsResponse DnsCreate(ctx, optional)



Creates DNS domain and server configurations for an SVM.<br/> ### Important notes - Each SVM can have only one DNS configuration. - The domain name and the servers fields cannot be empty. - IPv6 must be enabled if IPv6 family addresses are specified in the `servers` field. - Configuring more than one DNS server is recommended to avoid a single point of failure. - The DNS server specified using the `servers` field is validated during this operation.<br/> </br> The validation fails in the following scenarios:<br/> 1. The server is not a DNS server. 2. The server does not exist. 3. The server is unreachable.<br/>  ### Learn more * [`DOC /name-services/dns`](#docs-name-services-name-services_dns)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DnsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Dns**](Dns.md)| Info specification | 

### Return type

[**DnsResponse**](dns_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsDelete

> DnsDelete(ctx, svmUuid)



Deletes DNS domain configuration of the specified SVM. ### Related ONTAP commands * `vserver services name-service dns delete` ### Learn more * [`DOC /name-services/dns`](#docs-name-services-name-services_dns) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 

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


## DnsGet

> Dns DnsGet(ctx, svmUuid, optional)



Retrieves DNS domain and server configuration of an SVM. By default, both DNS domains and servers are displayed. DNS configuration for the cluster is retrieved and managed via [`/api/cluster`](#docs-cluster-cluster). ### Related ONTAP commands * `vserver services name-service dns show` * `vserver services name-service dns check` ### Learn more * [`DOC /name-services/dns`](#docs-name-services-name-services_dns) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***DnsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Dns**](dns.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsModify

> DnsModify(ctx, svmUuid, optional)



Updates DNS domain and server configurations of an SVM. ### Important notes - Both DNS domains and servers can be modified. - The domains and servers fields cannot be empty. - IPv6 must be enabled if IPv6 family addresses are specified for the `servers` field. - The DNS server specified using the `servers` field is validated during this operation.<br/> The validation fails in the following scenarios:<br/> 1. The server is not a DNS server. 2. The server does not exist. 3. The server is unreachable.<br/>  ### Learn more * [`DOC /name-services/dns`](#docs-name-services-name-services_dns)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***DnsModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DnsModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of Dns**](Dns.md)| Info specification | 

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


## LdapCollectionGet

> LdapServiceResponse LdapCollectionGet(ctx, optional)



Retrieves the LDAP configurations for all SVMs.  ### Learn more * [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LdapCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**LdapServiceResponse**](ldap_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapCreate

> LdapServiceResponse LdapCreate(ctx, optional)



Creates an LDAP configuration for an SVM. ### Important notes * Each SVM can have one LDAP configuration. * The LDAP servers and Active Directory domain are mutually exclusive fields. These fields cannot be empty. At any point in time, either the LDAP servers or Active Directory domain must be populated. * IPv6 must be enabled if IPv6 family addresses are specified.</br> #### The following parameters are optional: - preferred AD servers - schema - port - min_bind_level - bind_password - base_scope - use_start_tls - session_security</br> Configuring more than one LDAP server is recommended to avoid a single point of failure. Both FQDNs and IP addresses are supported for the \"servers\" field. The Acitve Directory domain or LDAP servers are validated as part of this operation.</br> LDAP validation fails in the following scenarios:<br/> 1. The server does not have LDAP installed. 2. The server or Active Directory domain is invalid. 3. The server or Active Directory domain is unreachable.<br/>  ### Learn more * [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LdapCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of LdapService**](LdapService.md)| Info specification | 

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


## LdapDelete

> LdapDelete(ctx, svmUuid)



Deletes the LDAP configuration of the specified SVM. LDAP can be removed as a source from the ns-switch if LDAP is not used as a source for lookups.  ### Learn more * [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 

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


## LdapGet

> LdapService LdapGet(ctx, svmUuid, optional)



Retrieves LDAP configuration for an SVM. All parameters for the LDAP configuration are displayed by default.  ### Learn more * [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***LdapGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LdapGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**LdapService**](ldap_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapModify

> LdapModify(ctx, svmUuid, optional)



Updates an LDAP configuration of an SVM. ### Important notes * Both mandatory and optional parameters of the LDAP configuration can be updated. * The LDAP servers and Active Directory domain are mutually exclusive fields. These fields cannot be empty. At any point in time, either the LDAP servers or Active Directory domain must be populated. * IPv6 must be enabled if IPv6 family addresses are specified.<br/> </br>Configuring more than one LDAP server is recommended to avoid a sinlge point of failure. Both FQDNs and IP addresses are supported for the \"servers\" field. The Active Directory domain or LDAP servers are validated as part of this operation.<br/> LDAP validation fails in the following scenarios:<br/> 1. The server does not have LDAP installed. 2. The server or Active Directory domain is invalid. 3. The server or Active Directory domain is unreachable<br/>  ### Learn more * [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***LdapModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LdapModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of LdapService**](LdapService.md)| Info specification | 

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


## NameMappingCollectionGet

> NameMappingResponse NameMappingCollectionGet(ctx, optional)



Retrieves the name mapping configuration for all SVMs. ### Related ONTAP commands * `vserver name-mapping show` ### Learn more * [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NameMappingCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NameMappingCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **clientMatch** | **optional.String**| Filter by client_match | 
 **replacement** | **optional.String**| Filter by replacement | 
 **direction** | **optional.String**| Filter by direction | 
 **pattern** | **optional.String**| Filter by pattern | 
 **index** | **optional.Int32**| Filter by index | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NameMappingResponse**](name_mapping_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NameMappingCreate

> NameMappingResponse NameMappingCreate(ctx, optional)



Creates name mappings for an SVM. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the name mapping. * `index` - Name mapping's position in the priority list. * `direction` - Direction of the name mapping. * `pattern` - Pattern to match to. Maximum length is 256 characters. * `replacement` - Replacement pattern to match to. Maximum length is 256 characters. ### Recommended optional properties * `client_match` - Hostname or IP address added to match the pattern to the client's workstation IP address. ### Related ONTAP commands * `vserver name-mapping create` * `vserver name-mapping insert` ### Learn more * [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NameMappingCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NameMappingCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of NameMapping**](NameMapping.md)| Info specification | 

### Return type

[**NameMappingResponse**](name_mapping_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NameMappingDelete

> NameMappingDelete(ctx, svmUuid, direction, index)



Deletes the name mapping configuration. ### Related ONTAP commands * `vserver name-mapping delete` ### Learn more * [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**direction** | **string**| Direction | 
**index** | **int32**| Position of the entry in the list | 

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


## NameMappingModify

> NameMappingModify(ctx, svmUuid, direction, index, optional)



Updates the name mapping configuration of an SVM. ### Related ONTAP commands * `vserver name-mapping insert` * `vserver name-mapping modify` ### Learn more * [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**direction** | **string**| Direction | 
**index** | **int32**| Position of the entry in the list | 
 **optional** | ***NameMappingModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NameMappingModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **info** | [**optional.Interface of NameMapping**](NameMapping.md)| Info specification | 

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


## NameMappingPositionGet

> NameMapping NameMappingPositionGet(ctx, svmUuid, direction, index, optional)



Retrieves the name mapping configuration of an SVM. ### Related ONTAP commands * `vserver name-mapping show` ### Learn more * [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**direction** | **string**| Direction | 
**index** | **int32**| Position of the entry in the list | 
 **optional** | ***NameMappingPositionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NameMappingPositionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NameMapping**](name_mapping.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NisCollectionGet

> NisServiceResponse NisCollectionGet(ctx, optional)



Retrieves NIS domain configurations of all the SVMs. The bound_servers field indicates the successfully bound NIS servers. Lookups and authentications fail if there are no bound servers. ### Related ONTAP commands * `vserver services name-service nis-domain show` * `vserver services name-service nis-domain show-bound` ### Learn more * [`DOC /name-services/nis`](#docs-name-services-name-services_nis) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NisCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NisCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NisServiceResponse**](nis_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NisCreate

> NisServiceResponse NisCreate(ctx, optional)



Creates an NIS domain and server confguration for a data SVM. NIS configuration for the cluster is managed via [`/api/security/authentication/cluster/nis`](#docs-security-security_authentication_cluster_nis).<br/> ### Important notes   - Each SVM can have one NIS domain configuration.   - Multiple SVMs can be configured with the same NIS domain. Specify the NIS domain and NIS servers as input.Domain name and servers fields cannot be empty.   - Both FQDNs and IP addresses are supported for the servers field.   - IPv6 must be enabled if IPv6 family addresses are specified in the servers field.   - A maximum of ten NIS servers are supported. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the NIS configuration. * `domain` - NIS domain to which the configuration belongs. * `servers` - List of NIS server IP addresses. ### Related ONTAP commands * `vserver services name-service nis-domain create` ### Learn more * [`DOC /name-services/nis`](#docs-name-services-name-services_nis) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NisCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NisCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of NisService**](NisService.md)| Info specification | 

### Return type

[**NisServiceResponse**](nis_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NisDelete

> NisDelete(ctx, svmUuid)



Deletes the NIS domain configuration of an SVM. NIS can be removed as a source from ns-switch if NIS is not used for lookups. ### Related ONTAP commands * `vserver services name-service nis-domain delete` ### Learn more * [`DOC /name-services/nis`](#docs-name-services-name-services_nis) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 

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


## NisGet

> NisService NisGet(ctx, svmUuid, optional)



Retrieves NIS domain and server configurations of an SVM. Both NIS domain and servers are displayed by default. The bound_servers field indicates the successfully bound NIS servers. ### Related ONTAP commands * `vserver services name-service nis-domain show` * `vserver services name-service nis-domain show-bound` ### Learn more * [`DOC /name-services/nis`](#docs-name-services-name-services_nis) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***NisGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NisGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NisService**](nis_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NisModify

> NisModify(ctx, svmUuid, optional)



Updates NIS domain and server configuration of an SVM.<br/> ### Important notes   - Both NIS domain and servers can be modified.   - Domains and servers cannot be empty.   - Both FQDNs and IP addresses are supported for the servers field.   - If the domain is modified, NIS servers must also be specified.   - IPv6 must be enabled if IPv6 family addresses are specified for the servers field. ### Related ONTAP commands * `vserver services name-service nis-domain modify` ### Learn more * [`DOC /name-services/nis`](#docs-name-services-name-services_nis) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***NisModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NisModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of NisService**](NisService.md)| Info specification | 

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

