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

type AllOfsupportBodyContext struct {
	Action string `json:"action"`
	// Domain code that is relevant to this transaction context
	Domain string `json:"domain,omitempty"`
	// The location where the transaction is intended to be fulfilled.
	Location string `json:"location,omitempty"`
	// Version of transaction protocol being used by the sender.
	Version string `json:"version,omitempty"`
	// Subscriber ID of the BAP
	BapId string `json:"bap_id,omitempty"`
	// Subscriber URL of the BAP for accepting callbacks from BPPs.
	BapUri string `json:"bap_uri,omitempty"`
	// Subscriber ID of the BPP
	BppId string `json:"bpp_id,omitempty"`
	// Subscriber URL of the BPP for accepting calls from BAPs.
	BppUri string `json:"bpp_uri,omitempty"`
	// This is a unique value which persists across all API calls from `search` through `confirm`. This is done to indicate an active user session across multiple requests. The BPPs can use this value to push personalized recommendations, and dynamic offerings related to an ongoing transaction despite being unaware of the user active on the BAP.
	TransactionId string `json:"transaction_id,omitempty"`
	// This is a unique value which persists during a request / callback cycle. Since beckn protocol APIs are asynchronous, BAPs need a common value to match an incoming callback from a BPP to an earlier call. This value can also be used to ignore duplicate messages coming from the BPP. It is recommended to generate a fresh message_id for every new interaction. When sending unsolicited callbacks, BPPs must generate a new message_id.
	MessageId string `json:"message_id,omitempty"`
	// Time of request generation in RFC3339 format
	Timestamp time.Time `json:"timestamp,omitempty"`
	// The encryption public key of the sender
	Key string `json:"key,omitempty"`
	// The duration in ISO8601 format after timestamp for which this message holds valid
	Ttl string `json:"ttl,omitempty"`
}
