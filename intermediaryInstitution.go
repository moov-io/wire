// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// IntermediaryInstitution is the intermediary institution
type IntermediaryInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewIntermediaryInstitution returns a new IntermediaryInstitution
func NewIntermediaryInstitution() IntermediaryInstitution {
	ii := IntermediaryInstitution{
		tag: TagIntermediaryInstitution,
	}
	return ii
}

// Parse takes the input string and parses the IntermediaryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ii *IntermediaryInstitution) Parse(record string) {
	ii.tag = record[:6]
	ii.CoverPayment.SwiftFieldTag = ii.parseStringField(record[6:11])
	ii.CoverPayment.SwiftLineOne = ii.parseStringField(record[11:46])
	ii.CoverPayment.SwiftLineTwo = ii.parseStringField(record[46:81])
	ii.CoverPayment.SwiftLineThree = ii.parseStringField(record[81:116])
	ii.CoverPayment.SwiftLineFour = ii.parseStringField(record[116:151])
	ii.CoverPayment.SwiftLineFive = ii.parseStringField(record[151:186])
}

// String writes IntermediaryInstitution
func (ii *IntermediaryInstitution) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(186)
	buf.WriteString(ii.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on IntermediaryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ii *IntermediaryInstitution) Validate() error {
	if err := ii.fieldInclusion(); err != nil {
		return err
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ii.CoverPayment.SwiftFieldTag)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ii.CoverPayment.SwiftLineOne)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ii.CoverPayment.SwiftLineTwo)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ii.CoverPayment.SwiftLineThree)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ii.CoverPayment.SwiftLineFour)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, ii.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ii *IntermediaryInstitution) fieldInclusion() error {
	return nil
}
