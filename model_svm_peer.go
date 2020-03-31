/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// SvmPeer An SVM peer relation object.
type SvmPeer struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// A list of applications for an SVM peer relation.
	Applications []SvmPeerApplications `json:"applications,omitempty"`
	// A peer SVM alias name to avoid a name conflict on the local cluster.
	Name string `json:"name,omitempty"`
	Peer SvmPeerPeer `json:"peer,omitempty"`
	// SVM peering state. To accept a pending SVM peer request, PATCH the state to \"peered\". To reject a pending SVM peer request, PATCH the state to \"rejected\". To suspend a peered SVM peer relation, PATCH the state to \"suspended\". To resume a suspended SVM peer relation, PATCH the state to \"peered\". The states \"initiated\", \"pending\", and \"initializing\" are system-generated and cannot be used for PATCH.
	State string `json:"state,omitempty"`
	Svm SvmPeerSvm `json:"svm,omitempty"`
	// SVM peer relation UUID
	Uuid string `json:"uuid,omitempty"`
}
