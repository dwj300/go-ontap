# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Svm** | [**ApplicationSvm**](application_svm.md) |  | [optional] 
**Uuid** | **string** | Application UUID. This field is generated when the application is created. Required in the URL | [optional] [readonly] 
**Name** | **string** | Application Name. This field is user supplied when the application is created. Required in the POST body | [optional] 
**CreationTimestamp** | **string** | The time when the application was created | [optional] [readonly] 
**Generation** | **int32** | The generation number of the application. This indicates which features are supported on the application. For example, generation 1 applications do not support snapshots. Support for snapshots was added at generation 2. Any future generation numbers and their feature set will be documented | [optional] [readonly] 
**ProtectionGranularity** | **string** | Protection granularity determines the scope of Snapshot operations for the application. Possible values are \&quot;application\&quot; and \&quot;component\&quot;. If the value is \&quot;application\&quot;, Snapshot operations are performed on the entire application. If the value is \&quot;component\&quot;, Snapshot operations are performed separately on the application components | [optional] [readonly] 
**Rpo** | [**ApplicationRpo**](application_rpo.md) |  | [optional] 
**State** | **string** | The state of the application. For full functionality, applications must be in the online state. Other states indicate that the application is in a transient state and not all operations are supported | [optional] [readonly] 
**Statistics** | [**ApplicationStatistics**](application_statistics.md) |  | [optional] 
**Template** | [**ApplicationTemplate1**](application_template_1.md) |  | [optional] 
**MaxdataOnSan** | [**MaxdataOnSan**](maxdata_on_san.md) |  | [optional] 
**MongoDbOnSan** | [**MongoDbOnSan**](mongo_db_on_san.md) |  | [optional] 
**Nas** | [**Nas**](nas.md) |  | [optional] 
**OracleOnNfs** | [**OracleOnNfs**](oracle_on_nfs.md) |  | [optional] 
**OracleOnSan** | [**OracleOnSan**](oracle_on_san.md) |  | [optional] 
**OracleRacOnNfs** | [**OracleRacOnNfs**](oracle_rac_on_nfs.md) |  | [optional] 
**OracleRacOnSan** | [**OracleRacOnSan**](oracle_rac_on_san.md) |  | [optional] 
**San** | [**San**](san.md) |  | [optional] 
**SqlOnSan** | [**SqlOnSan**](sql_on_san.md) |  | [optional] 
**SqlOnSmb** | [**SqlOnSmb**](sql_on_smb.md) |  | [optional] 
**VdiOnNas** | [**VdiOnNas**](vdi_on_nas.md) |  | [optional] 
**VdiOnSan** | [**VdiOnSan**](vdi_on_san.md) |  | [optional] 
**VsiOnNas** | [**VsiOnNas**](vsi_on_nas.md) |  | [optional] 
**VsiOnSan** | [**VsiOnSan**](vsi_on_san.md) |  | [optional] 
**Links** | [**ApplicationLinks**](application__links.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


