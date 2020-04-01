# PortLag

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePorts** | [**[]IpInterfaceLocationHomePort**](ip_interface_location_home_port.md) | Active ports of a LAG (ifgrp). (Some member ports may be inactive.) | [optional] [default to null]
**DistributionPolicy** | **string** | Policy for mapping flows to ports for outbound packets through a LAG (ifgrp). | [optional] [default to null]
**MemberPorts** | [**[]IpInterfaceLocationHomePort**](ip_interface_location_home_port.md) |  | [optional] [default to null]
**Mode** | **string** | Determines how the ports interact with the switch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


