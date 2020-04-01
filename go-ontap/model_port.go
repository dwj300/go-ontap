/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Port struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	BroadcastDomain *PortBroadcastDomain `json:"broadcast_domain,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Lag *PortLag `json:"lag,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	// MTU of the port in bytes. Set by broadcast domain.
	Mtu int32 `json:"mtu,omitempty"`
	// Portname, such as e0a, e1b-100 (VLAN on ethernet), a0c (LAG/ifgrp), a0d-200 (vlan on LAG/ifgrp)
	Name string `json:"name,omitempty"`
	Node *AutosupportMessageNode `json:"node,omitempty"`
	// Link speed in Mbps
	Speed int32 `json:"speed,omitempty"`
	// Operational state of the port.
	State string `json:"state,omitempty"`
	// Type of physical or virtual port
	Type_ string `json:"type,omitempty"`
	// Port UUID
	Uuid string `json:"uuid,omitempty"`
	Vlan *PortVlan `json:"vlan,omitempty"`
}