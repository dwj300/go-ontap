# NvmeSubsystemMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Anagrpid** | **string** | The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.&lt;br/&gt; The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;.&lt;br/&gt; There is an added cost to retrieving this property&#39;s value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [default to null]
**Namespace** | [***NvmeSubsystemMapNamespace**](nvme_subsystem_map_namespace.md) |  | [optional] [default to null]
**Nsid** | **string** | The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.&lt;br/&gt; The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;.  | [optional] [default to null]
**Subsystem** | [***NvmeSubsystemMapSubsystem**](nvme_subsystem_map_subsystem.md) |  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


