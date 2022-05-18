// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// FIToFI is financial institution to financial institution
type FIToFI struct {
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`
	// LineFour
	LineFour string `json:"lineFour,omitempty"`
	// LineFive
	LineFive string `json:"lineFive,omitempty"`
	// LineSix
	LineSix string `json:"lineSix,omitempty"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the FIBeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (f *FIToFI) Parse(record string) (length int, err error) {

	var read int

	if f.LineOne, read, err = f.parseVariableStringField(record[length:], 30); err != nil {
		return 0, fieldError("LineOne", err)
	}
	length += read

	if f.LineTwo, read, err = f.parseVariableStringField(record[length:], 33); err != nil {
		return 0, fieldError("LineTwo", err)
	}
	length += read

	if f.LineThree, read, err = f.parseVariableStringField(record[length:], 33); err != nil {
		return 0, fieldError("LineThree", err)
	}
	length += read

	if f.LineFour, read, err = f.parseVariableStringField(record[length:], 33); err != nil {
		return 0, fieldError("LineFour", err)
	}
	length += read

	if f.LineFive, read, err = f.parseVariableStringField(record[length:], 33); err != nil {
		return 0, fieldError("LineFive", err)
	}
	length += read

	if f.LineSix, read, err = f.parseVariableStringField(record[length:], 33); err != nil {
		return 0, fieldError("LineSix", err)
	}
	length += read

	return
}

// String writes FIToFI
func (f *FIToFI) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(195)

	buf.WriteString(f.LineOneField(isVariable))
	buf.WriteString(f.LineTwoField(isVariable))
	buf.WriteString(f.LineThreeField(isVariable))
	buf.WriteString(f.LineFourField(isVariable))
	buf.WriteString(f.LineFiveField(isVariable))
	buf.WriteString(f.LineSixField(isVariable))

	return buf.String()
}

// LineOneField gets a string of the LineOne field
func (f *FIToFI) LineOneField(isVariable bool) string {
	return f.alphaVariableField(f.LineOne, 30, isVariable)
}

// LineTwoField gets a string of the LineTwo field
func (f *FIToFI) LineTwoField(isVariable bool) string {
	return f.alphaVariableField(f.LineTwo, 33, isVariable)
}

// LineThreeField gets a string of the LineThree field
func (f *FIToFI) LineThreeField(isVariable bool) string {
	return f.alphaVariableField(f.LineThree, 33, isVariable)
}

// LineFourField gets a string of the LineFour field
func (f *FIToFI) LineFourField(isVariable bool) string {
	return f.alphaVariableField(f.LineFour, 33, isVariable)
}

// LineFiveField gets a string of the LineFive field
func (f *FIToFI) LineFiveField(isVariable bool) string {
	return f.alphaVariableField(f.LineFive, 33, isVariable)
}

// LineSixField gets a string of the LineSix field
func (f *FIToFI) LineSixField(isVariable bool) string {
	return f.alphaVariableField(f.LineSix, 33, isVariable)
}
