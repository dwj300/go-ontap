# \ApplicationApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationCollectionGet**](ApplicationApi.md#ApplicationCollectionGet) | **Get** /application/applications | 
[**ApplicationComponentCollectionGet**](ApplicationApi.md#ApplicationComponentCollectionGet) | **Get** /application/applications/{application.uuid}/components | 
[**ApplicationComponentGet**](ApplicationApi.md#ApplicationComponentGet) | **Get** /application/applications/{application.uuid}/components/{uuid} | 
[**ApplicationComponentSnapshotCollectionGet**](ApplicationApi.md#ApplicationComponentSnapshotCollectionGet) | **Get** /application/applications/{application.uuid}/components/{component.uuid}/snapshots | 
[**ApplicationComponentSnapshotCreate**](ApplicationApi.md#ApplicationComponentSnapshotCreate) | **Post** /application/applications/{application.uuid}/components/{component.uuid}/snapshots | 
[**ApplicationComponentSnapshotDelete**](ApplicationApi.md#ApplicationComponentSnapshotDelete) | **Delete** /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid} | 
[**ApplicationComponentSnapshotGet**](ApplicationApi.md#ApplicationComponentSnapshotGet) | **Get** /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid} | 
[**ApplicationComponentSnapshotRestore**](ApplicationApi.md#ApplicationComponentSnapshotRestore) | **Post** /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}/restore | 
[**ApplicationCreate**](ApplicationApi.md#ApplicationCreate) | **Post** /application/applications | 
[**ApplicationDelete**](ApplicationApi.md#ApplicationDelete) | **Delete** /application/applications/{uuid} | 
[**ApplicationGet**](ApplicationApi.md#ApplicationGet) | **Get** /application/applications/{uuid} | 
[**ApplicationModify**](ApplicationApi.md#ApplicationModify) | **Patch** /application/applications/{uuid} | 
[**ApplicationSnapshotCollectionGet**](ApplicationApi.md#ApplicationSnapshotCollectionGet) | **Get** /application/applications/{application.uuid}/snapshots | 
[**ApplicationSnapshotCreate**](ApplicationApi.md#ApplicationSnapshotCreate) | **Post** /application/applications/{application.uuid}/snapshots | 
[**ApplicationSnapshotDelete**](ApplicationApi.md#ApplicationSnapshotDelete) | **Delete** /application/applications/{application.uuid}/snapshots/{uuid} | 
[**ApplicationSnapshotGet**](ApplicationApi.md#ApplicationSnapshotGet) | **Get** /application/applications/{application.uuid}/snapshots/{uuid} | 
[**ApplicationSnapshotRestore**](ApplicationApi.md#ApplicationSnapshotRestore) | **Post** /application/applications/{application.uuid}/snapshots/{uuid}/restore | 
[**ApplicationTemplateCollectionGet**](ApplicationApi.md#ApplicationTemplateCollectionGet) | **Get** /application/templates | 
[**ApplicationTemplateGet**](ApplicationApi.md#ApplicationTemplateGet) | **Get** /application/templates/{name} | 



## ApplicationCollectionGet

> ApplicationResponse ApplicationCollectionGet(ctx, optional)



Retrieves applications. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `<template>` the property corresponding to the `template.name` of the application ### Query examples Numerous queries are available for classifying and sorting applications: 1. Return a list of applications sorted by name.     ```     GET /application/applications?order_by=name     ```     <br/> 2. Return a list of applications for a specific SVM.     ```     GET /application/applications?svm.name=<name>     ```     <br/> 3. Return a list of all SQL applications.     ```     GET /application/applications?template.name=sql*     ```     <br/> 4. Return a list of all applications that can be accessed via SAN.<br/>     ```     GET /application/applications?template.protocol=san     ```     <br/> 5. Return the top five applications consuming the most IOPS.<br/>     ```     GET /application/applications?order_by=statistics.iops.total desc&max_records=5     ``` <br/>The above examples are not comprehensive. There are many more properties available for queries. Also, multiple queries can be mixed and matched with other query parameters for a large variety of requests. See the per-property documentation below for the full list of supported query parameters. ### Learn more * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by UUID | 
 **name** | **optional.String**| Filter by name | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **templateName** | **optional.String**| Filter by template.name | 
 **templateVersion** | **optional.String**| Filter by template.version | 
 **templateProtocol** | **optional.String**| Filter by template.protocol | 
 **generation** | **optional.String**| Filter by generation | 
 **state** | **optional.String**| Filter by state | 
 **protectionGranularity** | **optional.String**| Filter by protection granularity | 
 **rpoIsSupported** | **optional.String**| Filter by rpo.is_supported | 
 **rpoLocalName** | **optional.String**| Filter by rpo.local.name | 
 **rpoLocalDescription** | **optional.String**| Filter by rpo.local.description | 
 **rpoRemoteName** | **optional.String**| Filter by rpo.remote.name | 
 **rpoRemoteDescription** | **optional.String**| Filter by rpo.remote.description | 
 **rpoComponentsName** | **optional.String**| Filter by rpo.components.name | 
 **rpoComponentsUuid** | **optional.String**| Filter by rpo.components.uuid | 
 **rpoComponentsLocalName** | **optional.String**| Filter by rpo.components.rpo.local.name | 
 **rpoComponentsRpoLocalDescription** | **optional.String**| Filter by rpo.components.rpo.local.description | 
 **rpoComponentsRpoRemoteName** | **optional.String**| Filter by rpo.components.rpo.remote.name | 
 **rpoComponentsRpoRemoteDescription** | **optional.String**| Filter by rpo.components.rpo.remote.description | 
 **statisticsSpaceProvisioned** | **optional.String**| Filter by statistics.space.provisioned | 
 **statisticsSpaceUsed** | **optional.String**| Filter by statistics.space.used | 
 **statisticsSpaceUsedPercent** | **optional.String**| Filter by statistics.space.used_percent | 
 **statisticsSpaceUsedExcludingReserves** | **optional.String**| Filter by statistics.space.used_excluding_reserves | 
 **statisticsSpaceLogicalUsed** | **optional.String**| Filter by statistics.space.logical_used | 
 **statisticsSpaceReservedUnused** | **optional.String**| Filter by statistics.space.reserved_unused | 
 **statisticsSpaceAvailable** | **optional.String**| Filter by statistics.space.available | 
 **statisticsSpaceSavings** | **optional.String**| Filter by statistics.space.savings | 
 **statisticsIopsTotal** | **optional.String**| Filter by statistics.iops.total | 
 **statisticsIopsPerTb** | **optional.String**| Filter by statistics.iops.per_tb | 
 **statisticsSnapshotReserve** | **optional.String**| Filter by statistics.snapshot.reserve | 
 **statisticsSnapshotUsed** | **optional.String**| Filter by statistics.snapshot.used | 
 **statisticsLatencyRaw** | **optional.String**| Filter by statistics.latency.raw | 
 **statisticsLatencyAverage** | **optional.String**| Filter by statistics.latency.average | 
 **statisticsStatisticsIncomplete** | **optional.String**| Filter by statistics.statistics_incomplete | 
 **statisticsSharedStoragePool** | **optional.String**| Filter by statistics.shared_storage_pool | 
 **statisticsComponentsName** | **optional.String**| Filter by statistics.components.name | 
 **statisticsComponentsUuid** | **optional.String**| Filter by statistics.components.uuid | 
 **statisticsComponentsStorageServiceName** | **optional.String**| Filter by statistics.components.storage_service.name | 
 **statisticsComponentsSpaceProvisioned** | **optional.String**| Filter by statistics.components.space.provisioned | 
 **statisticsComponentsSpaceUsed** | **optional.String**| Filter by statistics.components.space.used | 
 **statisticsComponentsSpaceUsedPercent** | **optional.String**| Filter by statistics.components.space.used_percent | 
 **statisticsComponentsSpaceUsedExcludingReserves** | **optional.String**| Filter by statistics.components.space.used_excluding_reserves | 
 **statisticsComponentsSpaceLogicalUsed** | **optional.String**| Filter by statistics.components.space.logical_used | 
 **statisticsComponentsSpaceReservedUnused** | **optional.String**| Filter by statistics.components.space.reserved_unused | 
 **statisticsComponentsSpaceAvailable** | **optional.String**| Filter by statistics.components.space.available | 
 **statisticsComponentsSpaceSavings** | **optional.String**| Filter by statistics.components.space.savings | 
 **statisticsComponentsIopsTotal** | **optional.String**| Filter by statistics.components.iops.total | 
 **statisticsComponentsIopsPerTb** | **optional.String**| Filter by statistics.components.iops.per_tb | 
 **statisticsComponentsSnapshotReserve** | **optional.String**| Filter by statistics.components.snapshot.reserve | 
 **statisticsComponentsSnapshotUsed** | **optional.String**| Filter by statistics.components.snapshot.used | 
 **statisticsComponentsLatencyRaw** | **optional.String**| Filter by statistics.components.latency.raw | 
 **statisticsComponentsLatencyAverage** | **optional.String**| Filter by statistics.components.latency.average | 
 **statisticsComponentsStatisticsIncomplete** | **optional.String**| Filter by statistics.components.statistics_incomplete | 
 **statisticsComponentsSharedStoragePool** | **optional.String**| Filter by statistics.components.shared_storage_pool | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ApplicationResponse**](application_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationComponentCollectionGet

> ApplicationComponentResponse ApplicationComponentCollectionGet(ctx, applicationUuid, optional)



Retrieves application components. ### Overview The application component object exposes how to access an application. Most application interfaces abstract away the underlying ONTAP storage elements, but this interface exposes what is necessary to connect to and uses the storage that is provisioned for an application. See the application component model for a detailed description of each property. ### Query examples Queries are limited on this API. Most of the details are nested under the `nfs_access`, `cifs_access`, or `san_access` properties, but those properties do not support queries, and properties nested under those properties cannot be requested individually in the current release.<br/> The following query returns all application components with names beginning in _secondary_.<br/><br/> ``` GET /application/applications/{application.uuid}/components?name=secondary* ``` <br/>The following query returns all application components at the _extreme_ storage service.<br/><br/> ``` GET /application/applications/{application.uuid}/components?storage_service.name=extreme ``` ### Learn more * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
 **optional** | ***ApplicationComponentCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationComponentCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uuid** | **optional.String**| Filter by UUID | 
 **name** | **optional.String**| Filter by name | 
 **storageServiceName** | **optional.String**| Filter by storage_service.name | 
 **storageServiceUuid** | **optional.String**| Filter by storage_service.uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ApplicationComponentResponse**](application_component_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationComponentGet

> ApplicationComponent ApplicationComponentGet(ctx, applicationUuid, uuid, optional)



Retrieves an application component. ### Overview The application component object exposes how to access an application. Most application interfaces abstract away the underlying ONTAP storage elements, but this interface exposes what is necessary to connect to and uses the storage that is provisioned for an application. See the application component model for a detailed description of each property. ### Access Each application component can be accessed via NFS, CIFS, or SAN. NFS and CIFS access can be enabled simultaneously. Each access section includes a `backing_storage` property. This property is used to correlate the storage elements with the access elements of the application. The `backing_storage` portion of the access section provides the `type` and `uuid` of the backing storage. There is another `backing_storage` property at the same level as the access properties which contains lists of backing storage elements corresponding to the types listed in the access section. ### Learn more * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**uuid** | **string**| Application component UUID | 
 **optional** | ***ApplicationComponentGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationComponentGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ApplicationComponent**](application_component.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationComponentSnapshotCollectionGet

> ApplicationComponentSnapshotResponse ApplicationComponentSnapshotCollectionGet(ctx, applicationUuid, componentUuid, optional)



Retrieves Snapshot copies of an application component.<br/> This endpoint is only supported for Maxdata template applications.<br/> Component Snapshot copies are essentially more granular application Snapshot copies. There is no difference beyond the scope of the operation. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`GET /application/applications/{uuid}/snapshots`](#operations-application-application_snapshot_collection_get) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**componentUuid** | **string**| Application Component UUID | 
 **optional** | ***ApplicationComponentSnapshotCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationComponentSnapshotCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentName** | **optional.String**| Filter by Application Component Name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **name** | **optional.String**| Filter by name | 
 **consistencyType** | **optional.String**| Filter by consistency_type | 
 **comment** | **optional.String**| Filter by comment | 
 **createTime** | **optional.String**| Filter by create_time | 
 **isPartial** | **optional.String**| Filter by is_partial | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ApplicationComponentSnapshotResponse**](application_component_snapshot_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationComponentSnapshotCreate

> JobLinkResponse ApplicationComponentSnapshotCreate(ctx, applicationUuid, componentUuid, optional)



Creates a Snapshot copy of an application component.<br/> This endpoint is only supported for Maxdata template applications.<br/> ### Required properties * `name` ### Recommended optional properties * `consistency_type` - Track whether this snapshot is _application_ or _crash_ consistent. Component Snapshot copies are essentially more granular application Snapshot copies. There is no difference beyond the scope of the operation. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`GET /application/applications/{uuid}/snapshots`](#operations-application-application_snapshot_create) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**componentUuid** | **string**| Application Component UUID | 
 **optional** | ***ApplicationComponentSnapshotCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationComponentSnapshotCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of ApplicationComponentSnapshot**](ApplicationComponentSnapshot.md)| Info specification | 

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


## ApplicationComponentSnapshotDelete

> JobLinkResponse ApplicationComponentSnapshotDelete(ctx, applicationUuid, componentUuid, uuid)



Delete a Snapshot copy of an application component.<br/> This endpoint is only supported for Maxdata template applications.<br/> Component Snapshot copies are essentially more granular application Snapshot copies. There is no difference beyond the scope of the operation. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`DELETE /application/applications/{uuid}/snapshots`](#operations-application-application_snapshot_delete) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**componentUuid** | **string**| Application Component UUID | 
**uuid** | **string**| Snapshot UUID | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationComponentSnapshotGet

> ApplicationComponentSnapshot ApplicationComponentSnapshotGet(ctx, applicationUuid, componentUuid, uuid, optional)



Retrieve a Snapshot copy of an application component.<br/> This endpoint is only supported for Maxdata template applications.<br/> Component Snapshot copies are essentially more granular application Snapshot copies. There is no difference beyond the scope of the operation. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`GET /application/applications/{uuid}/snapshots`](#operations-application-application_snapshot_get) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**componentUuid** | **string**| Application Component UUID | 
**uuid** | **string**| Snapshot UUID | 
 **optional** | ***ApplicationComponentSnapshotGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationComponentSnapshotGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ApplicationComponentSnapshot**](application_component_snapshot.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationComponentSnapshotRestore

> JobLinkResponse ApplicationComponentSnapshotRestore(ctx, applicationUuid, componentUuid, uuid)



Restore a Snapshot copy of an application component.<br/> This endpoint is only supported for Maxdata template applications.<br/> Component Snapshot copies are essentially more granular application Snapshot copies. There is no difference beyond the scope of the operation. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`POST /application/applications/{application.uuid}/snapshots/{uuid}/restore`](#operations-application-application_snapshot_restore) * [`DOC /application`](#docs-application-overview) * [`DOC Asynchronous operations`](#docs-docs-Synchronous-and-asynchronous-operations) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**componentUuid** | **string**| Application Component UUID | 
**uuid** | **string**| Snapshot copy UUID | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationCreate

> JobLinkResponse ApplicationCreate(ctx, info)



Creates an application. ### Template properties The application APIs appear to be complex and long in this documentation because we document every possible template, of which there are currently 14. When creating an application, only a single template is used, so it is best to focus only on the template of interest. Other than the properties for the chosen template, only the `name` and `svm` of the application must be provided. The following three sections provided guidelines on using the properties of the templates, but the whole idea behind the templates is to automatically follow the best practices of the given application, so the only way to determine the exact list of required properties and default values is to dig in to the model section of the template. The templates are all top level properties of the application object with names matching the values returned by [`GET /application/templates`](#operations-application-application_template_collection_get). ### Required properties * `svm.uuid` or `svm.name` - The existing SVM in which to create the application. * `name` - The name for the application. * `<template>` - Properties for one template must be provided. In general, the following properties are required, however the naming of these may vary slightly from template to template.   * `name` - The generic templates require names for the components of the application. Other templates name the components automatically.   * `size` - This generally refers to the size of an application component, which may be spread across multiple underlying storage objects (volumes, LUNs, etc...).   * `igroup_name` - All SAN applications require an initiator group to be specified in some way.   * `os_type` - All SAN applications require an os_type to be specified in some way. Some templates refer to this as the `hypervisor`. ### Recommended optional properties * `<template>` - The following properties are available in some templates.   * `new_igroups.*` - SAN applications can use existing initiator groups or create new ones. When creating new initiator groups, `new_igroups.name` is required and the other properties may be used to fully specify the new initiator group. ### Default property values If not specified in POST, the follow default property values are assigned. It is recommended that most of these properties be provided explicitly rather than relying upon the defaults. The defaults are intended to make it as easy as possible to provision and connect to an application, but likely provide more access to the application than is necessary. * `template.name` - Defaults to match the `<template>` provided. If specified, the value of this property must match the provided template properties. * `<template>` - The majority of template properties have default values. The defaults may vary from template to template. See the model of each template for complete details. In general the following patterns are common across all template properties. The location of these properties varies from template to template.   * `storage_service.name` - _value_   * `nfs_access.host` - _0.0.0.0/0_   * `nfs_access.access` - _rw_   * `cifs_access.user_or_group` - _everyone_   * `cifs_access.access` - _full_access_   * `protection_type.local_rpo` - _hourly_ (Hourly Snapshot copies)   * `protection_type.remote_rpo` - _none_ (Not MetroCluster)   * `new_igroups.os_type` - Defaults to match the `os_type` provided for the application, but may need to be provided explicitly when using virtualization. ### Optional components A common pattern across many templates are objects that are optional, but once any property in the object is specified, other properties within the object become required. Many applications have optional components. For example, provisioning a database without a component to store the logs is supported. If the properties related to the logs are omitted, no storage will be provisioned for logs. But when the additional component is desired, the size is required. Specifying any other property of a component without specifying the size is not supported. In the model of each template, this is documented in the description of each property. When a `size` property is listed as optional, that means the component itself is optional, and the size should be specified to include that component in the application. ### POST body examples 1. Create a generic SAN application that exposes four LUNs to an existing initiator group, _igroup_1_.<br/>     ```     {       \"name\": \"app1\",       \"svm\": { \"name\": \"svm1\" },       \"san\": {         \"os_type\": \"linux\",         \"application_components\": [           { \"name\": \"component1\", \"total_size\": \"10GB\", \"lun_count\": 4, \"igroup_name\": \"igroup_1\" }         ]       }     }     ```     <br/> 2. Create an SQL application that can be accessed via initiator _iqn.2017-01.com.example:foo_ from a new initiator group, _igroup_2_.<br/>     ```     {       \"name\": \"app2\",       \"svm\": { \"name\": \"svm1\" },       \"sql_on_san\": {         \"db\": { \"size\": \"5GB\" },         \"log\": { \"size\": \"1GB\" },         \"temp_db\": { \"size\": \"2GB\" },         \"igroup_name\": \"igroup_2\",         \"new_igroups\": [           { \"name\": \"igroup_2\", \"initiators\": [ \"iqn.2017-01.com.example:foo\" ] }         ]       }     }     ```     <br/> 3. The following body creates the exact same SQL application, but manually provides all the defaults that were excluded from the previous call. Note: The model of a _sql_on_san_ application documents all these default values.<br/>     ```     {       \"name\": \"app3\",       \"svm\": { \"name\": \"svm1\" },       \"template\": { \"name\": \"sql_on_san\" },       \"sql_on_san\": {         \"os_type\": \"windows_2008\",         \"server_cores_count\": 8,         \"db\": { \"size\": \"5GB\", \"storage_service\": { \"name\": \"value\" } },         \"log\": { \"size\": \"1GB\", \"storage_service\": { \"name\": \"value\" } },         \"temp_db\": { \"size\": \"2GB\", \"storage_service\": { \"name\": \"value\" } },         \"igroup_name\": \"igroup_2\",         \"new_igroups\": [           {             \"name\": \"igroup_2\",             \"protocol\": \"mixed\",             \"os_type\": \"windows\",             \"initiators\": [ \"iqn.a.new.initiator\" ]           }         ],         \"protection_type\": { \"local_rpo\": \"none\" }       }     }     ``` ### Learn more * [`DOC /application`](#docs-application-overview) * [`DOC Asynchronous operations`](#docs-docs-Synchronous-and-asynchronous-operations) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**Application**](Application.md)| Info specification | 

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


## ApplicationDelete

> JobLinkResponse ApplicationDelete(ctx, uuid)



Deletes an application and all associated data. ### Warning - this deletes it all, including your data This deletes everything created with the application, including any volumes, LUNs, NFS export policies, CIFS shares, and initiator groups. Initiator groups are only destroyed if they were created as part of an application and are no longer in use by other applications. ### Learn more * [`DOC /application`](#docs-application-overview) * [`DOC Asynchronous operations`](#docs-docs-Synchronous-and-asynchronous-operations) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Application UUID | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationGet

> Application ApplicationGet(ctx, uuid, optional)



Retrieves an application ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `<template>` the property corresponding to the `template.name` of the application ### Property overview An application includes three main groups or properties. * Generic properties - such as the `name`, `template.name`, and `state` of the application. These properties are all inexpensive to retrieve and their meaning is consistent for every type of application. * `statistics.*` - application statistics report live usage data about the application and its components. Various space and IOPS details are included at both the application level and at a per component level. The application model includes a detailed description of each property. These properties are slightly more expensive than the generic properties because live data must be collected from every storage element in the application. * `<template>` - the property corresponding to the value of the `template.name` returns the contents of the application in the same layout that was used to provision the application. This information is very expensive to retrieve because it requires collecting information about all the storage and access settings for every element of the application. There are a few notable limitations to what can be returned in the `<template>` section:   * The `new_igroups` array of many SAN templates is not returned by GET. This property allows igroup creation in the same call that creates an application, but is not a property of the application itself. The `new_igroups` array is allowed during PATCH operations, but that does not modify the `new_igroups` of the application. It is another way to allow igroup creation while updating the application to use a different igroup.   * The `vdi_on_san` and `vdi_on_nas` `desktops.count` property is rounded to the nearest 1000 during creation, and is reported with that rounding applied.   * The `mongo_db_on_san` `dataset.element_count` property is rounded up to an even number, and is reported with that rounding applied.   * The `sql_on_san` and `sql_on_smb` `server_cores_count` property is limited to 8 for GET operations. Higher values are accepted by POST, but the impact of the `server_cores_count` property on the application layout currently reaches its limit at 8. ### Learn more * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Application UUID | 
 **optional** | ***ApplicationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Application**](application.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationModify

> JobLinkResponse ApplicationModify(ctx, uuid, optional)



Updates the properties of an application. ### Overview Similar to creating an application, modification is done using the template properties of an application. The `storage_service`, `size`, and `igroup_name` of an application may be modified. ### `storage_service` Storage service modifications are processed in place, meaning that the storage can not be moved to a location with more performance headroom to accommodate the request. If the current backing storage of the application is in a location that can support increased performance, the QoS policies associated with the application will be modified to allow it. If not, an error will be returned. A storage service modification to a lower tier of performance is always allowed, but the reverse modification may not be supported if the cluster is over provisioned and the cluster is unlikely to be able to fulfil the original storage service. ### `size` Size modifications are processed in a variety of ways depending on the type of application. For NAS applications, volumes are grown or new volumes are added. For SAN applications, LUNs are grown, new LUNs are added to existing volumes, or new LUNs are added to new volumes. If new storage elements are created, they can be found using the [`GET /application/applications/{application.uuid}/components`](#operations-application-application_component_collection_get) interface. The creation time of each storage object is included, and the newly created objects will use the same naming scheme as the previous objects. Resize follows the best practices associated with the type of application being expanded. Reducing the size of an application is not supported. ### `igroup_name` Modification of the igroup name allows an entire application to be mapped from one initiator group to another. Data access will be interrupted as the LUNs are unmapped from the original igroup and remapped to the new one. ### Application state During a modification, the `state` property of the application updates to indicate `modifying`. In `modifying` state, statistics are not available and Snapshot copy operations are not allowed. If the modification fails, it is possible for the application to be left in an inconsistent state, with the underlying ONTAP storage elements not matching across a component. When this occurs, the application is left in the `modifying` state until the command is either retried and succeeds or a call to restore the original state is successful. ### Examples 1. Change the storage service of the database of the Oracle application to _extreme_ and resize the redo logs to _100GB_.     ```     {       \"oracle_on_nfs\": {         \"db\": {           \"storage_service\": {             \"name\": \"extreme\"           }         },         \"redo_log\": {           \"size\": \"100GB\"         }       }     }     ```     <br/> 2. Change the storage service, size, and igroup of a generic application by component name.     ```     {       \"san\": {         \"application_components\": [           {             \"name\": \"component1\",             \"storage_service\": {               \"name\": \"value\"             }           },           {             \"name\": \"component2\",             \"size\": \"200GB\"           },           {             \"name\": \"component3\",             \"igroup_name\": \"igroup5\"           }         ]       }     }     ```     <br/> ### Learn more * [`DOC /application`](#docs-application-overview) * [`DOC Asynchronous operations`](#docs-docs-Synchronous-and-asynchronous-operations) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Application UUID | 
 **optional** | ***ApplicationModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of Application**](Application.md)| Info specification | 

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


## ApplicationSnapshotCollectionGet

> ApplicationSnapshotResponse ApplicationSnapshotCollectionGet(ctx, applicationUuid, optional)



Retrieves Snapshot copies of an application. ### Query examples The following query returns all Snapshot copies from May 4, 2017 EST. For readability, the colon (`:`) is left in this example. For an actual call, they should be escaped as `%3A`.<br/><br/> ``` GET /application/applications/{application.uuid}/snapshots?create_time=2017-05-04T00:00:00-05:00..2017-05-04T23:59:59-05:00 ``` <br/>The following query returns all Snapshot copies that have been flagged as _application consistent_.<br/><br/> ``` GET /application/applications/{application.uuid}/snapshots?consistency_type=application ``` ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
 **optional** | ***ApplicationSnapshotCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationSnapshotCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uuid** | **optional.String**| Filter by UUID | 
 **name** | **optional.String**| Filter by name | 
 **consistencyType** | **optional.String**| Filter by consistency_type | 
 **componentsName** | **optional.String**| Filter by components.name | 
 **componentsUuid** | **optional.String**| Filter by components.uuid | 
 **comment** | **optional.String**| Filter by comment | 
 **createTime** | **optional.String**| Filter by create_time | 
 **isPartial** | **optional.String**| Filter by is_partial | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ApplicationSnapshotResponse**](application_snapshot_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationSnapshotCreate

> JobLinkResponse ApplicationSnapshotCreate(ctx, applicationUuid, optional)



Creates a Snapshot copy of the application. ### Required properties * `name` ### Recommended optional properties * `consistency_type` - Track whether this snapshot is _application_ or _crash_ consistent. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
 **optional** | ***ApplicationSnapshotCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationSnapshotCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of ApplicationSnapshot**](ApplicationSnapshot.md)| Info specification | 

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


## ApplicationSnapshotDelete

> JobLinkResponse ApplicationSnapshotDelete(ctx, applicationUuid, uuid)



Delete a Snapshot copy of an application ### Query examples Individual Snapshot copies can be destroyed with no query parameters, or a range of Snapshot copies can be destroyed at one time using a query.<br/> The following query deletes all application Snapshot copies created before May 4, 2017<br/><br/> ``` DELETE /application/applications/{application.uuid}/snapshots?create_time=<2017-05-04T00:00:00-05:00 ```  ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**uuid** | **string**| Snapshot copy UUID | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationSnapshotGet

> ApplicationSnapshot ApplicationSnapshotGet(ctx, applicationUuid, uuid, optional)



Retrieve a Snapshot copy of an application component.<br/> This endpoint is only supported for Maxdata template applications.<br/> Component Snapshot copies are essentially more granular application Snapshot copies. There is no difference beyond the scope of the operation. ### Learn more * [`DOC /application/applications/{application.uuid}/snapshots`](#docs-application-application_applications_{application.uuid}_snapshots) * [`GET /application/applications/{uuid}/snapshots`](#operations-application-application_snapshot_create) * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**uuid** | **string**| Snapshot copy UUID | 
 **optional** | ***ApplicationSnapshotGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationSnapshotGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ApplicationSnapshot**](application_snapshot.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationSnapshotRestore

> JobLinkResponse ApplicationSnapshotRestore(ctx, applicationUuid, uuid)



Restore an application snapshot<br/> Restoring an application Snapshot copy reverts all storage elements in the Snapshot copy to the state in which the Snapshot copy was in when the Snapshot copy was taken. This restoration does not apply to access settings that might have changed since the Snapshot copy was created. ### Learn more * [`DOC /application`](#docs-application-overview) * [`DOC Asynchronous operations`](#docs-docs-Synchronous-and-asynchronous-operations) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationUuid** | **string**| Application UUID | 
**uuid** | **string**| Snapshot copy UUID | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationTemplateCollectionGet

> ApplicationTemplateResponse ApplicationTemplateCollectionGet(ctx, optional)



Retrieves application templates. ### Query examples The most useful queries on this API allows searches by name or protocol access. The following query returns all templates that are used to provision an Oracle application.<br/><br/> ``` GET /application/templates?name=ora* ``` <br/>Similarly, the following query returns all templates that support SAN access.<br/><br/> ``` GET /application/templates?protocol=san ``` ### Learn more * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationTemplateCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationTemplateCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **protocol** | **optional.String**| Filter by protocol | 
 **description** | **optional.String**| Filter by description | 
 **missingPrerequisites** | **optional.String**| Filter by missing_prerequisites | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ApplicationTemplateResponse**](application_template_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationTemplateGet

> ApplicationTemplate ApplicationTemplateGet(ctx, name, optional)



Retrieves an application template. ### Template properties Each application template has a set of properties. These properties are always nested under a property with the same name as the template. For example, when using the `mongo_db_on_san` template, the properties are found nested inside the `mongo_db_on_san` property. The properties nested under the template property are all specific to the template. The model for the application template object includes all the available templates, but only the object that corresponds to the template's name is returned, and only one is provided in any application API.<br/> The model of each template includes a description of each property and its allowed values or usage. Default values are also indicated when available. The template properties returned by this API include an example value for each property. ### Template prerequisites Each template has a set of prerequisites required for its use. If any of these prerequisites are not met, the `missing_prerequisites` property indicates which prerequisite is missing. ### Learn more * [`DOC /application`](#docs-application-overview) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Template Name | 
 **optional** | ***ApplicationTemplateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationTemplateGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ApplicationTemplate**](application_template.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

