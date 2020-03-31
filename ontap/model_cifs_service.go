/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// CifsService struct for CifsService
type CifsService struct {
	Links    InlineResponse201Links `json:"_links,omitempty"`
	AdDomain AdDomain               `json:"ad_domain,omitempty"`
	// A descriptive text comment for the CIFS server. SMB clients can see the CIFS server comment when browsing servers on the network. If there is a space in the comment, you must enclose the entire string in quotation marks.
	Comment string `json:"comment,omitempty"`
	// Specifies the UNIX user to which any authenticated CIFS user is mapped to, if the normal user mapping rules fails.
	DefaultUnixUser string `json:"default_unix_user,omitempty"`
	// Specifies if the CIFS service is administratively enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the CIFS server.
	Name     string              `json:"name,omitempty"`
	Netbios  CifsNetbios         `json:"netbios,omitempty"`
	Security CifsServiceSecurity `json:"security,omitempty"`
	Svm      AuditSvm            `json:"svm,omitempty"`
}