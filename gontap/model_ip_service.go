/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IpService string

// List of ip_service
const (
	CLUSTER_CORE_IpService IpService = "cluster_core"
	INTERCLUSTER_CORE_IpService IpService = "intercluster_core"
	MANAGEMENT_CORE_IpService IpService = "management_core"
	MANAGEMENT_AUTOSUPPORT_IpService IpService = "management_autosupport"
	MANAGEMENT_BGP_IpService IpService = "management_bgp"
	MANAGEMENT_HTTPS_IpService IpService = "management_https"
	MANAGEMENT_SSH_IpService IpService = "management_ssh"
	DATA_CORE_IpService IpService = "data_core"
	DATA_NFS_IpService IpService = "data_nfs"
	DATA_CIFS_IpService IpService = "data_cifs"
	DATA_FLEXCACHE_IpService IpService = "data_flexcache"
	DATA_ISCSI_IpService IpService = "data_iscsi"
)
