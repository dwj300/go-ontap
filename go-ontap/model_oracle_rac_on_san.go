/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Oracle RAC using SAN.
type OracleRacOnSan struct {
	ArchiveLog *OracleOnNfsArchiveLog `json:"archive_log,omitempty"`
	Db *OracleOnNfsDb `json:"db,omitempty"`
	DbSids []OracleRacOnSanDbSids `json:"db_sids,omitempty"`
	GridBinary *OracleRacOnNfsGridBinary `json:"grid_binary,omitempty"`
	// The list of initiator groups to create. Optional in the POST or PATCH body
	NewIgroups []OracleRacOnSanNewIgroups `json:"new_igroups,omitempty"`
	OraHome *OracleOnNfsOraHome `json:"ora_home,omitempty"`
	OracleCrs *OracleRacOnNfsOracleCrs `json:"oracle_crs,omitempty"`
	// The name of the host OS running the application. Required in the POST body
	OsType string `json:"os_type,omitempty"`
	ProtectionType *MongoDbOnSanProtectionType `json:"protection_type,omitempty"`
	RedoLog *OracleOnNfsRedoLog `json:"redo_log,omitempty"`
}
