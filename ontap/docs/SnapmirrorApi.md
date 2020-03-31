# \SnapmirrorApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnapmirrorPoliciesGet**](SnapmirrorApi.md#SnapmirrorPoliciesGet) | **Get** /snapmirror/policies | 
[**SnapmirrorPolicyCreate**](SnapmirrorApi.md#SnapmirrorPolicyCreate) | **Post** /snapmirror/policies | 
[**SnapmirrorPolicyDelete**](SnapmirrorApi.md#SnapmirrorPolicyDelete) | **Delete** /snapmirror/policies/{uuid} | 
[**SnapmirrorPolicyGet**](SnapmirrorApi.md#SnapmirrorPolicyGet) | **Get** /snapmirror/policies/{uuid} | 
[**SnapmirrorPolicyModify**](SnapmirrorApi.md#SnapmirrorPolicyModify) | **Patch** /snapmirror/policies/{uuid} | 
[**SnapmirrorRelationshipCreate**](SnapmirrorApi.md#SnapmirrorRelationshipCreate) | **Post** /snapmirror/relationships | 
[**SnapmirrorRelationshipDelete**](SnapmirrorApi.md#SnapmirrorRelationshipDelete) | **Delete** /snapmirror/relationships/{uuid} | 
[**SnapmirrorRelationshipGet**](SnapmirrorApi.md#SnapmirrorRelationshipGet) | **Get** /snapmirror/relationships/{uuid} | 
[**SnapmirrorRelationshipModify**](SnapmirrorApi.md#SnapmirrorRelationshipModify) | **Patch** /snapmirror/relationships/{uuid} | 
[**SnapmirrorRelationshipTransferCreate**](SnapmirrorApi.md#SnapmirrorRelationshipTransferCreate) | **Post** /snapmirror/relationships/{relationship.uuid}/transfers | 
[**SnapmirrorRelationshipTransferGet**](SnapmirrorApi.md#SnapmirrorRelationshipTransferGet) | **Get** /snapmirror/relationships/{relationship.uuid}/transfers/{uuid} | 
[**SnapmirrorRelationshipTransferModify**](SnapmirrorApi.md#SnapmirrorRelationshipTransferModify) | **Patch** /snapmirror/relationships/{relationship.uuid}/transfers/{uuid} | 
[**SnapmirrorRelationshipTransfersGet**](SnapmirrorApi.md#SnapmirrorRelationshipTransfersGet) | **Get** /snapmirror/relationships/{relationship.uuid}/transfers | 
[**SnapmirrorRelationshipsGet**](SnapmirrorApi.md#SnapmirrorRelationshipsGet) | **Get** /snapmirror/relationships | 



## SnapmirrorPoliciesGet

> SnapmirrorPolicyResponse SnapmirrorPoliciesGet(ctx, optional)



Retrieves SnapMirror policies of type \"mirror-vault\", \"sync-mirror\" and \"strict-sync-mirror\". ### Related ONTAP commands * `snapmirror policy show` ### Example The following example shows how to retrieve a collection of SnapMirror policies. <br/> ``` GET \"/api/storage/snapmirror/policies\" ``` <br/> ### Learn more * [`DOC /snapmirror/policies`](#docs-snapmirror-snapmirror_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapmirrorPoliciesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorPoliciesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncType** | **optional.String**| Filter by sync_type | 
 **networkCompressionEnabled** | **optional.Bool**| Filter by network_compression_enabled | 
 **transferScheduleUuid** | **optional.String**| Filter by transfer_schedule.uuid | 
 **transferScheduleName** | **optional.String**| Filter by transfer_schedule.name | 
 **type_** | **optional.String**| Filter by type | 
 **throttle** | **optional.Int32**| Filter by throttle | 
 **syncCommonSnapshotScheduleUuid** | **optional.String**| Filter by sync_common_snapshot_schedule.uuid | 
 **syncCommonSnapshotScheduleName** | **optional.String**| Filter by sync_common_snapshot_schedule.name | 
 **scope** | **optional.String**| Filter by scope | 
 **comment** | **optional.String**| Filter by comment | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **identityPreservation** | **optional.String**| Filter by identity_preservation | 
 **retentionPrefix** | **optional.String**| Filter by retention.prefix | 
 **retentionCreationScheduleUuid** | **optional.String**| Filter by retention.creation_schedule.uuid | 
 **retentionCreationScheduleName** | **optional.String**| Filter by retention.creation_schedule.name | 
 **retentionLabel** | **optional.String**| Filter by retention.label | 
 **retentionCount** | **optional.Int32**| Filter by retention.count | 
 **uuid** | **optional.String**| Filter by uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SnapmirrorPolicyResponse**](snapmirror_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapmirrorPolicyCreate

> JobLinkResponse SnapmirrorPolicyCreate(ctx, info, optional)



Creates a SnapMirror policy. The parameter \"identity_preservation\" is applicable to only SnapMirror relationships with SVM endpoints and it indicates which configuration of the source SVM is replicated to the destination SVM.</br> It takes the following values: - `full` - indicates that the source SVM configuration is replicated to the destination SVM endpoint. - `exclude_network_config` - indicates that the source SVM configuration other than network configuration is replicated to the destination SVM endpoint. - `exclude_network_and_protocol_config` - indicates that the source SVM configuration is not replicated to the destination SVM endpoint.<br/> ### Important note - The property \"identity_preservation\" is applicable to only SnapMirror relationships with SVM endpoints and it indicates which configuration of the source SVM is replicated to the destination SVM. - The properties \"identity_preservation\", \"retention\" and \"transfer_schedule\" are not applicable for \"sync\" type policies. - The property \"sync_common_snapshot_schedule\" is not applicabe for an \"async\" type policy. - The property \"retention.count\" specifies the maximum number of Snapshot copies that are retained on the SnapMirror destination volume. - When the property \"retention.label\" is specified, the Snapshot copies that have a SnapMirror label matching this property is transferred to the SnapMirror destination. - When the property \"retention.creation_schedule\" is specified, Snapshot copies are directly created on the SnapMirror destination. The Snapshot copies created have the same content as the latest Snapshot copy already present on the SnapMirror destination. ### Required properties * `name` - Name of the new SnapMirror policy. ### Recommended optional properties * `svm.name` or `svm.uuid` - Name or UUID of the SVM that owns the SnapMirror policy. ### Default property values If not specified in POST, the following default property values are assigned: * `type` - _async_ * `sync_type` - _sync_ (when `type` is _sync_) * `network_compression_enabled` - _false_ * `throttle` - _0_ * `identity_preservation` - `_exclude_network_and_protocol_config_` ### Related ONTAP commands * `snapmirror policy create` ### Examples   Creating a SnapMirror policy of type \"sync\"    <br/>    ```    POST \"/api/snapmirror/policies/\" '{\"name\": \"policy1\", \"svm.name\": \"VS0\", \"type\": \"sync\", \"sync_type\": \"sync\"}'    ```    <br/>   Creating a SnapMirror policy of type \"async\" with retention values    <br/>    ```    POST \"/api/snapmirror/policies\" '{\"name\": \"policy_ret\", \"svm\": {\"name\": \"vs1\"}, \"retention\": {\"label\": [\"smcreate\"], \"count\": [\"2\"], \"creation_schedule\": [\"weekly\"]}}'    ```    <br/>   Creating a SnapMirror policy of type \"async\"    ```    POST \"/api/snapmirror/policies\" '{\"name\": \"newPolicy\", \"svm\":{\"name\" : \"vs1\"}, \"type\": \"async\"}'    ```    <br/> ### Learn more * [`DOC /snapmirror/policies`](#docs-snapmirror-snapmirror_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**SnapmirrorPolicy**](SnapmirrorPolicy.md)| Information on the SnapMirror policy | 
 **optional** | ***SnapmirrorPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorPolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]

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


## SnapmirrorPolicyDelete

> JobLinkResponse SnapmirrorPolicyDelete(ctx, uuid, optional)



Deletes a SnapMirror policy. ### Related ONTAP commands * `snapmirror policy delete` ### Example <br/> ``` DELETE \"/api/snapmirror/policies/510c15d4-f9e6-11e8-bdb5-0050568e12c2\" ``` <br/> ### Learn more * [`DOC /snapmirror/policies`](#docs-snapmirror-snapmirror_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Policy UUID | 
 **optional** | ***SnapmirrorPolicyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorPolicyDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]

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


## SnapmirrorPolicyGet

> SnapmirrorPolicy SnapmirrorPolicyGet(ctx, uuid, optional)



Retrieves a specific SnapMirror policy. ### Example <br/> ``` GET \"/api/snapmirror/policies/567aaac0-f863-11e8-a666-0050568e12c2\" ``` <br/> ### Learn more * [`DOC /snapmirror/policies`](#docs-snapmirror-snapmirror_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Policy UUID | 
 **optional** | ***SnapmirrorPolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorPolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SnapmirrorPolicy**](snapmirror_policy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapmirrorPolicyModify

> JobLinkResponse SnapmirrorPolicyModify(ctx, uuid, info)



Updates the SnapMirror policy. ### Important notes * The properties \"transfer_schedule\" and \"throttle\" can be modified only if all the SnapMirror relationships associated with the specified SnapMirror policy have the same values. * The properties \"retention.label\" and \"retention.count\" are mandatory if \"retention\" is provided in the input. The provided \"retention.label\" is the final list and is replaced with the existing values. * The value of the \"identity_preservation\" property cannot be changed if the SnapMirror relationships associated with the policy have different identity_preservation configurations. * If the SnapMirror policy \"identity_preservation\" value matches the \"identity_preservation\" value of the associated SnapMirror relationships, then the \"identity_preservation\" value can be changed from a higher \"identity_preservation\" threshold value to a lower \"identity_preservation\" threshold value but not vice-versa. For example, the threshold value of the \"identity_preservation\" property can be changed from \"full\" to \"exclude_network_config\" to \"exclude_network_and_protocol_config\", but could not be increased from \"exclude_network_and_protocol_config\" to \"exclude_network_config\" to \"full\".<br/> ### Related ONTAP commands * `snapmirror policy modify` ### Example   Updating the \"retention\" property    <br/>    ```    PATCH \"/api/snapmirror/policies/fe65686d-00dc-11e9-b5fb-0050568e3f83\" '{\"retention\" : {\"label\" : [\"sm_created\", \"lab2\"], \"count\": [\"1\",\"2\"], \"creation_schedule\": {\"name\": [\"weekly\"]}}}'    ```    <br/>   Updating \"transfer_schedule\", \"throttle\", and \"identity_preservation\" properties    <br/>    ```    PATCH \"/api/snapmirror/policies/8aef950b-3bef-11e9-80ac-0050568ea591\" '{\"transfer_schedule.name\" : \"weekly\", \"throttle\" : \"100\", \"identity_preservation\":\"exclude_network_and_protocol_config\"}'    ```    <br/> ### Learn more * [`DOC /snapmirror/policies`](#docs-snapmirror-snapmirror_policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Policy UUID | 
**info** | [**SnapmirrorPolicy**](SnapmirrorPolicy.md)| Information on the SnapMirror policy | 

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


## SnapmirrorRelationshipCreate

> JobLinkResponse SnapmirrorRelationshipCreate(ctx, info, optional)



Creates a SnapMirror relationship. This API must be executed on the cluster containing the destination endpoint. ### Required properties * `source.path` - Path to the source endpoint of the SnapMirror relationship. * `destination.path` - Path to the destination endpoint of the SnapMirror relationship. ### Recommended optional properties * `policy.name` or `policy.uuid` - Policy governing the SnapMirror relationship. ### Default property values If not specified in POST, the following default property values are assigned: * `policy.name` - _MirrorAndVault_ * `restore` - _false_ ### Related ONTAP commands * `snapmirror create` ### Examples The following examples show how to create FlexVol, FlexGroup and SVM SnapMirror relationships. Note that the source SVM name should be the local name of the peer SVM.</br>    Creating a FlexVol SnapMirror relationship of type XDP.    <br/>    ```    POST \"/api/snapmirror/relationships/\" '{\"source\": {\"path\": \"test_vserv_src:src_vol_rw\"}, \"destination\": { \"path\": \"test_vserv_dst:dst_vol_rw\"}}'    ```    <br/>    Creating a FlexGroup SnapMirror relationship of type XDP.    <br/>    ```    POST \"/api/snapmirror/relationships/\" '{\"source\": {\"path\": \"test_vserv_src:source_flexgrp\"}, \"destination\": { \"path\": \"test_vserv_dst:dest_flexgrp\"}}'    ```    <br/>    Creating a SVM SnapMirror relationship of type XDP.    <br/>    ```    POST \"/api/snapmirror/relationships/\" '{\"source\": { \"path\": \"src_svm:\"}, \"destination\": { \"path\": \"dst_svm:\"}}'    ```    <br/>    Creating a SnapMirror relationship in order to restore from a destination.    <br/>    ```    POST \"/api/snapmirror/relationships/\" '{\"source\": {\"path\": \"test_vserv_src:src_vol_rw\"}, \"destination\": { \"path\": \"test_vserv_dst:dst_vol_rw\"}, \"restore\": \"true\"}'    ```    <br/> ### Learn more * [`DOC /snapmirror/relationships`](#docs-snapmirror-snapmirror_relationships) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**SnapmirrorRelationship**](SnapmirrorRelationship.md)| Information on the SnapMirror relationship | 
 **optional** | ***SnapmirrorRelationshipCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]

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


## SnapmirrorRelationshipDelete

> JobLinkResponse SnapmirrorRelationshipDelete(ctx, uuid, optional)



Deletes a SnapMirror relationship. ### Important notes * The \"destination_only\", \"source_only\", and \"source_info_only\" flags are mutually exclusive. If no flag is specified, the relationship is deleted from both the source and destination and all common Snapshot copies between the source and destination are also deleted. * For a restore relationship, the call must be executed on the cluster containing the destination endpoint without specifying the destination_only, source_only, or source_info_only parameters. * Additionally, ensure that there are no ongoing transfers on a restore relationship before calling this API. ### Related ONTAP commands * `snapmirror delete` * `snapmirror release` ### Examples The following examples show how to delete the relationship from both the source and destination, the destination only, and the source only. <br/>    Deleting the relationship from both the source and destination. This API must be run on the cluster containing the destination endpoint.    <br/>    ```    DELETE \"/api/snapmirror/relationships/4512b2d2-fd60-11e8-8929-005056bbfe52\"    ```    <br/>    Deleting the relationship on the destination only. This API must be run on the cluster containing the destination endpoint.    <br/>    ```    DELETE \"/api/snapmirror/relationships/fd1e0697-02ba-11e9-acc7-005056a7697f/?destination_only=true\"    ```    <br/>    Deleting the relationship on the source only. This API must be run on the cluster containing the source endpoint.    <br/>    ```    DELETE \"/api/snapmirror/relationships/93e828ba-02bc-11e9-acc7-005056a7697f/?source_only=true\"    ```    <br/>    Deleting the source information only. This API must be run on the cluster containing the source endpoint. This does not delete the common Snapshot copies between the source and destination.    <br/>    ```    DELETE \"/api/snapmirror/relationships/caf545a2-fc60-11e8-aa13-005056a707ff/?source_info_only=true\"    ```    <br/> ### Learn more * [`DOC /snapmirror/relationships`](#docs-snapmirror-snapmirror_relationships) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Relationship UUID | 
 **optional** | ***SnapmirrorRelationshipDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationOnly** | **optional.Bool**| Deletes a relationship on the destination only. This parameter is applicable only when the call is executed on the cluster that contains the destination endpoint. | 
 **sourceOnly** | **optional.Bool**| Deletes a relationship on the source only. This parameter is applicable only when the call is executed on the cluster that contains the source endpoint. | 
 **sourceInfoOnly** | **optional.Bool**| Deletes relationship information on the source only. This parameter is applicable only when the call is executed on the cluster that contains the source endpoint. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]

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


## SnapmirrorRelationshipGet

> SnapmirrorRelationship SnapmirrorRelationshipGet(ctx, uuid, optional)



Retrieves a SnapMirror relationship. ### Related ONTAP commands * `snapmirror show` * `snapmirror list-destinations` ### Example <br/> ``` GET \"/api/snapmirror/relationships/caf545a2-fc60-11e8-aa13-005056a707ff/\" ``` <br/> ### Learn more * [`DOC /snapmirror/relationships`](#docs-snapmirror-snapmirror_relationships) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Relationship UUID | 
 **optional** | ***SnapmirrorRelationshipGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listDestinationsOnly** | **optional.Bool**| Set to true to show relationships from the source only. | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SnapmirrorRelationship**](snapmirror_relationship.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapmirrorRelationshipModify

> JobLinkResponse SnapmirrorRelationshipModify(ctx, uuid, info)



Updates a SnapMirror relationship. This API is used to initiate SnapMirror operations such as \"initialize\", \"resync\", \"break\", \"quiesce\", and \"resume\" by specifying the appropriate value for the \"state\" field. It is also used to modify the SnapMirror policy associated with the specified relationship. ### Related ONTAP commands * `snapmirror modify` * `snapmirror initialize` * `snapmirror resync` * `snapmirror break` * `snapmirror quiesce` * `snapmirror resume` ### Examples The following examples show how to perform the SnapMirror \"resync\", \"initialize\", \"resume\", \"quiesce\", and \"break\" operations. <br/>    Performing a SnapMirror \"resync\"    <br/>    ```    PATCH \"/api/snapmirror/relationships/98bb2608-fc60-11e8-aa13-005056a707ff/\" '{\"state\":\"snapmirrored\"}'    ```    <br/>    Performing a SnapMirror \"initialize\"    <br/>    ```    PATCH \"/api/snapmirror/relationships/98bb2608-fc60-11e8-aa13-005056a707ff/\" '{\"state\":\"snapmirrored\"}'    ```    <br/>    Performing a SnapMirror \"resume\"    <br/>    ```    PATCH \"/api/snapmirror/relationships/98bb2608-fc60-11e8-aa13-005056a707ff/\" '{\"state\":\"snapmirrored\"}'    ```    <br/>    Performing a SnapMirror \"quiesce\"    <br/>    ```    PATCH \"/api/snapmirror/relationships/98bb2608-fc60-11e8-aa13-005056a707ff\" '{\"state\":\"paused\"}'    ```    <br/>    Performing a SnapMirror \"break\"    <br/>    ```    PATCH \"/api/snapmirror/relationships/98bb2608-fc60-11e8-aa13-005056a707ff\" '{\"state\":\"broken_off\"}'    ```    <br/>    Updating an associated SnapMirror policy    <br/>    ```    PATCH \"/api/snapmirror/relationships/9e922e65-1818-11e9-8b22-005056bbee73/\" '{\"policy\": { \"name\" : \"MirrorAndVaultDiscardNetwork\"}}'    ```    <br/> ### Learn more * [`DOC /snapmirror/relationships`](#docs-snapmirror-snapmirror_relationships) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Relationship UUID | 
**info** | [**SnapmirrorRelationship**](SnapmirrorRelationship.md)| Information on the SnapMirror relationship | 

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


## SnapmirrorRelationshipTransferCreate

> SnapmirrorRelationshipTransferCreate(ctx, relationshipUuid, optional)



Starts a SnapMirror transfer operation. This API initiates a restore operation if the SnapMirror relationship is of type \"restore\". Otherwise, it intiates a SnapMirror \"initialize\" operation or \"update\" operation based on the current SnapMirror state. ### Default property values * `storage_efficiency_enabled` - _true_ ### Related ONTAP commands * `snapmirror update` * `snapmirror initialize` * `snapmirror restore` ### Examples The following examples show how to perform SnapMirror \"initialize\", \"update\", and \"restore\" operations. <br/>    Performing a SnapMirror initialize or update    <br/>    ```    POST \"/api/snapmirror/relationships/e4e7e130-0279-11e9-b566-0050568e9909/transfers\" '{}'    ```    <br/>    Performing a SnapMirror restore transfer    <br/>    ```    POST \"/api/snapmirror/relationships/c8c62a90-0fef-11e9-b09e-0050568e7067/transfers\" '{\"source-snapshot\": \"src\", \"files\": {\"source_path\": [\"/a1.txt.0\"], \"destination_path\": [\"/a1-renamed.txt.0\"]}}'    ```    <br/> ### Learn more * [`DOC /snapmirror/relationships/{relationship.uuid}/transfers`](#docs-snapmirror-snapmirror_relationships_{relationship.uuid}_transfers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipUuid** | **string**| Relationship UUID | 
 **optional** | ***SnapmirrorRelationshipTransferCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipTransferCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]
 **info** | [**optional.Interface of SnapmirrorTransfer**](SnapmirrorTransfer.md)| Information on the SnapMirror transfer | 

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


## SnapmirrorRelationshipTransferGet

> SnapmirrorTransfer SnapmirrorRelationshipTransferGet(ctx, relationshipUuid, uuid, optional)



Retrieves the attributes of a specific ongoing SnapMirror transfer. ### Related ONTAP commands * `snapmirror show` ### Example <br/> ``` GET \"/api/snapmirror/relationships/293baa53-e63d-11e8-bff1-005056a793dd/transfers/293baa53-e63d-11e8-bff1-005056a793dd\" ``` <br/> ### Learn more * [`DOC /snapmirror/relationships/{relationship.uuid}/transfers`](#docs-snapmirror-snapmirror_relationships_{relationship.uuid}_transfers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipUuid** | **string**| Relationship UUID | 
**uuid** | **string**| Transfer UUID | 
 **optional** | ***SnapmirrorRelationshipTransferGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipTransferGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SnapmirrorTransfer**](snapmirror_transfer.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapmirrorRelationshipTransferModify

> SnapmirrorRelationshipTransferModify(ctx, relationshipUuid, uuid, info)



Aborts an ongoing SnapMirror transfer. ### Related ONTAP commands * `snapmirror abort` ### Example <br/> ``` PATCH \"/api/snapmirror/relationships/293baa53-e63d-11e8-bff1-005056a793dd/transfers/293baa53-e63d-11e8-bff1-005056a793dd\" '{\"state\":\"aborted\"}' ``` <br/> ### Learn more * [`DOC /snapmirror/relationships/{relationship.uuid}/transfers`](#docs-snapmirror-snapmirror_relationships_{relationship.uuid}_transfers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipUuid** | **string**| Relationship UUID | 
**uuid** | **string**| Transfer UUID | 
**info** | [**SnapmirrorTransfer**](SnapmirrorTransfer.md)| Information on the SnapMirror transfer | 

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


## SnapmirrorRelationshipTransfersGet

> SnapmirrorTransferResponse SnapmirrorRelationshipTransfersGet(ctx, relationshipUuid, optional)



Retrieves the list of ongoing SnapMirror transfers for the specified relationship. ### Related ONTAP commands * `snapmirror show` ### Example <br/> ``` GET \"/api/snapmirror/relationships/293baa53-e63d-11e8-bff1-005056a793dd/transfers\" ``` ### Learn more * [`DOC /snapmirror/relationships/{relationship.uuid}/transfers`](#docs-snapmirror-snapmirror_relationships_{relationship.uuid}_transfers) <br/> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipUuid** | **string**| Relationship UUID | 
 **optional** | ***SnapmirrorRelationshipTransfersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipTransfersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relationshipRestore** | **optional.Bool**| Filter by relationship.restore | 
 **relationshipDestinationPath** | **optional.String**| Filter by relationship.destination.path | 
 **relationshipDestinationSvmUuid** | **optional.String**| Filter by relationship.destination.svm.uuid | 
 **relationshipDestinationSvmName** | **optional.String**| Filter by relationship.destination.svm.name | 
 **relationshipUuid2** | **optional.String**| Filter by relationship.uuid | 
 **uuid** | **optional.String**| Filter by uuid | 
 **checkpointSize** | **optional.Int32**| Filter by checkpoint_size | 
 **snapshot** | **optional.String**| Filter by snapshot | 
 **bytesTransferred** | **optional.Int32**| Filter by bytes_transferred | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SnapmirrorTransferResponse**](snapmirror_transfer_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapmirrorRelationshipsGet

> SnapmirrorRelationshipResponse SnapmirrorRelationshipsGet(ctx, optional)



Retrieves information for SnapMirror relationships whose destination endpoints are in the current SVM or the current cluster, depending on the cluster context. ### Related ONTAP commands * `snapmirror show` * `snapmirror list-destinations` ### Examples The following examples show how to retrieve the list of SnapMirror relationships and the list of SnapMirror destinations.    1. Retrieving the list of SnapMirror relationships. This API must be run on the cluster containing the destination endpoint.    <br/>    ```    GET \"/api/snapmirror/relationships/\"    ```    <br/>   2.  Retrieving the list of SnapMirror destinations on source. This must be run on the cluster containing the source endpoint.    <br/>    ```    GET \"/api/snapmirror/relationships/?list_destinations_only=true\"    ```    <br/> ### Learn more * [`DOC /snapmirror/relationships`](#docs-snapmirror-snapmirror_relationships) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapmirrorRelationshipsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapmirrorRelationshipsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listDestinationsOnly** | **optional.Bool**| Set to true to show relationships from the source only. | 
 **restore** | **optional.Bool**| Filter by restore | 
 **transferState** | **optional.String**| Filter by transfer.state | 
 **transferUuid** | **optional.String**| Filter by transfer.uuid | 
 **transferBytesTransferred** | **optional.Int32**| Filter by transfer.bytes_transferred | 
 **uuid** | **optional.String**| Filter by uuid | 
 **policyUuid** | **optional.String**| Filter by policy.uuid | 
 **policyType** | **optional.String**| Filter by policy.type | 
 **policyName** | **optional.String**| Filter by policy.name | 
 **exportedSnapshot** | **optional.String**| Filter by exported_snapshot | 
 **sourcePath** | **optional.String**| Filter by source.path | 
 **sourceSvmUuid** | **optional.String**| Filter by source.svm.uuid | 
 **sourceSvmName** | **optional.String**| Filter by source.svm.name | 
 **healthy** | **optional.Bool**| Filter by healthy | 
 **lagTime** | **optional.String**| Filter by lag_time | 
 **destinationPath** | **optional.String**| Filter by destination.path | 
 **destinationSvmUuid** | **optional.String**| Filter by destination.svm.uuid | 
 **destinationSvmName** | **optional.String**| Filter by destination.svm.name | 
 **unhealthyReasonParameters** | **optional.String**| Filter by unhealthy_reason.parameters | 
 **unhealthyReasonMessage** | **optional.String**| Filter by unhealthy_reason.message | 
 **unhealthyReasonCode** | **optional.Int32**| Filter by unhealthy_reason.code | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SnapmirrorRelationshipResponse**](snapmirror_relationship_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

