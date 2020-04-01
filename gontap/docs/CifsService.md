# CifsService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**AdDomain** | [***AdDomain**](ad_domain.md) |  | [optional] [default to null]
**Comment** | **string** | A descriptive text comment for the CIFS server. SMB clients can see the CIFS server comment when browsing servers on the network. If there is a space in the comment, you must enclose the entire string in quotation marks. | [optional] [default to null]
**DefaultUnixUser** | **string** | Specifies the UNIX user to which any authenticated CIFS user is mapped to, if the normal user mapping rules fails. | [optional] [default to null]
**Enabled** | **bool** | Specifies if the CIFS service is administratively enabled.  | [optional] [default to null]
**Name** | **string** | The name of the CIFS server. | [optional] [default to null]
**Netbios** | [***CifsNetbios**](cifs_netbios.md) |  | [optional] [default to null]
**Security** | [***CifsServiceSecurity**](cifs_service_security.md) |  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


