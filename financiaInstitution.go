// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (f *FinancialInstitution) Parse(record string) int {

	length := 1
	read := 0

	f.IdentificationCode = f.parseStringField(record[0:1])

	f.Identifier, read = f.parseVariableStringField(record[length:], 34)
	length += read

	f.Name, read = f.parseVariableStringField(record[length:], 35)
	length += read

	f.Address.AddressLineOne, read = f.parseVariableStringField(record[length:], 35)
	length += read

	f.Address.AddressLineTwo, read = f.parseVariableStringField(record[length:], 35)
	length += read

	f.Address.AddressLineThree, read = f.parseVariableStringField(record[length:], 35)
	length += read

	return length
}

// String writes FinancialInstitution
func (f *FinancialInstitution) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(175)

	buf.WriteString(f.IdentificationCodeField())
	buf.WriteString(f.IdentifierField(isVariable))
	buf.WriteString(f.NameField(isVariable))
	buf.WriteString(f.AddressLineOneField(isVariable))
	buf.WriteString(f.AddressLineTwoField(isVariable))
	buf.WriteString(f.AddressLineThreeField(isVariable))

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

// AddressLineOneField gets a string of AddressLineOne field
func (f *FinancialInstitution) AddressLineOneField(isVariable bool) string {
	return f.alphaVariableField(f.Address.AddressLineOne, 35, isVariable)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (f *FinancialInstitution) AddressLineTwoField(isVariable bool) string {
	return f.alphaVariableField(f.Address.AddressLineTwo, 35, isVariable)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (f *FinancialInstitution) AddressLineThreeField(isVariable bool) string {
	return f.alphaVariableField(f.Address.AddressLineThree, 35, isVariable)
}
