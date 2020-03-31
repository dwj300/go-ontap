/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// KerberosInterface struct for KerberosInterface
type KerberosInterface struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// Specifies if Kerberos is enabled.
	Enabled bool `json:"enabled,omitempty"`
	EncryptionTypes []string `json:"encryption_types,omitempty"`
	Interface ClusterClusterInterfaces `json:"interface,omitempty"`
	// Load keytab from URI
	KeytabUri string `json:"keytab_uri,omitempty"`
	// Organizational unit
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// Account creation password
	Password string `json:"password,omitempty"`
	// Service principal name. Valid in PATCH.
	Spn string `json:"spn,omitempty"`
	Svm AuditSvm `json:"svm,omitempty"`
	// Account creation user name
	User string `json:"user,omitempty"`
}
