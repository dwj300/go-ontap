# \ClusterApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChassisCollectionGet**](ClusterApi.md#ChassisCollectionGet) | **Get** /cluster/chassis | 
[**ChassisGet**](ClusterApi.md#ChassisGet) | **Get** /cluster/chassis/{id} | 
[**ClusterCollectionPerformanceMetricsGet**](ClusterApi.md#ClusterCollectionPerformanceMetricsGet) | **Get** /cluster/metrics | 
[**ClusterCreate**](ClusterApi.md#ClusterCreate) | **Post** /cluster | 
[**ClusterGet**](ClusterApi.md#ClusterGet) | **Get** /cluster | 
[**ClusterModify**](ClusterApi.md#ClusterModify) | **Patch** /cluster | 
[**ClusterPeerCreate**](ClusterApi.md#ClusterPeerCreate) | **Post** /cluster/peers | 
[**ClusterPeerDelete**](ClusterApi.md#ClusterPeerDelete) | **Delete** /cluster/peers/{uuid} | 
[**ClusterPeerGet**](ClusterApi.md#ClusterPeerGet) | **Get** /cluster/peers/{uuid} | 
[**ClusterPeerModify**](ClusterApi.md#ClusterPeerModify) | **Patch** /cluster/peers/{uuid} | 
[**JobCollectionGet**](ClusterApi.md#JobCollectionGet) | **Get** /cluster/jobs | 
[**JobGet**](ClusterApi.md#JobGet) | **Get** /cluster/jobs/{uuid} | 
[**JobModify**](ClusterApi.md#JobModify) | **Patch** /cluster/jobs/{uuid} | 
[**LicenseCreate**](ClusterApi.md#LicenseCreate) | **Post** /cluster/licensing/licenses | 
[**LicenseDelete**](ClusterApi.md#LicenseDelete) | **Delete** /cluster/licensing/licenses/{name} | 
[**LicenseGet**](ClusterApi.md#LicenseGet) | **Get** /cluster/licensing/licenses/{name} | 
[**LicensesGet**](ClusterApi.md#LicensesGet) | **Get** /cluster/licensing/licenses | 
[**NodeGet**](ClusterApi.md#NodeGet) | **Get** /cluster/nodes/{uuid} | 
[**NodeModify**](ClusterApi.md#NodeModify) | **Patch** /cluster/nodes/{uuid} | 
[**NodesCreate**](ClusterApi.md#NodesCreate) | **Post** /cluster/nodes | 
[**NodesGet**](ClusterApi.md#NodesGet) | **Get** /cluster/nodes | 
[**PeersCollectionGet**](ClusterApi.md#PeersCollectionGet) | **Get** /cluster/peers | 
[**ScheduleCollectionGet**](ClusterApi.md#ScheduleCollectionGet) | **Get** /cluster/schedules | 
[**ScheduleCreate**](ClusterApi.md#ScheduleCreate) | **Post** /cluster/schedules | 
[**ScheduleDelete**](ClusterApi.md#ScheduleDelete) | **Delete** /cluster/schedules/{uuid} | 
[**ScheduleGet**](ClusterApi.md#ScheduleGet) | **Get** /cluster/schedules/{uuid} | 
[**ScheduleModify**](ClusterApi.md#ScheduleModify) | **Patch** /cluster/schedules/{uuid} | 
[**SoftwareGet**](ClusterApi.md#SoftwareGet) | **Get** /cluster/software | 
[**SoftwareHistoryCollectionGet**](ClusterApi.md#SoftwareHistoryCollectionGet) | **Get** /cluster/software/history | 
[**SoftwareModify**](ClusterApi.md#SoftwareModify) | **Patch** /cluster/software | 
[**SoftwarePackageCreate**](ClusterApi.md#SoftwarePackageCreate) | **Post** /cluster/software/download | 
[**SoftwarePackageDelete**](ClusterApi.md#SoftwarePackageDelete) | **Delete** /cluster/software/packages/{version} | 
[**SoftwarePackageGet**](ClusterApi.md#SoftwarePackageGet) | **Get** /cluster/software/packages/{version} | 
[**SoftwarePackagesCollectionGet**](ClusterApi.md#SoftwarePackagesCollectionGet) | **Get** /cluster/software/packages | 


# **ChassisCollectionGet**
> ChassisResponse ChassisCollectionGet(ctx, optional)


Retrieves a collection of chassis. ### Related ONTAP commands * `system chassis show` * `system chassis fru show` ### Learn more * [`DOC /cluster/chassis`](#docs-cluster-cluster_chassis) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChassisCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChassisCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Filter by id | 
 **shelvesUid** | **optional.String**| Filter by shelves.uid | 
 **nodesUuid** | **optional.String**| Filter by nodes.uuid | 
 **nodesName** | **optional.String**| Filter by nodes.name | 
 **state** | **optional.String**| Filter by state | 
 **frusId** | **optional.String**| Filter by frus.id | 
 **frusType** | **optional.String**| Filter by frus.type | 
 **frusState** | **optional.String**| Filter by frus.state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ChassisResponse**](chassis_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChassisGet**
> Chassis ChassisGet(ctx, id, optional)


Retrieves a specific chassis. ### Related ONTAP commands * `system chassis show` * `system chassis fru show` ### Learn more * [`DOC /cluster/chassis`](#docs-cluster-cluster_chassis) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Chassis ID | 
 **optional** | ***ChassisGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChassisGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Chassis**](chassis.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterCollectionPerformanceMetricsGet**
> ClusterMetricsResponse ClusterCollectionPerformanceMetricsGet(ctx, optional)


Retrieves historical performance metrics for the cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterCollectionPerformanceMetricsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterCollectionPerformanceMetricsGetOpts struct

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
 **interval** | **optional.String**| The time range for the data. Examples can be 1h, 1d, 1m, 1w, 1y. The period for each time range is specified as follows: * 1h: Metrics over the most recent hour sampled over 15 second. * 1d: Metrics over the most recent day sampled over 4 minutes. * 1w: Metrics over the most recent week sampled over 30 minutes. * 1m: Metrics over the most recent month sampled over 2 hours. * 1y: Metrics over the most recent year sampled over a day.  | [default to 1h]

### Return type

[**ClusterMetricsResponse**](cluster_metrics_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterCreate**
> JobLinkResponse ClusterCreate(ctx, info, optional)


Sets up a cluster. ### Required properties * `name` * `password` ### Recommended optional properties * `location` * `contact` * `dns_domains` * `name_servers` * `ntp_servers` * `license` * `configuration_backup` * `management_interface` * `nodes` ### Learn more * [`DOC /cluster`](#docs-cluster-cluster) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**Cluster**](Cluster.md)| Cluster information | 
 **optional** | ***ClusterCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singleNodeCluster** | **optional.Bool**| Configures a single node cluster.  All cluster ports are reassigned to the default network.  The storage failover settings are configured to non-HA.  The node reboots during this operation. | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterGet**
> Cluster ClusterGet(ctx, optional)


Retrieves the cluster configuration. ### Learn more * [`DOC /cluster`](#docs-cluster-cluster)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Cluster**](cluster.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterModify**
> JobLinkResponse ClusterModify(ctx, info)


Updates the cluster configuration once the cluster has been created. ### Learn more * [`DOC /cluster`](#docs-cluster-cluster)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**Cluster**](Cluster.md)| Cluster information | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterPeerCreate**
> ClusterPeerSetupResponse ClusterPeerCreate(ctx, optional)


Creates a peering relationship and, optionally, the IP interfaces it will use. There are two ways to create a peering relationship. ### Provide remote IP Here the user provides the remote IP address. Creating a new cluster peer relationship with a specific remote cluster requires at least one remote intercluster IP address from that cluster. ### Required properties   * `remote.ip_addresses` - Addresses of the remote peers. The local peer must be able to reach and connect to these addresses for the request to succeed in creating a peer.   * Either set `generate_passphrase` to true or provide a passphrase in the body of the request; only one of them is required. ### Optional properties   The following fields are optional for a POST on /cluster/peer/. All fields must follow the structure in the cluster peer API definition.   * `name` - Name of the peering relationship.   * `passphrase` - User generated passphrase for use in authentication.   * `generate_passphrase` (true/false) - When this option is true, ONTAP automatically generates a passphrase to authenticate cluster peers.   * `ipspace` - IPspace of the local intercluster LIFs. Assumes Default IPspace if not provided.   * `initial_allowed_svms` - the local SVMs allowed to peer with the peer cluster's SVMs. This list can be modified until the remote cluster accepts this cluster peering relationship.   * `local_network` - fields to create a local intercluster LIF. See section on \"Creating local intercluster lifs\".   * `expiry_time` - Duration in ISO 8601 format for which the user-supplied or auto-generated passphrase is valid. Expiration time must not be greater than seven days into the future. ISO 8601 duration format is \"PnDTnHnMnS\" or \"PnW\" where n is a positive integer. The nD, nH, nM and nS fields can be dropped if zero. \"P\" should always be present and \"T\" should be present if there are any hours, minutes or seconds fields.   * `encryption_proposed` (none/tls-psk) - Encryption mechanism of the communication channel between the two peers. ### Do not provide remote IP   This method is used when the remote IP address is not provided. This method is used when the filer is ready to accept peering request from foreign clusters. ### Required properties   * `generate_passphrase` (true) - This option must be set to  true. ONTAP automatically generates a passphrase to authenticate cluster peers. Either set generate_passphrase to true or provide a passphrase in the body of the request; only one of them is required. ### Optional properties   The following fields are optional for a POST on /cluster/peer/. All fields must follow the structure in the cluster peer API definition.   * `name` - Name of the remote peer.   * `ipspace` - IPspace of the local intercluster LIFs. Assumes Default IPspace if not provided.   * `initial_allowed_svms` - Local SVMs allowed to peer with the peer cluster's SVMs. This list can be modified until the remote cluster accepts this cluster peering relationship.   * `local_network` - Fields to create a local intercluster LIF. See section on \"Creating local intercluster lifs\".   * `expiry_time` - Duration in ISO 8601 format for which the user-supplied or auto-generated passphrase is valid. Expiration time must not be greater than seven days into the future. ISO 8601 duration format is \"PnDTnHnMnS\" or \"PnW\" where n is a positive integer. The nD, nH, nM and nS fields can be dropped if zero. \"P\" should always be present and \"T\" should be present if there are any hours, minutes or seconds fields.   * `encryption_proposed` (none/tls-psk) - Encryption mechanism of the communication channel between the two peers. ## Additional information As with creating a cluster peer through the CLI, the combinations of options must be valid in order for the create operation to succeed. The following list shows the combinations that will succeed and those that will fail: * a passphrase only (fail) * a peer IP address (fail) * a passphrase with an expiration time > 7 days into the future (fail) * peer IP address and a passphrase (OK) * generate_passphrase=true (OK) * any proposed encryption protocol (OK) * an IPspace name or UUID (OK) * a passphrase, peer IP address, and any proposed encryption protocol (OK) * a non empty list initial allowed vserver peer names or UUIDs. (OK) ### Learn more * [`DOC /cluster/peers`](#docs-cluster-cluster_peers) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterPeerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterPeerCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of ClusterPeer**](ClusterPeer.md)| Info specification | 

### Return type

[**ClusterPeerSetupResponse**](cluster_peer_setup_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterPeerDelete**
> ClusterPeerDelete(ctx, uuid)


Deletes a cluster peer. ### Learn more * [`DOC /cluster/peers`](#docs-cluster-cluster_peers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Cluster peer relationship UUID | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterPeerGet**
> ClusterPeer ClusterPeerGet(ctx, uuid, optional)


Retrieves a specific cluster peer instance. ### Learn more * [`DOC /cluster/peers`](#docs-cluster-cluster_peers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Cluster peer relationship UUID | 
 **optional** | ***ClusterPeerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterPeerGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**ClusterPeer**](cluster_peer.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterPeerModify**
> ClusterPeerSetupResponse ClusterPeerModify(ctx, uuid, optional)


Updates a cluster peer instance. ### Learn more * [`DOC /cluster/peers`](#docs-cluster-cluster_peers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Cluster peer relationship UUID | 
 **optional** | ***ClusterPeerModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterPeerModifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of ClusterPeer**](ClusterPeer.md)| Info specification | 

### Return type

[**ClusterPeerSetupResponse**](cluster_peer_setup_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobCollectionGet**
> JobResponse JobCollectionGet(ctx, optional)


Retrieves a list of recently running asynchronous jobs. Once a job transitions to a failure or success state, it is deleted after a default time of 300 seconds. ### Learn more * [`DOC /cluster/jobs`](#docs-cluster-cluster_jobs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **optional.String**| Filter by start_time | 
 **endTime** | **optional.String**| Filter by end_time | 
 **uuid** | **optional.String**| Filter by uuid | 
 **code** | **optional.Int32**| Filter by code | 
 **description** | **optional.String**| Filter by description | 
 **message** | **optional.String**| Filter by message | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**JobResponse**](job_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobGet**
> Job JobGet(ctx, uuid, optional)


Retrieves the details of a specific asynchronous job. Once a job transitions to a failure or success state, it is deleted after a default time of 300 seconds. ### Learn more * [`DOC /cluster/jobs`](#docs-cluster-cluster_jobs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Job UUID | 
 **optional** | ***JobGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Job**](job.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobModify**
> JobModify(ctx, uuid, optional)


Updates the state of a specific asynchronous job. ### Learn more * [`DOC /cluster/jobs`](#docs-cluster-cluster_jobs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Job UUID | 
 **optional** | ***JobModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobModifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**| Request a job to pause, resume, or cancel. Note: Not all jobs support these actions. A job can only be resumed if it is in paused state. Upon successfully requesting a job to be cancelled, the job state changes to either success or failure.  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LicenseCreate**
> LicensePackageResponse LicenseCreate(ctx, info, optional)


Installs one or more feature licenses. ### Required properties * `keys` - Array containing a list of NLF or 26-character license keys. ### Related ONTAP commands * `system license add`  ### Learn more * [`DOC /cluster/licensing/licenses`](#docs-cluster-cluster_licensing_licenses)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**LicensePackage**](LicensePackage.md)| List of NLF or 26-character keys. | 
 **optional** | ***LicenseCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnRecords** | **optional.Bool**| The default is false.  If set to true, the records are returned. | [default to false]

### Return type

[**LicensePackageResponse**](license_package_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LicenseDelete**
> LicenseDelete(ctx, name, serialNumber)


Deletes a license. ### Related ONTAP commands * `system license delete`  ### Learn more * [`DOC /cluster/licensing/licenses/{name}`](#docs-cluster-cluster_licensing_licenses_{name})

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the license package to delete. | 
  **serialNumber** | **string**| Serial number of the license to delete. | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LicenseGet**
> LicensePackage LicenseGet(ctx, name, optional)


Retrieves a specific license package. ### Related ONTAP commands * `system license show` * `system license show-status`  ### Learn more * [`DOC /cluster/licensing/licenses/{name}`](#docs-cluster-cluster_licensing_licenses_{name})

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the license package. | 
 **optional** | ***LicenseGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**| Filter by state | 
 **name2** | **optional.String**| Filter by name | 
 **licensesCapacityMaximumSize** | **optional.Int32**| Filter by licenses.capacity.maximum_size | 
 **licensesCapacityUsedSize** | **optional.Int32**| Filter by licenses.capacity.used_size | 
 **licensesStartTime** | **optional.String**| Filter by licenses.start_time | 
 **licensesSerialNumber** | **optional.String**| Filter by licenses.serial_number | 
 **licensesOwner** | **optional.String**| Filter by licenses.owner | 
 **licensesEvaluation** | **optional.Bool**| Filter by licenses.evaluation | 
 **licensesActive** | **optional.Bool**| Filter by licenses.active | 
 **licensesExpiryTime** | **optional.String**| Filter by licenses.expiry_time | 
 **licensesComplianceState** | **optional.String**| Filter by licenses.compliance.state | 
 **scope** | **optional.String**| Filter by scope | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**LicensePackage**](license_package.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LicensesGet**
> LicensePackageResponse LicensesGet(ctx, optional)


Retrieves a collection of license packages. ### Related ONTAP commands * `system license show-status` * `system license show`  ### Learn more * [`DOC /cluster/licensing/licenses`](#docs-cluster-cluster_licensing_licenses)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LicensesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicensesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter by state | 
 **name** | **optional.String**| Filter by name | 
 **licensesCapacityMaximumSize** | **optional.Int32**| Filter by licenses.capacity.maximum_size | 
 **licensesCapacityUsedSize** | **optional.Int32**| Filter by licenses.capacity.used_size | 
 **licensesStartTime** | **optional.String**| Filter by licenses.start_time | 
 **licensesSerialNumber** | **optional.String**| Filter by licenses.serial_number | 
 **licensesOwner** | **optional.String**| Filter by licenses.owner | 
 **licensesEvaluation** | **optional.Bool**| Filter by licenses.evaluation | 
 **licensesActive** | **optional.Bool**| Filter by licenses.active | 
 **licensesExpiryTime** | **optional.String**| Filter by licenses.expiry_time | 
 **licensesComplianceState** | **optional.String**| Filter by licenses.compliance.state | 
 **scope** | **optional.String**| Filter by scope | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**LicensePackageResponse**](license_package_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeGet**
> Node NodeGet(ctx, uuid, optional)


Retrieves information for the node. ### Learn more * [`DOC /cluster/nodes`](#docs-cluster-cluster_nodes)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 
 **optional** | ***NodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Node**](node.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeModify**
> JobLinkResponse NodeModify(ctx, uuid, optional)


Updates the node information or performs shutdown/reboot actions on a node. ### Learn more * [`DOC /cluster/nodes`](#docs-cluster-cluster_nodes)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 
 **optional** | ***NodeModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeModifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**| The shutdown action shuts the node down and transfers storage control to its HA group if storage failover is enabled. The reboot action reboots the node and transfers storage control to its HA group if storage failover is enabled.  | 
 **shutdownRebootReason** | **optional.String**| Indicates the reason for the reboot or shutdown. This only applies when an action of reboot or shutdown is provided.  | 
 **allowDataOutage** | **optional.Bool**| This only applies when an action of reboot or shutdown is provided. It allows storage failover to be bypassed along with any failures related to mainintaing quorum in the cluster.  | [default to false]
 **info** | [**optional.Interface of Node**](Node.md)|  | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesCreate**
> JobLinkResponse NodesCreate(ctx, info)


Adds a node or nodes to the cluster ### Required properties * `cluster_interface.ip.address`  ### Learn more * [`DOC /cluster/nodes`](#docs-cluster-cluster_nodes)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **info** | [**Node**](Node.md)| An object containing an array of nodes. | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesGet**
> NodeResponse NodesGet(ctx, optional)


Retrieves the nodes in the cluster. ### Learn more * [`DOC /cluster/nodes`](#docs-cluster-cluster_nodes)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membership** | **optional.String**| Filter by membership | 
 **name** | **optional.String**| Filter by name | 
 **model** | **optional.String**| Filter by model | 
 **versionMajor** | **optional.Int32**| Filter by version.major | 
 **versionMinor** | **optional.Int32**| Filter by version.minor | 
 **versionGeneration** | **optional.Int32**| Filter by version.generation | 
 **versionFull** | **optional.String**| Filter by version.full | 
 **location** | **optional.String**| Filter by location | 
 **serviceProcessorIpv4InterfaceNetmask** | **optional.String**| Filter by service_processor.ipv4_interface.netmask | 
 **serviceProcessorIpv4InterfaceAddress** | **optional.String**| Filter by service_processor.ipv4_interface.address | 
 **serviceProcessorIpv4InterfaceGateway** | **optional.String**| Filter by service_processor.ipv4_interface.gateway | 
 **serviceProcessorMacAddress** | **optional.String**| Filter by service_processor.mac_address | 
 **serviceProcessorState** | **optional.String**| Filter by service_processor.state | 
 **serviceProcessorFirmwareVersion** | **optional.String**| Filter by service_processor.firmware_version | 
 **serviceProcessorIpv6InterfaceNetmask** | **optional.String**| Filter by service_processor.ipv6_interface.netmask | 
 **serviceProcessorIpv6InterfaceAddress** | **optional.String**| Filter by service_processor.ipv6_interface.address | 
 **serviceProcessorIpv6InterfaceGateway** | **optional.String**| Filter by service_processor.ipv6_interface.gateway | 
 **serviceProcessorLinkStatus** | **optional.String**| Filter by service_processor.link_status | 
 **serviceProcessorDhcpEnabled** | **optional.Bool**| Filter by service_processor.dhcp_enabled | 
 **uptime** | **optional.Int32**| Filter by uptime | 
 **controllerOverTemperature** | **optional.String**| Filter by controller.over_temperature | 
 **controllerFrusState** | **optional.String**| Filter by controller.frus.state | 
 **controllerFrusType** | **optional.String**| Filter by controller.frus.type | 
 **controllerFrusId** | **optional.Int32**| Filter by controller.frus.id | 
 **controllerFlashCacheState** | **optional.String**| Filter by controller.flash_cache.state | 
 **controllerFlashCacheModel** | **optional.String**| Filter by controller.flash_cache.model | 
 **controllerFlashCacheSerialNumber** | **optional.String**| Filter by controller.flash_cache.serial_number | 
 **controllerFlashCachePartNumber** | **optional.String**| Filter by controller.flash_cache.part_number | 
 **controllerFlashCacheHardwareRevision** | **optional.String**| Filter by controller.flash_cache.hardware_revision | 
 **controllerFlashCacheSlot** | **optional.String**| Filter by controller.flash_cache.slot | 
 **controllerFlashCacheFirmwareVersion** | **optional.String**| Filter by controller.flash_cache.firmware_version | 
 **controllerFlashCacheCapacity** | **optional.Int32**| Filter by controller.flash_cache.capacity | 
 **haEnabled** | **optional.Bool**| Filter by ha.enabled | 
 **haPartnersUuid** | **optional.String**| Filter by ha.partners.uuid | 
 **haPartnersName** | **optional.String**| Filter by ha.partners.name | 
 **haAutoGiveback** | **optional.Bool**| Filter by ha.auto_giveback | 
 **date** | **optional.String**| Filter by date | 
 **uuid** | **optional.String**| Filter by uuid | 
 **serialNumber** | **optional.String**| Filter by serial_number | 
 **clusterInterfacesName** | **optional.String**| Filter by cluster_interfaces.name | 
 **clusterInterfacesUuid** | **optional.String**| Filter by cluster_interfaces.uuid | 
 **clusterInterfacesIpAddress** | **optional.String**| Filter by cluster_interfaces.ip.address | 
 **managementInterfacesName** | **optional.String**| Filter by management_interfaces.name | 
 **managementInterfacesUuid** | **optional.String**| Filter by management_interfaces.uuid | 
 **managementInterfacesIpAddress** | **optional.String**| Filter by management_interfaces.ip.address | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**NodeResponse**](node_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PeersCollectionGet**
> ClusterPeerResponse PeersCollectionGet(ctx, optional)


Retrieve the collection of cluster peers. ### Learn more * [`DOC /cluster/peers`](#docs-cluster-cluster_peers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeersCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeersCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ClusterPeerResponse**](cluster_peer_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleCollectionGet**
> ScheduleResponse ScheduleCollectionGet(ctx, optional)


Retrieves a schedule. ### Learn more * [`DOC /cluster/schedules`](#docs-cluster-cluster_schedules)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduleCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by uuid | 
 **cronDays** | **optional.Int32**| Filter by cron.days | 
 **cronHours** | **optional.Int32**| Filter by cron.hours | 
 **cronMonths** | **optional.Int32**| Filter by cron.months | 
 **cronMinutes** | **optional.Int32**| Filter by cron.minutes | 
 **cronWeekdays** | **optional.Int32**| Filter by cron.weekdays | 
 **clusterName** | **optional.String**| Filter by cluster.name | 
 **clusterUuid** | **optional.String**| Filter by cluster.uuid | 
 **interval** | **optional.String**| Filter by interval | 
 **type_** | **optional.String**| Filter by type | 
 **name** | **optional.String**| Filter by name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**ScheduleResponse**](schedule_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleCreate**
> ScheduleCreate(ctx, optional)


Create a schedule. ### Required Fields * name - Name of the job schedule It is required to provide a minutes field for a cron schedule. An interval field is required for an interval schedule. You must not provide both a cron field and an interval field.  ### Learn more * [`DOC /cluster/schedules`](#docs-cluster-cluster_schedules)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduleCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of Schedule**](Schedule.md)| Info specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleDelete**
> ScheduleDelete(ctx, uuid)


Deletes a schedule. ### Learn more * [`DOC /cluster/schedules`](#docs-cluster-cluster_schedules)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Schedule UUID | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleGet**
> Schedule ScheduleGet(ctx, uuid, optional)


Retrieves a schedule. ### Learn more * [`DOC /cluster/schedules`](#docs-cluster-cluster_schedules)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Schedule UUID | 
 **optional** | ***ScheduleGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Schedule**](schedule.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleModify**
> ScheduleModify(ctx, uuid, optional)


Updates a schedule. Note that you cannot modify a cron field of an interval schedule, or the interval field of a cron schedule. ### Learn more * [`DOC /cluster/schedules`](#docs-cluster-cluster_schedules)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Schedule UUID | 
 **optional** | ***ScheduleModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleModifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of Schedule**](Schedule.md)| Info specification | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwareGet**
> SoftwareReference SoftwareGet(ctx, optional)


Retrieves the software profile of a cluster. ### Related ONTAP commands * `cluster image show` * `cluster image show-update-progress` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]

### Return type

[**SoftwareReference**](software_reference.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwareHistoryCollectionGet**
> SoftwareHistoryResponse SoftwareHistoryCollectionGet(ctx, optional)


Retrieves the history details for software installation requests. ### Related ONTAP commands * `cluster image show-update-history` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareHistoryCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareHistoryCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SoftwareHistoryResponse**](software_history_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwareModify**
> JobLinkResponse SoftwareModify(ctx, optional)


Upgrades the cluster software version. Setting `version`  triggers the installation of the package to start. To validate the package for installation but not perform the installation, use the `validate_only` field on request. Important note:   * Setting 'version' triggers the package installation.   * To validate the package for installation but not perform the installation, use the validate_only field on the request. ### Required properties * `version` - Software version to be installed on the cluster ### Recommended optional parameters * `validate_only` - Required to validate a software package before an upgrade * `skip_warnings` - Used to skip validation warnings when starting a software upgrade * `action` - Used to pause, resume, or cancel an ongoing software upgrade ### Related ONTAP commands * `cluster image validate` * `cluster image update` * `cluster image pause-update` * `cluster image resume-update` * `cluster image cancel-update` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwareModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwareModifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateOnly** | **optional.Bool**| Validate the operation and its parameters, without actually performing the operation. | 
 **skipWarnings** | **optional.Bool**| Ignore warnings and proceed with the install. | 
 **action** | **optional.String**| Requests an upgrade to pause, resume, or cancel. Note that not all upgrades support these actions. An upgrade can only be resumed if it is in the paused state. When a request to cancel an upgrade is successful, the upgrade state changes to either success or failure.  | 
 **info** | [**optional.Interface of SoftwareReference**](SoftwareReference.md)|  | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwarePackageCreate**
> JobLinkResponse SoftwarePackageCreate(ctx, optional)


Downloads a software package from the server. ### Required properties * `url` - URL location of the software package ### Recommended optional parameters * `username` - Username of HTTPS/FTP server * `password` - Password of HTTPS/FTP server ### Related ONTAP commands * `cluster image package get` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwarePackageCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwarePackageCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]
 **info** | [**optional.Interface of SoftwarePackageDownload**](SoftwarePackageDownload.md)| Info specification | 

### Return type

[**JobLinkResponse**](job_link_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwarePackageDelete**
> SoftwarePackageDelete(ctx, version)


Deletes a software package from the cluster. The delete operation fails if the package is currently installed. ### Related ONTAP commands * `cluster image package delete` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwarePackageGet**
> SoftwarePackage SoftwarePackageGet(ctx, version, optional)


Retrieves the software package information. ### Related ONTAP commands * `cluster image package show-repository` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***SoftwarePackageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwarePackageGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**SoftwarePackage**](software_package.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SoftwarePackagesCollectionGet**
> SoftwarePackageResponse SoftwarePackagesCollectionGet(ctx, optional)


Retrieves the software packages for a cluster. ### Related ONTAP commands * `cluster image package show-repository` ### Learn more * [`DOC /cluster/software`](#docs-cluster-cluster_software) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoftwarePackagesCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoftwarePackagesCollectionGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SoftwarePackageResponse**](software_package_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

 - **Content-Type**: application/json, application/hal+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

