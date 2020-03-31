/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// AuditSchedule Rotates the audit logs based on a schedule by using the time-based rotation parameters in any combination. The rotation schedule is calculated by using all the time-related values. This is mutually exclusive with log size.
type AuditSchedule struct {
	// Specifies the day of the month schedule to rotate audit log. Leave empty for all.
	Days []int32 `json:"days,omitempty"`
	// Specifies the hourly schedule to rotate audit log. Leave empty for all.
	Hours []int32 `json:"hours,omitempty"`
	// Specifies the minutes schedule to rotate the audit log.
	Minutes []int32 `json:"minutes,omitempty"`
	// Specifies the months schedule to rotate audit log. Leave empty for all.
	Months []int32 `json:"months,omitempty"`
	// Specifies the weekdays schedule to rotate audit log. Leave empty for all.
	Weekdays []int32 `json:"weekdays,omitempty"`
}
