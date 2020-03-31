# LunMovement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxThroughput** | **string** | The maximum data throughput that should be utilized in support of the LUN movement. This property can be used to throttle a transfer and limit its impact on the performance of the source and destination nodes. The specified value will be rounded up to the nearest megabyte.&lt;br/&gt; If this property is not specified in a POST that begins a LUN movement, throttling is not applied to the data transfer.&lt;br/&gt; For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.&lt;br/&gt; This property is valid only in a POST that begins a LUN movement or a PATCH when a LUN movement is already in process.  | [optional] 
**Paths** | [**LunMovementPaths**](lun_movement_paths.md) |  | [optional] 
**Progress** | [**LunMovementProgress**](lun_movement_progress.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


