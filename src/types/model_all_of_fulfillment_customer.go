/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// The person that will ultimately receive the order
type AllOfFulfillmentCustomer struct {
	Person Person `json:"person,omitempty"`

	Contact Contact `json:"contact,omitempty"`
}
