/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// EmsDestinationResponseRecords struct for EmsDestinationResponseRecords
type EmsDestinationResponseRecords struct {
	Links       InlineResponse201Links    `json:"_links,omitempty"`
	Certificate EmsDestinationCertificate `json:"certificate,omitempty"`
	// Event destination
	Destination string                  `json:"destination,omitempty"`
	Filters     []EmsDestinationFilters `json:"filters,omitempty"`
	// Destination name.  Valid in POST.
	Name string `json:"name,omitempty"`
	// Type of destination. Valid in POST.
	Type string `json:"type,omitempty"`
}
