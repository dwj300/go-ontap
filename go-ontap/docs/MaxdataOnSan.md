# MaxdataOnSan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppType** | **string** | Type of the application that is being deployed on the L2. Required in the POST body | [optional] [default to null]
**ApplicationComponents** | [**[]MaxdataOnSanApplicationComponents**](maxdata_on_san_application_components.md) | application-components. Optional in the POST or PATCH body | [optional] [default to null]
**Metadata** | [**[]MaxdataOnSanMetadata**](maxdata_on_san_metadata.md) |  | [optional] [default to null]
**NewIgroups** | [**[]MaxdataOnSanNewIgroups**](maxdata_on_san_new_igroups.md) | The list of initiator groups to create. Optional in the POST or PATCH body | [optional] [default to null]
**OcsmUrl** | **string** | The OnCommand System Manager URL for this application | [optional] [default to null]
**OsType** | **string** | The name of the host OS running the application. Required in the POST body | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


