// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewErrorWire returns a new ErrorWire
func NewErrorWire() ErrorWire {
	we := ErrorWire{
		tag: TagErrorWire,
	}
	return we
}

// Parse takes the input string and parses the ErrorWire values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (we *ErrorWire) Parse(record string) {
}

// String writes ErrorWire
func (we *ErrorWire) String() string {
	var buf strings.Builder
	buf.Grow(39)
	buf.WriteString(we.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ErrorWire and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (we *ErrorWire) Validate() error {
	if err := we.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (we *ErrorWire) fieldInclusion() error {
	return nil
}
