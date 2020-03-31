# \CloudApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudTargetCollectionGet**](CloudApi.md#CloudTargetCollectionGet) | **Get** /cloud/targets | 
[**CloudTargetCreate**](CloudApi.md#CloudTargetCreate) | **Post** /cloud/targets | 
[**CloudTargetDelete**](CloudApi.md#CloudTargetDelete) | **Delete** /cloud/targets/{uuid} | 
[**CloudTargetGet**](CloudApi.md#CloudTargetGet) | **Get** /cloud/targets/{uuid} | 
[**CloudTargetModify**](CloudApi.md#CloudTargetModify) | **Patch** /cloud/targets/{uuid} | 



## CloudTargetCollectionGet

> CloudTargetResponse CloudTargetCollectionGet(ctx, optional)



Retrieves the collection of cloud targets in the cluster. ### Related ONTAP commands * `storage aggregate object-store config show`  ### Learn more * [`DOC /cloud/targets`](#docs-cloud-cloud_targets)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudTargetCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudTargetCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **authenticationType** | **optional.String**| Filter by authentication_type | 
 **port** | **optional.Int32**| Filter by port | 
 **used** | **optional.Int32**| Filter by used | 
 **capUrl** | **optional.String**| Filter by cap_url | 
 **snapmirrorUse** | **optional.String**| Filter by snapmirror_use | 
 **ipspaceUuid** | **optional.String**| Filter by ipspace.uuid | 
 **ipspaceName** | **optional.String**| Filter by ipspace.name | 
 **azureAccount** | **optional.String**| Filter by azure_account | 
 **owner** | **optional.String**| Filter by owner | 
 **container** | **optional.String**| Filter by container | 
 **server** | **optional.String**| Filter by server | 
 **uuid** | **optional.String**| Filter by uuid | 
 **sslEnabled** | **optional.Bool**| Filter by ssl_enabled | 
 **certificateValidationEnabled** | **optional.Bool**| Filter by certificate_validation_enabled | 
 **providerType** | **optional.String**| Filter by provider_type | 
 **accessKey** | **optional.String**| Filter by access_key | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CloudTargetResponse**](cloud_target_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudTargetCreate

> JobLinkResponse CloudTargetCreate(ctx, optional)



Creates a cloud target. ### Required properties * `name` - Name for the cloud target. * `owner` - Owner of the target: _fabricpool_, _snapmirror_. * `provider_type` - Type of cloud provider: _AWS_S3_, _Azure_Cloud_, _SGWS_, _IBM_COS_, _AliCloud_, _GoogleCloud_. * `server` - Fully qualified domain name of the object store server. Required when `provider_type` is one of the following: _SGWS_, _IBM_COS_, _AliCloud_. * `container` - Data bucket/container name. * `access_key` - Access key ID if `provider_type` is not _Azure_Cloud_ and `authentication_type` is _key_. * `secret_password` - Secret access key if `provider_type` is not _Azure_Cloud_ and `authentication_type` is _key_. * `azure_account` - Azure account if `provider_type` is _Azure_Cloud_. * `azure_private_key` - Azure access key if `provider_type` is _Azure_Cloud_. * `cap_url` - Full URL of the request to a CAP server for retrieving temporary credentials if `authentication_type` is _cap_. * `svm.name` or `svm.uuid` - Name or UUID of SVM if `owner` is _snapmirror_. * `snapmirror_use` - Use of the cloud target if `owner` is _snapmirror_: data, metadata. ### Recommended optional properties * `authentication_type` - Authentication used to access the target: _key_, _cap_, _ec2_iam_. * `ssl_enabled` - SSL/HTTPS enabled or disabled. * `port` - Port number of the object store that ONTAP uses when establishing a connection. * `ipspace` - IPspace to use in order to reach the cloud target. ### Default property values * `authentication_type`   - _ec2_iam_ - if running in Cloud Volumes ONTAP in AWS   - _key_  - in all other cases. * `server`   - _s3.amazonaws.com_ - if `provider_type` is _AWS_S3_   - _blob.core.windows.net_ - if `provider_type` is _Azure_Cloud_   - _storage.googleapis.com_ - if `provider_type` is _GoogleCloud_ * `ssl_enabled` - _true_ * `port`   - _443_ if `ssl_enabled` is _true_ and `provider_type` is not _SGWS_   - _8082_ if `ssl_enabled` is _true_ and `provider_type` is _SGWS_   - _80_ if `ssl_enabled` is _false_ and `provider_type` is not _SGWS_   - _8084_ if `ssl_enabled` is _false_ and `provider_type` is _SGWS_ * `ipspace` - _Default_ * `certificate_validation_enabled` - _true_ * `ignore_warnings` - _false_ * `check_only` - _false_ ### Related ONTAP commands * `storage aggregate object-store config create`  ### Learn more * [`DOC /cloud/targets`](#docs-cloud-cloud_targets)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudTargetCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudTargetCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoreWarnings** | **optional.Bool**| Specifies whether or not warning codes should be ignored. | 
 **checkOnly** | **optional.Bool**| Do not create the target configuration, only check that the POST request succeeds. | 
 **info** | [**optional.Interface of CloudTarget**](CloudTarget.md)| Info specification | 

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


## CloudTargetDelete

> JobLinkResponse CloudTargetDelete(ctx, uuid)



Deletes the cloud target specified by the UUID. This request starts a job and returns a link to that job. ### Related ONTAP commands * `storage aggregate object-store config delete`  ### Learn more * [`DOC /cloud/targets`](#docs-cloud-cloud_targets)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Cloud target UUID | 

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


## CloudTargetGet

> CloudTarget CloudTargetGet(ctx, uuid, optional)



Retrieves the cloud target specified by the UUID. ### Related ONTAP commands * `storage aggregate object-store config show`  ### Learn more * [`DOC /cloud/targets`](#docs-cloud-cloud_targets)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Cloud target UUID | 
 **optional** | ***CloudTargetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudTargetGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CloudTarget**](cloud_target.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudTargetModify

> JobLinkResponse CloudTargetModify(ctx, uuid, optional)



Updates the cloud target specified by the UUID with the fields in the body. This request starts a job and returns a link to that job. ### Related ONTAP commands * `storage aggregate object-store config modify`  ### Learn more * [`DOC /cloud/targets`](#docs-cloud-cloud_targets)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Cloud target UUID | 
 **optional** | ***CloudTargetModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudTargetModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ignoreWarnings** | **optional.Bool**| Specifies whether or not warnings should be ignored. | 
 **checkOnly** | **optional.Bool**| Do not modify the configuration, only check that the PATCH request succeeds. | 
 **info** | [**optional.Interface of CloudTarget**](CloudTarget.md)| Info specification | 

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

