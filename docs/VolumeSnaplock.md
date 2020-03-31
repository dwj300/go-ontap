# VolumeSnaplock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppendModeEnabled** | **bool** | Specifies if the volume append mode is enabled or disabled. When it is enabled, all the files created with write permissions on the volume are, by default, WORM appendable files. The user can append the data to a WORM appendable file but cannot modify the existing contents of the file nor delete the file until it expires. | [optional] [default to false]
**AutocommitPeriod** | **string** | Specifies the autocommit period for SnapLock volume. All files which are not modified for a period greater than the autocommit period of the volume are committed to the WORM state. The autocommit period value represents a duration and must be specified in the ISO-8601 duration format. The autocommit period can be in years, months, days, hours, and minutes. A period specified for years, months, and days is represented in the ISO-8601 format as \&quot;P&lt;num&gt;Y\&quot;, \&quot;P&lt;num&gt;M\&quot;, \&quot;P&lt;num&gt;D\&quot; respectively, for example \&quot;P10Y\&quot; represents a duration of 10 years. A duration in hours and minutes is represented by \&quot;PT&lt;num&gt;H\&quot; and \&quot;PT&lt;num&gt;M\&quot; respectively. The period string must contain only a single time element that is, either years, months, days, hours, or minutes. A duration which combines different periods is not supported, for example \&quot;P1Y10M\&quot; is not supported. Apart from the duration specified in the ISO-8601 format, the autocommit field also accepts the string \&quot;none\&quot;. | [optional] [default to none]
**ComplianceClockTime** | [**time.Time**](time.Time.md) | This is the volume compliance clock time which is used to manage the SnapLock objects in the volume. | [optional] [readonly] 
**ExpiryTime** | [**time.Time**](time.Time.md) | Expiry time of the volume. | [optional] [readonly] 
**IsAuditLog** | **bool** | Indicates if this volume has been configured as SnapLock audit log volume for the SVM . | [optional] [readonly] 
**LitigationCount** | **int32** | Litigation count indicates the number of active legal-holds on the volume. | [optional] [readonly] 
**PrivilegedDelete** | **string** | Specifies the privileged-delete attribute of a SnapLock volume. On a SnapLock Enterprise (SLE) volume, a designated privileged user can selectively delete files irrespective of the retention time of the file. SLE volumes can have privileged delete as disabled, enabled or permanently_disabled and for SnapLock Compliance (SLC) volumes it is always permanently_disabled. | [optional] 
**Retention** | [**VolumeSnaplockRetention**](volume_snaplock_retention.md) |  | [optional] 
**Type** | **string** | The SnapLock type of the volume. &lt;br&gt;compliance &amp;dash; A SnapLock Compliance(SLC) volume provides the highest level of WORM protection and an administrator cannot destroy a SLC volume if it contains unexpired WORM files. &lt;br&gt; enterprise &amp;dash; An administrator can delete a SnapLock Enterprise(SLE) volume.&lt;br&gt; non_snaplock &amp;dash; Indicates the volume is non-snaplock. | [optional] [readonly] [default to TYPE_NON_SNAPLOCK]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


