# NvmeNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**AutoDelete** | **bool** | This property marks the NVMe namespace for auto deletion when the volume containing the namespace runs out of space. This is most commonly set on namespace clones.&lt;br/&gt; When set to _true_, the NVMe namespace becomes eligible for automatic deletion when the volume runs out of space. Auto deletion only occurs when the volume containing the namespace is also configured for auto deletion and free space in the volume decreases below a particular threshold.&lt;br/&gt; This property is optional in POST and PATCH. The default value for a new NVMe namespace is _false_.&lt;br/&gt; There is an added cost to retrieving this property&#39;s value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [default to false]
**Clone** | [**NvmeNamespaceClone**](nvme_namespace_clone.md) |  | [optional] 
**Comment** | **string** | A configurable comment available for use by the administrator. Valid in POST and PATCH.  | [optional] 
**Enabled** | **bool** | The enabled state of the NVMe namespace. Namespaces can be disabled to prevent access to the namespace. Certain error conditions also cause the namespace to become disabled. If the namespace is disabled, you can consult the &#x60;state&#x60; property to determine if the namespace is administratively disabled (_offline_) or has become disabled as a result of an error. A namespace in an error condition can be brought online by setting the &#x60;enabled&#x60; property to _true_ or brought administratively offline by setting the &#x60;enabled&#x60; property to _false_. Upon creation, an NVMe namespace is enabled by default. Valid in PATCH.  | [optional] 
**Location** | [**NvmeNamespaceLocation**](nvme_namespace_location.md) |  | [optional] 
**Name** | **string** | The fully qualified path name of the NVMe namespace composed of a \&quot;/vol\&quot; prefix, the volume name, the (optional) qtree name and base name of the namespace. Valid in POST.&lt;br/&gt; NVMe namespaces do not support rename, or movement between volumes.  | [optional] 
**OsType** | **string** | The operating system type of the NVMe namespace.&lt;br/&gt; Required in POST when creating an NVMe namespace that is not a clone of another. Disallowed in POST when creating a namespace clone.  | [optional] 
**Space** | [**NvmeNamespaceSpace**](nvme_namespace_space.md) |  | [optional] 
**Status** | [**NvmeNamespaceStatus**](nvme_namespace_status.md) |  | [optional] 
**SubsystemMap** | [**NvmeNamespaceSubsystemMap**](nvme_namespace_subsystem_map.md) |  | [optional] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Uuid** | **string** | The unique identifier of the NVMe namespace.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


