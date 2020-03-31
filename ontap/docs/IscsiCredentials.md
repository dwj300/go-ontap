# IscsiCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**AuthenticationType** | **string** | The iSCSI authentication type. Required in POST and optional in PATCH. | [optional] 
**Chap** | [**IscsiCredentialsChap**](iscsi_credentials_chap.md) |  | [optional] 
**Initiator** | **string** | The iSCSI initiator to which the credentials apply. Required in POST.  | [optional] 
**InitiatorAddress** | [**IscsiCredentialsInitiatorAddress**](iscsi_credentials_initiator_address.md) |  | [optional] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


