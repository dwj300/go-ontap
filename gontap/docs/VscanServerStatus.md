# VscanServerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisconnectedReason** | **string** | Specifies the server disconnected reason. The following is a list of the possible reasons: * unknown                   - Disconnected, unknown reason. * vscan_disabled            - Disconnected, Vscan is disabled on the SVM. * no_data_lif               - Disconnected, SVM does not have data LIF. * session_uninitialized     - Disconnected, session is not initialized. * remote_closed             - Disconnected, server has closed the connection. * invalid_protocol_msg      - Disconnected, invalid protocol message received. * invalid_session_id        - Disconnected, invalid session ID received. * inactive_connection       - Disconnected, no activity on connection. * invalid_user              - Connection request by an invalid user. * server_removed            - Disconnected, server has been removed from the active Scanners List. enum:   - unknown   - vscan_disabled   - no_data_lif   - session_uninitialized   - remote_closed   - invalid_protocol_msg   - invalid_session_id   - inactive_connection   - invalid_user   - server_removed  | [optional] [default to null]
**Ip** | **string** | IP address of the Vscan server. | [optional] [default to null]
**Node** | [***InlineResponse201Node**](inline_response_201_node.md) |  | [optional] [default to null]
**State** | **string** | Specifies the server connection state indicating if it is in the connected or disconnected state. The following is a list of the possible states: * connected                 - Connected * disconnected              - Disconnected enum:   - connected   - disconnected  | [optional] [default to null]
**Svm** | [***AuditSvm**](audit_svm.md) |  | [optional] [default to null]
**Type_** | **string** | Server type. The possible values are:   * primary - Primary server   * backup  - Backup server  | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | Specifies the time the server is in the connected or disconnected state. | [optional] [default to null]
**Vendor** | **string** | Name of the connected virus-scanner vendor. | [optional] [default to null]
**Version** | **string** | Version of the connected virus-scanner. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


