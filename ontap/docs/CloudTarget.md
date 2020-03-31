# CloudTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**AccessKey** | **string** | Access key ID for AWS_S3 and other S3 compatible provider types. | [optional] 
**AuthenticationType** | **string** | Authentication used to access the target. Snapmirror does not yet support CAP. Required in POST. | [optional] 
**AzureAccount** | **string** | Azure account | [optional] 
**AzurePrivateKey** | **string** | Azure access key | [optional] 
**CapUrl** | **string** | This parameter is available only when auth-type is CAP. It specifies a full URL of the request to a CAP server for retrieving temporary credentials (access-key, secret-pasword, and session token) for accessing the object store. | [optional] 
**CertificateValidationEnabled** | **bool** | Is SSL/TLS certificate validation enabled? The default value is true. This can only be modified for SGWS and IBM_COS provider types. | [optional] 
**Container** | **string** | Data bucket/container name | [optional] 
**Ipspace** | [**CloudTargetIpspace**](cloud_target_ipspace.md) |  | [optional] 
**Name** | **string** | Cloud target name | [optional] 
**Owner** | **string** | Owner of the target. Allowed values are FabricPool or SnapMirror. A target can be used by only one feature. | [optional] 
**Port** | **int32** | Port number of the object store that ONTAP uses when establishing a connection. Required in POST. | [optional] 
**ProviderType** | **string** | Type of cloud provider. Allowed values depend on owner type. For FabricPool, AliCloud, AWS_S3, Azure_Cloud, GoggleCloud, IBM_COS, and SGWS are allowed. For SnapMirror, the valid values are AWS_S3 or SGWS. | [optional] 
**SecretPassword** | **string** | Secret access key for AWS_S3 and other S3 compatible provider types. | [optional] 
**Server** | **string** | Fully qualified domain name of the object store server. Required on POST.  For Amazon S3, server name must be an AWS regional endpoint in the format s3.amazonaws.com or s3-&lt;region&gt;.amazonaws.com, for example, s3-us-west-2.amazonaws.com. The region of the server and the bucket must match. For Azure, if the server is a \&quot;blob.core.windows.net\&quot; or a \&quot;blob.core.usgovcloudapi.net\&quot;, then a value of azure-account followed by a period is added in front of the server. | [optional] 
**SnapmirrorUse** | **string** | Use of the cloud target by SnapMirror. | [optional] 
**SslEnabled** | **bool** | SSL/HTTPS enabled or not | [optional] [default to true]
**Svm** | [**CloudTargetSvm**](cloud_target_svm.md) |  | [optional] 
**Used** | **int32** | The amount of cloud space used by all the aggregates attached to the target, in bytes. This field is only populated for FabricPool targets. The value is recalculated once every 5 minutes. | [optional] [readonly] 
**Uuid** | **string** | Cloud target UUID | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


