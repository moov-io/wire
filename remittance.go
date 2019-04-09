// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// Remittance is the remittance information
type Remittance struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittance returns a new Remittance
func NewRemittance() Remittance {
	ri := Remittance{
		tag: TagRemittance,
	}
	return ri
}

// Parse takes the input string and parses the Remittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ri *Remittance) Parse(record string) {
	ri.tag = record[:6]
	ri.CoverPayment.SwiftFieldTag = ri.parseStringField(record[6:11])
	ri.CoverPayment.SwiftLineOne = ri.parseStringField(record[11:46])
	ri.CoverPayment.SwiftLineTwo = ri.parseStringField(record[46:81])
	ri.CoverPayment.SwiftLineThree = ri.parseStringField(record[81:116])
	ri.CoverPayment.SwiftLineFour = ri.parseStringField(record[116:151])
}

// String writes Remittance
func (ri *Remittance) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(151)
	buf.WriteString(ri.tag)
	buf.WriteString(ri.SwiftFieldTagField())
	buf.WriteString(ri.SwiftLineOneField())
	buf.WriteString(ri.SwiftLineTwoField())
	buf.WriteString(ri.SwiftLineThreeField())
	buf.WriteString(ri.SwiftLineFourField())
	return buf.String()
}

// Validate performs WIRE format rule checks on Remittance and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ri *Remittance) Validate() error {
	if err := ri.fieldInclusion(); err != nil {
		return err
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ri.CoverPayment.SwiftFieldTag)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ri.CoverPayment.SwiftLineOne)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ri.CoverPayment.SwiftLineTwo)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ri.CoverPayment.SwiftLineThree)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ri.CoverPayment.SwiftLineFour)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ri *Remittance) fieldInclusion() error {
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (ri *Remittance) SwiftFieldTagField() string {
	return ri.alphaField(ri.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ri *Remittance) SwiftLineOneField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ri *Remittance) SwiftLineTwoField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ri *Remittance) SwiftLineThreeField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ri *Remittance) SwiftLineFourField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineFour, 35)
}
