# FcPortFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** | Reports if the physical port has established a connection with the FC fabric.  | [optional] [readonly] 
**ConnectedSpeed** | **int32** | The negotiated data rate between the target FC port and the fabric in gigabits per second.  | [optional] [readonly] 
**Name** | **string** | The name of the fabric to which the port is connected. This is only available when the FC port is connected to a fabric.&lt;br/&gt; There is an added cost to retrieving this property&#39;s value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [readonly] 
**PortAddress** | **string** | The FC port address of the host bus adapter (HBA) physical port.&lt;br/&gt; Each FC port in an FC switched fabric has its own unique FC port address for routing purposes. The FC port address is assigned by a switch in the fabric when that port logs in to the fabric. This property refers to the FC port address given to the physical host bus adapter (HBA) port when the port performs a fabric login (FLOGI).&lt;br/&gt; This is useful for obtaining statistics and diagnostic information from FC switches.&lt;br/&gt; This is a six-digit hexadecimal encoded numeric value.  | [optional] [readonly] 
**SwitchPort** | **string** | The switch port to which the FC port is connected.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


