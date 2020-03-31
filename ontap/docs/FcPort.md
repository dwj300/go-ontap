# FcPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Description** | **string** | A description of the FC port.  | [optional] [readonly] 
**Enabled** | **bool** | The administrative state of the FC port. If this property is set to _false_, all FC connectivity to FC interfaces are blocked. Optional in PATCH.  | [optional] 
**Fabric** | [**FcPortFabric**](fc_port_fabric.md) |  | [optional] 
**Name** | **string** | The FC port name.  | [optional] [readonly] 
**Node** | [**InlineResponse201Node**](inline_response_201_node.md) |  | [optional] 
**PhysicalProtocol** | **string** | The physical network protocol of the FC port.  | [optional] [readonly] 
**Speed** | [**FcPortSpeed**](fc_port_speed.md) |  | [optional] 
**State** | **string** | The operational state of the FC port. - startup - The port is booting up. - link_not_connected - The port has finished initialization, but a link with the fabric is not established. - online - The port is initialized and a link with the fabric has been established. - link_disconnected - The link was present at one point on this port but is currently not established. - offlined_by_user - The port is administratively disabled. - offlined_by_system - The port is set to offline by the system. This happens when the port encounters too many errors. - node_offline - The state information for the port cannot be retrieved. The node is offline or inaccessible.  | [optional] [readonly] 
**SupportedProtocols** | **[]string** | The network protocols supported by the FC port.  | [optional] [readonly] 
**Transceiver** | [**FcPortTransceiver**](fc_port_transceiver.md) |  | [optional] 
**Uuid** | **string** | The unique identifier of the FC port.  | [optional] [readonly] 
**Wwnn** | **string** | The base world wide node name (WWNN) for the FC port.  | [optional] [readonly] 
**Wwpn** | **string** | The base world wide port name (WWPN) for the FC port.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


