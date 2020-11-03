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
	if utf8.RuneCountInString(record) != 181 {
		return NewTagWrongLengthErr(181, len(record))
	}
	ofi.tag = record[:6]
	ofi.FinancialInstitution.IdentificationCode = ofi.parseStringField(record[6:7])
	ofi.FinancialInstitution.Identifier = ofi.parseStringField(record[7:41])
	ofi.FinancialInstitution.Name = ofi.parseStringField(record[41:76])
	ofi.FinancialInstitution.Address.AddressLineOne = ofi.parseStringField(record[76:111])
	ofi.FinancialInstitution.Address.AddressLineTwo = ofi.parseStringField(record[111:146])
	ofi.FinancialInstitution.Address.AddressLineThree = ofi.parseStringField(record[146:181])
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
func (ofi *OriginatorFI) String() string {
	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(ofi.tag)
	buf.WriteString(ofi.IdentificationCodeField())
	buf.WriteString(ofi.IdentifierField())
	buf.WriteString(ofi.NameField())
	buf.WriteString(ofi.AddressLineOneField())
	buf.WriteString(ofi.AddressLineTwoField())
	buf.WriteString(ofi.AddressLineThreeField())
	return buf.String()
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
func (ofi *OriginatorFI) IdentificationCodeField() string {
	return ofi.alphaField(ofi.FinancialInstitution.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (ofi *OriginatorFI) IdentifierField() string {
	return ofi.alphaField(ofi.FinancialInstitution.Identifier, 34)
}

// NameField gets a string of the Name field
func (ofi *OriginatorFI) NameField() string {
	return ofi.alphaField(ofi.FinancialInstitution.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (ofi *OriginatorFI) AddressLineOneField() string {
	return ofi.alphaField(ofi.FinancialInstitution.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (ofi *OriginatorFI) AddressLineTwoField() string {
	return ofi.alphaField(ofi.FinancialInstitution.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (ofi *OriginatorFI) AddressLineThreeField() string {
	return ofi.alphaField(ofi.FinancialInstitution.Address.AddressLineThree, 35)
}
