/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

type TrackBody struct {
	Context *AllOftrackBodyContext `json:"context"`

	Message *TrackMessage `json:"message"`
}
