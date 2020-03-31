# LunSpaceGuarantee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | **bool** | The requested space reservation policy for the LUN. If _true_, a space reservation is requested for the LUN; if _false_, the LUN is thin provisioned. Guaranteeing a space reservation request for a LUN requires that the volume in which the LUN resides is also space reserved and that the fractional reserve for the volume is 100%. Valid in POST and PATCH.  | [optional] [default to false]
**Reserved** | **bool** | Reports if the LUN is space guaranteed.&lt;br/&gt; If _true_, a space guarantee is requested and the containing volume and aggregate support the request. If _false_, a space guarantee is not requested or a space guarantee is requested and either the containing volume or aggregate do not support the request.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


