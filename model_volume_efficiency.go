/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// VolumeEfficiency struct for VolumeEfficiency
type VolumeEfficiency struct {
	// The system can be enabled/disabled compaction.<br>inline &dash; Data will be compacted first and written to the volume.<br>none &dash; None<br>mixed &dash; Read only field for FlexGroups, where some of the constituent volumes are compaction enabled and some are disabled.
	Compaction string `json:"compaction,omitempty"`
	// The system can be enabled/disabled compression.<br>inline &dash; Data will be compressed first and written to the volume.<br>background &dash; Data will be written to the volume and compressed later.<br>both &dash; Inline compression compresses the data and write to the volume, background compression compresses only the blocks on which inline compression is not run.<br>none &dash; None<br>mixed &dash; Read only field for FlexGroups, where some of the constituent volumes are compression enabled and some are disabled.
	Compression string `json:"compression,omitempty"`
	// The system can be enabled/disabled cross volume dedupe. it can be enabled only when dedupe is enabled.<br>inline &dash; Data will be cross volume deduped first and written to the volume.<br>background &dash; Data will be written to the volume and cross volume deduped later.<br>both &dash; Inline cross volume dedupe dedupes the data and write to the volume, background cross volume dedupe dedupes only the blocks on which inline dedupe is not run.<br>none &dash; None<br>mixed &dash; Read only field for FlexGroups, where some of the constituent volumes are cross volume dedupe enabled and some are disabled.
	CrossVolumeDedupe string `json:"cross_volume_dedupe,omitempty"`
	// The system can be enabled/disabled dedupe.<br>inline &dash; Data will be deduped first and written to the volume.<br>background &dash; Data will be written to the volume and deduped later.<br>both &dash; Inline dedupe dedupes the data and write to the volume, background dedupe dedupes only the blocks on which inline dedupe is not run.<br>none &dash; None<br>mixed &dash; Read only field for FlexGroups, where some of the constituent volumes are dedupe enabled and some are disabled.
	Dedupe string `json:"dedupe,omitempty"`
}