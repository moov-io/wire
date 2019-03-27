// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// SenderReference is the SenderReference of the wire
type SenderReference struct {
	// tag
	tag string
	// SenderReference
	SenderReference string `json:"senderReference,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSenderReference returns a new SenderReference
func NewSenderReference() SenderReference {
	sr := SenderReference{
		tag: TagSenderReference,
	}
	return sr
}

// Parse takes the input string and parses the SenderReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sr *SenderReference) Parse(record string) {
}

// String writes SenderReference
func (sr *SenderReference) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(16)
	buf.WriteString(sr.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderReference and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sr *SenderReference) Validate() error {
	if err := sr.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sr *SenderReference) fieldInclusion() error {
	return nil
}
