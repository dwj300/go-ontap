# FcLoginInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | **[]string** | The logged in initiator world wide port name (WWPN) aliases.  | [optional] [readonly] 
**PortAddress** | **string** | The port address of the initiator&#39;s FC port.&lt;br/&gt; Each port in an FC switched fabric has its own unique port address for routing purposes. The port address is assigned by a switch in the fabric when that port logs in to the fabric. This property refers to the address given by a switch to the initiator port.&lt;br/&gt; This is useful for obtaining statistics and diagnostic information from FC switches.&lt;br/&gt; This is a hexadecimal encoded numeric value.  | [optional] [readonly] 
**Wwnn** | **string** | The logged in initiator world wide node name (WWNN).  | [optional] [readonly] 
**Wwpn** | **string** | The logged in initiator WWPN.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


