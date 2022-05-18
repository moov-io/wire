// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
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

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the FinancialInstitution values
func (f *FinancialInstitution) Parse(record string) (length int, err error) {

	var read int

	f.IdentificationCode = f.parseStringField(record[0:1])

	if f.Identifier, read, err = f.parseVariableStringField(record[length:], 34); err != nil {
		return 0, fieldError("Identifier", err)
	}
	length += read

	if f.Name, read, err = f.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("Name", err)
	}
	length += read

	if read, err = f.Address.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return
}

// String writes FinancialInstitution
func (f *FinancialInstitution) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(175)

	buf.WriteString(f.IdentificationCodeField())
	buf.WriteString(f.IdentifierField(isVariable))
	buf.WriteString(f.NameField(isVariable))
	buf.WriteString(f.Address.String(isVariable))

	return buf.String()
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (f *FinancialInstitution) IdentificationCodeField() string {
	return f.alphaField(f.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (f *FinancialInstitution) IdentifierField(isVariable bool) string {
	return f.alphaVariableField(f.Identifier, 34, isVariable)
}

// NameField gets a string of the Name field
func (f *FinancialInstitution) NameField(isVariable bool) string {
	return f.alphaVariableField(f.Name, 35, isVariable)
}
