/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// Describes the price of a product or service
type Price struct {
	Currency string `json:"currency,omitempty"`

	Value string `json:"value,omitempty"`

	EstimatedValue string `json:"estimated_value,omitempty"`

	ComputedValue string `json:"computed_value,omitempty"`

	ListedValue string `json:"listed_value,omitempty"`

	OfferedValue string `json:"offered_value,omitempty"`

	MinimumValue string `json:"minimum_value,omitempty"`

	MaximumValue string `json:"maximum_value,omitempty"`
}
