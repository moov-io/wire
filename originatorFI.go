// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func NewOriginatorFI() OriginatorFI {
	ofi := OriginatorFI{
		tag: TagOriginatorFI,
	}
	return ofi
}

// Parse takes the input string and parses the OriginatorFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ofi *OriginatorFI) Parse(record string) {
	ofi.tag = record[:6]
	ofi.FinancialInstitution.IdentificationCode = ofi.parseStringField(record[6:7])
	ofi.FinancialInstitution.Identifier = ofi.parseStringField(record[7:41])
	ofi.FinancialInstitution.Name = ofi.parseStringField(record[41:76])
	ofi.FinancialInstitution.Address.AddressLineOne = ofi.parseStringField(record[76:111])
	ofi.FinancialInstitution.Address.AddressLineTwo = ofi.parseStringField(record[111:146])
	ofi.FinancialInstitution.Address.AddressLineThree = ofi.parseStringField(record[146:181])
}

// String writes OriginatorFI
func (ofi *OriginatorFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(ofi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OriginatorFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ofi *OriginatorFI) Validate() error {
	if err := ofi.fieldInclusion(); err != nil {
		return err
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
	return nil
}
