// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// FIBeneficiary is the financial institution beneficiary
type FIBeneficiary struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiary returns a new FIBeneficiary
func NewFIBeneficiary() *FIBeneficiary {
	fib := &FIBeneficiary{
		tag: TagFIBeneficiary,
	}
	return fib
}

// Parse takes the input string and parses the FIBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fib *FIBeneficiary) Parse(record string) error {
	if utf8.RuneCountInString(record) != 201 {
		return NewTagWrongLengthErr(201, len(record))
	}
	fib.tag = record[:6]
	fib.FIToFI.LineOne = fib.parseStringField(record[6:36])
	fib.FIToFI.LineTwo = fib.parseStringField(record[36:69])
	fib.FIToFI.LineThree = fib.parseStringField(record[69:102])
	fib.FIToFI.LineFour = fib.parseStringField(record[102:135])
	fib.FIToFI.LineFive = fib.parseStringField(record[135:168])
	fib.FIToFI.LineSix = fib.parseStringField(record[168:201])
	return nil
}

// String writes FIBeneficiary
func (fib *FIBeneficiary) String() string {
	var buf strings.Builder
	buf.Grow(201)
	buf.WriteString(fib.tag)
	buf.WriteString(fib.LineOneField())
	buf.WriteString(fib.LineTwoField())
	buf.WriteString(fib.LineThreeField())
	buf.WriteString(fib.LineFourField())
	buf.WriteString(fib.LineFiveField())
	buf.WriteString(fib.LineSixField())
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fib *FIBeneficiary) Validate() error {
	if err := fib.isAlphanumeric(fib.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fib.FIToFI.LineOne)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fib.FIToFI.LineTwo)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fib.FIToFI.LineThree)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fib.FIToFI.LineFour)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fib.FIToFI.LineFive)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fib.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fib *FIBeneficiary) LineOneField() string {
	return fib.alphaField(fib.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (fib *FIBeneficiary) LineTwoField() string {
	return fib.alphaField(fib.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fib *FIBeneficiary) LineThreeField() string {
	return fib.alphaField(fib.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fib *FIBeneficiary) LineFourField() string {
	return fib.alphaField(fib.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fib *FIBeneficiary) LineFiveField() string {
	return fib.alphaField(fib.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fib *FIBeneficiary) LineSixField() string {
	return fib.alphaField(fib.FIToFI.LineSix, 33)
}
