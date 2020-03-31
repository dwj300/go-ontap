/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// NvmeNamespaceLocation The location of the NVMe namespace within the ONTAP cluster. Valid in POST.<br/> NVMe namespaces do not support rename, or movement between volumes. 
type NvmeNamespaceLocation struct {
	// The base name component of the NVMe namespace. Valid in POST.<br/> If properties `name` and `location.namespace` are specified in the same request, they must refer to the base name.<br/> NVMe namespaces do not support rename. 
	Namespace string `json:"namespace,omitempty"`
	Qtree NvmeNamespaceLocationQtree `json:"qtree,omitempty"`
	Volume NvmeNamespaceLocationVolume `json:"volume,omitempty"`
}
