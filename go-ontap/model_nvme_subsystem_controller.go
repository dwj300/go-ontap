/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A Non-Volatile Memory Express (NVMe) subsystem controller represents a connection between a host and a storage solution.<br/> An NVMe subsystem controller is identified by the NVMe subsystem UUID and the controller ID. 
type NvmeSubsystemController struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	AdminQueue *NvmeSubsystemControllerAdminQueue `json:"admin_queue,omitempty"`
	Host *NvmeSubsystemControllerHost `json:"host,omitempty"`
	// The identifier of the subsystem controller. This field consists of 4 zero-filled hexadecimal digits followed by an 'h'. 
	Id string `json:"id,omitempty"`
	Interface_ *NvmeSubsystemControllerInterface `json:"interface,omitempty"`
	IoQueue *NvmeSubsystemControllerIoQueue `json:"io_queue,omitempty"`
	Node *InlineResponse201Node `json:"node,omitempty"`
	Subsystem *NvmeNamespaceSubsystemMapSubsystem `json:"subsystem,omitempty"`
	Svm *AuditSvm `json:"svm,omitempty"`
}
