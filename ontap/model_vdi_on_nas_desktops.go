/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// VdiOnNasDesktops struct for VdiOnNasDesktops
type VdiOnNasDesktops struct {
	// The number of desktops to support. Optional in the POST or PATCH body
	Count int32 `json:"count,omitempty"`
	// The size of the desktops. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]} Required in the POST body
	Size           int32                          `json:"size,omitempty"`
	StorageService VdiOnNasDesktopsStorageService `json:"storage_service,omitempty"`
}
