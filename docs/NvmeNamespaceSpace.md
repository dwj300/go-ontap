# NvmeNamespaceSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSize** | **int32** | The size of blocks in the namespace in bytes.&lt;br/&gt; Valid in POST when creating an NVMe namespace that is not a clone of another. Disallowed in POST when creating a namespace clone.  Valid in POST.  | [optional] 
**Guarantee** | [**NvmeNamespaceSpaceGuarantee**](nvme_namespace_space_guarantee.md) |  | [optional] 
**Size** | **int32** | The total provisioned size of the NVMe namespace.&lt;br/&gt; NVMe namespaces do not support resize.&lt;br/&gt; For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.  | [optional] 
**Used** | **int32** | The amount of space consumed by the main data stream of the NVMe namespace.&lt;br/&gt; This value is the total space consumed in the volume by the NVMe namespace, including filesystem overhead, but excluding prefix and suffix streams. Due to internal filesystem overhead and the many ways NVMe filesystems and applications utilize blocks within a namespace, this value does not necessarily reflect actual consumption/availability from the perspective of the filesystem or application. Without specific knowledge of how the namespace blocks are utilized outside of ONTAP, this property should not be used and an indicator for an out-of-space condition.&lt;br/&gt; For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


