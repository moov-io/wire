// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

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
func NewIntermediaryInstitution() *IntermediaryInstitution {
	ii := &IntermediaryInstitution{
		tag: TagIntermediaryInstitution,
	}
	return ii
}

// Parse takes the input string and parses the IntermediaryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ii *IntermediaryInstitution) Parse(record string) error {
	if utf8.RuneCountInString(record) != 186 {
		return NewTagWrongLengthErr(186, len(record))
	}
	ii.tag = record[:6]
	ii.CoverPayment.SwiftFieldTag = ii.parseStringField(record[6:11])
	ii.CoverPayment.SwiftLineOne = ii.parseStringField(record[11:46])
	ii.CoverPayment.SwiftLineTwo = ii.parseStringField(record[46:81])
	ii.CoverPayment.SwiftLineThree = ii.parseStringField(record[81:116])
	ii.CoverPayment.SwiftLineFour = ii.parseStringField(record[116:151])
	ii.CoverPayment.SwiftLineFive = ii.parseStringField(record[151:186])
	return nil
}

// String writes IntermediaryInstitution
func (ii *IntermediaryInstitution) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(ii.tag)
	buf.WriteString(ii.SwiftFieldTagField())
	buf.WriteString(ii.SwiftLineOneField())
	buf.WriteString(ii.SwiftLineTwoField())
	buf.WriteString(ii.SwiftLineThreeField())
	buf.WriteString(ii.SwiftLineFourField())
	buf.WriteString(ii.SwiftLineFiveField())
	return buf.String()
}

// Validate performs WIRE format rule checks on IntermediaryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ii *IntermediaryInstitution) Validate() error {
	if err := ii.fieldInclusion(); err != nil {
		return err
	}
	if ii.tag != TagIntermediaryInstitution {
		return fieldError("tag", ErrValidTagForType, ii.tag)
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
	if ii.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, ii.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (ii *IntermediaryInstitution) SwiftFieldTagField() string {
	return ii.alphaField(ii.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ii *IntermediaryInstitution) SwiftLineOneField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ii *IntermediaryInstitution) SwiftLineTwoField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ii *IntermediaryInstitution) SwiftLineThreeField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ii *IntermediaryInstitution) SwiftLineFourField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (ii *IntermediaryInstitution) SwiftLineFiveField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineFive, 35)
}
