/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// The state where the organization's address is registered
type AllOfOrganizationState struct {
	// Name of the state
	Name string `json:"name,omitempty"`
	// State code as per country or international standards
	Code string `json:"code,omitempty"`
}
