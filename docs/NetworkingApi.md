# \NetworkingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FcInterfaceCollectionGet**](NetworkingApi.md#FcInterfaceCollectionGet) | **Get** /network/fc/interfaces | 
[**FcInterfaceCreate**](NetworkingApi.md#FcInterfaceCreate) | **Post** /network/fc/interfaces | 
[**FcInterfaceDelete**](NetworkingApi.md#FcInterfaceDelete) | **Delete** /network/fc/interfaces/{uuid} | 
[**FcInterfaceGet**](NetworkingApi.md#FcInterfaceGet) | **Get** /network/fc/interfaces/{uuid} | 
[**FcInterfaceModify**](NetworkingApi.md#FcInterfaceModify) | **Patch** /network/fc/interfaces/{uuid} | 
[**FcPortCollectionGet**](NetworkingApi.md#FcPortCollectionGet) | **Get** /network/fc/ports | 
[**FcPortGet**](NetworkingApi.md#FcPortGet) | **Get** /network/fc/ports/{uuid} | 
[**FcPortModify**](NetworkingApi.md#FcPortModify) | **Patch** /network/fc/ports/{uuid} | 
[**IpspaceDelete**](NetworkingApi.md#IpspaceDelete) | **Delete** /network/ipspaces/{uuid} | 
[**IpspaceGet**](NetworkingApi.md#IpspaceGet) | **Get** /network/ipspaces/{uuid} | 
[**IpspaceModify**](NetworkingApi.md#IpspaceModify) | **Patch** /network/ipspaces/{uuid} | 
[**IpspacesCreate**](NetworkingApi.md#IpspacesCreate) | **Post** /network/ipspaces | 
[**IpspacesGet**](NetworkingApi.md#IpspacesGet) | **Get** /network/ipspaces | 
[**NetworkEthernetBroadcastDomainDelete**](NetworkingApi.md#NetworkEthernetBroadcastDomainDelete) | **Delete** /network/ethernet/broadcast-domains/{uuid} | 
[**NetworkEthernetBroadcastDomainGet**](NetworkingApi.md#NetworkEthernetBroadcastDomainGet) | **Get** /network/ethernet/broadcast-domains/{uuid} | 
[**NetworkEthernetBroadcastDomainModify**](NetworkingApi.md#NetworkEthernetBroadcastDomainModify) | **Patch** /network/ethernet/broadcast-domains/{uuid} | 
[**NetworkEthernetBroadcastDomainsCreate**](NetworkingApi.md#NetworkEthernetBroadcastDomainsCreate) | **Post** /network/ethernet/broadcast-domains | 
[**NetworkEthernetBroadcastDomainsGet**](NetworkingApi.md#NetworkEthernetBroadcastDomainsGet) | **Get** /network/ethernet/broadcast-domains | 
[**NetworkEthernetPortDelete**](NetworkingApi.md#NetworkEthernetPortDelete) | **Delete** /network/ethernet/ports/{uuid} | 
[**NetworkEthernetPortGet**](NetworkingApi.md#NetworkEthernetPortGet) | **Get** /network/ethernet/ports/{uuid} | 
[**NetworkEthernetPortModify**](NetworkingApi.md#NetworkEthernetPortModify) | **Patch** /network/ethernet/ports/{uuid} | 
[**NetworkEthernetPortsCreate**](NetworkingApi.md#NetworkEthernetPortsCreate) | **Post** /network/ethernet/ports | 
[**NetworkEthernetPortsGet**](NetworkingApi.md#NetworkEthernetPortsGet) | **Get** /network/ethernet/ports | 
[**NetworkIpInterfaceDelete**](NetworkingApi.md#NetworkIpInterfaceDelete) | **Delete** /network/ip/interfaces/{uuid} | 
[**NetworkIpInterfaceGet**](NetworkingApi.md#NetworkIpInterfaceGet) | **Get** /network/ip/interfaces/{uuid} | 
[**NetworkIpInterfaceModify**](NetworkingApi.md#NetworkIpInterfaceModify) | **Patch** /network/ip/interfaces/{uuid} | 
[**NetworkIpInterfacesCreate**](NetworkingApi.md#NetworkIpInterfacesCreate) | **Post** /network/ip/interfaces | 
[**NetworkIpInterfacesGet**](NetworkingApi.md#NetworkIpInterfacesGet) | **Get** /network/ip/interfaces | 
[**NetworkIpRouteDelete**](NetworkingApi.md#NetworkIpRouteDelete) | **Delete** /network/ip/routes/{uuid} | 
[**NetworkIpRouteGet**](NetworkingApi.md#NetworkIpRouteGet) | **Get** /network/ip/routes/{uuid} | 
[**NetworkIpRoutesCreate**](NetworkingApi.md#NetworkIpRoutesCreate) | **Post** /network/ip/routes | 
[**NetworkIpRoutesGet**](NetworkingApi.md#NetworkIpRoutesGet) | **Get** /network/ip/routes | 
[**NetworkIpServicePoliciesGet**](NetworkingApi.md#NetworkIpServicePoliciesGet) | **Get** /network/ip/service-policies | 
[**NetworkIpServicePolicyGet**](NetworkingApi.md#NetworkIpServicePolicyGet) | **Get** /network/ip/service-policies/{uuid} | 



## FcInterfaceCollectionGet

> FcInterfaceResponse FcInterfaceCollectionGet(ctx, optional)



Retrieves FC interfaces. ### Related ONTAP commands * `network interface show` * `vserver fcp interface show` ### Learn more * [`DOC /network/fc/interfaces`](#docs-networking-network_fc_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FcInterfaceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FcInterfaceCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by uuid | 
 **dataProtocol** | **optional.String**| Filter by data_protocol | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **comment** | **optional.String**| Filter by comment | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **portAddress** | **optional.String**| Filter by port_address | 
 **state** | **optional.String**| Filter by state | 
 **wwpn** | **optional.String**| Filter by wwpn | 
 **locationPortName** | **optional.String**| Filter by location.port.name | 
 **locationPortNodeName** | **optional.String**| Filter by location.port.node.name | 
 **locationPortUuid** | **optional.String**| Filter by location.port.uuid | 
 **locationNodeUuid** | **optional.String**| Filter by location.node.uuid | 
 **locationNodeName** | **optional.String**| Filter by location.node.name | 
 **wwnn** | **optional.String**| Filter by wwnn | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FcInterfaceResponse**](fc_interface_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FcInterfaceCreate

> FcInterfaceResponse FcInterfaceCreate(ctx, info)



Creates an FC interface. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the FC interface. * `name` - Name of the FC interface. * `location.port.uuid` or both `location.port.name` and `location.port.node.name` - FC port on which to create the FC interface. * `data_protocol` - Data protocol for the FC interface. ### Default property values If not specified in POST, the following default property values are assigned. * `enabled` - _true_ ### Related ONTAP commands * `network interface create` ### Learn more * [`DOC /network/fc/interfaces`](#docs-networking-network_fc_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**FcInterface**](FcInterface.md)| The property values for the new FC interface.  | 

### Return type

[**FcInterfaceResponse**](fc_interface_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FcInterfaceDelete

> FcInterfaceDelete(ctx, uuid)



Deletes an FC interface. ### Related ONTAP commands * `network interface delete` ### Learn more * [`DOC /network/fc/interfaces`](#docs-networking-network_fc_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| The unique identifier for the FC interface.  | 

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


## FcInterfaceGet

> FcInterface FcInterfaceGet(ctx, uuid, optional)



Retrieves a Fibre Channel interface. ### Related ONTAP commands * `network interface show` * `vserver fcp interface show` ### Learn more * [`DOC /network/fc/interfaces`](#docs-networking-network_fc_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| The unique identifier for the FC interface.  | 
 **optional** | ***FcInterfaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FcInterfaceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FcInterface**](fc_interface.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FcInterfaceModify

> FcInterfaceModify(ctx, uuid, info)



Updates an FC interface. ### Related ONTAP commands * `network interface modify` ### Learn more * [`DOC /network/fc/interfaces`](#docs-networking-network_fc_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| The unique identifier for the FC interface.  | 
**info** | [**FcInterface**](FcInterface.md)| The new property values for the FC interface.  | 

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


## FcPortCollectionGet

> FcPortResponse FcPortCollectionGet(ctx, optional)



Retrieves FC ports.<br/> ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `fabric.name` ### Related ONTAP commands * `network fcp adapter show` ### Learn more * [`DOC /network/fc/ports`](#docs-networking-network_fc_ports) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FcPortCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FcPortCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **transceiverManufacturer** | **optional.String**| Filter by transceiver.manufacturer | 
 **transceiverCapabilities** | **optional.Int32**| Filter by transceiver.capabilities | 
 **transceiverFormFactor** | **optional.String**| Filter by transceiver.form-factor | 
 **transceiverPartNumber** | **optional.String**| Filter by transceiver.part_number | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **description** | **optional.String**| Filter by description | 
 **uuid** | **optional.String**| Filter by uuid | 
 **supportedProtocols** | **optional.String**| Filter by supported_protocols | 
 **wwpn** | **optional.String**| Filter by wwpn | 
 **wwnn** | **optional.String**| Filter by wwnn | 
 **physicalProtocol** | **optional.String**| Filter by physical_protocol | 
 **fabricConnected** | **optional.Bool**| Filter by fabric.connected | 
 **fabricPortAddress** | **optional.String**| Filter by fabric.port_address | 
 **fabricConnectedSpeed** | **optional.Int32**| Filter by fabric.connected_speed | 
 **fabricSwitchPort** | **optional.String**| Filter by fabric.switch_port | 
 **fabricName** | **optional.String**| Filter by fabric.name | 
 **state** | **optional.String**| Filter by state | 
 **speedConfigured** | **optional.String**| Filter by speed.configured | 
 **speedMaximum** | **optional.String**| Filter by speed.maximum | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FcPortResponse**](fc_port_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FcPortGet

> FcPort FcPortGet(ctx, uuid, optional)



Retrieves an FC port. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `fabric.name` ### Related ONTAP commands * `network fcp adapter show` ### Learn more * [`DOC /network/fc/ports`](#docs-networking-network_fc_ports) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| The unique identifier for the FC port.  | 
 **optional** | ***FcPortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FcPortGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FcPort**](fc_port.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FcPortModify

> FcPortModify(ctx, uuid, info)



Updates an FC port. ### Related ONTAP commands * `network fcp adapter modify` ### Learn more * [`DOC /network/fc/ports`](#docs-networking-network_fc_ports) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| The unique identifier for the FC port.  | 
**info** | [**FcPort**](FcPort.md)| The new property values for the FC port.  | 

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


## IpspaceDelete

> IpspaceDelete(ctx, uuid)



Deletes an IPspace object. ### Related ONTAP commands * `network ipspace delete`  ### Learn more * [`DOC /network/ipspaces`](#docs-networking-network_ipspaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| IPspace UUID | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpspaceGet

> Ipspace IpspaceGet(ctx, uuid, optional)



Retrieves information about a specific IPspace. ### Related ONTAP commands * `network ipspace show`  ### Learn more * [`DOC /network/ipspaces`](#docs-networking-network_ipspaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| IPspace UUID | 
 **optional** | ***IpspaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IpspaceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Ipspace**](ipspace.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpspaceModify

> IpspaceModify(ctx, uuid, optional)



Updates an IPspace object. ### Related ONTAP commands * `network ipspace rename`  ### Learn more * [`DOC /network/ipspaces`](#docs-networking-network_ipspaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| IPspace UUID | 
 **optional** | ***IpspaceModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IpspaceModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipspaceInfo** | [**optional.Interface of Ipspace**](Ipspace.md)|  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpspacesCreate

> IpspacesCreate(ctx, ipspace)



Creates a new domain within which IP addresses are unique. SVMs, ports, and networks are scoped within a single IPspace. ### Required properties * `name` - Name of the ipspace to create. ### Related ONTAP commands * `network ipspace create`  ### Learn more * [`DOC /network/ipspaces`](#docs-networking-network_ipspaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipspace** | [**Ipspace**](Ipspace.md)| IPspace identifiers | 

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


## IpspacesGet

> IpspaceResponse IpspacesGet(ctx, optional)



Retrieves a collection of IPspaces for the entire cluster. ### Related ONTAP commands * `network ipspace show`  ### Learn more * [`DOC /network/ipspaces`](#docs-networking-network_ipspaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IpspacesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IpspacesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IpspaceResponse**](ipspace_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEthernetBroadcastDomainDelete

> NetworkEthernetBroadcastDomainDelete(ctx, uuid)



Deletes a broadcast domain. ### Related ONTAP commands * `network port broadcast-domain delete`  ### Learn more * [`DOC /network/ethernet/broadcast-domains`](#docs-networking-network_ethernet_broadcast-domains)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Broadcast domain UUID | 

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


## NetworkEthernetBroadcastDomainGet

> BroadcastDomain NetworkEthernetBroadcastDomainGet(ctx, uuid, optional)



Retrieves details of a broadcast domain. ### Related ONTAP commands * `network port broadcast-domain show`  ### Learn more * [`DOC /network/ethernet/broadcast-domains`](#docs-networking-network_ethernet_broadcast-domains)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Broadcast domain UUID | 
 **optional** | ***NetworkEthernetBroadcastDomainGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkEthernetBroadcastDomainGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**BroadcastDomain**](broadcast_domain.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEthernetBroadcastDomainModify

> NetworkEthernetBroadcastDomainModify(ctx, uuid, optional)



Updates the properties of a broadcast domain. ### Related ONTAP commands * `network port broadcast-domain modify` * `network port broadcast-domain rename`  ### Learn more * [`DOC /network/ethernet/broadcast-domains`](#docs-networking-network_ethernet_broadcast-domains)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Broadcast domain UUID | 
 **optional** | ***NetworkEthernetBroadcastDomainModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkEthernetBroadcastDomainModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainInfo** | [**optional.Interface of BroadcastDomain**](BroadcastDomain.md)|  | 

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


## NetworkEthernetBroadcastDomainsCreate

> NetworkEthernetBroadcastDomainsCreate(ctx, broadcastDomain)



Creates a new broadcast domain.<br/> ### Required properties * `name` - Name of the broadcast-domain to create. * `mtu` - Maximum transmission unit of the broadcast domain. ### Recommended optional properties * `ipspace.name` or `ipspace.uuid` - IPspace the broadcast domain belongs to. ### Default property values If not specified in POST, the following default property values are assigned: * `ipspace` - _Default_ ### Related ONTAP commands * `network port broadcast-domain create`  ### Learn more * [`DOC /network/ethernet/broadcast-domains`](#docs-networking-network_ethernet_broadcast-domains)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastDomain** | [**BroadcastDomain**](BroadcastDomain.md)| Broadcast_domain parameters | 

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


## NetworkEthernetBroadcastDomainsGet

> BroadcastDomainResponse NetworkEthernetBroadcastDomainsGet(ctx, optional)



Retrieves a collection of broadcast domains for the entire cluster. ### Related ONTAP commands * `network port broadcast-domain show`  ### Learn more * [`DOC /network/ethernet/broadcast-domains`](#docs-networking-network_ethernet_broadcast-domains)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkEthernetBroadcastDomainsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkEthernetBroadcastDomainsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipspaceUuid** | **optional.String**| Filter by ipspace.uuid | 
 **ipspaceName** | **optional.String**| Filter by ipspace.name | 
 **name** | **optional.String**| Filter by name | 
 **portsUuid** | **optional.String**| Filter by ports.uuid | 
 **portsNodeName** | **optional.String**| Filter by ports.node.name | 
 **portsName** | **optional.String**| Filter by ports.name | 
 **mtu** | **optional.Int32**| Filter by mtu | 
 **uuid** | **optional.String**| Filter by uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**BroadcastDomainResponse**](broadcast_domain_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEthernetPortDelete

> NetworkEthernetPortDelete(ctx, uuid)



Deletes a VLAN or LAG (ifgrp). ### Related ONTAP commands * `network port ifgrp delete` * `network port vlan delete`  ### Learn more * [`DOC /network/ethernet/ports`](#docs-networking-network_ethernet_ports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Port UUID | 

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


## NetworkEthernetPortGet

> Port NetworkEthernetPortGet(ctx, uuid, optional)



Retrieves the details of a physical port, VLAN, or LAG. ### Related ONTAP commands * `network port show` * `network port ifgrp show` * `network port vlan show`  ### Learn more * [`DOC /network/ethernet/ports`](#docs-networking-network_ethernet_ports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Port UUID | 
 **optional** | ***NetworkEthernetPortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkEthernetPortGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Port**](port.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEthernetPortModify

> NetworkEthernetPortModify(ctx, uuid, optional)



Updates a port. ### Related ONTAP commands * `network port broadcast-domain add-ports` * `network port broadcast-domain remove-ports` * `network port ifgrp modify` * `network port modify` * `network port vlan modify`  ### Learn more * [`DOC /network/ethernet/ports`](#docs-networking-network_ethernet_ports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Port UUID | 
 **optional** | ***NetworkEthernetPortModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkEthernetPortModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portInfo** | [**optional.Interface of Port**](Port.md)|  | 

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


## NetworkEthernetPortsCreate

> PortResponse NetworkEthernetPortsCreate(ctx, port)



Creates a new VLAN (such as node1:e0a-100) or LAG (ifgrp, such as node2:a0a). ### Required properties * `node` - Node the port will be created on. * `broadcast_domain` - Broadcast domain the port is associated with. * `type` - Defines if a VLAN or LAG will be created:   * VLAN     * `vlan.base_port` - Physical port or LAG the VLAN will be created on.     * `vlan.tag` - Tag used to identify VLAN on the base port.   * LAG     * `lag.mode` - Policy for the LAG that will be created.     * `lag.distribution_policy` - Indicates how the packets are distributed between ports.     * `lag.member_ports` - Set of ports the LAG consists of. ### Related ONTAP commands * `network port ifgrp create` * `network port vlan create`  ### Learn more * [`DOC /network/ethernet/ports`](#docs-networking-network_ethernet_ports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | [**Port**](Port.md)| Port parameters | 

### Return type

[**PortResponse**](port_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkEthernetPortsGet

> PortResponse NetworkEthernetPortsGet(ctx, optional)



Retrieves a collection of ports (physical, VLAN and LAG) for an entire cluster. ### Related ONTAP commands * `network port show` * `network port ifgrp show` * `network port vlan show`  ### Learn more * [`DOC /network/ethernet/ports`](#docs-networking-network_ethernet_ports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkEthernetPortsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkEthernetPortsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by uuid | 
 **macAddress** | **optional.String**| Filter by mac_address | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **name** | **optional.String**| Filter by name | 
 **lagDistributionPolicy** | **optional.String**| Filter by lag.distribution_policy | 
 **lagActivePortsUuid** | **optional.String**| Filter by lag.active_ports.uuid | 
 **lagActivePortsNodeName** | **optional.String**| Filter by lag.active_ports.node.name | 
 **lagActivePortsName** | **optional.String**| Filter by lag.active_ports.name | 
 **lagMode** | **optional.String**| Filter by lag.mode | 
 **lagMemberPortsUuid** | **optional.String**| Filter by lag.member_ports.uuid | 
 **lagMemberPortsNodeName** | **optional.String**| Filter by lag.member_ports.node.name | 
 **lagMemberPortsName** | **optional.String**| Filter by lag.member_ports.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **speed** | **optional.Int32**| Filter by speed | 
 **broadcastDomainUuid** | **optional.String**| Filter by broadcast_domain.uuid | 
 **broadcastDomainIpspaceName** | **optional.String**| Filter by broadcast_domain.ipspace.name | 
 **broadcastDomainName** | **optional.String**| Filter by broadcast_domain.name | 
 **mtu** | **optional.Int32**| Filter by mtu | 
 **state** | **optional.String**| Filter by state | 
 **vlanBasePortUuid** | **optional.String**| Filter by vlan.base_port.uuid | 
 **vlanBasePortNodeName** | **optional.String**| Filter by vlan.base_port.node.name | 
 **vlanBasePortName** | **optional.String**| Filter by vlan.base_port.name | 
 **vlanTag** | **optional.Int32**| Filter by vlan.tag | 
 **type_** | **optional.String**| Filter by type | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**PortResponse**](port_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpInterfaceDelete

> NetworkIpInterfaceDelete(ctx, uuid)



Deletes an IP interface. ### Related ONTAP commands * `network interface delete`  ### Learn more * [`DOC /network/ip/interfaces`](#docs-networking-network_ip_interfaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| IP interface UUID | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpInterfaceGet

> IpInterface NetworkIpInterfaceGet(ctx, uuid, optional)



Retrieves details for a specific IP interface. ### Related ONTAP commands * `network interface show`  ### Learn more * [`DOC /network/ip/interfaces`](#docs-networking-network_ip_interfaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| IP interface UUID | 
 **optional** | ***NetworkIpInterfaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpInterfaceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**IpInterface**](ip_interface.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpInterfaceModify

> NetworkIpInterfaceModify(ctx, uuid, optional)



Updates an IP interface. ### Related ONTAP commands * `network interface migrate` * `network interface modify` * `network interface rename` * `network interface revert`  ### Learn more * [`DOC /network/ip/interfaces`](#docs-networking-network_ip_interfaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| IP interface UUID | 
 **optional** | ***NetworkIpInterfaceModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpInterfaceModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interfaceParameters** | [**optional.Interface of IpInterface**](IpInterface.md)|  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpInterfacesCreate

> NetworkIpInterfacesCreate(ctx, interface_)



Creates a new cluster-scoped or svm-scoped interface.<br/> ### Required properties * `name` - Name of the interface to create. * `ip.address` - IP address for the interface. * `ip.netmask` - IP subnet of the interface. * `ipspace.name` or `ipspace.uuid`   * Required for cluster-scoped interfaces.   * Optional for svm-scoped interfaces. * `svm.name` or `svm.uuid`   * Required for a svm-scoped interface.   * Invalid for a cluster-scoped interface. * `location.home_port` or `location.home_node` or `location.broadcast_domain` - One of these properties must be set to a value to define where the interface will be located. ### Default property values If not specified in POST, the following default property values are assigned: * `scope`   * _svm_ if svm parameter is specified.   * _cluster_ if svm parameter is not specified * `enabled` - _true_ * `location.auto_revert` - _true_ * `service_policy`   * _default-data-files_ if scope is `svm`   * _default-management_ if scope is `cluster` and IPspace is not `Cluster`   * _default-cluster_ if scope is `svm` and IPspace is `Cluster` * `failover` - Selects the least restrictive failover policy supported by all the services in the service policy. ### Related ONTAP commands * `network interface create`  ### Learn more * [`DOC /network/ip/interfaces`](#docs-networking-network_ip_interfaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interface_** | [**IpInterface**](IpInterface.md)| IP interface parameters | 

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


## NetworkIpInterfacesGet

> IpInterfaceResponse NetworkIpInterfacesGet(ctx, optional)



Retrieves the details of all IP interfaces. ### Related ONTAP Commands * `network interface show`  ### Learn more * [`DOC /network/ip/interfaces`](#docs-networking-network_ip_interfaces)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkIpInterfacesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpInterfacesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by uuid | 
 **services** | **optional.String**| Filter by services | 
 **servicePolicyName** | **optional.String**| Filter by service_policy.name | 
 **servicePolicyUuid** | **optional.String**| Filter by service_policy.uuid | 
 **scope** | **optional.String**| Filter by scope | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **vip** | **optional.Bool**| Filter by vip | 
 **name** | **optional.String**| Filter by name | 
 **state** | **optional.String**| Filter by state | 
 **ipAddress** | **optional.String**| Filter by ip.address | 
 **ipNetmask** | **optional.String**| Filter by ip.netmask | 
 **ipFamily** | **optional.String**| Filter by ip.family | 
 **ipspaceUuid** | **optional.String**| Filter by ipspace.uuid | 
 **ipspaceName** | **optional.String**| Filter by ipspace.name | 
 **locationAutoRevert** | **optional.Bool**| Filter by location.auto_revert | 
 **locationIsHome** | **optional.Bool**| Filter by location.is_home | 
 **locationFailover** | **optional.String**| Filter by location.failover | 
 **locationPortUuid** | **optional.String**| Filter by location.port.uuid | 
 **locationPortNodeName** | **optional.String**| Filter by location.port.node.name | 
 **locationPortName** | **optional.String**| Filter by location.port.name | 
 **locationNodeUuid** | **optional.String**| Filter by location.node.uuid | 
 **locationNodeName** | **optional.String**| Filter by location.node.name | 
 **locationHomePortUuid** | **optional.String**| Filter by location.home_port.uuid | 
 **locationHomePortNodeName** | **optional.String**| Filter by location.home_port.node.name | 
 **locationHomePortName** | **optional.String**| Filter by location.home_port.name | 
 **locationHomeNodeUuid** | **optional.String**| Filter by location.home_node.uuid | 
 **locationHomeNodeName** | **optional.String**| Filter by location.home_node.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IpInterfaceResponse**](ip_interface_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpRouteDelete

> NetworkIpRouteDelete(ctx, uuid)



Deletes a specific IP route. ### Related ONTAP commands * `network route delete`  ### Learn more * [`DOC /network/ip/routes`](#docs-networking-network_ip_routes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Route UUID | 

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


## NetworkIpRouteGet

> NetworkRoute NetworkIpRouteGet(ctx, uuid, optional)



Retrieves the details of a specific IP route. ### Related ONTAP commands * `network route show`  ### Learn more * [`DOC /network/ip/routes`](#docs-networking-network_ip_routes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Route UUID | 
 **optional** | ***NetworkIpRouteGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpRouteGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NetworkRoute**](network_route.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpRoutesCreate

> NetworkRouteResponse NetworkIpRoutesCreate(ctx, route)



Creates a cluster-scoped or SVM-scoped static route. ### Required properties * `gateway` - IP address to route packets to. * SVM-scoped routes   * `svm.name` or `svm.uuid` - SVM that route is applied to. * cluster-scoped routes   * There are no additional required fields for cluster-scoped routes. ### Default property values If not specified in POST, the following default property values are assigned: * `destination` - _0.0.0.0/0_ for IPv4 or _::/0_ for IPv6. * `ipspace.name`   * _Default_ for cluster-scoped routes.   * Name of the SVM's IPspace for SVM-scoped routes. ### Related ONTAP commands * `network route create`  ### Learn more * [`DOC /network/ip/routes`](#docs-networking-network_ip_routes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**route** | [**NetworkRoute**](NetworkRoute.md)| Route parameters | 

### Return type

[**NetworkRouteResponse**](network_route_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpRoutesGet

> NetworkRouteResponse NetworkIpRoutesGet(ctx, optional)



Retrieves the collection of IP routes. ### Related ONTAP commands * `network route show`  ### Learn more * [`DOC /network/ip/routes`](#docs-networking-network_ip_routes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkIpRoutesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpRoutesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipspaceUuid** | **optional.String**| Filter by ipspace.uuid | 
 **ipspaceName** | **optional.String**| Filter by ipspace.name | 
 **scope** | **optional.String**| Filter by scope | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **destinationAddress** | **optional.String**| Filter by destination.address | 
 **destinationNetmask** | **optional.String**| Filter by destination.netmask | 
 **destinationFamily** | **optional.String**| Filter by destination.family | 
 **uuid** | **optional.String**| Filter by uuid | 
 **gateway** | **optional.String**| Filter by gateway | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NetworkRouteResponse**](network_route_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpServicePoliciesGet

> IpServicePolicyResponse NetworkIpServicePoliciesGet(ctx, optional)



Retrieves a collection of service policies. ### Related ONTAP commands * `network interface service-policy show`  ### Learn more * [`DOC /network/ip/service-policies`](#docs-networking-network_ip_service-policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkIpServicePoliciesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpServicePoliciesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Filter by scope | 
 **ipspaceUuid** | **optional.String**| Filter by ipspace.uuid | 
 **ipspaceName** | **optional.String**| Filter by ipspace.name | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **services** | **optional.String**| Filter by services | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IpServicePolicyResponse**](ip_service_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkIpServicePolicyGet

> IpServicePolicy NetworkIpServicePolicyGet(ctx, uuid, optional)



Retrieves a specific service policy. ### Related ONTAP commands * `network interface service-policy show`  ### Learn more * [`DOC /network/ip/service-policies`](#docs-networking-network_ip_service-policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Service policy UUID | 
 **optional** | ***NetworkIpServicePolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NetworkIpServicePolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**IpServicePolicy**](ip_service_policy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

