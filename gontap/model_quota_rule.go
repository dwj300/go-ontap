/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QuotaRule struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	Files *QuotaRuleFiles `json:"files,omitempty"`
	Group *QuotaRuleGroup `json:"group,omitempty"`
	Qtree *QuotaRuleQtree `json:"qtree,omitempty"`
	Space *QuotaRuleSpace `json:"space,omitempty"`
	Svm *FlexcacheRelationshipSvm `json:"svm,omitempty"`
	// This parameter specifies the quota policy rule type. This is required in POST only and can take either one of the \\\"user\\\", \\\"group\\\" or \\\"tree\\\" values.
	Type_ string `json:"type,omitempty"`
	// This parameter enables user mapping for user quota policy rules. This is valid in POST or PATCH for user quota policy rules only.
	UserMapping bool `json:"user_mapping,omitempty"`
	// This parameter specifies the target user to which the user quota policy rule applies. This parameter takes single or multiple user names or identifiers. This parameter is valid only for the POST operation of a user quota policy rule. If this parameter is used as an input to create a group or a tree quota policy rule, the POST operation will fail with an appropriate error. For POST, this input parameter takes either a user name or a user identifier, not both. For default quota rules, the user name must be chosen and specified as \"\". For explicit user quota rules, this parameter can indicate either a user name or user identifier. The user name can be a UNIX user name or a Windows user name. If a name contains a space, enclose the entire value in quotes. A UNIX user name cannot include a backslash (\\) or an @ sign; user names with these characters are treated as Windows names. The user identifer can be a UNIX user identifier or a Windows security identifier. For multi-user quota, this parameter can contain multiple user targets separated by a comma.
	Users []QuotaReportUsers `json:"users,omitempty"`
	// Unique identifier for the quota policy rule. This field is generated when the quota policy rule is created.
	Uuid string `json:"uuid,omitempty"`
	Volume *QuotaRuleVolume `json:"volume,omitempty"`
}
