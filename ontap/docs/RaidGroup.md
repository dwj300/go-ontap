# RaidGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheTier** | **bool** | RAID group is a cache tier | [optional] [readonly] 
**Degraded** | **bool** | RAID group is degraded. A RAID group is degraded when at least one disk from that group has failed or is offline. | [optional] [readonly] 
**Disks** | [**[]RaidGroupDisk**](raid_group_disk.md) |  | [optional] [readonly] 
**Name** | **string** | RAID group name | [optional] [readonly] 
**RecomputingParity** | [**RaidGroupRecomputingParity**](raid_group_recomputing_parity.md) |  | [optional] 
**Reconstruct** | [**RaidGroupReconstruct**](raid_group_reconstruct.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


