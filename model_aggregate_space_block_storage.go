/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// AggregateSpaceBlockStorage struct for AggregateSpaceBlockStorage
type AggregateSpaceBlockStorage struct {
	// Space available in bytes
	Available int32 `json:"available,omitempty"`
	// The aggregate used percentage at which 'monitor.volume.full' EMS is generated.
	FullThresholdPercent int32 `json:"full_threshold_percent,omitempty"`
	// The size that is physically used in the block storage and has a cold temperature, in bytes. This property is only supported if the aggregate is either attached to a cloud store or can be attached to a cloud store. This is an advanced property; there is an added cost to retrieving its value. The field is not populated for either a collection GET or an instance GET unless it is explicitly requested using the <i>fields</i> query parameter containing either block_storage.inactive_user_data or **. 
	InactiveUserData int32 `json:"inactive_user_data,omitempty"`
	// Total usable space in bytes, not including WAFL reserve and aggregate Snapshot copy reserve.
	Size int32 `json:"size,omitempty"`
	// Space used or reserved in bytes. Includes volume guarantees and aggregate metadata.
	Used int32 `json:"used,omitempty"`
}
