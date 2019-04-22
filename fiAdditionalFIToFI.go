// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIAdditionalFIToFI is the financial institution beneficiary financial institution
type FIAdditionalFIToFI struct {
	// tag
	tag string
	// AdditionalFiToFi is additional financial institution to financial institution information
	AdditionalFIToFI AdditionalFIToFI `json:"additionalFiToFi,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIAdditionalFIToFI returns a new FIAdditionalFIToFI
func NewFIAdditionalFIToFI() *FIAdditionalFIToFI {
	fifi := &FIAdditionalFIToFI{
		tag: TagFIAdditionalFIToFI,
	}
	return fifi
}

// Parse takes the input string and parses the FIAdditionalFIToFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fifi *FIAdditionalFIToFI) Parse(record string) {
	fifi.tag = record[:6]
	fifi.AdditionalFIToFI.LineOne = fifi.parseStringField(record[6:36])
	fifi.AdditionalFIToFI.LineTwo = fifi.parseStringField(record[36:71])
	fifi.AdditionalFIToFI.LineThree = fifi.parseStringField(record[71:106])
	fifi.AdditionalFIToFI.LineFour = fifi.parseStringField(record[106:141])
	fifi.AdditionalFIToFI.LineFive = fifi.parseStringField(record[141:176])
	fifi.AdditionalFIToFI.LineSix = fifi.parseStringField(record[176:211])
}

// String writes FIAdditionalFIToFI
func (fifi *FIAdditionalFIToFI) String() string {
	var buf strings.Builder
	buf.Grow(211)
	buf.WriteString(fifi.tag)
	buf.WriteString(fifi.LineOneField())
	buf.WriteString(fifi.LineTwoField())
	buf.WriteString(fifi.LineThreeField())
	buf.WriteString(fifi.LineFourField())
	buf.WriteString(fifi.LineFiveField())
	buf.WriteString(fifi.LineSixField())
	return buf.String()
}

// Validate performs WIRE format rule checks on FIAdditionalFIToFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fifi *FIAdditionalFIToFI) Validate() error {
	if err := fifi.fieldInclusion(); err != nil {
		return err
	}
	if err:= fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineOne); err!= nil {
		return fieldError("LineOne", err, fifi.AdditionalFIToFI.LineOne)
	}
	if err:= fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineTwo); err!= nil {
		return fieldError("LineTwo", err, fifi.AdditionalFIToFI.LineTwo)
	}
	if err:= fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineThree); err!= nil {
		return fieldError("LineThree", err, fifi.AdditionalFIToFI.LineThree)
	}
	if err:= fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineFour); err!= nil {
		return fieldError("LineFour", err, fifi.AdditionalFIToFI.LineFour)
	}
	if err:= fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineFive); err!= nil {
		return fieldError("LineFive", err, fifi.AdditionalFIToFI.LineFive)
	}
	if err:= fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineSix); err!= nil {
		return fieldError("LineSix", err, fifi.AdditionalFIToFI.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (fifi *FIAdditionalFIToFI) fieldInclusion() error {
	return nil
}

// ToDo:  Add functions for writing

// LineOneField gets a string of the LineOne field
func (fifi *FIAdditionalFIToFI) LineOneField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (fifi *FIAdditionalFIToFI) LineTwoField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (fifi *FIAdditionalFIToFI) LineThreeField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (fifi *FIAdditionalFIToFI) LineFourField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (fifi *FIAdditionalFIToFI) LineFiveField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (fifi *FIAdditionalFIToFI) LineSixField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineSix, 35)
}