/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// The city where the billable entity resides.
type AllOfBillingCity struct {
	// Name of the city
	Name string `json:"name,omitempty"`
	// City code
	Code string `json:"code,omitempty"`
}
