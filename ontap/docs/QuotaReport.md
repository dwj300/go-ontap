# QuotaReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Files** | [**QuotaReportFiles**](quota_report_files.md) |  | [optional] 
**Group** | [**QuotaReportGroup**](quota_report_group.md) |  | [optional] 
**Index** | **int32** | Index that identifies a unique quota record. Valid in URL. | [optional] [readonly] 
**Qtree** | [**QuotaReportQtree**](quota_report_qtree.md) |  | [optional] 
**Space** | [**QuotaReportSpace**](quota_report_space.md) |  | [optional] 
**Specifier** | **string** | Quota specifier | [optional] [readonly] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Type** | **string** | Quota type associated with the quota record. | [optional] [readonly] 
**Users** | [**[]QuotaReportUsers**](quota_report_users.md) | This parameter specifies the target user or users associated with the given quota report record. This parameter is available for user quota records and is not available for group or tree quota records. The target user or users are identified by a user name and user identifier. The user name can be a UNIX user name or a Windows user name, and the identifer can be a UNIX user identifier or a Windows security identifier. | [optional] 
**Volume** | [**CifsShareVolume**](cifs_share_volume.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


