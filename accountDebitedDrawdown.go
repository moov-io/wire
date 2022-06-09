// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (debitDD *AccountDebitedDrawdown) Parse(record string) error {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagMinLengthErr(9, len(record))
	}

	debitDD.tag = record[:6]
	debitDD.IdentificationCode = record[6:7]
	length := 7

	value, read, err := debitDD.parseVariableStringField(record[length:], 34)
	if err != nil {
		return fieldError("Identifier", err)
	}
	debitDD.Identifier = value
	length += read

	value, read, err = debitDD.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Name", err)
	}
	debitDD.Name = value
	length += read

	value, read, err = debitDD.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	debitDD.Address.AddressLineOne = value
	length += read

	value, read, err = debitDD.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	debitDD.Address.AddressLineTwo = value
	length += read

	value, read, err = debitDD.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	debitDD.Address.AddressLineThree = value
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (debitDD *AccountDebitedDrawdown) UnmarshalJSON(data []byte) error {
	type Alias AccountDebitedDrawdown
	aux := struct {
		*Alias
	}{
		(*Alias)(debitDD),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	debitDD.tag = TagAccountDebitedDrawdown
	return nil
}

// String writes AccountDebitedDrawdown
func (debitDD *AccountDebitedDrawdown) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(debitDD.tag)
	buf.WriteString(debitDD.IdentificationCodeField())
	buf.WriteString(debitDD.IdentifierField(options...))
	buf.WriteString(debitDD.NameField(options...))
	buf.WriteString(debitDD.AddressLineOneField(options...))
	buf.WriteString(debitDD.AddressLineTwoField(options...))
	buf.WriteString(debitDD.AddressLineThreeField(options...))

	if debitDD.parseFirstOption(options) {
		return debitDD.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on AccountDebitedDrawdown and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (debitDD *AccountDebitedDrawdown) Validate() error {
	if err := debitDD.fieldInclusion(); err != nil {
		return err
	}
	if debitDD.tag != TagAccountDebitedDrawdown {
		return fieldError("tag", ErrValidTagForType, debitDD.tag)
	}
	if err := debitDD.isIdentificationCode(debitDD.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, debitDD.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch debitDD.IdentificationCode {
	case
		DemandDepositAccountNumber:
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
	if debitDD.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (debitDD *AccountDebitedDrawdown) IdentificationCodeField() string {
	return debitDD.alphaField(debitDD.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (debitDD *AccountDebitedDrawdown) IdentifierField(options ...bool) string {
	return debitDD.alphaVariableField(debitDD.Identifier, 34, debitDD.parseFirstOption(options))
}

// NameField gets a string of the Name field
func (debitDD *AccountDebitedDrawdown) NameField(options ...bool) string {
	return debitDD.alphaVariableField(debitDD.Name, 35, debitDD.parseFirstOption(options))
}

// AddressLineOneField gets a string of AddressLineOne field
func (debitDD *AccountDebitedDrawdown) AddressLineOneField(options ...bool) string {
	return debitDD.alphaVariableField(debitDD.Address.AddressLineOne, 35, debitDD.parseFirstOption(options))
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (debitDD *AccountDebitedDrawdown) AddressLineTwoField(options ...bool) string {
	return debitDD.alphaVariableField(debitDD.Address.AddressLineTwo, 35, debitDD.parseFirstOption(options))
}

// AddressLineThreeField gets a string of AddressLineThree field
func (debitDD *AccountDebitedDrawdown) AddressLineThreeField(options ...bool) string {
	return debitDD.alphaVariableField(debitDD.Address.AddressLineThree, 35, debitDD.parseFirstOption(options))
}
