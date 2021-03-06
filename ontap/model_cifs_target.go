/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// CifsTarget struct for CifsTarget
type CifsTarget struct {
	// Specify if the destination share is a home directory.
	HomeDirectory bool `json:"home_directory,omitempty"`
	// Specifies whether the CIFS symbolic link is a local link or wide link. The following values are supported: * local - Local symbolic link maps only to the same CIFS share. * widelink - Wide symbolic link maps to any CIFS share on the network.
	Locality string `json:"locality,omitempty"`
	// Specifies the CIFS path on the destination to which the symbolic link maps. The final path is generated by concatenating the CIFS server name, the share name, the cifs-path and the remaining path in the symbolic link left after the prefix match. This value is specified by using a UNIX-style path name. The trailing forward slash is required for the full path name to be properly interpreted.
	Path string `json:"path,omitempty"`
	// Specifies the destination CIFS server where the UNIX symbolic link is pointing. This field is mandatory if the locality of the symbolic link is 'widelink'. You can specify the value in any of the following formats:   * DNS name of the CIFS server.   * IP address of the CIFS server.   * NetBIOS name of the CIFS server.
	Server string `json:"server,omitempty"`
	// Specifies the CIFS share name on the destination CIFS server to which the UNIX symbolic link is pointing.
	Share string `json:"share,omitempty"`
}
