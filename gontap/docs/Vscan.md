# Vscan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**CacheClear** | **bool** | Discards the cached information of the files that have been successfully scanned. Once the cache is cleared, files are scanned again when they are accessed. PATCH only | [optional] [default to null]
**Enabled** | **bool** | Specifies whether or not Vscan is enabled on the SVM. | [optional] [default to null]
**OnAccessPolicies** | [**[]VscanOnAccess**](vscan_on_access.md) |  | [optional] [default to null]
**OnDemandPolicies** | [**[]VscanOnDemand**](vscan_on_demand.md) |  | [optional] [default to null]
**ScannerPools** | [**[]VscanScannerPool**](vscan_scanner_pool.md) |  | [optional] [default to null]
**Svm** | [***FlexcacheRelationshipSvm**](flexcache_relationship_svm.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


