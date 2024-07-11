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

// Every API call in beckn protocol has a context. It provides a high-level overview to the receiver about the nature of the intended transaction. Typically, it is the BAP that sets the transaction context based on the consumer's location and action on their UI. But sometimes, during unsolicited callbacks, the BPP also sets the transaction context but it is usually the same as the context of a previous full-cycle, request-callback interaction between the BAP and the BPP. The context object contains four  package types of fields. <ol><li>Demographic information about the transaction using fields like `domain`, `country`, and `region`.</li><li>Addressing details like the sending and receiving platform's ID and API URL.</li><li>Interoperability information like the protocol version that implemented by the sender and,</li><li>Transaction details like the method being called at the receiver's endpoint, the transaction_id that represents an end-to-end user session at the BAP, a message ID to pair requests with callbacks, a timestamp to capture sending times, a ttl to specifiy the validity of the request, and a key to encrypt information if necessary.</li></ol> This object must be passed in every interaction between a BAP and a BPP. In HTTP/S implementations, it is not necessary to send the context during the synchronous response. However, in asynchronous protocols, the context must be sent during all interactions,
type Context struct {
	// Domain code that is relevant to this transaction context
	Domain *AllOfContextDomain `json:"domain,omitempty"`
	// The location where the transaction is intended to be fulfilled.
	Location *AllOfContextLocation `json:"location,omitempty"`
	// The Beckn protocol method being called by the sender and executed at the receiver.
	Action string `json:"action,omitempty"`
	// Version of transaction protocol being used by the sender.
	Version string `json:"version,omitempty"`
	// Subscriber ID of the BAP
	BapId *AllOfContextBapId `json:"bap_id,omitempty"`
	// Subscriber URL of the BAP for accepting callbacks from BPPs.
	BapUri *AllOfContextBapUri `json:"bap_uri,omitempty"`
	// Subscriber ID of the BPP
	BppId *AllOfContextBppId `json:"bpp_id,omitempty"`
	// Subscriber URL of the BPP for accepting calls from BAPs.
	BppUri *AllOfContextBppUri `json:"bpp_uri,omitempty"`
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
