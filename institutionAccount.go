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
}

// String writes InstitutionAccount
func (ia *InstitutionAccount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(180)
	buf.WriteString(ia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on InstitutionAccount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ia *InstitutionAccount) Validate() error {
	if err := ia.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ia *InstitutionAccount) fieldInclusion() error {
	return nil
}
