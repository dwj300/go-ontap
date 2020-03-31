/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// FcPortFabric Properties of the fabric to which the FC port is attached.
type FcPortFabric struct {
	// Reports if the physical port has established a connection with the FC fabric.
	Connected bool `json:"connected,omitempty"`
	// The negotiated data rate between the target FC port and the fabric in gigabits per second.
	ConnectedSpeed int32 `json:"connected_speed,omitempty"`
	// The name of the fabric to which the port is connected. This is only available when the FC port is connected to a fabric.<br/> There is an added cost to retrieving this property's value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter. See [`DOC Requesting specific fields`](#docs-docs-Requesting-specific-fields) to learn more.
	Name string `json:"name,omitempty"`
	// The FC port address of the host bus adapter (HBA) physical port.<br/> Each FC port in an FC switched fabric has its own unique FC port address for routing purposes. The FC port address is assigned by a switch in the fabric when that port logs in to the fabric. This property refers to the FC port address given to the physical host bus adapter (HBA) port when the port performs a fabric login (FLOGI).<br/> This is useful for obtaining statistics and diagnostic information from FC switches.<br/> This is a six-digit hexadecimal encoded numeric value.
	PortAddress string `json:"port_address,omitempty"`
	// The switch port to which the FC port is connected.
	SwitchPort string `json:"switch_port,omitempty"`
}