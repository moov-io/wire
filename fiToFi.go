// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (f *FIToFI) Parse(record string) int {

	length := 0
	read := 0

	f.LineOne, read = f.parseVariableStringField(record[length:], 30)
	length += read

	f.LineTwo, read = f.parseVariableStringField(record[length:], 33)
	length += read

	f.LineThree, read = f.parseVariableStringField(record[length:], 33)
	length += read

	f.LineFour, read = f.parseVariableStringField(record[length:], 33)
	length += read

	f.LineFive, read = f.parseVariableStringField(record[length:], 33)
	length += read

	f.LineSix, read = f.parseVariableStringField(record[length:], 33)
	length += read

	return length
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
