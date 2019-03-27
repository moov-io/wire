// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// MessageDisposition is the message disposition of the wire
type MessageDisposition struct {
	// tag
	tag string
	// FormatVersion 30
	FormatVersion string `json:"formatVersion,omitempty"`
	// TestTestProductionCode determines if test or production
	TestProductionCode string `json:"testProductionCode,omitempty"`
	// MessageDuplicationCode  * ` ` - Original Message * `R` - Retrieval of an original message * `P` - Resend
	MessageDuplicationCode string `json:"messageDuplicationCode,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewMessageDisposition returns a new MessageDisposition
func NewMessageDisposition() MessageDisposition {
	md := MessageDisposition{
		tag: TagMessageDisposition,
	}
	return md
}

// Parse takes the input string and parses the MessageDisposition values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (md *MessageDisposition) Parse(record string) {
}

// String writes MessageDisposition
func (md *MessageDisposition) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(5)
	buf.WriteString(md.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on MessageDisposition and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (md *MessageDisposition) Validate() error {
	if err := md.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (md *MessageDisposition) fieldInclusion() error {
	return nil
}
