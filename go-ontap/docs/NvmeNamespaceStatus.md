# NvmeNamespaceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerState** | **string** | The state of the volume and aggregate that contain the NVMe namespace. Namespaces are only available when their containers are available.  | [optional] [default to null]
**Mapped** | **bool** | Reports if the NVMe namespace is mapped to an NVMe subsystem.&lt;br/&gt; There is an added cost to retrieving this property&#39;s value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [default to null]
**ReadOnly** | **bool** | Reports if the NVMe namespace allows only read access.  | [optional] [default to null]
**State** | **string** | The state of the NVMe namespace. Normal states for a namespace are _online_ and _offline_. Other states indicate errors.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


