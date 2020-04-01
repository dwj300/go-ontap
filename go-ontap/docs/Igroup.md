# Igroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**DeleteOnUnmap** | **bool** | An option that causes the initiator group to be deleted when the last LUN map associated with it is deleted. Optional in PATCH only; not available in POST. This property defaults to _false_ when the initiator group is created.  | [optional] [default to null]
**Initiators** | [**[]IgroupInitiatorNoRecords**](igroup_initiator_no_records.md) | The initiators that are members of the group. Optional in POST.&lt;br/&gt; Zero or more initiators can be supplied when the initiator group is created. After creation, initiators can be added or removed from the initiator group using the &#x60;/protocols/san/igroups/{igroup.uuid}/initiators&#x60; endpoint. See [&#x60;POST /protocols/san/igroups/{igroup.uuid}/initiators&#x60;](#/SAN/igroup_initiator_create) and [&#x60;DELETE /protocols/san/igroups/{igroup.uuid}/initiators/{name}&#x60;](#/SAN/igroup_initiator_delete) for more details.  | [optional] [default to null]
**LunMaps** | [**[]IgroupLunMaps**](igroup_lun_maps.md) | All LUN maps with which the initiator is associated.&lt;br/&gt; There is an added cost to retrieving property values for &#x60;lun_maps&#x60;. They not populated for either a collection GET or an instance GET unless explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [default to null]
**Name** | **string** | The name of the initiator group. Required in POST; optional in PATCH.&lt;br/&gt; Note that renaming an initiator group must be done in a PATCH request separate from any other modifications.  | [optional] [default to null]
**OsType** | **string** | The host operating system of the initiator group. All initiators in the group should be hosts of the same operating system. Required in POST; optional in PATCH.  | [optional] [default to null]
**Protocol** | **string** | The protocols supported by the initiator group. This restricts the type of initiators that can be added to the initiator group. Optional in POST; if not supplied, this defaults to _mixed_.&lt;br/&gt; The protocol of an initiator group cannot be changed after creation of the group.  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | The unique identifier of the initiator group.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


