/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// Contact details of the stop
type AllOfStopContact struct {
	Phone string `json:"phone,omitempty"`

	Email string `json:"email,omitempty"`
	// A Jcard object as per draft-ietf-jcardcal-jcard-03 specification
	Jcard *interface{} `json:"jcard,omitempty"`
}
