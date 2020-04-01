/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method. 
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Complete cluster information
type Cluster struct {
	Links *InlineResponse201Links `json:"_links,omitempty"`
	ConfigurationBackup *ClusterConfigurationBackup `json:"configuration_backup,omitempty"`
	Contact string `json:"contact,omitempty"`
	// A list of DNS domains. Domain names have the following requirements: * The name must contain only the following characters: A through Z,   a through z, 0 through 9, \".\", \"-\" or \"_\". * The first character of each label, delimited by \".\", must be one   of the following characters: A through Z or a through z or 0   through 9. * The last character of each label, delimited by \".\", must be one of   the following characters: A through Z, a through z, or 0 through 9. * The top level domain must contain only the following characters: A   through Z, a through z. * The system reserves the following names:\"all\", \"local\", and \"localhost\". 
	DnsDomains []string `json:"dns_domains,omitempty"`
	License *ClusterLicense `json:"license,omitempty"`
	Location string `json:"location,omitempty"`
	ManagementInterface *ClusterManagementInterface `json:"management_interface,omitempty"`
	ManagementInterfaces []ApplicationSanAccessIscsiEndpointInterface `json:"management_interfaces,omitempty"`
	Metric *ClusterMetric `json:"metric,omitempty"`
	Name string `json:"name,omitempty"`
	// The list of IP addresses of the DNS servers. Addresses can be either IPv4 or IPv6 addresses. 
	NameServers []string `json:"name_servers,omitempty"`
	Nodes []ClusterNodes `json:"nodes,omitempty"`
	// Host name, IPv4 address, or IPv6 address for the external NTP time servers.
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Initial admin password used to create the cluster.
	Password string `json:"password,omitempty"`
	Statistics *ClusterStatistics `json:"statistics,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Version *ClusterVersion `json:"version,omitempty"`
}
