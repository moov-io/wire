// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// BeneficiaryFI is the financial institution of the beneficiary
type BeneficiaryFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryFI returns a new BeneficiaryFI
func NewBeneficiaryFI() *BeneficiaryFI {
	bfi := &BeneficiaryFI{
		tag: TagBeneficiaryFI,
	}
	return bfi
}

// Parse takes the input string and parses the BeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bfi *BeneficiaryFI) Parse(record string) error {
	if utf8.RuneCountInString(record) < 7 {
		return NewTagMinLengthErr(7, len(record))
	}

	bfi.tag = record[:6]
	bfi.FinancialInstitution.IdentificationCode = bfi.parseStringField(record[6:7])
	length := 7

	value, read, err := bfi.parseVariableStringField(record[length:], 34)
	if err != nil {
		return fieldError("Identifier", err)
	}
	bfi.FinancialInstitution.Identifier = value
	length += read

	value, read, err = bfi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Name", err)
	}
	bfi.FinancialInstitution.Name = value
	length += read

	value, read, err = bfi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	bfi.FinancialInstitution.Address.AddressLineOne = value
	length += read

	value, read, err = bfi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	bfi.FinancialInstitution.Address.AddressLineTwo = value
	length += read

	value, read, err = bfi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	bfi.FinancialInstitution.Address.AddressLineThree = value
	length += read

	if err := bfi.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width BeneficiaryFI record
func (bfi *BeneficiaryFI) String() string {
	return bfi.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a BeneficiaryFI record formatted according to the FormatOptions
func (bfi *BeneficiaryFI) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(bfi.tag)
	buf.WriteString(bfi.IdentificationCodeField())
	buf.WriteString(bfi.FormatIdentifier(options) + Delimiter)
	buf.WriteString(bfi.FormatName(options) + Delimiter)
	buf.WriteString(bfi.FormatAddressLineOne(options) + Delimiter)
	buf.WriteString(bfi.FormatAddressLineTwo(options) + Delimiter)
	buf.WriteString(bfi.FormatAddressLineThree(options) + Delimiter)

	if options.VariableLengthFields {
		return bfi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// IdentificationCodeField gets a string of the IdentificationCode field
func (bfi *BeneficiaryFI) IdentificationCodeField() string {
	return bfi.alphaField(bfi.FinancialInstitution.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (bfi *BeneficiaryFI) IdentifierField() string {
	return bfi.alphaField(bfi.FinancialInstitution.Identifier, 34)
}

// NameField gets a string of the Name field
func (bfi *BeneficiaryFI) NameField() string {
	return bfi.alphaField(bfi.FinancialInstitution.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (bfi *BeneficiaryFI) AddressLineOneField() string {
	return bfi.alphaField(bfi.FinancialInstitution.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (bfi *BeneficiaryFI) AddressLineTwoField() string {
	return bfi.alphaField(bfi.FinancialInstitution.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (bfi *BeneficiaryFI) AddressLineThreeField() string {
	return bfi.alphaField(bfi.FinancialInstitution.Address.AddressLineThree, 35)
}

// FormatIdentifier returns FinancialInstitution.Identifier formatted according to the FormatOptions
func (bfi *BeneficiaryFI) FormatIdentifier(options FormatOptions) string {
	return bfi.formatAlphaField(bfi.FinancialInstitution.Identifier, 34, options)
}

// FormatName returns FinancialInstitution.Name formatted according to the FormatOptions
func (bfi *BeneficiaryFI) FormatName(options FormatOptions) string {
	return bfi.formatAlphaField(bfi.FinancialInstitution.Name, 35, options)
}

// FormatAddressLineOne returns FinancialInstitution.Address.AddressLineOne formatted according to the FormatOptions
func (bfi *BeneficiaryFI) FormatAddressLineOne(options FormatOptions) string {
	return bfi.formatAlphaField(bfi.FinancialInstitution.Address.AddressLineOne, 35, options)
}

// FormatAddressLineTwo returns FinancialInstitution.Address.AddressLineTwo formatted according to the FormatOptions
func (bfi *BeneficiaryFI) FormatAddressLineTwo(options FormatOptions) string {
	return bfi.formatAlphaField(bfi.FinancialInstitution.Address.AddressLineTwo, 35, options)
}

// FormatAddressLineThree returns FinancialInstitution.Address.AddressLineThree formatted according to the FormatOptions
func (bfi *BeneficiaryFI) FormatAddressLineThree(options FormatOptions) string {
	return bfi.formatAlphaField(bfi.FinancialInstitution.Address.AddressLineThree, 35, options)
}
