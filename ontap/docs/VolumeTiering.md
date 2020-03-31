# VolumeTiering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | Policy that determines whether the user data blocks of a volume in a FabricPool will be tiered to the cloud store when they become cold. FabricPool combines flash (performance tier) with a cloud store into a single aggregate. Temperature of a volume block increases if it is accessed frequently and decreases when it is not. Valid in POST or PATCH.&lt;br&gt;all &amp;dash; This policy allows tiering of both Snapshot copies and active file system user data to the cloud store as soon as possible by ignoring the temperature on the volume blocks.&lt;br&gt;auto &amp;dash; This policy allows tiering of both snapshot and active file system user data to the cloud store&lt;br&gt;none &amp;dash; Volume blocks will not be tiered to the cloud store.&lt;br&gt;snapshot_only &amp;dash; This policy allows tiering of only the volume Snapshot copies not associated with the active file system. The default tiering policy is \&quot;snapshot-only\&quot; for a FlexVol and \&quot;none\&quot; for a FlexGroup. | [optional] 
**Supported** | **bool** | This parameter specifies whether or not FabricPools are selected when provisioning a FlexGroup without specifying \&quot;aggregates.name\&quot; or \&quot;aggregates.uuid\&quot;. Only FabricPool aggregates are used if this parameter is set to true and only non FabricPool aggregates are used if this parameter is set to false. Tiering support for a FlexGroup can be changed by moving all of the constituents to the required aggregates. Note that in order to tier data, not only does the volume need to support tiering by using FabricPools, the tiering \&quot;policy\&quot; must not be &#39;none&#39;. A volume that uses FabricPools but has a tiering \&quot;policy\&quot; of &#39;none&#39; supports tiering, but will not tier any data. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


