/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StoragePort struct {
	BoardName string `json:"board_name,omitempty"`
	Cable *ShelfCable `json:"cable,omitempty"`
	Description string `json:"description,omitempty"`
	Error_ *StoragePortError `json:"error,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	Name string `json:"name,omitempty"`
	Node *InlineResponse201Node `json:"node,omitempty"`
	PartNumber string `json:"part_number,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	// Operational port speed in Gbps
	Speed float32 `json:"speed,omitempty"`
	State string `json:"state,omitempty"`
	// World Wide Name
	Wwn string `json:"wwn,omitempty"`
}
