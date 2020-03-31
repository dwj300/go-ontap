/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// SnapmirrorRelationship SnapMirror relationship information
type SnapmirrorRelationship struct {
	Destination SnapmirrorEndpoint `json:"destination,omitempty"`
	// Snapshot copy exported to clients on destination.
	ExportedSnapshot string `json:"exported_snapshot,omitempty"`
	// Is the relationship healthy?
	Healthy bool `json:"healthy,omitempty"`
	// Time since the exported Snapshot copy was created.
	LagTime string `json:"lag_time,omitempty"`
	Policy SnapmirrorRelationshipPolicy `json:"policy,omitempty"`
	// Set to true on resync to preserve Snapshot copies on the destination that are newer than the latest common Snapshot copy. This field is applicable only for relationships with FlexGroup or FlexVol endpoints and when the PATCH state is \"snapmirrored\".
	Preserve bool `json:"preserve,omitempty"`
	// Set to true to reduce resync time by not preserving storage efficiency. This field is applicable only for relationships with FlexVol endpoints and when the PATCH state is \"snapmirrored\".
	QuickResync bool `json:"quick_resync,omitempty"`
	// Set to true to recover from a failed SnapMirror break operation on a FlexGroup relationship. This restores all destination FlexGroup constituents to the latest Snapshot copy, and any writes to the read-write constituents are lost. This field is applicable only for SnapMirror relationships with FlexGroup endpoints and when the PATCH state is \"broken_off\".
	RecoverAfterBreak bool `json:"recover_after_break,omitempty"`
	// Set to true to create a relationship for restore. To trigger restore-transfer, use transfers POST on the restore relationship.
	Restore bool `json:"restore,omitempty"`
	// Specifies the Snapshot copy to restore to on the destination after a break operation. This field is applicable only for SnapMirror relationships with FlexVol endpoints and when the PATCH state is \"broken_off\".
	RestoreToSnapshot string `json:"restore_to_snapshot,omitempty"`
	Source SnapmirrorEndpoint `json:"source,omitempty"`
	// State of the relationship. To initialize the relationship, PATCH the state to \"snapmirrored\". To break the relationship, PATCH the state to \"broken_off\". To resync the broken relationship, PATCH the state to \"snapmirrored\" for relationships with async policy type or \"in_sync\" for relationships with sync policy type. To pause the relationship, suspending further transfers, PATCH the state to \"paused\". To resume transfers for a paused relationship, PATCH the state to \"snapmirrored\" or \"in_sync\". The entries \"in_sync\", \"out_of_sync\", and \"synchronizing\" are only applicable to relationships of the sync policy type. A PATCH call on the state change only triggers the transition to the specified state. You must poll on the \"state\", \"healthy\" and \"unhealthy_reason\" fields using GET to determine if the transition is successful.
	State string `json:"state,omitempty"`
	Transfer SnapmirrorRelationshipTransfer `json:"transfer,omitempty"`
	// Reason the relationship is not healthy. It is a concatenation of up to four levels of error messages.
	UnhealthyReason []SnapmirrorError `json:"unhealthy_reason,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
