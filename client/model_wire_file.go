/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WireFile struct for WireFile
type WireFile struct {
	// File ID
	ID             string         `json:"ID,omitempty"`
	FedWireMessage FedWireMessage `json:"fedWireMessage"`
}
