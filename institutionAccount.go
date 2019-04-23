// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// InstitutionAccount is the institution account
type InstitutionAccount struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInstitutionAccount returns a new InstitutionAccount
func NewInstitutionAccount() *InstitutionAccount {
	iAccount := &InstitutionAccount{
		tag: TagInstitutionAccount,
	}
	return iAccount
}

// Parse takes the input string and parses the InstitutionAccount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (iAccount *InstitutionAccount) Parse(record string) {
	iAccount.tag = record[:6]
	iAccount.CoverPayment.SwiftFieldTag = iAccount.parseStringField(record[6:11])
	iAccount.CoverPayment.SwiftLineOne = iAccount.parseStringField(record[11:46])
	iAccount.CoverPayment.SwiftLineTwo = iAccount.parseStringField(record[46:81])
	iAccount.CoverPayment.SwiftLineThree = iAccount.parseStringField(record[81:116])
	iAccount.CoverPayment.SwiftLineFour = iAccount.parseStringField(record[116:151])
	iAccount.CoverPayment.SwiftLineFive = iAccount.parseStringField(record[151:186])
}

// String writes InstitutionAccount
func (iAccount *InstitutionAccount) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(iAccount.tag)
	buf.WriteString(iAccount.SwiftFieldTagField())
	buf.WriteString(iAccount.SwiftLineOneField())
	buf.WriteString(iAccount.SwiftLineTwoField())
	buf.WriteString(iAccount.SwiftLineThreeField())
	buf.WriteString(iAccount.SwiftLineFourField())
	buf.WriteString(iAccount.SwiftLineFiveField())
	return buf.String()
}

// Validate performs WIRE format rule checks on InstitutionAccount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (iAccount *InstitutionAccount) Validate() error {
	if err := iAccount.fieldInclusion(); err != nil {
		return err
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, iAccount.CoverPayment.SwiftFieldTag)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, iAccount.CoverPayment.SwiftLineOne)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, iAccount.CoverPayment.SwiftLineTwo)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, iAccount.CoverPayment.SwiftLineThree)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, iAccount.CoverPayment.SwiftLineFour)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, iAccount.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (iAccount *InstitutionAccount) fieldInclusion() error {
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (iAccount *InstitutionAccount) SwiftFieldTagField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (iAccount *InstitutionAccount) SwiftLineOneField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (iAccount *InstitutionAccount) SwiftLineTwoField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (iAccount *InstitutionAccount) SwiftLineThreeField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (iAccount *InstitutionAccount) SwiftLineFourField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (iAccount *InstitutionAccount) SwiftLineFiveField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineFive, 35)
}
