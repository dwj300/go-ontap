/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A named set of privileges that defines the rights an account has when it is assigned the role.
type Role struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	// Indicates if this is a built-in (pre-defined) role which cannot be modified or deleted.
	Builtin bool `json:"builtin,omitempty"`
	// Role name
	Name string `json:"name,omitempty"`
	Owner *RoleOwner `json:"owner,omitempty"`
	// The list of privileges that this role has been granted.
	Privileges []RolePrivilege `json:"privileges,omitempty"`
	// Scope of the entity. set to \"cluster\" for cluster owned objects and to \"svm\" for SVM owned objects.
	Scope string `json:"scope,omitempty"`
}
