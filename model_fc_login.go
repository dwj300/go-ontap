/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// FcLogin A Fibre Channel (FC) login represents a connection formed by an FC initiator that has successfully logged in to ONTAP. This represents the FC login on which higher-level protocols such as Fibre Channel Protocol and NVMe over Fibre Channel (NVMe/FC) rely.
type FcLogin struct {
	Links AccountResponseLinks `json:"_links,omitempty"`
	// The initiator groups in which the initiator is a member.
	Igroups   []FcLoginIgroups `json:"igroups,omitempty"`
	Initiator FcLoginInitiator `json:"initiator,omitempty"`
	Interface FcLoginInterface `json:"interface,omitempty"`
	// The data protocol used to perform the login.
	Protocol string   `json:"protocol,omitempty"`
	Svm      AuditSvm `json:"svm,omitempty"`
}
