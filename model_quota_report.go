/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// QuotaReport struct for QuotaReport
type QuotaReport struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	Files QuotaReportFiles `json:"files,omitempty"`
	Group QuotaReportGroup `json:"group,omitempty"`
	// Index that identifies a unique quota record. Valid in URL.
	Index int32 `json:"index,omitempty"`
	Qtree QuotaReportQtree `json:"qtree,omitempty"`
	Space QuotaReportSpace `json:"space,omitempty"`
	// Quota specifier
	Specifier string `json:"specifier,omitempty"`
	Svm AuditSvm `json:"svm,omitempty"`
	// Quota type associated with the quota record.
	Type string `json:"type,omitempty"`
	// This parameter specifies the target user or users associated with the given quota report record. This parameter is available for user quota records and is not available for group or tree quota records. The target user or users are identified by a user name and user identifier. The user name can be a UNIX user name or a Windows user name, and the identifer can be a UNIX user identifier or a Windows security identifier.
	Users []QuotaReportUsers `json:"users,omitempty"`
	Volume CifsShareVolume `json:"volume,omitempty"`
}
