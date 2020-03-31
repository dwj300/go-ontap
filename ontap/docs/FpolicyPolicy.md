# FpolicyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Specifies if the policy is enabled on the SVM or not. If no value is mentioned for this field but priority is set, then this policy will be enabled.  | [optional] 
**Engine** | [**FpolicyEngineReference**](fpolicy_engine_reference.md) |  | [optional] 
**Events** | [**[]FpolicyEventReference**](fpolicy_event_reference.md) |  | [optional] 
**Mandatory** | **bool** | Specifies what action to take on a file access event in a case when all primary and secondary servers are down or no response is received from the FPolicy servers within a given timeout period. When this parameter is set to true, file access events will be denied under these circumstances. | [optional] [default to true]
**Name** | **string** | Specifies the name of the policy. | [optional] 
**Priority** | **int32** | Specifies the priority that is assigned to this policy. | [optional] 
**Scope** | [**FpolicyPolicyScope**](fpolicy_policy_scope.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


