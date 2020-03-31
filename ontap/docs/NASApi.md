# \NASApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditCollectionGet**](NASApi.md#AuditCollectionGet) | **Get** /protocols/audit | 
[**AuditCreate**](NASApi.md#AuditCreate) | **Post** /protocols/audit | 
[**AuditDelete**](NASApi.md#AuditDelete) | **Delete** /protocols/audit/{svm.uuid} | 
[**AuditGet**](NASApi.md#AuditGet) | **Get** /protocols/audit/{svm.uuid} | 
[**AuditModify**](NASApi.md#AuditModify) | **Patch** /protocols/audit/{svm.uuid} | 
[**CifsHomedirSearchPathGet**](NASApi.md#CifsHomedirSearchPathGet) | **Get** /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index} | 
[**CifsSearchPathCollectionGet**](NASApi.md#CifsSearchPathCollectionGet) | **Get** /protocols/cifs/home-directory/search-paths | 
[**CifsSearchPathCreate**](NASApi.md#CifsSearchPathCreate) | **Post** /protocols/cifs/home-directory/search-paths | 
[**CifsSearchPathDelete**](NASApi.md#CifsSearchPathDelete) | **Delete** /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index} | 
[**CifsSearchPathModify**](NASApi.md#CifsSearchPathModify) | **Patch** /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index} | 
[**CifsServiceCollectionGet**](NASApi.md#CifsServiceCollectionGet) | **Get** /protocols/cifs/services | 
[**CifsServiceCreate**](NASApi.md#CifsServiceCreate) | **Post** /protocols/cifs/services | 
[**CifsServiceDelete**](NASApi.md#CifsServiceDelete) | **Delete** /protocols/cifs/services/{svm.uuid} | 
[**CifsServiceGet**](NASApi.md#CifsServiceGet) | **Get** /protocols/cifs/services/{svm.uuid} | 
[**CifsServiceModify**](NASApi.md#CifsServiceModify) | **Patch** /protocols/cifs/services/{svm.uuid} | 
[**CifsShareAclCollectionGet**](NASApi.md#CifsShareAclCollectionGet) | **Get** /protocols/cifs/shares/{svm.uuid}/{share}/acls | 
[**CifsShareAclCreate**](NASApi.md#CifsShareAclCreate) | **Post** /protocols/cifs/shares/{svm.uuid}/{share}/acls | 
[**CifsShareAclDelete**](NASApi.md#CifsShareAclDelete) | **Delete** /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type} | 
[**CifsShareAclGet**](NASApi.md#CifsShareAclGet) | **Get** /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type} | 
[**CifsShareAclModify**](NASApi.md#CifsShareAclModify) | **Patch** /protocols/cifs/shares/{svm.uuid}/{share}/acls/{user_or_group}/{type} | 
[**CifsShareCollectionGet**](NASApi.md#CifsShareCollectionGet) | **Get** /protocols/cifs/shares | 
[**CifsShareCreate**](NASApi.md#CifsShareCreate) | **Post** /protocols/cifs/shares | 
[**CifsShareDelete**](NASApi.md#CifsShareDelete) | **Delete** /protocols/cifs/shares/{svm.uuid}/{name} | 
[**CifsShareGet**](NASApi.md#CifsShareGet) | **Get** /protocols/cifs/shares/{svm.uuid}/{name} | 
[**CifsShareModify**](NASApi.md#CifsShareModify) | **Patch** /protocols/cifs/shares/{svm.uuid}/{name} | 
[**CifsSymlinkMappingCollectionGet**](NASApi.md#CifsSymlinkMappingCollectionGet) | **Get** /protocols/cifs/unix-symlink-mapping | 
[**CifsSymlinkMappingCreate**](NASApi.md#CifsSymlinkMappingCreate) | **Post** /protocols/cifs/unix-symlink-mapping | 
[**CifsSymlinkMappingDelete**](NASApi.md#CifsSymlinkMappingDelete) | **Delete** /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path} | 
[**CifsSymlinkMappingGet**](NASApi.md#CifsSymlinkMappingGet) | **Get** /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path} | 
[**CifsSymlinkMappingModify**](NASApi.md#CifsSymlinkMappingModify) | **Patch** /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path} | 
[**ExportPolicyCollectionGet**](NASApi.md#ExportPolicyCollectionGet) | **Get** /protocols/nfs/export-policies | 
[**ExportPolicyCreate**](NASApi.md#ExportPolicyCreate) | **Post** /protocols/nfs/export-policies | 
[**ExportPolicyDelete**](NASApi.md#ExportPolicyDelete) | **Delete** /protocols/nfs/export-policies/{id} | 
[**ExportPolicyGet**](NASApi.md#ExportPolicyGet) | **Get** /protocols/nfs/export-policies/{id} | 
[**ExportPolicyModify**](NASApi.md#ExportPolicyModify) | **Patch** /protocols/nfs/export-policies/{id} | 
[**ExportRuleClientsCreate**](NASApi.md#ExportRuleClientsCreate) | **Post** /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients | 
[**ExportRuleClientsDelete**](NASApi.md#ExportRuleClientsDelete) | **Delete** /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients/{match} | 
[**ExportRuleClientsGet**](NASApi.md#ExportRuleClientsGet) | **Get** /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients | 
[**ExportRuleCollectionGet**](NASApi.md#ExportRuleCollectionGet) | **Get** /protocols/nfs/export-policies/{policy.id}/rules | 
[**ExportRuleCreate**](NASApi.md#ExportRuleCreate) | **Post** /protocols/nfs/export-policies/{policy.id}/rules | 
[**ExportRuleDelete**](NASApi.md#ExportRuleDelete) | **Delete** /protocols/nfs/export-policies/{policy.id}/rules/{index} | 
[**ExportRuleGet**](NASApi.md#ExportRuleGet) | **Get** /protocols/nfs/export-policies/{policy.id}/rules/{index} | 
[**ExportRuleModify**](NASApi.md#ExportRuleModify) | **Patch** /protocols/nfs/export-policies/{policy.id}/rules/{index} | 
[**FpolicyCollectionGet**](NASApi.md#FpolicyCollectionGet) | **Get** /protocols/fpolicy | 
[**FpolicyCreate**](NASApi.md#FpolicyCreate) | **Post** /protocols/fpolicy | 
[**FpolicyDelete**](NASApi.md#FpolicyDelete) | **Delete** /protocols/fpolicy/{svm.uuid} | 
[**FpolicyEngineCollectionGet**](NASApi.md#FpolicyEngineCollectionGet) | **Get** /protocols/fpolicy/{svm.uuid}/engines | 
[**FpolicyEngineCreate**](NASApi.md#FpolicyEngineCreate) | **Post** /protocols/fpolicy/{svm.uuid}/engines | 
[**FpolicyEngineDelete**](NASApi.md#FpolicyEngineDelete) | **Delete** /protocols/fpolicy/{svm.uuid}/engines/{name} | 
[**FpolicyEngineGet**](NASApi.md#FpolicyEngineGet) | **Get** /protocols/fpolicy/{svm.uuid}/engines/{name} | 
[**FpolicyEngineModify**](NASApi.md#FpolicyEngineModify) | **Patch** /protocols/fpolicy/{svm.uuid}/engines/{name} | 
[**FpolicyEventCollectionGet**](NASApi.md#FpolicyEventCollectionGet) | **Get** /protocols/fpolicy/{svm.uuid}/events | 
[**FpolicyEventCreate**](NASApi.md#FpolicyEventCreate) | **Post** /protocols/fpolicy/{svm.uuid}/events | 
[**FpolicyEventDelete**](NASApi.md#FpolicyEventDelete) | **Delete** /protocols/fpolicy/{svm.uuid}/events/{name} | 
[**FpolicyEventModify**](NASApi.md#FpolicyEventModify) | **Patch** /protocols/fpolicy/{svm.uuid}/events/{name} | 
[**FpolicyEventsGet**](NASApi.md#FpolicyEventsGet) | **Get** /protocols/fpolicy/{svm.uuid}/events/{name} | 
[**FpolicyGet**](NASApi.md#FpolicyGet) | **Get** /protocols/fpolicy/{svm.uuid} | 
[**FpolicyPolicyCollectionGet**](NASApi.md#FpolicyPolicyCollectionGet) | **Get** /protocols/fpolicy/{svm.uuid}/policies | 
[**FpolicyPolicyCreate**](NASApi.md#FpolicyPolicyCreate) | **Post** /protocols/fpolicy/{svm.uuid}/policies | 
[**FpolicyPolicyDelete**](NASApi.md#FpolicyPolicyDelete) | **Delete** /protocols/fpolicy/{svm.uuid}/policies/{name} | 
[**FpolicyPolicyGet**](NASApi.md#FpolicyPolicyGet) | **Get** /protocols/fpolicy/{svm.uuid}/policies/{name} | 
[**FpolicyPolicyModify**](NASApi.md#FpolicyPolicyModify) | **Patch** /protocols/fpolicy/{svm.uuid}/policies/{name} | 
[**KerberosInterfaceCollectionGet**](NASApi.md#KerberosInterfaceCollectionGet) | **Get** /protocols/nfs/kerberos/interfaces | 
[**KerberosInterfaceGet**](NASApi.md#KerberosInterfaceGet) | **Get** /protocols/nfs/kerberos/interfaces/{uuid} | 
[**KerberosInterfaceModify**](NASApi.md#KerberosInterfaceModify) | **Patch** /protocols/nfs/kerberos/interfaces/{uuid} | 
[**KerberosRealmCollectionGet**](NASApi.md#KerberosRealmCollectionGet) | **Get** /protocols/nfs/kerberos/realms | 
[**KerberosRealmCreate**](NASApi.md#KerberosRealmCreate) | **Post** /protocols/nfs/kerberos/realms | 
[**KerberosRealmDelete**](NASApi.md#KerberosRealmDelete) | **Delete** /protocols/nfs/kerberos/realms/{svm.uuid}/{name} | 
[**KerberosRealmGet**](NASApi.md#KerberosRealmGet) | **Get** /protocols/nfs/kerberos/realms/{svm.uuid}/{name} | 
[**KerberosRealmModify**](NASApi.md#KerberosRealmModify) | **Patch** /protocols/nfs/kerberos/realms/{svm.uuid}/{name} | 
[**NfsCollectionGet**](NASApi.md#NfsCollectionGet) | **Get** /protocols/nfs/services | 
[**NfsCreate**](NASApi.md#NfsCreate) | **Post** /protocols/nfs/services | 
[**NfsDelete**](NASApi.md#NfsDelete) | **Delete** /protocols/nfs/services/{svm.uuid} | 
[**NfsGet**](NASApi.md#NfsGet) | **Get** /protocols/nfs/services/{svm.uuid} | 
[**NfsModify**](NASApi.md#NfsModify) | **Patch** /protocols/nfs/services/{svm.uuid} | 
[**VscanCollectionGet**](NASApi.md#VscanCollectionGet) | **Get** /protocols/vscan | 
[**VscanConfigDelete**](NASApi.md#VscanConfigDelete) | **Delete** /protocols/vscan/{svm.uuid} | 
[**VscanCreate**](NASApi.md#VscanCreate) | **Post** /protocols/vscan | 
[**VscanGet**](NASApi.md#VscanGet) | **Get** /protocols/vscan/{svm.uuid} | 
[**VscanModify**](NASApi.md#VscanModify) | **Patch** /protocols/vscan/{svm.uuid} | 
[**VscanOnAccessCreate**](NASApi.md#VscanOnAccessCreate) | **Post** /protocols/vscan/{svm.uuid}/on-access-policies | 
[**VscanOnAccessDelete**](NASApi.md#VscanOnAccessDelete) | **Delete** /protocols/vscan/{svm.uuid}/on-access-policies/{name} | 
[**VscanOnAccessGet**](NASApi.md#VscanOnAccessGet) | **Get** /protocols/vscan/{svm.uuid}/on-access-policies/{name} | 
[**VscanOnAccessModify**](NASApi.md#VscanOnAccessModify) | **Patch** /protocols/vscan/{svm.uuid}/on-access-policies/{name} | 
[**VscanOnAccessPolicyCollectionGet**](NASApi.md#VscanOnAccessPolicyCollectionGet) | **Get** /protocols/vscan/{svm.uuid}/on-access-policies | 
[**VscanOnDemandCreate**](NASApi.md#VscanOnDemandCreate) | **Post** /protocols/vscan/{svm.uuid}/on-demand-policies | 
[**VscanOnDemandDelete**](NASApi.md#VscanOnDemandDelete) | **Delete** /protocols/vscan/{svm.uuid}/on-demand-policies/{name} | 
[**VscanOnDemandGet**](NASApi.md#VscanOnDemandGet) | **Get** /protocols/vscan/{svm.uuid}/on-demand-policies/{name} | 
[**VscanOnDemandModify**](NASApi.md#VscanOnDemandModify) | **Patch** /protocols/vscan/{svm.uuid}/on-demand-policies/{name} | 
[**VscanOnDemandPolicyCollectionGet**](NASApi.md#VscanOnDemandPolicyCollectionGet) | **Get** /protocols/vscan/{svm.uuid}/on-demand-policies | 
[**VscanScannerCollectionGet**](NASApi.md#VscanScannerCollectionGet) | **Get** /protocols/vscan/{svm.uuid}/scanner-pools | 
[**VscanScannerCreate**](NASApi.md#VscanScannerCreate) | **Post** /protocols/vscan/{svm.uuid}/scanner-pools | 
[**VscanScannerDelete**](NASApi.md#VscanScannerDelete) | **Delete** /protocols/vscan/{svm.uuid}/scanner-pools/{name} | 
[**VscanScannerGet**](NASApi.md#VscanScannerGet) | **Get** /protocols/vscan/{svm.uuid}/scanner-pools/{name} | 
[**VscanScannerModify**](NASApi.md#VscanScannerModify) | **Patch** /protocols/vscan/{svm.uuid}/scanner-pools/{name} | 
[**VscanServerStatusGet**](NASApi.md#VscanServerStatusGet) | **Get** /protocols/vscan/server-status | 



## AuditCollectionGet

> AuditResponse AuditCollectionGet(ctx, optional)



Retrieves audit configurations. ### Related ONTAP commands * `vserver audit show` ### Learn more * [`DOC /protocols/audit`](#docs-NAS-protocols_audit) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logRetentionDuration** | **optional.String**| Filter by log.retention.duration | 
 **logRetentionCount** | **optional.Int32**| Filter by log.retention.count | 
 **logRotationSize** | **optional.Int32**| Filter by log.rotation.size | 
 **logRotationScheduleDays** | **optional.Int32**| Filter by log.rotation.schedule.days | 
 **logRotationScheduleWeekdays** | **optional.Int32**| Filter by log.rotation.schedule.weekdays | 
 **logRotationScheduleMinutes** | **optional.Int32**| Filter by log.rotation.schedule.minutes | 
 **logRotationScheduleMonths** | **optional.Int32**| Filter by log.rotation.schedule.months | 
 **logRotationScheduleHours** | **optional.Int32**| Filter by log.rotation.schedule.hours | 
 **logFormat** | **optional.String**| Filter by log.format | 
 **logPath** | **optional.String**| Filter by log_path | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **eventsFileShare** | **optional.Bool**| Filter by events.file_share | 
 **eventsAuthorizationPolicy** | **optional.Bool**| Filter by events.authorization_policy | 
 **eventsCifsLogonLogoff** | **optional.Bool**| Filter by events.cifs_logon_logoff | 
 **eventsCapStaging** | **optional.Bool**| Filter by events.cap_staging | 
 **eventsUserAccount** | **optional.Bool**| Filter by events.user_account | 
 **eventsSecurityGroup** | **optional.Bool**| Filter by events.security_group | 
 **eventsFileOperations** | **optional.Bool**| Filter by events.file_operations | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**AuditResponse**](audit_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditCreate

> AuditResponse AuditCreate(ctx, optional)



Creates an audit configuration. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM to which audit configuration is to be created. * `log_path` - Path in the owning SVM namespace that is used to store audit logs. ### Default property values If not specified in POST, the following default property values are assigned: * `enabled` - _true_ * `events.authorization_policy` - _false_ * `events.cap_staging` - _false_ * `events.file_share` - _false_ * `events.security_group` - _false_ * `events.user_account` - _false_ * `events.cifs_logon_logoff` - _true_ * `events.file_operations` - _true_ * `log.format` - _evtx_ * `log.retention.count` - _0_ * `log.retention.duration` - _PT0S_ * `log.rotation.size` - _100MB_ * `log.rotation.now` - _false_ ### Related ONTAP commands * `vserver audit create` * `vserver audit enable` ### Learn more * [`DOC /protocols/audit`](#docs-NAS-protocols_audit) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Audit**](Audit.md)| Info specification | 

### Return type

[**AuditResponse**](audit_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditDelete

> AuditDelete(ctx, svmUuid)



Deletes an audit configuration. ### Related ONTAP commands * `vserver audit disable` * `vserver audit delete` ### Learn more * [`DOC /protocols/audit`](#docs-NAS-protocols_audit) 

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


## AuditGet

> Audit AuditGet(ctx, svmUuid, optional)



Retrieves an audit configuration for an SVM. ### Related ONTAP commands * `vserver audit show` ### Learn more * [`DOC /protocols/audit`](#docs-NAS-protocols_audit) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***AuditGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Audit**](audit.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditModify

> AuditModify(ctx, svmUuid, optional)



Updates an audit configuration for an SVM. ### Related ONTAP commands * `vserver audit modify` ### Learn more * [`DOC /protocols/audit`](#docs-NAS-protocols_audit) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***AuditModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuditModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of Audit**](Audit.md)| Info specification | 

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


## CifsHomedirSearchPathGet

> CifsSearchPath CifsHomedirSearchPathGet(ctx, svmUuid, index, optional)



Retrieves a CIFS home directory search path of an SVM. ### Related ONTAP commands * `cifs server home-directory search-path show` ### Learn more * [`DOC /protocols/cifs/home-directory/search-paths`](#docs-NAS-protocols_cifs_home-directory_search-paths) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**index** | **int32**| Home directory search path index | 
 **optional** | ***CifsHomedirSearchPathGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsHomedirSearchPathGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CifsSearchPath**](cifs_search_path.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsSearchPathCollectionGet

> CifsSearchPathResponse CifsSearchPathCollectionGet(ctx, optional)



Retrieves CIFS home directory search paths. ### Related ONTAP commands * `cifs server home-directory search-path show` ### Learn more * [`DOC /protocols/cifs/home-directory/search-paths`](#docs-NAS-protocols_cifs_home-directory_search-paths) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsSearchPathCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSearchPathCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **index** | **optional.Int32**| Filter by index | 
 **path** | **optional.String**| Filter by path | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CifsSearchPathResponse**](cifs_search_path_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsSearchPathCreate

> CifsSearchPathResponse CifsSearchPathCreate(ctx, optional)



Creates a home directory search path. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the home directory search path. * `path` - Path in the owning SVM namespace that is used to search for home directories. ### Related ONTAP commands * `cifs server home-directory search-path add` ### Learn more * [`DOC /protocols/cifs/home-directory/search-paths`](#docs-NAS-protocols_cifs_home-directory_search-paths) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsSearchPathCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSearchPathCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of CifsSearchPath**](CifsSearchPath.md)| Info specification | 

### Return type

[**CifsSearchPathResponse**](cifs_search_path_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsSearchPathDelete

> CifsSearchPathDelete(ctx, svmUuid, index)



Deletes a CIFS home directory search path. ### Related ONTAP commands * `cifs server home-directory search-path remove` ### Learn more * [`DOC /protocols/cifs/home-directory/search-paths`](#docs-NAS-protocols_cifs_home-directory_search-paths) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**index** | **int32**| Home directory search path index | 

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


## CifsSearchPathModify

> CifsSearchPathModify(ctx, svmUuid, index, optional)



Reorders a CIFS home directory search path. ### Related ONTAP commands * `cifs server home-directory search-path reorder` ### Learn more * [`DOC /protocols/cifs/home-directory/search-paths`](#docs-NAS-protocols_cifs_home-directory_search-paths) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**index** | **int32**| Home directory search path index | 
 **optional** | ***CifsSearchPathModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSearchPathModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newIndex** | **optional.Int32**| New position for the home directory search path | 

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


## CifsServiceCollectionGet

> CifsServiceResponse CifsServiceCollectionGet(ctx, optional)



Retrieves CIFS servers. ### Related ONTAP commands * `vserver cifs server show` * `vserver cifs server options show` * `vserver cifs server security show` ### Learn more * [`DOC /protocols/cifs/services`](#docs-NAS-protocols_cifs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsServiceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsServiceCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **comment** | **optional.String**| Filter by comment | 
 **defaultUnixUser** | **optional.String**| Filter by default_unix_user | 
 **securityRestrictAnonymous** | **optional.String**| Filter by security.restrict_anonymous | 
 **securitySmbEncryption** | **optional.Bool**| Filter by security.smb_encryption | 
 **securityKdcEncryption** | **optional.Bool**| Filter by security.kdc_encryption | 
 **securitySmbSigning** | **optional.Bool**| Filter by security.smb_signing | 
 **adDomainUser** | **optional.String**| Filter by ad_domain.user | 
 **adDomainFqdn** | **optional.String**| Filter by ad_domain.fqdn | 
 **adDomainOrganizationalUnit** | **optional.String**| Filter by ad_domain.organizational_unit | 
 **netbiosWinsServers** | **optional.String**| Filter by netbios.wins_servers | 
 **netbiosAliases** | **optional.String**| Filter by netbios.aliases | 
 **netbiosEnabled** | **optional.Bool**| Filter by netbios.enabled | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CifsServiceResponse**](cifs_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsServiceCreate

> JobLinkResponse CifsServiceCreate(ctx, optional)



Creates a CIFS server. Each SVM can have one CIFS server.</br> ### Important notes - The CIFS server name might or might not be the same as the SVM name. - The CIFS server name can contain up to 15 characters. - The CIFS server name does not support the following characters: @ # * ( ) = + [ ] | ; : \" , < >  / ? ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the CIFS server. * `name` -  Name of the CIFS server. * `ad_domain.fqdn` - Fully qualified domain name of the Windows Active Directory to which this CIFS server belongs. * `ad_domain.user` - User account with the access to add the CIFS server to the Active Directory. * `ad_domain.password` - Account password used to add this CIFS server to the Active Directory. ### Recommended optional properties * `comment` - Add a text comment of up to 48 characters about the CIFS server. * `netbios.aliases` - Add a comma-delimited list of one or more NetBIOS aliases for the CIFS server. * `netbios.wins_servers` - Add a list of Windows Internet Name Server (WINS) addresses that manage and map the NetBIOS name of the CIFS server to their network IP addresses. The IP addresses must be IPv4 addresses. ### Default property values If not specified in POST, the following default property values are assigned: * `ad_domain.organizational_unit` - _CN=Computers_ * `enabled` - _true_ * `restrict_anonymous` - _no_enumeration_ * `smb_signing` - _false_ * `smb_encryption` - _false_ * `kdc_encryption` - _false_ * `default_unix_user` - _pcuser_ * `netbios_enabled` - _false_ However, if either \"netbios.wins-server\" or \"netbios.aliases\" is set during POST and if `netbios_enabled` is not specified then `netbios_enabled` is set to true. ### Related ONTAP commands * `vserver cifs server create` * `vserver cifs server options modify` * `vserver cifs security modify` * `vserver cifs server add-netbios-aliases` ### Learn more * [`DOC /protocols/cifs/services`](#docs-NAS-protocols_cifs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsServiceCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of CifsService**](CifsService.md)| Info specification | 

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


## CifsServiceDelete

> JobLinkResponse CifsServiceDelete(ctx, svmUuid, optional)



Deletes a CIFS server and related CIFS configurations. ### Related ONTAP commands * `vserver cifs server delete` * `vserver cifs remove-netbios-aliases` ### Learn more * [`DOC /protocols/cifs/services`](#docs-NAS-protocols_cifs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***CifsServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsServiceDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of CifsServiceDelete**](CifsServiceDelete.md)| Info specification | 

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


## CifsServiceGet

> CifsService CifsServiceGet(ctx, svmUuid, optional)



Retrieves a CIFS server. ### Related ONTAP commands * `vserver cifs server show` * `vserver cifs server options show` * `vserver cifs server security show` ### Learn more * [`DOC /protocols/cifs/services`](#docs-NAS-protocols_cifs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**|  | 
 **optional** | ***CifsServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsServiceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CifsService**](cifs_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsServiceModify

> JobLinkResponse CifsServiceModify(ctx, svmUuid, optional)



Updates both the mandatory and optional parameters of the CIFS configuration. Ensure the CIFS server is administratively disabled when renaming the CIFS server or modifying the <i>ad_domain</i> properties. ### Related ONTAP commands * `vserver cifs server modify` * `vserver cifs server options modify` * `vserver cifs security modify` * `vserver cifs server add-netbios-aliases` * `vserver cifs server remove-netbios-aliases` ### Learn more * [`DOC /protocols/cifs/services`](#docs-NAS-protocols_cifs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***CifsServiceModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsServiceModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of CifsService**](CifsService.md)| Info specification | 

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


## CifsShareAclCollectionGet

> CifsShareAclResponse CifsShareAclCollectionGet(ctx, svmUuid, share, optional)



Retrieves the share-level ACL on a CIFS share. ### Related ONTAP commands * `vserver cifs share access-control show` ### Learn more * [`DOC /protocols/cifs/shares/{svm.uuid}/{share}/acls`](#docs-NAS-protocols_cifs_shares_{svm.uuid}_{share}_acls) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**share** | **string**| CIFS Share Name | 
 **optional** | ***CifsShareAclCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareAclCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permission** | **optional.String**| Filter by permission | 
 **type_** | **optional.String**| Filter by type | 
 **userOrGroup** | **optional.String**| Filter by user_or_group | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CifsShareAclResponse**](cifs_share_acl_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsShareAclCreate

> CifsShareAclCreate(ctx, svmUuid, share, optional)



Creates a share-level ACL on a CIFS share. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the share acl. * `share` - Existing CIFS share in which to create the share acl. * `user_or_group` - Existing user or group name for which the acl is added on the CIFS share. * `permission` - Access rights that a user or group has on the defined CIFS share. ### Default property values * `type` - _windows_ ### Related ONTAP commands * `vserver cifs share access-control create` ### Learn more * [`DOC /protocols/cifs/shares/{svm.uuid}/{share}/acls`](#docs-NAS-protocols_cifs_shares_{svm.uuid}_{share}_acls) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**share** | **string**| CIFS Share Name | 
 **optional** | ***CifsShareAclCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareAclCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of CifsShareAcl**](CifsShareAcl.md)| Info specification | 

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


## CifsShareAclDelete

> CifsShareAclDelete(ctx, svmUuid, share, userOrGroup, type_)



Deletes a share-level ACL on a CIFS share. ### Related ONTAP commands * `vserver cifs share access-control delete` ### Learn more * [`DOC /protocols/cifs/shares/{svm.uuid}/{share}/acls`](#docs-NAS-protocols_cifs_shares_{svm.uuid}_{share}_acls) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**share** | **string**| Share name | 
**userOrGroup** | **string**| User or group name | 
**type_** | **string**| User or group type | 

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


## CifsShareAclGet

> CifsShareAcl CifsShareAclGet(ctx, svmUuid, share, userOrGroup, type_, optional)



Retrieves the share-level ACL on CIFS share for a specified user or group. ### Related ONTAP commands * `vserver cifs share access-control show` ### Learn more * [`DOC /protocols/cifs/shares/{svm.uuid}/{share}/acls`](#docs-NAS-protocols_cifs_shares_{svm.uuid}_{share}_acls) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**share** | **string**| Share name | 
**userOrGroup** | **string**| User or group name | 
**type_** | **string**| User or group type | 
 **optional** | ***CifsShareAclGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareAclGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CifsShareAcl**](cifs_share_acl.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsShareAclModify

> CifsShareAclModify(ctx, svmUuid, share, userOrGroup, type_, optional)



Updates a share-level ACL on a CIFS share. ### Related ONTAP commands * `vserver cifs share access-control modify` ### Learn more * [`DOC /protocols/cifs/shares/{svm.uuid}/{share}/acls`](#docs-NAS-protocols_cifs_shares_{svm.uuid}_{share}_acls) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**share** | **string**| Share name | 
**userOrGroup** | **string**| User or group name | 
**type_** | **string**| User or group type | 
 **optional** | ***CifsShareAclModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareAclModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **info** | [**optional.Interface of CifsShareAcl**](CifsShareAcl.md)| Info specification | 

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


## CifsShareCollectionGet

> CifsShareResponse CifsShareCollectionGet(ctx, optional)



Retrieves CIFS shares. ### Related ONTAP commands * `vserver cifs share show` * `vserver cifs share properties show` ### Learn more * [`DOC /protocols/cifs/shares`](#docs-NAS-protocols_cifs_shares) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsShareCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **unixSymlink** | **optional.String**| Filter by unix_symlink | 
 **comment** | **optional.String**| Filter by comment | 
 **volumeName** | **optional.String**| Filter by volume.name | 
 **volumeUuid** | **optional.String**| Filter by volume.uuid | 
 **aclsPermission** | **optional.String**| Filter by acls.permission | 
 **aclsType** | **optional.String**| Filter by acls.type | 
 **aclsUserOrGroup** | **optional.String**| Filter by acls.user_or_group | 
 **oplocks** | **optional.Bool**| Filter by oplocks | 
 **path** | **optional.String**| Filter by path | 
 **encryption** | **optional.Bool**| Filter by encryption | 
 **homeDirectory** | **optional.Bool**| Filter by home_directory | 
 **accessBasedEnumeration** | **optional.Bool**| Filter by access_based_enumeration | 
 **changeNotify** | **optional.Bool**| Filter by change_notify | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CifsShareResponse**](cifs_share_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsShareCreate

> CifsShareCreate(ctx, optional)



Creates a CIFS share. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the CIFS share. * `name` - Name of the CIFS share. * `path` - Path in the owning SVM namespace that is shared through this share. ### Recommended optional properties * `comment` - Optionally choose to add a text comment of up to 256 characters about the CIFS share. * `acls` - Optionally choose to add share permissions that users and groups have on the CIFS share. ### Default property values If not specified in POST, the following default property values are assigned: * `home_directory` - _false_ * `oplocks` - _true_ * `access_based_enumeration` - _false_ * `change_notify` - _true_ * `encryption` - _false_ * `unix_symlink` - _local_ ### Related ONTAP commands * `vserver cifs share create` * `vserver cifs share properties add` * `vserver cifs share access-control create` ### Learn more * [`DOC /protocols/cifs/shares`](#docs-NAS-protocols_cifs_shares) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsShareCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of CifsShare**](CifsShare.md)| Info specification | 

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


## CifsShareDelete

> CifsShareDelete(ctx, svmUuid, name)



Deletes a CIFS share. ### Related ONTAP commands * `vserver cifs share delete` ### Learn more * [`DOC /protocols/cifs/shares`](#docs-NAS-protocols_cifs_shares) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**| Share Name | 

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


## CifsShareGet

> CifsShare CifsShareGet(ctx, svmUuid, name, optional)



Retrieves a CIFS share. ### Related ONTAP commands * `vserver cifs share show` * `vserver cifs share properties show` ### Learn more * [`DOC /protocols/cifs/shares`](#docs-NAS-protocols_cifs_shares) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**| Share Name | 
 **optional** | ***CifsShareGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CifsShare**](cifs_share.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsShareModify

> CifsShareModify(ctx, svmUuid, name, optional)



Updates a CIFS share. ### Related ONTAP commands * `vserver cifs share modify` * `vserver cifs share properties add` * `vserver cifs share properties remove` ### Learn more * [`DOC /protocols/cifs/shares`](#docs-NAS-protocols_cifs_shares) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**| Share Name | 
 **optional** | ***CifsShareModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsShareModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of CifsShare**](CifsShare.md)| Info specification | 

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


## CifsSymlinkMappingCollectionGet

> CifsSymlinkMappingResponse CifsSymlinkMappingCollectionGet(ctx, optional)



Retrieves UNIX symbolic link mappings for CIFS clients. ### Related ONTAP commands * `vserver cifs symlink show` ### Learn more * [`DOC /protocols/cifs/unix-symlink-mapping`](#docs-NAS-protocols_cifs_unix-symlink-mapping) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsSymlinkMappingCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSymlinkMappingCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetLocality** | **optional.String**| Filter by target.locality | 
 **targetHomeDirectory** | **optional.Bool**| Filter by target.home_directory | 
 **targetServer** | **optional.String**| Filter by target.server | 
 **targetShare** | **optional.String**| Filter by target.share | 
 **targetPath** | **optional.String**| Filter by target.path | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **unixPath** | **optional.String**| Filter by unix_path | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CifsSymlinkMappingResponse**](cifs_symlink_mapping_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsSymlinkMappingCreate

> CifsSymlinkMappingResponse CifsSymlinkMappingCreate(ctx, optional)



Creates a UNIX symbolic link mapping for a CIFS client. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the CIFS unix-symlink-mapping. * `unix_path` - UNIX path to which the CIFS symlink mapping to be created. * `target.share` - CIFS share name on the destination CIFS server to which the UNIX symbolic link is pointing. * `target.path` - CIFS path on the destination to which the symbolic link maps. ### Default property values * `target.server` - _Local_NetBIOS_Server_Name_ * `locality` - _local_ * `home_directory` - _false_ ### Related ONTAP commands * `vserver cifs symlink create` ### Learn more * [`DOC /protocols/cifs/unix-symlink-mapping`](#docs-NAS-protocols_cifs_unix-symlink-mapping) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CifsSymlinkMappingCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSymlinkMappingCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of CifsSymlinkMapping**](CifsSymlinkMapping.md)| Info specification | 

### Return type

[**CifsSymlinkMappingResponse**](cifs_symlink_mapping_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsSymlinkMappingDelete

> CifsSymlinkMappingDelete(ctx, svmUuid, unixPath)



Deletes the UNIX symbolic link mapping for CIFS clients. ### Related ONTAP commands * `vserver cifs symlink delete` ### Learn more * [`DOC /protocols/cifs/unix-symlink-mapping`](#docs-NAS-protocols_cifs_unix-symlink-mapping) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**unixPath** | **string**| UNIX symbolic link path | 

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


## CifsSymlinkMappingGet

> CifsSymlinkMapping CifsSymlinkMappingGet(ctx, svmUuid, unixPath, optional)



Retrieves a UNIX symbolic link mapping for CIFS clients. ### Related ONTAP commands * `vserver cifs symlink show` ### Learn more * [`DOC /protocols/cifs/unix-symlink-mapping`](#docs-NAS-protocols_cifs_unix-symlink-mapping) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**unixPath** | **string**| UNIX symbolic link path | 
 **optional** | ***CifsSymlinkMappingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSymlinkMappingGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CifsSymlinkMapping**](cifs_symlink_mapping.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CifsSymlinkMappingModify

> CifsSymlinkMappingModify(ctx, svmUuid, unixPath, optional)



Updates the UNIX symbolic link mapping for CIFS clients. ### Related ONTAP commands * `vserver cifs symlink modify` ### Learn more * [`DOC /protocols/cifs/unix-symlink-mapping`](#docs-NAS-protocols_cifs_unix-symlink-mapping) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**unixPath** | **string**| UNIX symbolic link path | 
 **optional** | ***CifsSymlinkMappingModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CifsSymlinkMappingModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of CifsSymlinkMapping**](CifsSymlinkMapping.md)| Info specification | 

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


## ExportPolicyCollectionGet

> ExportPolicyResponse ExportPolicyCollectionGet(ctx, optional)



Retrieves export policies. ### Related ONTAP commands * `vserver export-policy show` * `vserver export-policy rule show` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportPolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportPolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rulesProtocols** | **optional.String**| Filter by rules.protocols | 
 **rulesAnonymousUser** | **optional.String**| Filter by rules.anonymous_user | 
 **rulesRwRule** | **optional.String**| Filter by rules.rw_rule | 
 **rulesIndex** | **optional.Int32**| Filter by rules.index | 
 **rulesSuperuser** | **optional.String**| Filter by rules.superuser | 
 **rulesRoRule** | **optional.String**| Filter by rules.ro_rule | 
 **rulesClientsMatch** | **optional.String**| Filter by rules.clients.match | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **id** | **optional.Int32**| Filter by id | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ExportPolicyResponse**](export_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPolicyCreate

> ExportPolicyResponse ExportPolicyCreate(ctx, optional)



Creates an export policy. An SVM can have any number of export policies to define rules for which clients can access data exported by the SVM. A policy with no rules prohibits access. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create an export policy. * `name`  - Name of the export policy. ### Recommended optional properties * `rules`  - Rule(s) of an export policy. Used to create the export rule and populate the export policy with export rules in a single request. ### Related ONTAP commands * `vserver export-policy create` * `vserver export-policy rule create` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportPolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ExportPolicy**](ExportPolicy.md)| Info specification | 

### Return type

[**ExportPolicyResponse**](export_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPolicyDelete

> ExportPolicyDelete(ctx, id)



Deletes an export policy. ### Related ONTAP commands * `vserver export-policy delete` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Export Policy ID | 

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


## ExportPolicyGet

> ExportPolicy ExportPolicyGet(ctx, id, optional)



Retrieves an export policy. ### Related ONTAP commands * `vserver export-policy show` * `vserver export-policy rule show` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Export Policy ID | 
 **optional** | ***ExportPolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportPolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ExportPolicy**](export_policy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPolicyModify

> ExportPolicyModify(ctx, id, optional)



Updates the properties of an export policy to change an export policy name or replace all export policy rules. ### Related ONTAP commands * `vserver export-policy rename` * `vserver export-policy rule delete` * `vserver export-policy rule create` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| Export Policy ID | 
 **optional** | ***ExportPolicyModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportPolicyModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of ExportPolicy**](ExportPolicy.md)| Info specification | 

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


## ExportRuleClientsCreate

> ExportClientResponse ExportRuleClientsCreate(ctx, policyId, index, optional)



Creates an export policy rule client ### Required properties * `policy.id` - Existing export policy that contains export policy rules for the client being added. * `index`  - Existing export policy rule for which to create an export client. * `match`  - Base name for the export policy client. ### Related ONTAP commands * `vserver export-policy rule add-clientmatches` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
**index** | **int32**| Export Rule Index | 
 **optional** | ***ExportRuleClientsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportRuleClientsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of ExportClient**](ExportClient.md)| Info specification | 

### Return type

[**ExportClientResponse**](export_client_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRuleClientsDelete

> ExportRuleClientsDelete(ctx, policyId, index, match)



Deletes an export policy client ### Related ONTAP commands * `vserver export-policy rule remove-clientmatches` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
**index** | **int32**| Export Rule Index | 
**match** | **string**| Export Client Match | 

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


## ExportRuleClientsGet

> ExportClientResponse ExportRuleClientsGet(ctx, policyId, index)



Retrieves export policy rule clients. ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
**index** | **int32**| Export Rule Index | 

### Return type

[**ExportClientResponse**](export_client_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRuleCollectionGet

> ExportRuleResponse ExportRuleCollectionGet(ctx, policyId, optional)



Retrieves export policy rules. ### Related ONTAP commands * `vserver export-policy rule show` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
 **optional** | ***ExportRuleCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportRuleCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **protocols** | **optional.String**| Filter by protocols | 
 **anonymousUser** | **optional.String**| Filter by anonymous_user | 
 **rwRule** | **optional.String**| Filter by rw_rule | 
 **index** | **optional.Int32**| Filter by index | 
 **superuser** | **optional.String**| Filter by superuser | 
 **roRule** | **optional.String**| Filter by ro_rule | 
 **clientsMatch** | **optional.String**| Filter by clients.match | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ExportRuleResponse**](export_rule_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRuleCreate

> ExportRuleResponse ExportRuleCreate(ctx, policyId, optional)



Creates an export policy rule. ### Required properties * `policy.id`  - Existing export policy for which to create an export rule. * `clients.match`  - List of clients (hostnames, ipaddresses, netgroups, domains) to which the export rule applies. * `ro_rule`  - Used to specify the security type for read-only access to volumes that use the export rule. * `rw_rule`  - Used to specify the security type for read-write access to volumes that use the export rule. ### Default property values If not specified in POST, the following default property values are assigned: * `protocols` - _any_ * `anonymous_user` - _none_ * `superuser` - _any_ ### Related ONTAP commands * `vserver export-policy rule create` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
 **optional** | ***ExportRuleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportRuleCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of ExportRule**](ExportRule.md)| Info specification | 

### Return type

[**ExportRuleResponse**](export_rule_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRuleDelete

> ExportRuleDelete(ctx, policyId, index)



Deletes an export policy rule. ### Related ONTAP commands * `vserver export-policy rule delete` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
**index** | **int32**| Export Rule Index | 

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


## ExportRuleGet

> ExportRule ExportRuleGet(ctx, policyId, index, optional)



Retrieves an export policy rule ### Related ONTAP commands * `vserver export-policy rule show` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
**index** | **int32**| Export Rule Index | 
 **optional** | ***ExportRuleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportRuleGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ExportRule**](export_rule.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRuleModify

> ExportRuleModify(ctx, policyId, index, optional)



Updates the properties of an export policy rule to change an export policy rule's index or fields. ### Related ONTAP commands * `vserver export-policy rule modify` * `vserver export-policy rule setindex` ### Learn more * [`DOC /protocols/nfs/export-policies`](#docs-NAS-protocols_nfs_export-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32**| Export Policy ID | 
**index** | **int32**| Export Rule Index | 
 **optional** | ***ExportRuleModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportRuleModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newIndex** | **optional.Int32**| New Export Rule Index | 
 **info** | [**optional.Interface of ExportRule**](ExportRule.md)| Info specification | 

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


## FpolicyCollectionGet

> FpolicyResponse FpolicyCollectionGet(ctx, optional)



Retrieves an FPolicy configuration. ### Related ONTAP commands * `fpolicy show` * `fpolicy policy show` * `fpolicy policy scope show` * `fpolicy policy event show` * `fpolicy policy external-engine show` ### Learn more * [`DOC /protocols/fpolicy`](#docs-NAS-protocols_fpolicy) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FpolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enginesSecondaryServers** | **optional.String**| Filter by engines.secondary_servers | 
 **enginesName** | **optional.String**| Filter by engines.name | 
 **enginesType** | **optional.String**| Filter by engines.type | 
 **enginesPrimaryServers** | **optional.String**| Filter by engines.primary_servers | 
 **enginesPort** | **optional.Int32**| Filter by engines.port | 
 **policiesEnabled** | **optional.Bool**| Filter by policies.enabled | 
 **policiesScopeExcludeExportPolicies** | **optional.String**| Filter by policies.scope.exclude_export_policies | 
 **policiesScopeExcludeShares** | **optional.String**| Filter by policies.scope.exclude_shares | 
 **policiesScopeIncludeExtension** | **optional.String**| Filter by policies.scope.include_extension | 
 **policiesScopeExcludeExtension** | **optional.String**| Filter by policies.scope.exclude_extension | 
 **policiesScopeIncludeShares** | **optional.String**| Filter by policies.scope.include_shares | 
 **policiesScopeIncludeVolumes** | **optional.String**| Filter by policies.scope.include_volumes | 
 **policiesScopeIncludeExportPolicies** | **optional.String**| Filter by policies.scope.include_export_policies | 
 **policiesScopeExcludeVolumes** | **optional.String**| Filter by policies.scope.exclude_volumes | 
 **policiesPriority** | **optional.Int32**| Filter by policies.priority | 
 **policiesEngineName** | **optional.String**| Filter by policies.engine.name | 
 **policiesEventsName** | **optional.String**| Filter by policies.events.name | 
 **policiesName** | **optional.String**| Filter by policies.name | 
 **policiesMandatory** | **optional.Bool**| Filter by policies.mandatory | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **eventsVolumeMonitoring** | **optional.Bool**| Filter by events.volume_monitoring | 
 **eventsProtocol** | **optional.String**| Filter by events.protocol | 
 **eventsFiltersMonitorAds** | **optional.Bool**| Filter by events.filters.monitor_ads | 
 **eventsFiltersWriteWithSizeChange** | **optional.Bool**| Filter by events.filters.write_with_size_change | 
 **eventsFiltersFirstRead** | **optional.Bool**| Filter by events.filters.first_read | 
 **eventsFiltersSetattrWithModifyTimeChange** | **optional.Bool**| Filter by events.filters.setattr_with_modify_time_change | 
 **eventsFiltersOfflineBit** | **optional.Bool**| Filter by events.filters.offline_bit | 
 **eventsFiltersOpenWithWriteIntent** | **optional.Bool**| Filter by events.filters.open_with_write_intent | 
 **eventsFiltersFirstWrite** | **optional.Bool**| Filter by events.filters.first_write | 
 **eventsFiltersSetattrWithGroupChange** | **optional.Bool**| Filter by events.filters.setattr_with_group_change | 
 **eventsFiltersSetattrWithAccessTimeChange** | **optional.Bool**| Filter by events.filters.setattr_with_access_time_change | 
 **eventsFiltersSetattrWithModeChange** | **optional.Bool**| Filter by events.filters.setattr_with_mode_change | 
 **eventsFiltersExcludeDirectory** | **optional.Bool**| Filter by events.filters.exclude_directory | 
 **eventsFiltersSetattrWithOwnerChange** | **optional.Bool**| Filter by events.filters.setattr_with_owner_change | 
 **eventsFiltersSetattrWithDaclChange** | **optional.Bool**| Filter by events.filters.setattr_with_dacl_change | 
 **eventsFiltersCloseWithModification** | **optional.Bool**| Filter by events.filters.close_with_modification | 
 **eventsFiltersOpenWithDeleteIntent** | **optional.Bool**| Filter by events.filters.open_with_delete_intent | 
 **eventsFiltersSetattrWithSaclChange** | **optional.Bool**| Filter by events.filters.setattr_with_sacl_change | 
 **eventsFiltersSetattrWithAllocationSizeChange** | **optional.Bool**| Filter by events.filters.setattr_with_allocation_size_change | 
 **eventsFiltersCloseWithoutModification** | **optional.Bool**| Filter by events.filters.close_without_modification | 
 **eventsFiltersSetattrWithSizeChange** | **optional.Bool**| Filter by events.filters.setattr_with_size_change | 
 **eventsFiltersCloseWithRead** | **optional.Bool**| Filter by events.filters.close_with_read | 
 **eventsFiltersSetattrWithCreationTimeChange** | **optional.Bool**| Filter by events.filters.setattr_with_creation_time_change | 
 **eventsName** | **optional.String**| Filter by events.name | 
 **eventsFileOperationsLookup** | **optional.Bool**| Filter by events.file_operations.lookup | 
 **eventsFileOperationsDeleteDir** | **optional.Bool**| Filter by events.file_operations.delete_dir | 
 **eventsFileOperationsCreate** | **optional.Bool**| Filter by events.file_operations.create | 
 **eventsFileOperationsRead** | **optional.Bool**| Filter by events.file_operations.read | 
 **eventsFileOperationsRenameDir** | **optional.Bool**| Filter by events.file_operations.rename_dir | 
 **eventsFileOperationsWrite** | **optional.Bool**| Filter by events.file_operations.write | 
 **eventsFileOperationsGetattr** | **optional.Bool**| Filter by events.file_operations.getattr | 
 **eventsFileOperationsSymlink** | **optional.Bool**| Filter by events.file_operations.symlink | 
 **eventsFileOperationsLink** | **optional.Bool**| Filter by events.file_operations.link | 
 **eventsFileOperationsRename** | **optional.Bool**| Filter by events.file_operations.rename | 
 **eventsFileOperationsSetattr** | **optional.Bool**| Filter by events.file_operations.setattr | 
 **eventsFileOperationsCreateDir** | **optional.Bool**| Filter by events.file_operations.create_dir | 
 **eventsFileOperationsOpen** | **optional.Bool**| Filter by events.file_operations.open | 
 **eventsFileOperationsDelete** | **optional.Bool**| Filter by events.file_operations.delete | 
 **eventsFileOperationsClose** | **optional.Bool**| Filter by events.file_operations.close | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FpolicyResponse**](fpolicy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyCreate

> FpolicyResponse FpolicyCreate(ctx, optional)



Creates an FPolicy configuration. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the FPolicy configuration. ### Recommended optional properties * `engines` -  External server to which the notifications will be sent. * `events` - File operations to monitor. * `policies` - Policy configuration which acts as a container for FPolicy event and FPolicy engine. * `scope` - Scope of the policy. Can be limited to exports, volumes, shares or file extensions. ### Default property values If not specified in POST, the following default property values are assigned: * `engines.type` - _synchronous_ * `policies.engine` - _native_ * `policies.mandatory` -  _true_ * `events.volume_monitoring` - _false_ * `events.file_operations.*` - _false_ * `events.filters.*` - _false_ ### Related ONTAP commands * `fpolicy policy event create` * `fpolicy policy external-engine create` * `fpolicy policy create` * `fpolicy policy scope create` * `fpolicy enable` ### Learn more * [`DOC /protocols/fpolicy`](#docs-NAS-protocols_fpolicy) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FpolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Fpolicy**](Fpolicy.md)| Info specification | 

### Return type

[**FpolicyResponse**](fpolicy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyDelete

> FpolicyDelete(ctx, svmUuid)



Deletes the FPolicy configuration for the specified SVM. Before deleting the FPolicy configuration, ensure that all policies belonging to the SVM are disabled. ### Related ONTAP commands * `fpolicy delete` * `fpolicy policy scope delete` * `fpolicy policy delete` * `fpolicy policy event delete` * `fpolicy policy external-engine delete` ### Learn more * [`DOC /protocols/fpolicy`](#docs-NAS-protocols_fpolicy) 

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


## FpolicyEngineCollectionGet

> FpolicyEngineResponse FpolicyEngineCollectionGet(ctx, svmUuid, optional)



Retrieves FPolicy engine configurations of all the engines for a specified SVM. ONTAP allows creation of cluster-level FPolicy engines that act as a template for all the SVMs belonging to the cluster. These cluster-level FPolicy engines are also retrieved for the specified SVM. ### Related ONTAP commands * `fpolicy policy external-engine show` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/engines`](#docs-NAS-protocols_fpolicy_{svm.uuid}_engines) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyEngineCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEngineCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secondaryServers** | **optional.String**| Filter by secondary_servers | 
 **name** | **optional.String**| Filter by name | 
 **type_** | **optional.String**| Filter by type | 
 **primaryServers** | **optional.String**| Filter by primary_servers | 
 **port** | **optional.Int32**| Filter by port | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FpolicyEngineResponse**](fpolicy_engine_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyEngineCreate

> FpolicyEngineResponse FpolicyEngineCreate(ctx, svmUuid, optional)



Creates an FPolicy engine configuration for a specified SVM. FPolicy engine creation is allowed only on data SVMs. ### Required properties * `svm.uuid` - Existing SVM in which to create the FPolicy engine. * `name` - Name of external engine. * `port` - Port number of the FPolicy server application. * `primary_servers` - List of primary FPolicy servers to which the node will send notifications. ### Recommended optional properties * `secondary_servers` - It is recommended to configure secondary FPolicy server to which the node will send notifications when the primary server is down. ### Default property values * `type` - _synchronous_ ### Related ONTAP commands * `fpolicy policy external-engine create` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/engines`](#docs-NAS-protocols_fpolicy_{svm.uuid}_engines) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyEngineCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEngineCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of FpolicyEngine**](FpolicyEngine.md)| Info specification | 

### Return type

[**FpolicyEngineResponse**](fpolicy_engine_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyEngineDelete

> FpolicyEngineDelete(ctx, svmUuid, name)



Deletes the FPolicy external engine configuration. Deletion of an FPolicy engine that is attached to one or more FPolicy policies is not allowed. ### Related ONTAP commands * `fpolicy policy external-engine modify` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/engines`](#docs-NAS-protocols_fpolicy_{svm.uuid}_engines) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 

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


## FpolicyEngineGet

> FpolicyEngine FpolicyEngineGet(ctx, svmUuid, name, optional)



Retrieves a particular FPolicy engine configuration of a specifed SVM. A cluster-level FPolicy engine configuration cannot be retrieved for a data SVM. ### Related ONTAP commands * `fpolicy policy external-engine show` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/engines`](#docs-NAS-protocols_fpolicy_{svm.uuid}_engines) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***FpolicyEngineGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEngineGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FpolicyEngine**](fpolicy_engine.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyEngineModify

> FpolicyEngineModify(ctx, svmUuid, name, optional)



Updates a specific FPolicy engine configuration of an SVM. Modification of an FPolicy engine that is attached to one or more enabled FPolicy policies is not allowed. ### Related ONTAP commands * `fpolicy policy external-engine modify` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/engines`](#docs-NAS-protocols_fpolicy_{svm.uuid}_engines) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***FpolicyEngineModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEngineModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of FpolicyEngine**](FpolicyEngine.md)| Info specification | 

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


## FpolicyEventCollectionGet

> FpolicyEventResponse FpolicyEventCollectionGet(ctx, svmUuid, optional)



Retrieves FPolicy event configurations for all events for a specified SVM. ONTAP allows the creation of cluster-level FPolicy events that act as a template for all the data SVMs belonging to the cluster. These cluster-level FPolicy events are also retrieved for the specified SVM. ### Related ONTAP commands * `fpolicy policy event show` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/events`](#docs-NAS-protocols_fpolicy_{svm.uuid}_events) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyEventCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEventCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeMonitoring** | **optional.Bool**| Filter by volume_monitoring | 
 **protocol** | **optional.String**| Filter by protocol | 
 **filtersMonitorAds** | **optional.Bool**| Filter by filters.monitor_ads | 
 **filtersWriteWithSizeChange** | **optional.Bool**| Filter by filters.write_with_size_change | 
 **filtersFirstRead** | **optional.Bool**| Filter by filters.first_read | 
 **filtersSetattrWithModifyTimeChange** | **optional.Bool**| Filter by filters.setattr_with_modify_time_change | 
 **filtersOfflineBit** | **optional.Bool**| Filter by filters.offline_bit | 
 **filtersOpenWithWriteIntent** | **optional.Bool**| Filter by filters.open_with_write_intent | 
 **filtersFirstWrite** | **optional.Bool**| Filter by filters.first_write | 
 **filtersSetattrWithGroupChange** | **optional.Bool**| Filter by filters.setattr_with_group_change | 
 **filtersSetattrWithAccessTimeChange** | **optional.Bool**| Filter by filters.setattr_with_access_time_change | 
 **filtersSetattrWithModeChange** | **optional.Bool**| Filter by filters.setattr_with_mode_change | 
 **filtersExcludeDirectory** | **optional.Bool**| Filter by filters.exclude_directory | 
 **filtersSetattrWithOwnerChange** | **optional.Bool**| Filter by filters.setattr_with_owner_change | 
 **filtersSetattrWithDaclChange** | **optional.Bool**| Filter by filters.setattr_with_dacl_change | 
 **filtersCloseWithModification** | **optional.Bool**| Filter by filters.close_with_modification | 
 **filtersOpenWithDeleteIntent** | **optional.Bool**| Filter by filters.open_with_delete_intent | 
 **filtersSetattrWithSaclChange** | **optional.Bool**| Filter by filters.setattr_with_sacl_change | 
 **filtersSetattrWithAllocationSizeChange** | **optional.Bool**| Filter by filters.setattr_with_allocation_size_change | 
 **filtersCloseWithoutModification** | **optional.Bool**| Filter by filters.close_without_modification | 
 **filtersSetattrWithSizeChange** | **optional.Bool**| Filter by filters.setattr_with_size_change | 
 **filtersCloseWithRead** | **optional.Bool**| Filter by filters.close_with_read | 
 **filtersSetattrWithCreationTimeChange** | **optional.Bool**| Filter by filters.setattr_with_creation_time_change | 
 **name** | **optional.String**| Filter by name | 
 **fileOperationsLookup** | **optional.Bool**| Filter by file_operations.lookup | 
 **fileOperationsDeleteDir** | **optional.Bool**| Filter by file_operations.delete_dir | 
 **fileOperationsCreate** | **optional.Bool**| Filter by file_operations.create | 
 **fileOperationsRead** | **optional.Bool**| Filter by file_operations.read | 
 **fileOperationsRenameDir** | **optional.Bool**| Filter by file_operations.rename_dir | 
 **fileOperationsWrite** | **optional.Bool**| Filter by file_operations.write | 
 **fileOperationsGetattr** | **optional.Bool**| Filter by file_operations.getattr | 
 **fileOperationsSymlink** | **optional.Bool**| Filter by file_operations.symlink | 
 **fileOperationsLink** | **optional.Bool**| Filter by file_operations.link | 
 **fileOperationsRename** | **optional.Bool**| Filter by file_operations.rename | 
 **fileOperationsSetattr** | **optional.Bool**| Filter by file_operations.setattr | 
 **fileOperationsCreateDir** | **optional.Bool**| Filter by file_operations.create_dir | 
 **fileOperationsOpen** | **optional.Bool**| Filter by file_operations.open | 
 **fileOperationsDelete** | **optional.Bool**| Filter by file_operations.delete | 
 **fileOperationsClose** | **optional.Bool**| Filter by file_operations.close | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FpolicyEventResponse**](fpolicy_event_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyEventCreate

> FpolicyEventResponse FpolicyEventCreate(ctx, svmUuid, optional)



Creates an FPolicy event configuration for a specified SVM. FPolicy event creation is allowed only on data SVMs. When a protocol is specified, you must specify a file operation or a file operation and filters. ### Required properties * `svm.uuid` - Existing SVM in which to create the FPolicy event. * `name` - Name of the FPolicy event. ### Recommended optional properties * `file-operations` - List of file operations to monitor. * `protocol` - Protocol for which the file operations should be monitored. * `filters` - List of filters for the specified file operations. ### Default property values If not specified in POST, the following default property values are assigned: * `file_operations.*` - _false_ * `filters.*` - _false_ * `volume-monitoring` - _false_ ### Related ONTAP commands * `fpolicy policy event create` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/events`](#docs-NAS-protocols_fpolicy_{svm.uuid}_events) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyEventCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEventCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of FpolicyEvent**](FpolicyEvent.md)| Info specification | 

### Return type

[**FpolicyEventResponse**](fpolicy_event_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyEventDelete

> FpolicyEventDelete(ctx, svmUuid, name)



Deletes a specific FPolicy event configuration for an SVM. A cluster-level FPolicy event configuration cannot be modified for a data SVM through REST. An FPolicy event that is attached to an FPolicy policy cannot be deleted. ### Related ONTAP commands * `fpolicy policy event delete` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/events`](#docs-NAS-protocols_fpolicy_{svm.uuid}_events) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 

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


## FpolicyEventModify

> FpolicyEventModify(ctx, svmUuid, name, optional)



Updates a specific FPolicy event configuration for an SVM. A cluster-level FPolicy event configuration cannot be modified for a data SVM through REST. When the file operations and filters fields are modified, the previous values are retained and new values are added to the list of previous values. To remove a particular file operation or filter, set its value to false in the request. ### Related ONTAP commands * `fpolicy policy event modify` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/events`](#docs-NAS-protocols_fpolicy_{svm.uuid}_events) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***FpolicyEventModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEventModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of FpolicyEvent**](FpolicyEvent.md)| Info specification | 

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


## FpolicyEventsGet

> FpolicyEvent FpolicyEventsGet(ctx, svmUuid, name, optional)



Retrieves a specific FPolicy event configuration for an SVM. A cluster-level FPolicy event configuration cannot be retrieved for a data SVM through a REST API. ### Related ONTAP commands * `fpolicy policy event show` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/events`](#docs-NAS-protocols_fpolicy_{svm.uuid}_events) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***FpolicyEventsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyEventsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FpolicyEvent**](fpolicy_event.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyGet

> Fpolicy FpolicyGet(ctx, svmUuid, optional)



Retrieves an FPolicy configuration of an SVM. ### Related ONTAP commands * `fpolicy show` * `fpolicy policy show` * `fpolicy policy scope show` * `fpolicy policy event show` * `fpolicy policy external-engine show` ### Learn more * [`DOC /protocols/fpolicy`](#docs-NAS-protocols_fpolicy) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Fpolicy**](fpolicy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyPolicyCollectionGet

> FpolicyPolicyResponse FpolicyPolicyCollectionGet(ctx, svmUuid, optional)



Retrieves the FPolicy policy configuration of an SVM. ONTAP allows the creation of a cluster level FPolicy policy that acts as a template for all the data SVMs belonging to the cluster. This cluster level FPolicy policy is also retrieved for the specified SVM. ### Related ONTAP commands * `fpolicy policy show` * `fpolicy policy scope show` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/policies`](#docs-NAS-protocols_fpolicy_{svm.uuid}_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyPolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyPolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **optional.Bool**| Filter by enabled | 
 **scopeExcludeExportPolicies** | **optional.String**| Filter by scope.exclude_export_policies | 
 **scopeExcludeShares** | **optional.String**| Filter by scope.exclude_shares | 
 **scopeIncludeExtension** | **optional.String**| Filter by scope.include_extension | 
 **scopeExcludeExtension** | **optional.String**| Filter by scope.exclude_extension | 
 **scopeIncludeShares** | **optional.String**| Filter by scope.include_shares | 
 **scopeIncludeVolumes** | **optional.String**| Filter by scope.include_volumes | 
 **scopeIncludeExportPolicies** | **optional.String**| Filter by scope.include_export_policies | 
 **scopeExcludeVolumes** | **optional.String**| Filter by scope.exclude_volumes | 
 **priority** | **optional.Int32**| Filter by priority | 
 **engineName** | **optional.String**| Filter by engine.name | 
 **eventsName** | **optional.String**| Filter by events.name | 
 **name** | **optional.String**| Filter by name | 
 **mandatory** | **optional.Bool**| Filter by mandatory | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**FpolicyPolicyResponse**](fpolicy_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyPolicyCreate

> FpolicyPolicyResponse FpolicyPolicyCreate(ctx, svmUuid, optional)



Creates an FPolicy policy configuration for the specified SVM. To create an FPolicy policy, you must specify the policy scope and the FPolicy events to be monitored. </br>Important notes: * A single policy can monitor multiple events. * An FPolicy engine is an optional field whose default value is set to native. A native engine can be used to simply block the file access based on the file extensions specified in the policy scope. * To enable a policy, the policy priority  must be specified. If the priority is not specified, the policy is created but it is not enabled. * The \"mandatory\" field, if set to true, blocks the file access when the primary or secondary FPolicy servers are down. ### Required properties * `svm.uuid` - Existing SVM in which to create the FPolicy policy. * `events` - Name of the events to monitior. * `name` - Name of the FPolicy policy. * `scope` - Scope of the policy. Can be limited to exports, volumes, shares or file extensions. * `priority`- Priority of the policy (ranging from 1 to 10). ### Default property values * `mandatory` - _true_ * `engine` - _native_ ### Related ONTAP commands * `fpolicy policy scope create` * `fpolicy policy create` * `fpolicy enable` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/policies`](#docs-NAS-protocols_fpolicy_{svm.uuid}_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***FpolicyPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyPolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of FpolicyPolicy**](FpolicyPolicy.md)| Info specification | 

### Return type

[**FpolicyPolicyResponse**](fpolicy_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyPolicyDelete

> FpolicyPolicyDelete(ctx, svmUuid, name)



Deletes a particular FPolicy policy configuration for a specified SVM. To delete a policy, you must first disable the policy. ### Related ONTAP commands * `fpolicy policy scope delete` * `fpolicy policy delete` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/policies`](#docs-NAS-protocols_fpolicy_{svm.uuid}_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 

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


## FpolicyPolicyGet

> FpolicyPolicy FpolicyPolicyGet(ctx, svmUuid, name, optional)



Retrieves a particular FPolicy policy configuration for a specified SVM. Cluster-level FPolicy policy configuration details cannot be retrieved for a data SVM. ### Related ONTAP commands * `fpolicy policy show` * `fpolicy policy scope show` * `fpolicy show` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/policies`](#docs-NAS-protocols_fpolicy_{svm.uuid}_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***FpolicyPolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyPolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FpolicyPolicy**](fpolicy_policy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FpolicyPolicyModify

> FpolicyPolicyModify(ctx, svmUuid, name, optional)



Updates a particular FPolicy policy configuration for a specified SVM. PATCH can be used to enable or disable the policy. When enabling a policy, you must specify the policy priority. The policy priority of the policy is not required when disabling the policy. If the policy is enabled, the FPolicy policy engine cannot be modified. ### Related ONTAP commands * `fpolicy policy modify` * `fpolicy policy scope modify` * `fpolicy enable` * `fpolicy disable` ### Learn more * [`DOC /protocols/fpolicy/{svm.uuid}/policies`](#docs-NAS-protocols_fpolicy_{svm.uuid}_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***FpolicyPolicyModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FpolicyPolicyModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of FpolicyPolicy**](FpolicyPolicy.md)| Info specification | 

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


## KerberosInterfaceCollectionGet

> KerberosInterfaceResponse KerberosInterfaceCollectionGet(ctx, optional)



Retrieves Kerberos interfaces. ### Related ONTAP commands * `vserver nfs kerberos interface show` ### Learn more * [`DOC /protocols/nfs/kerberos/interfaces`](#docs-NAS-protocols_nfs_kerberos_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KerberosInterfaceCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosInterfaceCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptionTypes** | **optional.String**| Filter by encryption_types | 
 **interfaceName** | **optional.String**| Filter by interface.name | 
 **interfaceUuid** | **optional.String**| Filter by interface.uuid | 
 **interfaceIpAddress** | **optional.String**| Filter by interface.ip.address | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **spn** | **optional.String**| Filter by spn | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**KerberosInterfaceResponse**](kerberos_interface_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosInterfaceGet

> KerberosInterface KerberosInterfaceGet(ctx, uuid, optional)



Retrieves a Kerberos interface. ### Related ONTAP commands * `vserver nfs kerberos interface show` ### Learn more * [`DOC /protocols/nfs/kerberos/interfaces`](#docs-NAS-protocols_nfs_kerberos_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Network interface UUID | 
 **optional** | ***KerberosInterfaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosInterfaceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**KerberosInterface**](kerberos_interface.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosInterfaceModify

> KerberosInterfaceModify(ctx, uuid, optional)



Updates the properties of a Kerberos interface. ### Related ONTAP commands * `vserver nfs kerberos interface modify` * `vserver nfs kerberos interface enable` * `vserver nfs kerberos interface disable` ### Learn more * [`DOC /protocols/nfs/kerberos/interfaces`](#docs-NAS-protocols_nfs_kerberos_interfaces) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Network interface UUID | 
 **optional** | ***KerberosInterfaceModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosInterfaceModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of KerberosInterface**](KerberosInterface.md)| Info specification | 

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


## KerberosRealmCollectionGet

> KerberosRealmResponse KerberosRealmCollectionGet(ctx, optional)



Retrieves Kerberos realms. ### Related ONTAP commands * `vserver nfs kerberos realm show` ### Learn more * [`DOC /protocols/nfs/kerberos/realms`](#docs-NAS-protocols_nfs_kerberos_realms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KerberosRealmCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosRealmCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adServerName** | **optional.String**| Filter by ad_server.name | 
 **adServerAddress** | **optional.String**| Filter by ad_server.address | 
 **kdcVendor** | **optional.String**| Filter by kdc.vendor | 
 **kdcIp** | **optional.String**| Filter by kdc.ip | 
 **kdcPort** | **optional.Int32**| Filter by kdc.port | 
 **comment** | **optional.String**| Filter by comment | 
 **encryptionTypes** | **optional.String**| Filter by encryption_types | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**KerberosRealmResponse**](kerberos_realm_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmCreate

> KerberosRealmCreate(ctx, optional)



Creates a Kerberos realm. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM on which to create the Kerberos realm. * `name` - Base name for the Kerberos realm. * `kdc.vendor` - Vendor of the Key Distribution Center (KDC) server for this Kerberos realm. If the configuration uses a Microsoft Active Directory domain for authentication, this field nust be `microsoft`. * `kdc.ip` - IP address of the KDC server for this Kerberos realm. ### Recommended optional properties * `ad_server.name` - Host name of the Active Directory Domain Controller (DC). This is a mandatory parameter if the kdc-vendor is `microsoft`. * `ad_server.address` - IP address of the Active Directory Domain Controller (DC). This is a mandatory parameter if the kdc-vendor is `microsoft`. ### Default property values If not specified in POST, the following default property value is assigned: * `kdc.port` - _88_ ### Related ONTAP commands * `vserver nfs kerberos realm create` ### Learn more * [`DOC /protocols/nfs/kerberos/realms`](#docs-NAS-protocols_nfs_kerberos_realms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KerberosRealmCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosRealmCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of KerberosRealm**](KerberosRealm.md)| Info specification | 

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


## KerberosRealmDelete

> KerberosRealmDelete(ctx, svmUuid, name)



Deletes a Kerberos realm. * `vserver nfs kerberos realm delete` ### Learn more * [`DOC /protocols/nfs/kerberos/realms`](#docs-NAS-protocols_nfs_kerberos_realms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| SVM UUID | 
**name** | **string**| Kerberos realm | 

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


## KerberosRealmGet

> KerberosRealm KerberosRealmGet(ctx, svmUuid, name, optional)



Retrieves a Kerberos realm. * `vserver nfs kerberos realm show` ### Learn more * [`DOC /protocols/nfs/kerberos/realms`](#docs-NAS-protocols_nfs_kerberos_realms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| SVM UUID | 
**name** | **string**| Kerberos realm | 
 **optional** | ***KerberosRealmGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosRealmGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**KerberosRealm**](kerberos_realm.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KerberosRealmModify

> KerberosRealmModify(ctx, svmUuid, name, optional)



Updates the properties of a Kerberos realm. * `vserver nfs kerberos realm modify` ### Learn more * [`DOC /protocols/nfs/kerberos/realms`](#docs-NAS-protocols_nfs_kerberos_realms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| SVM UUID | 
**name** | **string**| Kerberos realm | 
 **optional** | ***KerberosRealmModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KerberosRealmModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of KerberosRealm**](KerberosRealm.md)| Info specification | 

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


## NfsCollectionGet

> NfsServiceResponse NfsCollectionGet(ctx, optional)



Retrieves the NFS configuration of SVMs. ### Related ONTAP commands * `vserver nfs show` * `vserver nfs status` ### Learn more * [`DOC /protocols/nfs/services`](#docs-NAS-protocols_nfs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NfsCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NfsCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter by state | 
 **protocolV40Enabled** | **optional.Bool**| Filter by protocol.v40_enabled | 
 **protocolV4IdDomain** | **optional.String**| Filter by protocol.v4_id_domain | 
 **protocolV3Enabled** | **optional.Bool**| Filter by protocol.v3_enabled | 
 **protocolV41Enabled** | **optional.Bool**| Filter by protocol.v41_enabled | 
 **protocolV40FeaturesReadDelegationEnabled** | **optional.Bool**| Filter by protocol.v40_features.read_delegation_enabled | 
 **protocolV40FeaturesWriteDelegationEnabled** | **optional.Bool**| Filter by protocol.v40_features.write_delegation_enabled | 
 **protocolV40FeaturesAclEnabled** | **optional.Bool**| Filter by protocol.v40_features.acl_enabled | 
 **protocolV41FeaturesAclEnabled** | **optional.Bool**| Filter by protocol.v41_features.acl_enabled | 
 **protocolV41FeaturesReadDelegationEnabled** | **optional.Bool**| Filter by protocol.v41_features.read_delegation_enabled | 
 **protocolV41FeaturesPnfsEnabled** | **optional.Bool**| Filter by protocol.v41_features.pnfs_enabled | 
 **protocolV41FeaturesWriteDelegationEnabled** | **optional.Bool**| Filter by protocol.v41_features.write_delegation_enabled | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **vstorageEnabled** | **optional.Bool**| Filter by vstorage_enabled | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **transportTcpEnabled** | **optional.Bool**| Filter by transport.tcp_enabled | 
 **transportUdpEnabled** | **optional.Bool**| Filter by transport.udp_enabled | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NfsServiceResponse**](nfs_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfsCreate

> NfsServiceResponse NfsCreate(ctx, optional)



Creates an NFS configuration for an SVM. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM for which to create the NFS configuration. ### Default property values If not specified in POST, the following default property values are assigned: * `enabled` - _true_ * `state` - online * `transport.udp_enabled` - _true_ * `transport.tcp_enabled` - _true_ * `protocol.v3_enabled` - _true_ * `protocol.v4_id_domain` - defaultv4iddomain.com * `protocol.v4_enabled` - _false_ * `protocol.v41_enabled` - _false_ * `protocol.v40_features.acl_enabled` - _false_ * `protocol.v40_features.read_delegation_enabled` - _false_ * `protocol.v40_features.write_delegation_enabled` - _false_ * `protocol.v41_features.acl_enabled` - _false_ * `protocol.v41_features.read_delegation_enabled` - _false_ * `protocol.v41_features.write_delegation_enabled` - _false_ * `protocol.v41_features.pnfs_enabled` - _false_ * `vstorage_enabled` - _false_ ### Related ONTAP commands * `vserver nfs create` ### Learn more * [`DOC /protocols/nfs/services`](#docs-NAS-protocols_nfs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NfsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NfsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of NfsService**](NfsService.md)| Info Specification | 

### Return type

[**NfsServiceResponse**](nfs_service_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfsDelete

> NfsDelete(ctx, svmUuid)



Deletes the NFS configuration of an SVM. ### Related ONTAP commands * `vserver nfs delete` ### Learn more * [`DOC /protocols/nfs/services`](#docs-NAS-protocols_nfs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**|  | 

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


## NfsGet

> NfsService NfsGet(ctx, svmUuid, optional)



Retrieves the NFS configuration of an SVM. ### Related ONTAP commands * `vserver nfs show` * `vserver nfs status` ### Learn more * [`DOC /protocols/nfs/services`](#docs-NAS-protocols_nfs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**|  | 
 **optional** | ***NfsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NfsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**NfsService**](nfs_service.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfsModify

> NfsModify(ctx, svmUuid, optional)



Updates the NFS configuration of an SVM. ### Related ONTAP commands * `vserver nfs modify` * `vserver nfs on` * `vserver nfs off` * `vserver nfs start` * `vserver nfs stop` ### Learn more * [`DOC /protocols/nfs/services`](#docs-NAS-protocols_nfs_services) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**|  | 
 **optional** | ***NfsModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NfsModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of NfsService**](NfsService.md)| Info Specification | 

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


## VscanCollectionGet

> VscanResponse VscanCollectionGet(ctx, optional)



Retrieves the Vscan configuration. This includes scanner-pools, On-Access policies, On-Demand policies, and information about whether a Vscan is enabled or disabled on an SVM.<br/> Important notes: * There can be only one Vscan configuration enabled for an SVM at any time. * You can only query using `svm.uuid` or `svm.name`. ### Related ONTAP commands * `vserver vscan show` * `vserver vscan scanner-pool show` * `vserver vscan scanner-pool servers show` * `vserver vscan scanner-pool privileged-users show` * `vserver vscan scanner-pool show-active` * `vserver vscan on-access-policy show` * `vserver vscan on-access-policy file-ext-to-exclude show` * `vserver vscan on-access-policy file-ext-to-include show` * `vserver vscan on-access-policy paths-to-exclude show` * `vserver vscan on-demand-task show` ### Learn more * [`DOC /protocols/vscan`](#docs-NAS-protocols_vscan) * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VscanCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onDemandPoliciesScopeScanWithoutExtension** | **optional.Bool**| Filter by on_demand_policies.scope.scan_without_extension | 
 **onDemandPoliciesScopeExcludePaths** | **optional.String**| Filter by on_demand_policies.scope.exclude_paths | 
 **onDemandPoliciesScopeExcludeExtensions** | **optional.String**| Filter by on_demand_policies.scope.exclude_extensions | 
 **onDemandPoliciesScopeIncludeExtensions** | **optional.String**| Filter by on_demand_policies.scope.include_extensions | 
 **onDemandPoliciesScopeMaxFileSize** | **optional.Int32**| Filter by on_demand_policies.scope.max_file_size | 
 **onDemandPoliciesName** | **optional.String**| Filter by on_demand_policies.name | 
 **onDemandPoliciesScanPaths** | **optional.String**| Filter by on_demand_policies.scan_paths | 
 **onDemandPoliciesLogPath** | **optional.String**| Filter by on_demand_policies.log_path | 
 **onDemandPoliciesScheduleUuid** | **optional.String**| Filter by on_demand_policies.schedule.uuid | 
 **onDemandPoliciesScheduleName** | **optional.String**| Filter by on_demand_policies.schedule.name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **scannerPoolsPrivilegedUsers** | **optional.String**| Filter by scanner_pools.privileged_users | 
 **scannerPoolsName** | **optional.String**| Filter by scanner_pools.name | 
 **scannerPoolsClusterUuid** | **optional.String**| Filter by scanner_pools.cluster.uuid | 
 **scannerPoolsClusterName** | **optional.String**| Filter by scanner_pools.cluster.name | 
 **scannerPoolsServers** | **optional.String**| Filter by scanner_pools.servers | 
 **scannerPoolsRole** | **optional.String**| Filter by scanner_pools.role | 
 **onAccessPoliciesScopeScanReadonlyVolumes** | **optional.Bool**| Filter by on_access_policies.scope.scan_readonly_volumes | 
 **onAccessPoliciesScopeIncludeExtensions** | **optional.String**| Filter by on_access_policies.scope.include_extensions | 
 **onAccessPoliciesScopeMaxFileSize** | **optional.Int32**| Filter by on_access_policies.scope.max_file_size | 
 **onAccessPoliciesScopeExcludeExtensions** | **optional.String**| Filter by on_access_policies.scope.exclude_extensions | 
 **onAccessPoliciesScopeOnlyExecuteAccess** | **optional.Bool**| Filter by on_access_policies.scope.only_execute_access | 
 **onAccessPoliciesScopeScanWithoutExtension** | **optional.Bool**| Filter by on_access_policies.scope.scan_without_extension | 
 **onAccessPoliciesScopeExcludePaths** | **optional.String**| Filter by on_access_policies.scope.exclude_paths | 
 **onAccessPoliciesEnabled** | **optional.Bool**| Filter by on_access_policies.enabled | 
 **onAccessPoliciesName** | **optional.String**| Filter by on_access_policies.name | 
 **onAccessPoliciesMandatory** | **optional.Bool**| Filter by on_access_policies.mandatory | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**VscanResponse**](vscan_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanConfigDelete

> VscanConfigDelete(ctx, svmUuid)



Deletes a Vscan configuration.<br/> Important notes: * The Vscan DELETE endpoint deletes all of the Vscan configuration of an SVM. It first disables the Vscan and then deletes all of the SVM scanner-pools, On-Access policies, and On-Demand policies. * Any active Vscan On-Access policy must first be disabled on an SVM before performing the Vscan delete operation on that SVM. ### Related ONTAP commands * `vserver vscan scanner-pool delete` * `vserver vscan on-access-policy delete` * `vserver vscan on-demand-policy delete` ### Learn more * [`DOC /protocols/vscan`](#docs-NAS-protocols_vscan) * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

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


## VscanCreate

> VscanResponse VscanCreate(ctx, optional)



Creates a Vscan configuration, which includes a list of scanner-pools, Vscan On-Access policies and Vscan On-Demand policies. Defines whether the Vscan configuration you’re creating is enabled or disabled for a specified SVM.<br/> Important notes: * There can be only one Vscan configuration enabled for an SVM at any time. * There needs to be at least one active scanner-pool and one enabled On-Access policy for Vscan to be enabled successfully. * By default, a Vscan is enabled when it’s created. * By default, the Vscan On-Access policies created from this endpoint are in the disabled state. You can use the On-Access policy PATCH endpoint to enable a particular On-Access policy. In ONTAP 9.6, only one Vscan On-Access policy can be enabled and only one Vscan On-Demand policy can be scheduled on an SVM. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the Vscan configuration. ### Recommended optional properties * `scanner_pools` - There must be at least one active scanner-pool for Vscan configuration. Created either through Vscan POST operation or scanner-pools POST operation. ### Default property values If not specified in POST, the following default property value is assigned: * `enabled` - _true_ ### Related ONTAP commands * `vserver vscan enable` * `vserver vscan scanner-pool create` * `vserver vscan scanner-pool apply-policy` * `vserver vscan scanner-pool servers add` * `vserver vscan scanner-pool privileged-users add` * `vserver vscan on-access-policy create` * `vserver vscan on-access-policy file-ext-to-exclude add` * `vserver vscan on-access-policy file-ext-to-include add` * `vserver vscan on-access-policy paths-to-exclude add` * `vserver vscan on-demand-task create` ### Learn more * [`DOC /protocols/vscan`](#docs-NAS-protocols_vscan) * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VscanCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Vscan**](Vscan.md)| Info specification | 

### Return type

[**VscanResponse**](vscan_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanGet

> Vscan VscanGet(ctx, svmUuid, optional)



Retrieves the Vscan configuration for a specified SVM. This includes scanner-pools, On-Access policies, On-Demand policies, and information about whether a Vscan is enabled or disabled on an SVM.<br/> Important note: * There can be only one Vscan configuration enabled for an SVM at any time. ### Related ONTAP commands * `vserver vscan show` * `vserver vscan scanner-pool show` * `vserver vscan scanner-pool servers show` * `vserver vscan scanner-pool privileged-users show` * `vserver vscan scanner-pool show-active` * `vserver vscan on-access-policy show` * `vserver vscan on-access-policy file-ext-to-exclude show` * `vserver vscan on-access-policy file-ext-to-include show` * `vserver vscan on-access-policy paths-to-exclude show` * `vserver vscan on-demand-task show` ### Learn more * [`DOC /protocols/vscan`](#docs-NAS-protocols_vscan) * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Vscan**](vscan.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanModify

> VscanModify(ctx, svmUuid, optional)



Updates the Vscan configuration of an SVM. Allows you to either enable or disable a Vscan, and allows you to clear the Vscan cache that stores the past scanning data for an SVM.<br/> Important note: * The Vscan PATCH endpoint does not allow you to modify scanner-pools, On-Demand policies or On-Access policies. Those modifications can only be done through their respective endpoints. ### Related ONTAP commands * `vserver vscan enable` * `vserver vscan disable` * `vserver vscan reset` ### Learn more * [`DOC /protocols/vscan`](#docs-NAS-protocols_vscan) * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of Vscan**](Vscan.md)| Info specification | 

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


## VscanOnAccessCreate

> VscanOnAccessResponse VscanOnAccessCreate(ctx, svmUuid, optional)



Creates a Vscan On-Access policy. Created only on a data SVM. </b>Important notes: * The policy needs to be enabled on an SVM before its files can be scanned. * Only one On-Access policy can be enabled on an SVM at a time. By default, the policy is enabled on creation. * If the Vscan On-Access policy has been created successfully on an SVM but cannot be enabled due to an error, the Vscan On-Access policy configurations are saved. The Vscan On-Access policy is then enabled using the PATCH operation. ### Required properties * `svm.uuid` - Existing SVM in which to create the Vscan On-Access policy. * `name` - Name of the Vscan On-Access policy. Maximum length is 256 characters. ### Default property values If not specified in POST, the following default property values are assigned: * `enabled` - _true_ * `mandatory` - _true_ * `include_extensions` - _*_ * `max_file_size` - _2147483648_ * `only_execute_access` - _false_ * `scan_readonly_volumes` - _false_ * `scan_without_extension` - _true_ ### Related ONTAP commands * `vserver vscan on-access-policy create` * `vserver vscan on-access-policy enable` * `vserver vscan on-access-policy disable` * `vserver vscan on-access-policy file-ext-to-include add` * `vserver vscan on-access-policy file-ext-to-exclude add` * `vserver vscan on-access-policy paths-to-exclude add` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-access-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-access-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanOnAccessCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnAccessCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of VscanOnAccess**](VscanOnAccess.md)| Info specification | 

### Return type

[**VscanOnAccessResponse**](vscan_on_access_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanOnAccessDelete

> VscanOnAccessDelete(ctx, svmUuid, name)



Deletes the anti-virus On-Access policy configuration. ### Related ONTAP commands * `vserver vscan on-access-policy delete` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-access-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-access-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 

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


## VscanOnAccessGet

> VscanOnAccess VscanOnAccessGet(ctx, svmUuid, name, optional)



Retrieves the Vscan On-Access policy configuration of an SVM. ### Related ONTAP commands * `vserver vscan on-access-policy show` * `vserver vscan on-access-policy file-ext-to-include show` * `vserver vscan on-access-policy file-ext-to-exclude show` * `vserver vscan on-access-policy paths-to-exclude show` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-access-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-access-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***VscanOnAccessGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnAccessGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**VscanOnAccess**](vscan_on_access.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanOnAccessModify

> VscanOnAccessModify(ctx, svmUuid, name, optional)



Updates the Vscan On-Access policy configuration and/or enables/disables the Vscan On-Access policy of an SVM. Configurations for an On-Access policy associated with an administrative SVM cannot be modified, although the policy associated with an administrative SVM can be enabled or disabled. ### Related ONTAP commands * `vserver vscan on-access-policy modify` * `vserver vscan on-access-policy enable` * `vserver vscan on-access-policy disable` * `vserver vscan on-access-policy file-ext-to-include add` * `vserver vscan on-access-policy file-ext-to-exclude add` * `vserver vscan on-access-policy paths-to-exclude add` * `vserver vscan on-access-policy file-ext-to-include remove` * `vserver vscan on-access-policy file-ext-to-exclude remove` * `vserver vscan on-access-policy paths-to-exclude remove` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-access-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-access-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***VscanOnAccessModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnAccessModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of VscanOnAccess**](VscanOnAccess.md)| Info specification | 

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


## VscanOnAccessPolicyCollectionGet

> VscanOnAccessResponse VscanOnAccessPolicyCollectionGet(ctx, svmUuid, optional)



Retrieves the Vscan On-Access policy. ### Related ONTAP commands * `vserver vscan on-access-policy show` * `vserver vscan on-access-policy file-ext-to-include show` * `vserver vscan on-access-policy file-ext-to-exclude show` * `vserver vscan on-access-policy paths-to-exclude show` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-access-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-access-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanOnAccessPolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnAccessPolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeScanReadonlyVolumes** | **optional.Bool**| Filter by scope.scan_readonly_volumes | 
 **scopeIncludeExtensions** | **optional.String**| Filter by scope.include_extensions | 
 **scopeMaxFileSize** | **optional.Int32**| Filter by scope.max_file_size | 
 **scopeExcludeExtensions** | **optional.String**| Filter by scope.exclude_extensions | 
 **scopeOnlyExecuteAccess** | **optional.Bool**| Filter by scope.only_execute_access | 
 **scopeScanWithoutExtension** | **optional.Bool**| Filter by scope.scan_without_extension | 
 **scopeExcludePaths** | **optional.String**| Filter by scope.exclude_paths | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **name** | **optional.String**| Filter by name | 
 **mandatory** | **optional.Bool**| Filter by mandatory | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**VscanOnAccessResponse**](vscan_on_access_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanOnDemandCreate

> VscanOnDemandResponse VscanOnDemandCreate(ctx, svmUuid, optional)



Creates a Vscan On-Demand policy. Created only on a data SVM. </br> Important notes:   * Only one policy can be scheduled at a time on an SVM. Use schedule name or schedule uuid to schedule an On-Demand policy.   * Scanning must be enabled on the SVM before the policy is scheduled to run.   * The exclude_extensions setting overrides the include_extensions setting. Set scan_without_extension to true to scan files without extensions. ### Required properties * `svm.uuid` - Existing SVM in which to create the Vscan On-Demand policy. * `name` - Name of the Vscan On-Demand policy. Maximum length is 256 characters. * `log_path` - Path from the Vserver root where the On-Demand policy report is created. * `scan_paths` - List of paths that need to be scanned. ### Recommended optional properties * `schedule` - Scan schedule. It is recommended to set the schedule property, as it dictates when to scan for viruses. ### Default property values If not specified in POST, the following default property values are assigned: * `include_extensions` - _*_ * `max_file_size` - _10737418240_ * `scan_without_extension` - _true_ ### Related ONTAP commands * `vserver vscan on-demand-task create` * `vserver vscan on-demand-task schedule` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-demand-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-demand-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanOnDemandCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnDemandCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of VscanOnDemand**](VscanOnDemand.md)| Info specification | 

### Return type

[**VscanOnDemandResponse**](vscan_on_demand_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanOnDemandDelete

> VscanOnDemandDelete(ctx, svmUuid, name)



Deletes the Vscan On-Demand configuration. ### Related ONTAP commands * `vserver vscan on-demand-task delete` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-demand-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-demand-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 

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


## VscanOnDemandGet

> VscanOnDemand VscanOnDemandGet(ctx, svmUuid, name, optional)



Retrieves the Vscan On-Demand configuration of an SVM. ### Related ONTAP commands * `vserver vscan on-demand-task show` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-demand-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-demand-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***VscanOnDemandGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnDemandGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**VscanOnDemand**](vscan_on_demand.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanOnDemandModify

> VscanOnDemandModify(ctx, svmUuid, name, optional)



Updates the Vscan On-Demand policy configuration of an SVM. Use schedule name or schedule UUID to schedule an On-Demand scan. ### Related ONTAP commands * `vserver vscan on-demand-task modify` * `vserver vscan on-demand-task schedule` * `vserver vscan on-demand-task unschedule` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-demand-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-demand-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***VscanOnDemandModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnDemandModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of VscanOnDemand**](VscanOnDemand.md)| Info specification | 

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


## VscanOnDemandPolicyCollectionGet

> VscanOnDemandResponse VscanOnDemandPolicyCollectionGet(ctx, svmUuid, optional)



Retrieves the Vscan On-Demand policy. ### Related ONTAP commands * `vserver vscan on-demand-task show` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/on-demand-policies`](#docs-NAS-protocols_vscan_{svm.uuid}_on-demand-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanOnDemandPolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanOnDemandPolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeScanWithoutExtension** | **optional.Bool**| Filter by scope.scan_without_extension | 
 **scopeExcludePaths** | **optional.String**| Filter by scope.exclude_paths | 
 **scopeExcludeExtensions** | **optional.String**| Filter by scope.exclude_extensions | 
 **scopeIncludeExtensions** | **optional.String**| Filter by scope.include_extensions | 
 **scopeMaxFileSize** | **optional.Int32**| Filter by scope.max_file_size | 
 **name** | **optional.String**| Filter by name | 
 **scanPaths** | **optional.String**| Filter by scan_paths | 
 **logPath** | **optional.String**| Filter by log_path | 
 **scheduleUuid** | **optional.String**| Filter by schedule.uuid | 
 **scheduleName** | **optional.String**| Filter by schedule.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**VscanOnDemandResponse**](vscan_on_demand_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanScannerCollectionGet

> VscanScannerPoolResponse VscanScannerCollectionGet(ctx, svmUuid, optional)



Retrieves the Vscan scanner-pool configuration of an SVM. ### Related ONTAP commands * `vserver vscan scanner-pool show` * `vserver vscan scanner-pool privileged-users show` * `vserver vscan scanner-pool servers show` * `vserver vscan scanner-pool show-active` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanScannerCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanScannerCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privilegedUsers** | **optional.String**| Filter by privileged_users | 
 **name** | **optional.String**| Filter by name | 
 **clusterUuid** | **optional.String**| Filter by cluster.uuid | 
 **clusterName** | **optional.String**| Filter by cluster.name | 
 **servers** | **optional.String**| Filter by servers | 
 **role** | **optional.String**| Filter by role | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**VscanScannerPoolResponse**](vscan_scanner_pool_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanScannerCreate

> VscanScannerPoolResponse VscanScannerCreate(ctx, svmUuid, optional)



Creates a Vscan scanner-pool configuration for a specified SVM. A scanner-pool can be created with all fields specified or only mandatory fields specified.<br/> Important notes: * A scanner-pool must have servers and privileged users specified. * If the role or cluster is not specified, the scanner-pool is created on the local cluster with the role set as primary. *`Only one of the fields cluster-uuid or cluster-name is required. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the Vscan configuration. * `name` - Scanner-pool name. * `privileged_users` - List of privileged users. * `servers` - List of server IP addresses or FQDNs. ### Recommended optional properties * `role` - Setting a role for a scanner-pool is recommended. * `cluster` - Passing the cluster name or UUID (or both) in a multi-cluster environment is recommended. ### Default property values If not specified in POST, the following default property values are assigned: * `role` - _primary_ * `cluster.name` - Local cluster name. * `cluster.uuid` - Local cluster UUID. ### Related ONTAP commands * `vserver vscan scanner-pool create` * `vserver vscan scanner-pool apply-policy` * `vserver vscan scanner-pool privileged-users add` * `vserver vscan scanner-pool servers add` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
 **optional** | ***VscanScannerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanScannerCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of VscanScannerPool**](VscanScannerPool.md)| Info specification | 

### Return type

[**VscanScannerPoolResponse**](vscan_scanner_pool_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanScannerDelete

> VscanScannerDelete(ctx, svmUuid, name)



Deletes a Vscan scanner-pool configuration.<br/> Important notes: * The Vscan scanner-pool DELETE endpoint deletes all of the Vscan scanner-pools for a specified SVM. * If a Vscan is enabled, it requires at least one scanner-pool to be in the active state. Therefore, Vscan must be disabled on the specified SVM so that all of the scanner-pools configured on that SVM can be deleted. ### Related ONTAP commands * `vserver vscan scanner-pool delete` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 

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


## VscanScannerGet

> VscanScannerPool VscanScannerGet(ctx, svmUuid, name, optional)



Retrieves the configuration of a specified scanner-pool of an SVM. ### Related ONTAP commands * `vserver vscan scanner-pool show` * `vserver vscan scanner-pool privileged-users show` * `vserver vscan scanner-pool servers show` * `vserver vscan scanner-pool show-active` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***VscanScannerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanScannerGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**VscanScannerPool**](vscan_scanner_pool.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VscanScannerModify

> VscanScannerModify(ctx, svmUuid, name, optional)



Updates the Vscan scanner-pool configuration of an SVM.<br/> Important notes: * Along with servers and privileged-users, the role of a scanner-pool can also be updated with the cluster on which a scanner-pool is allowed. * If role is specified and cluster isn't, then role is applied to the local cluster. ### Related ONTAP commands * `vserver vscan scanner-pool modify` * `vserver vscan scanner-pool apply-policy` * `vserver vscan scanner-pool privileged-users add` * `vserver vscan scanner-pool privileged-users remove` * `vserver vscan scanner-pool servers remove` * `vserver vscan scanner-pool servers add` ### Learn more * [`DOC /protocols/vscan/{svm.uuid}/scanner-pools`](#docs-NAS-protocols_vscan_{svm.uuid}_scanner-pools) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svmUuid** | **string**| UUID of the SVM to which this object belongs. | 
**name** | **string**|  | 
 **optional** | ***VscanScannerModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanScannerModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of VscanScannerPool**](VscanScannerPool.md)| Info specification | 

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


## VscanServerStatusGet

> VscanServerStatusResponse VscanServerStatusGet(ctx, optional)



Retrieves a Vscan server status. ### Related ONTAP commands * `vserver vscan connection-status show-all` ### Learn more * [`DOC /protocols/vscan/server-status`](#docs-NAS-protocols_vscan_server-status) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VscanServerStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VscanServerStatusGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter by state | 
 **ip** | **optional.String**| Filter by ip | 
 **version** | **optional.String**| Filter by version | 
 **disconnectedReason** | **optional.String**| Filter by disconnected_reason | 
 **type_** | **optional.String**| Filter by type | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **vendor** | **optional.String**| Filter by vendor | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **updateTime** | **optional.String**| Filter by update_time | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**VscanServerStatusResponse**](vscan_server_status_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

