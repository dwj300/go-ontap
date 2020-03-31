/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// Disk struct for Disk
type Disk struct {
	// List of aggregates sharing this disk
	Aggregates []DiskAggregates `json:"aggregates,omitempty"`
	// Disk shelf bay
	Bay string `json:"bay,omitempty"`
	// Disk class
	Class string `json:"class,omitempty"`
	// Type of overlying disk container
	ContainerType   string                `json:"container_type,omitempty"`
	DrNode          DiskDrNode            `json:"dr_node,omitempty"`
	Drawer          DiskDrawer            `json:"drawer,omitempty"`
	FirmwareVersion string                `json:"firmware_version,omitempty"`
	HomeNode        InlineResponse201Node `json:"home_node,omitempty"`
	Model           string                `json:"model,omitempty"`
	// Cluster-wide disk name
	Name string                `json:"name,omitempty"`
	Node InlineResponse201Node `json:"node,omitempty"`
	// Pool to which disk is assigned
	Pool string `json:"pool,omitempty"`
	// Percentage of rated life used
	RatedLifeUsedPercent int32 `json:"rated_life_used_percent,omitempty"`
	// Revolutions per minute
	Rpm          int32          `json:"rpm,omitempty"`
	SerialNumber string         `json:"serial_number,omitempty"`
	Shelf        ShelfReference `json:"shelf,omitempty"`
	// State
	State string `json:"state,omitempty"`
	// Disk interface type
	Type string `json:"type,omitempty"`
	// The unique identifier for a disk
	Uid        string `json:"uid,omitempty"`
	UsableSize int32  `json:"usable_size,omitempty"`
	Vendor     string `json:"vendor,omitempty"`
}
