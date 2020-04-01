/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The initiator group to which the LUN is mapped. Required in POST by supplying either the `igroup.uuid`, `igroup.name`, or both. 
type LunMapIgroup struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	// The initiators that are members of the initiator group. 
	Initiators []string `json:"initiators,omitempty"`
	// The name of the initiator group. Valid in POST. 
	Name string `json:"name,omitempty"`
	// The host operating system of the initiator group. All initiators in the group should be hosts of the same operating system. 
	OsType string `json:"os_type,omitempty"`
	// The protocols supported by the initiator group. This restricts the type of initiators that can be added to the initiator group. 
	Protocol string `json:"protocol,omitempty"`
	// The unique identifier of the initiator group. Valid in POST. 
	Uuid string `json:"uuid,omitempty"`
}
