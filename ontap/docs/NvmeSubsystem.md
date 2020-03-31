# NvmeSubsystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Comment** | **string** | A configurable comment for the NVMe subsystem. Optional in POST and PATCH.  | [optional] 
**Hosts** | [**[]NvmeSubsystemHosts**](nvme_subsystem_hosts.md) |  | [optional] [readonly] 
**IoQueue** | [**NvmeSubsystemIoQueue**](nvme_subsystem_io_queue.md) |  | [optional] 
**Name** | **string** | The name of the NVMe subsystem. Once created, an NVMe subsystem cannot be renamed. Required in POST.  | [optional] 
**OsType** | **string** | The host operating system of the NVMe subsystem&#39;s hosts. Required in POST.  | [optional] 
**SerialNumber** | **string** | The serial number of the NVMe subsystem.  | [optional] [readonly] 
**SubsystemMaps** | [**[]NvmeSubsystemSubsystemMaps**](nvme_subsystem_subsystem_maps.md) | The NVMe namespaces mapped to the NVMe subsystem.&lt;br/&gt; There is an added cost to retrieving property values for &#x60;subsystem_maps&#x60;. They are not populated for either a collection GET or an instance GET unless explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [readonly] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**TargetNqn** | **string** | The NVMe qualified name (NQN) used to identify the NVMe storage target.  | [optional] [readonly] 
**Uuid** | **string** | The unique identifier of the NVMe subsystem.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


