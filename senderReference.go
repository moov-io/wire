// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &SenderReference{}

// SenderReference is the SenderReference of the wire
type SenderReference struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// SenderReference
	SenderReference string `json:"senderReference,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSenderReference returns a new SenderReference
func NewSenderReference(isVariable bool) *SenderReference {
	sr := &SenderReference{
		tag:              TagSenderReference,
		isVariableLength: isVariable,
	}
	return sr
}

// Parse takes the input string and parses the SenderReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sr *SenderReference) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 7 {
		return 0, NewTagWrongLengthErr(7, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if sr.tag, read, err = sr.parseTag(record); err != nil {
		return 0, fieldError("SenderReference.Tag", err)
	}
	length += read

	if sr.SenderReference, read, err = sr.parseVariableStringField(record[length:], 16); err != nil {
		return 0, fieldError("SenderReference", err)
	}
	length += read

	return length, nil
}

func (sr *SenderReference) UnmarshalJSON(data []byte) error {
	type Alias SenderReference
	aux := struct {
		*Alias
	}{
		(*Alias)(sr),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	sr.tag = TagSenderReference
	return nil
}

// String writes SenderReference
func (sr *SenderReference) String() string {
	var buf strings.Builder
	buf.Grow(22)
	buf.WriteString(sr.tag)
	buf.WriteString(sr.SenderReferenceField())
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderReference and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sr *SenderReference) Validate() error {
	if sr.tag != TagSenderReference {
		return fieldError("tag", ErrValidTagForType, sr.tag)
	}
	if err := sr.isAlphanumeric(sr.SenderReference); err != nil {
		return fieldError("SenderReference", err, sr.SenderReference)
	}
	return nil
}

// SenderReferenceField gets a string of SenderReference field
func (sr *SenderReference) SenderReferenceField() string {
	return sr.alphaVariableField(sr.SenderReference, 16, sr.isVariableLength)
}
