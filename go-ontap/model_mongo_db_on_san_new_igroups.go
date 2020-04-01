/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The list of initiator groups to create.
type MongoDbOnSanNewIgroups struct {
	// The name of the new initiator group. Required in the POST body and optional in the PATCH body
	Name string `json:"name,omitempty"`
	Initiators []string `json:"initiators,omitempty"`
	// The name of the host OS accessing the application. The default value is the host OS that is running the application. Optional in the POST or PATCH body
	OsType string `json:"os_type,omitempty"`
	// The protocol of the new initiator group. Optional in the POST or PATCH body
	Protocol string `json:"protocol,omitempty"`
}
