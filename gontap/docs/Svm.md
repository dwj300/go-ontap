# Svm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Aggregates** | [**[]DiskAggregates**](disk_aggregates.md) | List of allowed aggregates for SVM volumes. An administrator is allowed to create volumes on these aggregates. | [optional] [default to null]
**Cifs** | [***SvmCifs**](svm_cifs.md) |  | [optional] [default to null]
**Comment** | **string** | Comment | [optional] [default to null]
**Dns** | [***SvmDns**](svm_dns.md) |  | [optional] [default to null]
**Fcp** | [***SvmFcp**](svm_fcp.md) |  | [optional] [default to null]
**IpInterfaces** | [**[]IpInterfaceSvm**](ip_interface_svm.md) | IP interfaces for the SVM | [optional] [default to null]
**Ipspace** | [***SvmIpspace**](svm_ipspace.md) |  | [optional] [default to null]
**Iscsi** | [***SvmIscsi**](svm_iscsi.md) |  | [optional] [default to null]
**Language** | **string** | Default volume language code. UTF-8 encoded languages are valid in POST or PATCH. Non UTF-8 language encodings are for backward compatibility and are not valid input for POST and PATCH requests. | [optional] [default to null]
**Ldap** | [***SvmLdap**](svm_ldap.md) |  | [optional] [default to null]
**Name** | **string** | The name of the SVM.  | [optional] [default to null]
**Nfs** | [***SvmNfs**](svm_nfs.md) |  | [optional] [default to null]
**Nis** | [***SvmNis**](svm_nis.md) |  | [optional] [default to null]
**Nsswitch** | [***SvmNsswitch**](svm_nsswitch.md) |  | [optional] [default to null]
**Nvme** | [***SvmNvme**](svm_nvme.md) |  | [optional] [default to null]
**Routes** | [**[]NetworkRouteForSvm**](network_route_for_svm.md) | Optional array of routes for the SVM | [optional] [default to null]
**SnapshotPolicy** | [***SvmSnapshotPolicy**](svm_snapshot_policy.md) |  | [optional] [default to null]
**State** | **string** | SVM State | [optional] [default to null]
**Subtype** | **string** | SVM subtype. The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster. A POST request with sync_destination as SVM subtype is invalid. | [optional] [default to null]
**Uuid** | **string** | The unique identifier of the SVM.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


