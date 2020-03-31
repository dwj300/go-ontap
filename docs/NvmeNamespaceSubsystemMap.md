# NvmeNamespaceSubsystemMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Anagrpid** | **string** | The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.&lt;br/&gt; The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;.  | [optional] [readonly] 
**Nsid** | **string** | The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.&lt;br/&gt; The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;.  | [optional] [readonly] 
**Subsystem** | [**NvmeNamespaceSubsystemMapSubsystem**](nvme_namespace_subsystem_map_subsystem.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


