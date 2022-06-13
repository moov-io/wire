// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// OriginatorToBeneficiary is the OriginatorToBeneficiary of the wire
type OriginatorToBeneficiary struct {
	// tag
	tag string
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`
	// LineFour
	LineFour string `json:"lineFour,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginatorToBeneficiary returns a new OriginatorToBeneficiary
func NewOriginatorToBeneficiary() *OriginatorToBeneficiary {
	ob := &OriginatorToBeneficiary{
		tag: TagOriginatorToBeneficiary,
	}
	return ob
}

// Parse takes the input string and parses the OriginatorToBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ob *OriginatorToBeneficiary) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	ob.tag = record[:6]
	length := 6

	value, read, err := ob.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineOne", err)
	}
	ob.LineOne = value
	length += read

	value, read, err = ob.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	ob.LineTwo = value
	length += read

	value, read, err = ob.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineThree", err)
	}
	ob.LineThree = value
	length += read

	value, read, err = ob.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineFour", err)
	}
	ob.LineFour = value
	length += read

	if !ob.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (ob *OriginatorToBeneficiary) UnmarshalJSON(data []byte) error {
	type Alias OriginatorToBeneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(ob),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ob.tag = TagOriginatorToBeneficiary
	return nil
}

// String writes OriginatorToBeneficiary
func (ob *OriginatorToBeneficiary) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(146)

	buf.WriteString(ob.tag)
	buf.WriteString(ob.LineOneField(options...))
	buf.WriteString(ob.LineTwoField(options...))
	buf.WriteString(ob.LineThreeField(options...))
	buf.WriteString(ob.LineFourField(options...))

	if ob.parseFirstOption(options) {
		return ob.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on OriginatorToBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// See latest version of the FAIM manual for Line Limits for Tags {6000} to {6500}.
func (ob *OriginatorToBeneficiary) Validate() error {
	if ob.tag != TagOriginatorToBeneficiary {
		return fieldError("tag", ErrValidTagForType, ob.tag)
	}
	if err := ob.isAlphanumeric(ob.LineOne); err != nil {
		return fieldError("LineOne", err, ob.LineOne)
	}
	if err := ob.isAlphanumeric(ob.LineTwo); err != nil {
		return fieldError("LineTwo", err, ob.LineTwo)
	}
	if err := ob.isAlphanumeric(ob.LineThree); err != nil {
		return fieldError("LineThree", err, ob.LineThree)
	}
	if err := ob.isAlphanumeric(ob.LineFour); err != nil {
		return fieldError("LineFour", err, ob.LineFour)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (ob *OriginatorToBeneficiary) LineOneField(options ...bool) string {
	return ob.alphaVariableField(ob.LineOne, 35, ob.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (ob *OriginatorToBeneficiary) LineTwoField(options ...bool) string {
	return ob.alphaVariableField(ob.LineTwo, 35, ob.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (ob *OriginatorToBeneficiary) LineThreeField(options ...bool) string {
	return ob.alphaVariableField(ob.LineThree, 35, ob.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (ob *OriginatorToBeneficiary) LineFourField(options ...bool) string {
	return ob.alphaVariableField(ob.LineFour, 35, ob.parseFirstOption(options))
}
