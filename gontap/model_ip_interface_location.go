/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Current or home location can be modified. Specifying a port implies a node. Specifying a node allows an appropriate port to be automatically selected. Ports are not valid and not shown for VIP interfaces. For POST, broadcast_domain can be specified alone or with home_node.
type IpInterfaceLocation struct {
	AutoRevert bool `json:"auto_revert,omitempty"`
	BroadcastDomain *IpInterfaceLocationBroadcastDomain `json:"broadcast_domain,omitempty"`
	Failover *FailoverScope `json:"failover,omitempty"`
	HomeNode *IpInterfaceLocationHomeNode `json:"home_node,omitempty"`
	HomePort *IpInterfaceLocationHomePort `json:"home_port,omitempty"`
	IsHome bool `json:"is_home,omitempty"`
	Node *IpInterfaceLocationNode `json:"node,omitempty"`
	Port *IpInterfaceLocationPort `json:"port,omitempty"`
}