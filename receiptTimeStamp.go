// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
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
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	rts.tag = record[:6]
	length := 6

	value, read, err := rts.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("ReceiptDate", err)
	}
	rts.ReceiptDate = value
	length += read

	value, read, err = rts.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("ReceiptTime", err)
	}
	rts.ReceiptTime = value
	length += read

	value, read, err = rts.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("ReceiptApplicationIdentification", err)
	}
	rts.ReceiptApplicationIdentification = value
	length += read

	if !rts.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (rts *ReceiptTimeStamp) UnmarshalJSON(data []byte) error {
	type Alias ReceiptTimeStamp
	aux := struct {
		*Alias
	}{
		(*Alias)(rts),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rts.tag = TagReceiptTimeStamp
	return nil
}

// String writes ReceiptTimeStamp
func (rts *ReceiptTimeStamp) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(18)

	buf.WriteString(rts.tag)
	buf.WriteString(rts.ReceiptDateField(options...))
	buf.WriteString(rts.ReceiptTimeField(options...))
	buf.WriteString(rts.ReceiptApplicationIdentificationField(options...))

	if rts.parseFirstOption(options) {
		return rts.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
func (rts *ReceiptTimeStamp) ReceiptDateField(options ...bool) string {
	return rts.alphaVariableField(rts.ReceiptDate, 4, rts.parseFirstOption(options))
}

// ReceiptTimeField gets a string of the ReceiptTime field
func (rts *ReceiptTimeStamp) ReceiptTimeField(options ...bool) string {
	return rts.alphaVariableField(rts.ReceiptTime, 4, rts.parseFirstOption(options))
}

// ReceiptApplicationIdentificationField gets a string of the ReceiptApplicationIdentification field
func (rts *ReceiptTimeStamp) ReceiptApplicationIdentificationField(options ...bool) string {
	return rts.alphaVariableField(rts.ReceiptApplicationIdentification, 4, rts.parseFirstOption(options))
}
