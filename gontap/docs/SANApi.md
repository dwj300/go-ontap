# \SANApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FcLoginCollectionGet**](SANApi.md#FcLoginCollectionGet) | **Get** /network/fc/logins | 
[**FcLoginGet**](SANApi.md#FcLoginGet) | **Get** /network/fc/logins/{interface.uuid}/{initiator.wwpn} | 
[**FcpServiceCollectionGet**](SANApi.md#FcpServiceCollectionGet) | **Get** /protocols/san/fcp/services | 
[**FcpServiceCreate**](SANApi.md#FcpServiceCreate) | **Post** /protocols/san/fcp/services | 
[**FcpServiceDelete**](SANApi.md#FcpServiceDelete) | **Delete** /protocols/san/fcp/services/{svm.uuid} | 
[**FcpServiceGet**](SANApi.md#FcpServiceGet) | **Get** /protocols/san/fcp/services/{svm.uuid} | 
[**FcpServiceModify**](SANApi.md#FcpServiceModify) | **Patch** /protocols/san/fcp/services/{svm.uuid} | 
[**IgroupCollectionGet**](SANApi.md#IgroupCollectionGet) | **Get** /protocols/san/igroups | 
[**IgroupCreate**](SANApi.md#IgroupCreate) | **Post** /protocols/san/igroups | 
[**IgroupDelete**](SANApi.md#IgroupDelete) | **Delete** /protocols/san/igroups/{uuid} | 
[**IgroupGet**](SANApi.md#IgroupGet) | **Get** /protocols/san/igroups/{uuid} | 
[**IgroupInitiatorCollectionGet**](SANApi.md#IgroupInitiatorCollectionGet) | **Get** /protocols/san/igroups/{igroup.uuid}/initiators | 
[**IgroupInitiatorCreate**](SANApi.md#IgroupInitiatorCreate) | **Post** /protocols/san/igroups/{igroup.uuid}/initiators | 
[**IgroupInitiatorDelete**](SANApi.md#IgroupInitiatorDelete) | **Delete** /protocols/san/igroups/{igroup.uuid}/initiators/{name} | 
[**IgroupInitiatorGet**](SANApi.md#IgroupInitiatorGet) | **Get** /protocols/san/igroups/{igroup.uuid}/initiators/{name} | 
[**IgroupModify**](SANApi.md#IgroupModify) | **Patch** /protocols/san/igroups/{uuid} | 
[**IscsiCredentialsCollectionGet**](SANApi.md#IscsiCredentialsCollectionGet) | **Get** /protocols/san/iscsi/credentials | 
[**IscsiCredentialsCreate**](SANApi.md#IscsiCredentialsCreate) | **Post** /protocols/san/iscsi/credentials | 
[**IscsiCredentialsDelete**](SANApi.md#IscsiCredentialsDelete) | **Delete** /protocols/san/iscsi/credentials/{svm.uuid}/{initiator} | 
[**IscsiCredentialsGet**](SANApi.md#IscsiCredentialsGet) | **Get** /protocols/san/iscsi/credentials/{svm.uuid}/{initiator} | 
[**IscsiCredentialsModify**](SANApi.md#IscsiCredentialsModify) | **Patch** /protocols/san/iscsi/credentials/{svm.uuid}/{initiator} | 
[**IscsiServiceCollectionGet**](SANApi.md#IscsiServiceCollectionGet) | **Get** /protocols/san/iscsi/services | 
[**IscsiServiceCreate**](SANApi.md#IscsiServiceCreate) | **Post** /protocols/san/iscsi/services | 
[**IscsiServiceDelete**](SANApi.md#IscsiServiceDelete) | **Delete** /protocols/san/iscsi/services/{svm.uuid} | 
[**IscsiServiceGet**](SANApi.md#IscsiServiceGet) | **Get** /protocols/san/iscsi/services/{svm.uuid} | 
[**IscsiServiceModify**](SANApi.md#IscsiServiceModify) | **Patch** /protocols/san/iscsi/services/{svm.uuid} | 
[**IscsiSessionCollectionGet**](SANApi.md#IscsiSessionCollectionGet) | **Get** /protocols/san/iscsi/sessions | 
[**IscsiSessionGet**](SANApi.md#IscsiSessionGet) | **Get** /protocols/san/iscsi/sessions/{svm.uuid}/{tpgroup}/{tsih} | 
[**LunCollectionGet**](SANApi.md#LunCollectionGet) | **Get** /storage/luns | 
[**LunCreate**](SANApi.md#LunCreate) | **Post** /storage/luns | 
[**LunDelete**](SANApi.md#LunDelete) | **Delete** /storage/luns/{uuid} | 
[**LunGet**](SANApi.md#LunGet) | **Get** /storage/luns/{uuid} | 
[**LunMapCollectionGet**](SANApi.md#LunMapCollectionGet) | **Get** /protocols/san/lun-maps | 
[**LunMapCreate**](SANApi.md#LunMapCreate) | **Post** /protocols/san/lun-maps | 
[**LunMapDelete**](SANApi.md#LunMapDelete) | **Delete** /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid} | 
[**LunMapGet**](SANApi.md#LunMapGet) | **Get** /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid} | 
[**LunModify**](SANApi.md#LunModify) | **Patch** /storage/luns/{uuid} | 
[**WwpnAliasCollectionGet**](SANApi.md#WwpnAliasCollectionGet) | **Get** /network/fc/wwpn-aliases | 
[**WwpnAliasCreate**](SANApi.md#WwpnAliasCreate) | **Post** /network/fc/wwpn-aliases | 
[**WwpnAliasDelete**](SANApi.md#WwpnAliasDelete) | **Delete** /network/fc/wwpn-aliases/{svm.uuid}/{alias} | 
[**WwpnAliasGet**](SANApi.md#WwpnAliasGet) | **Get** /network/fc/wwpn-aliases/{svm.uuid}/{alias} | 


# **FcLoginCollectionGet**
> FcLoginResponse FcLoginCollectionGet(ctx, optional)


Retrieves FC logins. ### Related ONTAP commands * `vserver fcp initiator show` ### Learn more * SAN: [`DOC /network/fc/logins`](#docs-SAN-network_fc_logins) * NVMe: [`DOC /network/fc/logins`](#docs-NVMe-network_fc_logins) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FcLoginCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FcLoginCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protocol** | **optional.String**| Filter by protocol | 
 **initiatorPortAddress** | **optional.String**| Filter by initiator.port_address | 
 **initiatorWwpn** | **optional.String**| Filter by initiator.wwpn | 
 **initiatorWwnn** | **optional.String**| Filter by initiator.wwnn | 
 **initiatorAliases** | **optional.String**| Filter by initiator.aliases | 
 **igroupsUuid** | **optional.String**| Filter by igroups.uuid | 
 **igroupsName** | **optional.String**| Filter by igroups.name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **interfaceName** | **optional.String**| Filter by interface.name | 
 **interfaceWwpn** | **optional.String**| Filter by interface.wwpn | 
 **interfaceUuid** | **optional.String**| Filter by interface.uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FcLoginResponse**](fc_login_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FcLoginGet**
> FcLogin FcLoginGet(ctx, interfaceUuid, initiatorWwpn, optional)


Retrieves an FC login. ### Related ONTAP commands * `vserver fcp initiator show` ### Learn more * SAN: [`DOC /network/fc/logins`](#docs-SAN-network_fc_logins) * NVMe: [`DOC /network/fc/logins`](#docs-NVMe-network_fc_logins) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceUuid** | **string**| The unique identifier of the FC interface through which the initiator logged in.  | 
  **initiatorWwpn** | **string**| The world wide port name (WWPN) of the initiator.  | 
 **optional** | ***FcLoginGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FcLoginGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FcLogin**](fc_login.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FcpServiceCollectionGet**
> FcpServiceResponse FcpServiceCollectionGet(ctx, optional)


Retrieves FC Protocol services. ### Related ONTAP commands * `vserver fcp show` ### Learn more * [`DOC /protocols/san/fcp/services`](#docs-SAN-protocols_san_fcp_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FcpServiceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FcpServiceCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **targetName** | **optional.String**| Filter by target.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FcpServiceResponse**](fcp_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FcpServiceCreate**
> FcpServiceResponse FcpServiceCreate(ctx, info)


Creates an FC Protocol service. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the FC service. ### Related ONTAP commands * `vserver fcp create` ### Learn more * [`DOC /protocols/san/fcp/services`](#docs-SAN-protocols_san_fcp_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**FcpService**](FcpService.md)| The property values for the new FC Protocol service.  | 

### Return type

[**FcpServiceResponse**](fcp_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FcpServiceDelete**
> FcpServiceDelete(ctx, svmUuid)


Deletes an FC Protocol service. An FC Protocol service must be disabled before it can be deleted. ### Related ONTAP commands * `vserver fcp delete` ### Learn more * [`DOC /protocols/san/fcp/services`](#docs-SAN-protocols_san_fcp_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM for which to delete the FC Protocol service.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FcpServiceGet**
> FcpService FcpServiceGet(ctx, svmUuid, optional)


Retrieves a Fibre Channel Protocol service. ### Related ONTAP commands * `vserver fcp show` ### Learn more * [`DOC /protocols/san/fcp/services`](#docs-SAN-protocols_san_fcp_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM for which to retrieve the FC Protocol service.  | 
 **optional** | ***FcpServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FcpServiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FcpService**](fcp_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FcpServiceModify**
> FcpServiceModify(ctx, svmUuid, info)


Updates an FC Protocol service. ### Related ONTAP commands * `vserver fcp modify` * `vserver fcp start` * `vserver fcp stop` ### Learn more * [`DOC /protocols/san/fcp/services`](#docs-SAN-protocols_san_fcp_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM whose FC Protocol service is to be patched.  | 
  **info** | [**FcpService**](FcpService.md)| The new property values for the FC Protocol service.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupCollectionGet**
> IgroupResponse IgroupCollectionGet(ctx, optional)


Retrieves initiator groups. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `lun_maps.*` ### Related ONTAP commands * `lun igroup show` * `lun mapping show` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IgroupCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgroupCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteOnUnmap** | **optional.Bool**| Filter by delete_on_unmap | 
 **protocol** | **optional.String**| Filter by protocol | 
 **lunMapsLogicalUnitNumber** | **optional.Int32**| Filter by lun_maps.logical_unit_number | 
 **lunMapsLunNodeUuid** | **optional.String**| Filter by lun_maps.lun.node.uuid | 
 **lunMapsLunNodeName** | **optional.String**| Filter by lun_maps.lun.node.name | 
 **lunMapsLunUuid** | **optional.String**| Filter by lun_maps.lun.uuid | 
 **lunMapsLunName** | **optional.String**| Filter by lun_maps.lun.name | 
 **initiatorsIgroupUuid** | **optional.String**| Filter by initiators.igroup.uuid | 
 **initiatorsName** | **optional.String**| Filter by initiators.name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **osType** | **optional.String**| Filter by os_type | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IgroupResponse**](igroup_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupCreate**
> IgroupResponse IgroupCreate(ctx, info)


Creates an initiator group. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the initiator group. * `name` - Name of the initiator group. * `os_type` - Operating system of the initiator group's initiators. ### Recommended optional properties * `initiators.name` - Name(s) of initiator group's initiators. This property can be used to create the initiator group and populate it with initiators in a single request. ### Default property values If not specified in POST, the following default property values are assigned. * `protocol` - _mixed_ - Data protocol of the initiator group's initiators. ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**Igroup**](Igroup.md)| The property values for the new initiator group.  | 

### Return type

[**IgroupResponse**](igroup_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupDelete**
> IgroupDelete(ctx, uuid, optional)


Deletes an initiator group. ### Related ONTAP commands * `lun igroup delete` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the initiator group.  | 
 **optional** | ***IgroupDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgroupDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowDeleteWhileMapped** | **optional.Bool**| Allow deletion of a mapped initiator group.&lt;br/&gt; Deleting a mapped initiator group makes the LUNs to which the initiator group is mapped no longer available. This might cause a disruption in the availability of data.&lt;br/&gt; &lt;b&gt;This parameter should be used with caution.&lt;/b&gt;  | [default to false]

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupGet**
> Igroup IgroupGet(ctx, uuid, optional)


Retrieves an initiator group. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `lun_maps.*` ### Related ONTAP commands * `lun igroup show` * `lun mapping show` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the initiator group.  | 
 **optional** | ***IgroupGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgroupGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Igroup**](igroup.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupInitiatorCollectionGet**
> IgroupInitiatorResponse IgroupInitiatorCollectionGet(ctx, igroupUuid, optional)


Retrieves initiators of an initiator group. ### Related ONTAP commands * `lun igroup show` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **igroupUuid** | **string**| The unique identifier of the initiator group.  | 
 **optional** | ***IgroupInitiatorCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgroupInitiatorCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IgroupInitiatorResponse**](igroup_initiator_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupInitiatorCreate**
> IgroupInitiatorResponse IgroupInitiatorCreate(ctx, igroupUuid, info)


Adds one or more initiators to an initiator group. ### Required properties * `name` or `records.name` - Initiator name(s) to add to the initiator group. ### Related ONTAP commands * `lun igroup add` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **igroupUuid** | **string**| The unique identifier of the initiator group.  | 
  **info** | [**IgroupInitiator**](IgroupInitiator.md)| The properties of the initiator to add to the initiator group.  | 

### Return type

[**IgroupInitiatorResponse**](igroup_initiator_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupInitiatorDelete**
> IgroupInitiatorDelete(ctx, igroupUuid, name, optional)


Deletes an initiator from an initiator group. ### Related ONTAP commands * `lun igroup remove` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **igroupUuid** | **string**| The unique identifier of the initiator group.  | 
  **name** | **string**| The initiator name.  | 
 **optional** | ***IgroupInitiatorDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgroupInitiatorDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allowDeleteWhileMapped** | **optional.Bool**| Allow deletion of an initiator from of a mapped initiator group.&lt;br/&gt; Deleting an initiator from a mapped initiator group makes the LUNs to which the initiator group is mapped no longer available to the initiator. This might cause a disruption in the availability of data.&lt;br/&gt; &lt;b&gt;This parameter should be used with caution.&lt;/b&gt;  | [default to false]

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupInitiatorGet**
> IgroupInitiator IgroupInitiatorGet(ctx, igroupUuid, name, optional)


Retrieves an initiator of an initiator group. ### Related ONTAP commands * `lun igroup show` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **igroupUuid** | **string**| The unique identifier of the initiator group.  | 
  **name** | **string**| The initiator name.  | 
 **optional** | ***IgroupInitiatorGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgroupInitiatorGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**IgroupInitiator**](igroup_initiator.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IgroupModify**
> IgroupModify(ctx, uuid, info)


Updates an initiator group. ### Related ONTAP commands * `lun igroup modify` * `lun igroup rename` ### Learn more * [`DOC /protocols/san/igroups`](#docs-SAN-protocols_san_igroups) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the initiator group.  | 
  **info** | [**Igroup**](Igroup.md)| The new property values for the initiator group.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiCredentialsCollectionGet**
> IscsiCredentialsResponse IscsiCredentialsCollectionGet(ctx, optional)


Retrieves iSCSI credentials. ### Related ONTAP commands * `vserver iscsi security show` ### Learn more * [`DOC /protocols/san/iscsi/credentials`](#docs-SAN-protocols_san_iscsi_credentials) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IscsiCredentialsCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiCredentialsCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initiatorAddressMasksAddress** | **optional.String**| Filter by initiator_address.masks.address | 
 **initiatorAddressMasksNetmask** | **optional.String**| Filter by initiator_address.masks.netmask | 
 **initiatorAddressMasksFamily** | **optional.String**| Filter by initiator_address.masks.family | 
 **initiatorAddressRangesFamily** | **optional.String**| Filter by initiator_address.ranges.family | 
 **initiatorAddressRangesStart** | **optional.String**| Filter by initiator_address.ranges.start | 
 **initiatorAddressRangesEnd** | **optional.String**| Filter by initiator_address.ranges.end | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **initiator** | **optional.String**| Filter by initiator | 
 **authenticationType** | **optional.String**| Filter by authentication_type | 
 **chapInboundUser** | **optional.String**| Filter by chap.inbound.user | 
 **chapOutboundUser** | **optional.String**| Filter by chap.outbound.user | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IscsiCredentialsResponse**](iscsi_credentials_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiCredentialsCreate**
> IscsiCredentialsResponse IscsiCredentialsCreate(ctx, info)


Creates iSCSI credentials. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the iSCSI credentials. * `initiator` - Initiator for which the iSCSI credentials are to be created. * `authentication_type` - Type of authentication to use for the credentials. ### Recommended optional properties * `chap.inbound.user` - In-bound CHAP authentication user name. * `chap.inbound.password` - In-bound CHAP authentication password. * `chap.outbound.user` - Out-bound CHAP authentication user name. * `chap.outbound.password` - Out-bound CHAP authentication password. ### Related ONTAP commands * `vserver iscsi security create` ### Learn more * [`DOC /protocols/san/iscsi/credentials`](#docs-SAN-protocols_san_iscsi_credentials) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**IscsiCredentials**](IscsiCredentials.md)| The property values for the new iSCSI credentials object.  | 

### Return type

[**IscsiCredentialsResponse**](iscsi_credentials_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiCredentialsDelete**
> IscsiCredentialsDelete(ctx, svmUuid, initiator)


Deletes specified iSCSI credentials. ### Related ONTAP commands * `vserver iscsi security delete` ### Learn more * [`DOC /protocols/san/iscsi/credentials`](#docs-SAN-protocols_san_iscsi_credentials) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of an SVM.  | 
  **initiator** | **string**| The iSCSI initiator of the credentials object.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiCredentialsGet**
> IscsiCredentials IscsiCredentialsGet(ctx, svmUuid, initiator, optional)


Retrieves specified iSCSI credentials. ### Related ONTAP commands * `vserver iscsi security show` ### Learn more * [`DOC /protocols/san/iscsi/credentials`](#docs-SAN-protocols_san_iscsi_credentials) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of an SVM.  | 
  **initiator** | **string**| The iSCSI initiator of the credentials object.  | 
 **optional** | ***IscsiCredentialsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiCredentialsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**IscsiCredentials**](iscsi_credentials.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiCredentialsModify**
> IscsiCredentialsModify(ctx, svmUuid, initiator, info, optional)


Updates specified iSCSI credentials. ### Related ONTAP commands * `vserver iscsi security add-initiator-address-ranges` * `vserver iscsi security default` * `vserver iscsi security modify` * `vserver iscsi security remove-initiator-address-ranges` ### Learn more * [`DOC /protocols/san/iscsi/credentials`](#docs-SAN-protocols_san_iscsi_credentials) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of an SVM.  | 
  **initiator** | **string**| The iSCSI initiator of the credentials object.  | 
  **info** | [**IscsiCredentials**](IscsiCredentials.md)| The new property values for the iSCSI credentials object.  | 
 **optional** | ***IscsiCredentialsModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiCredentialsModifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addInitiatorAddresses** | **optional.Bool**| If _true_, the initiator addresses in the body merge into the existing addresses in the iSCSI security object rather than replace the existing addresses.  | [default to false]
 **removeInitiatorAddresses** | **optional.Bool**| If _true_, the initiator addresses in the body are removed from the existing addresses in the iSCSI security object rather than replace the existing addresses.  | [default to false]

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiServiceCollectionGet**
> IscsiServiceResponse IscsiServiceCollectionGet(ctx, optional)


Retrieves iSCSI services. ### Related ONTAP commands * `vserver iscsi show` ### Learn more * [`DOC /protocols/san/iscsi/services`](#docs-SAN-protocols_san_iscsi_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IscsiServiceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiServiceCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **targetName** | **optional.String**| Filter by target.name | 
 **targetAlias** | **optional.String**| Filter by target.alias | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IscsiServiceResponse**](iscsi_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiServiceCreate**
> IscsiServiceResponse IscsiServiceCreate(ctx, info)


Creates an iSCSI service. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the iSCSI service. ### Related ONTAP commands * `vserver iscsi create` ### Learn more * [`DOC /protocols/san/iscsi/services`](#docs-SAN-protocols_san_iscsi_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**IscsiService**](IscsiService.md)| The property values for the new iSCSI service.  | 

### Return type

[**IscsiServiceResponse**](iscsi_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiServiceDelete**
> IscsiServiceDelete(ctx, svmUuid)


Deletes an iSCSI service. An iSCSI service must be disabled before it can be deleted. ### Related ONTAP commands * `vserver iscsi delete` ### Learn more * [`DOC /protocols/san/iscsi/services`](#docs-SAN-protocols_san_iscsi_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM for which to delete the iSCSI service.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiServiceGet**
> IscsiService IscsiServiceGet(ctx, svmUuid, optional)


Retrieves an iSCSI service. ### Related ONTAP commands * `vserver iscsi show` ### Learn more * [`DOC /protocols/san/iscsi/services`](#docs-SAN-protocols_san_iscsi_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM for which to retrieve the iSCSI service.  | 
 **optional** | ***IscsiServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiServiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**IscsiService**](iscsi_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiServiceModify**
> IscsiServiceModify(ctx, svmUuid, info)


Updates an iSCSI service. ### Related ONTAP commands * `vserver iscsi modify` * `vserver iscsi start` * `vserver iscsi stop` ### Learn more * [`DOC /protocols/san/iscsi/services`](#docs-SAN-protocols_san_iscsi_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM for which to update the iSCSI service.  | 
  **info** | [**IscsiService**](IscsiService.md)| The new property values for the iSCSI service.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiSessionCollectionGet**
> IscsiSessionResponse IscsiSessionCollectionGet(ctx, optional)


Retrieves iSCSI sessions. ### Related ONTAP commands * `vserver iscsi connection show` * `vserver iscsi session parameter show` * `vserver iscsi session show` ### Learn more * [`DOC /protocols/san/iscsi/sessions`](#docs-SAN-protocols_san_iscsi_sessions) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IscsiSessionCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiSessionCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **igroupsUuid** | **optional.String**| Filter by igroups.uuid | 
 **igroupsName** | **optional.String**| Filter by igroups.name | 
 **connectionsAuthenticationType** | **optional.String**| Filter by connections.authentication_type | 
 **connectionsCid** | **optional.Int32**| Filter by connections.cid | 
 **connectionsInterfaceIpAddress** | **optional.String**| Filter by connections.interface.ip.address | 
 **connectionsInterfaceIpPort** | **optional.Int32**| Filter by connections.interface.ip.port | 
 **connectionsInterfaceUuid** | **optional.String**| Filter by connections.interface.uuid | 
 **connectionsInterfaceName** | **optional.String**| Filter by connections.interface.name | 
 **connectionsInitiatorAddressPort** | **optional.Int32**| Filter by connections.initiator_address.port | 
 **connectionsInitiatorAddressAddress** | **optional.String**| Filter by connections.initiator_address.address | 
 **targetPortalGroup** | **optional.String**| Filter by target_portal_group | 
 **tsih** | **optional.Int32**| Filter by tsih | 
 **initiatorName** | **optional.String**| Filter by initiator.name | 
 **initiatorAlias** | **optional.String**| Filter by initiator.alias | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **isid** | **optional.String**| Filter by isid | 
 **targetPortalGroupTag** | **optional.Int32**| Filter by target_portal_group_tag | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IscsiSessionResponse**](iscsi_session_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IscsiSessionGet**
> IscsiSession IscsiSessionGet(ctx, svmUuid, tpgroup, tsih, optional)


Retrieves an iSCSI session. ### Related ONTAP commands * `vserver iscsi connection show` * `vserver iscsi session parameter show` * `vserver iscsi session show` ### Learn more * [`DOC /protocols/san/iscsi/sessions`](#docs-SAN-protocols_san_iscsi_sessions) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM of the iSCSI session.  | 
  **tpgroup** | **string**| The target portal group of the iSCSI session.  | 
  **tsih** | **int32**| The target session identifying handle.  | 
 **optional** | ***IscsiSessionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IscsiSessionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**IscsiSession**](iscsi_session.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunCollectionGet**
> LunResponse LunCollectionGet(ctx, optional)


Retrieves LUNs. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `auto_delete` * `lun_maps.*` * `movement.*` * `status.mapped` ### Related ONTAP commands * `lun mapping show` * `lun move show` * `lun show` * `volume file clone show-autodelete` ### Learn more * [`DOC /storage/luns`](#docs-SAN-storage_luns) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LunCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LunCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationLogicalUnit** | **optional.String**| Filter by location.logical_unit | 
 **locationQtreeName** | **optional.String**| Filter by location.qtree.name | 
 **locationQtreeId** | **optional.Int32**| Filter by location.qtree.id | 
 **locationVolumeName** | **optional.String**| Filter by location.volume.name | 
 **locationVolumeUuid** | **optional.String**| Filter by location.volume.uuid | 
 **statusContainerState** | **optional.String**| Filter by status.container_state | 
 **statusMapped** | **optional.Bool**| Filter by status.mapped | 
 **statusState** | **optional.String**| Filter by status.state | 
 **statusReadOnly** | **optional.Bool**| Filter by status.read_only | 
 **qosPolicyUuid** | **optional.String**| Filter by qos_policy.uuid | 
 **qosPolicyName** | **optional.String**| Filter by qos_policy.name | 
 **class** | **optional.String**| Filter by class | 
 **movementMaxThroughput** | **optional.String**| Filter by movement.max_throughput | 
 **movementPathsSource** | **optional.String**| Filter by movement.paths.source | 
 **movementPathsDestination** | **optional.String**| Filter by movement.paths.destination | 
 **movementProgressState** | **optional.String**| Filter by movement.progress.state | 
 **movementProgressFailureCode** | **optional.String**| Filter by movement.progress.failure.code | 
 **movementProgressFailureArgumentsMessage** | **optional.String**| Filter by movement.progress.failure.arguments.message | 
 **movementProgressFailureArgumentsCode** | **optional.String**| Filter by movement.progress.failure.arguments.code | 
 **movementProgressFailureMessage** | **optional.String**| Filter by movement.progress.failure.message | 
 **movementProgressFailureTarget** | **optional.String**| Filter by movement.progress.failure.target | 
 **movementProgressPercentComplete** | **optional.Int32**| Filter by movement.progress.percent_complete | 
 **movementProgressVolumeSnapshotBlocked** | **optional.Bool**| Filter by movement.progress.volume_snapshot_blocked | 
 **movementProgressElapsed** | **optional.Int32**| Filter by movement.progress.elapsed | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **lunMapsLogicalUnitNumber** | **optional.Int32**| Filter by lun_maps.logical_unit_number | 
 **lunMapsIgroupUuid** | **optional.String**| Filter by lun_maps.igroup.uuid | 
 **lunMapsIgroupName** | **optional.String**| Filter by lun_maps.igroup.name | 
 **spaceSize** | **optional.Int32**| Filter by space.size | 
 **spaceGuaranteeRequested** | **optional.Bool**| Filter by space.guarantee.requested | 
 **spaceGuaranteeReserved** | **optional.Bool**| Filter by space.guarantee.reserved | 
 **spaceUsed** | **optional.Int32**| Filter by space.used | 
 **osType** | **optional.String**| Filter by os_type | 
 **comment** | **optional.String**| Filter by comment | 
 **autoDelete** | **optional.Bool**| Filter by auto_delete | 
 **uuid** | **optional.String**| Filter by uuid | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**LunResponse**](lun_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunCreate**
> LunResponse LunCreate(ctx, info)


Creates a LUN. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the LUN. * `name`, `location.volume.name` or `location.volume.uuid` - Existing volume in which to create the LUN. * `name` or `location.logical_unit` - Base name of the LUN. * `os_type` - Operating system from which the LUN will be accessed. Required when creating a non-clone LUN and disallowed when creating a clone of an existing LUN. A clone's `os_type` is taken from the source LUN. * `space.size` - Size of the LUN. Required when creating a non-clone LUN and disallowed when creating a clone of an existing LUN. A clone's size is taken from the source LUN. ### Recommended optional properties * `qos_policy.name` or `qos_policy.uuid` - Existing traditional or adaptive QoS policy to be applied to the LUN. All LUNs should be managed by a QoS policy at the volume or LUN level. ### Default property values If not specified in POST, the follow default property values are assigned. * `auto_delete` - _false_ ### Related ONTAP commands * `lun create` * `volume file clone autodelete` * `volume file clone create` ### Learn more * [`DOC /storage/luns`](#docs-SAN-storage_luns) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**Lun**](Lun.md)| The property values for the new LUN.  | 

### Return type

[**LunResponse**](lun_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunDelete**
> LunDelete(ctx, uuid, optional)


Deletes a LUN. ### Related ONTAP commands * `lun delete` ### Learn more * [`DOC /storage/luns`](#docs-SAN-storage_luns) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the LUN.  | 
 **optional** | ***LunDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LunDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowDeleteWhileMapped** | **optional.Bool**| Allow deletion of a mapped LUN. A mapped LUN might be in use. Deleting a mapped LUN also deletes the LUN map and makes the data no longer available. This might cause a disruption in the availability of data. **This parameter should be used with caution.**  | [default to false]

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunGet**
> Lun LunGet(ctx, uuid, optional)


Retrieves a LUN. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `auto_delete` * `lun_maps.*` * `movement.*` * `status.mapped` ### Related ONTAP commands * `lun mapping show` * `lun move show` * `lun show` * `volume file clone show-autodelete` ### Learn more * [`DOC /storage/luns`](#docs-SAN-storage_luns) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the LUN to retrieve.  | 
 **optional** | ***LunGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LunGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Lun**](lun.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunMapCollectionGet**
> LunMapResponse LunMapCollectionGet(ctx, optional)


Retrieves LUN maps. ### Related ONTAP commands * `lun mapping show` * [`DOC /protocols/san/lun-maps`](#docs-SAN-protocols_san_lun-maps)  ### Learn more * [`DOC /protocols/san/lun-maps`](#docs-SAN-protocols_san_lun-maps)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LunMapCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LunMapCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **logicalUnitNumber** | **optional.Int32**| Filter by logical_unit_number | 
 **igroupProtocol** | **optional.String**| Filter by igroup.protocol | 
 **igroupUuid** | **optional.String**| Filter by igroup.uuid | 
 **igroupInitiators** | **optional.String**| Filter by igroup.initiators | 
 **igroupOsType** | **optional.String**| Filter by igroup.os_type | 
 **igroupName** | **optional.String**| Filter by igroup.name | 
 **lunName** | **optional.String**| Filter by lun.name | 
 **lunNodeName** | **optional.String**| Filter by lun.node.name | 
 **lunNodeUuid** | **optional.String**| Filter by lun.node.uuid | 
 **lunUuid** | **optional.String**| Filter by lun.uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**LunMapResponse**](lun_map_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunMapCreate**
> LunMapResponse LunMapCreate(ctx, info)


Creates a LUN map. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the LUN map. * `igroup.uuid` or `igroup.name` - Existing initiator group to map to the specified LUN. * `lun.uuid` or `lun.name` - Existing LUN to map to the specified initiator group. ### Default property values If not specified in POST, the following default property values are assigned. * `logical_unit_number` - If no value is provided, ONTAP assigns the lowest available value. ### Related ONTAP commands * `lun mapping create` ### Learn more * [`DOC /protocols/san/lun-maps`](#docs-SAN-protocols_san_lun-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**LunMap**](LunMap.md)| The property values for the new LUN map.  | 

### Return type

[**LunMapResponse**](lun_map_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunMapDelete**
> LunMapDelete(ctx, lunUuid, igroupUuid)


Deletes a LUN map. ### Related ONTAP commands * `lun mapping delete` ### Learn more * [`DOC /protocols/san/lun-maps`](#docs-SAN-protocols_san_lun-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lunUuid** | **string**| The unique identifier of the LUN.  | 
  **igroupUuid** | **string**| The unique identifier of the igroup.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunMapGet**
> LunMap LunMapGet(ctx, lunUuid, igroupUuid, optional)


Retrieves a LUN map. ### Related ONTAP commands * `lun mapping show` ### Learn more * [`DOC /protocols/san/lun-maps`](#docs-SAN-protocols_san_lun-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lunUuid** | **string**| The unique identifier of the LUN.  | 
  **igroupUuid** | **string**| The unique identifier of the igroup.  | 
 **optional** | ***LunMapGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LunMapGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**LunMap**](lun_map.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LunModify**
> LunModify(ctx, uuid, info)


Updates the properties of a LUN. PATCH can also be be used to overwrite the contents of a LUN as a clone of another, to begin movement of a LUN between volumes, and to pause and resume the movement of a LUN between volumes. ### Related ONTAP commands * `lun modify` * `lun move modify` * `lun move pause` * `lun move resume` * `lun move start` * `lun resize` * `volume file clone autodelete` ### Learn more * [`DOC /storage/luns`](#docs-SAN-storage_luns) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the LUN to update.  | 
  **info** | [**Lun**](Lun.md)| The new property values for the LUN.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WwpnAliasCollectionGet**
> WwpnAliasResponse WwpnAliasCollectionGet(ctx, optional)


Retrieves FC WWPN aliases. ### Related ONTAP commands * `vserver fcp wwpn-alias show` ### Learn more * [`DOC /network/fc/wwpn-aliases`](#docs-SAN-network_fc_wwpn-aliases) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WwpnAliasCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WwpnAliasCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wwpn** | **optional.String**| Filter by wwpn | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **alias** | **optional.String**| Filter by alias | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**WwpnAliasResponse**](wwpn_alias_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WwpnAliasCreate**
> WwpnAliasResponse WwpnAliasCreate(ctx, info)


Creates an FC WWPN alias. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the FC alias. * `alias` - Name of the FC alias. * `wwpn` - FC WWPN for which to create the alias. ### Related ONTAP commands * `vserver fcp wwpn-alias set` ### Learn more * [`DOC /network/fc/wwpn-aliases`](#docs-SAN-network_fc_wwpn-aliases) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**WwpnAlias**](WwpnAlias.md)| The property values for the new WWPN alias.  | 

### Return type

[**WwpnAliasResponse**](wwpn_alias_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WwpnAliasDelete**
> WwpnAliasDelete(ctx, svmUuid, alias)


Deletes an FC WWPN alias. ### Related ONTAP commands * `vserver fcp wwpn-alias remove` ### Learn more * [`DOC /network/fc/wwpn-aliases`](#docs-SAN-network_fc_wwpn-aliases) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM.  | 
  **alias** | **string**| The name of FC WWPN alias.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WwpnAliasGet**
> WwpnAlias WwpnAliasGet(ctx, svmUuid, alias, optional)


Retrieves an FC WWPN alias. ### Related ONTAP commands * `vserver fcp wwpn-alias show` ### Learn more * [`DOC /network/fc/wwpn-aliases`](#docs-SAN-network_fc_wwpn-aliases) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM in which the alias is found.  | 
  **alias** | **string**| The name of FC WWPN alias.  | 
 **optional** | ***WwpnAliasGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WwpnAliasGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**WwpnAlias**](wwpn_alias.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

