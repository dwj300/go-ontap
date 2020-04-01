# Application

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Svm** | [***ApplicationSvm**](application_svm.md) |  | [optional] [default to null]
**Uuid** | **string** | Application UUID. This field is generated when the application is created. Required in the URL | [optional] [default to null]
**Name** | **string** | Application Name. This field is user supplied when the application is created. Required in the POST body | [optional] [default to null]
**CreationTimestamp** | **string** | The time when the application was created | [optional] [default to null]
**Generation** | **int32** | The generation number of the application. This indicates which features are supported on the application. For example, generation 1 applications do not support snapshots. Support for snapshots was added at generation 2. Any future generation numbers and their feature set will be documented | [optional] [default to null]
**ProtectionGranularity** | **string** | Protection granularity determines the scope of Snapshot operations for the application. Possible values are \&quot;application\&quot; and \&quot;component\&quot;. If the value is \&quot;application\&quot;, Snapshot operations are performed on the entire application. If the value is \&quot;component\&quot;, Snapshot operations are performed separately on the application components | [optional] [default to null]
**Rpo** | [***ApplicationRpo**](application_rpo.md) |  | [optional] [default to null]
**State** | **string** | The state of the application. For full functionality, applications must be in the online state. Other states indicate that the application is in a transient state and not all operations are supported | [optional] [default to null]
**Statistics** | [***ApplicationStatistics**](application_statistics.md) |  | [optional] [default to null]
**Template** | [***ApplicationTemplate1**](application_template_1.md) |  | [optional] [default to null]
**MaxdataOnSan** | [***MaxdataOnSan**](maxdata_on_san.md) |  | [optional] [default to null]
**MongoDbOnSan** | [***MongoDbOnSan**](mongo_db_on_san.md) |  | [optional] [default to null]
**Nas** | [***Nas**](nas.md) |  | [optional] [default to null]
**OracleOnNfs** | [***OracleOnNfs**](oracle_on_nfs.md) |  | [optional] [default to null]
**OracleOnSan** | [***OracleOnSan**](oracle_on_san.md) |  | [optional] [default to null]
**OracleRacOnNfs** | [***OracleRacOnNfs**](oracle_rac_on_nfs.md) |  | [optional] [default to null]
**OracleRacOnSan** | [***OracleRacOnSan**](oracle_rac_on_san.md) |  | [optional] [default to null]
**San** | [***San**](san.md) |  | [optional] [default to null]
**SqlOnSan** | [***SqlOnSan**](sql_on_san.md) |  | [optional] [default to null]
**SqlOnSmb** | [***SqlOnSmb**](sql_on_smb.md) |  | [optional] [default to null]
**VdiOnNas** | [***VdiOnNas**](vdi_on_nas.md) |  | [optional] [default to null]
**VdiOnSan** | [***VdiOnSan**](vdi_on_san.md) |  | [optional] [default to null]
**VsiOnNas** | [***VsiOnNas**](vsi_on_nas.md) |  | [optional] [default to null]
**VsiOnSan** | [***VsiOnSan**](vsi_on_san.md) |  | [optional] [default to null]
**Links** | [***ApplicationLinks**](application__links.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


