/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The storage space related properties of the NVMe namespace. 
type NvmeNamespaceSpace struct {
	// The size of blocks in the namespace in bytes.<br/> Valid in POST when creating an NVMe namespace that is not a clone of another. Disallowed in POST when creating a namespace clone.  Valid in POST. 
	BlockSize int32 `json:"block_size,omitempty"`
	Guarantee *NvmeNamespaceSpaceGuarantee `json:"guarantee,omitempty"`
	// The total provisioned size of the NVMe namespace.<br/> NVMe namespaces do not support resize.<br/> For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation. 
	Size int32 `json:"size,omitempty"`
	// The amount of space consumed by the main data stream of the NVMe namespace.<br/> This value is the total space consumed in the volume by the NVMe namespace, including filesystem overhead, but excluding prefix and suffix streams. Due to internal filesystem overhead and the many ways NVMe filesystems and applications utilize blocks within a namespace, this value does not necessarily reflect actual consumption/availability from the perspective of the filesystem or application. Without specific knowledge of how the namespace blocks are utilized outside of ONTAP, this property should not be used and an indicator for an out-of-space condition.<br/> For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation. 
	Used int32 `json:"used,omitempty"`
}
