# FpolicyPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Specifies if the policy is enabled on the SVM or not. If no value is mentioned for this field but priority is set, then this policy will be enabled.  | [optional] [default to null]
**Engine** | [***FpolicyEngineReference**](fpolicy_engine_reference.md) |  | [optional] [default to null]
**Events** | [**[]FpolicyEventReference**](fpolicy_event_reference.md) |  | [optional] [default to null]
**Mandatory** | **bool** | Specifies what action to take on a file access event in a case when all primary and secondary servers are down or no response is received from the FPolicy servers within a given timeout period. When this parameter is set to true, file access events will be denied under these circumstances. | [optional] [default to null]
**Name** | **string** | Specifies the name of the policy. | [optional] [default to null]
**Priority** | **int32** | Specifies the priority that is assigned to this policy. | [optional] [default to null]
**Scope** | [***FpolicyPolicyScope**](fpolicy_policy_scope.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


