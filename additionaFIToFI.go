// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// AdditionalFIToFI is additional financial institution to financial institution information
type AdditionalFIToFI struct {
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

// Parse takes the input string and parses the AdditionalFIToFI values
func (a *AdditionalFIToFI) Parse(record string) (length int, err error) {

	read := 0

	if a.LineOne, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineOne", err)
	}
	length += read

	if a.LineTwo, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineTwo", err)
	}
	length += read

	if a.LineThree, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineThree", err)
	}
	length += read

	if a.LineFour, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineFour", err)
	}
	length += read

	if a.LineFive, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineFive", err)
	}
	length += read

	if a.LineSix, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineSix", err)
	}
	length += read

	return
}

// String writes AdditionalFIToFI
func (a *AdditionalFIToFI) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(210)

	buf.WriteString(a.LineOneField(isVariable))
	buf.WriteString(a.LineTwoField(isVariable))
	buf.WriteString(a.LineThreeField(isVariable))
	buf.WriteString(a.LineFourField(isVariable))
	buf.WriteString(a.LineFiveField(isVariable))
	buf.WriteString(a.LineSixField(isVariable))

	return buf.String()
}

// LineOneField gets a string of the LineOne field
func (a *AdditionalFIToFI) LineOneField(isVariable bool) string {
	return a.alphaVariableField(a.LineOne, 35, isVariable)
}

// LineTwoField gets a string of the LineTwo field
func (a *AdditionalFIToFI) LineTwoField(isVariable bool) string {
	return a.alphaVariableField(a.LineTwo, 35, isVariable)
}

// LineThreeField gets a string of the LineThree field
func (a *AdditionalFIToFI) LineThreeField(isVariable bool) string {
	return a.alphaVariableField(a.LineThree, 35, isVariable)
}

// LineFourField gets a string of the LineFour field
func (a *AdditionalFIToFI) LineFourField(isVariable bool) string {
	return a.alphaVariableField(a.LineFour, 35, isVariable)
}

// LineFiveField gets a string of the LineFive field
func (a *AdditionalFIToFI) LineFiveField(isVariable bool) string {
	return a.alphaVariableField(a.LineFive, 35, isVariable)
}

// LineSixField gets a string of the LineSix field
func (a *AdditionalFIToFI) LineSixField(isVariable bool) string {
	return a.alphaVariableField(a.LineSix, 35, isVariable)
}
