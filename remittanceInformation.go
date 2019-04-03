// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// RemittanceInformation is the remittance information
type RemittanceInformation struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceInformation returns a new RemittanceInformation
func NewRemittanceInformation() RemittanceInformation {
	ri := RemittanceInformation{
		tag: TagRemittanceInformation,
	}
	return ri
}

// Parse takes the input string and parses the RemittanceInformation values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ri *RemittanceInformation) Parse(record string) {
	ri.tag = record[:6]
	ri.CoverPayment.SwiftFieldTag = ri.parseStringField(record[6:11])
	ri.CoverPayment.SwiftLineOne = ri.parseStringField(record[11:46])
	ri.CoverPayment.SwiftLineTwo = ri.parseStringField(record[46:81])
	ri.CoverPayment.SwiftLineThree = ri.parseStringField(record[81:116])
	ri.CoverPayment.SwiftLineFour = ri.parseStringField(record[116:151])
}

// String writes RemittanceInformation
func (ri *RemittanceInformation) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(151)
	buf.WriteString(ri.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceInformation and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ri *RemittanceInformation) Validate() error {
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
func (ri *RemittanceInformation) fieldInclusion() error {
	return nil
}
