# License

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | A flag indicating whether the license is currently being enforced. | [optional] [default to null]
**Capacity** | [***LicenseCapacity**](license_capacity.md) |  | [optional] [default to null]
**Compliance** | [***LicenseCompliance**](license_compliance.md) |  | [optional] [default to null]
**Evaluation** | **bool** | A flag indicating whether the license is in evaluation mode. | [optional] [default to null]
**ExpiryTime** | [**time.Time**](time.Time.md) | Date and time when the license expires. | [optional] [default to null]
**Owner** | **string** | Cluster, node or license manager that owns the license. | [optional] [default to null]
**SerialNumber** | **string** | Serial number of the license. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | Date and time when the license starts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


