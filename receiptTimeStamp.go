// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

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
	// validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewReceiptTimeStamp returns a new ReceiptTimeStamp
func NewReceiptTimeStamp() *ReceiptTimeStamp {
	rts := &ReceiptTimeStamp{
		tag: TagReceiptTimeStamp,
	}
	return rts
}

// Parse takes the input string and parses the ReceiptTimeStamp values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rts *ReceiptTimeStamp) Parse(record string) error {
	if utf8.RuneCountInString(record) != 18 {
		return NewTagWrongLengthErr(18, utf8.RuneCountInString(record))
	}

	rts.tag = record[:6]
	rts.ReceiptDate = rts.parseStringField(record[6:10])
	rts.ReceiptTime = rts.parseStringField(record[10:14])
	rts.ReceiptApplicationIdentification = rts.parseStringField(record[14:18])
	return nil
}

// String writes ReceiptTimeStamp
func (rts *ReceiptTimeStamp) String() string {
	var buf strings.Builder
	buf.Grow(18)
	buf.WriteString(rts.tag)
	buf.WriteString(rts.ReceiptDateField())
	buf.WriteString(rts.ReceiptTimeField())
	buf.WriteString(rts.ReceiptApplicationIdentificationField())
	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiptTimeStamp and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rts *ReceiptTimeStamp) Validate() error {
	// Currently no validation as the FED is responsible for the values
	if rts.tag != TagReceiptTimeStamp {
		return fieldError("tag", ErrValidTagForType, rts.tag)
	}
	return nil
}

// ReceiptDateField gets a string of the ReceiptDate field
func (rts *ReceiptTimeStamp) ReceiptDateField() string {
	return rts.alphaField(rts.ReceiptDate, 4)
}

// ReceiptTimeField gets a string of the ReceiptTime field
func (rts *ReceiptTimeStamp) ReceiptTimeField() string {
	return rts.alphaField(rts.ReceiptTime, 4)
}

// ReceiptApplicationIdentificationField gets a string of the ReceiptApplicationIdentification field
func (rts *ReceiptTimeStamp) ReceiptApplicationIdentificationField() string {
	return rts.alphaField(rts.ReceiptApplicationIdentification, 4)
}
