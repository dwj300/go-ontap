# ClusterNodes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**ClusterInterface** | [***ClusterClusterInterface**](cluster_cluster_interface.md) |  | [optional] [default to null]
**ClusterInterfaces** | [**[]ClusterClusterInterfaces**](cluster_cluster_interfaces.md) |  | [optional] [default to null]
**Controller** | [***ClusterController**](cluster_controller.md) |  | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Specifies the ISO-8601 format date and time on the node. | [optional] [default to null]
**Ha** | [***ClusterHa**](cluster_ha.md) |  | [optional] [default to null]
**Location** | **string** |  | [optional] [default to null]
**ManagementInterface** | [***ClusterManagementInterface1**](cluster_management_interface_1.md) |  | [optional] [default to null]
**ManagementInterfaces** | [**[]ClusterClusterInterfaces**](cluster_cluster_interfaces.md) |  | [optional] [default to null]
**Membership** | **string** | Possible values: * &lt;i&gt;available&lt;/i&gt; - If a node is available, this means it is detected on the internal cluster network and can be added to the cluster.  Nodes that have a membership of \&quot;available\&quot; are not returned when a GET request is called when the cluster exists. A query on the \&quot;membership\&quot; property for &lt;i&gt;available&lt;/i&gt; must be provided to scan for nodes on the cluster network. Nodes that have a membership of \&quot;available\&quot; are returned automatically before a cluster is created. * &lt;i&gt;joining&lt;/i&gt; - Joining nodes are in the process of being added to the cluster. The node may be progressing through the steps to become a member or might have failed. The job to add the node or create the cluster provides details on the current progress of the node. * &lt;i&gt;member&lt;/i&gt; - Nodes that are members have successfully joined the cluster.  | [optional] [default to null]
**Model** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**SerialNumber** | **string** |  | [optional] [default to null]
**ServiceProcessor** | [***ClusterServiceProcessor**](cluster_service_processor.md) |  | [optional] [default to null]
**Uptime** | **int32** | The total time, in seconds, that the node has been up. | [optional] [default to null]
**Uuid** | **string** |  | [optional] [default to null]
**Version** | [***ClusterVersion**](cluster_version.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


