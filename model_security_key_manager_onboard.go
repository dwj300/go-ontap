/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// SecurityKeyManagerOnboard Configures onboard key management. After configuring onboard key management, save the encrypted configuration data in a safe location so that you can use it if you need to perform a manual recovery operation.
type SecurityKeyManagerOnboard struct {
	// Is the onboard key manager enabled?
	Enabled bool `json:"enabled,omitempty"`
	// The cluster-wide passphrase. This is not audited.
	ExistingPassphrase string `json:"existing_passphrase,omitempty"`
	// The cluster-wide passphrase. This is not audited.
	Passphrase string `json:"passphrase,omitempty"`
}
