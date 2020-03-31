# IscsiServiceTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | The iSCSI target alias of the iSCSI service.&lt;br/&gt; The target alias can contain one (1) to 128 characters and feature any printable character except space (\&quot; \&quot;). A PATCH request with an empty alias (\&quot;\&quot;) clears the alias.&lt;br/&gt; Optional in POST and PATCH. In POST, this defaults to the name of the SVM.  | [optional] 
**Name** | **string** | The iSCSI target name of the iSCSI service. This is generated for the SVM during POST.&lt;br/&gt; If required, the target name can be modified using the ONTAP command line.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


