# Svm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Aggregates** | [**[]DiskAggregates**](disk_aggregates.md) | List of allowed aggregates for SVM volumes. An administrator is allowed to create volumes on these aggregates. | [optional] 
**Cifs** | [**SvmCifs**](svm_cifs.md) |  | [optional] 
**Comment** | **string** | Comment | [optional] 
**Dns** | [**SvmDns**](svm_dns.md) |  | [optional] 
**Fcp** | [**SvmFcp**](svm_fcp.md) |  | [optional] 
**IpInterfaces** | [**[]IpInterfaceSvm**](ip_interface_svm.md) | IP interfaces for the SVM | [optional] 
**Ipspace** | [**SvmIpspace**](svm_ipspace.md) |  | [optional] 
**Iscsi** | [**SvmIscsi**](svm_iscsi.md) |  | [optional] 
**Language** | **string** | Default volume language code. UTF-8 encoded languages are valid in POST or PATCH. Non UTF-8 language encodings are for backward compatibility and are not valid input for POST and PATCH requests. | [optional] 
**Ldap** | [**SvmLdap**](svm_ldap.md) |  | [optional] 
**Name** | **string** | The name of the SVM.  | [optional] 
**Nfs** | [**SvmNfs**](svm_nfs.md) |  | [optional] 
**Nis** | [**SvmNis**](svm_nis.md) |  | [optional] 
**Nsswitch** | [**SvmNsswitch**](svm_nsswitch.md) |  | [optional] 
**Nvme** | [**SvmNvme**](svm_nvme.md) |  | [optional] 
**Routes** | [**[]NetworkRouteForSvm**](network_route_for_svm.md) | Optional array of routes for the SVM | [optional] 
**SnapshotPolicy** | [**SvmSnapshotPolicy**](svm_snapshot_policy.md) |  | [optional] 
**State** | **string** | SVM State | [optional] 
**Subtype** | **string** | SVM subtype. The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster. A POST request with sync_destination as SVM subtype is invalid. | [optional] 
**Uuid** | **string** | The unique identifier of the SVM.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


