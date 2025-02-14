/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// Describes the billing details of an entity.<br>This has properties like name,organization,address,email,phone,time,tax_number, created_at,updated_at
type Billing struct {
	// Name of the billable entity
	Name string `json:"name,omitempty"`
	// Details of the organization being billed.
	Organization *AllOfBillingOrganization `json:"organization,omitempty"`
	// The address of the billable entity
	Address *AllOfBillingAddress `json:"address,omitempty"`
	// The state where the billable entity resides. This is important for state-level tax calculation
	State *AllOfBillingState `json:"state,omitempty"`
	// The city where the billable entity resides.
	City *AllOfBillingCity `json:"city,omitempty"`
	// Email address where the bill is sent to
	Email string `json:"email,omitempty"`
	// Phone number of the billable entity
	Phone string `json:"phone,omitempty"`
	// Details regarding the billing period
	Time *AllOfBillingTime `json:"time,omitempty"`
	// ID of the billable entity as recognized by the taxation authority
	TaxId string `json:"tax_id,omitempty"`
}
