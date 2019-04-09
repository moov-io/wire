// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIBeneficiaryFI is the financial institution beneficiary financial institution
type FIBeneficiaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFI returns a new FIBeneficiaryFI
func NewFIBeneficiaryFI() FIBeneficiaryFI {
	fibfi := FIBeneficiaryFI{
		tag: TagFIBeneficiaryFI,
	}
	return fibfi
}

// Parse takes the input string and parses the FIBeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfi *FIBeneficiaryFI) Parse(record string) {
	fibfi.FIToFI.LineOne = fibfi.parseStringField(record[6:36])
	fibfi.FIToFI.LineTwo = fibfi.parseStringField(record[36:69])
	fibfi.FIToFI.LineThree = fibfi.parseStringField(record[69:104])
	fibfi.FIToFI.LineFour = fibfi.parseStringField(record[104:139])
	fibfi.FIToFI.LineFive = fibfi.parseStringField(record[139:174])
	fibfi.FIToFI.LineSix = fibfi.parseStringField(record[174:209])
}

// String writes FIBeneficiaryFI
func (fibfi *FIBeneficiaryFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(209)
	buf.WriteString(fibfi.LineOneField())
	buf.WriteString(fibfi.LineTwoField())
	buf.WriteString(fibfi.LineThreeField())
	buf.WriteString(fibfi.LineFourField())
	buf.WriteString(fibfi.LineFiveField())
	buf.WriteString(fibfi.LineSixField())
	buf.WriteString(fibfi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfi *FIBeneficiaryFI) Validate() error {
	if err := fibfi.fieldInclusion(); err != nil {
		return err
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fibfi.FIToFI.LineOne)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfi.FIToFI.LineTwo)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fibfi.FIToFI.LineThree)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fibfi.FIToFI.LineFour)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fibfi.FIToFI.LineFive)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fibfi.FIToFI.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (fibfi *FIBeneficiaryFI) fieldInclusion() error {
	return nil
}

// LineOneField gets a string of the LineOne field
func (fibfi *FIBeneficiaryFI) LineOneField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (fibfi *FIBeneficiaryFI) LineTwoField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (fibfi *FIBeneficiaryFI) LineThreeField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (fibfi *FIBeneficiaryFI) LineFourField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (fibfi *FIBeneficiaryFI) LineFiveField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (fibfi *FIBeneficiaryFI) LineSixField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineSix, 35)
}
