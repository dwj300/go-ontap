/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap
// ApplicationStatisticsSpace1 struct for ApplicationStatisticsSpace1
type ApplicationStatisticsSpace1 struct {
	// The available amount of space left in the application. Note that this field has limited meaning for SAN applications. Space may be considered used from ONTAP's perspective while the host filesystem still considers it available
	Available int32 `json:"available,omitempty"`
	// The amount of space that would currently be used if no space saving features were enabled. For example, if compression were the only space saving feature enabled, this field would represent the uncompressed amount of space used
	LogicalUsed int32 `json:"logical_used,omitempty"`
	// The originally requested amount of space that was provisioned for the application
	Provisioned int32 `json:"provisioned,omitempty"`
	// The amount of space reserved for system features such as snapshots that has not yet been used
	ReservedUnused int32 `json:"reserved_unused,omitempty"`
	// The amount of space saved by all enabled space saving features
	Savings int32 `json:"savings,omitempty"`
	// The amount of space that is currently being used by the application. Note that this includes any space reserved by the system for features such as snapshots
	Used int32 `json:"used,omitempty"`
	// The amount of space that is currently being used, excluding any space that is reserved by the system for features such as snapshots
	UsedExcludingReserves int32 `json:"used_excluding_reserves,omitempty"`
	// The percentage of the originally provisioned space that is currently being used by the application
	UsedPercent int32 `json:"used_percent,omitempty"`
}
