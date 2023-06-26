// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func NewSenderReference() *SenderReference {
	sr := &SenderReference{
		tag: TagSenderReference,
	}
	return sr
}

// Parse takes the input string and parses the SenderReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sr *SenderReference) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	sr.tag = record[:6]
	length := 6

	value, read, err := sr.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("SenderReference", err)
	}
	sr.SenderReference = value
	length += read

	if err := sr.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width SenderReference record
func (sr *SenderReference) String() string {
	return sr.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a SenderReference record formatted according to the FormatOptions
func (sr *SenderReference) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(22)

	buf.WriteString(sr.tag)
	buf.WriteString(sr.FormatSenderReference(options) + Delimiter)

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

// FormatSenderReference returns SenderReference formatted according to the FormatOptions
func (sr *SenderReference) FormatSenderReference(options FormatOptions) string {
	return sr.formatAlphaField(sr.SenderReference, 16, options)
}
