/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// NvmeNamespaceSubsystemMap The NVMe subsystem with which the NVMe namespace is associated. A namespace can be mapped to zero (0) or one (1) subsystems.<br/> There is an added cost to retrieving property values for `subsystem_map`. They are not populated for either a collection GET or an instance GET unless explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more. 
type NvmeNamespaceSubsystemMap struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.<br/> The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case \"h\". 
	Anagrpid string `json:"anagrpid,omitempty"`
	// The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.<br/> The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \"h\". 
	Nsid string `json:"nsid,omitempty"`
	Subsystem NvmeNamespaceSubsystemMapSubsystem `json:"subsystem,omitempty"`
}
