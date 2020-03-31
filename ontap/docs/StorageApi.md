# \StorageApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateCollectionGet**](StorageApi.md#AggregateCollectionGet) | **Get** /storage/aggregates | 
[**AggregateCreate**](StorageApi.md#AggregateCreate) | **Post** /storage/aggregates | 
[**AggregateDelete**](StorageApi.md#AggregateDelete) | **Delete** /storage/aggregates/{uuid} | 
[**AggregateGet**](StorageApi.md#AggregateGet) | **Get** /storage/aggregates/{uuid} | 
[**AggregateModify**](StorageApi.md#AggregateModify) | **Patch** /storage/aggregates/{uuid} | 
[**CloudStoreCollectionGet**](StorageApi.md#CloudStoreCollectionGet) | **Get** /storage/aggregates/{aggregate.uuid}/cloud-stores | 
[**CloudStoreCreate**](StorageApi.md#CloudStoreCreate) | **Post** /storage/aggregates/{aggregate.uuid}/cloud-stores | 
[**CloudStoreDelete**](StorageApi.md#CloudStoreDelete) | **Delete** /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid} | 
[**CloudStoreGet**](StorageApi.md#CloudStoreGet) | **Get** /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid} | 
[**CloudStoreModify**](StorageApi.md#CloudStoreModify) | **Patch** /storage/aggregates/{aggregate.uuid}/cloud-stores/{target.uuid} | 
[**DiskCollectionGet**](StorageApi.md#DiskCollectionGet) | **Get** /storage/disks | 
[**DiskGet**](StorageApi.md#DiskGet) | **Get** /storage/disks/{name} | 
[**FileCloneCreate**](StorageApi.md#FileCloneCreate) | **Post** /storage/file/clone | 
[**FlexcacheCollectionGet**](StorageApi.md#FlexcacheCollectionGet) | **Get** /storage/flexcache/flexcaches | 
[**FlexcacheCreate**](StorageApi.md#FlexcacheCreate) | **Post** /storage/flexcache/flexcaches | 
[**FlexcacheDelete**](StorageApi.md#FlexcacheDelete) | **Delete** /storage/flexcache/flexcaches/{uuid} | 
[**FlexcacheGet**](StorageApi.md#FlexcacheGet) | **Get** /storage/flexcache/flexcaches/{uuid} | 
[**FlexcacheOriginCollectionGet**](StorageApi.md#FlexcacheOriginCollectionGet) | **Get** /storage/flexcache/origins | 
[**FlexcacheOriginGet**](StorageApi.md#FlexcacheOriginGet) | **Get** /storage/flexcache/origins/{uuid} | 
[**PlexCollectionGet**](StorageApi.md#PlexCollectionGet) | **Get** /storage/aggregates/{aggregate.uuid}/plexes | 
[**PlexGet**](StorageApi.md#PlexGet) | **Get** /storage/aggregates/{aggregate.uuid}/plexes/{name} | 
[**PortCollectionGet**](StorageApi.md#PortCollectionGet) | **Get** /storage/ports | 
[**PortGet**](StorageApi.md#PortGet) | **Get** /storage/ports/{node.uuid}/{name} | 
[**QosPolicyCollectionGet**](StorageApi.md#QosPolicyCollectionGet) | **Get** /storage/qos/policies | 
[**QosPolicyCreate**](StorageApi.md#QosPolicyCreate) | **Post** /storage/qos/policies | 
[**QosPolicyDelete**](StorageApi.md#QosPolicyDelete) | **Delete** /storage/qos/policies/{policy.uuid} | 
[**QosPolicyGet**](StorageApi.md#QosPolicyGet) | **Get** /storage/qos/policies/{policy.uuid} | 
[**QosPolicyModify**](StorageApi.md#QosPolicyModify) | **Patch** /storage/qos/policies/{policy.uuid} | 
[**QtreeCollectionGet**](StorageApi.md#QtreeCollectionGet) | **Get** /storage/qtrees | 
[**QtreeCreate**](StorageApi.md#QtreeCreate) | **Post** /storage/qtrees | 
[**QtreeDelete**](StorageApi.md#QtreeDelete) | **Delete** /storage/qtrees/{volume.uuid}/{id} | 
[**QtreeGet**](StorageApi.md#QtreeGet) | **Get** /storage/qtrees/{volume.uuid}/{id} | 
[**QtreeModify**](StorageApi.md#QtreeModify) | **Patch** /storage/qtrees/{volume.uuid}/{id} | 
[**QuotaReportCollectionGet**](StorageApi.md#QuotaReportCollectionGet) | **Get** /storage/quota/reports | 
[**QuotaReportGet**](StorageApi.md#QuotaReportGet) | **Get** /storage/quota/reports/{volume.uuid}/{index} | 
[**QuotaRuleCollectionGet**](StorageApi.md#QuotaRuleCollectionGet) | **Get** /storage/quota/rules | 
[**QuotaRuleCreate**](StorageApi.md#QuotaRuleCreate) | **Post** /storage/quota/rules | 
[**QuotaRuleDelete**](StorageApi.md#QuotaRuleDelete) | **Delete** /storage/quota/rules/{rule.uuid} | 
[**QuotaRuleGet**](StorageApi.md#QuotaRuleGet) | **Get** /storage/quota/rules/{rule.uuid} | 
[**QuotaRuleModify**](StorageApi.md#QuotaRuleModify) | **Patch** /storage/quota/rules/{rule.uuid} | 
[**ShelfCollectionGet**](StorageApi.md#ShelfCollectionGet) | **Get** /storage/shelves | 
[**ShelfGet**](StorageApi.md#ShelfGet) | **Get** /storage/shelves/{uid} | 
[**SnapshotCollectionGet**](StorageApi.md#SnapshotCollectionGet) | **Get** /storage/volumes/{volume.uuid}/snapshots | 
[**SnapshotCreate**](StorageApi.md#SnapshotCreate) | **Post** /storage/volumes/{volume.uuid}/snapshots | 
[**SnapshotDelete**](StorageApi.md#SnapshotDelete) | **Delete** /storage/volumes/{volume.uuid}/snapshots/{uuid} | 
[**SnapshotGet**](StorageApi.md#SnapshotGet) | **Get** /storage/volumes/{volume.uuid}/snapshots/{uuid} | 
[**SnapshotModify**](StorageApi.md#SnapshotModify) | **Patch** /storage/volumes/{volume.uuid}/snapshots/{uuid} | 
[**SnapshotPolicyCollectionGet**](StorageApi.md#SnapshotPolicyCollectionGet) | **Get** /storage/snapshot-policies | 
[**SnapshotPolicyCreate**](StorageApi.md#SnapshotPolicyCreate) | **Post** /storage/snapshot-policies | 
[**SnapshotPolicyDelete**](StorageApi.md#SnapshotPolicyDelete) | **Delete** /storage/snapshot-policies/{uuid} | 
[**SnapshotPolicyGet**](StorageApi.md#SnapshotPolicyGet) | **Get** /storage/snapshot-policies/{uuid} | 
[**SnapshotPolicyModify**](StorageApi.md#SnapshotPolicyModify) | **Patch** /storage/snapshot-policies/{uuid} | 
[**StorageClusterGet**](StorageApi.md#StorageClusterGet) | **Get** /storage/cluster | 
[**VolumeCollectionGet**](StorageApi.md#VolumeCollectionGet) | **Get** /storage/volumes | 
[**VolumeCollectionPerformanceMetricsGet**](StorageApi.md#VolumeCollectionPerformanceMetricsGet) | **Get** /storage/volumes/{uuid}/metrics | 
[**VolumeCreate**](StorageApi.md#VolumeCreate) | **Post** /storage/volumes | 
[**VolumeDelete**](StorageApi.md#VolumeDelete) | **Delete** /storage/volumes/{uuid} | 
[**VolumeGet**](StorageApi.md#VolumeGet) | **Get** /storage/volumes/{uuid} | 
[**VolumeModify**](StorageApi.md#VolumeModify) | **Patch** /storage/volumes/{uuid} | 



## AggregateCollectionGet

> AggregateResponse AggregateCollectionGet(ctx, optional)



Retrieves the collection of aggregates for the entire cluster. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `space.block_storage.inactive_user_data` * `space.footprint` ### Related ONTAP commands * `storage aggregate show`  ### Learn more * [`DOC /storage/aggregates`](#docs-storage-storage_aggregates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AggregateCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AggregateCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommend** | **optional.Bool**| If set to &#39;true&#39;, it queries the system for the recommended optimal layout for creating new aggregates. The default setting is &#39;false&#39;. | 
 **showSpares** | **optional.Bool**| If set to &#39;true&#39;, the spares object is returned instead of records to show the spare groups in the cluster. The default setting is &#39;false&#39;. | 
 **snaplockType** | **optional.String**| Filter by snaplock_type | 
 **name** | **optional.String**| Filter by name | 
 **homeNodeUuid** | **optional.String**| Filter by home_node.uuid | 
 **homeNodeName** | **optional.String**| Filter by home_node.name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **drHomeNodeName** | **optional.String**| Filter by dr_home_node.name | 
 **drHomeNodeUuid** | **optional.String**| Filter by dr_home_node.uuid | 
 **createTime** | **optional.String**| Filter by create_time | 
 **spaceEfficiencyLogicalUsed** | **optional.Int32**| Filter by space.efficiency.logical_used | 
 **spaceEfficiencyRatio** | **optional.Float32**| Filter by space.efficiency.ratio | 
 **spaceEfficiencySavings** | **optional.Int32**| Filter by space.efficiency.savings | 
 **spaceFootprint** | **optional.Int32**| Filter by space.footprint | 
 **spaceEfficiencyWithoutSnapshotsLogicalUsed** | **optional.Int32**| Filter by space.efficiency_without_snapshots.logical_used | 
 **spaceEfficiencyWithoutSnapshotsRatio** | **optional.Float32**| Filter by space.efficiency_without_snapshots.ratio | 
 **spaceEfficiencyWithoutSnapshotsSavings** | **optional.Int32**| Filter by space.efficiency_without_snapshots.savings | 
 **spaceBlockStorageFullThresholdPercent** | **optional.Int32**| Filter by space.block_storage.full_threshold_percent | 
 **spaceBlockStorageUsed** | **optional.Int32**| Filter by space.block_storage.used | 
 **spaceBlockStorageAvailable** | **optional.Int32**| Filter by space.block_storage.available | 
 **spaceBlockStorageSize** | **optional.Int32**| Filter by space.block_storage.size | 
 **spaceBlockStorageInactiveUserData** | **optional.Int32**| Filter by space.block_storage.inactive_user_data | 
 **spaceCloudStorageUsed** | **optional.Int32**| Filter by space.cloud_storage.used | 
 **blockStoragePrimaryRaidSize** | **optional.Int32**| Filter by block_storage.primary.raid_size | 
 **blockStoragePrimaryRaidType** | **optional.String**| Filter by block_storage.primary.raid_type | 
 **blockStoragePrimaryChecksumStyle** | **optional.String**| Filter by block_storage.primary.checksum_style | 
 **blockStoragePrimaryDiskClass** | **optional.String**| Filter by block_storage.primary.disk_class | 
 **blockStoragePrimaryDiskCount** | **optional.Int32**| Filter by block_storage.primary.disk_count | 
 **blockStoragePlexesName** | **optional.String**| Filter by block_storage.plexes.name | 
 **blockStorageHybridCacheRaidType** | **optional.String**| Filter by block_storage.hybrid_cache.raid_type | 
 **blockStorageHybridCacheUsed** | **optional.Int32**| Filter by block_storage.hybrid_cache.used | 
 **blockStorageHybridCacheEnabled** | **optional.Bool**| Filter by block_storage.hybrid_cache.enabled | 
 **blockStorageHybridCacheSize** | **optional.Int32**| Filter by block_storage.hybrid_cache.size | 
 **blockStorageHybridCacheDiskCount** | **optional.Int32**| Filter by block_storage.hybrid_cache.disk_count | 
 **blockStorageMirrorState** | **optional.String**| Filter by block_storage.mirror.state | 
 **blockStorageMirrorEnabled** | **optional.Bool**| Filter by block_storage.mirror.enabled | 
 **dataEncryptionSoftwareEncryptionEnabled** | **optional.Bool**| Filter by data_encryption.software_encryption_enabled | 
 **dataEncryptionDriveProtectionEnabled** | **optional.Bool**| Filter by data_encryption.drive_protection_enabled | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**AggregateResponse**](aggregate_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateCreate

> JobLinkResponse AggregateCreate(ctx, optional)



Automatically creates aggregates based on an optimal layout recommended by the system. Alternatively, properties can be provided to create an aggregate according to the requested specification. This request starts a job and returns a link to that job. ### Required properties Properties are not required for this API. The following properties are only required if you want to specify properties for aggregate creation: * `name` - Name of the aggregate. * `node.name` or `node.uuid` - Node on which the aggregate will be created. * `block_storage.primary.disk_count` - Number of disks to be used to create the aggregate. ### Default values If not specified in POST, the following default values are assigned. The remaining unspecified properties will receive system dependent default values. * `block_storage.mirror.enabled` - _false_ * `snaplock_type` - _non_snaplock_ ### Related ONTAP commands * `storage aggregate auto-provision` * `storage aggregate create` ### Example: ``` POST /api/storage/aggregates {\"node\": {\"name\": \"node1\"}, \"name\": \"test\", \"block_storage\": {\"primary\": {\"disk_count\": \"10\"}}} ```  ### Learn more * [`DOC /storage/aggregates`](#docs-storage-storage_aggregates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AggregateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AggregateCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskSize** | **optional.Int32**| If set, POST only selects disks of the specified size. | 
 **info** | [**optional.Interface of Aggregate**](Aggregate.md)| Info specification | 

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


## AggregateDelete

> JobLinkResponse AggregateDelete(ctx, uuid)



Deletes the aggregate specified by the UUID. This request starts a job and returns a link to that job. ### Related ONTAP commands * `storage aggregate delete`  ### Learn more * [`DOC /storage/aggregates`](#docs-storage-storage_aggregates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Aggregate UUID | 

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


## AggregateGet

> Aggregate AggregateGet(ctx, uuid, optional)



Retrieves the aggregate specified by the UUID. The recommend query cannot be used for this operation. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `space.block_storage.inactive_user_data` * `space.footprint` ### Related ONTAP commands * `storage aggregate show`  ### Learn more * [`DOC /storage/aggregates`](#docs-storage-storage_aggregates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Aggregate UUID | 
 **optional** | ***AggregateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AggregateGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Aggregate**](aggregate.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateModify

> JobLinkResponse AggregateModify(ctx, uuid, optional)



Updates the aggregate specified by the UUID with the properties in the body. This request starts a job and returns a link to that job. ### Related ONTAP commands * `storage aggregate add-disks` * `storage aggregate mirror` * `storage aggregate modify` * `storage aggregate relocation start` * `storage aggregate rename`  ### Learn more * [`DOC /storage/aggregates`](#docs-storage-storage_aggregates)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Aggregate UUID | 
 **optional** | ***AggregateModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AggregateModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskSize** | **optional.Int32**| If set, PATCH only selects disks of the specified size. | 
 **info** | [**optional.Interface of Aggregate**](Aggregate.md)| Info specification | 

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


## CloudStoreCollectionGet

> CloudStoreResponse CloudStoreCollectionGet(ctx, aggregateUuid, optional)



Retrieves the collection of cloud stores used by an aggregate. ### Related ONTAP commands * `storage aggregate object-store show` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
 **optional** | ***CloudStoreCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudStoreCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorDegraded** | **optional.Bool**| Filter by mirror_degraded | 
 **targetName** | **optional.String**| Filter by target.name | 
 **targetUuid** | **optional.String**| Filter by target.uuid | 
 **used** | **optional.Int32**| Filter by used | 
 **availability** | **optional.String**| Filter by availability | 
 **unreclaimedSpaceThreshold** | **optional.Int32**| Filter by unreclaimed_space_threshold | 
 **primary** | **optional.Bool**| Filter by primary | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**CloudStoreResponse**](cloud_store_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudStoreCreate

> JobLinkResponse CloudStoreCreate(ctx, aggregateUuid, optional)



Attaches an object store to an aggregate, or adds a second object store as a mirror. ### Required properties * `target.uuid` or `target.name` - UUID or name of the cloud target. ### Recommended optional properties * `primary` - _true_ if the object store is primary or _false_ if it is a mirror. * `allow_flexgroups` - Allow attaching object store to an aggregate containing FlexGroup constituents. * `check_only` - Validate only and do not add the cloud store. ### Default property values * `primary` - _true_ * `allow_flexgroups` - _false_ * `check_only` - _false_ ### Related ONTAP commands * `storage aggregate object-store attach` * `storage aggregate object-store mirror` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
 **optional** | ***CloudStoreCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudStoreCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowFlexgroups** | **optional.Bool**| This optional parameter allows attaching object store to an aggregate containing FlexGroup constituents. The default value is false. Mixing FabricPools and non-FabricPools within a FlexGroup is not recommended. All aggregates hosting constituents of a FlexGroup should be attached to the object store. | 
 **checkOnly** | **optional.Bool**| Validate only and do not add the cloud store. | 
 **info** | [**optional.Interface of CloudStore**](CloudStore.md)| Info specification | 

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


## CloudStoreDelete

> JobLinkResponse CloudStoreDelete(ctx, aggregateUuid, targetUuid)



Removes the specified cloud target from the aggregate. Only removal of a mirror is allowed. The primary cannot be removed. This request starts a job and returns a link to that job. ### Related ONTAP commands * `storage aggregate object-store unmirror` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
**targetUuid** | **string**| Cloud target UUID | 

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


## CloudStoreGet

> CloudStoreResponse CloudStoreGet(ctx, aggregateUuid, targetUuid, optional)



Retrieves the cloud store for the aggregate using the specified cloud target UUID. ### Related ONTAP commands * `storage aggregate object-store show` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
**targetUuid** | **string**| Cloud target UUID | 
 **optional** | ***CloudStoreGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudStoreGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**CloudStoreResponse**](cloud_store_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudStoreModify

> JobLinkResponse CloudStoreModify(ctx, aggregateUuid, targetUuid, optional)



Updates the cloud store specified by the UUID with the fields in the body. This request starts a job and returns a link to that job. ### Related ONTAP commands * `storage aggregate object-store modify` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
**targetUuid** | **string**| Cloud target UUID | 
 **optional** | ***CloudStoreModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudStoreModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of CloudStore**](CloudStore.md)| Info specification | 

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


## DiskCollectionGet

> DiskResponse DiskCollectionGet(ctx, optional)



Retrieves a collection of disks. ### Related ONTAP commands * `storage disk show` ### Learn more * [`DOC /storage/disks`](#docs-storage-storage_disks) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiskCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiskCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpm** | **optional.Int32**| Filter by rpm | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **containerType** | **optional.String**| Filter by container_type | 
 **uid** | **optional.String**| Filter by uid | 
 **vendor** | **optional.String**| Filter by vendor | 
 **firmwareVersion** | **optional.String**| Filter by firmware_version | 
 **shelfUid** | **optional.String**| Filter by shelf.uid | 
 **pool** | **optional.String**| Filter by pool | 
 **state** | **optional.String**| Filter by state | 
 **class** | **optional.String**| Filter by class | 
 **aggregatesUuid** | **optional.String**| Filter by aggregates.uuid | 
 **aggregatesName** | **optional.String**| Filter by aggregates.name | 
 **bay** | **optional.String**| Filter by bay | 
 **drawerId** | **optional.Int32**| Filter by drawer.id | 
 **drawerSlot** | **optional.Int32**| Filter by drawer.slot | 
 **usableSize** | **optional.Int32**| Filter by usable_size | 
 **model** | **optional.String**| Filter by model | 
 **type_** | **optional.String**| Filter by type | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **homeNodeUuid** | **optional.String**| Filter by home_node.uuid | 
 **homeNodeName** | **optional.String**| Filter by home_node.name | 
 **drNodeName** | **optional.String**| Filter by dr_node.name | 
 **drNodeUuid** | **optional.String**| Filter by dr_node.uuid | 
 **name** | **optional.String**| Filter by name | 
 **ratedLifeUsedPercent** | **optional.Int32**| Filter by rated_life_used_percent | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**DiskResponse**](disk_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiskGet

> Disk DiskGet(ctx, name, optional)



Retrieves a specific disk. ### Related ONTAP commands * `storage disk show` ### Learn more * [`DOC /storage/disks`](#docs-storage-storage_disks) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Disk name | 
 **optional** | ***DiskGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiskGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Disk**](disk.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileCloneCreate

> JobLinkResponse FileCloneCreate(ctx, optional)



Creates a clone of the file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileCloneCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FileCloneCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of FileClone**](FileClone.md)| Info specification | 

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


## FlexcacheCollectionGet

> FlexcacheResponse FlexcacheCollectionGet(ctx, optional)



Retrieves FlexCaches in the cluster. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `origins.ip_address` - IP address of origin. * `origins.size` - Physical size of origin. * `origins.state` - State of origin. * `size` - Physical size of FlexCache. * `aggregates.name` or `aggregates.uuid` - Name or UUID of aggregrate of FlexCache volume. * `path` - Fully-qualified path of the owning SVM's namespace where the FlexCache is mounted. ### Related ONTAP commands * `volume flexcache show` ### Learn more * [`DOC /storage/flexcache/flexcaches`](#docs-storage-storage_flexcache_flexcaches) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlexcacheCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlexcacheCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **size** | **optional.Int32**| Filter by size | 
 **path** | **optional.String**| Filter by path | 
 **uuid** | **optional.String**| Filter by uuid | 
 **constituentsPerAggregate** | **optional.Int32**| Filter by constituents_per_aggregate | 
 **aggregatesUuid** | **optional.String**| Filter by aggregates.uuid | 
 **aggregatesName** | **optional.String**| Filter by aggregates.name | 
 **originsIpAddress** | **optional.String**| Filter by origins.ip_address | 
 **originsCreateTime** | **optional.String**| Filter by origins.create_time | 
 **originsSvmUuid** | **optional.String**| Filter by origins.svm.uuid | 
 **originsSvmName** | **optional.String**| Filter by origins.svm.name | 
 **originsClusterUuid** | **optional.String**| Filter by origins.cluster.uuid | 
 **originsClusterName** | **optional.String**| Filter by origins.cluster.name | 
 **originsVolumeName** | **optional.String**| Filter by origins.volume.name | 
 **originsVolumeUuid** | **optional.String**| Filter by origins.volume.uuid | 
 **originsState** | **optional.String**| Filter by origins.state | 
 **originsSize** | **optional.Int32**| Filter by origins.size | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]

### Return type

[**FlexcacheResponse**](flexcache_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlexcacheCreate

> JobLinkResponse FlexcacheCreate(ctx, info)



Creates a FlexCache in the cluster. ### Required properties * `name` - Name of FlexCache volume. * `origins.volume.name` or `origins.volume.uuid` - Name or UUID of origin volume. * `origins.svm.name` - Name of origin Vserver. * `svm.name` or `svm.uuid` - Name or UUID of Vserver where FlexCache will be created. ### Recommended optional properties * `path` - Path to mount the FlexCache volume ### Default property values If not specified in POST, the following default property values are assigned: * `size` - 10% of origin volume size or 1GB per constituent, whichever is greater. * `constituents_per_aggregate` - 4 if aggregates.name or aggregates.uuid is used. ### Related ONTAP commands * `volume flexcache create` ### Learn more * [`DOC /storage/flexcache/flexcaches`](#docs-storage-storage_flexcache_flexcaches) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**Flexcache**](Flexcache.md)| Info specification | 

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


## FlexcacheDelete

> JobLinkResponse FlexcacheDelete(ctx, uuid)



Deletes a FlexCache. If a FlexCache volume is online, it is offlined before deletion. ### Related ONTAP commands * `volume flexcache delete` ### Learn more * [`DOC /storage/flexcache/flexcaches`](#docs-storage-storage_flexcache_flexcaches) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of the FlexCache. | 

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


## FlexcacheGet

> Flexcache FlexcacheGet(ctx, uuid, optional)



Retrieves attributes of the FlexCache in the cluster. ### Expensive properties There is an added cost to retrieving values for these properties. They are included by default in GET. The recommended method to use this API is to use filter and retrieve only the required fields. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `origins.ip_address` - IP address of origin. * `origins.size` - Physical size of origin. * `origins.state` - State of origin. * `size` - Physical size of FlexCache. * `aggregates.name` or `aggregates.uuid` - Name or UUID of aggregrate of FlexCache volume. * `path` - Fully-qualified path of the owning SVM's namespace where the FlexCache is mounted. ### Related ONTAP commands * `volume flexcache show` ### Learn more * [`DOC /storage/flexcache/flexcaches`](#docs-storage-storage_flexcache_flexcaches) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of FlexCache. | 
 **optional** | ***FlexcacheGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlexcacheGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Flexcache**](flexcache.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlexcacheOriginCollectionGet

> FlexcacheOriginResponse FlexcacheOriginCollectionGet(ctx, optional)



Retrieves origin of FlexCaches in the cluster. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `flexcaches.ip_address` - IP address of FlexCache. * `flexcaches.size` - Physical size of FlexCache. * `flexcaches.state` - State of FlexCache. ### Related ONTAP commands * `volume flexcache origin show-caches` ### Learn more * [`DOC /storage/flexcache/origins`](#docs-storage-storage_flexcache_origins) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlexcacheOriginCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlexcacheOriginCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **flexcachesIpAddress** | **optional.String**| Filter by flexcaches.ip_address | 
 **flexcachesCreateTime** | **optional.String**| Filter by flexcaches.create_time | 
 **flexcachesSvmUuid** | **optional.String**| Filter by flexcaches.svm.uuid | 
 **flexcachesSvmName** | **optional.String**| Filter by flexcaches.svm.name | 
 **flexcachesClusterUuid** | **optional.String**| Filter by flexcaches.cluster.uuid | 
 **flexcachesClusterName** | **optional.String**| Filter by flexcaches.cluster.name | 
 **flexcachesVolumeName** | **optional.String**| Filter by flexcaches.volume.name | 
 **flexcachesVolumeUuid** | **optional.String**| Filter by flexcaches.volume.uuid | 
 **flexcachesState** | **optional.String**| Filter by flexcaches.state | 
 **flexcachesSize** | **optional.Int32**| Filter by flexcaches.size | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]

### Return type

[**FlexcacheOriginResponse**](flexcache_origin_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlexcacheOriginGet

> FlexcacheOrigin FlexcacheOriginGet(ctx, uuid, optional)



Retrieves attributes of the origin of a FlexCache in the cluster. ### Expensive properties There is an added cost to retrieving values for these properties. They are included by default in GET results. The recommended method to use this API is to use filter and retrieve only the required fields. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `flexcaches.ip_address` - IP address of FlexCache. * `flexcaches.size` - Physical size of FlexCache. * `flexcaches.state` - State of FlexCache. ### Related ONTAP commands * `volume flexcache origin show-caches` ### Learn more * [`DOC /storage/flexcache/origins`](#docs-storage-storage_flexcache_origins) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of origin of FlexCache. | 
 **optional** | ***FlexcacheOriginGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlexcacheOriginGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**FlexcacheOrigin**](flexcache_origin.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlexCollectionGet

> PlexResponse PlexCollectionGet(ctx, aggregateUuid, name, optional)



Retrieves the collection of plexes for the specified aggregate. ### Related ONTAP commands * `storage aggregate plex show`  ### Learn more * [`DOC /storage/aggregates/{aggregate.uuid}/plexes`](#docs-storage-storage_aggregates_{aggregate.uuid}_plexes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
**name** | **string**| Filter by name | 
 **optional** | ***PlexCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlexCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregateUuid2** | **optional.String**| Filter by aggregate.uuid | 
 **aggregateName** | **optional.String**| Filter by aggregate.name | 
 **resyncPercent** | **optional.Int32**| Filter by resync.percent | 
 **resyncLevel** | **optional.String**| Filter by resync.level | 
 **resyncActive** | **optional.Bool**| Filter by resync.active | 
 **online** | **optional.Bool**| Filter by online | 
 **state** | **optional.String**| Filter by state | 
 **pool** | **optional.String**| Filter by pool | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**PlexResponse**](plex_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlexGet

> Plex PlexGet(ctx, aggregateUuid, name, optional)



Retrieves the plex specified by the aggregate UUID and plex name. ### Related ONTAP commands * `storage aggregate plex show`  ### Learn more * [`DOC /storage/aggregates/{aggregate.uuid}/plexes`](#docs-storage-storage_aggregates_{aggregate.uuid}_plexes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregateUuid** | **string**| Aggregate UUID | 
**name** | **string**| Plex name | 
 **optional** | ***PlexGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlexGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Plex**](plex.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortCollectionGet

> StoragePortResponse PortCollectionGet(ctx, optional)



Retrieves a collection of storage ports. ### Related ONTAP commands * `storage port show` ### Learn more * [`DOC /storage/ports`](#docs-storage-storage_ports) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PortCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **optional.String**| Filter by description | 
 **macAddress** | **optional.String**| Filter by mac_address | 
 **partNumber** | **optional.String**| Filter by part_number | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **nodeUuid** | **optional.String**| Filter by node.uuid | 
 **nodeName** | **optional.String**| Filter by node.name | 
 **errorMessage** | **optional.String**| Filter by error.message | 
 **errorCorrectiveAction** | **optional.String**| Filter by error.corrective_action | 
 **boardName** | **optional.String**| Filter by board_name | 
 **cableSerialNumber** | **optional.String**| Filter by cable.serial_number | 
 **cablePartNumber** | **optional.String**| Filter by cable.part_number | 
 **cableLength** | **optional.String**| Filter by cable.length | 
 **cableIdentifier** | **optional.String**| Filter by cable.identifier | 
 **name** | **optional.String**| Filter by name | 
 **state** | **optional.String**| Filter by state | 
 **speed** | **optional.Float32**| Filter by speed | 
 **wwn** | **optional.String**| Filter by wwn | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**StoragePortResponse**](storage_port_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortGet

> StoragePort PortGet(ctx, nodeUuid, name, optional)



Retrieves a specific storage port. ### Related ONTAP commands * `storage port show` ### Learn more * [`DOC /storage/ports`](#docs-storage-storage_ports) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeUuid** | **string**| Node UUID | 
**name** | **string**| Port name | 
 **optional** | ***PortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PortGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**StoragePort**](storage_port.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QosPolicyCollectionGet

> QosPolicyResponse QosPolicyCollectionGet(ctx, optional)



Retrieves a collection of QoS policies. ### Learn more * [`DOC /storage/qos/policies`](#docs-storage-storage_qos_policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QosPolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QosPolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectCount** | **optional.Int32**| Filter by object_count | 
 **adaptiveExpectedIops** | **optional.Int32**| Filter by adaptive.expected_iops | 
 **adaptiveAbsoluteMinIops** | **optional.Int32**| Filter by adaptive.absolute_min_iops | 
 **adaptivePeakIops** | **optional.Int32**| Filter by adaptive.peak_iops | 
 **name** | **optional.String**| Filter by name | 
 **fixedMaxThroughputIops** | **optional.Int32**| Filter by fixed.max_throughput_iops | 
 **fixedCapacityShared** | **optional.Bool**| Filter by fixed.capacity_shared | 
 **fixedMaxThroughputMbps** | **optional.Int32**| Filter by fixed.max_throughput_mbps | 
 **fixedMinThroughputIops** | **optional.Int32**| Filter by fixed.min_throughput_iops | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**QosPolicyResponse**](qos_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QosPolicyCreate

> JobLinkResponse QosPolicyCreate(ctx, optional)



Creates a QoS policy. ### Required properties * `svm.uuid` or `svm.name` - The existing SVM owning the QoS policy. * `name` - The name of the QoS policy. * `fixed.*` or `adaptive.*` - Either of the fixed or adaptive parameters. ### Default property values * If `fixed.*` parameters are specified, then capacity.shared is set to false by default. ### Related ONTAP commands * `qos policy-group create` * `qos adaptive-policy-group create`  ### Learn more * [`DOC /storage/qos/policies`](#docs-storage-storage_qos_policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QosPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QosPolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]
 **info** | [**optional.Interface of QosPolicy**](QosPolicy.md)| Info specification | 

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


## QosPolicyDelete

> JobLinkResponse QosPolicyDelete(ctx, policyUuid, optional)



Deletes a QoS policy. All QoS workloads associated with the policy are removed. ### Related ONTAP commands * `qos policy-group delete` * `qos adaptive-policy-group delete`  ### Learn more * [`DOC /storage/qos/policies`](#docs-storage-storage_qos_policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string**|  | 
 **optional** | ***QosPolicyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QosPolicyDeleteOpts struct


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


## QosPolicyGet

> QosPolicy QosPolicyGet(ctx, policyUuid, optional)



Retrieves a specific QoS policy. ### Related ONTAP commands * `qos policy-group show` * `qos adaptive-policy-group show`  ### Learn more * [`DOC /storage/qos/policies`](#docs-storage-storage_qos_policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string**|  | 
 **optional** | ***QosPolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QosPolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**QosPolicy**](qos_policy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QosPolicyModify

> JobLinkResponse QosPolicyModify(ctx, policyUuid, optional)



Update a specific QoS policy. ### Related ONTAP commands * `qos policy-group modify` * `qos adaptive-policy-group modify`  ### Learn more * [`DOC /storage/qos/policies`](#docs-storage-storage_qos_policies)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string**|  | 
 **optional** | ***QosPolicyModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QosPolicyModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]
 **info** | [**optional.Interface of QosPolicy**](QosPolicy.md)| Info specification | 

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


## QtreeCollectionGet

> QtreeResponse QtreeCollectionGet(ctx, optional)



Retrieves qtrees configured for all FlexVols or FlexGroup volumes. <br/> Use the `fields` query parameter to retrieve all properties of the qtree. If the `fields` query parameter is not used, then GET returns the qtree `name` and qtree `id` only. ### Related ONTAP commands * `qtree show`  ### Learn more * [`DOC /storage/qtrees`](#docs-storage-storage_qtrees)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QtreeCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QtreeCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **optional.String**| Filter by path | 
 **exportPolicyName** | **optional.String**| Filter by export_policy.name | 
 **exportPolicyId** | **optional.Int32**| Filter by export_policy.id | 
 **volumeName** | **optional.String**| Filter by volume.name | 
 **volumeUuid** | **optional.String**| Filter by volume.uuid | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **unixPermissions** | **optional.Int32**| Filter by unix_permissions | 
 **securityStyle** | **optional.String**| Filter by security_style | 
 **id** | **optional.Int32**| Filter by id | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**QtreeResponse**](qtree_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QtreeCreate

> JobLinkResponse QtreeCreate(ctx, info)



Creates a qtree in a FlexVol or a FlexGroup volume. <br/> After a qtree is created, the new qtree is assigned an identifier. This identifier is obtained using a qtree GET request. This identifier is used in the API path for the qtree PATCH and DELETE operations. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the qtree. * `volume.uuid` or `volume.name` - Existing volume in which to create the qtree. * `name` - Name for the qtree. ### Recommended optional properties If not specified in POST, the values are inherited from the volume. * `security_style` - Security style for the qtree. * `unix_permissions` - UNIX permissions for the qtree. * `export_policy.name or export_policy.id` - Export policy of the SVM for the qtree. ### Related ONTAP commands * `qtree create`  ### Learn more * [`DOC /storage/qtrees`](#docs-storage-storage_qtrees)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**Qtree**](Qtree.md)| Info specification | 

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


## QtreeDelete

> JobLinkResponse QtreeDelete(ctx, volumeUuid, id)



Deletes a qtree. ### Related ONTAP commands * `qtree delete`  ### Learn more * [`DOC /storage/qtrees`](#docs-storage-storage_qtrees)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
**id** | **string**| Qtree ID | 

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


## QtreeGet

> Qtree QtreeGet(ctx, volumeUuid, id, optional)



Retrieves properties for a specific qtree identified by the `volume.uuid` and the `id` in the api path. ### Related ONTAP commands * `qtree show`  ### Learn more * [`DOC /storage/qtrees`](#docs-storage-storage_qtrees)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
**id** | **string**| Qtree ID | 
 **optional** | ***QtreeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QtreeGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Qtree**](qtree.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QtreeModify

> JobLinkResponse QtreeModify(ctx, volumeUuid, id, info)



Updates properties for a specific qtree. ### Related ONTAP commands * `qtree modify` * `qtree rename`  ### Learn more * [`DOC /storage/qtrees`](#docs-storage-storage_qtrees)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
**id** | **string**| Qtree ID | 
**info** | [**Qtree**](Qtree.md)| The new property values for the qtree.  | 

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


## QuotaReportCollectionGet

> QuotaReportResponse QuotaReportCollectionGet(ctx, optional)



Retrieves the quota report records for all FlexVols and FlexGroup volumes. ### Related ONTAP commands * `quota report`  ### Learn more * [`DOC /storage/quota/reports`](#docs-storage-storage_quota_reports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotaReportCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotaReportCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **volumeName** | **optional.String**| Filter by volume.name | 
 **volumeUuid** | **optional.String**| Filter by volume.uuid | 
 **usersId** | **optional.String**| Filter by users.id | 
 **usersName** | **optional.String**| Filter by users.name | 
 **groupId** | **optional.String**| Filter by group.id | 
 **groupName** | **optional.String**| Filter by group.name | 
 **type_** | **optional.String**| Filter by type | 
 **filesUsedTotal** | **optional.Int32**| Filter by files.used.total | 
 **filesUsedHardLimitPercent** | **optional.String**| Filter by files.used.hard_limit_percent | 
 **filesUsedSoftLimitPercent** | **optional.String**| Filter by files.used.soft_limit_percent | 
 **filesHardLimit** | **optional.Int32**| Filter by files.hard_limit | 
 **filesSoftLimit** | **optional.Int32**| Filter by files.soft_limit | 
 **spaceSoftLimit** | **optional.Int32**| Filter by space.soft_limit | 
 **spaceHardLimit** | **optional.Int32**| Filter by space.hard_limit | 
 **spaceUsedSoftLimitPercent** | **optional.String**| Filter by space.used.soft_limit_percent | 
 **spaceUsedTotal** | **optional.Int32**| Filter by space.used.total | 
 **spaceUsedHardLimitPercent** | **optional.String**| Filter by space.used.hard_limit_percent | 
 **index** | **optional.Int32**| Filter by index | 
 **specifier** | **optional.String**| Filter by specifier | 
 **qtreeId** | **optional.Int32**| Filter by qtree.id | 
 **qtreeName** | **optional.String**| Filter by qtree.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**QuotaReportResponse**](quota_report_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaReportGet

> QuotaReport QuotaReportGet(ctx, volumeUuid, index, optional)



Retrieves a specific quota report record. ### Related ONTAP commands * `quota report`  ### Learn more * [`DOC /storage/quota/reports`](#docs-storage-storage_quota_reports)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
**index** | **int32**| Quota report index | 
 **optional** | ***QuotaReportGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotaReportGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**QuotaReport**](quota_report.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaRuleCollectionGet

> QuotaRuleResponse QuotaRuleCollectionGet(ctx, optional)



Retrieves quota policy rules configured for all FlexVols and FlexGroup volumes. ### Related ONTAP commands * `quota policy rule show`  ### Learn more * [`DOC /storage/quota/rules`](#docs-storage-storage_quota_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotaRuleCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotaRuleCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **volumeName** | **optional.String**| Filter by volume.name | 
 **volumeUuid** | **optional.String**| Filter by volume.uuid | 
 **uuid** | **optional.String**| Filter by uuid | 
 **groupName** | **optional.String**| Filter by group.name | 
 **groupId** | **optional.String**| Filter by group.id | 
 **type_** | **optional.String**| Filter by type | 
 **usersName** | **optional.String**| Filter by users.name | 
 **usersId** | **optional.String**| Filter by users.id | 
 **userMapping** | **optional.Bool**| Filter by user_mapping | 
 **filesHardLimit** | **optional.Int32**| Filter by files.hard_limit | 
 **filesSoftLimit** | **optional.Int32**| Filter by files.soft_limit | 
 **spaceSoftLimit** | **optional.Int32**| Filter by space.soft_limit | 
 **spaceHardLimit** | **optional.Int32**| Filter by space.hard_limit | 
 **qtreeId** | **optional.Int32**| Filter by qtree.id | 
 **qtreeName** | **optional.String**| Filter by qtree.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**QuotaRuleResponse**](quota_rule_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaRuleCreate

> JobLinkResponse QuotaRuleCreate(ctx, info)



Creates a quota policy rule for a FlexVol or a FlexGroup volume.<br/> Important notes: * Unlike CLI/ONTAPI, the `quota policy` input is not needed for POST. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the qtree. * `volume.uuid` or `volume.name` - Existing volume in which to create the qtree. * `type` - Quota type for the rule. This type can be `user`, `group`, or `tree`. * `users.name` or `user.id` -  If the quota type is user, this property takes the user name or user ID. For default user quota rules, the user name must be specified as \"\". * `group.name` or `group.id` - If the quota type is group, this property takes the group name or group ID. For default group quota rules, the group name must be specified as \"\". * `qtree.name` - Qtree for which to create the rule. For default tree rules, the qtree name must be specified as \"\". ### Recommended optional properties * `space.hard_limit` - Specifies the space hard limit, in bytes. If less than 1024 bytes, the value is rounded up to 1024 bytes. * `space.soft_limit` - Specifies the space soft limit, in bytes. If less than 1024 bytes, the value is rounded up to 1024 bytes. * `files.hard_limit` - Specifies the hard limit for files. * `files.hard_limit` - Specifies the soft limit for files. * `user_mapping` - Specifies the user_mapping. This property is valid only for quota policy rules of type `user`. ### Related ONTAP commands * `quota policy rule create`  ### Learn more * [`DOC /storage/quota/rules`](#docs-storage-storage_quota_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**QuotaRule**](QuotaRule.md)| Info specification | 

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


## QuotaRuleDelete

> JobLinkResponse QuotaRuleDelete(ctx, ruleUuid)



Deletes a quota policy rule. ### Related ONTAP commands * `quota policy rule delete`  ### Learn more * [`DOC /storage/quota/rules`](#docs-storage-storage_quota_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleUuid** | **string**| Rule UUID | 

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


## QuotaRuleGet

> QuotaRule QuotaRuleGet(ctx, ruleUuid, optional)



Retrieves properties for a specific quota policy rule. ### Related ONTAP commands * `quota policy rule show`  ### Learn more * [`DOC /storage/quota/rules`](#docs-storage-storage_quota_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleUuid** | **string**| Rule UUID | 
 **optional** | ***QuotaRuleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotaRuleGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**QuotaRule**](quota_rule.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaRuleModify

> JobLinkResponse QuotaRuleModify(ctx, ruleUuid, optional)



Updates properties of a specific quota policy rule. <br> Important notes: * The quota resize functionality is supported with the PATCH operation. * Quota resize allows you to modify the quota limits, directly in the filesystem. * The quota must be enabled on a FlexVol or a FlexGroup volume for `quota resize` to take effect. * If the quota is disabled on the volume, the quota policy rule PATCH API modifies the rule, but this does not affect the limits in the filesystem. ### Related ONTAP commands * `quota policy rule modify` * `quota resize`  ### Learn more * [`DOC /storage/quota/rules`](#docs-storage-storage_quota_rules)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleUuid** | **string**| Rule UUID | 
 **optional** | ***QuotaRuleModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotaRuleModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of QuotaRule**](QuotaRule.md)| Info specification | 

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


## ShelfCollectionGet

> ShelfResponse ShelfCollectionGet(ctx, optional)



Retrieves a collection of shelves. ### Related ONTAP commands * `storage shelf show` * `storage shelf port show` * `storage shelf drawer show` ### Learn more * [`DOC /storage/shelves`](#docs-storage-storage_shelves) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShelfCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ShelfCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portsMacAddress** | **optional.String**| Filter by ports.mac_address | 
 **portsInternal** | **optional.Bool**| Filter by ports.internal | 
 **portsCableIdentifier** | **optional.String**| Filter by ports.cable.identifier | 
 **portsCableLength** | **optional.String**| Filter by ports.cable.length | 
 **portsCableSerialNumber** | **optional.String**| Filter by ports.cable.serial_number | 
 **portsCablePartNumber** | **optional.String**| Filter by ports.cable.part_number | 
 **portsModuleId** | **optional.String**| Filter by ports.module_id | 
 **portsState** | **optional.String**| Filter by ports.state | 
 **portsDesignator** | **optional.String**| Filter by ports.designator | 
 **portsRemoteChassis** | **optional.String**| Filter by ports.remote.chassis | 
 **portsRemoteWwn** | **optional.String**| Filter by ports.remote.wwn | 
 **portsRemotePhy** | **optional.String**| Filter by ports.remote.phy | 
 **portsRemoteMacAddress** | **optional.String**| Filter by ports.remote.mac_address | 
 **portsRemotePort** | **optional.String**| Filter by ports.remote.port | 
 **portsId** | **optional.Int32**| Filter by ports.id | 
 **portsWwn** | **optional.String**| Filter by ports.wwn | 
 **model** | **optional.String**| Filter by model | 
 **id** | **optional.String**| Filter by id | 
 **state** | **optional.String**| Filter by state | 
 **diskCount** | **optional.Int32**| Filter by disk_count | 
 **moduleType** | **optional.String**| Filter by module_type | 
 **drawersState** | **optional.String**| Filter by drawers.state | 
 **drawersSerialNumber** | **optional.String**| Filter by drawers.serial_number | 
 **drawersDiskCount** | **optional.Int32**| Filter by drawers.disk_count | 
 **drawersPartNumber** | **optional.String**| Filter by drawers.part_number | 
 **drawersError** | **optional.String**| Filter by drawers.error | 
 **drawersId** | **optional.Int32**| Filter by drawers.id | 
 **drawersClosed** | **optional.Bool**| Filter by drawers.closed | 
 **name** | **optional.String**| Filter by name | 
 **connectionType** | **optional.String**| Filter by connection_type | 
 **pathsName** | **optional.String**| Filter by paths.name | 
 **pathsNodeUuid** | **optional.String**| Filter by paths.node.uuid | 
 **pathsNodeName** | **optional.String**| Filter by paths.node.name | 
 **uid** | **optional.String**| Filter by uid | 
 **baysId** | **optional.Int32**| Filter by bays.id | 
 **baysHasDisk** | **optional.Bool**| Filter by bays.has_disk | 
 **baysType** | **optional.String**| Filter by bays.type | 
 **baysState** | **optional.String**| Filter by bays.state | 
 **internal** | **optional.Bool**| Filter by internal | 
 **frusSerialNumber** | **optional.String**| Filter by frus.serial_number | 
 **frusState** | **optional.String**| Filter by frus.state | 
 **frusPartNumber** | **optional.String**| Filter by frus.part_number | 
 **frusType** | **optional.String**| Filter by frus.type | 
 **frusFirmwareVersion** | **optional.String**| Filter by frus.firmware_version | 
 **frusId** | **optional.Int32**| Filter by frus.id | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ShelfResponse**](shelf_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShelfGet

> Shelf ShelfGet(ctx, uid, optional)



Retrieves a specific shelf. ### Related ONTAP commands * `storage shelf show` * `storage shelf port show` * `storage shelf drawer show` ### Learn more * [`DOC /storage/shelves`](#docs-storage-storage_shelves) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| Shelf UID | 
 **optional** | ***ShelfGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ShelfGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Shelf**](shelf.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotCollectionGet

> SnapshotResponse SnapshotCollectionGet(ctx, volumeUuid, optional)



Retrieves a collection of volume Snapshot copies. ### Related ONTAP commands * `snapshot show` ### Learn more * [`DOC /storage/volumes/{volume.uuid}/snapshots`](#docs-storage-storage_volumes_{volume.uuid}_snapshots) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume | 
 **optional** | ***SnapshotCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snaplockExpiryTime** | **optional.String**| Filter by snaplock_expiry_time | 
 **createTime** | **optional.String**| Filter by create_time | 
 **state** | **optional.String**| Filter by state | 
 **expiryTime** | **optional.String**| Filter by expiry_time | 
 **comment** | **optional.String**| Filter by comment | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **uuid** | **optional.String**| Filter by uuid | 
 **volumeName** | **optional.String**| Filter by volume.name | 
 **volumeUuid2** | **optional.String**| Filter by volume.uuid | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SnapshotResponse**](snapshot_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotCreate

> JobLinkResponse SnapshotCreate(ctx, volumeUuid, optional)



Creates a volume Snapshot copy. ### Required properties * `name` - Name of the Snapshot copy to be created. ### Recommended optional properties * `comment` - Comment associated with the Snapshot copy. * `expiry_time` - Snapshot copies with an expiry time set are not allowed to be deleted until the retention time is reached. ### Related ONTAP commands * `snapshot create` ### Learn more * [`DOC /storage/volumes/{volume.uuid}/snapshots`](#docs-storage-storage_volumes_{volume.uuid}_snapshots) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
 **optional** | ***SnapshotCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of Snapshot**](Snapshot.md)| Info specification | 

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


## SnapshotDelete

> JobLinkResponse SnapshotDelete(ctx, volumeUuid, uuid)



Deletes a Volume Snapshot copy. ### Related ONTAP commands * `snapshot delete` ### Learn more * [`DOC /storage/volumes/{volume.uuid}/snapshots`](#docs-storage-storage_volumes_{volume.uuid}_snapshots) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
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


## SnapshotGet

> Snapshot SnapshotGet(ctx, volumeUuid, uuid, optional)



Retrieves details of a specific volume Snapshot copy. ### Related ONTAP commands * `snapshot show` ### Learn more * [`DOC /storage/volumes/{volume.uuid}/snapshots`](#docs-storage-storage_volumes_{volume.uuid}_snapshots) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
**uuid** | **string**| Snapshot copy UUID | 
 **optional** | ***SnapshotGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Snapshot**](snapshot.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotModify

> JobLinkResponse SnapshotModify(ctx, volumeUuid, uuid, optional)



Updates a Volume Snapshot copy. ### Related ONTAP commands * `snapshot modify` * `snapshot rename` ### Learn more * [`DOC /storage/volumes/{volume.uuid}/snapshots`](#docs-storage-storage_volumes_{volume.uuid}_snapshots) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeUuid** | **string**| Volume UUID | 
**uuid** | **string**| Snapshot copy UUID | 
 **optional** | ***SnapshotModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of Snapshot**](Snapshot.md)| Info specification | 

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


## SnapshotPolicyCollectionGet

> SnapshotPolicyResponse SnapshotPolicyCollectionGet(ctx, optional)



Retrieves a collection of Snapshot copy policies. ### Related ONTAP commands * `snapshot policy show` ### Learn more * [`DOC /storage/snapshot-policies`](#docs-storage-storage_snapshot-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotPolicyCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotPolicyCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by uuid | 
 **copiesPrefix** | **optional.String**| Filter by copies.prefix | 
 **copiesCount** | **optional.Int32**| Filter by copies.count | 
 **copiesSnapmirrorLabel** | **optional.String**| Filter by copies.snapmirror_label | 
 **copiesScheduleName** | **optional.String**| Filter by copies.schedule.name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **name** | **optional.String**| Filter by name | 
 **scope** | **optional.String**| Filter by scope | 
 **comment** | **optional.String**| Filter by comment | 
 **enabled** | **optional.Bool**| Filter by enabled | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SnapshotPolicyResponse**](snapshot_policy_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotPolicyCreate

> SnapshotPolicyCreate(ctx, optional)



Creates a Snapshot copy policy. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the Snapshot copy policy. * `name` - Name for the Snapshot copy policy. * `copies.schedule` - Schedule at which Snapshot copies are captured on the volume. * `copies.count` - Number of Snapshot copies to maintain for this schedule. ### Recommended optional properties * `copies.prefix` - Prefix to use when creating Snapshot copies at regular intervals. ### Default property values If not specified in POST, the following default property values are assigned: * `enabled` - _true_ * `copies.prefix` - Value of `schedule.name` ### Related ONTAP commands * `snapshot policy create` ### Learn more * [`DOC /storage/snapshot-policies`](#docs-storage-storage_snapshot-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotPolicyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotPolicyCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SnapshotPolicy**](SnapshotPolicy.md)| Info specification | 

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


## SnapshotPolicyDelete

> SnapshotPolicyDelete(ctx, uuid)



Deletes a Snapshot copy policy ### Related ONTAP commands * `snapshot policy delete` ### Learn more * [`DOC /storage/snapshot-policies`](#docs-storage-storage_snapshot-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Snapshot policy UUID | 

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


## SnapshotPolicyGet

> SnapshotPolicy SnapshotPolicyGet(ctx, uuid, optional)



Retrieves details of a specific Snapshot copy policy. ### Related ONTAP commands * `snapshot policy show` ### Learn more * [`DOC /storage/snapshot-policies`](#docs-storage-storage_snapshot-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Snapshot copy policy ID | 
 **optional** | ***SnapshotPolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotPolicyGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SnapshotPolicy**](snapshot_policy.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotPolicyModify

> SnapshotPolicyModify(ctx, uuid, optional)



Updates a Snapshot copy policy ### Related ONTAP commands * `snapshot policy modify` * `snapshot policy modify-schedule` * `snapshot policy add-schedule` ### Learn more * [`DOC /storage/snapshot-policies`](#docs-storage-storage_snapshot-policies) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Snapshot policy UUID | 
 **optional** | ***SnapshotPolicyModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SnapshotPolicyModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of SnapshotPolicy**](SnapshotPolicy.md)| Info specification | 

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


## StorageClusterGet

> ClusterSpace StorageClusterGet(ctx, optional)



Reports cluster wide storage details across different tiers. By default, this endpoint returns all fields. Supports the following roles: admin, and readonly. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageClusterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StorageClusterGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ClusterSpace**](cluster_space.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeCollectionGet

> VolumeResponse VolumeCollectionGet(ctx, optional)



Retrieves volumes. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `application.*` * `encryption.*` * `clone.parent_snapshot.name` * `clone.parent_snapshot.uuid` * `clone.parent_svm.name` * `clone.parent_svm.uuid` * `clone.parent_volume.name` * `clone.parent_volume.uuid` * `clone.split_complete_percent` * `clone.split_estimate` * `clone.split_initiated` * `efficiency.*` * `error_state.*` * `files.*` * `nas.export_policy.id` * `nas.gid` * `nas.path` * `nas.security_style` * `nas.uid` * `nas.unix_permissions` * `snaplock.*` * `restore_to.*` * `snapshot_policy.uuid` * `quota.*` * `qos.*` * `flexcache_endpoint_type` * `space.block_storage_inactive_user_data` * `space.capacity_tier_footprint` * `space.footprint` * `space.over_provisioned` * `space.metadata` * `space.logical_space.*` * `space.snapshot.*` * `guarantee.*` * `autosize.*` * `movement.*` * `statistics.*` ### Related ONTAP commands * `volume show` * `volume clone show` * `volume efficiency show` * `volume encryption show` * `volume flexcache show` * `volume flexgroup show` * `volume move show` * `volume quota show` * `volume show-space` * `volume snaplock show`  ### Learn more * [`DOC /storage/volumes`](#docs-storage-storage_volumes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VolumeCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nasUnixPermissions** | **optional.Int32**| Filter by nas.unix_permissions | 
 **nasSecurityStyle** | **optional.String**| Filter by nas.security_style | 
 **nasGid** | **optional.Int32**| Filter by nas.gid | 
 **nasExportPolicyName** | **optional.String**| Filter by nas.export_policy.name | 
 **nasExportPolicyId** | **optional.Int32**| Filter by nas.export_policy.id | 
 **nasPath** | **optional.String**| Filter by nas.path | 
 **nasUid** | **optional.Int32**| Filter by nas.uid | 
 **cloneIsFlexclone** | **optional.Bool**| Filter by clone.is_flexclone | 
 **cloneParentVolumeName** | **optional.String**| Filter by clone.parent_volume.name | 
 **cloneParentVolumeUuid** | **optional.String**| Filter by clone.parent_volume.uuid | 
 **cloneParentSvmUuid** | **optional.String**| Filter by clone.parent_svm.uuid | 
 **cloneParentSvmName** | **optional.String**| Filter by clone.parent_svm.name | 
 **cloneSplitEstimate** | **optional.Int32**| Filter by clone.split_estimate | 
 **cloneSplitCompletePercent** | **optional.Int32**| Filter by clone.split_complete_percent | 
 **cloneSplitInitiated** | **optional.Bool**| Filter by clone.split_initiated | 
 **cloneParentSnapshotUuid** | **optional.String**| Filter by clone.parent_snapshot.uuid | 
 **cloneParentSnapshotName** | **optional.String**| Filter by clone.parent_snapshot.name | 
 **errorStateIsInconsistent** | **optional.Bool**| Filter by error_state.is_inconsistent | 
 **errorStateHasBadBlocks** | **optional.Bool**| Filter by error_state.has_bad_blocks | 
 **flexcacheEndpointType** | **optional.String**| Filter by flexcache_endpoint_type | 
 **statisticsLatencyRawTotal** | **optional.Int32**| Filter by statistics.latency_raw.total | 
 **statisticsLatencyRawOther** | **optional.Int32**| Filter by statistics.latency_raw.other | 
 **statisticsLatencyRawWrite** | **optional.Int32**| Filter by statistics.latency_raw.write | 
 **statisticsLatencyRawRead** | **optional.Int32**| Filter by statistics.latency_raw.read | 
 **statisticsStatus** | **optional.String**| Filter by statistics.status | 
 **statisticsTimestamp** | **optional.String**| Filter by statistics.timestamp | 
 **statisticsIopsRawTotal** | **optional.Int32**| Filter by statistics.iops_raw.total | 
 **statisticsIopsRawOther** | **optional.Int32**| Filter by statistics.iops_raw.other | 
 **statisticsIopsRawWrite** | **optional.Int32**| Filter by statistics.iops_raw.write | 
 **statisticsIopsRawRead** | **optional.Int32**| Filter by statistics.iops_raw.read | 
 **statisticsThroughputRawTotal** | **optional.Int32**| Filter by statistics.throughput_raw.total | 
 **statisticsThroughputRawOther** | **optional.Int32**| Filter by statistics.throughput_raw.other | 
 **statisticsThroughputRawWrite** | **optional.Int32**| Filter by statistics.throughput_raw.write | 
 **statisticsThroughputRawRead** | **optional.Int32**| Filter by statistics.throughput_raw.read | 
 **qosPolicyUuid** | **optional.String**| Filter by qos.policy.uuid | 
 **qosPolicyMinThroughputIops** | **optional.Int32**| Filter by qos.policy.min_throughput_iops | 
 **qosPolicyMaxThroughputMbps** | **optional.Int32**| Filter by qos.policy.max_throughput_mbps | 
 **qosPolicyName** | **optional.String**| Filter by qos.policy.name | 
 **qosPolicyMaxThroughputIops** | **optional.Int32**| Filter by qos.policy.max_throughput_iops | 
 **spaceSize** | **optional.Int32**| Filter by space.size | 
 **spaceOverProvisioned** | **optional.Int32**| Filter by space.over_provisioned | 
 **spaceAvailable** | **optional.Int32**| Filter by space.available | 
 **spaceMetadata** | **optional.Int32**| Filter by space.metadata | 
 **spaceLogicalSpaceUsedByAfs** | **optional.Int32**| Filter by space.logical_space.used_by_afs | 
 **spaceLogicalSpaceReporting** | **optional.Bool**| Filter by space.logical_space.reporting | 
 **spaceLogicalSpaceEnforcement** | **optional.Bool**| Filter by space.logical_space.enforcement | 
 **spaceLogicalSpaceAvailable** | **optional.Int32**| Filter by space.logical_space.available | 
 **spaceBlockStorageInactiveUserData** | **optional.Int32**| Filter by space.block_storage_inactive_user_data | 
 **spaceCapacityTierFootprint** | **optional.Int32**| Filter by space.capacity_tier_footprint | 
 **spaceUsed** | **optional.Int32**| Filter by space.used | 
 **spaceFootprint** | **optional.Int32**| Filter by space.footprint | 
 **spaceSnapshotReservePercent** | **optional.Int32**| Filter by space.snapshot.reserve_percent | 
 **spaceSnapshotUsed** | **optional.Int32**| Filter by space.snapshot.used | 
 **name** | **optional.String**| Filter by name | 
 **movementPercentComplete** | **optional.String**| Filter by movement.percent_complete | 
 **movementState** | **optional.String**| Filter by movement.state | 
 **movementDestinationAggregateUuid** | **optional.String**| Filter by movement.destination_aggregate.uuid | 
 **movementDestinationAggregateName** | **optional.String**| Filter by movement.destination_aggregate.name | 
 **movementCutoverWindow** | **optional.Int32**| Filter by movement.cutover_window | 
 **snapshotPolicyName** | **optional.String**| Filter by snapshot_policy.name | 
 **snapshotPolicyUuid** | **optional.String**| Filter by snapshot_policy.uuid | 
 **size** | **optional.Int32**| Filter by size | 
 **snaplockType** | **optional.String**| Filter by snaplock.type | 
 **snaplockComplianceClockTime** | **optional.String**| Filter by snaplock.compliance_clock_time | 
 **snaplockPrivilegedDelete** | **optional.String**| Filter by snaplock.privileged_delete | 
 **snaplockAutocommitPeriod** | **optional.String**| Filter by snaplock.autocommit_period | 
 **snaplockIsAuditLog** | **optional.Bool**| Filter by snaplock.is_audit_log | 
 **snaplockAppendModeEnabled** | **optional.Bool**| Filter by snaplock.append_mode_enabled | 
 **snaplockExpiryTime** | **optional.String**| Filter by snaplock.expiry_time | 
 **snaplockRetentionMaximum** | **optional.String**| Filter by snaplock.retention.maximum | 
 **snaplockRetentionMinimum** | **optional.String**| Filter by snaplock.retention.minimum | 
 **snaplockRetentionDefault** | **optional.String**| Filter by snaplock.retention.default | 
 **snaplockLitigationCount** | **optional.Int32**| Filter by snaplock.litigation_count | 
 **createTime** | **optional.String**| Filter by create_time | 
 **efficiencyDedupe** | **optional.String**| Filter by efficiency.dedupe | 
 **efficiencyCrossVolumeDedupe** | **optional.String**| Filter by efficiency.cross_volume_dedupe | 
 **efficiencyCompression** | **optional.String**| Filter by efficiency.compression | 
 **efficiencyCompaction** | **optional.String**| Filter by efficiency.compaction | 
 **autosizeShrinkThreshold** | **optional.Int32**| Filter by autosize.shrink_threshold | 
 **autosizeMode** | **optional.String**| Filter by autosize.mode | 
 **autosizeGrowThreshold** | **optional.Int32**| Filter by autosize.grow_threshold | 
 **autosizeMaximum** | **optional.Int32**| Filter by autosize.maximum | 
 **autosizeMinimum** | **optional.Int32**| Filter by autosize.minimum | 
 **comment** | **optional.String**| Filter by comment | 
 **guaranteeType** | **optional.String**| Filter by guarantee.type | 
 **guaranteeHonored** | **optional.Bool**| Filter by guarantee.honored | 
 **uuid** | **optional.String**| Filter by uuid | 
 **metricTimestamp** | **optional.String**| Filter by metric.timestamp | 
 **metricIopsTotal** | **optional.Int32**| Filter by metric.iops.total | 
 **metricIopsOther** | **optional.Int32**| Filter by metric.iops.other | 
 **metricIopsWrite** | **optional.Int32**| Filter by metric.iops.write | 
 **metricIopsRead** | **optional.Int32**| Filter by metric.iops.read | 
 **metricDuration** | **optional.String**| Filter by metric.duration | 
 **metricThroughputTotal** | **optional.Int32**| Filter by metric.throughput.total | 
 **metricThroughputOther** | **optional.Int32**| Filter by metric.throughput.other | 
 **metricThroughputWrite** | **optional.Int32**| Filter by metric.throughput.write | 
 **metricThroughputRead** | **optional.Int32**| Filter by metric.throughput.read | 
 **metricLatencyTotal** | **optional.Int32**| Filter by metric.latency.total | 
 **metricLatencyOther** | **optional.Int32**| Filter by metric.latency.other | 
 **metricLatencyWrite** | **optional.Int32**| Filter by metric.latency.write | 
 **metricLatencyRead** | **optional.Int32**| Filter by metric.latency.read | 
 **metricStatus** | **optional.String**| Filter by metric.status | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **quotaState** | **optional.String**| Filter by quota.state | 
 **tieringPolicy** | **optional.String**| Filter by tiering.policy | 
 **style** | **optional.String**| Filter by style | 
 **type_** | **optional.String**| Filter by type | 
 **encryptionState** | **optional.String**| Filter by encryption.state | 
 **encryptionKeyId** | **optional.String**| Filter by encryption.key_id | 
 **encryptionRekey** | **optional.Bool**| Filter by encryption.rekey | 
 **encryptionStatusCode** | **optional.String**| Filter by encryption.status.code | 
 **encryptionStatusMessage** | **optional.String**| Filter by encryption.status.message | 
 **encryptionEnabled** | **optional.Bool**| Filter by encryption.enabled | 
 **encryptionType** | **optional.String**| Filter by encryption.type | 
 **filesUsed** | **optional.Int32**| Filter by files.used | 
 **filesMaximum** | **optional.Int32**| Filter by files.maximum | 
 **aggregatesUuid** | **optional.String**| Filter by aggregates.uuid | 
 **aggregatesName** | **optional.String**| Filter by aggregates.name | 
 **language** | **optional.String**| Filter by language | 
 **applicationName** | **optional.String**| Filter by application.name | 
 **applicationUuid** | **optional.String**| Filter by application.uuid | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**VolumeResponse**](volume_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeCollectionPerformanceMetricsGet

> VolumeMetricsResponse VolumeCollectionPerformanceMetricsGet(ctx, uuid, optional)



Retrieves historical performance metrics for a volume.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of the volume. | 
 **optional** | ***VolumeCollectionPerformanceMetricsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VolumeCollectionPerformanceMetricsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **optional.String**| Filter by timestamp | 
 **iopsTotal** | **optional.Int32**| Filter by iops.total | 
 **iopsOther** | **optional.Int32**| Filter by iops.other | 
 **iopsWrite** | **optional.Int32**| Filter by iops.write | 
 **iopsRead** | **optional.Int32**| Filter by iops.read | 
 **duration** | **optional.String**| Filter by duration | 
 **throughputTotal** | **optional.Int32**| Filter by throughput.total | 
 **throughputOther** | **optional.Int32**| Filter by throughput.other | 
 **throughputWrite** | **optional.Int32**| Filter by throughput.write | 
 **throughputRead** | **optional.Int32**| Filter by throughput.read | 
 **latencyTotal** | **optional.Int32**| Filter by latency.total | 
 **latencyOther** | **optional.Int32**| Filter by latency.other | 
 **latencyWrite** | **optional.Int32**| Filter by latency.write | 
 **latencyRead** | **optional.Int32**| Filter by latency.read | 
 **status** | **optional.String**| Filter by status | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **interval** | **optional.String**| The time range for the data. Examples can be 1h, 1d, 1m, 1w, 1y. The period for each time range is as follows: * 1h: Metrics over the most recent hour sampled over 15 seconds. * 1d: Metrics over the most recent day sampled over 4 minutes. * 1w: Metrics over the most recent week sampled over 30 minutes. * 1m: Metrics over the most recent month sampled over 2 hours. * 1y: Metrics over the most recent year sampled over a day.  | [default to 1h]

### Return type

[**VolumeMetricsResponse**](volume_metrics_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeCreate

> JobLinkResponse VolumeCreate(ctx, info)



Creates a volume on a specified SVM and storage aggregates. ### Required properties * `svm.uuid` or `svm.name` - Existing SVM in which to create the volume. * `name` - Name of the volume. * `aggregates.name` or `aggregates.uuid` - Existing aggregates in which to create the volume. ### Default property values * `state` -  _online_ * `size` - _20MB_ * `style` - _flexvol_ * `type` - _rw_ * `encryption.enabled` - _false_ * `snapshot_policy.name` - _default_ * `gaurantee.type` - _volume_ ### Related ONTAP commands * `volume create` * `volume clone create`  ### Learn more * [`DOC /storage/volumes`](#docs-storage-storage_volumes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**info** | [**Volume**](Volume.md)| Info specification | 

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


## VolumeDelete

> JobLinkResponse VolumeDelete(ctx, uuid)



Deletes a volume. If the UUID belongs to a volume, all of its blocks are freed and returned to its containing aggregate. If a volume is online, it is offlined before deletion. ### Related ONTAP commands * `volume delete` * `volume clone delete`  ### Learn more * [`DOC /storage/volumes`](#docs-storage-storage_volumes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of the volume. | 

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


## VolumeGet

> Volume VolumeGet(ctx, uuid, optional)



Retrieves a volume. The GET API can be used to retrieve the quota state for a FlexVol or a FlexGroup volume. ### Expensive properties There is an added cost to retrieving values for these properties. They are not included by default in GET results and must be explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. * `application.*` * `encryption.*` * `clone.parent_snapshot.name` * `clone.parent_snapshot.uuid` * `clone.parent_svm.name` * `clone.parent_svm.uuid` * `clone.parent_volume.name` * `clone.parent_volume.uuid` * `clone.split_complete_percent` * `clone.split_estimate` * `clone.split_initiated` * `efficiency.*` * `error_state.*` * `files.*` * `nas.export_policy.id` * `nas.gid` * `nas.path` * `nas.security_style` * `nas.uid` * `nas.unix_permissions` * `snaplock.*` * `restore_to.*` * `snapshot_policy.uuid` * `quota.*` * `qos.*` * `flexcache_endpoint_type` * `space.block_storage_inactive_user_data` * `space.capacity_tier_footprint` * `space.footprint` * `space.over_provisioned` * `space.metadata` * `space.logical_space.*` * `space.snapshot.*` * `guarantee.*` * `autosize.*` * `movement.*` * `statistics.*` ### Related ONTAP commands * `volume show` * `volume clone show` * `volume efficiency show` * `volume encryption show` * `volume flexcache show` * `volume flexgroup show` * `volume move show` * `volume quota show` * `volume show-space` * `volume snaplock show`  ### Learn more * [`DOC /storage/volumes`](#docs-storage-storage_volumes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of the volume. | 
 **optional** | ***VolumeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VolumeGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Volume**](volume.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeModify

> JobLinkResponse VolumeModify(ctx, uuid, optional)



Updates the attributes of a volume. For movement, use the \"validate_only\" field on the request to validate but not perform the operation. The PATCH API can be used to enable or disable quotas for a FlexVol or a FlexGroup volume. ### Related ONTAP commands * `volume modify` * `volume clone modify` * `volume efficiency modify` * `volume quota on` * `volume quota off`  ### Learn more * [`DOC /storage/volumes`](#docs-storage-storage_volumes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Unique identifier of the volume. | 
 **optional** | ***VolumeModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VolumeModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreToSnapshotUuid** | **optional.String**| UUID of the Snapshot copy to restore volume to the point in time the Snapshot copy was taken. | 
 **restoreToSnapshotName** | **optional.String**| Name of the Snapshot copy to restore volume to the point in time the Snapshot copy was taken. | 
 **sizingMethod** | **optional.String**| Represents the method to modify the size of a Flexgroup. The following methods are supported: * use_existing_resources - Increases or decreases the size of the FlexGroup by increasing or decreasing the size of the current FlexGroup resources * add_new_resources - Increases the size of the FlexGroup by adding new resources  | [default to use_existing_resources]
 **movementDestinationAggregateName** | **optional.String**| Name of the aggregates to which the specified volume can be moved. | 
 **movementDestinationAggregateUuid** | **optional.String**| UUID of the aggregates to which the specified volume can be moved. | 
 **validateOnly** | **optional.Bool**| Validate the operation and its parameters, without actually performing the operation. | 
 **info** | [**optional.Interface of Volume**](Volume.md)| Info specification | 

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

