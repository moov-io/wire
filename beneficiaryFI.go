// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &BeneficiaryFI{}

// BeneficiaryFI is the financial institution of the beneficiary
type BeneficiaryFI struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryFI returns a new BeneficiaryFI
func NewBeneficiaryFI(isVariable bool) *BeneficiaryFI {
	bfi := &BeneficiaryFI{
		tag:              TagBeneficiaryFI,
		isVariableLength: isVariable,
	}
	return bfi
}

// Parse takes the input string and parses the BeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bfi *BeneficiaryFI) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 12 {
		return NewTagWrongLengthErr(12, len(record)), 0
	}
	bfi.tag = record[:6]

	return nil, 6 + bfi.FinancialInstitution.Parse(record[6:])
}

func (bfi *BeneficiaryFI) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(bfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bfi.tag = TagBeneficiaryFI
	return nil
}

// String writes BeneficiaryFI
func (bfi *BeneficiaryFI) String() string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(bfi.tag)
	buf.WriteString(bfi.FinancialInstitution.String(bfi.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bfi *BeneficiaryFI) Validate() error {
	if err := bfi.fieldInclusion(); err != nil {
		return err
	}
	if bfi.tag != TagBeneficiaryFI {
		return fieldError("tag", ErrValidTagForType, bfi.tag)
	}
	if err := bfi.isIdentificationCode(bfi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, bfi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch bfi.FinancialInstitution.IdentificationCode {
	case
		SWIFTBankIdentifierCode,
		CHIPSParticipant,
		DemandDepositAccountNumber,
		FEDRoutingNumber,
		CHIPSIdentifier:
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, bfi.FinancialInstitution.IdentificationCode)
	}
	if err := bfi.isAlphanumeric(bfi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, bfi.FinancialInstitution.Identifier)
	}
	if err := bfi.isAlphanumeric(bfi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, bfi.FinancialInstitution.Name)
	}
	if err := bfi.isAlphanumeric(bfi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, bfi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := bfi.isAlphanumeric(bfi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, bfi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := bfi.isAlphanumeric(bfi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, bfi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bfi *BeneficiaryFI) fieldInclusion() error {
	if bfi.FinancialInstitution.IdentificationCode != "" && bfi.FinancialInstitution.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if bfi.FinancialInstitution.IdentificationCode == "" && bfi.FinancialInstitution.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}
