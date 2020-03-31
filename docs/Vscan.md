# Vscan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**InlineResponse201Links**](inline_response_201__links.md) |  | [optional] 
**CacheClear** | **bool** | Discards the cached information of the files that have been successfully scanned. Once the cache is cleared, files are scanned again when they are accessed. PATCH only | [optional] [default to false]
**Enabled** | **bool** | Specifies whether or not Vscan is enabled on the SVM. | [optional] [default to true]
**OnAccessPolicies** | [**[]VscanOnAccess**](vscan_on_access.md) |  | [optional] 
**OnDemandPolicies** | [**[]VscanOnDemand**](vscan_on_demand.md) |  | [optional] 
**ScannerPools** | [**[]VscanScannerPool**](vscan_scanner_pool.md) |  | [optional] 
**Svm** | [**AuditSvm**](audit_svm.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


