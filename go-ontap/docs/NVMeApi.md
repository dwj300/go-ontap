# \NVMeApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FcLoginCollectionGet**](NVMeApi.md#FcLoginCollectionGet) | **Get** /network/fc/logins | 
[**FcLoginGet**](NVMeApi.md#FcLoginGet) | **Get** /network/fc/logins/{interface.uuid}/{initiator.wwpn} | 
[**NvmeInterfaceCollectionGet**](NVMeApi.md#NvmeInterfaceCollectionGet) | **Get** /protocols/nvme/interfaces | 
[**NvmeInterfaceGet**](NVMeApi.md#NvmeInterfaceGet) | **Get** /protocols/nvme/interfaces/{uuid} | 
[**NvmeNamespaceCollectionGet**](NVMeApi.md#NvmeNamespaceCollectionGet) | **Get** /storage/namespaces | 
[**NvmeNamespaceCreate**](NVMeApi.md#NvmeNamespaceCreate) | **Post** /storage/namespaces | 
[**NvmeNamespaceDelete**](NVMeApi.md#NvmeNamespaceDelete) | **Delete** /storage/namespaces/{uuid} | 
[**NvmeNamespaceGet**](NVMeApi.md#NvmeNamespaceGet) | **Get** /storage/namespaces/{uuid} | 
[**NvmeNamespaceModify**](NVMeApi.md#NvmeNamespaceModify) | **Patch** /storage/namespaces/{uuid} | 
[**NvmeServiceCollectionGet**](NVMeApi.md#NvmeServiceCollectionGet) | **Get** /protocols/nvme/services | 
[**NvmeServiceCreate**](NVMeApi.md#NvmeServiceCreate) | **Post** /protocols/nvme/services | 
[**NvmeServiceDelete**](NVMeApi.md#NvmeServiceDelete) | **Delete** /protocols/nvme/services/{svm.uuid} | 
[**NvmeServiceGet**](NVMeApi.md#NvmeServiceGet) | **Get** /protocols/nvme/services/{svm.uuid} | 
[**NvmeServiceModify**](NVMeApi.md#NvmeServiceModify) | **Patch** /protocols/nvme/services/{svm.uuid} | 
[**NvmeSubsystemCollectionGet**](NVMeApi.md#NvmeSubsystemCollectionGet) | **Get** /protocols/nvme/subsystems | 
[**NvmeSubsystemControllerCollectionGet**](NVMeApi.md#NvmeSubsystemControllerCollectionGet) | **Get** /protocols/nvme/subsystem-controllers | 
[**NvmeSubsystemControllerGet**](NVMeApi.md#NvmeSubsystemControllerGet) | **Get** /protocols/nvme/subsystem-controllers/{subsystem.uuid}/{id} | 
[**NvmeSubsystemCreate**](NVMeApi.md#NvmeSubsystemCreate) | **Post** /protocols/nvme/subsystems | 
[**NvmeSubsystemDelete**](NVMeApi.md#NvmeSubsystemDelete) | **Delete** /protocols/nvme/subsystems/{uuid} | 
[**NvmeSubsystemGet**](NVMeApi.md#NvmeSubsystemGet) | **Get** /protocols/nvme/subsystems/{uuid} | 
[**NvmeSubsystemHostCollectionGet**](NVMeApi.md#NvmeSubsystemHostCollectionGet) | **Get** /protocols/nvme/subsystems/{subsystem.uuid}/hosts | 
[**NvmeSubsystemHostCreate**](NVMeApi.md#NvmeSubsystemHostCreate) | **Post** /protocols/nvme/subsystems/{subsystem.uuid}/hosts | 
[**NvmeSubsystemHostDelete**](NVMeApi.md#NvmeSubsystemHostDelete) | **Delete** /protocols/nvme/subsystems/{subsystem.uuid}/hosts/{nqn} | 
[**NvmeSubsystemHostGet**](NVMeApi.md#NvmeSubsystemHostGet) | **Get** /protocols/nvme/subsystems/{subsystem.uuid}/hosts/{nqn} | 
[**NvmeSubsystemMapCollectionGet**](NVMeApi.md#NvmeSubsystemMapCollectionGet) | **Get** /protocols/nvme/subsystem-maps | 
[**NvmeSubsystemMapCreate**](NVMeApi.md#NvmeSubsystemMapCreate) | **Post** /protocols/nvme/subsystem-maps | 
[**NvmeSubsystemMapDelete**](NVMeApi.md#NvmeSubsystemMapDelete) | **Delete** /protocols/nvme/subsystem-maps/{subsystem.uuid}/{namespace.uuid} | 
[**NvmeSubsystemMapGet**](NVMeApi.md#NvmeSubsystemMapGet) | **Get** /protocols/nvme/subsystem-maps/{subsystem.uuid}/{namespace.uuid} | 
[**NvmeSubsystemModify**](NVMeApi.md#NvmeSubsystemModify) | **Patch** /protocols/nvme/subsystems/{uuid} | 


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

# **NvmeInterfaceCollectionGet**
> NvmeInterfaceResponse NvmeInterfaceCollectionGet(ctx, optional)


Retrieves NVMe interfaces. ### Related ONTAP commands * `vserver nvme show-interface` ### Learn more * [`DOC /protocols/nvme/interfaces`](#docs-NVMe-protocols_nvme_interfaces) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvmeInterfaceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeInterfaceCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transportAddress** | **optional.String**| Filter by transport_address | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **fcInterfaceWwpn** | **optional.String**| Filter by fc_interface.wwpn | 
 **fcInterfaceWwnn** | **optional.String**| Filter by fc_interface.wwnn | 
 **fcInterfacePortName** | **optional.String**| Filter by fc_interface.port.name | 
 **fcInterfacePortNodeName** | **optional.String**| Filter by fc_interface.port.node.name | 
 **fcInterfacePortUuid** | **optional.String**| Filter by fc_interface.port.uuid | 
 **uuid** | **optional.String**| Filter by uuid | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeInterfaceResponse**](nvme_interface_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeInterfaceGet**
> NvmeInterface NvmeInterfaceGet(ctx, uuid)


Retrieves an NVMe interface. ### Related ONTAP commands * `vserver nvme show-interface` ### Learn more * [`DOC /protocols/nvme/interfaces`](#docs-NVMe-protocols_nvme_interfaces) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe interface.  | 

### Return type

[**NvmeInterface**](nvme_interface.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeNamespaceCollectionGet**
> NvmeNamespaceResponse NvmeNamespaceCollectionGet(ctx, optional)


Retrieves NVMe namespaces. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `auto_delete` * `subsystem_map.*` * `status.mapped` ### Related ONTAP commands * `vserver nvme namespace show` * `vserver nvme subsystem map show` ### Learn more * [`DOC /storage/namespaces`](#docs-NVMe-storage_namespaces) to learn more and examples. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvmeNamespaceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeNamespaceCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **optional.Bool**| Filter by enabled | 
 **comment** | **optional.String**| Filter by comment | 
 **subsystemMapNsid** | **optional.String**| Filter by subsystem_map.nsid | 
 **subsystemMapAnagrpid** | **optional.String**| Filter by subsystem_map.anagrpid | 
 **subsystemMapSubsystemName** | **optional.String**| Filter by subsystem_map.subsystem.name | 
 **subsystemMapSubsystemUuid** | **optional.String**| Filter by subsystem_map.subsystem.uuid | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **osType** | **optional.String**| Filter by os_type | 
 **uuid** | **optional.String**| Filter by uuid | 
 **autoDelete** | **optional.Bool**| Filter by auto_delete | 
 **statusContainerState** | **optional.String**| Filter by status.container_state | 
 **statusMapped** | **optional.Bool**| Filter by status.mapped | 
 **statusState** | **optional.String**| Filter by status.state | 
 **statusReadOnly** | **optional.Bool**| Filter by status.read_only | 
 **locationQtreeName** | **optional.String**| Filter by location.qtree.name | 
 **locationQtreeId** | **optional.Int32**| Filter by location.qtree.id | 
 **locationVolumeName** | **optional.String**| Filter by location.volume.name | 
 **locationVolumeUuid** | **optional.String**| Filter by location.volume.uuid | 
 **locationNamespace** | **optional.String**| Filter by location.namespace | 
 **spaceGuaranteeReserved** | **optional.Bool**| Filter by space.guarantee.reserved | 
 **spaceGuaranteeRequested** | **optional.Bool**| Filter by space.guarantee.requested | 
 **spaceBlockSize** | **optional.Int32**| Filter by space.block_size | 
 **spaceSize** | **optional.Int32**| Filter by space.size | 
 **spaceUsed** | **optional.Int32**| Filter by space.used | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeNamespaceResponse**](nvme_namespace_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeNamespaceCreate**
> NvmeNamespaceResponse NvmeNamespaceCreate(ctx, info)


Creates an NVMe namespace. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the NVMe namespace. * `name`, `location.volume.name` or `location.volume.uuid` - Existing volume in which to create the NVMe namespace. * `name` or `location.namespace` - Base name for the NVMe namespace. * `os_type` - Operating system from which the NVMe namespace will be accessed. (Not used for clones, which are created based on the `os_type` of the source NVMe namespace.) * `space.size` - Size for the NVMe namespace. (Not used for clones, which are created based on the size of the source NVMe namespace.) ### Default property values If not specified in POST, the following default property values are assigned: * `auto_delete` - _false_ * `space.block_size` - _4096_ ### Related ONTAP commands * `volume file clone autodelete` * `volume file clone create` * `vserver nvme namespace create` ### Learn more * [`DOC /storage/namespaces`](#docs-NVMe-storage_namespaces) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**NvmeNamespace**](NvmeNamespace.md)| The property values for the new NVMe namespace.  | 

### Return type

[**NvmeNamespaceResponse**](nvme_namespace_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeNamespaceDelete**
> NvmeNamespaceDelete(ctx, uuid, optional)


Deletes an NVMe namespace. ### Related ONTAP commands * `vserver nvme namespace delete` ### Learn more * [`DOC /storage/namespaces`](#docs-NVMe-storage_namespaces) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe namespace to delete.  | 
 **optional** | ***NvmeNamespaceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeNamespaceDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowDeleteWhileMapped** | **optional.Bool**| Allows deletion of a mapped NVMe namespace. A mapped NVMe namespace might be in use. Deleting a mapped namespace also deletes the namespace map and makes the data no longer available, possibly causing a disruption in the availability of data. **This parameter should be used with caution.**  | [default to false]

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeNamespaceGet**
> NvmeNamespace NvmeNamespaceGet(ctx, uuid, optional)


Retrieves an NVMe namespace. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `auto_delete` * `subsystem_map.*` * `status.mapped` ### Related ONTAP commands * `vserver nvme namespace show` * `vserver nvme subsystem map show` ### Learn more * [`DOC /storage/namespaces`](#docs-NVMe-storage_namespaces) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe namespace to retrieve.  | 
 **optional** | ***NvmeNamespaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeNamespaceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NvmeNamespace**](nvme_namespace.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeNamespaceModify**
> NvmeNamespaceModify(ctx, uuid, info)


Updates an NVMe namespace. ### Related ONTAP commands * `volume file clone autodelete` * `vserver nvme namespace modify` ### Learn more * [`DOC /storage/namespaces`](#docs-NVMe-storage_namespaces) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe namespace to update.  | 
  **info** | [**NvmeNamespace**](NvmeNamespace.md)| The new property values for the NVMe namespace.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeServiceCollectionGet**
> NvmeServiceResponse NvmeServiceCollectionGet(ctx, optional)


Retrieves NVMe services. ### Related ONTAP commands * `vserver nvme show` ### Learn more * [`DOC /protocols/nvme/services`](#docs-NVMe-protocols_nvme_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvmeServiceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeServiceCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeServiceResponse**](nvme_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeServiceCreate**
> NvmeServiceResponse NvmeServiceCreate(ctx, info)


Creates an NVMe service. ### Required properties * `svm.uuid` or `svm.name` - The existing SVM in which to create the NVMe service. ### Related ONTAP commands * `vserver nvme create` ### Learn more * [`DOC /protocols/nvme/services`](#docs-NVMe-protocols_nvme_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**NvmeService**](NvmeService.md)| The property values for the new NVMe service.  | 

### Return type

[**NvmeServiceResponse**](nvme_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeServiceDelete**
> NvmeServiceDelete(ctx, svmUuid)


Deletes an NVMe service. An NVMe service must be disabled before it can be deleted. In addition, all NVMe interfaces, subsystems, and subsystem maps associated with the SVM must first be deleted. ### Related ONTAP commands * `vserver nvme delete` ### Learn more * [`DOC /protocols/nvme/services`](#docs-NVMe-protocols_nvme_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM whose NVMe service is to be deleted.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeServiceGet**
> NvmeService NvmeServiceGet(ctx, svmUuid, optional)


Retrieves an NVMe service. ### Related ONTAP commands * `vserver nvme show` ### Learn more * [`DOC /protocols/nvme/services`](#docs-NVMe-protocols_nvme_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM whose NVMe service is to be retrieved.  | 
 **optional** | ***NvmeServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeServiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NvmeService**](nvme_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeServiceModify**
> NvmeServiceModify(ctx, svmUuid, info)


Updates an NVMe service. ### Related ONTAP commands * `vserver nvme modify` ### Learn more * [`DOC /protocols/nvme/services`](#docs-NVMe-protocols_nvme_services) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **svmUuid** | **string**| The unique identifier of the SVM whose NVMe service is to be updated.  | 
  **info** | [**NvmeService**](NvmeService.md)| The new property values for the NVMe service.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemCollectionGet**
> NvmeSubsystemResponse NvmeSubsystemCollectionGet(ctx, optional)


Retrieves NVMe subsystems. ### Related ONTAP commands * `vserver nvme subsystem host show` * `vserver nvme subsystem map show` * `vserver nvme subsystem show` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvmeSubsystemCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **optional.String**| Filter by comment | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **ioQueueDefaultCount** | **optional.Int32**| Filter by io_queue.default.count | 
 **ioQueueDefaultDepth** | **optional.Int32**| Filter by io_queue.default.depth | 
 **osType** | **optional.String**| Filter by os_type | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **uuid** | **optional.String**| Filter by uuid | 
 **hostsNqn** | **optional.String**| Filter by hosts.nqn | 
 **subsystemMapsNamespaceName** | **optional.String**| Filter by subsystem_maps.namespace.name | 
 **subsystemMapsNamespaceUuid** | **optional.String**| Filter by subsystem_maps.namespace.uuid | 
 **subsystemMapsAnagrpid** | **optional.String**| Filter by subsystem_maps.anagrpid | 
 **subsystemMapsNsid** | **optional.String**| Filter by subsystem_maps.nsid | 
 **targetNqn** | **optional.String**| Filter by target_nqn | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeSubsystemResponse**](nvme_subsystem_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemControllerCollectionGet**
> NvmeSubsystemControllerResponse NvmeSubsystemControllerCollectionGet(ctx, optional)


Retrieves NVMe subsystem controllers. ### Related ONTAP commands * `vserver nvme subsystem controller show` ### Learn more * [`DOC /protocols/nvme/subsystem-controllers`](#docs-NVMe-protocols_nvme_subsystem-controllers) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvmeSubsystemControllerCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemControllerCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interfaceUuid** | **optional.String**| Filter by interface.uuid | 
 **interfaceName** | **optional.String**| Filter by interface.name | 
 **interfaceTransportAddress** | **optional.String**| Filter by interface.transport_address | 
 **id** | **optional.String**| Filter by id | 
 **adminQueueDepth** | **optional.Int32**| Filter by admin_queue.depth | 
 **hostId** | **optional.String**| Filter by host.id | 
 **hostTransportAddress** | **optional.String**| Filter by host.transport_address | 
 **hostNqn** | **optional.String**| Filter by host.nqn | 
 **ioQueueCount** | **optional.Int32**| Filter by io_queue.count | 
 **ioQueueDepth** | **optional.Int32**| Filter by io_queue.depth | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **subsystemName** | **optional.String**| Filter by subsystem.name | 
 **subsystemUuid** | **optional.String**| Filter by subsystem.uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeSubsystemControllerResponse**](nvme_subsystem_controller_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemControllerGet**
> NvmeSubsystemController NvmeSubsystemControllerGet(ctx, subsystemUuid, id, optional)


Retrieves an NVMe subsystem controller. ### Related ONTAP commands * `vserver nvme subsystem controller show` ### Learn more * [`DOC /protocols/nvme/subsystem-controllers`](#docs-NVMe-protocols_nvme_subsystem-controllers) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **id** | **string**| The unique identifier of the NVMe subsystem controller.  | 
 **optional** | ***NvmeSubsystemControllerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemControllerGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NvmeSubsystemController**](nvme_subsystem_controller.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemCreate**
> NvmeSubsystemResponse NvmeSubsystemCreate(ctx, info)


Creates an NVMe subsystem. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the NVMe subsystem. * `name` - Name for NVMe subsystem. Once created, an NVMe subsytem cannot be renamed. * `os_type` - Operating system of the NVMe subsystem's hosts. ### Related ONTAP commands * `vserver nvme subsystem create` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**NvmeSubsystem**](NvmeSubsystem.md)| The property values for the new NVMe subsystem.  | 

### Return type

[**NvmeSubsystemResponse**](nvme_subsystem_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemDelete**
> NvmeSubsystemDelete(ctx, uuid, optional)


Removes an NVMe subsystem. ### Related ONTAP commands * `vserver nvme subsystem delete` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe subsystem.  | 
 **optional** | ***NvmeSubsystemDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowDeleteWhileMapped** | **optional.Bool**| Allows for the deletion of a mapped NVMe subsystem.  | 
 **allowDeleteWithHosts** | **optional.Bool**| Allows for the deletion of an NVMe subsystem with NVMe hosts.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemGet**
> NvmeSubsystem NvmeSubsystemGet(ctx, uuid, optional)


Retrieves an NVMe subsystem. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `subsystem_maps.*` ### Related ONTAP commands * `vserver nvme subsystem host show` * `vserver nvme subsystem map show` * `vserver nvme subsystem show` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe subsystem.  | 
 **optional** | ***NvmeSubsystemGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NvmeSubsystem**](nvme_subsystem.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemHostCollectionGet**
> NvmeSubsystemHostResponse NvmeSubsystemHostCollectionGet(ctx, subsystemUuid, optional)


Retrieves the NVMe subsystem hosts of an NVMe subsystem. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `subsystem_maps.*` ### Related ONTAP commands * `vserver nvme subsystem map show` * `vserver nvme subsystem show` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
 **optional** | ***NvmeSubsystemHostCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemHostCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeSubsystemHostResponse**](nvme_subsystem_host_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemHostCreate**
> NvmeSubsystemHostResponse NvmeSubsystemHostCreate(ctx, subsystemUuid, info)


Adds NVMe subsystem host(s) to an NVMe subsystem. ### Required properties * `nqn` or `records.nqn` - NVMe host(s) NQN(s) to add to the NVMe subsystem. ### Related ONTAP commands * `vserver nvme subsystem host add` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **info** | [**NvmeSubsystemHost**](NvmeSubsystemHost.md)| The property values for the NVMe subsystem host to add to the NVMe subsystem.  | 

### Return type

[**NvmeSubsystemHostResponse**](nvme_subsystem_host_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemHostDelete**
> NvmeSubsystemHostDelete(ctx, subsystemUuid, nqn)


Deletes an NVMe subsystem host from an NVMe subsystem. ### Related ONTAP commands * `vserver nvme subsystem host remove` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **nqn** | **string**| The NVMe qualified name (NQN) used to identify the NVMe subsystem host.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemHostGet**
> NvmeSubsystemHost NvmeSubsystemHostGet(ctx, subsystemUuid, nqn, optional)


Retrieves an NVMe subsystem host of an NVMe subsystem. ### Related ONTAP commands * `vserver nvme subsystem host show` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **nqn** | **string**| The NVMe qualified name (NQN) used to identify the NVMe subsystem host.  | 
 **optional** | ***NvmeSubsystemHostGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemHostGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NvmeSubsystemHost**](nvme_subsystem_host.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemMapCollectionGet**
> NvmeSubsystemMapResponse NvmeSubsystemMapCollectionGet(ctx, optional)


Retrieves NVMe subsystem maps. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `anagrpid` ### Related ONTAP commands * `vserver nvme subsystem map show` ### Learn more * [`DOC /protocols/nvme/subsystem-maps`](#docs-NVMe-protocols_nvme_subsystem-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvmeSubsystemMapCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemMapCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **nsid** | **optional.String**| Filter by nsid | 
 **namespaceName** | **optional.String**| Filter by namespace.name | 
 **namespaceUuid** | **optional.String**| Filter by namespace.uuid | 
 **namespaceNodeUuid** | **optional.String**| Filter by namespace.node.uuid | 
 **namespaceNodeName** | **optional.String**| Filter by namespace.node.name | 
 **subsystemName** | **optional.String**| Filter by subsystem.name | 
 **subsystemUuid** | **optional.String**| Filter by subsystem.uuid | 
 **anagrpid** | **optional.String**| Filter by anagrpid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NvmeSubsystemMapResponse**](nvme_subsystem_map_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemMapCreate**
> NvmeSubsystemMapResponse NvmeSubsystemMapCreate(ctx, info)


Creates an NVMe subsystem map. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the NVMe subsystem map. * `namespace.uuid` or `namespace.name` - Existing NVMe namespace to map to the specified NVme subsystem. * `subsystem.uuid` or `subsystem.name` - Existing NVMe subsystem to map to the specified NVMe namespace. ### Related ONTAP commands * `vserver nvme subsystem map create` ### Learn more * [`DOC /protocols/nvme/subsystem-maps`](#docs-NVMe-protocols_nvme_subsystem-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**NvmeSubsystemMap**](NvmeSubsystemMap.md)| The property values for the new NVMe subsystem map.  | 

### Return type

[**NvmeSubsystemMapResponse**](nvme_subsystem_map_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemMapDelete**
> NvmeSubsystemMapResponse NvmeSubsystemMapDelete(ctx, subsystemUuid, namespaceUuid)


Deletes an NVMe subsystem map. ### Related ONTAP commands * `vserver nvme subsystem map delete` ### Learn more * [`DOC /protocols/nvme/subsystem-maps`](#docs-NVMe-protocols_nvme_subsystem-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **namespaceUuid** | **string**| The unique identifier of the NVMe namespace.  | 

### Return type

[**NvmeSubsystemMapResponse**](nvme_subsystem_map_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemMapGet**
> NvmeSubsystemMap NvmeSubsystemMapGet(ctx, subsystemUuid, namespaceUuid, optional)


Retrieves an NVMe subsystem map. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `anagrpid` ### Related ONTAP commands * `vserver nvme subsystem map show` ### Learn more * [`DOC /protocols/nvme/subsystem-maps`](#docs-NVMe-protocols_nvme_subsystem-maps) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subsystemUuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **namespaceUuid** | **string**| The unique identifier of the NVMe namespace.  | 
 **optional** | ***NvmeSubsystemMapGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvmeSubsystemMapGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NvmeSubsystemMap**](nvme_subsystem_map.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvmeSubsystemModify**
> NvmeSubsystemModify(ctx, uuid, info)


Updates an NVMe subsystem. ### Related ONTAP commands * `vserver nvme subsystem modify` ### Learn more * [`DOC /protocols/nvme/subsystems`](#docs-NVMe-protocols_nvme_subsystems) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The unique identifier of the NVMe subsystem.  | 
  **info** | [**NvmeSubsystem**](NvmeSubsystem.md)| The new property values for the NVMe subsystem.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

