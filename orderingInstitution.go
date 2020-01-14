// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// OrderingInstitution is the ordering institution
type OrderingInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingInstitution returns a new OrderingInstitution
func NewOrderingInstitution() *OrderingInstitution {
	oi := &OrderingInstitution{
		tag: TagOrderingInstitution,
	}
	return oi
}

// Parse takes the input string and parses the OrderingInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oi *OrderingInstitution) Parse(record string) error {
	if utf8.RuneCountInString(record) != 186 {
		return NewTagWrongLengthErr(186, len(record))
	}
	oi.tag = record[:6]
	oi.CoverPayment.SwiftFieldTag = oi.parseStringField(record[6:11])
	oi.CoverPayment.SwiftLineOne = oi.parseStringField(record[11:46])
	oi.CoverPayment.SwiftLineTwo = oi.parseStringField(record[46:81])
	oi.CoverPayment.SwiftLineThree = oi.parseStringField(record[81:116])
	oi.CoverPayment.SwiftLineFour = oi.parseStringField(record[116:151])
	oi.CoverPayment.SwiftLineFive = oi.parseStringField(record[151:186])
	return nil
}

// String writes OrderingInstitution
func (oi *OrderingInstitution) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(oi.tag)
	buf.WriteString(oi.SwiftFieldTagField())
	buf.WriteString(oi.SwiftLineOneField())
	buf.WriteString(oi.SwiftLineTwoField())
	buf.WriteString(oi.SwiftLineThreeField())
	buf.WriteString(oi.SwiftLineFourField())
	buf.WriteString(oi.SwiftLineFiveField())
	return buf.String()
}

// Validate performs WIRE format rule checks on OrderingInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oi *OrderingInstitution) Validate() error {
	if err := oi.fieldInclusion(); err != nil {
		return err
	}
	if oi.tag != TagOrderingInstitution {
		return fieldError("tag", ErrValidTagForType, oi.tag)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, oi.CoverPayment.SwiftFieldTag)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, oi.CoverPayment.SwiftLineOne)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, oi.CoverPayment.SwiftLineTwo)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, oi.CoverPayment.SwiftLineThree)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, oi.CoverPayment.SwiftLineFour)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, oi.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oi *OrderingInstitution) fieldInclusion() error {
	if oi.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, oi.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (oi *OrderingInstitution) SwiftFieldTagField() string {
	return oi.alphaField(oi.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (oi *OrderingInstitution) SwiftLineOneField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (oi *OrderingInstitution) SwiftLineTwoField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (oi *OrderingInstitution) SwiftLineThreeField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (oi *OrderingInstitution) SwiftLineFourField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (oi *OrderingInstitution) SwiftLineFiveField() string {
	return oi.alphaField(oi.CoverPayment.SwiftLineFive, 35)
}
