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
func NewRemittanceFreeText() *RemittanceFreeText {
	rft := &RemittanceFreeText{
		tag: TagRemittanceFreeText,
	}
	return rft
}

// Parse takes the input string and parses the RemittanceFreeText values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rft *RemittanceFreeText) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 9 {
		return 0, NewTagWrongLengthErr(9, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if rft.tag, read, err = rft.parseTag(record); err != nil {
		return 0, fieldError("RemittanceFreeText.Tag", err)
	}
	length += read

	if rft.LineOne, read, err = rft.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineOne", err)
	}
	length += read

	if rft.LineTwo, read, err = rft.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineTwo", err)
	}
	length += read

	if rft.LineThree, read, err = rft.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineThree", err)
	}
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
func (rft *RemittanceFreeText) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder

	buf.Grow(426)
	buf.WriteString(rft.tag)
	buf.WriteString(rft.LineOneField(isCompressed))
	buf.WriteString(rft.LineTwoField(isCompressed))
	buf.WriteString(rft.LineThreeField(isCompressed))

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
func (rft *RemittanceFreeText) LineOneField(isCompressed bool) string {
	return rft.alphaVariableField(rft.LineOne, 35, isCompressed)
}

// LineTwoField gets a string of the LineTwo field
func (rft *RemittanceFreeText) LineTwoField(isCompressed bool) string {
	return rft.alphaVariableField(rft.LineTwo, 35, isCompressed)
}

// LineThreeField gets a string of the LineThree field
func (rft *RemittanceFreeText) LineThreeField(isCompressed bool) string {
	return rft.alphaVariableField(rft.LineThree, 35, isCompressed)
}
