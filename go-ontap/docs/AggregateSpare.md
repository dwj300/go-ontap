# AggregateSpare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChecksumStyle** | **string** | The checksum type that has been assigned to the spares | [optional] [default to null]
**DiskClass** | **string** | Disk class of spares | [optional] [default to null]
**LayoutRequirements** | [**[]LayoutRequirement**](layout_requirement.md) | Available RAID protections and their restrictions | [optional] [default to null]
**Node** | [***AggregateSpareNode**](aggregate_spare_node.md) |  | [optional] [default to null]
**Size** | **int32** | Usable size of each spare in bytes | [optional] [default to null]
**SyncmirrorPool** | **string** | SyncMirror spare pool | [optional] [default to null]
**Usable** | **int32** | Total number of usable spares | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


