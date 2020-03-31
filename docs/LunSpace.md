# LunSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guarantee** | [**LunSpaceGuarantee**](lun_space_guarantee.md) |  | [optional] 
**Size** | **int32** | The total provisioned size of the LUN. The LUN size can be increased but not be made smaller using the REST interface.&lt;br/&gt; For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.  | [optional] 
**Used** | **int32** | The amount of space consumed by the main data stream of the LUN.&lt;br/&gt; This value is the total space consumed in the volume by the LUN, including filesystem overhead, but excluding prefix and suffix streams. Due to internal filesystem overhead and the many ways SAN filesystems and applications utilize blocks within a LUN, this value does not necessarily reflect actual consumption/availability from the perspective of the filesystem or application. Without specific knowledge of how the LUN blocks are utilized outside of ONTAP, this property should not be used as an indicator for an out-of-space condition.&lt;br/&gt; For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


