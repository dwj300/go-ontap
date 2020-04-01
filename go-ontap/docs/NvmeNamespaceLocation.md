# NvmeNamespaceLocation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | The base name component of the NVMe namespace. Valid in POST.&lt;br/&gt; If properties &#x60;name&#x60; and &#x60;location.namespace&#x60; are specified in the same request, they must refer to the base name.&lt;br/&gt; NVMe namespaces do not support rename.  | [optional] [default to null]
**Qtree** | [***NvmeNamespaceLocationQtree**](nvme_namespace_location_qtree.md) |  | [optional] [default to null]
**Volume** | [***NvmeNamespaceLocationVolume**](nvme_namespace_location_volume.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


