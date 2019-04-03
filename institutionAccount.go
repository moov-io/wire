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
func NewInstitutionAccount() InstitutionAccount {
	ia := InstitutionAccount{
		tag: TagInstitutionAccount,
	}
	return ia
}

// Parse takes the input string and parses the InstitutionAccount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ia *InstitutionAccount) Parse(record string) {
	ia.tag = record[:6]
	ia.CoverPayment.SwiftFieldTag = ia.parseStringField(record[6:11])
	ia.CoverPayment.SwiftLineOne = ia.parseStringField(record[11:46])
	ia.CoverPayment.SwiftLineTwo = ia.parseStringField(record[46:81])
	ia.CoverPayment.SwiftLineThree = ia.parseStringField(record[81:116])
	ia.CoverPayment.SwiftLineFour = ia.parseStringField(record[116:151])
	ia.CoverPayment.SwiftLineFive = ia.parseStringField(record[151:186])
}

// String writes InstitutionAccount
func (ia *InstitutionAccount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(186)
	buf.WriteString(ia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on InstitutionAccount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ia *InstitutionAccount) Validate() error {
	if err := ia.fieldInclusion(); err != nil {
		return err
	}
	if err := ia.isAlphanumeric(ia.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ia.CoverPayment.SwiftFieldTag)
	}
	if err := ia.isAlphanumeric(ia.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ia.CoverPayment.SwiftLineOne)
	}
	if err := ia.isAlphanumeric(ia.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ia.CoverPayment.SwiftLineTwo)
	}
	if err := ia.isAlphanumeric(ia.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ia.CoverPayment.SwiftLineThree)
	}
	if err := ia.isAlphanumeric(ia.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ia.CoverPayment.SwiftLineFour)
	}
	if err := ia.isAlphanumeric(ia.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, ia.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ia *InstitutionAccount) fieldInclusion() error {
	return nil
}
