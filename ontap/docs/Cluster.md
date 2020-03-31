# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**ConfigurationBackup** | [**ClusterConfigurationBackup**](cluster_configuration_backup.md) |  | [optional] 
**Contact** | **string** |  | [optional] 
**DnsDomains** | **[]string** | A list of DNS domains. Domain names have the following requirements: * The name must contain only the following characters: A through Z,   a through z, 0 through 9, \&quot;.\&quot;, \&quot;-\&quot; or \&quot;_\&quot;. * The first character of each label, delimited by \&quot;.\&quot;, must be one   of the following characters: A through Z or a through z or 0   through 9. * The last character of each label, delimited by \&quot;.\&quot;, must be one of   the following characters: A through Z, a through z, or 0 through 9. * The top level domain must contain only the following characters: A   through Z, a through z. * The system reserves the following names:\&quot;all\&quot;, \&quot;local\&quot;, and \&quot;localhost\&quot;.  | [optional] 
**License** | [**ClusterLicense**](cluster_license.md) |  | [optional] 
**Location** | **string** |  | [optional] 
**ManagementInterface** | [**ClusterManagementInterface**](cluster_management_interface.md) |  | [optional] 
**ManagementInterfaces** | [**[]ApplicationSanAccessIscsiEndpointInterface**](application_san_access_iscsi_endpoint_interface.md) |  | [optional] [readonly] 
**Metric** | [**ClusterMetric**](cluster_metric.md) |  | [optional] 
**Name** | **string** |  | [optional] 
**NameServers** | **[]string** | The list of IP addresses of the DNS servers. Addresses can be either IPv4 or IPv6 addresses.  | [optional] 
**Nodes** | [**[]ClusterNodes**](cluster_nodes.md) |  | [optional] 
**NtpServers** | **[]string** | Host name, IPv4 address, or IPv6 address for the external NTP time servers. | [optional] 
**Password** | **string** | Initial admin password used to create the cluster. | [optional] 
**Statistics** | [**ClusterStatistics**](cluster_statistics.md) |  | [optional] 
**Uuid** | **string** |  | [optional] [readonly] 
**Version** | [**ClusterVersion**](cluster_version.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


