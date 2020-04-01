# LunMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Igroup** | [***LunMapIgroup**](lun_map_igroup.md) |  | [optional] [default to null]
**LogicalUnitNumber** | **int32** | The logical unit number assigned to the LUN when mapped to the specified initiator group. The number is used to identify the LUN to initiators in the initiator group when communicating through Fibre Channel Protocol or iSCSI. Optional in POST; if no value is provided, ONTAP assigns the lowest available value.  | [optional] [default to null]
**Lun** | [***LunMapLun**](lun_map_lun.md) |  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


