# \SvmApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SvmCollectionGet**](SvmApi.md#SvmCollectionGet) | **Get** /svm/svms | 
[**SvmCreate**](SvmApi.md#SvmCreate) | **Post** /svm/svms | 
[**SvmDelete**](SvmApi.md#SvmDelete) | **Delete** /svm/svms/{uuid} | 
[**SvmGet**](SvmApi.md#SvmGet) | **Get** /svm/svms/{uuid} | 
[**SvmPatchModify**](SvmApi.md#SvmPatchModify) | **Patch** /svm/svms/{uuid} | 
[**SvmPeerCollectionGet**](SvmApi.md#SvmPeerCollectionGet) | **Get** /svm/peers | 
[**SvmPeerCreate**](SvmApi.md#SvmPeerCreate) | **Post** /svm/peers | 
[**SvmPeerDelete**](SvmApi.md#SvmPeerDelete) | **Delete** /svm/peers/{peer.uuid} | 
[**SvmPeerInstanceGet**](SvmApi.md#SvmPeerInstanceGet) | **Get** /svm/peers/{peer.uuid} | 
[**SvmPeerModify**](SvmApi.md#SvmPeerModify) | **Patch** /svm/peers/{peer.uuid} | 
[**SvmPeerPermissionCollectionGet**](SvmApi.md#SvmPeerPermissionCollectionGet) | **Get** /svm/peer-permissions | 
[**SvmPeerPermissionDelete**](SvmApi.md#SvmPeerPermissionDelete) | **Delete** /svm/peer-permissions/{cluster.uuid}/{svm.uuid} | 
[**SvmPeerPermissionInstanceGet**](SvmApi.md#SvmPeerPermissionInstanceGet) | **Get** /svm/peer-permissions/{cluster.uuid}/{svm.uuid} | 
[**SvmPeerPermissionModify**](SvmApi.md#SvmPeerPermissionModify) | **Patch** /svm/peer-permissions/{cluster.uuid}/{svm.uuid} | 
[**SvmPeerPermissionPost**](SvmApi.md#SvmPeerPermissionPost) | **Post** /svm/peer-permissions | 



## SvmCollectionGet

> SvmResponse SvmCollectionGet(ctx, optional)



Retrieves a list of SVMs and individual SVM properties. This includes protocol configurations such as CIFS and NFS, export policies, name service configurations, and network services. ### Important notes * The SVM object includes a large set of fields and can be expensive to retrieve. Use this API to list the collection of SVMs, and to retrieve only the full details of individual SVMs as needed. * It is not recommended to create or delete more than five SVMs in parallel. * REST APIs only expose a data SVM as an SVM. ### Related ONTAP commands * `vserver show` ### Examples 1. Retrieves a list of SVMs in the cluster sorted by name     <br/>     ```     GET \"/api/svm/svms?order_by=name\"     ```     <br/> 2. Retrieves a list of SVMs in the cluster that have the NFS protocol enabled     <br/>     ```     GET \"/api/svm/svms?nfs.enabled=true\"     ```     <br/> 3. Retrieves a list of SVMs in the cluster that have the CIFS protocol enabled     <br/>     ```     GET \"/api/svm/svms?cifs.enabled=true\"     ```     <br/> ### Learn more * [`DOC /svm/svms`](#docs-svm-svm_svms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SvmCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nisEnabled** | **optional.Bool**| Filter by nis.enabled | 
 **nisServers** | **optional.String**| Filter by nis.servers | 
 **nisDomain** | **optional.String**| Filter by nis.domain | 
 **nvmeEnabled** | **optional.Bool**| Filter by nvme.enabled | 
 **language** | **optional.String**| Filter by language | 
 **nfsEnabled** | **optional.Bool**| Filter by nfs.enabled | 
 **comment** | **optional.String**| Filter by comment | 
 **aggregatesName** | **optional.String**| Filter by aggregates.name | 
 **aggregatesUuid** | **optional.String**| Filter by aggregates.uuid | 
 **subtype** | **optional.String**| Filter by subtype | 
 **dnsServers** | **optional.String**| Filter by dns.servers | 
 **dnsDomains** | **optional.String**| Filter by dns.domains | 
 **fcpEnabled** | **optional.Bool**| Filter by fcp.enabled | 
 **iscsiEnabled** | **optional.Bool**| Filter by iscsi.enabled | 
 **name** | **optional.String**| Filter by name | 
 **ipspaceUuid** | **optional.String**| Filter by ipspace.uuid | 
 **ipspaceName** | **optional.String**| Filter by ipspace.name | 
 **ldapBaseDn** | **optional.String**| Filter by ldap.base_dn | 
 **ldapServers** | **optional.String**| Filter by ldap.servers | 
 **ldapEnabled** | **optional.Bool**| Filter by ldap.enabled | 
 **ldapBindDn** | **optional.String**| Filter by ldap.bind_dn | 
 **ldapAdDomain** | **optional.String**| Filter by ldap.ad_domain | 
 **uuid** | **optional.String**| Filter by uuid | 
 **cifsName** | **optional.String**| Filter by cifs.name | 
 **cifsAdDomainFqdn** | **optional.String**| Filter by cifs.ad_domain.fqdn | 
 **cifsAdDomainOrganizationalUnit** | **optional.String**| Filter by cifs.ad_domain.organizational_unit | 
 **cifsEnabled** | **optional.Bool**| Filter by cifs.enabled | 
 **nsswitchNetgroup** | **optional.String**| Filter by nsswitch.netgroup | 
 **nsswitchGroup** | **optional.String**| Filter by nsswitch.group | 
 **nsswitchHosts** | **optional.String**| Filter by nsswitch.hosts | 
 **nsswitchNamemap** | **optional.String**| Filter by nsswitch.namemap | 
 **nsswitchPasswd** | **optional.String**| Filter by nsswitch.passwd | 
 **ipInterfacesIpAddress** | **optional.String**| Filter by ip_interfaces.ip.address | 
 **ipInterfacesName** | **optional.String**| Filter by ip_interfaces.name | 
 **ipInterfacesUuid** | **optional.String**| Filter by ip_interfaces.uuid | 
 **snapshotPolicyUuid** | **optional.String**| Filter by snapshot_policy.uuid | 
 **snapshotPolicyName** | **optional.String**| Filter by snapshot_policy.name | 
 **state** | **optional.String**| Filter by state | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SvmResponse**](svm_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmCreate

> JobLinkResponse SvmCreate(ctx, optional)



Creates and provisions an SVM. If no IPspace is provided, then the SVM is created on the `Default` IPspace. The number of parallel SVMs that can be created should not be greater than 5. When the sixth SVM POST request is issued, the error message \"Maximum allowed SVM jobs exceeded. Wait for the existing SVM jobs to complete and try again.\" will be returned. ### Required properties * `name` - Name of the SVM to be created. ### Recommended optional properties * `ipspace.name` or `ipspace.uuid` - IPspace of the SVM   * `ip_interfaces` - If provided, the following fields are required:   * `ip_interfaces.name` - Name of the interface   * `ip_interfaces.ip.address` - IP address   * `ip_interfaces.ip.netmask` - Netmask length or IP address   * `ip_interfaces.location.broadcast_domain.uuid` or `ip_interfaces.location.broadcast_domain.name` - Broadcast domain name or UUID belonging to the same IPspace of the SVM. * `routes` - If provided, the following field is required:   * `routes.gateway` - Gateway IP address * `cifs` - If provided, interfaces, routes and DNS must be provided. The following fields are also required:   * `cifs.name` - Name of the CIFS server to be created for the SVM.   * `cifs.ad_domain.fqdn` - Fully qualified domain name   * `cifs.ad_domain.user` - Administrator username   * `cifs.ad_domain.password` - User password * `ldap` - If provided, the following fields are required:   * `ldap.servers` or `ldap.ad_domain` - LDAP server list or Active directory domain   * `ldap.bind_dn` - Bind DN   * `ldap.base_dn` - Base DN * `nis` - If provided, the following fields are required:   * `nis.servers` - NIS servers   * `nis.domain` - NIS domain * `dns` - If provided, the following fields are required:   * `dns.servers` - Name servers   * `dns.domains` - Domains ### Default property values If not specified in POST, the following default property values are assigned: * `language` - _C.UTF-8_ * `ipspace.name` - _Default_ * `snapshot_policy.name` - _Default_ * `subtype` - _Default_ ( _sync-source_ if MetroCluster configuration ) ### Related ONTAP commands * `vserver create` * `vserver add-aggregates` * `network interface create` * `network route create` * `vserver services name-service dns create` * `vserver nfs create` * `vserver services name-service ldap client create` * `vserver cifs create` * `vserver services name-service nis-domain create` * `vserver iscsi create` * `vserver nvme create` * `vserver fcp create` * `vserver services name-service ns-switch create` ### Examples 1. Creates an SVM with default \"snapshot_policy\"     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"snapshot_policy\":{\"name\":\"default\"}}'     ```     <br/> 2. Creates an SVM and configures NFS, ISCSI and FCP     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"nfs\":{\"enabled\":\"true\"}, \"fcp\":{\"enabled\":\"true\"}, \"iscsi\":{\"enabled\":\"true\"}}'     ```     <br/> 3. Creates an SVM and configures NVMe     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"nvme\":{\"enabled\":\"true\"}}'     ```     <br/> 4. Creates an SVM and configures LDAP     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"snapshot_policy\":{\"name\":\"default\"}, \"ldap\":{\"servers\":[\"10.140.101.1\",\"10.140.101.2\"], \"ad_domain\":\"abc.com\", \"base_dn\":\"dc=netapp,dc=com\", \"bind_dn\":\"dc=netapp,dc=com\"}}'     ```     <br/> 5. Creates an SVM and configures NIS     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"snapshot_policy\":{\"name\":\"default\"}, \"nis\":{\"enabled\":\"true\", \"domain\":\"def.com\",\"servers\":[\"10.224.223.130\", \"10.224.223.131\"]}}'     ```     <br/> 6. Creates an SVM and configures DNS     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"snapshot_policy\":{\"name\":\"default\"}, \"dns\":{\"domains\":[\"abc.com\",\"def.com\"], \"servers\":[\"10.224.223.130\", \"10.224.223.131\"]}}'     ```     <br/> 7. Creates an SVM and configures a LIF     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"ip_interfaces\": [{\"name\":\"lif1\", \"ip\":{\"address\":\"10.10.10.7\", \"netmask\": \"255.255.255.0\"}, \"location\":{\"broadcast_domain\":{\"name\":\"bd1\"}, \"home_node\":{\"name\":\"node1\"}}, \"service_policy\": \"default-management\"}]}'     ```     <br/> 8. Creates an SVM and configures a LIF with IPV6 address     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"ip_interfaces\": [{\"name\":\"lif2\", \"ip\":{\"address\":\"fd22:8b1e:b255:202:2a0:98ff:fe01:7d5b\", \"netmask\":\"24\"}, \"location\":{\"broadcast_domain\":{\"name\":\"bd1\"}, \"home_node\":{\"name\":\"node1\"}}, \"service_policy\": \"default-management\"}]}'     ```     <br/> 9. Creates an SVM and configures CIFS     <br/>     ```     POST \"/api/svm/svms\" '{\"name\":\"testVs\", \"cifs\":{\"name\":\"CIFDOC\", \"ad_domain\":{\"fqdn\":\"abc.def.com\", \"organizational_unit\":\"CN=Computers\", \"user\":\"cif_admin\", \"password\":\"abc123\"}}, \"ip_interfaces\":[{\"name\":\"lif1\", \"ip\":{\"address\":\"10.10.10.7\", \"netmask\": \"255.255.255.0\"}, \"location\":{\"broadcast_domain\":{\"name\":\"bd1\"}, \"home_node\":{\"name\":\"node1\"}}, \"service_policy\": \"default-management\"}],\"routes\": [{\"destination\": {\"address\": \"0.0.0.0\", \"netmask\": \"0\"}, \"gateway\": \"10.10.10.7\"}], \"dns\":{\"domains\":[\"abc.def.com\", \"def.com\"], \"servers\":[\"10.224.223.130\", \"10.224.223.131\"]}}'     ```     <br/> ### Learn more * [`DOC /svm/svms`](#docs-svm-svm_svms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SvmCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]
 **info** | [**optional.Interface of Svm**](Svm.md)| Info specification | 

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


## SvmDelete

> JobLinkResponse SvmDelete(ctx, uuid, optional)



Deletes an SVM. As a prerequisite, SVM objects must be deleted first. SnapMirror relations must be deleted and data volumes must be offline and deleted. The number of parallel SVMs that can be deleted should not be greater than 5. When the sixth SVM DELETE request is issued, the error message \"Maximum allowed SVM jobs exceeded. Wait for the existing SVM jobs to complete and try again.\" will be returned. ### Related ONTAP commands * `vserver delete` ### Example Deletes an individual SVM in the cluster.   <br/>   ```   DELETE \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\"   ```   <br/> ### Learn more * [`DOC /svm/svms`](#docs-svm-svm_svms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Filter by UUID | 
 **optional** | ***SvmDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmDeleteOpts struct


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


## SvmGet

> Svm SvmGet(ctx, uuid, optional)



Retrieves the properties for an individual SVM. This includes protocol configurations such as CIFS and NFS, export policies, name service configurations, and network services. ### Important note * The SVM object includes a large set of fields and can be expensive to retrieve. * REST APIs only expose a data SVM as an SVM. ### Example     Retrieves an individual SVM in the cluster     <br/>     ```     GET \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\"     ```     <br/>  ### Learn more * [`DOC /svm/svms`](#docs-svm-svm_svms)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Filter by UUID | 
 **optional** | ***SvmGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 

### Return type

[**Svm**](svm.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPatchModify

> JobLinkResponse SvmPatchModify(ctx, uuid, optional)



Updates one or more of the following properties of an individual SVM: SVM name, SVM default volume language code, SVM comment, and SVM state. ### Related ONTAP commands * `vserver modify` * `vserver rename` * `vserver start` * `vserver stop` ### Examples 1.  Stops an SVM and updates the \"comment\" field for an individual SVM     <br/>     ```     PATCH \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\" '{\"state\":\"stopped\", \"comment\":\"This SVM is stopped.\"}'     ```     <br/> 2.  Starts an SVM and updates the \"comment\" field for an individual SVM     <br/>     ```     PATCH \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\" '{\"state\":\"running\", \"comment\":\"This SVM is running.\"}'     ```     <br/> 3.  Updates the \"language\" field for an individual SVM     <br/>     ```     PATCH \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\" '{\"language\":\"en.UTF-8\"}'     ```     <br/> 4.  Updates the \"name\" field for an SVM or renames the SVM     <br/>     ```     PATCH \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\" '{\"name\":\"svm_new\"}'     ```     <br/> 5.  Updates the aggregates for an individual SVM     <br/>     ```     PATCH \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\" '{\"aggregates\":{\"name\":[\"aggr1\",\"aggr2\",\"aggr3\"]}}'     ```     <br/> 6.  Updates the Snapshot copy policy for an individual SVM     <br/>     ```     PATCH \"/api/svm/svms/f16f0935-5281-11e8-b94d-005056b46485\" '{\"snapshot_policy\":{\"name\":\"custom1\"}}'     ```     <br/> ### Learn more * [`DOC /svm/svms`](#docs-svm-svm_svms) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Filter by UUID | 
 **optional** | ***SvmPatchModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPatchModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202. | [default to 0]
 **info** | [**optional.Interface of Svm**](Svm.md)| Info specification | 

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


## SvmPeerCollectionGet

> SvmPeerResponse SvmPeerCollectionGet(ctx, optional)



Retrieves the list of SVM peer relationships. ### Related ONTAP commands * `vserver peer show` ### Examples The following examples show how to retrieve a collection of SVM peer relationships based on a query. 1. Retrieves a list of SVM peers of a specific local SVM    <br/>    ```    GET \"/api/svm/peers/?svm.name=VS1\"    ```    <br/> 2. Retrieves a list of SVM peers of a specific cluster peer    <br/>    ```    GET \"/api/svm/peers/?peer.cluster.name=cluster2\"    ```    <br/> ### Learn more * [`DOC /svm/peers`](#docs-svm-svm_peers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SvmPeerCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**| Filter by uuid | 
 **peerClusterUuid** | **optional.String**| Filter by peer.cluster.uuid | 
 **peerClusterName** | **optional.String**| Filter by peer.cluster.name | 
 **peerSvmUuid** | **optional.String**| Filter by peer.svm.uuid | 
 **peerSvmName** | **optional.String**| Filter by peer.svm.name | 
 **state** | **optional.String**| Filter by state | 
 **applications** | **optional.String**| Filter by applications | 
 **name** | **optional.String**| Filter by name | 
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SvmPeerResponse**](svm_peer_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPeerCreate

> SvmPeer SvmPeerCreate(ctx, optional)



Creates a new SVM peer relationship. ### Important notes   * The create request accepts peer SVM name as input instead of peer SVM UUID as the local cluster cannot validate peer SVM based on UUID.   * The input parameter `name` refers to the local name of the peer SVM. The `peer cluster name` parameter is optional for creating intracluster SVM peer relationships. ### Required properties * `svm.name` or `svm.uuid` - SVM name or SVM UUID * `peer.svm.name` or `peer.svm.uuid` - Peer SVM name or Peer SVM UUID * `peer.cluster.name` or `peer.cluster.uuid` - Peer cluster name or peer cluster UUID * `applications` - Peering applications ### Related ONTAP commands * `vserver peer create` ### Example Creates a new SVM peer relationship. <br/> ``` POST \"/api/svm/peers\" '{\"svm\":{\"name\":\"vs1\", \"peer.cluster.name\":\"cluster2\", \"peer.svm.name\":\"VS1\", \"applications\":[\"snapmirror\"]}' ``` <br/> ### Learn more * [`DOC /svm/peers`](#docs-svm-svm_peers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SvmPeerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SvmPeer**](SvmPeer.md)| Info specification | 

### Return type

[**SvmPeer**](svm_peer.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPeerDelete

> SvmPeerDelete(ctx, peerUuid)



Deletes the SVM peer relationship. ### Related ONTAP commands * `vserver peer delete` ### Example Deletes an SVM peer relationship. <br/> ``` DELETE \"/api/svm/peers/d3268a74-ee76-11e8-a9bb-005056ac6dc9\" ``` <br/> ### Learn more * [`DOC /svm/peers`](#docs-svm-svm_peers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerUuid** | **string**| SVM peer relationship UUID | 

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


## SvmPeerInstanceGet

> SvmPeer SvmPeerInstanceGet(ctx, peerUuid, optional)



Retrieves the SVM peer relationship instance. ### Related ONTAP commands * `vserver peer show` ### Example Retrieves the parameters of an SVM peer relationship. <br/> ``` GET \"/api/svm/peers/d3268a74-ee76-11e8-a9bb-005056ac6dc9\" ``` <br/> ### Learn more * [`DOC /svm/peers`](#docs-svm-svm_peers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerUuid** | **string**| SVM peer relation UUID | 
 **optional** | ***SvmPeerInstanceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerInstanceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]

### Return type

[**SvmPeer**](svm_peer.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPeerModify

> SvmPeerModify(ctx, peerUuid, optional)



Updates the SVM peer relationship. ### Related ONTAP commands * `vserver peer modify` ### Examples The following examples show how to update an SVM peer relationship. The input parameter 'name' refers to the local name of the peer SVM. <br/> 1. Accepts an SVM peer relationship    <br/>    ```    PATCH \"/api/svm/peers/d3268a74-ee76-11e8-a9bb-005056ac6dc9\" '{\"state\":\"peered\"}'    ```    <br/> 2. Updates the local name of an SVM peer relationship    <br/>    ```    PATCH \"/api/svm/peers/d3268a74-ee76-11e8-a9bb-005056ac6dc9\" '{\"name\":\"vs2\"}'    ```    <br/> ### Learn more * [`DOC /svm/peers`](#docs-svm-svm_peers) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerUuid** | **string**| SVM peer relationship UUID | 
 **optional** | ***SvmPeerModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**optional.Interface of SvmPeer**](SvmPeer.md)| Info specification | 

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


## SvmPeerPermissionCollectionGet

> SvmPeerPermissionResponse SvmPeerPermissionCollectionGet(ctx, optional)



Retrieves the list of SVM peer permissions. ### Related ONTAP commands * `vserver peer permission show` ### Examples The following examples show how to retrieve a collection of SVM peer permissions based on a query. <br/> 1. Retrieves a list of SVM peer permissions of a specific local SVM    <br/>    ```    GET \"/api/svm/peer-permissions/?svm.name=VS1\"    ```    <br/> 2. Retrieves a list of SVM peer permissions of a specific cluster peer    <br/>    ```    GET \"/api/svm/peer-permissions/?cluster_peer.name=cluster2\"    ```    <br/> ### Learn more * [`DOC /svm/peer-permissions`](#docs-svm-svm_peer-permissions) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SvmPeerPermissionCollectionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerPermissionCollectionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svmUuid** | **optional.String**| Filter by svm.uuid | 
 **svmName** | **optional.String**| Filter by svm.name | 
 **applications** | **optional.String**| Filter by applications | 
 **clusterPeerUuid** | **optional.String**| Filter by cluster_peer.uuid | 
 **clusterPeerName** | **optional.String**| Filter by cluster_peer.name | 
 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **maxRecords** | **optional.Int32**| Limit the number of records returned. | 
 **returnRecords** | **optional.Bool**| The default is true for GET calls.  When set to false, only the number of records is returned. | [default to true]
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]
 **orderBy** | [**optional.Interface of []string**](string.md)| Order results by specified fields and optional [asc|desc] direction. Default direction is &#39;asc&#39; for ascending. | 

### Return type

[**SvmPeerPermissionResponse**](svm_peer_permission_response.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPeerPermissionDelete

> SvmPeerPermissionDelete(ctx, clusterUuid, svmUuid)



Deletes the SVM peer permissions. ### Related ONTAP commands * `verver peer permission delete` ### Example Deletes an SVM peer permission. <br/> ``` DELETE \"/api/svm/peer-permissions/d3268a74-ee76-11e8-a9bb-005056ac6dc9/8f467b93-f2f1-11e8-9027-005056ac81fc\" ``` <br/> ### Learn more * [`DOC /svm/peer-permissions`](#docs-svm-svm_peer-permissions) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string**| Peer cluster UUID | 
**svmUuid** | **string**| SVM UUID | 

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


## SvmPeerPermissionInstanceGet

> SvmPeerPermission SvmPeerPermissionInstanceGet(ctx, clusterUuid, svmUuid, optional)



Retrieves the SVM peer permission instance. ### Related ONTAP commands * `vserver peer permission show` ### Example The following example shows how to retrieve the parameters for an SVM peer permission. <br/> ``` GET \"/api/svm/peer-permissions/d3268a74-ee76-11e8-a9bb-005056ac6dc9/8f467b93-f2f1-11e8-9027-005056ac81fc\" ``` <br/> ### Learn more * [`DOC /svm/peer-permissions`](#docs-svm-svm_peer-permissions) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string**| Peer cluster UUID | 
**svmUuid** | **string**| SVM UUID | 
 **optional** | ***SvmPeerPermissionInstanceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerPermissionInstanceGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| Specify the fields to return. | 
 **returnTimeout** | **optional.Int32**| The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached. | [default to 15]

### Return type

[**SvmPeerPermission**](svm_peer_permission.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPeerPermissionModify

> SvmPeerPermission SvmPeerPermissionModify(ctx, clusterUuid, svmUuid, optional)



Updates the SVM peer permissions. ### Related ONTAP commands * `vserver peer permission modify` ### Example Updates an SVM peer permission. <br/> ``` PATCH \"/api/svm/peer-permissions/d3268a74-ee76-11e8-a9bb-005056ac6dc9/8f467b93-f2f1-11e8-9027-005056ac81fc\" '{\"applications\":[\"flexcache\"]}' ``` <br/> ### Learn more * [`DOC /svm/peer-permissions`](#docs-svm-svm_peer-permissions) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string**| Peer cluster UUID | 
**svmUuid** | **string**| SVM UUID | 
 **optional** | ***SvmPeerPermissionModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerPermissionModifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **info** | [**optional.Interface of SvmPeerPermission**](SvmPeerPermission.md)| Info specification | 

### Return type

[**SvmPeerPermission**](svm_peer_permission.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SvmPeerPermissionPost

> SvmPeerPermission SvmPeerPermissionPost(ctx, optional)



Creates an SVM peer permission. ### Required properties * `svm.name` or `svm.uuid` - SVM name * `cluster_peer.uuid` or `cluster_peer.name` - Peer cluster name or peer cluster UUID * `applications` - Peering applications ### Related ONTAP commands * `vserver peer permission create` ### Examples The following examples show how to create SVM peer permissions. <br/> 1. Creates an SVM peer permission entry with the local SVM and cluster peer names    <br/>    ```    POST \"/api/svm/peer-permissions\" '{\"cluster_peer\":{\"name\":\"cluster2\"}, \"svm\":{\"name\":\"VS1\"}, \"applications\":[\"snapmirror\"]}'    ```    <br/> 2. Creates an SVM peer permission entry with the local SVM and cluster peer UUID    <br/>    ```    POST \"/api/svm/peer-permissions\" '{\"cluster_peer\":{\"uuid\":\"d3268a74-ee76-11e8-a9bb-005056ac6dc9\"}, \"svm\":{\"uuid\":\"8f467b93-f2f1-11e8-9027-005056ac81fc\"}, \"applications\":[\"snapmirror\"]}'    ```    <br/> 3. Creates an SVM peer permission entry with all SVMs and the cluster peer name    <br/>    ```    POST \"/api/svm/peer-permissions\" '{\"cluster_peer\":{\"name\":\"cluster2\"}, \"svm\":{\"name\":\"*\"}, \"applications\":[\"snapmirror\"]}'    ```    <br/> ### Learn more * [`DOC /svm/peer-permissions`](#docs-svm-svm_peer-permissions) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SvmPeerPermissionPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SvmPeerPermissionPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **info** | [**optional.Interface of SvmPeerPermission**](SvmPeerPermission.md)| Info specification | 

### Return type

[**SvmPeerPermission**](svm_peer_permission.md)

### Authorization

[simple](../README.md#simple)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

