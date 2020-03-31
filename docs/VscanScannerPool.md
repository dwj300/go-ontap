# VscanScannerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterReference**](cluster_reference.md) |  | [optional] 
**Name** | **string** | Specifies the name of the scanner pool. Scanner pool name can be up to 256 characters long and is a string that can only contain any combination of ASCII-range alphanumeric characters a-z, A-Z, 0-9), \&quot;_\&quot;, \&quot;-\&quot; and \&quot;.\&quot;. | [optional] 
**PrivilegedUsers** | **[]string** | Specifies a list of privileged users. A valid form of privileged user-name is \&quot;domain-name\\user-name\&quot;. Privileged user-names are stored and treated as case-insensitive strings. Virus scanners must use one of the registered privileged users for connecting to clustered Data ONTAP for exchanging virus-scanning protocol messages and to access file for scanning, remedying and quarantining operations. | [optional] 
**Role** | **string** | Specifies the role of the scanner pool. The possible values are:   * primary   - Always active.   * secondary - Active only when none of the primary external virus-scanning servers are connected.   * idle      - Always inactive.  | [optional] [default to ROLE_PRIMARY]
**Servers** | **[]string** | Specifies a list of IP addresses or FQDN for each Vscan server host names which are allowed to connect to clustered ONTAP. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


