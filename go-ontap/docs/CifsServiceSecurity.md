# CifsServiceSecurity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KdcEncryption** | **bool** | Specifies whether AES-128 and AES-256 encryption is enabled for all Kerberos-based communication with the Active Directory KDC. To take advantage of the strongest security with Kerberos-based communication, AES-256 and AES-128 encryption can be enabled on the CIFS server. Kerberos-related communication for CIFS is used during CIFS server creation on the SVM, as well as during the SMB session setup phase. The CIFS server supports the following encryption types for Kerberos communication:     * RC4-HMAC     * DES     * AES When the CIFS server is created, the domain controller creates a computer machine account in Active Directory. After a newly created machine account authenticates, the KDC and the CIFS server negotiates encryption types. At this time, the KDC becomes aware of the encryption capabilities of the particular machine account and uses those capabilities in subsequent communication with the CIFS server. In addition to negotiating encryption types during CIFS server creation, the encryption types are renegotiated when a machine account password is reset.  | [optional] [default to null]
**RestrictAnonymous** | **string** | Specifies what level of access an anonymous user is granted. An anonymous user (also known as a \&quot;null user\&quot;) can list or enumerate certain types of system information from Windows hosts on the network, including user names and details, account policies, and share names. Access for the anonymous user can be controlled by specifying one of three access restriction settings.  The available values are:  * no_restriction   - No access restriction for an anonymous user.  * no_enumeration   - Enumeration is restricted for an anonymous user.  * no_access        - All access is restricted for an anonymous user.  | [optional] [default to null]
**SmbEncryption** | **bool** | Specifies whether encryption is required for incoming CIFS traffic. | [optional] [default to null]
**SmbSigning** | **bool** | Specifies whether signing is required for incoming CIFS traffic. SMB signing helps to ensure that network traffic between the CIFS server and the client is not compromised.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


