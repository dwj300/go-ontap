# IscsiCredentials

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**AuthenticationType** | **string** | The iSCSI authentication type. Required in POST and optional in PATCH. | [optional] [default to null]
**Chap** | [***IscsiCredentialsChap**](iscsi_credentials_chap.md) |  | [optional] [default to null]
**Initiator** | **string** | The iSCSI initiator to which the credentials apply. Required in POST.  | [optional] [default to null]
**InitiatorAddress** | [***IscsiCredentialsInitiatorAddress**](iscsi_credentials_initiator_address.md) |  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


