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
	fiifi := FIIntermediaryFI{
		tag: TagFIIntermediaryFI,
	}
	return fiifi
}

// Parse takes the input string and parses the FIIntermediaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifi *FIIntermediaryFI) Parse(record string) {
	fiifi.tag = record[:6]
	fiifi.FIToFI.LineOne = fiifi.parseStringField(record[6:36])
	fiifi.FIToFI.LineTwo = fiifi.parseStringField(record[36:69])
	fiifi.FIToFI.LineThree = fiifi.parseStringField(record[69:104])
	fiifi.FIToFI.LineFour = fiifi.parseStringField(record[104:139])
	fiifi.FIToFI.LineFive = fiifi.parseStringField(record[139:174])
	fiifi.FIToFI.LineSix = fiifi.parseStringField(record[174:209])
}

// String writes FIIntermediaryFI
func (fiifi *FIIntermediaryFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(209)
	buf.WriteString(fiifi.tag)
	buf.WriteString(fiifi.LineOneField())
	buf.WriteString(fiifi.LineTwoField())
	buf.WriteString(fiifi.LineThreeField())
	buf.WriteString(fiifi.LineFourField())
	buf.WriteString(fiifi.LineFiveField())
	buf.WriteString(fiifi.LineSixField())
	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiifi *FIIntermediaryFI) Validate() error {
	if err := fiifi.fieldInclusion(); err != nil {
		return err
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fiifi.FIToFI.LineOne)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiifi.FIToFI.LineTwo)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fiifi.FIToFI.LineThree)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fiifi.FIToFI.LineFour)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fiifi.FIToFI.LineFive)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fiifi.FIToFI.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (fiifi *FIIntermediaryFI) fieldInclusion() error {
	return nil
}

// LineOneField gets a string of the LineOne field
func (fiifi *FIIntermediaryFI) LineOneField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (fiifi *FIIntermediaryFI) LineTwoField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (fiifi *FIIntermediaryFI) LineThreeField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (fiifi *FIIntermediaryFI) LineFourField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (fiifi *FIIntermediaryFI) LineFiveField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (fiifi *FIIntermediaryFI) LineSixField() string {
	return fiifi.alphaField(firfi.FIToFI.LineSix, 35)
}
