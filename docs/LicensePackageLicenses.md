# LicensePackageLicenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | A flag indicating whether the license is currently being enforced. | [optional] [readonly] 
**Capacity** | [**LicenseCapacity**](license_capacity.md) |  | [optional] 
**Compliance** | [**LicenseCompliance**](license_compliance.md) |  | [optional] 
**Evaluation** | **bool** | A flag indicating whether the license is in evaluation mode. | [optional] [readonly] 
**ExpiryTime** | [**time.Time**](time.Time.md) | Date and time when the license expires. | [optional] [readonly] 
**Owner** | **string** | Cluster, node or license manager that owns the license. | [optional] [readonly] 
**SerialNumber** | **string** | Serial number of the license. | [optional] [readonly] 
**StartTime** | [**time.Time**](time.Time.md) | Date and time when the license starts. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


