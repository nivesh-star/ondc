/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// Describes the acknowledgement sent in response to an API call. If the implementation uses HTTP/S, then Ack must be returned in the same session. Every API call to a BPP must be responded to with an Ack whether the BPP intends to respond with a callback or not. This has one property called `status` that indicates the status of the Acknowledgement.
type Ack struct {
	// The status of the acknowledgement. If the request passes the validation criteria of the BPP, then this is set to ACK. If a BPP responds with status = `ACK` to a request, it is required to respond with a callback. If the request fails the validation criteria, then this is set to NACK. Additionally, if a BPP does not intend to respond with a callback even after the request meets the validation criteria, it should set this value to `NACK`.
	Status string `json:"status,omitempty"`
	// A list of tags containing any additional information sent along with the Acknowledgement.
	Tags []TagGroup `json:"tags,omitempty"`
}
