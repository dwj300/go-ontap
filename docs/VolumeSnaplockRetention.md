# VolumeSnaplockRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **string** | Specifies the default retention period that is applied to files while committing them to the WORM state without an associated retention period. The retention value represents a duration and must be specified in the ISO-8601 duration format. The retention period can be in years, months, days, hours, and minutes. A duration specified for years, months, and days is represented in the ISO-8601 format as \&quot;P&lt;num&gt;Y\&quot;, \&quot;P&lt;num&gt;M\&quot;, \&quot;P&lt;num&gt;D\&quot; respectively, for example \&quot;P10Y\&quot; represents a duration of 10 years. A duration in hours and minutes is represented by \&quot;PT&lt;num&gt;H\&quot; and \&quot;PT&lt;num&gt;M\&quot; respectively. The retention string must contain only a single time element that is, either years, months, days, hours, or minutes. A duration which combines different periods is not supported, for example \&quot;P1Y10M\&quot; is not supported. Apart from the duration specified in the ISO-8601 format, the duration field also accepts the string \&quot;infinite\&quot; to set an infinite retention period. | [optional] 
**Maximum** | **string** | Specifies the maximum allowed retention period for files committed to the WORM state on the volume. The retention value represents a duration and must be specified in the ISO-8601 duration format. The retention period can be in years, months, days, hours, and minutes. A duration specified for years, months, and days is represented in the ISO-8601 format as \&quot;P&lt;num&gt;Y\&quot;, \&quot;P&lt;num&gt;M\&quot;, \&quot;P&lt;num&gt;D\&quot; respectively, for example \&quot;P10Y\&quot; represents a duration of 10 years. A duration in hours and minutes is represented by \&quot;PT&lt;num&gt;H\&quot; and \&quot;PT&lt;num&gt;M\&quot; respectively. The retention string must contain only a single time element that is, either years, months, days, hours, or minutes. A duration which combines different periods is not supported, for example \&quot;P1Y10M\&quot; is not supported. Apart from the duration specified in the ISO-8601 format, the duration field also accepts the string \&quot;infinite\&quot; to set an infinite retention period. | [optional] [default to P30Y]
**Minimum** | **string** | Specifies the minimum allowed retention period for files committed to the WORM state on the volume. The retention value represents a duration and must be specified in the ISO-8601 duration format. The retention period can be in years, months, days, hours, and minutes. A duration specified for years, month,s and days is represented in the ISO-8601 format as \&quot;P&lt;num&gt;Y\&quot;, \&quot;P&lt;num&gt;M\&quot;, \&quot;P&lt;num&gt;D\&quot; respectively, for example \&quot;P10Y\&quot; represents a duration of 10 years. A duration in hours and minutes is represented by \&quot;PT&lt;num&gt;H\&quot; and \&quot;PT&lt;num&gt;M\&quot; respectively. The retention string must contain only a single time element that is, either years, months, days, hours, or minutes. A duration which combines different periods is not supported, for example \&quot;P1Y10M\&quot; is not supported. Apart from the duration specified in the ISO-8601 format, the duration field also accepts the string \&quot;infinite\&quot; to set an infinite retention period. | [optional] [default to P0Y]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

