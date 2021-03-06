/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// VscanOnDemandScope struct for VscanOnDemandScope
type VscanOnDemandScope struct {
	// List of file extensions for which scanning is not performed.
	ExcludeExtensions []string `json:"exclude_extensions,omitempty"`
	// List of file paths for which scanning must not be performed.
	ExcludePaths []string `json:"exclude_paths,omitempty"`
	// List of file extensions to be scanned.
	IncludeExtensions []string `json:"include_extensions,omitempty"`
	// Maximum file size, in bytes, allowed for scanning.
	MaxFileSize int32 `json:"max_file_size,omitempty"`
	// Specifies whether or not files without any extension can be scanned.
	ScanWithoutExtension bool `json:"scan_without_extension,omitempty"`
}
