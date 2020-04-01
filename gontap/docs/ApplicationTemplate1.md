# ApplicationTemplate1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the template that was used to provision this application. Optional in the POST body | [optional] [default to null]
**Protocol** | **string** | The protocol access of the template that was used to provision this application | [optional] [default to null]
**Version** | **int32** | The version of the template that was used to provision this application. The template version changes only if the layout of the application changes over time. For example, redo logs in Oracle RAC templates were updated and provisioned differently in DATA ONTAP 9.3.0 compared to prior releases, so the version number was increased. If layouts change in the future, the changes will be documented along with the corresponding version numbers | [optional] [default to null]
**Links** | [***SelfLink**](self_link.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


