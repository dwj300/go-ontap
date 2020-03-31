# QuotaRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Files** | [**QuotaRuleFiles**](quota_rule_files.md) |  | [optional] 
**Group** | [**QuotaRuleGroup**](quota_rule_group.md) |  | [optional] 
**Qtree** | [**QuotaReportQtree**](quota_report_qtree.md) |  | [optional] 
**Space** | [**QuotaRuleSpace**](quota_rule_space.md) |  | [optional] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Type** | **string** | This parameter specifies the quota policy rule type. This is required in POST only and can take either one of the \\\&quot;user\\\&quot;, \\\&quot;group\\\&quot; or \\\&quot;tree\\\&quot; values. | [optional] 
**UserMapping** | **bool** | This parameter enables user mapping for user quota policy rules. This is valid in POST or PATCH for user quota policy rules only. | [optional] 
**Users** | [**[]QuotaReportUsers**](quota_report_users.md) | This parameter specifies the target user to which the user quota policy rule applies. This parameter takes single or multiple user names or identifiers. This parameter is valid only for the POST operation of a user quota policy rule. If this parameter is used as an input to create a group or a tree quota policy rule, the POST operation will fail with an appropriate error. For POST, this input parameter takes either a user name or a user identifier, not both. For default quota rules, the user name must be chosen and specified as \&quot;\&quot;. For explicit user quota rules, this parameter can indicate either a user name or user identifier. The user name can be a UNIX user name or a Windows user name. If a name contains a space, enclose the entire value in quotes. A UNIX user name cannot include a backslash (\\) or an @ sign; user names with these characters are treated as Windows names. The user identifer can be a UNIX user identifier or a Windows security identifier. For multi-user quota, this parameter can contain multiple user targets separated by a comma. | [optional] 
**Uuid** | **string** | Unique identifier for the quota policy rule. This field is generated when the quota policy rule is created. | [optional] [readonly] 
**Volume** | [**CifsShareVolume**](cifs_share_volume.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


