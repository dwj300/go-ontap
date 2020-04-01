/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Complete node information
type ClusterNodes struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	ClusterInterface *ClusterClusterInterface `json:"cluster_interface,omitempty"`
	ClusterInterfaces []ClusterClusterInterfaces `json:"cluster_interfaces,omitempty"`
	Controller *ClusterController `json:"controller,omitempty"`
	// Specifies the ISO-8601 format date and time on the node.
	Date time.Time `json:"date,omitempty"`
	Ha *ClusterHa `json:"ha,omitempty"`
	Location string `json:"location,omitempty"`
	ManagementInterface *ClusterManagementInterface1 `json:"management_interface,omitempty"`
	ManagementInterfaces []ClusterClusterInterfaces `json:"management_interfaces,omitempty"`
	// Possible values: * <i>available</i> - If a node is available, this means it is detected on the internal cluster network and can be added to the cluster.  Nodes that have a membership of \"available\" are not returned when a GET request is called when the cluster exists. A query on the \"membership\" property for <i>available</i> must be provided to scan for nodes on the cluster network. Nodes that have a membership of \"available\" are returned automatically before a cluster is created. * <i>joining</i> - Joining nodes are in the process of being added to the cluster. The node may be progressing through the steps to become a member or might have failed. The job to add the node or create the cluster provides details on the current progress of the node. * <i>member</i> - Nodes that are members have successfully joined the cluster. 
	Membership string `json:"membership,omitempty"`
	Model string `json:"model,omitempty"`
	Name string `json:"name,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	ServiceProcessor *ClusterServiceProcessor `json:"service_processor,omitempty"`
	// The total time, in seconds, that the node has been up.
	Uptime int32 `json:"uptime,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Version *ClusterVersion `json:"version,omitempty"`
}
