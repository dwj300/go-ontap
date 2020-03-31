# IpInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Enabled** | **bool** | The administrative state of the interface. | [optional] [default to true]
**Ip** | [**IpInfo**](ip_info.md) |  | [optional] 
**Ipspace** | [**IpInterfaceIpspace**](ip_interface_ipspace.md) |  | [optional] 
**Location** | [**IpInterfaceLocation**](ip_interface_location.md) |  | [optional] 
**Name** | **string** | Interface name | [optional] 
**Scope** | **string** | Set to \&quot;svm\&quot; for interfaces owned by an SVM. Otherwise, set to \&quot;cluster\&quot;. | [optional] 
**ServicePolicy** | [**IpInterfaceServicePolicy**](ip_interface_service_policy.md) |  | [optional] 
**Services** | [**[]IpService**](ip_service.md) | The services associated with the interface. | [optional] [readonly] 
**State** | **string** | The operational state of the interface. | [optional] [readonly] 
**Svm** | [**IpInterfaceSvm1**](ip_interface_svm_1.md) |  | [optional] 
**Uuid** | **string** | The UUID that uniquely identifies the interface. | [optional] [readonly] 
**Vip** | **bool** | True for a VIP interface, whose location is announced via BGP. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


