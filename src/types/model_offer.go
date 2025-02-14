/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

type Locationpropertiesid string
type Itempropertiesid string

// An offer associated with a catalog. This is typically used to promote a particular product and enable more purchases.
type Offer struct {
	Id string `json:"id,omitempty"`

	Descriptor Descriptor `json:"descriptor,omitempty"`

	LocationIds []Locationpropertiesid `json:"location_ids,omitempty"`

	CategoryIds []Categorypropertiesid `json:"category_ids,omitempty"`

	ItemIds []Itempropertiesid `json:"item_ids,omitempty"`

	Time Time `json:"time,omitempty"`

	Tags []TagGroup `json:"tags,omitempty"`
}
