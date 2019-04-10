// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// Beneficiary is the beneficiary of the wire
type Beneficiary struct {
	// tag
	tag string
	// Personal
	Personal Personal `json:"personal,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiary returns a new Beneficiary
func NewBeneficiary() *Beneficiary {
	b := &Beneficiary{
		tag: TagBeneficiary,
	}
	return b
}

// Parse takes the input string and parses the Beneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (b *Beneficiary) Parse(record string) {
	b.tag = record[:6]
	b.Personal.IdentificationCode = b.parseStringField(record[6:7])
	b.Personal.Identifier = b.parseStringField(record[7:41])
	b.Personal.Name = b.parseStringField(record[41:76])
	b.Personal.Address.AddressLineOne = b.parseStringField(record[76:111])
	b.Personal.Address.AddressLineTwo = b.parseStringField(record[111:146])
	b.Personal.Address.AddressLineThree = b.parseStringField(record[146:181])
}

// String writes Beneficiary
func (b *Beneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(181)
	buf.WriteString(b.tag)
	buf.WriteString(b.IdentificationCodeField())
	buf.WriteString(b.IdentifierField())
	buf.WriteString(b.AddressLineOneField())
	buf.WriteString(b.AddressLineTwoField())
	buf.WriteString(b.AddressLineThreeField())
	return buf.String()
}

// Validate performs WIRE format rule checks on Beneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (b *Beneficiary) Validate() error {
	if err := b.fieldInclusion(); err != nil {
		return err
	}
	// Can be any Identification Code
	if err := b.isIdentificationCode(b.Personal.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, b.Personal.IdentificationCode)
	}
	if err := b.isAlphanumeric(b.Personal.Identifier); err != nil {
		return fieldError("Identifier", err, b.Personal.Identifier)
	}
	if err := b.isAlphanumeric(b.Personal.Name); err != nil {
		return fieldError("Name", err, b.Personal.Name)
	}
	if err := b.isAlphanumeric(b.Personal.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, b.Personal.Address.AddressLineOne)
	}
	if err := b.isAlphanumeric(b.Personal.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, b.Personal.Address.AddressLineTwo)
	}
	if err := b.isAlphanumeric(b.Personal.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, b.Personal.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (b *Beneficiary) fieldInclusion() error {
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (b *Beneficiary) IdentificationCodeField() string {
	return b.alphaField(b.Personal.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (b *Beneficiary) IdentifierField() string {
	return b.alphaField(b.Personal.Identifier, 34)
}

// NameField gets a string of the Name field
func (b *Beneficiary) NameField() string {
	return b.alphaField(b.Personal.Name, 35)
}

// AddressLineOneField gets a string of AddressLineOne field
func (b *Beneficiary) AddressLineOneField() string {
	return b.alphaField(b.Personal.Address.AddressLineOne, 35)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (b *Beneficiary) AddressLineTwoField() string {
	return b.alphaField(b.Personal.Address.AddressLineTwo, 35)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (b *Beneficiary) AddressLineThreeField() string {
	return b.alphaField(b.Personal.Address.AddressLineThree, 35)
}
