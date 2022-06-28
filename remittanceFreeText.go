// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (rft *RemittanceFreeText) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	rft.tag = record[:6]
	length := 6

	value, read, err := rft.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("LineOne", err)
	}
	rft.LineOne = value
	length += read

	value, read, err = rft.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	rft.LineTwo = value
	length += read

	value, read, err = rft.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("LineThree", err)
	}
	rft.LineThree = value
	length += read

	if !rft.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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

// String returns a fixed-width RemittanceFreeText record
func (rft *RemittanceFreeText) String() string {
	return rft.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a RemittanceFreeText record formatted according to the FormatOptions
func (rft *RemittanceFreeText) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(426)

	buf.WriteString(rft.tag)
	buf.WriteString(rft.FormatLineOne(options))
	buf.WriteString(rft.FormatLineTwo(options))
	buf.WriteString(rft.FormatLineThree(options))

	if options.VariableLengthFields {
		return rft.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// FormatLineOne returns LineOne formatted according to the FormatOptions
func (rft *RemittanceFreeText) FormatLineOne(options FormatOptions) string {
	return rft.formatAlphaField(rft.LineOne, 140, options)
}

// FormatLineTwo returns LineTwo formatted according to the FormatOptions
func (rft *RemittanceFreeText) FormatLineTwo(options FormatOptions) string {
	return rft.formatAlphaField(rft.LineTwo, 140, options)
}

// FormatLineThree returns LineThree formatted according to the FormatOptions
func (rft *RemittanceFreeText) FormatLineThree(options FormatOptions) string {
	return rft.formatAlphaField(rft.LineThree, 140, options)
}
