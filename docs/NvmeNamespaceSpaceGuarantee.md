# NvmeNamespaceSpaceGuarantee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | **bool** | The requested space reservation policy for the NVMe namespace. If _true_, a space reservation is requested for the namespace; if _false_, the namespace is thin provisioned. Guaranteeing a space reservation request for a namespace requires that the volume in which the namespace resides also be space reserved and that the fractional reserve for the volume be 100%.&lt;br/&gt; The space reservation policy for an NVMe namespace is determined by ONTAP.  | [optional] [readonly] 
**Reserved** | **bool** | Reports if the NVMe namespace is space guaranteed.&lt;br/&gt; This property is _true_ if a space guarantee is requested and the containing volume and aggregate support the request. This property is _false_ if a space guarantee is not requested or if a space guarantee is requested and either the containing volume and aggregate do not support the request.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


