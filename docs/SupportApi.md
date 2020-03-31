# \SupportApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutosupportCreate**](SupportApi.md#AutosupportCreate) | **Post** /support/autosupport/messages | 
[**AutosupportGet**](SupportApi.md#AutosupportGet) | **Get** /support/autosupport | 
[**AutosupportMessageCollectionGet**](SupportApi.md#AutosupportMessageCollectionGet) | **Get** /support/autosupport/messages | 
[**AutosupportModify**](SupportApi.md#AutosupportModify) | **Patch** /support/autosupport | 
[**ConfigurationBackupGet**](SupportApi.md#ConfigurationBackupGet) | **Get** /support/configuration-backup | 
[**ConfigurationBackupModify**](SupportApi.md#ConfigurationBackupModify) | **Patch** /support/configuration-backup | 
[**EmsConfigGet**](SupportApi.md#EmsConfigGet) | **Get** /support/ems | 
[**EmsConfigModify**](SupportApi.md#EmsConfigModify) | **Patch** /support/ems | 
[**EmsDestinationCollectionGet**](SupportApi.md#EmsDestinationCollectionGet) | **Get** /support/ems/destinations | 
[**EmsDestinationCreate**](SupportApi.md#EmsDestinationCreate) | **Post** /support/ems/destinations | 
[**EmsDestinationDelete**](SupportApi.md#EmsDestinationDelete) | **Delete** /support/ems/destinations/{name} | 
[**EmsDestinationGet**](SupportApi.md#EmsDestinationGet) | **Get** /support/ems/destinations/{name} | 
[**EmsDestinationModify**](SupportApi.md#EmsDestinationModify) | **Patch** /support/ems/destinations/{name} | 
[**EmsEventCollectionGet**](SupportApi.md#EmsEventCollectionGet) | **Get** /support/ems/events | 
[**EmsFilterCollectionGet**](SupportApi.md#EmsFilterCollectionGet) | **Get** /support/ems/filters | 
[**EmsFilterCreate**](SupportApi.md#EmsFilterCreate) | **Post** /support/ems/filters | 
[**EmsFilterDelete**](SupportApi.md#EmsFilterDelete) | **Delete** /support/ems/filters/{name} | 
[**EmsFilterGet**](SupportApi.md#EmsFilterGet) | **Get** /support/ems/filters/{name} | 
[**EmsFilterModify**](SupportApi.md#EmsFilterModify) | **Patch** /support/ems/filters/{name} | 
[**EmsFilterRuleCollectionGet**](SupportApi.md#EmsFilterRuleCollectionGet) | **Get** /support/ems/filters/{name}/rules | 
[**EmsFilterRuleDelete**](SupportApi.md#EmsFilterRuleDelete) | **Delete** /support/ems/filters/{name}/rules/{index} | 
[**EmsFilterRuleGet**](SupportApi.md#EmsFilterRuleGet) | **Get** /support/ems/filters/{name}/rules/{index} | 
[**EmsFilterRuleModify**](SupportApi.md#EmsFilterRuleModify) | **Patch** /support/ems/filters/{name}/rules/{index} | 
[**EmsFiltersRulesCreate**](SupportApi.md#EmsFiltersRulesCreate) | **Post** /support/ems/filters/{name}/rules | 
[**EmsMessageCollectionGet**](SupportApi.md#EmsMessageCollectionGet) | **Get** /support/ems/messages | 



## AutosupportCreate

> InlineResponse201 AutosupportCreate(ctx, optional)



Generates and sends an AutoSupport message with the provided input parameters.<p/> Important note: * By default, the response is an empty object. If `return_records=true` is passed in the request, the response includes information about the node and the index of the invoked AutoSupport message. ### Recommended optional properties * `message` - Message included in the AutoSupport subject. Use to identify the generated AutoSupport message. ### Default property values If not specified in POST, the following default property values are assigned: * `type` - _all_ * `node.name` or `node.uuid` - Not specifying any of these properties invokes the AutoSupport on all nodes in the cluster. ### Related ONTAP commands * `system node autosupport invoke` ### Learn more * [`DOC /support/autosupport/messages`](#docs-support-support_autosupport_messages) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutosupportCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutosupportCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnRecords** | **optional.Bool**| The default is false.  If set to true, the records are returned. | [default to false]
 **info** | [**optional.Interface of AutosupportMessage**](AutosupportMessage.md)| Info specification | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutosupportGet

> Autosupport AutosupportGet(ctx, optional)



Retrieves the AutoSupport configuration of the cluster and if requested, returns any connectivity issues with the AutoSupport configuration.<p/> </br>Important note: * The **issues** field consists of a list of objects containing details of the node that has a connectivity issue, the issue description, and corrective action you can take to address the issue. When not empty, this indicates a connection issue to the **HTTP/S**, **SMTP**, or **AutoSupport On Demand** server. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `issues` ### Related ONTAP commands * `system node autosupport show -instance` * `system node autosupport check show-details` ### Learn more * [`DOC /support/autosupport`](#docs-support-support_autosupport) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutosupportGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutosupportGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Autosupport**](autosupport.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutosupportMessageCollectionGet

> AutosupportMessageResponse AutosupportMessageCollectionGet(ctx, optional)



Retrieves AutoSupport message history from all nodes in the cluster.<p/> There can be a short delay on invoked AutoSupport messages showing in history, dependent on processing of other AutoSupports in the queue. ### Related ONTAP commands * `system node autosupport history show` ### Learn more * [`DOC /support/autosupport/messages`](#docs-support-support_autosupport_messages) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutosupportMessageCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutosupportMessageCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **destination** | **optional.String**| Filter by destination | 
 **state** | **optional.String**| Filter by state | 
 **generatedOn** | **optional.String**| Filter by generated_on | 
 **index** | **optional.Int32**| Filter by index | 
 **subject** | **optional.String**| Filter by subject | 
 **errorMessage** | **optional.String**| Filter by error.message | 
 **errorCode** | **optional.Int32**| Filter by error.code | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]

### Return type

[**AutosupportMessageResponse**](autosupport_message_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutosupportModify

> AutosupportModify(ctx, optional)



Updates the AutoSupport configuration for the entire cluster. ### Related ONTAP commands * `system node autosupport modify` ### Learn more * [`DOC /support/autosupport`](#docs-support-support_autosupport) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutosupportModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutosupportModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Autosupport**](Autosupport.md)| Info specification | 

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


## ConfigurationBackupGet

> ConfigurationBackup ConfigurationBackupGet(ctx, optional)



Retrieves the cluster configuration backup information. ### Learn more * [`DOC /support/configuration-backup`](#docs-support-support_configuration-backup)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationBackupGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigurationBackupGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ConfigurationBackup**](configuration_backup.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationBackupModify

> ConfigurationBackupModify(ctx, optional)



Updates the cluster configuration backup information.  ### Learn more * [`DOC /support/configuration-backup`](#docs-support-support_configuration-backup)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationBackupModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigurationBackupModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ConfigurationBackup**](ConfigurationBackup.md)| Cluster configuration backup information | 

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


## EmsConfigGet

> EmsConfig EmsConfigGet(ctx, optional)



Retrieves the EMS configuration. ### Related ONTAP commands * `event config show`  ### Learn more * [`DOC /support/ems`](#docs-support-support_ems)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsConfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsConfigGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**EmsConfig**](ems_config.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsConfigModify

> EmsConfigModify(ctx, optional)



Updates the EMS configuration. ### Related ONTAP commands * `event config modify`  ### Learn more * [`DOC /support/ems`](#docs-support-support_ems)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsConfigModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsConfigModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of EmsConfig**](EmsConfig.md)| Information specification | 

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


## EmsDestinationCollectionGet

> EmsDestinationResponse EmsDestinationCollectionGet(ctx, optional)



Retrieves a collection of event destinations. ### Related ONTAP commands * `event notification destination show` * `event notification show`  ### Learn more * [`DOC /support/ems/destinations`](#docs-support-support_ems_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsDestinationCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsDestinationCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersName** | **optional.String**| Filter by filters.name | 
 **destination** | **optional.String**| Filter by destination | 
 **type_** | **optional.String**| Filter by type | 
 **name** | **optional.String**| Filter by name | 
 **certificateSerialNumber** | **optional.String**| Filter by certificate.serial_number | 
 **certificateCa** | **optional.String**| Filter by certificate.ca | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**EmsDestinationResponse**](ems_destination_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsDestinationCreate

> EmsDestinationResponse EmsDestinationCreate(ctx, optional)



Creates an event destination. ### Required properties * `name` - String that uniquely identifies the destination. * `type` - Type of destination that is to be created. * `destination` - String that identifies the destination. The contents of this field changes depending on type. ### Recommended optional properties * `filters.name` - List of filter names that should direct to this destination. * `certificate` - When specifying a rest api destination, a client certificate can be provided. ### Related ONTAP commands * `event notification destination create` * `event notification create`  ### Learn more * [`DOC /support/ems/destinations`](#docs-support-support_ems_destinations)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsDestinationCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsDestinationCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of EmsDestination**](EmsDestination.md)| Information specification | 

### Return type

[**EmsDestinationResponse**](ems_destination_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsDestinationDelete

> EmsDestinationDelete(ctx, name)



Deletes an event destination. ### Related ONTAP commands * `event notification destination delete` * `event notification delete`  ### Learn more * [`DOC /support/ems/destinations/{name}`](#docs-support-support_ems_destinations_{name})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Destination name | 

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


## EmsDestinationGet

> EmsDestination EmsDestinationGet(ctx, name, optional)



Retrieves event destinations. ### Related ONTAP commands * `event notification destination show` * `event notification show`  ### Learn more * [`DOC /support/ems/destinations/{name}`](#docs-support-support_ems_destinations_{name})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Destination name | 
 **optional** | ***EmsDestinationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsDestinationGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**EmsDestination**](ems_destination.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsDestinationModify

> EmsDestination EmsDestinationModify(ctx, name, optional)



Updates an event destination. ### Recommended optional properties * `filters.name` - New list of filters that should direct to this destination. The existing list is discarded. * `certificate` - New certificate parameters when the destination type is `rest api`. ### Related ONTAP commands * `event notification destination modify` * `event notification modify`  ### Learn more * [`DOC /support/ems/destinations/{name}`](#docs-support-support_ems_destinations_{name})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Destination name | 
 **optional** | ***EmsDestinationModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsDestinationModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of EmsDestination**](EmsDestination.md)| Information specification | 

### Return type

[**EmsDestination**](ems_destination.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsEventCollectionGet

> EmsEventResponse EmsEventCollectionGet(ctx, optional)



Retrieves a collection of observed events. ### Related ONTAP commands * `event log show`  ### Learn more * [`DOC /support/ems/events`](#docs-support-support_ems_events)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsEventCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsEventCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **logMessage** | **optional.String**| Filter by log_message | 
 **messageSeverity** | **optional.String**| Filter by message.severity | 
 **messageName** | **optional.String**| Filter by message.name | 
 **parametersValue** | **optional.String**| Filter by parameters.value | 
 **parametersName** | **optional.String**| Filter by parameters.name | 
 **source** | **optional.String**| Filter by source | 
 **time** | **optional.String**| Filter by time | 
 **index** | **optional.Int32**| Filter by index | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**EmsEventResponse**](ems_event_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsFilterCollectionGet

> EmsFilterResponse EmsFilterCollectionGet(ctx, optional)



Retrieves a collection of event filters. ### Related ONTAP commands * `event filter show`  ### Learn more * [`DOC /support/ems/filters`](#docs-support-support_ems_filters)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsFilterCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **rulesIndex** | **optional.Int32**| Filter by rules.index | 
 **rulesMessageCriteriaSnmpTrapTypes** | **optional.String**| Filter by rules.message_criteria.snmp_trap_types | 
 **rulesMessageCriteriaNamePattern** | **optional.String**| Filter by rules.message_criteria.name_pattern | 
 **rulesMessageCriteriaSeverities** | **optional.String**| Filter by rules.message_criteria.severities | 
 **rulesType** | **optional.String**| Filter by rules.type | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**EmsFilterResponse**](ems_filter_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsFilterCreate

> EmsFilterResponse EmsFilterCreate(ctx, optional)



Creates an event filter. ### Required properties * `name` - String that uniquely identifies the filter. ### Recommended optional properties * `rules` - List of criteria which is used to match a filter with an event. ### Related ONTAP commands * `event filter create`  ### Learn more * [`DOC /support/ems/filters`](#docs-support-support_ems_filters)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsFilterCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of EmsFilter**](EmsFilter.md)| Information specification | 

### Return type

[**EmsFilterResponse**](ems_filter_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsFilterDelete

> EmsFilterDelete(ctx, name)



Deletes an event filter. ### Related ONTAP commands * `event filter delete`  ### Learn more * [`DOC /support/ems/filters/{name}`](#docs-support-support_ems_filters_{name})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 

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


## EmsFilterGet

> EmsFilter EmsFilterGet(ctx, name, optional)



Retrieves an event filter. ### Related ONTAP commands * `event filter show`  ### Learn more * [`DOC /support/ems/filters/{name}`](#docs-support-support_ems_filters_{name})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 
 **optional** | ***EmsFilterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**EmsFilter**](ems_filter.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsFilterModify

> EmsFilterModify(ctx, name, optional)



Updates an event filter. ### Recommended optional properties * `new_name` - New string that uniquely identifies a filter. * `rules` - New list of criteria used to match the filter with an event. The existing list is discarded. ### Related ONTAP commands * `event filter create` * `event filter delete` * `event filter rename` * `event filter rule add` * `event filter rule delete` * `event filter rule reorder`  ### Learn more * [`DOC /support/ems/filters/{name}`](#docs-support-support_ems_filters_{name})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 
 **optional** | ***EmsFilterModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newName** | **optional.String**| New filter name for renames. Valid in PATCH. | 
 **info** | [**optional.Interface of EmsFilter**](EmsFilter.md)| Information specification | 

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


## EmsFilterRuleCollectionGet

> EmsFilterRuleResponse EmsFilterRuleCollectionGet(ctx, name, optional)



Retrieves event filter rules. ### Related ONTAP commands * `event filter show`  ### Learn more * [`DOC /support/ems/filters/{name}/rules`](#docs-support-support_ems_filters_{name}_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter Name | 
 **optional** | ***EmsFilterRuleCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterRuleCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **index** | **optional.Int32**| Filter by index | 
 **messageCriteriaSnmpTrapTypes** | **optional.String**| Filter by message_criteria.snmp_trap_types | 
 **messageCriteriaNamePattern** | **optional.String**| Filter by message_criteria.name_pattern | 
 **messageCriteriaSeverities** | **optional.String**| Filter by message_criteria.severities | 
 **type_** | **optional.String**| Filter by type | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**EmsFilterRuleResponse**](ems_filter_rule_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsFilterRuleDelete

> EmsFilterRuleDelete(ctx, name, index)



Deletes an event filter rule. ### Related ONTAP commands * `event filter rule delete`  ### Learn more * [`DOC /support/ems/filters/{name}/rules/{index}`](#docs-support-support_ems_filters_{name}_rules_{index})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 
**index** | **string**| Filter index | 

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


## EmsFilterRuleGet

> EmsFilterRule EmsFilterRuleGet(ctx, name, index, optional)



Retrieves an event filter rule. ### Related ONTAP commands * `event filter show`  ### Learn more * [`DOC /support/ems/filters/{name}/rules/{index}`](#docs-support-support_ems_filters_{name}_rules_{index})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 
**index** | **string**| Filter index | 
 **optional** | ***EmsFilterRuleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterRuleGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**EmsFilterRule**](ems_filter_rule.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsFilterRuleModify

> EmsFilterRuleModify(ctx, name, index, optional)



Updates an event filter rule. ### Recommended optional properties * `message_criteria` - New criteria on which a rule is to match an event. ### Related ONTAP commands * `event filter rule add` * `event filter rule delete`  ### Learn more * [`DOC /support/ems/filters/{name}/rules/{index}`](#docs-support-support_ems_filters_{name}_rules_{index})

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 
**index** | **string**| Filter index | 
 **optional** | ***EmsFilterRuleModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFilterRuleModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newIndex** | **optional.Int32**| New position for the filter rule index | 
 **info** | [**optional.Interface of EmsFilterRule**](EmsFilterRule.md)| Information specification | 

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


## EmsFiltersRulesCreate

> EmsFilterRuleResponse EmsFiltersRulesCreate(ctx, name, optional)



Creates an event filter rule. ### Required properties * `message_criteria` - Criteria on which a rule is to match an event. ### Recommended optional properties * `index` - One-based position index of the new rule. ### Related ONTAP commands * `event filter rule add`  ### Learn more * [`DOC /support/ems/filters/{name}/rules`](#docs-support-support_ems_filters_{name}_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Filter name | 
 **optional** | ***EmsFiltersRulesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsFiltersRulesCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of EmsFilterRule**](EmsFilterRule.md)| Information specification | 

### Return type

[**EmsFilterRuleResponse**](ems_filter_rule_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmsMessageCollectionGet

> EmsMessageResponse EmsMessageCollectionGet(ctx, optional)



Retrieves the event catalog definitions. ### Related ONTAP commands * `event catalog show`  ### Learn more * [`DOC /support/ems/messages`](#docs-support-support_ems_messages)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmsMessageCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmsMessageCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **severity** | **optional.String**| Filter by severity | 
 **deprecated** | **optional.Bool**| Filter by deprecated | 
 **description** | **optional.String**| Filter by description | 
 **correctiveAction** | **optional.String**| Filter by corrective_action | 
 **name** | **optional.String**| Filter by name | 
 **snmpTrapType** | **optional.String**| Filter by snmp_trap_type | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**EmsMessageResponse**](ems_message_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

