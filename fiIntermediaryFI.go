// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIIntermediaryFI is the financial institution intermediary financial institution
type FIIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFI returns a new FIIntermediaryFI
func NewFIIntermediaryFI() FIIntermediaryFI {
	ifi := FIIntermediaryFI{
		tag: TagFIIntermediaryFI,
	}
	return ifi
}

// Parse takes the input string and parses the FIIntermediaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ifi *FIIntermediaryFI) Parse(record string) {
	ifi.tag = record[:6]
	ifi.FIToFI.LineOne = ifi.parseStringField(record[6:36])
	ifi.FIToFI.LineTwo = ifi.parseStringField(record[36:69])
	ifi.FIToFI.LineThree = ifi.parseStringField(record[69:104])
	ifi.FIToFI.LineFour = ifi.parseStringField(record[104:139])
	ifi.FIToFI.LineFive = ifi.parseStringField(record[139:174])
	ifi.FIToFI.LineSix = ifi.parseStringField(record[174:209])
}

// String writes FIIntermediaryFI
func (ifi *FIIntermediaryFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(209)
	buf.WriteString(ifi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ifi *FIIntermediaryFI) Validate() error {
	if err := ifi.fieldInclusion(); err != nil {
		return err
	}
	if err:= ifi.isAlphanumeric(ifi.FIToFI.LineOne); err!= nil {
		return fieldError("LineOne", err, ifi.FIToFI.LineOne)
	}
	if err:= ifi.isAlphanumeric(ifi.FIToFI.LineTwo); err!= nil {
		return fieldError("LineTwo", err, ifi.FIToFI.LineTwo)
	}
	if err:= ifi.isAlphanumeric(ifi.FIToFI.LineThree); err!= nil {
		return fieldError("LineThree", err, ifi.FIToFI.LineThree)
	}
	if err:= ifi.isAlphanumeric(ifi.FIToFI.LineFour); err!= nil {
		return fieldError("LineFour", err, ifi.FIToFI.LineFour)
	}
	if err:= ifi.isAlphanumeric(ifi.FIToFI.LineFive); err!= nil {
		return fieldError("LineFive", err, ifi.FIToFI.LineFive)
	}
	if err:= ifi.isAlphanumeric(ifi.FIToFI.LineSix); err!= nil {
		return fieldError("LineSix", err, ifi.FIToFI.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ifi *FIIntermediaryFI) fieldInclusion() error {
	return nil
}
