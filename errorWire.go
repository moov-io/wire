// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// ErrorWire is a wire error with the fedwire message
type ErrorWire struct {
	// tag
	tag string
	//  * `E` - Data Error * `F` - Insufficient Balance * `H` - Accountability Error * `I` - In Process or Intercepted * `W` - Cutoff Hour Error * `X` - Duplicate IMAD
	ErrorCategory string `json:"errorCategory,omitempty"`
	// ErrorCode
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorDescription
	ErrorDescription string `json:"errorDescription,omitempty"`

	// validator is composed for data validation
	// validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewErrorWire returns a new ErrorWire
func NewErrorWire() *ErrorWire {
	ew := &ErrorWire{
		tag: TagErrorWire,
	}
	return ew
}

// Parse takes the input string and parses the ErrorWire values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ew *ErrorWire) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	ew.tag = record[:6]

	var err error
	length := 6
	read := 0

	if ew.ErrorCategory, read, err = ew.parseVariableStringField(record[length:], 1); err != nil {
		return fieldError("ErrorCategory", err)
	}
	length += read

	if ew.ErrorCode, read, err = ew.parseVariableStringField(record[length:], 3); err != nil {
		return fieldError("ErrorCode", err)
	}
	length += read

	if ew.ErrorDescription, read, err = ew.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("ErrorDescription", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (ew *ErrorWire) UnmarshalJSON(data []byte) error {
	type Alias ErrorWire
	aux := struct {
		*Alias
	}{
		(*Alias)(ew),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ew.tag = TagErrorWire
	return nil
}

// String writes ErrorWire
func (ew *ErrorWire) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(45)
	buf.WriteString(ew.tag)

	buf.WriteString(ew.ErrorCategoryField(options...))
	buf.WriteString(ew.ErrorCodeField(options...))
	buf.WriteString(ew.ErrorDescriptionField(options...))

	if ew.parseFirstOption(options) {
		return ew.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on ErrorWire and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ew *ErrorWire) Validate() error {
	// Currently no validation as the FED is responsible for the values
	return nil
}

// ErrorCategoryField gets a string of the ErrorCategory field
func (ew *ErrorWire) ErrorCategoryField(options ...bool) string {
	return ew.alphaVariableField(ew.ErrorCategory, 1, ew.parseFirstOption(options))
}

// ErrorCodeField gets a string of the ErrorCode field
func (ew *ErrorWire) ErrorCodeField(options ...bool) string {
	return ew.alphaVariableField(ew.ErrorCode, 3, ew.parseFirstOption(options))
}

// ErrorDescriptionField gets a string of the ErrorDescription field
func (ew *ErrorWire) ErrorDescriptionField(options ...bool) string {
	return ew.alphaVariableField(ew.ErrorDescription, 35, ew.parseFirstOption(options))
}
