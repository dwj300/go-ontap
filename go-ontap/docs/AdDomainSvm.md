# AdDomainSvm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | The fully qualified domain name of the Windows Active Directory to which this CIFS server belongs. A CIFS server appears as a member of Windows server object in the Active Directory store.  | [optional] [default to null]
**OrganizationalUnit** | **string** | Specifies the organizational unit within the Active Directory domain to associate with the CIFS server.  | [optional] [default to null]
**Password** | **string** | The account password used to add this CIFS server to the Active Directory. This is not audited. Valid in POST only.  | [optional] [default to null]
**User** | **string** | The user account used to add this CIFS server to the Active Directory. Valid in POST only.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


