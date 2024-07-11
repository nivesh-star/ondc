/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// The replacement policy of an item or an order
type ReplacementTerm struct {
	// The state of fulfillment during which this term is applicable.
	FulfillmentState AllOfReplacementTermFulfillmentState `json:"fulfillment_state,omitempty"`
	// Applicable only for buyer managed returns where the buyer has to replace the item before a certain date-time, failing which they will not be eligible for replacement
	ReplaceWithin AllOfReplacementTermReplaceWithin `json:"replace_within,omitempty"`

	ExternalRef MediaFile `json:"external_ref,omitempty"`
}
