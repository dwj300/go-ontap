# Volume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***InlineResponse201Links**](inline_response_201__links.md) |  | [optional] [default to null]
**Aggregates** | [**[]DiskAggregates**](disk_aggregates.md) | Aggregate hosting the volume. Required on POST. | [optional] [default to null]
**Application** | [***VolumeApplication**](volume_application.md) |  | [optional] [default to null]
**Autosize** | [***VolumeAutosize**](volume_autosize.md) |  | [optional] [default to null]
**Clone** | [***VolumeClone**](volume_clone.md) |  | [optional] [default to null]
**Comment** | **string** | A comment for the volume. Valid in POST or PATCH. | [optional] [default to null]
**ConstituentsPerAggregate** | **int32** | Specifies the number of times to iterate over the aggregates listed with the \&quot;aggregates.name\&quot; or \&quot;aggregates.uuid\&quot; when creating or expanding a FlexGroup. If a volume is being created on a single aggregate, the system will create a flexible volume if the \&quot;constituents_per_aggregate\&quot; field is not specified, and a FlexGroup if it is specified.  If a volume is being created on multiple aggregates, the system will always create a FlexGroup. | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | Creation time of the volume. This field is generated when the volume is created. | [optional] [default to null]
**Efficiency** | [***VolumeEfficiency**](volume_efficiency.md) |  | [optional] [default to null]
**Encryption** | [***VolumeEncryption**](volume_encryption.md) |  | [optional] [default to null]
**ErrorState** | [***VolumeErrorState**](volume_error_state.md) |  | [optional] [default to null]
**Files** | [***VolumeFiles**](volume_files.md) |  | [optional] [default to null]
**FlexcacheEndpointType** | **string** | FlexCache endpoint type. &lt;br&gt;none &amp;dash; The volume is neither a FlexCache nor origin of any FlexCache. &lt;br&gt;cache &amp;dash; The volume is a FlexCache volume. &lt;br&gt;origin &amp;dash; The volume is origin of a FlexCache volume. | [optional] [default to null]
**Guarantee** | [***VolumeGuarantee**](volume_guarantee.md) |  | [optional] [default to null]
**Language** | **string** | Language encoding setting for volume. If no language is specified, the volume inherits its SVM language encoding setting. | [optional] [default to null]
**Metric** | [***ClusterMetric**](cluster_metric.md) |  | [optional] [default to null]
**Movement** | [***VolumeMovement**](volume_movement.md) |  | [optional] [default to null]
**Name** | **string** | Volume name. The name of volume must start with an alphabetic character (a to z or A to Z) or an underscore (_). The name must be 197 or fewer characters in length for FlexGroups, and 203 or fewer characters in length for all other types of volumes. Volume names must be unique within an SVM. Required on POST. | [optional] [default to null]
**Nas** | [***VolumeNas**](volume_nas.md) |  | [optional] [default to null]
**Qos** | [***VolumeQos**](volume_qos.md) |  | [optional] [default to null]
**Quota** | [***VolumeQuota**](volume_quota.md) |  | [optional] [default to null]
**Size** | **int32** | Physical size of the volume. The minimum size for a FlexVol volume is 20MB and the minimum size for a FlexGroup volume is 200MB per constituent. The recommended size for a FlexGroup volume is a minimum of 100GB per constituent. For all volumes, the default size is equal to the minimum size. | [optional] [default to null]
**Snaplock** | [***VolumeSnaplock**](volume_snaplock.md) |  | [optional] [default to null]
**SnapshotPolicy** | [***SvmSnapshotPolicy**](svm_snapshot_policy.md) |  | [optional] [default to null]
**Space** | [***VolumeSpace**](volume_space.md) |  | [optional] [default to null]
**State** | **string** | Volume state. A volume can only be brought online if it is offline. The &#39;mixed&#39; state applies to FlexGroup volumes only and cannot be specified as a target state. An &#39;error&#39; state implies that the volume is not in a state to serve data. | [optional] [default to null]
**Statistics** | [***ClusterStatistics**](cluster_statistics.md) |  | [optional] [default to null]
**Style** | **string** | The style of the volume. If \&quot;style\&quot; is not specified, the volume type is determined based on the specified aggregates. Specifying a single aggregate, without \&quot;constituents_per_aggregate\&quot; creates a flexible volume. Specifying multiple aggregates, or a single aggregate with \&quot;constituents_per_aggregate\&quot; creates a FlexGroup. If \&quot;style\&quot; is specified, a volume of that type is created. That is, if style is \&quot;flexvol\&quot;, a single aggregate must be specified. If style is \&quot;flexgroup\&quot;, the system either uses the specified aggregates, or automatically provisions if no aggregates are specified.&lt;br&gt;flexvol &amp;dash; flexible volumes and FlexClone volumes&lt;br&gt;flexgroup &amp;dash; FlexGroups. | [optional] [default to null]
**Svm** | [***VolumeSvm**](volume_svm.md) |  | [optional] [default to null]
**Tiering** | [***VolumeTiering**](volume_tiering.md) |  | [optional] [default to null]
**Type_** | **string** | Type of the volume.&lt;br&gt;rw &amp;dash; read-write volume.&lt;br&gt;dp &amp;dash; data-protection volume.&lt;br&gt;ls &amp;dash; load-sharing &#x60;dp&#x60; volume. Valid in GET. | [optional] [default to null]
**UseMirroredAggregates** | **bool** | Specifies whether mirrored aggregates are selected when provisioning a FlexGroup without specifying \&quot;aggregates.name\&quot; or \&quot;aggregates.uuid\&quot;. Only mirrored aggregates are used if this parameter is set to &#39;true&#39; and only unmirrored aggregates are used if this parameter is set to &#39;false&#39;. Aggregate level mirroring for a FlexGroup can be changed by moving all of the constituents to the required aggregates. The default value is &#39;true&#39; for a MetroCluster configuration and is &#39;false&#39; for a non-MetroCluster configuration. | [optional] [default to null]
**Uuid** | **string** | Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


