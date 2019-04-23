// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// ReceiptTimeStamp is the receipt time stamp of the wire
type ReceiptTimeStamp struct {
	// tag
	tag string
	// ReceiptDate is the receipt date
	ReceiptDate string `json:"receiptDate,omitempty"`
	// ReceiptTime is the receipt time
	ReceiptTime string `json:"receiptTime,omitempty"`
	// ApplicationIdentification
	ReceiptApplicationIdentification string `json:"receiptApplicationIdentification,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewReceiptTimeStamp returns a new ReceiptTimeStamp
func NewReceiptTimeStamp() ReceiptTimeStamp {
	rts := ReceiptTimeStamp{
		tag: TagReceiptTimeStamp,
	}
	return rts
}

// Parse takes the input string and parses the ReceiptTimeStamp values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rts *ReceiptTimeStamp) Parse(record string) {
}

// String writes ReceiptTimeStamp
func (rts *ReceiptTimeStamp) String() string {
	var buf strings.Builder
	buf.Grow(12)
	buf.WriteString(rts.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiptTimeStamp and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rts *ReceiptTimeStamp) Validate() error {
	if err := rts.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rts *ReceiptTimeStamp) fieldInclusion() error {
	return nil
}
