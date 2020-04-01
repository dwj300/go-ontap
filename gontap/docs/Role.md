# Role

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Builtin** | **bool** | Indicates if this is a built-in (pre-defined) role which cannot be modified or deleted. | [optional] [default to null]
**Name** | **string** | Role name | [optional] [default to null]
**Owner** | [***RoleOwner**](role_owner.md) |  | [optional] [default to null]
**Privileges** | [**[]RolePrivilege**](role_privilege.md) | The list of privileges that this role has been granted. | [optional] [default to null]
**Scope** | **string** | Scope of the entity. set to \&quot;cluster\&quot; for cluster owned objects and to \&quot;svm\&quot; for SVM owned objects. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


