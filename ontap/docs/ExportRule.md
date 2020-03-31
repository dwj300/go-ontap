# ExportRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**AnonymousUser** | **string** | User ID To Which Anonymous Users Are Mapped. | [optional] 
**Clients** | [**[]ExportClient**](export_client.md) | Array of client matches | [optional] 
**Index** | **int32** | Index of the rule within the export policy.  | [optional] [readonly] 
**Protocols** | **[]string** |  | [optional] 
**RoRule** | [**[]ExportAuthenticationFlavor**](export_authentication_flavor.md) | Authentication flavors that the read-only access rule governs  | [optional] 
**RwRule** | [**[]ExportAuthenticationFlavor**](export_authentication_flavor.md) | Authentication flavors that the read/write access rule governs  | [optional] 
**Superuser** | [**[]ExportAuthenticationFlavor**](export_authentication_flavor.md) | Authentication flavors that the superuser security type governs  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


