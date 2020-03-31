# VolumeAutosize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrowThreshold** | **int32** | Used space threshold size, in percentage, for the automatic growth of the volume. When the amount of used space in the volume becomes greater than this threhold, the volume automatically grows unless it has reached the maximum size. The volume grows when &#39;space.used&#39; is greater than this percent of &#39;space.size&#39;. The &#39;grow_threshold&#39; size cannot be less than or equal to the &#39;shrink_threshold&#39; size.. | [optional] 
**Maximum** | **int32** | Maximum size in bytes up to which a volume grows automatically. This size cannot be less than the current volume size, or less than or equal to the minimum size of volume. | [optional] 
**Minimum** | **int32** | Minimum size in bytes up to which the volume shrinks automatically. This size cannot be greater than or equal to the maximum size of volume. | [optional] 
**Mode** | **string** | Autosize mode for the volume.&lt;br&gt;grow &amp;dash; Volume automatically grows when the amount of used space is above the &#39;grow_threshold&#39; value.&lt;br&gt;grow_shrink &amp;dash; Volume grows or shrinks in response to the amount of space used.&lt;br&gt;off &amp;dash; Autosizing of the volume is disabled. | [optional] 
**ShrinkThreshold** | **int32** | Used space threshold size, in percentage, for the automatic shrinkage of the volume.  When the amount of used space in the volume drops below this threshold, the volume automatically shrinks unless it has reached the minimum size. The volume shrinks when the &#39;space.used&#39; is less than the &#39;shrink_threshold&#39; percent of &#39;space.size&#39;. The &#39;shrink_threshold&#39; size cannot be greater than or equal to the &#39;grow_threshold&#39; size. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


