# IscsiSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Connections** | [**[]IscsiConnection**](iscsi_connection.md) | The iSCSI connections that make up the iSCSI session.  | [optional] [readonly] 
**Igroups** | [**[]FcLoginIgroups**](fc_login_igroups.md) | The initiator groups in which the initiator is a member.  | [optional] [readonly] 
**Initiator** | [**IscsiSessionInitiator**](iscsi_session_initiator.md) |  | [optional] 
**Isid** | **string** | The initiator portion of the session identifier specified by the initiator during login.  | [optional] [readonly] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**TargetPortalGroup** | **string** | The target portal group to which the session belongs.  | [optional] [readonly] 
**TargetPortalGroupTag** | **int32** | The target portal group tag of the session.  | [optional] [readonly] 
**Tsih** | **int32** | The target session identifier handle (TSIH) of the session.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


