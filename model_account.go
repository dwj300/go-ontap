/*
 * ONTAP REST API
 *
 * ONTAP 9.6 adds support for an expansive RESTful API. The documentation below provides information about the types of API calls available to you, as well as details about using each API endpoint. You can learn more about the ONTAP REST API and ONTAP in the ONTAP 9 Documentation Center:  http://docs.netapp.com/ontap-9/topic/com.netapp.doc.dot-rest-api/home.html. NetApp welcomes your comments and suggestions about the ONTAP REST API and the documentation for its use.</br> **Using the ONTAP 9.6 REST API online documentation** Each API method includes usage examples, as well as a model that displays all the required and optional properties supported by the method. Click the _Model_ link, available with each API method, to see all the required and optional properties supported by each method.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ontap

// Account struct for Account
type Account struct {
	Links        InlineResponse201Links `json:"_links,omitempty"`
	Applications []AccountApplication   `json:"applications,omitempty"`
	// Optional comment for the user account.
	Comment string `json:"comment,omitempty"`
	// Locked status of the account.
	Locked bool `json:"locked,omitempty"`
	// User or group account name
	Name  string       `json:"name,omitempty"`
	Owner AccountOwner `json:"owner,omitempty"`
	// Password for the account. The password can contain a mix of lower and upper case alphabetic characters, digits, and special characters.
	Password string        `json:"password,omitempty"`
	Role     RoleReference `json:"role,omitempty"`
	// Scope of the entity. set to \"cluster\" for cluster owned objects and to \"svm\" for SVM owned objects.
	Scope string `json:"scope,omitempty"`
}
