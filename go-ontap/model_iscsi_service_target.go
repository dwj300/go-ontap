/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IscsiServiceTarget struct {
	// The iSCSI target alias of the iSCSI service.<br/> The target alias can contain one (1) to 128 characters and feature any printable character except space (\" \"). A PATCH request with an empty alias (\"\") clears the alias.<br/> Optional in POST and PATCH. In POST, this defaults to the name of the SVM. 
	Alias string `json:"alias,omitempty"`
	// The iSCSI target name of the iSCSI service. This is generated for the SVM during POST.<br/> If required, the target name can be modified using the ONTAP command line. 
	Name string `json:"name,omitempty"`
}
