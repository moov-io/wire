// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &ErrorWire{}

// ErrorWire is a wire error with the fedwire message
type ErrorWire struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
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
func NewErrorWire(isVariable bool) *ErrorWire {
	ew := &ErrorWire{
		tag:              TagErrorWire,
		isVariableLength: isVariable,
	}
	return ew
}

// Parse takes the input string and parses the ErrorWire values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ew *ErrorWire) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagWrongLengthErr(9, len(record)), 0
	}

	ew.tag = record[:6]
	ew.ErrorCategory = ew.parseStringField(record[6:7])

	length := 7
	read := 0

	ew.ErrorCode, read = ew.parseVariableStringField(record[length:], 3)
	length += read

	ew.ErrorDescription, read = ew.parseVariableStringField(record[length:], 35)
	length += read

	return nil, length
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
func (ew *ErrorWire) String() string {
	var buf strings.Builder
	buf.Grow(45)
	buf.WriteString(ew.tag)
	buf.WriteString(ew.ErrorCategoryField())
	buf.WriteString(ew.ErrorCodeField())
	buf.WriteString(ew.ErrorDescriptionField())
	return buf.String()
}

// Validate performs WIRE format rule checks on ErrorWire and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ew *ErrorWire) Validate() error {
	// Currently no validation as the FED is responsible for the values
	return nil
}

// ErrorCategoryField gets a string of the ErrorCategory field
func (ew *ErrorWire) ErrorCategoryField() string {
	return ew.alphaField(ew.ErrorCategory, 1)
}

// ErrorCodeField gets a string of the ErrorCode field
func (ew *ErrorWire) ErrorCodeField() string {
	return ew.alphaVariableField(ew.ErrorCode, 3, ew.isVariableLength)
}

// ErrorDescriptionField gets a string of the ErrorDescription field
func (ew *ErrorWire) ErrorDescriptionField() string {
	return ew.alphaVariableField(ew.ErrorDescription, 35, ew.isVariableLength)
}
