/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// details on the offer the customer wants to avail
type AllOfIntentOffer struct {
	Id string `json:"id,omitempty"`

	Descriptor Descriptor `json:"descriptor,omitempty"`

	LocationIds []Locationpropertiesid `json:"location_ids,omitempty"`

	CategoryIds []Categorypropertiesid `json:"category_ids,omitempty"`

	ItemIds []Itempropertiesid `json:"item_ids,omitempty"`

	Time Time `json:"time,omitempty"`

	Tags []TagGroup `json:"tags,omitempty"`
}
