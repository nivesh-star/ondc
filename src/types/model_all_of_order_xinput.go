/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// Additional input required from the customer to confirm this order
type AllOfOrderXinput struct {
	Form *Form `json:"form,omitempty"`
	// Indicates whether the form data is mandatorily required by the BPP to confirm the order.
	Required bool `json:"required,omitempty"`
}
