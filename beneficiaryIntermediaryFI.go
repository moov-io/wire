// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// BeneficiaryIntermediaryFI {4000}
type BeneficiaryIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryIntermediaryFI returns a new BeneficiaryIntermediaryFI
func NewBeneficiaryIntermediaryFI() BeneficiaryIntermediaryFI {
	bifi := BeneficiaryIntermediaryFI{
		tag: TagBeneficiaryIntermediaryFI,
	}
	return bifi
}

// Parse takes the input string and parses the ReceiverDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bifi *BeneficiaryIntermediaryFI) Parse(record string) {
	bifi.tag = record[:6]
	bifi.FinancialInstitution.IdentificationCode = bifi.parseStringField(record[6:7])
	bifi.FinancialInstitution.Identifier = bifi.parseStringField(record[7:41])
	bifi.FinancialInstitution.Name = bifi.parseStringField(record[41:76])
	bifi.FinancialInstitution.Address.AddressLineOne = bifi.parseStringField(record[76:111])
	bifi.FinancialInstitution.Address.AddressLineTwo = bifi.parseStringField(record[111:146])
	bifi.FinancialInstitution.Address.AddressLineThree = bifi.parseStringField(record[146:181])
}

// String writes BeneficiaryIntermediaryFI
func (bifi *BeneficiaryIntermediaryFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(181)
	buf.WriteString(bifi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bifi *BeneficiaryIntermediaryFI) Validate() error {
	if err := bifi.fieldInclusion(); err != nil {
		return err
	}
	if err := bifi.isIdentificationCode(bifi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, bifi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch bifi.FinancialInstitution.IdentificationCode {
	case
		"B", "C", "D", "F", "U":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, bifi.FinancialInstitution.Identifier)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, bifi.FinancialInstitution.Name)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, bifi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, bifi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, bifi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bifi *BeneficiaryIntermediaryFI) fieldInclusion() error {
	return nil
}
