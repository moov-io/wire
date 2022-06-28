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
	length := 6

	value, read, err := ew.parseVariableStringField(record[length:], 1)
	if err != nil {
		return fieldError("ErrorCategory", err)
	}
	ew.ErrorCategory = value
	length += read

	value, read, err = ew.parseVariableStringField(record[length:], 3)
	if err != nil {
		return fieldError("ErrorCode", err)
	}
	ew.ErrorCode = value
	length += read

	value, read, err = ew.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ErrorDescription", err)
	}
	ew.ErrorDescription = value
	length += read

	if !ew.verifyDataWithReadLength(record, length) {
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

// String returns a fixed-width ErrorWire record
func (ew *ErrorWire) String() string {
	return ew.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a ErrorWire record formatted according to the FormatOptions
func (ew *ErrorWire) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(45)
	buf.WriteString(ew.tag)

	buf.WriteString(ew.FormatErrorCategory(options))
	buf.WriteString(ew.FormatErrorCode(options))
	buf.WriteString(ew.FormatErrorDescription(options))

	if options.VariableLengthFields {
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
func (ew *ErrorWire) ErrorCategoryField() string {
	return ew.alphaField(ew.ErrorCategory, 1)
}

// ErrorCodeField gets a string of the ErrorCode field
func (ew *ErrorWire) ErrorCodeField() string {
	return ew.alphaField(ew.ErrorCode, 3)
}

// ErrorDescriptionField gets a string of the ErrorDescription field
func (ew *ErrorWire) ErrorDescriptionField() string {
	return ew.alphaField(ew.ErrorDescription, 35)
}

// FormatErrorCategory returns ErrorCategory formatted according to the FormatOptions
func (ew *ErrorWire) FormatErrorCategory(options FormatOptions) string {
	return ew.formatAlphaField(ew.ErrorCategory, 1, options)
}

// FormatErrorCode returns ErrorCode formatted according to the FormatOptions
func (ew *ErrorWire) FormatErrorCode(options FormatOptions) string {
	return ew.formatAlphaField(ew.ErrorCode, 3, options)
}

// FormatErrorDescription returns ErrorDescription formatted according to the FormatOptions
func (ew *ErrorWire) FormatErrorDescription(options FormatOptions) string {
	return ew.formatAlphaField(ew.ErrorDescription, 35, options)
}
