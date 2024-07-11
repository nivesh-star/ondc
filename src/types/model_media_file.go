/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package types

// This object contains a url to a media file.
type MediaFile struct {
	// indicates the nature and format of the document, file, or assortment of bytes. MIME  package types are defined and standardized in IETF's RFC 6838
	Mimetype string `json:"mimetype,omitempty"`
	// The URL of the file
	Url string `json:"url,omitempty"`
	// The digital signature of the file signed by the sender
	Signature string `json:"signature,omitempty"`
	// The signing algorithm used by the sender
	Dsa string `json:"dsa,omitempty"`
}
