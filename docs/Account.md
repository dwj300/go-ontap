# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Applications** | [**[]AccountApplication**](account_application.md) |  | [optional] 
**Comment** | **string** | Optional comment for the user account. | [optional] 
**Locked** | **bool** | Locked status of the account. | [optional] [default to false]
**Name** | **string** | User or group account name | [optional] 
**Owner** | [**AccountOwner**](account_owner.md) |  | [optional] 
**Password** | **string** | Password for the account. The password can contain a mix of lower and upper case alphabetic characters, digits, and special characters. | [optional] 
**Role** | [**RoleReference**](role_reference.md) |  | [optional] 
**Scope** | **string** | Scope of the entity. set to \&quot;cluster\&quot; for cluster owned objects and to \&quot;svm\&quot; for SVM owned objects. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


