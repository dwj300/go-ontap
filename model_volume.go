/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

import (
	"time"
)

// Volume struct for Volume
type Volume struct {
	Links InlineResponse201Links `json:"_links,omitempty"`
	// Aggregate hosting the volume. Required on POST.
	Aggregates  []DiskAggregates  `json:"aggregates,omitempty"`
	Application VolumeApplication `json:"application,omitempty"`
	Autosize    VolumeAutosize    `json:"autosize,omitempty"`
	Clone       VolumeClone       `json:"clone,omitempty"`
	// A comment for the volume. Valid in POST or PATCH.
	Comment string `json:"comment,omitempty"`
	// Specifies the number of times to iterate over the aggregates listed with the \"aggregates.name\" or \"aggregates.uuid\" when creating or expanding a FlexGroup. If a volume is being created on a single aggregate, the system will create a flexible volume if the \"constituents_per_aggregate\" field is not specified, and a FlexGroup if it is specified.  If a volume is being created on multiple aggregates, the system will always create a FlexGroup.
	ConstituentsPerAggregate int32 `json:"constituents_per_aggregate,omitempty"`
	// Creation time of the volume. This field is generated when the volume is created.
	CreateTime time.Time        `json:"create_time,omitempty"`
	Efficiency VolumeEfficiency `json:"efficiency,omitempty"`
	Encryption VolumeEncryption `json:"encryption,omitempty"`
	ErrorState VolumeErrorState `json:"error_state,omitempty"`
	Files      VolumeFiles      `json:"files,omitempty"`
	// FlexCache endpoint type. <br>none &dash; The volume is neither a FlexCache nor origin of any FlexCache. <br>cache &dash; The volume is a FlexCache volume. <br>origin &dash; The volume is origin of a FlexCache volume.
	FlexcacheEndpointType string          `json:"flexcache_endpoint_type,omitempty"`
	Guarantee             VolumeGuarantee `json:"guarantee,omitempty"`
	// Language encoding setting for volume. If no language is specified, the volume inherits its SVM language encoding setting.
	Language string         `json:"language,omitempty"`
	Metric   ClusterMetric  `json:"metric,omitempty"`
	Movement VolumeMovement `json:"movement,omitempty"`
	// Volume name. The name of volume must start with an alphabetic character (a to z or A to Z) or an underscore (_). The name must be 197 or fewer characters in length for FlexGroups, and 203 or fewer characters in length for all other types of volumes. Volume names must be unique within an SVM. Required on POST.
	Name  string      `json:"name,omitempty"`
	Nas   VolumeNas   `json:"nas,omitempty"`
	Qos   VolumeQos   `json:"qos,omitempty"`
	Quota VolumeQuota `json:"quota,omitempty"`
	// Physical size of the volume. The minimum size for a FlexVol volume is 20MB and the minimum size for a FlexGroup volume is 200MB per constituent. The recommended size for a FlexGroup volume is a minimum of 100GB per constituent. For all volumes, the default size is equal to the minimum size.
	Size           int32             `json:"size,omitempty"`
	Snaplock       VolumeSnaplock    `json:"snaplock,omitempty"`
	SnapshotPolicy SvmSnapshotPolicy `json:"snapshot_policy,omitempty"`
	Space          VolumeSpace       `json:"space,omitempty"`
	// Volume state. A volume can only be brought online if it is offline. The 'mixed' state applies to FlexGroup volumes only and cannot be specified as a target state. An 'error' state implies that the volume is not in a state to serve data.
	State      string            `json:"state,omitempty"`
	Statistics ClusterStatistics `json:"statistics,omitempty"`
	// The style of the volume. If \"style\" is not specified, the volume type is determined based on the specified aggregates. Specifying a single aggregate, without \"constituents_per_aggregate\" creates a flexible volume. Specifying multiple aggregates, or a single aggregate with \"constituents_per_aggregate\" creates a FlexGroup. If \"style\" is specified, a volume of that type is created. That is, if style is \"flexvol\", a single aggregate must be specified. If style is \"flexgroup\", the system either uses the specified aggregates, or automatically provisions if no aggregates are specified.<br>flexvol &dash; flexible volumes and FlexClone volumes<br>flexgroup &dash; FlexGroups.
	Style   string        `json:"style,omitempty"`
	Svm     VolumeSvm     `json:"svm,omitempty"`
	Tiering VolumeTiering `json:"tiering,omitempty"`
	// Type of the volume.<br>rw &dash; read-write volume.<br>dp &dash; data-protection volume.<br>ls &dash; load-sharing `dp` volume. Valid in GET.
	Type string `json:"type,omitempty"`
	// Specifies whether mirrored aggregates are selected when provisioning a FlexGroup without specifying \"aggregates.name\" or \"aggregates.uuid\". Only mirrored aggregates are used if this parameter is set to 'true' and only unmirrored aggregates are used if this parameter is set to 'false'. Aggregate level mirroring for a FlexGroup can be changed by moving all of the constituents to the required aggregates. The default value is 'true' for a MetroCluster configuration and is 'false' for a non-MetroCluster configuration.
	UseMirroredAggregates bool `json:"use_mirrored_aggregates,omitempty"`
	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	Uuid string `json:"uuid,omitempty"`
}
