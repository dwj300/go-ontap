/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SecurityAuditLogForward struct {
	// Destination syslog|splunk host to forward audit records to. This can be an IP address (IPv4|IPv6) or a hostname.
	Address string `json:"address,omitempty"`
	// This is the standard Syslog Facility value that is used when sending audit records to a remote server.
	Facility string `json:"facility,omitempty"`
	// Destination Port. The default port depends on the protocol chosen: For un-encrypted destinations the default port is 514. For encrypted destinations the default port is 6514. 
	Port int32 `json:"port,omitempty"`
	// Log forwarding protocol
	Protocol string `json:"protocol,omitempty"`
	// This is only applicable when the protocol is tcp_encrypted. This controls whether the remote server's certificate is validated. Setting \"verify_server\" to \"true\" will enforce validation of remote server's certificate. Setting \"verify_server\" to \"false\" will not enforce validation of remote server's certificate.
	VerifyServer bool `json:"verify_server,omitempty"`
}
