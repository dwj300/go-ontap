# NvmeSubsystemHost

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**IoQueue** | [***NvmeSubsystemHostIoQueue**](nvme_subsystem_host_io_queue.md) |  | [optional] [default to null]
**Nqn** | **string** | The NVMe qualified name (NQN) used to identify the NVMe storage target. Not allowed in POST when the &#x60;records&#x60; property is used.  | [optional] [default to null]
**Records** | [**[]NvmeSubsystemHostRecords**](nvme_subsystem_host_records.md) | An array of NVMe hosts specified to add multiple NVMe hosts to an NVMe subsystem in a single API call. Valid in POST only.  | [optional] [default to null]
**Subsystem** | [***NvmeSubsystemHostSubsystem**](nvme_subsystem_host_subsystem.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


