// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// OriginatorToBeneficiary is the OriginatorToBeneficiary of the wire
type OriginatorToBeneficiary struct {
	// tag
	tag string
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`
	// LineFour
	LineFour string `json:"lineFour,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginatorToBeneficiary returns a new OriginatorToBeneficiary
func NewOriginatorToBeneficiary() OriginatorToBeneficiary {
	ob := OriginatorToBeneficiary{
		tag: TagOriginatorToBeneficiary,
	}
	return ob
}

// Parse takes the input string and parses the OriginatorToBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ob *OriginatorToBeneficiary) Parse(record string) {
	ob.tag = record[:6]
	ob.LineOne = ob.parseStringField(record[6:41])
	ob.LineTwo = ob.parseStringField(record[41:76])
	ob.LineThree = ob.parseStringField(record[76:111])
	ob.LineFour = ob.parseStringField(record[111:146])
}

// String writes OriginatorToBeneficiary
func (ob *OriginatorToBeneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(146)
	buf.WriteString(ob.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OriginatorToBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ob *OriginatorToBeneficiary) Validate() error {
	if err := ob.fieldInclusion(); err != nil {
		return err
	}
	if err := ob.isAlphanumeric(ob.LineOne); err != nil {
		return fieldError("LineOne", err, ob.LineOne)
	}
	if err := ob.isAlphanumeric(ob.LineTwo); err != nil {
		return fieldError("LineTwo", err, ob.LineTwo)
	}
	if err := ob.isAlphanumeric(ob.LineThree); err != nil {
		return fieldError("LineThree", err, ob.LineThree)
	}
	if err := ob.isAlphanumeric(ob.LineFour); err != nil {
		return fieldError("LineFour", err, ob.LineFour)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ob *OriginatorToBeneficiary) fieldInclusion() error {
	return nil
}
