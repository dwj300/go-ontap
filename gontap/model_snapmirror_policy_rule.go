/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// SnapMirror policy rule for retention.
type SnapmirrorPolicyRule struct {
	// Number of Snapshot copies to be kept for retention.
	Count int32 `json:"count,omitempty"`
	CreationSchedule *SnapmirrorPolicyRuleCreationSchedule `json:"creation_schedule,omitempty"`
	// Snapshot copy label
	Label string `json:"label,omitempty"`
	// Specifies the prefix for the Snapshot copy name to be created as per the schedule. If no value is specified, then the label is used as the prefix.
	Prefix string `json:"prefix,omitempty"`
}
