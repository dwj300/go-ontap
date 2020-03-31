# ClusterNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**ClusterInterface** | [**ClusterClusterInterface**](cluster_cluster_interface.md) |  | [optional] 
**ClusterInterfaces** | [**[]ClusterClusterInterfaces**](cluster_cluster_interfaces.md) |  | [optional] [readonly] 
**Controller** | [**ClusterController**](cluster_controller.md) |  | [optional] 
**Date** | [**time.Time**](time.Time.md) | Specifies the ISO-8601 format date and time on the node. | [optional] [readonly] 
**Ha** | [**ClusterHa**](cluster_ha.md) |  | [optional] 
**Location** | **string** |  | [optional] 
**ManagementInterface** | [**ClusterManagementInterface1**](cluster_management_interface_1.md) |  | [optional] 
**ManagementInterfaces** | [**[]ClusterClusterInterfaces**](cluster_cluster_interfaces.md) |  | [optional] [readonly] 
**Membership** | **string** | Possible values: * &lt;i&gt;available&lt;/i&gt; - If a node is available, this means it is detected on the internal cluster network and can be added to the cluster.  Nodes that have a membership of \&quot;available\&quot; are not returned when a GET request is called when the cluster exists. A query on the \&quot;membership\&quot; property for &lt;i&gt;available&lt;/i&gt; must be provided to scan for nodes on the cluster network. Nodes that have a membership of \&quot;available\&quot; are returned automatically before a cluster is created. * &lt;i&gt;joining&lt;/i&gt; - Joining nodes are in the process of being added to the cluster. The node may be progressing through the steps to become a member or might have failed. The job to add the node or create the cluster provides details on the current progress of the node. * &lt;i&gt;member&lt;/i&gt; - Nodes that are members have successfully joined the cluster.  | [optional] [readonly] 
**Model** | **string** |  | [optional] [readonly] 
**Name** | **string** |  | [optional] 
**SerialNumber** | **string** |  | [optional] [readonly] 
**ServiceProcessor** | [**ClusterServiceProcessor**](cluster_service_processor.md) |  | [optional] 
**Uptime** | **int32** | The total time, in seconds, that the node has been up. | [optional] [readonly] 
**Uuid** | **string** |  | [optional] [readonly] 
**Version** | [**ClusterVersion**](cluster_version.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


