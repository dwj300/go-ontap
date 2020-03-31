# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**BroadcastDomain** | [**PortBroadcastDomain**](port_broadcast_domain.md) |  | [optional] 
**Enabled** | **bool** |  | [optional] 
**Lag** | [**PortLag**](port_lag.md) |  | [optional] 
**MacAddress** | **string** |  | [optional] [readonly] 
**Mtu** | **int32** | MTU of the port in bytes. Set by broadcast domain. | [optional] [readonly] 
**Name** | **string** | Portname, such as e0a, e1b-100 (VLAN on ethernet), a0c (LAG/ifgrp), a0d-200 (vlan on LAG/ifgrp) | [optional] [readonly] 
**Node** | [**InlineResponse201Node**](inline_response_201_node.md) |  | [optional] 
**Speed** | **int32** | Link speed in Mbps | [optional] [readonly] 
**State** | **string** | Operational state of the port. | [optional] [readonly] 
**Type** | **string** | Type of physical or virtual port | [optional] 
**Uuid** | **string** | Port UUID | [optional] [readonly] 
**Vlan** | [**PortVlan**](port_vlan.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


