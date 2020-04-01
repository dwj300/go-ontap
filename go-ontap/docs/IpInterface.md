# IpInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Enabled** | **bool** | The administrative state of the interface. | [optional] [default to null]
**Ip** | [***IpInfo**](ip_info.md) |  | [optional] [default to null]
**Ipspace** | [***IpInterfaceIpspace**](ip_interface_ipspace.md) |  | [optional] [default to null]
**Location** | [***IpInterfaceLocation**](ip_interface_location.md) |  | [optional] [default to null]
**Name** | **string** | Interface name | [optional] [default to null]
**Scope** | **string** | Set to \&quot;svm\&quot; for interfaces owned by an SVM. Otherwise, set to \&quot;cluster\&quot;. | [optional] [default to null]
**ServicePolicy** | [***IpInterfaceServicePolicy**](ip_interface_service_policy.md) |  | [optional] [default to null]
**Services** | [**[]IpService**](ip_service.md) | The services associated with the interface. | [optional] [default to null]
**State** | **string** | The operational state of the interface. | [optional] [default to null]
**Svm** | [***IpInterfaceSvm1**](ip_interface_svm_1.md) |  | [optional] [default to null]
**Uuid** | **string** | The UUID that uniquely identifies the interface. | [optional] [default to null]
**Vip** | **bool** | True for a VIP interface, whose location is announced via BGP. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


