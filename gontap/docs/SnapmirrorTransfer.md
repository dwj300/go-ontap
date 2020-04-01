# SnapmirrorTransfer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**BytesTransferred** | **int32** | Bytes transferred | [optional] [default to null]
**CheckpointSize** | **int32** | Amount of data transferred in bytes as recorded in the restart checkpoint. | [optional] [default to null]
**Files** | [**[]SnapmirrorTransferFiles**](snapmirror_transfer_files.md) | This is supported for transfer of restore relationship only. This specifies the list of files or LUNs to be restored. Can contain up to eight files or LUNs. | [optional] [default to null]
**Relationship** | [***SnapmirrorTransferRelationship**](snapmirror_transfer_relationship.md) |  | [optional] [default to null]
**Snapshot** | **string** | Name of Snapshot copy being transferred. | [optional] [default to null]
**SourceSnapshot** | **string** | Specifies the Snapshot copy on the source to be transferred to the destination. | [optional] [default to null]
**State** | **string** | Status of the transfer. Set PATCH state to \&quot;aborted\&quot; to abort the transfer. Set PATCH state to \&quot;hard_aborted\&quot; to abort the transfer and discard the restart checkpoint. | [optional] [default to null]
**StorageEfficiencyEnabled** | **bool** | This is supported for transfer of restore relationship only. Set this field to &#39;false&#39; to turn off storage efficiency for data transferred over the wire and written to the destination. | [optional] [default to null]
**Uuid** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


