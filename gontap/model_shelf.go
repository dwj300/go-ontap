/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Shelf struct {
	Bays []ShelfBays `json:"bays,omitempty"`
	ConnectionType string `json:"connection_type,omitempty"`
	DiskCount int32 `json:"disk_count,omitempty"`
	Drawers []ShelfDrawers `json:"drawers,omitempty"`
	Frus []ShelfFrus `json:"frus,omitempty"`
	Id string `json:"id,omitempty"`
	Internal bool `json:"internal,omitempty"`
	Model string `json:"model,omitempty"`
	ModuleType string `json:"module_type,omitempty"`
	Name string `json:"name,omitempty"`
	Paths []ShelfPaths `json:"paths,omitempty"`
	Ports []ShelfPorts `json:"ports,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	State string `json:"state,omitempty"`
	Uid string `json:"uid,omitempty"`
}
