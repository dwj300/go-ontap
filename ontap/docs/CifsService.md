# CifsService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**AdDomain** | [**AdDomain**](ad_domain.md) |  | [optional] 
**Comment** | **string** | A descriptive text comment for the CIFS server. SMB clients can see the CIFS server comment when browsing servers on the network. If there is a space in the comment, you must enclose the entire string in quotation marks. | [optional] 
**DefaultUnixUser** | **string** | Specifies the UNIX user to which any authenticated CIFS user is mapped to, if the normal user mapping rules fails. | [optional] [default to pcuser]
**Enabled** | **bool** | Specifies if the CIFS service is administratively enabled.  | [optional] [default to true]
**Name** | **string** | The name of the CIFS server. | [optional] 
**Netbios** | [**CifsNetbios**](cifs_netbios.md) |  | [optional] 
**Security** | [**CifsServiceSecurity**](cifs_service_security.md) |  | [optional] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


