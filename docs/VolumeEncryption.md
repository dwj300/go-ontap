# VolumeEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Encrypts an unencrypted volume. When set to &#39;true&#39;, a new key is generated and used to encrypt the given volume. The underlying SVM must be configured with the key manager. | [optional] 
**KeyId** | **string** | The key ID used for creating encrypted volume. A new key-id is generated for creating an encrypted volume. This key-id is associated with the generated key. | [optional] [readonly] 
**Rekey** | **bool** | If set to &#39;true&#39;, re-encrypts the volume with a new key. Valid in PATCH. | [optional] 
**State** | **string** | Volume encryption state.&lt;br&gt;encrypted &amp;dash; The volume is completely encrypted.&lt;br&gt;encrypting &amp;dash; Encryption operation is in progress.&lt;br&gt;partial &amp;dash; Some constituents are encrypted and some are not. Applicable only for FlexGroup volume.&lt;br&gt;rekeying. Encryption of volume with a new key is in progress.&lt;br&gt;unencrypted &amp;dash; The volume is a plain-text one. | [optional] [readonly] 
**Status** | [**VolumeEncryptionStatus**](volume_encryption_status.md) |  | [optional] 
**Type** | **string** | Volume encryption type.&lt;br&gt;none &amp;dash; The volume is a plain-text one.&lt;br&gt;volume &amp;dash; The volume is encrypted with volume key (NVE volume).&lt;br&gt;aggregate &amp;dash; The volume is encrypted with aggregate key (NAE volume). | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


