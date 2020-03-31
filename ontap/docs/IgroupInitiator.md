# IgroupInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Igroup** | [**IgroupInitiatorIgroup**](igroup_initiator_igroup.md) |  | [optional] 
**Name** | **string** | The FC WWPN, iSCSI IQN, or iSCSI EUI that identifies the host initiator. Valid in POST only and not allowed when the &#x60;records&#x60; property is used.&lt;br/&gt; An FC WWPN consist of 16 hexadecimal digits grouped as 8 pairs separated by colons. The format for an iSCSI IQN is _iqn.yyyy-mm.reverse_domain_name:any_. The iSCSI EUI format consists of the _eui._ prefix followed by 16 hexadecimal characters.  | [optional] 
**Records** | [**[]IgroupInitiatorRecords**](igroup_initiator_records.md) | An array of initiators specified to add multiple initiators to an initiator group in a single API call. Valid in POST only and not allowed when the &#x60;name&#x60; property is used.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


