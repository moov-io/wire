// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (rft *RemittanceFreeText) Parse(record string) {
	rft.tag = record[:6]
	rft.LineOne = rft.parseStringField(record[6:146])
	rft.LineTwo = rft.parseStringField(record[146:286])
	rft.LineThree = rft.parseStringField(record[286:426])
}

// String writes RemittanceFreeText
func (rft *RemittanceFreeText) String() string {
	var buf strings.Builder
	// ToDo: Separator
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
	if err := rft.fieldInclusion(); err != nil {
		return err
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

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rft *RemittanceFreeText) fieldInclusion() error {
	return nil
}

// LineOneField gets a string of the LineOne field
func (rft *RemittanceFreeText) LineOneField() string {
	return rft.alphaField(rft.LineOne, 140)
}

// LineTwoField gets a string of the LineTwo field
func (rft *RemittanceFreeText) LineTwoField() string {
	return rft.alphaField(rft.LineTwo, 140)
}

// LineThreeField gets a string of the LineThree field
func (rft *RemittanceFreeText) LineThreeField() string {
	return rft.alphaField(rft.LineThree, 140)
}
