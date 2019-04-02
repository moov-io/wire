// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// PreviousMessageIdentifier is the PreviousMessageIdentifier of the wire
type PreviousMessageIdentifier struct {
	// tag
	tag string
	// PreviousMessageIdentifier
	PreviousMessageIdentifier string `json:"PreviousMessageIdentifier,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewPreviousMessageIdentifier returns a new PreviousMessageIdentifier
func NewPreviousMessageIdentifier() PreviousMessageIdentifier {
	pmi := PreviousMessageIdentifier{
		tag: TagPreviousMessageIdentifier,
	}
	return pmi
}

// Parse takes the input string and parses the PreviousMessageIdentifier values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pmi *PreviousMessageIdentifier) Parse(record string) {
	pmi.tag = record[:6]
	pmi.PreviousMessageIdentifier = pmi.parseStringField(record[6:28])
}

// String writes PreviousMessageIdentifier
func (pmi *PreviousMessageIdentifier) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(28)
	buf.WriteString(pmi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on PreviousMessageIdentifier and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (pmi *PreviousMessageIdentifier) Validate() error {
	if err := pmi.fieldInclusion(); err != nil {
		return err
	}
	if err := pmi.isAlphanumeric(pmi.PreviousMessageIdentifier); err != nil {
		return fieldError("PreviousMessageIdentifier", err, pmi.PreviousMessageIdentifier)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (pmi *PreviousMessageIdentifier) fieldInclusion() error {
	return nil
}
