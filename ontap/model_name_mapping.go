/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// NameMapping Name mapping is used to map CIFS identities to UNIX identities, Kerberos identities to UNIX identities, and UNIX identities to CIFS identities. It needs this information to obtain user credentials and provide proper file access regardless of whether they are connecting from an NFS client or a CIFS client.
type NameMapping struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// Client workstation IP Address which is matched when searching for the pattern.   You can specify the value in any of the following formats: * As an IPv4 address with a subnet mask expressed as a number of bits; for instance, 10.1.12.0/24 * As an IPv6 address with a subnet mask expressed as a number of bits; for instance, fd20:8b1e:b255:4071::/64 * As an IPv4 address with a network mask; for instance, 10.1.16.0/255.255.255.0 * As a hostname
	ClientMatch string `json:"client_match,omitempty"`
	// Direction in which the name mapping is applied. The possible values are:   * krb_unix  - Kerberos principal name to UNIX user name   * win_unix  - Windows user name to UNIX user name   * unix_win  - UNIX user name to Windows user name mapping
	Direction string `json:"direction,omitempty"`
	// Position in the list of name mappings.
	Index int32 `json:"index,omitempty"`
	// Pattern used to match the name while searching for a name that can be used as a replacement. The pattern is a UNIX-style regular expression. Regular expressions are case-insensitive when mapping from Windows to UNIX, and they are case-sensitive for mappings from Kerberos to UNIX and UNIX to Windows.
	Pattern string `json:"pattern,omitempty"`
	// The name that is used as a replacement, if the pattern associated with this entry matches.
	Replacement string `json:"replacement,omitempty"`
	Svm         DnsSvm `json:"svm,omitempty"`
}
