/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// Autosupport struct for Autosupport
type Autosupport struct {
	// Specifies whether to send the AutoSupport messages to vendor support.
	ContactSupport bool `json:"contact_support,omitempty"`
	// Specifies whether the AutoSupport daemon is enabled.  When this setting is disabled, delivery of all AutoSupport messages is turned off.
	Enabled bool `json:"enabled,omitempty"`
	// The e-mail address from which the AutoSupport messages are sent. To generate node-specific 'from' addresses, enable '-node-specific-from' parameter via ONTAP CLI.
	From string `json:"from,omitempty"`
	// Specifies whether the system information is collected in compliant form, to remove private data or in complete form, to enhance diagnostics.
	IsMinimal bool `json:"is_minimal,omitempty"`
	// A list of nodes in the cluster with connectivity issues to HTTP/SMTP/AOD AutoSupport destinations along with the corresponding error descriptions and corrective actions.
	Issues []AutosupportIssues `json:"issues,omitempty"`
	// The names of the mail servers used to deliver AutoSupport messages via SMTP.
	MailHosts []string `json:"mail_hosts,omitempty"`
	// The list of partner addresses.
	PartnerAddresses []string `json:"partner_addresses,omitempty"`
	// Proxy server for AutoSupport message delivery via HTTP/S. Optionally specify a username/password for authentication with the proxy server.
	ProxyUrl string `json:"proxy_url,omitempty"`
	// The e-mail addresses to which the AutoSupport messages are sent.
	To []string `json:"to,omitempty"`
	// The name of the transport protocol used to deliver AutoSupport messages.
	Transport string `json:"transport,omitempty"`
}
