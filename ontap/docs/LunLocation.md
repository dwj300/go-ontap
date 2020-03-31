# LunLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalUnit** | **string** | The base name component of the LUN. Valid in POST and PATCH.&lt;br/&gt; If properties &#x60;name&#x60; and &#x60;location.logical_unit&#x60; are specified in the same request, they must refer to the base name.&lt;br/&gt; A PATCH that modifies the base name of the LUN is considered a rename operation.  | [optional] 
**Qtree** | [**LunLocationQtree**](lun_location_qtree.md) |  | [optional] 
**Volume** | [**LunLocationVolume**](lun_location_volume.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


