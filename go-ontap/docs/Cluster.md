# Cluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**ConfigurationBackup** | [***ClusterConfigurationBackup**](cluster_configuration_backup.md) |  | [optional] [default to null]
**Contact** | **string** |  | [optional] [default to null]
**DnsDomains** | **[]string** | A list of DNS domains. Domain names have the following requirements: * The name must contain only the following characters: A through Z,   a through z, 0 through 9, \&quot;.\&quot;, \&quot;-\&quot; or \&quot;_\&quot;. * The first character of each label, delimited by \&quot;.\&quot;, must be one   of the following characters: A through Z or a through z or 0   through 9. * The last character of each label, delimited by \&quot;.\&quot;, must be one of   the following characters: A through Z, a through z, or 0 through 9. * The top level domain must contain only the following characters: A   through Z, a through z. * The system reserves the following names:\&quot;all\&quot;, \&quot;local\&quot;, and \&quot;localhost\&quot;.  | [optional] [default to null]
**License** | [***ClusterLicense**](cluster_license.md) |  | [optional] [default to null]
**Location** | **string** |  | [optional] [default to null]
**ManagementInterface** | [***ClusterManagementInterface**](cluster_management_interface.md) |  | [optional] [default to null]
**ManagementInterfaces** | [**[]ApplicationSanAccessIscsiEndpointInterface**](application_san_access_iscsi_endpoint_interface.md) |  | [optional] [default to null]
**Metric** | [***ClusterMetric**](cluster_metric.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NameServers** | **[]string** | The list of IP addresses of the DNS servers. Addresses can be either IPv4 or IPv6 addresses.  | [optional] [default to null]
**Nodes** | [**[]ClusterNodes**](cluster_nodes.md) |  | [optional] [default to null]
**NtpServers** | **[]string** | Host name, IPv4 address, or IPv6 address for the external NTP time servers. | [optional] [default to null]
**Password** | **string** | Initial admin password used to create the cluster. | [optional] [default to null]
**Statistics** | [***ClusterStatistics**](cluster_statistics.md) |  | [optional] [default to null]
**Uuid** | **string** |  | [optional] [default to null]
**Version** | [***ClusterVersion**](cluster_version.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


