/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// NvmeInterface NVMe interfaces are network interfaces configured to support an NVMe over Fabrics (NVMe-oF) protocol. The NVMe interfaces are Fibre Channel interfaces supporting an NVMe-oF data protocol. Regardless of the underlying physical and data protocol, NVMe interfaces are treated equally for host-side application configuration. This endpoint provides a consolidated view of all NVMe interfaces for the purpose of configuring host-side applications.<br/> NVMe interfaces must be created using the protocol-specific endpoints for Fibre Channel interfaces. See [`POST /network/fc/interfaces`](#/networking/fc_interface_create). After creation, the interfaces are available via this interface.
type NvmeInterface struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// The administrative state of the NVMe interface.
	Enabled     bool                     `json:"enabled,omitempty"`
	FcInterface NvmeInterfaceFcInterface `json:"fc_interface,omitempty"`
	// The name of the NVMe interface.
	Name string                `json:"name,omitempty"`
	Node InlineResponse201Node `json:"node,omitempty"`
	Svm  AuditSvm              `json:"svm,omitempty"`
	// The transport address of the NVMe interface.
	TransportAddress string `json:"transport_address,omitempty"`
	// The unique identifier of the NVMe interface.
	Uuid string `json:"uuid,omitempty"`
}
