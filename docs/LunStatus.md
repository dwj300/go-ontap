# LunStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerState** | **string** | The state of the volume and aggregate that contain the LUN. LUNs are only available when their containers are available.  | [optional] [readonly] 
**Mapped** | **bool** | Reports if the LUN is mapped to one or more initiator groups.&lt;br/&gt; There is an added cost to retrieving this property&#39;s value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the &#x60;fields&#x60; query parameter. See [&#x60;DOC Requesting specific fields&#x60;](#docs-docs-Requesting-specific-fields) to learn more.  | [optional] [readonly] 
**ReadOnly** | **bool** | Reports if the LUN allows only read access.  | [optional] [readonly] 
**State** | **string** | The state of the LUN. Normal states for a LUN are _online_ and _offline_. Other states indicate errors.  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


