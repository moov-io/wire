// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &PreviousMessageIdentifier{}

// PreviousMessageIdentifier is the PreviousMessageIdentifier of the wire
type PreviousMessageIdentifier struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// PreviousMessageIdentifier
	PreviousMessageIdentifier string `json:"PreviousMessageIdentifier,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewPreviousMessageIdentifier returns a new PreviousMessageIdentifier
func NewPreviousMessageIdentifier(isVariable bool) *PreviousMessageIdentifier {
	pmi := &PreviousMessageIdentifier{
		tag:              TagPreviousMessageIdentifier,
		isVariableLength: isVariable,
	}
	return pmi
}

// Parse takes the input string and parses the PreviousMessageIdentifier values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pmi *PreviousMessageIdentifier) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 7 {
		return 0, NewTagWrongLengthErr(7, len(record))
	}

	pmi.tag = record[:6]

	length := 6
	read := 0

	pmi.PreviousMessageIdentifier, read = pmi.parseVariableStringField(record[length:], 22)
	length += read

	return length, nil
}

func (pmi *PreviousMessageIdentifier) UnmarshalJSON(data []byte) error {
	type Alias PreviousMessageIdentifier
	aux := struct {
		*Alias
	}{
		(*Alias)(pmi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	pmi.tag = TagPreviousMessageIdentifier
	return nil
}

// String writes PreviousMessageIdentifier
func (pmi *PreviousMessageIdentifier) String() string {
	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(pmi.tag)
	buf.WriteString(pmi.PreviousMessageIdentifierField())

	return buf.String()
}

// Validate performs WIRE format rule checks on PreviousMessageIdentifier and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (pmi *PreviousMessageIdentifier) Validate() error {
	if pmi.tag != TagPreviousMessageIdentifier {
		return fieldError("tag", ErrValidTagForType, pmi.tag)
	}
	if err := pmi.isAlphanumeric(pmi.PreviousMessageIdentifier); err != nil {
		return fieldError("PreviousMessageIdentifier", err, pmi.PreviousMessageIdentifier)
	}
	return nil
}

// PreviousMessageIdentifierField gets a string of PreviousMessageIdentifier field
func (pmi *PreviousMessageIdentifier) PreviousMessageIdentifierField() string {
	return pmi.alphaVariableField(pmi.PreviousMessageIdentifier, 22, pmi.isVariableLength)
}
