// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &AccountDebitedDrawdown{}

// AccountDebitedDrawdown is the account which is debited in a drawdown
type AccountDebitedDrawdown struct {
	// tag
	tag string
	// Identification Code * `D` - Debit
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name string `json:"name"`
	// Address
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
func (debitDD *AccountDebitedDrawdown) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if debitDD.tag, read, err = debitDD.parseTag(record); err != nil {
		return 0, fieldError("AccountDebitedDrawdown.Tag", err)
	}
	length += read

	debitDD.IdentificationCode = debitDD.parseStringField(record[length : length+1])
	length += 1

	if debitDD.Identifier, read, err = debitDD.parseVariableStringField(record[length:], 34); err != nil {
		return 0, fieldError("Identifier", err)
	}
	length += read

	if debitDD.Name, read, err = debitDD.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("Name", err)
	}
	length += read

	if read, err = debitDD.Address.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(debitDD.tag)

	buf.WriteString(debitDD.IdentificationCodeField())
	buf.WriteString(debitDD.IdentifierField(isCompressed))
	buf.WriteString(debitDD.NameField(isCompressed))
	buf.WriteString(debitDD.Address.String(isCompressed))

	return buf.String()
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
func (debitDD *AccountDebitedDrawdown) IdentifierField(isCompressed bool) string {
	return debitDD.alphaVariableField(debitDD.Identifier, 34, isCompressed)
}

// NameField gets a string of the Name field
func (debitDD *AccountDebitedDrawdown) NameField(isCompressed bool) string {
	return debitDD.alphaVariableField(debitDD.Name, 35, isCompressed)
}
