// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &Beneficiary{}

// Beneficiary is the beneficiary of the wire
type Beneficiary struct {
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

// NewBeneficiary returns a new Beneficiary
func NewBeneficiary(isVariable bool) *Beneficiary {
	ben := &Beneficiary{
		tag:              TagBeneficiary,
		isVariableLength: isVariable,
	}
	return ben
}

// Parse takes the input string and parses the Beneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ben *Beneficiary) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if ben.tag, read, err = ben.parseTag(record); err != nil {
		return 0, fieldError("Beneficiary.Tag", err)
	}
	length += read

	if read, err = ben.Personal.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (ben *Beneficiary) UnmarshalJSON(data []byte) error {
	type Alias Beneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(ben),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ben.tag = TagBeneficiary
	return nil
}

// String writes Beneficiary
func (ben *Beneficiary) String() string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(ben.tag)
	buf.WriteString(ben.Personal.String(ben.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on Beneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (ben *Beneficiary) Validate() error {
	if err := ben.fieldInclusion(); err != nil {
		return err
	}
	if ben.tag != TagBeneficiary {
		return fieldError("tag", ErrValidTagForType, ben.tag)
	}
	// Can be any Identification Code
	if err := ben.isIdentificationCode(ben.Personal.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, ben.Personal.IdentificationCode)
	}
	if err := ben.isAlphanumeric(ben.Personal.Identifier); err != nil {
		return fieldError("Identifier", err, ben.Personal.Identifier)
	}
	if err := ben.isAlphanumeric(ben.Personal.Name); err != nil {
		return fieldError("Name", err, ben.Personal.Name)
	}
	if err := ben.isAlphanumeric(ben.Personal.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ben.Personal.Address.AddressLineOne)
	}
	if err := ben.isAlphanumeric(ben.Personal.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ben.Personal.Address.AddressLineTwo)
	}
	if err := ben.isAlphanumeric(ben.Personal.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ben.Personal.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ben *Beneficiary) fieldInclusion() error {
	if ben.Personal.IdentificationCode != "" && ben.Personal.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if ben.Personal.IdentificationCode == "" && ben.Personal.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}
