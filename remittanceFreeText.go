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
func NewRemittanceFreeText() RemittanceFreeText  {
	rft := RemittanceFreeText {
		tag: TagRemittanceFreeText,
	}
	return rft
}

// Parse takes the input string and parses the RemittanceFreeText values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rft *RemittanceFreeText) Parse(record string) {
}

// String writes RemittanceFreeText
func (rft *RemittanceFreeText) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(420)
	buf.WriteString(rft.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceFreeText and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rft *RemittanceFreeText) Validate() error {
	if err := rft.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rft *RemittanceFreeText) fieldInclusion() error {
	return nil
}
