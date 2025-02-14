/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

type Categorypropertiesid string

// A label under which a collection of items can be grouped.
type Category struct {
	// ID of the category
	Id string `json:"id,omitempty"`

	ParentCategoryId Categorypropertiesid `json:"parent_category_id,omitempty"`

	Descriptor Descriptor `json:"descriptor,omitempty"`

	Time Time `json:"time,omitempty"`
	// Time to live for an instance of this schema
	Ttl string `json:"ttl,omitempty"`

	Tags []TagGroup `json:"tags,omitempty"`
}
