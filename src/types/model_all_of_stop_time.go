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

// Timings applicable at the stop.
type AllOfStopTime struct {
	Label string `json:"label,omitempty"`

	Timestamp time.Time `json:"timestamp,omitempty"`

	Duration string `json:"duration,omitempty"`

	Range_ TimeRange `json:"range,omitempty"`
	// comma separated values representing days of the week
	Days string `json:"days,omitempty"`

	Schedule Schedule `json:"schedule,omitempty"`
}
