# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Builtin** | **bool** | Indicates if this is a built-in (pre-defined) role which cannot be modified or deleted. | [optional] [readonly] 
**Name** | **string** | Role name | [optional] 
**Owner** | [**RoleOwner**](role_owner.md) |  | [optional] 
**Privileges** | [**[]RolePrivilege**](role_privilege.md) | The list of privileges that this role has been granted. | [optional] 
**Scope** | **string** | Scope of the entity. set to \&quot;cluster\&quot; for cluster owned objects and to \&quot;svm\&quot; for SVM owned objects. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


