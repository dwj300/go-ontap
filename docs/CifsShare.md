# CifsShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**AccessBasedEnumeration** | **bool** | If enabled, all folders inside this share are visible to a user based on that individual user access right; prevents the display of folders or other shared resources that the user does not have access to.  | [optional] [default to false]
**Acls** | [**[]CifsShareAcl**](cifs_share_acl.md) |  | [optional] 
**ChangeNotify** | **bool** | Specifies whether CIFS clients can request for change notifications for directories on this share. | [optional] [default to true]
**Comment** | **string** | Specify the CIFS share descriptions. | [optional] 
**Encryption** | **bool** | Specifies that SMB encryption must be used when accessing this share. Clients that do not support encryption are not able to access this share.  | [optional] [default to false]
**HomeDirectory** | **bool** | Specifies whether or not the share is a home directory share, where the share and path names are dynamic. ONTAP home directory functionality automatically offer each user a dynamic share to their home directory without creating an individual SMB share for each user. The ONTAP CIFS home directory feature enable us to configure a share that maps to different directories based on the user that connects to it. Instead of creating a separate shares for each user, a single share with a home directory parameters can be created. In a home directory share, ONTAP dynamically generates the share-name and share-path by substituting %w, %u, and %d variables with the corresponding Windows user name, UNIX user name, and domain name, respectively.  | [optional] [default to false]
**Name** | **string** | Specifies the name of the CIFS share that you want to create. If this is a home directory share then the share name includes the pattern as %w (Windows user name), %u (UNIX user name) and %d (Windows domain name) variables in any combination with this parameter to generate shares dynamically.  | [optional] 
**Oplocks** | **bool** | Specify whether opportunistic locks are enabled on this share. \&quot;Oplocks\&quot; allow clients to lock files and cache content locally, which can increase performance for file operations.  | [optional] [default to true]
**Path** | **string** | The fully-qualified pathname in the owning SVM namespace that is shared through this share. If this is a home directory share then the path should be dynamic by specifying the pattern %w (Windows user name), %u (UNIX user name), or %d (domain name) variables in any combination. ONTAP generates the path dynamically for the connected user and this path is appended to each search path to find the full Home Directory path.  | [optional] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**UnixSymlink** | **string** | Controls the access of UNIX symbolic links to CIFS clients. The supported values are:     * local - Enables only local symbolic links which is within the same CIFS share.     * widelink - Enables both local symlinks and widelinks.     * disable - Disables local symlinks and widelinks.  | [optional] [default to UNIX_SYMLINK_LOCAL]
**Volume** | [**CifsShareVolume**](cifs_share_volume.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


