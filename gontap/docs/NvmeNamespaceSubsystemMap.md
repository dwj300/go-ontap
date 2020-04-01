# NvmeNamespaceSubsystemMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Anagrpid** | **string** | The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.&lt;br/&gt; The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;.  | [optional] [default to null]
**Nsid** | **string** | The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.&lt;br/&gt; The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;.  | [optional] [default to null]
**Subsystem** | [***NvmeNamespaceSubsystemMapSubsystem**](nvme_namespace_subsystem_map_subsystem.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


