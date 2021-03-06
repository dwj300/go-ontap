/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Specifies the file operations for the FPolicy event. You must specify a valid protocol in the protocol parameter. The event will check the operations specified from all client requests using the protocol. 
type FpolicyEventFileOperations struct {
	// File close operations
	Close bool `json:"close,omitempty"`
	// File create operations
	Create bool `json:"create,omitempty"`
	// Directory create operations
	CreateDir bool `json:"create_dir,omitempty"`
	// File delete operations
	Delete bool `json:"delete,omitempty"`
	// Directory delete operations
	DeleteDir bool `json:"delete_dir,omitempty"`
	// Get attribute operations
	Getattr bool `json:"getattr,omitempty"`
	// Link operations
	Link bool `json:"link,omitempty"`
	// Lookup operations
	Lookup bool `json:"lookup,omitempty"`
	// File open operations
	Open bool `json:"open,omitempty"`
	// File read operations
	Read bool `json:"read,omitempty"`
	// File rename operations
	Rename bool `json:"rename,omitempty"`
	// Directory rename operations
	RenameDir bool `json:"rename_dir,omitempty"`
	// Set attribute operations
	Setattr bool `json:"setattr,omitempty"`
	// Symbolic link operations
	Symlink bool `json:"symlink,omitempty"`
	// File write operations
	Write bool `json:"write,omitempty"`
}
