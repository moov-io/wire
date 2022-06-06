// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// IntermediaryInstitution is the intermediary institution
type IntermediaryInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewIntermediaryInstitution returns a new IntermediaryInstitution
func NewIntermediaryInstitution() *IntermediaryInstitution {
	ii := &IntermediaryInstitution{
		tag: TagIntermediaryInstitution,
	}
	return ii
}

// Parse takes the input string and parses the IntermediaryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ii *IntermediaryInstitution) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	ii.tag = record[:6]

	var err error
	length := 6
	read := 0

	if ii.CoverPayment.SwiftFieldTag, read, err = ii.parseVariableStringField(record[length:], 5); err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	length += read

	if ii.CoverPayment.SwiftLineOne, read, err = ii.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("SwiftLineOne", err)
	}
	length += read

	if ii.CoverPayment.SwiftLineTwo, read, err = ii.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	length += read

	if ii.CoverPayment.SwiftLineThree, read, err = ii.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("SwiftLineThree", err)
	}
	length += read

	if ii.CoverPayment.SwiftLineFour, read, err = ii.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("SwiftLineFour", err)
	}
	length += read

	if ii.CoverPayment.SwiftLineFive, read, err = ii.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("SwiftLineFive", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (ii *IntermediaryInstitution) UnmarshalJSON(data []byte) error {
	type Alias IntermediaryInstitution
	aux := struct {
		*Alias
	}{
		(*Alias)(ii),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ii.tag = TagIntermediaryInstitution
	return nil
}

// String writes IntermediaryInstitution
func (ii *IntermediaryInstitution) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(ii.tag)
	buf.WriteString(ii.SwiftFieldTagField(options...))
	buf.WriteString(ii.SwiftLineOneField(options...))
	buf.WriteString(ii.SwiftLineTwoField(options...))
	buf.WriteString(ii.SwiftLineThreeField(options...))
	buf.WriteString(ii.SwiftLineFourField(options...))
	buf.WriteString(ii.SwiftLineFiveField(options...))

	if ii.parseFirstOption(options) {
		return ii.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on IntermediaryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ii *IntermediaryInstitution) Validate() error {
	if err := ii.fieldInclusion(); err != nil {
		return err
	}
	if ii.tag != TagIntermediaryInstitution {
		return fieldError("tag", ErrValidTagForType, ii.tag)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ii.CoverPayment.SwiftFieldTag)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ii.CoverPayment.SwiftLineOne)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ii.CoverPayment.SwiftLineTwo)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ii.CoverPayment.SwiftLineThree)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ii.CoverPayment.SwiftLineFour)
	}
	if err := ii.isAlphanumeric(ii.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, ii.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ii *IntermediaryInstitution) fieldInclusion() error {
	if ii.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, ii.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (ii *IntermediaryInstitution) SwiftFieldTagField(options ...bool) string {
	return ii.alphaVariableField(ii.CoverPayment.SwiftFieldTag, 5, ii.parseFirstOption(options))
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ii *IntermediaryInstitution) SwiftLineOneField(options ...bool) string {
	return ii.alphaVariableField(ii.CoverPayment.SwiftLineOne, 35, ii.parseFirstOption(options))
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ii *IntermediaryInstitution) SwiftLineTwoField(options ...bool) string {
	return ii.alphaVariableField(ii.CoverPayment.SwiftLineTwo, 35, ii.parseFirstOption(options))
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ii *IntermediaryInstitution) SwiftLineThreeField(options ...bool) string {
	return ii.alphaVariableField(ii.CoverPayment.SwiftLineThree, 35, ii.parseFirstOption(options))
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ii *IntermediaryInstitution) SwiftLineFourField(options ...bool) string {
	return ii.alphaVariableField(ii.CoverPayment.SwiftLineFour, 35, ii.parseFirstOption(options))
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (ii *IntermediaryInstitution) SwiftLineFiveField(options ...bool) string {
	return ii.alphaVariableField(ii.CoverPayment.SwiftLineFive, 35, ii.parseFirstOption(options))
}
