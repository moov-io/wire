/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ReceiptTimeStamp struct for ReceiptTimeStamp
type ReceiptTimeStamp struct {
	// ReceiptDate is based on the calendar date. (Format MMDD - M=Month, D=Day)
	ReceiptDate string `json:"receiptDate,omitempty"`
	// ReceiptTime is based on a 24-hour clock, Eastern Time. (Format HHmm - H=Hour, m=Minute)
	ReceiptTime string `json:"receiptTime,omitempty"`
	// ApplicationIdentification
	ReceiptApplicationIdentification string `json:"receiptApplicationIdentification,omitempty"`
}
