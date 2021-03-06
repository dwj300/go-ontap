/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// VolumeNas struct for VolumeNas
type VolumeNas struct {
	ExportPolicy QtreeExportPolicy `json:"export_policy,omitempty"`
	// The UNIX group ID of the volume. Valid in POST or PATCH.
	Gid int32 `json:"gid,omitempty"`
	// The fully-qualified path in the owning SVM's namespace at which the volume is mounted. The path is case insensitive and must be unique within a SVM's namespace. Path must begin with '/' and must not end with '/'. Only one volume can be mounted at any given junction path. An empty path in POST creates an unmounted volume. An empty path in PATCH deactivates and unmounts the volume. This attribute is reported in GET only when the volume is mounted.
	Path string `json:"path,omitempty"`
	// Security style associated with the volume. Valid in POST or PATCH.<br>mixed &dash; Mixed-style security<br>ntfs &dash; NTFS/WIndows-style security<br>unified &dash; Unified-style security, unified UNIX, NFS and CIFS permissions<br>unix &dash; Unix-style security.
	SecurityStyle string `json:"security_style,omitempty"`
	// The UNIX user ID of the volume. Valid in POST or PATCH.
	Uid int32 `json:"uid,omitempty"`
	// UNIX permissions to be viewed as an octal number. It consists of 4 digits derived by adding up bits 4 (read), 2 (write) and 1 (execute). First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. The second digit selects permission for the owner of the file; the third selects permissions for other users in the same group; the fourth for other users not in the group. Valid in POST or PATCH. For security style \"mixed\" or \"unix\", the default setting is 0755 in octal (493 in decimal) and for security style \"ntfs\", the default setting is 0000. In cases where only owner, group and other permissions are given (as in 755, representing the second, third and fourth dight), first digit is assumed to be zero.
	UnixPermissions int32 `json:"unix_permissions,omitempty"`
}
