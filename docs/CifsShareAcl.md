# CifsShareAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Permission** | **string** | Specifies the access rights that a user or group has on the defined CIFS Share. The following values are allowed: * no_access    - User does not have CIFS share access * read         - User has only read access * change       - User has change access * full_control - User has full_control access  | [optional] 
**Type** | **string** | Specifies the type of the user or group to add to the access control list of a CIFS share. The following values are allowed: * windows    - Windows user or group * unix_user  - UNIX user * unix_group - UNIX group  | [optional] [default to TYPE_WINDOWS]
**UserOrGroup** | **string** | Specifies the user or group name to add to the access control list of a CIFS share. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


