// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// OriginatorFI is the originator Financial Institution
type OriginatorFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginatorFI returns a new OriginatorFI
func NewOriginatorFI() *OriginatorFI {
	ofi := &OriginatorFI{
		tag: TagOriginatorFI,
	}
	return ofi
}

// Parse takes the input string and parses the OriginatorFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ofi *OriginatorFI) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	ofi.tag = record[:6]
	length := 6

	value, read, err := ofi.parseVariableStringField(record[length:], 1)
	if err != nil {
		return fieldError("IdentificationCode", err)
	}
	ofi.FinancialInstitution.IdentificationCode = value
	length += read

	value, read, err = ofi.parseVariableStringField(record[length:], 34)
	if err != nil {
		return fieldError("Identifier", err)
	}
	ofi.FinancialInstitution.Identifier = value
	length += read

	value, read, err = ofi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Name", err)
	}
	ofi.FinancialInstitution.Name = value
	length += read

	value, read, err = ofi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	ofi.FinancialInstitution.Address.AddressLineOne = value
	length += read

	value, read, err = ofi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	ofi.FinancialInstitution.Address.AddressLineTwo = value
	length += read

	value, read, err = ofi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	ofi.FinancialInstitution.Address.AddressLineThree = value
	length += read

	if !ofi.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (ofi *OriginatorFI) UnmarshalJSON(data []byte) error {
	type Alias OriginatorFI
	aux := struct {
		*Alias
	}{
		(*Alias)(ofi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ofi.tag = TagOriginatorFI
	return nil
}

// String writes OriginatorFI
func (ofi *OriginatorFI) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(ofi.tag)

	buf.WriteString(ofi.IdentificationCodeField(options...))
	buf.WriteString(ofi.IdentifierField(options...))
	buf.WriteString(ofi.NameField(options...))
	buf.WriteString(ofi.AddressLineOneField(options...))
	buf.WriteString(ofi.AddressLineTwoField(options...))
	buf.WriteString(ofi.AddressLineThreeField(options...))

	if ofi.parseFirstOption(options) {
		return ofi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on OriginatorFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (ofi *OriginatorFI) Validate() error {
	if err := ofi.fieldInclusion(); err != nil {
		return err
	}
	if ofi.tag != TagOriginatorFI {
		return fieldError("tag", ErrValidTagForType, ofi.tag)
	}
	if err := ofi.isIdentificationCode(ofi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, ofi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch ofi.FinancialInstitution.IdentificationCode {
	case
		"B", "C", "D", "F", "U":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, ofi.FinancialInstitution.IdentificationCode)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, ofi.FinancialInstitution.Identifier)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, ofi.FinancialInstitution.Name)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ofi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ofi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ofi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ofi *OriginatorFI) fieldInclusion() error {
	if ofi.FinancialInstitution.IdentificationCode != "" && ofi.FinancialInstitution.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if ofi.FinancialInstitution.IdentificationCode == "" && ofi.FinancialInstitution.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ofi *OriginatorFI) IdentificationCodeField(options ...bool) string {
	return ofi.alphaVariableField(ofi.FinancialInstitution.IdentificationCode, 1, ofi.parseFirstOption(options))
}

// IdentifierField gets a string of the Identifier field
func (ofi *OriginatorFI) IdentifierField(options ...bool) string {
	return ofi.alphaVariableField(ofi.FinancialInstitution.Identifier, 34, ofi.parseFirstOption(options))
}

// NameField gets a string of the Name field
func (ofi *OriginatorFI) NameField(options ...bool) string {
	return ofi.alphaVariableField(ofi.FinancialInstitution.Name, 35, ofi.parseFirstOption(options))
}

// AddressLineOneField gets a string of AddressLineOne field
func (ofi *OriginatorFI) AddressLineOneField(options ...bool) string {
	return ofi.alphaVariableField(ofi.FinancialInstitution.Address.AddressLineOne, 35, ofi.parseFirstOption(options))
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (ofi *OriginatorFI) AddressLineTwoField(options ...bool) string {
	return ofi.alphaVariableField(ofi.FinancialInstitution.Address.AddressLineTwo, 35, ofi.parseFirstOption(options))
}

// AddressLineThreeField gets a string of AddressLineThree field
func (ofi *OriginatorFI) AddressLineThreeField(options ...bool) string {
	return ofi.alphaVariableField(ofi.FinancialInstitution.Address.AddressLineThree, 35, ofi.parseFirstOption(options))
}
