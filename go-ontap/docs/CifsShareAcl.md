# CifsShareAcl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Permission** | **string** | Specifies the access rights that a user or group has on the defined CIFS Share. The following values are allowed: * no_access    - User does not have CIFS share access * read         - User has only read access * change       - User has change access * full_control - User has full_control access  | [optional] [default to null]
**Type_** | **string** | Specifies the type of the user or group to add to the access control list of a CIFS share. The following values are allowed: * windows    - Windows user or group * unix_user  - UNIX user * unix_group - UNIX group  | [optional] [default to null]
**UserOrGroup** | **string** | Specifies the user or group name to add to the access control list of a CIFS share. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


