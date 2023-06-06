// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func NewPreviousMessageIdentifier() *PreviousMessageIdentifier {
	pmi := &PreviousMessageIdentifier{
		tag: TagPreviousMessageIdentifier,
	}
	return pmi
}

// Parse takes the input string and parses the PreviousMessageIdentifier values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pmi *PreviousMessageIdentifier) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	pmi.tag = record[:6]
	length := 6

	value, read, err := pmi.parseVariableStringField(record[length:], 22)
	if err != nil {
		return fieldError("PreviousMessageIdentifier", err)
	}
	pmi.PreviousMessageIdentifier = value
	length += read

	if err := pmi.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width PreviousMessageIdentifier record
func (pmi *PreviousMessageIdentifier) String() string {
	return pmi.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a PreviousMessageIdentifier record formatted according to the FormatOptions
func (pmi *PreviousMessageIdentifier) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(pmi.tag)
	buf.WriteString(pmi.FormatPreviousMessageIdentifier(options))

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
	return pmi.alphaField(pmi.PreviousMessageIdentifier, 22)
}

// FormatPreviousMessageIdentifier returns PreviousMessageIdentifier formatted according to the FormatOptions
func (pmi *PreviousMessageIdentifier) FormatPreviousMessageIdentifier(options FormatOptions) string {
	return pmi.formatAlphaField(pmi.PreviousMessageIdentifier, 22, options)
}
