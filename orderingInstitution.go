// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func NewOrderingInstitution() OrderingInstitution {
	oi := OrderingInstitution{
		tag: TagOrderingInstitution,
	}
	return oi
}

// Parse takes the input string and parses the OrderingInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oi *OrderingInstitution) Parse(record string) {
	oi.tag = record[:6]
	oi.CoverPayment.SwiftFieldTag = oi.parseStringField(record[6:11])
	oi.CoverPayment.SwiftLineOne = oi.parseStringField(record[11:46])
	oi.CoverPayment.SwiftLineTwo = oi.parseStringField(record[46:81])
	oi.CoverPayment.SwiftLineThree = oi.parseStringField(record[81:116])
	oi.CoverPayment.SwiftLineFour = oi.parseStringField(record[116:151])
	oi.CoverPayment.SwiftLineFive = oi.parseStringField(record[151:186])
}

// String writes OrderingInstitution
func (oi *OrderingInstitution) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(186)
	buf.WriteString(oi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OrderingInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oi *OrderingInstitution) Validate() error {
	if err := oi.fieldInclusion(); err != nil {
		return err
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
	return nil
}
