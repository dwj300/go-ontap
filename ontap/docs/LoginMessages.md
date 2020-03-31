# LoginMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**Banner** | **string** | The login banner text. This message is displayed during SSH and console device login just before the password prompt displays. When configured, a cluster-level login banner is used for every incoming connection. Each data SVM can override the cluster-level banner to instead display when you log into the SVM. To restore the default setting for a data SVM, set the banner to an empty string. New lines are supplied as either LF or CRLF but are always returned as LF. Optional in the PATCH body.  | [optional] 
**Message** | **string** | The message of the day (MOTD). This message appears just before the clustershell prompt after a successful login. When configured, the cluster message displays first. If you log in as a data SVM administrator, the SVM message is then printed. The cluster-level MOTD can be disabled for a given data SVM using the \&quot;show_cluster_message\&quot; property. New lines are supplied as either LF or CRLF but are always returned as LF. Optional in the PATCH body.  | [optional] 
**Scope** | [**NetworkScope**](network_scope.md) |  | [optional] 
**ShowClusterMessage** | **bool** | Specifies whether to show a cluster-level message before the SVM message when logging in as an SVM administrator. This setting can only be modified by the cluster administrator. Optional in the PATCH body.  | [optional] [default to true]
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 
**Uuid** | **string** | The unique identifier (ID) of the login messages configuration.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


