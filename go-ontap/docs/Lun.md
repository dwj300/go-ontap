# Lun

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**AutoDelete** | **bool** | This property marks the LUN for auto deletion when the volume containing the LUN runs out of space. This is most commonly set on LUN clones.&lt;br/&gt; When set to _true_, the LUN becomes eligible for automatic deletion when the volume runs out of space. Auto deletion only occurs when the volume containing the LUN is also configured for auto deletion and free space in the volume decreases below a particular threshold.&lt;br/&gt; This property is optional in POST and PATCH. The default value for a new LUN is _false_.&lt;br/&gt; There is an added cost to retrieving this property&#39;s value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [default to null]
**Class** | **string** | The class of LUN. Only _regular_ LUNs can be created using the REST API.  | [optional] [default to null]
**Clone** | [***LunClone**](lun_clone.md) |  | [optional] [default to null]
**Comment** | **string** | A configurable comment available for use by the administrator. Valid in POST and PATCH.  | [optional] [default to null]
**Enabled** | **bool** | The enabled state of the LUN. LUNs can be disabled to prevent access to the LUN. Certain error conditions also cause the LUN to become disabled. If the LUN is disabled, you can consult the &#x60;state&#x60; property to determine if the LUN is administratively disabled (_offline_) or has become disabled as a result of an error. A LUN in an error condition can be brought online by setting the &#x60;enabled&#x60; property to _true_ or brought administratively offline by setting the &#x60;enabled&#x60; property to _false_. Upon creation, a LUN is enabled by default. Valid in PATCH.  | [optional] [default to null]
**Location** | [***LunLocation**](lun_location.md) |  | [optional] [default to null]
**LunMaps** | [**[]LunLunMaps**](lun_lun_maps.md) | The LUN maps with which the LUN is associated.&lt;br/&gt; There is an added cost to retrieving property values for &#x60;lun_maps&#x60;. They are not populated for either a collection GET or an instance GET unless explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [default to null]
**Movement** | [***LunMovement**](lun_movement.md) |  | [optional] [default to null]
**Name** | **string** | The fully qualified path name of the LUN composed of a \&quot;/vol\&quot; prefix, the volume name, the (optional) qtree name, and base name of the LUN. Valid in POST and PATCH.&lt;br/&gt; A PATCH that modifies the qtree and/or base name portion of the LUN path is considered a rename operation.&lt;br/&gt; A PATCH that modifies the volume portion of the LUN path begins an asynchronous LUN movement operation.  | [optional] [default to null]
**OsType** | **string** | The operating system type of the LUN.&lt;br/&gt; Required in POST when creating a LUN that is not a clone of another. Disallowed in POST when creating a LUN clone.  | [optional] [default to null]
**QosPolicy** | [***LunQosPolicy**](lun_qos_policy.md) |  | [optional] [default to null]
**SerialNumber** | **string** | The LUN serial number. The serial number is generated by ONTAP when the LUN is created.  | [optional] [default to null]
**Space** | [***LunSpace**](lun_space.md) |  | [optional] [default to null]
**Status** | [***LunStatus**](lun_status.md) |  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | The unique identifier of the LUN.  The UUID is generated by ONTAP when the LUN is created.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


