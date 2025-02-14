/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

import (
	"time"
)

// The provider from which the customer wants to place to the order from
type AllOfIntentProvider struct {
	// Id of the provider
	Id string `json:"id,omitempty"`

	Descriptor Descriptor `json:"descriptor,omitempty"`
	// Category Id of the provider at the BPP-level catalog
	CategoryId string `json:"category_id,omitempty"`

	Rating string `json:"rating,omitempty"`

	Time Time `json:"time,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	Fulfillments []Fulfillment `json:"fulfillments,omitempty"`

	Payments []Payment `json:"payments,omitempty"`

	Locations []Location `json:"locations,omitempty"`

	Offers []Offer `json:"offers,omitempty"`

	Items []Item `json:"items,omitempty"`
	// Time after which catalog has to be refreshed
	Exp time.Time `json:"exp,omitempty"`
	// Whether this provider can be rated or not
	Rateable bool `json:"rateable,omitempty"`
	// The time-to-live in seconds, for this object. This can be overriden at deeper levels. A value of -1 indicates that this object is not cacheable.
	Ttl int32 `json:"ttl,omitempty"`

	Tags []TagGroup `json:"tags,omitempty"`
}
