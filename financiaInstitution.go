// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "slices"

var (
	financialInstitutionIDCodes = []string{
		SWIFTBankIdentifierCode,
		CHIPSParticipant,
		DemandDepositAccountNumber,
		FEDRoutingNumber,
		CHIPSIdentifier,
	}
)

// FinancialInstitution is demographic information for a financial institution
type FinancialInstitution struct {
	// IdentificationCode:  * `B` - SWIFT Bank Identifier Code (BIC) * `C` - CHIPS Participant * `D` - Demand Deposit Account (DDA) Number * `F` - Fed Routing Number * `T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * `U` - CHIPS Identifier
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name string `json:"name"`
	// Address
	Address Address `json:"address"`

	validator
}

func (fi FinancialInstitution) Validate() error {
	if err := fi.fieldInclusion(); err != nil {
		return err
	}

	// if ID Code is present, make sure it's a valid value
	if fi.IdentificationCode != "" && !slices.Contains(financialInstitutionIDCodes, fi.IdentificationCode) {
		return fieldError("IdentificationCode", ErrIdentificationCode, fi.IdentificationCode)
	}

	if err := fi.isAlphanumeric(fi.Identifier); err != nil {
		return fieldError("Identifier", err, fi.Identifier)
	}
	if err := fi.isAlphanumeric(fi.Name); err != nil {
		return fieldError("Name", err, fi.Name)
	}
	if err := fi.isAlphanumeric(fi.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, fi.Address.AddressLineOne)
	}
	if err := fi.isAlphanumeric(fi.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, fi.Address.AddressLineTwo)
	}
	if err := fi.isAlphanumeric(fi.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, fi.Address.AddressLineThree)
	}

	return nil
}

func (fi FinancialInstitution) fieldInclusion() error {
	// if Identifier is present, IdentificationCode must be provided.
	if fi.Identifier != "" && fi.IdentificationCode == "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}

	// If IdentificationCode is present, Identifier must be present
	if fi.IdentificationCode != "" && fi.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}

	return nil
}
