# IscsiSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Connections** | [**[]IscsiConnection**](iscsi_connection.md) | The iSCSI connections that make up the iSCSI session.  | [optional] [default to null]
**Igroups** | [**[]FcLoginIgroups**](fc_login_igroups.md) | The initiator groups in which the initiator is a member.  | [optional] [default to null]
**Initiator** | [***IscsiSessionInitiator**](iscsi_session_initiator.md) |  | [optional] [default to null]
**Isid** | **string** | The initiator portion of the session identifier specified by the initiator during login.  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**TargetPortalGroup** | **string** | The target portal group to which the session belongs.  | [optional] [default to null]
**TargetPortalGroupTag** | **int32** | The target portal group tag of the session.  | [optional] [default to null]
**Tsih** | **int32** | The target session identifier handle (TSIH) of the session.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


