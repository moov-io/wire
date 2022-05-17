// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &Originator{}

// Originator is the originator of the wire
type Originator struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Personal
	Personal Personal `json:"personal,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginator returns a new Originator
func NewOriginator(isVariable bool) *Originator {
	o := &Originator{
		tag:              TagOriginator,
		isVariableLength: isVariable,
	}
	return o
}

// Parse takes the input string and parses the Originator values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (o *Originator) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}
	o.tag = record[:6]

	return 6 + o.Personal.Parse(record[6:]), nil
}

func (o *Originator) UnmarshalJSON(data []byte) error {
	type Alias Originator
	aux := struct {
		*Alias
	}{
		(*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.tag = TagOriginator
	return nil
}

// String writes Originator
func (o *Originator) String() string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(o.tag)
	buf.WriteString(o.Personal.String(o.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on Originator and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (o *Originator) Validate() error {
	if err := o.fieldInclusion(); err != nil {
		return err
	}
	if o.tag != TagOriginator {
		return fieldError("tag", ErrValidTagForType, o.tag)
	}
	// Can be any Identification Code
	if err := o.isIdentificationCode(o.Personal.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, o.Personal.IdentificationCode)
	}
	if err := o.isAlphanumeric(o.Personal.Identifier); err != nil {
		return fieldError("Identifier", err, o.Personal.Identifier)
	}
	if err := o.isAlphanumeric(o.Personal.Name); err != nil {
		return fieldError("Name", err, o.Personal.Name)
	}
	if err := o.isAlphanumeric(o.Personal.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, o.Personal.Address.AddressLineOne)
	}
	if err := o.isAlphanumeric(o.Personal.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, o.Personal.Address.AddressLineTwo)
	}
	if err := o.isAlphanumeric(o.Personal.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, o.Personal.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (o *Originator) fieldInclusion() error {
	if o.Personal.IdentificationCode != "" && o.Personal.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if o.Personal.IdentificationCode == "" && o.Personal.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}
