/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// VolumeClone struct for VolumeClone
type VolumeClone struct {
	// Specifies if this volume is a normal FlexVol or FlexClone. This field needs to be set when creating a FlexClone. Valid in POST.
	IsFlexclone    bool              `json:"is_flexclone,omitempty"`
	ParentSnapshot SnapshotReference `json:"parent_snapshot,omitempty"`
	ParentSvm      AuditSvm          `json:"parent_svm,omitempty"`
	ParentVolume   CifsShareVolume   `json:"parent_volume,omitempty"`
	// Percentage of FlexClone blocks split from its parent volume.
	SplitCompletePercent int32 `json:"split_complete_percent,omitempty"`
	// Space required by the containing-aggregate to split the FlexClone volume.
	SplitEstimate int32 `json:"split_estimate,omitempty"`
	// This field is set when split is executed on any FlexClone, that is when the FlexClone volume is split from its parent FlexVol. This field needs to be set for splitting a FlexClone form FlexVol. Valid in PATCH.
	SplitInitiated bool `json:"split_initiated,omitempty"`
}
