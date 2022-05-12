// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// Advice is financial institution advice information
type Advice struct {
	AdviceCode string `json:"adviceCode,omitempty"`
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

// Parse takes the input string and parses the Advice values
func (a *Advice) Parse(record string) int {

	length := 0
	read := 0

	a.AdviceCode, read = a.parseVariableStringField(record[length:], 3)
	length += read

	a.LineOne, read = a.parseVariableStringField(record[length:], 26)
	length += read

	a.LineTwo, read = a.parseVariableStringField(record[length:], 33)
	length += read

	a.LineThree, read = a.parseVariableStringField(record[length:], 33)
	length += read

	a.LineFour, read = a.parseVariableStringField(record[length:], 33)
	length += read

	a.LineFive, read = a.parseVariableStringField(record[length:], 33)
	length += read

	a.LineSix, read = a.parseVariableStringField(record[length:], 33)
	length += read

	return length
}

// String writes Advice
func (a *Advice) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(194)

	buf.WriteString(a.AdviceCodeField(isVariable))
	buf.WriteString(a.LineOneField(isVariable))
	buf.WriteString(a.LineTwoField(isVariable))
	buf.WriteString(a.LineThreeField(isVariable))
	buf.WriteString(a.LineFourField(isVariable))
	buf.WriteString(a.LineFiveField(isVariable))
	buf.WriteString(a.LineSixField(isVariable))

	return buf.String()
}

// AdviceCodeField gets a string of the AdviceCode field
func (a *Advice) AdviceCodeField(isVariable bool) string {
	return a.alphaVariableField(a.AdviceCode, 3, isVariable)
}

// LineOneField gets a string of the LineOne field
func (a *Advice) LineOneField(isVariable bool) string {
	return a.alphaVariableField(a.LineOne, 26, isVariable)
}

// LineTwoField gets a string of the LineTwo field
func (a *Advice) LineTwoField(isVariable bool) string {
	return a.alphaVariableField(a.LineTwo, 33, isVariable)
}

// LineThreeField gets a string of the LineThree field
func (a *Advice) LineThreeField(isVariable bool) string {
	return a.alphaVariableField(a.LineThree, 33, isVariable)
}

// LineFourField gets a string of the LineFour field
func (a *Advice) LineFourField(isVariable bool) string {
	return a.alphaVariableField(a.LineFour, 33, isVariable)
}

// LineFiveField gets a string of the LineFive field
func (a *Advice) LineFiveField(isVariable bool) string {
	return a.alphaVariableField(a.LineFive, 33, isVariable)
}

// LineSixField gets a string of the LineSix field
func (a *Advice) LineSixField(isVariable bool) string {
	return a.alphaVariableField(a.LineSix, 33, isVariable)
}
