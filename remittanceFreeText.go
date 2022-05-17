// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &RemittanceFreeText{}

// RemittanceFreeText is the remittance free text
type RemittanceFreeText struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceFreeText returns a new RemittanceFreeText
func NewRemittanceFreeText(isVariable bool) *RemittanceFreeText {
	rft := &RemittanceFreeText{
		tag:              TagRemittanceFreeText,
		isVariableLength: isVariable,
	}
	return rft
}

// Parse takes the input string and parses the RemittanceFreeText values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rft *RemittanceFreeText) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) != 426 {
		return 0, NewTagWrongLengthErr(426, utf8.RuneCountInString(record))
	}

	rft.tag = record[:6]

	length := 6
	read := 0

	rft.LineOne, read = rft.parseVariableStringField(record[length:], 140)
	length += read

	rft.LineTwo, read = rft.parseVariableStringField(record[length:], 140)
	length += read

	rft.LineThree, read = rft.parseVariableStringField(record[length:], 140)
	length += read

	return length, nil
}

func (rft *RemittanceFreeText) UnmarshalJSON(data []byte) error {
	type Alias RemittanceFreeText
	aux := struct {
		*Alias
	}{
		(*Alias)(rft),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rft.tag = TagRemittanceFreeText
	return nil
}

// String writes RemittanceFreeText
func (rft *RemittanceFreeText) String() string {
	var buf strings.Builder

	buf.Grow(426)
	buf.WriteString(rft.tag)
	buf.WriteString(rft.LineOneField())
	buf.WriteString(rft.LineTwoField())
	buf.WriteString(rft.LineThreeField())

	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceFreeText and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rft *RemittanceFreeText) Validate() error {
	if rft.tag != TagRemittanceFreeText {
		return fieldError("tag", ErrValidTagForType, rft.tag)
	}
	if err := rft.isAlphanumeric(rft.LineOne); err != nil {
		return fieldError("LineOne", err, rft.LineOne)
	}
	if err := rft.isAlphanumeric(rft.LineTwo); err != nil {
		return fieldError("LineTwo", err, rft.LineTwo)
	}
	if err := rft.isAlphanumeric(rft.LineThree); err != nil {
		return fieldError("LineThree", err, rft.LineThree)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (rft *RemittanceFreeText) LineOneField() string {
	return rft.alphaVariableField(rft.LineOne, 35, rft.isVariableLength)
}

// LineTwoField gets a string of the LineTwo field
func (rft *RemittanceFreeText) LineTwoField() string {
	return rft.alphaVariableField(rft.LineTwo, 35, rft.isVariableLength)
}

// LineThreeField gets a string of the LineThree field
func (rft *RemittanceFreeText) LineThreeField() string {
	return rft.alphaVariableField(rft.LineThree, 35, rft.isVariableLength)
}
