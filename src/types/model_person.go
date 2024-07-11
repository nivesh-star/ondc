/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

type AllOfPersonAge string

// Describes a person as any individual
type Person struct {
	// Describes the identity of the person
	Id string `json:"id,omitempty"`
	// Profile url of the person
	Url string `json:"url,omitempty"`
	// the name of the person
	Name string `json:"name,omitempty"`

	Image Image `json:"image,omitempty"`
	// Age of the person
	Age AllOfPersonAge `json:"age,omitempty"`
	// Date of birth of the person
	Dob string `json:"dob,omitempty"`
	// Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender.Allowed values for this field can be published in the network policy
	Gender string `json:"gender,omitempty"`

	Creds []Credential `json:"creds,omitempty"`

	Languages []PersonLanguages `json:"languages,omitempty"`

	Skills []PersonSkills `json:"skills,omitempty"`

	Tags []TagGroup `json:"tags,omitempty"`
}
