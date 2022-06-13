// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// Beneficiary is the beneficiary of the wire
type Beneficiary struct {
	// tag
	tag string
	// Personal
	Personal Personal `json:"personal,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiary returns a new Beneficiary
func NewBeneficiary() *Beneficiary {
	ben := &Beneficiary{
		tag: TagBeneficiary,
	}
	return ben
}

// Parse takes the input string and parses the Beneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ben *Beneficiary) Parse(record string) error {
	if utf8.RuneCountInString(record) < 7 {
		return NewTagMinLengthErr(7, len(record))
	}

	ben.tag = record[:6]
	ben.Personal.IdentificationCode = ben.parseStringField(record[6:7])
	length := 7

	value, read, err := ben.parseVariableStringField(record[length:], 34)
	if err != nil {
		return fieldError("Identifier", err)
	}
	ben.Personal.Identifier = value
	length += read

	value, read, err = ben.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Name", err)
	}
	ben.Personal.Name = value
	length += read

	value, read, err = ben.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	ben.Personal.Address.AddressLineOne = value
	length += read

	value, read, err = ben.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	ben.Personal.Address.AddressLineTwo = value
	length += read

	value, read, err = ben.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	ben.Personal.Address.AddressLineThree = value
	length += read

	if !ben.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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
func (ben *Beneficiary) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(ben.tag)
	buf.WriteString(ben.IdentificationCodeField())
	buf.WriteString(ben.IdentifierField(options...))
	buf.WriteString(ben.NameField(options...))
	buf.WriteString(ben.AddressLineOneField(options...))
	buf.WriteString(ben.AddressLineTwoField(options...))
	buf.WriteString(ben.AddressLineThreeField(options...))

	if ben.parseFirstOption(options) {
		return ben.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// IdentificationCodeField gets a string of the IdentificationCode field
func (ben *Beneficiary) IdentificationCodeField() string {
	return ben.alphaField(ben.Personal.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (ben *Beneficiary) IdentifierField(options ...bool) string {
	return ben.alphaVariableField(ben.Personal.Identifier, 34, ben.parseFirstOption(options))
}

// NameField gets a string of the Name field
func (ben *Beneficiary) NameField(options ...bool) string {
	return ben.alphaVariableField(ben.Personal.Name, 35, ben.parseFirstOption(options))
}

// AddressLineOneField gets a string of AddressLineOne field
func (ben *Beneficiary) AddressLineOneField(options ...bool) string {
	return ben.alphaVariableField(ben.Personal.Address.AddressLineOne, 35, ben.parseFirstOption(options))
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (ben *Beneficiary) AddressLineTwoField(options ...bool) string {
	return ben.alphaVariableField(ben.Personal.Address.AddressLineTwo, 35, ben.parseFirstOption(options))
}

// AddressLineThreeField gets a string of AddressLineThree field
func (ben *Beneficiary) AddressLineThreeField(options ...bool) string {
	return ben.alphaVariableField(ben.Personal.Address.AddressLineThree, 35, ben.parseFirstOption(options))
}
