/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// SvmCifsAdDomain struct for SvmCifsAdDomain
type SvmCifsAdDomain struct {
	// The fully qualified domain name of the Windows Active Directory to which this CIFS server belongs. A CIFS server appears as a member of Windows server object in the Active Directory store.
	Fqdn string `json:"fqdn,omitempty"`
	// Specifies the organizational unit within the Active Directory domain to associate with the CIFS server.
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// The account password used to add this CIFS server to the Active Directory. This is not audited. Valid in POST only.
	Password string `json:"password,omitempty"`
	// The user account used to add this CIFS server to the Active Directory. Valid in POST only.
	User string `json:"user,omitempty"`
}
