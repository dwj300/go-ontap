# ExportRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**AnonymousUser** | **string** | User ID To Which Anonymous Users Are Mapped. | [optional] [default to null]
**Clients** | [**[]ExportClient**](export_client.md) | Array of client matches | [optional] [default to null]
**Index** | **int32** | Index of the rule within the export policy.  | [optional] [default to null]
**Protocols** | **[]string** |  | [optional] [default to null]
**RoRule** | [**[]ExportAuthenticationFlavor**](export_authentication_flavor.md) | Authentication flavors that the read-only access rule governs  | [optional] [default to null]
**RwRule** | [**[]ExportAuthenticationFlavor**](export_authentication_flavor.md) | Authentication flavors that the read/write access rule governs  | [optional] [default to null]
**Superuser** | [**[]ExportAuthenticationFlavor**](export_authentication_flavor.md) | Authentication flavors that the superuser security type governs  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


