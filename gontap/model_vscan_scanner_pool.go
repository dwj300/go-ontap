/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Scanner pool is a set of attributes which are used to validate and manage connections between clustered ONTAP and external virus-scanning server, or \"Vscan server\".
type VscanScannerPool struct {
	Cluster *ClusterReference `json:"cluster,omitempty"`
	// Specifies the name of the scanner pool. Scanner pool name can be up to 256 characters long and is a string that can only contain any combination of ASCII-range alphanumeric characters a-z, A-Z, 0-9), \"_\", \"-\" and \".\".
	Name string `json:"name,omitempty"`
	// Specifies a list of privileged users. A valid form of privileged user-name is \"domain-name\\user-name\". Privileged user-names are stored and treated as case-insensitive strings. Virus scanners must use one of the registered privileged users for connecting to clustered Data ONTAP for exchanging virus-scanning protocol messages and to access file for scanning, remedying and quarantining operations.
	PrivilegedUsers []string `json:"privileged_users,omitempty"`
	// Specifies the role of the scanner pool. The possible values are:   * primary   - Always active.   * secondary - Active only when none of the primary external virus-scanning servers are connected.   * idle      - Always inactive. 
	Role string `json:"role,omitempty"`
	// Specifies a list of IP addresses or FQDN for each Vscan server host names which are allowed to connect to clustered ONTAP.
	Servers []string `json:"servers,omitempty"`
}
