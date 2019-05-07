// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// LocalInstrument is the LocalInstrument of the wire
type LocalInstrument struct {
	// tag
	tag string
	// LocalInstrumentCode is local instrument code
	LocalInstrumentCode string `json:"LocalInstrument,omitempty"`
	// ProprietaryCode is proprietary code
	ProprietaryCode string `json:"proprietaryCode,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewLocalInstrument returns a new LocalInstrument
func NewLocalInstrument() *LocalInstrument {
	li := &LocalInstrument{
		tag: TagLocalInstrument,
	}
	return li
}

// Parse takes the input string and parses the LocalInstrument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (li *LocalInstrument) Parse(record string) error {
	if utf8.RuneCountInString(record) != 45  {
		return NewTagWrongLengthErr(45, len(record))
	}
	li.tag = record[:6]
	li.LocalInstrumentCode = li.parseStringField(record[6:10])
	li.ProprietaryCode = li.parseStringField(record[10:45])
	return nil
}

// String writes LocalInstrument
func (li *LocalInstrument) String() string {
	var buf strings.Builder
	buf.Grow(45)
	buf.WriteString(li.tag)
	buf.WriteString(li.LocalInstrumentCodeField())
	buf.WriteString(li.ProprietaryCodeField())
	return buf.String()
}

// Validate performs WIRE format rule checks on LocalInstrument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (li *LocalInstrument) Validate() error {
	if err := li.fieldInclusion(); err != nil {
		return err
	}
	if err := li.isLocalInstrumentCode(li.LocalInstrumentCode); err != nil {
		return fieldError("LocalInstrumentCode", err, li.LocalInstrumentCode)
	}
	if err := li.isAlphanumeric(li.ProprietaryCode); err != nil {
		return fieldError("ProprietaryCode", err, li.ProprietaryCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (li *LocalInstrument) fieldInclusion() error {
	if li.LocalInstrumentCode != ProprietaryLocalInstrumentCode && li.ProprietaryCode != "" {
		return fieldError("ProprietaryCode", ErrInvalidProperty, li.ProprietaryCode)
	}
	return nil
}

// LocalInstrumentCodeField gets a string of LocalInstrumentCode field
func (li *LocalInstrument) LocalInstrumentCodeField() string {
	return li.alphaField(li.LocalInstrumentCode, 4)
}

// ProprietaryCodeField gets a string of ProprietaryCode field
func (li *LocalInstrument) ProprietaryCodeField() string {
	return li.alphaField(li.ProprietaryCode, 35)
}
