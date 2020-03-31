/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// IscsiService An iSCSI service defines the properties of the iSCSI target for an SVM. There can be at most one iSCSI service for an SVM. An SVM's iSCSI service must be created before iSCSI initiators can log in to the SVM.<br/> An iSCSI service is identified by the UUID of its SVM. 
type IscsiService struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// The administrative state of the iSCSI service. The iSCSI service can be disabled to block all iSCSI connectivity to the SVM.<br/> Optional in POST and PATCH. The default setting is _true_ (enabled) in POST. 
	Enabled bool `json:"enabled,omitempty"`
	Svm AuditSvm `json:"svm,omitempty"`
	Target IscsiServiceTarget `json:"target,omitempty"`
}