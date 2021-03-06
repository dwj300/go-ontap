/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// Aggregate struct for Aggregate
type Aggregate struct {
	Links        InlineResponse201Links `json:"_links,omitempty"`
	BlockStorage AggregateBlockStorage  `json:"block_storage,omitempty"`
	CloudStorage AggregateCloudStorage  `json:"cloud_storage,omitempty"`
	// Timestamp of aggregate creation
	CreateTime     string                  `json:"create_time,omitempty"`
	DataEncryption AggregateDataEncryption `json:"data_encryption,omitempty"`
	DrHomeNode     AggregateDrHomeNode     `json:"dr_home_node,omitempty"`
	HomeNode       AggregateHomeNode       `json:"home_node,omitempty"`
	// Aggregate name
	Name string        `json:"name,omitempty"`
	Node AggregateNode `json:"node,omitempty"`
	// SnapLock type
	SnaplockType string         `json:"snaplock_type,omitempty"`
	Space        AggregateSpace `json:"space,omitempty"`
	// Operational state of the aggregate
	State string `json:"state,omitempty"`
	// Aggregate UUID
	Uuid string `json:"uuid,omitempty"`
}
