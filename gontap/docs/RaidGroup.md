# RaidGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheTier** | **bool** | RAID group is a cache tier | [optional] [default to null]
**Degraded** | **bool** | RAID group is degraded. A RAID group is degraded when at least one disk from that group has failed or is offline. | [optional] [default to null]
**Disks** | [**[]RaidGroupDisk**](raid_group_disk.md) |  | [optional] [default to null]
**Name** | **string** | RAID group name | [optional] [default to null]
**RecomputingParity** | [***RaidGroupRecomputingParity**](raid_group_recomputing_parity.md) |  | [optional] [default to null]
**Reconstruct** | [***RaidGroupReconstruct**](raid_group_reconstruct.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


