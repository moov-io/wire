// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// AccountDebitedDrawdown is the account which is debited in a drawdown
type AccountDebitedDrawdown struct {
	// tag
	tag string
	// Identification Code * `D` - Debit
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name    string  `json:"name"`
	Address Address `json:"address,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAccountDebitedDrawdown returns a new AccountDebitedDrawdown
func NewAccountDebitedDrawdown() *AccountDebitedDrawdown {
	debitDD := &AccountDebitedDrawdown{
		tag: TagAccountDebitedDrawdown,
	}
	return debitDD
}

// Parse takes the input string and parses the AccountDebitedDrawdown values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (debitDD *AccountDebitedDrawdown) Parse(record string) {
	debitDD.tag = record[:6]
	debitDD.IdentificationCode = debitDD.parseStringField(record[6:7])
	debitDD.Identifier = debitDD.parseStringField(record[7:41])
	debitDD.Name = debitDD.parseStringField(record[41:76])
	debitDD.Address.AddressLineOne = debitDD.parseStringField(record[76:111])
	debitDD.Address.AddressLineTwo = debitDD.parseStringField(record[111:146])
	debitDD.Address.AddressLineThree = debitDD.parseStringField(record[146:181])
}

// String writes AccountDebitedDrawdown
func (debitDD *AccountDebitedDrawdown) String() string {
	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(debitDD.tag)
	buf.WriteString(debitDD.IdentificationCodeField())
	buf.WriteString(debitDD.IdentifierField())
	buf.WriteString(debitDD.NameField())
	buf.WriteString(debitDD.AddressLineOneField())
	buf.WriteString(debitDD.AddressLineTwoField())
	buf.WriteString(debitDD.AddressLineThreeField())
	return buf.String()
}

// Validate performs WIRE format rule checks on AccountDebitedDrawdown and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (debitDD *AccountDebitedDrawdown) Validate() error {
	if err := debitDD.fieldInclusion(); err != nil {
		return err
	}
	if err := debitDD.isIdentificationCode(debitDD.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, debitDD.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch debitDD.IdentificationCode {
	case
		"D":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, debitDD.IdentificationCode)
	}
	if err := debitDD.isAlphanumeric(debitDD.Identifier); err != nil {
		return fieldError("Identifier", err, debitDD.Identifier)
	}
	if err := debitDD.isAlphanumeric(debitDD.Name); err != nil {
		return fieldError("Name", err, debitDD.Name)
	}
	if err := debitDD.isAlphanumeric(debitDD.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, debitDD.Address.AddressLineOne)
	}
	if err := debitDD.isAlphanumeric(debitDD.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, debitDD.Address.AddressLineTwo)
	}
	if err := debitDD.isAlphanumeric(debitDD.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, debitDD.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (debitDD *AccountDebitedDrawdown) fieldInclusion() error {
	if debitDD.IdentificationCode == "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	if debitDD.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if debitDD.Name== "" {
		return fieldError("Name", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (debitDD *AccountDebitedDrawdown) IdentificationCodeField() string {
	return debitDD.alphaField(debitDD.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (debitDD *AccountDebitedDrawdown) IdentifierField() string {
	return debitDD.alphaField(debitDD.Identifier, 34)
}

// NameField gets a string of the Name field
func (debitDD *AccountDebitedDrawdown) NameField() string {
	return debitDD.alphaField(debitDD.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (debitDD *AccountDebitedDrawdown) AddressLineOneField() string {
	return debitDD.alphaField(debitDD.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (debitDD *AccountDebitedDrawdown) AddressLineTwoField() string {
	return debitDD.alphaField(debitDD.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (debitDD *AccountDebitedDrawdown) AddressLineThreeField() string {
	return debitDD.alphaField(debitDD.Address.AddressLineThree, 35)
}
