/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
import (
	"time"
)
// Job struct for Job
type Job struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// If the state indicates \"failure\", this is the final error code.
	Code int32 `json:"code,omitempty"`
	// The description of the job to help identify it independent of the UUID.
	Description string `json:"description,omitempty"`
	// The time the job ended.
	EndTime time.Time `json:"end_time,omitempty"`
	// A message corresponding to the state of the job providing additional details about the current state.
	Message string `json:"message,omitempty"`
	// The time the job started.
	StartTime time.Time `json:"start_time,omitempty"`
	// The state of the job.
	State string `json:"state,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
