/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ServiceProcessor struct {
	// Set to true to use DHCP to configure an IPv4 interface.
	DhcpEnabled bool `json:"dhcp_enabled,omitempty"`
	// The version of firmware installed.
	FirmwareVersion string `json:"firmware_version,omitempty"`
	Ipv4Interface *ClusterManagementInterfaceIp `json:"ipv4_interface,omitempty"`
	Ipv6Interface *ClusterServiceProcessorIpv6Interface `json:"ipv6_interface,omitempty"`
	LinkStatus string `json:"link_status,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	State string `json:"state,omitempty"`
}
