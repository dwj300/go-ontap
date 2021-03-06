/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Set of ports that will receive a broadcast Ethernet packet from any of them
type BroadcastDomain struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	Ipspace *BroadcastDomainIpspace `json:"ipspace,omitempty"`
	// Maximum transmission unit, largest packet size on this network
	Mtu int32 `json:"mtu,omitempty"`
	// Name of the broadcast domain, scoped to its IPspace
	Name string `json:"name,omitempty"`
	// Ports that belong to the broadcast domain
	Ports []BroadcastDomainPorts `json:"ports,omitempty"`
	// Broadcast domain UUID
	Uuid string `json:"uuid,omitempty"`
}
