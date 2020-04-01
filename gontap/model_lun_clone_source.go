/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The source LUN for a LUN clone operation. This can be specified using property `clone.source.uuid` or `clone.source.name`. If both properties are supplied, they must refer to the same LUN.<br/> Valid in POST to create a new LUN as a clone of the source.<br/> Valid in PATCH to overwrite an existing LUN's data as a clone of another. 
type LunCloneSource struct {
	// The fully qualified path name of the clone source LUN composed of a \"/vol\" prefix, the volume name, the (optional) qtree name, and base name of the LUN. Valid in POST and PATCH. 
	Name string `json:"name,omitempty"`
	// The unique identifier of the clone source LUN. Valid in POST and PATCH. 
	Uuid string `json:"uuid,omitempty"`
}
