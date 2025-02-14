/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

type TrackMessage struct {
	OrderId string `json:"order_id"`

	CallbackUrl string `json:"callback_url,omitempty"`
}
