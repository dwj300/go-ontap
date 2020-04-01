# VolumeNas

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportPolicy** | [***QtreeExportPolicy**](qtree_export_policy.md) |  | [optional] [default to null]
**Gid** | **int32** | The UNIX group ID of the volume. Valid in POST or PATCH. | [optional] [default to null]
**Path** | **string** | The fully-qualified path in the owning SVM&#39;s namespace at which the volume is mounted. The path is case insensitive and must be unique within a SVM&#39;s namespace. Path must begin with &#39;/&#39; and must not end with &#39;/&#39;. Only one volume can be mounted at any given junction path. An empty path in POST creates an unmounted volume. An empty path in PATCH deactivates and unmounts the volume. This attribute is reported in GET only when the volume is mounted. | [optional] [default to null]
**SecurityStyle** | **string** | Security style associated with the volume. Valid in POST or PATCH.&lt;br&gt;mixed &amp;dash; Mixed-style security&lt;br&gt;ntfs &amp;dash; NTFS/WIndows-style security&lt;br&gt;unified &amp;dash; Unified-style security, unified UNIX, NFS and CIFS permissions&lt;br&gt;unix &amp;dash; Unix-style security. | [optional] [default to null]
**Uid** | **int32** | The UNIX user ID of the volume. Valid in POST or PATCH. | [optional] [default to null]
**UnixPermissions** | **int32** | UNIX permissions to be viewed as an octal number. It consists of 4 digits derived by adding up bits 4 (read), 2 (write) and 1 (execute). First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. The second digit selects permission for the owner of the file; the third selects permissions for other users in the same group; the fourth for other users not in the group. Valid in POST or PATCH. For security style \&quot;mixed\&quot; or \&quot;unix\&quot;, the default setting is 0755 in octal (493 in decimal) and for security style \&quot;ntfs\&quot;, the default setting is 0000. In cases where only owner, group and other permissions are given (as in 755, representing the second, third and fourth dight), first digit is assumed to be zero. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


