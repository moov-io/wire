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
	length := 6

	value, read, err := ii.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	ii.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = ii.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	ii.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = ii.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	ii.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = ii.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	ii.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = ii.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	ii.CoverPayment.SwiftLineFour = value
	length += read

	value, read, err = ii.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFive", err)
	}
	ii.CoverPayment.SwiftLineFive = value
	length += read

	if err := ii.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width IntermediaryInstitution record
func (ii *IntermediaryInstitution) String() string {
	return ii.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a IntermediaryInstitution record formatted according to the FormatOptions
func (ii *IntermediaryInstitution) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(ii.tag)
	buf.WriteString(ii.FormatSwiftFieldTag(options) + Delimiter)
	buf.WriteString(ii.FormatSwiftLineOne(options) + Delimiter)
	buf.WriteString(ii.FormatSwiftLineTwo(options) + Delimiter)
	buf.WriteString(ii.FormatSwiftLineThree(options) + Delimiter)
	buf.WriteString(ii.FormatSwiftLineFour(options) + Delimiter)
	buf.WriteString(ii.FormatSwiftLineFive(options) + Delimiter)

	if options.VariableLengthFields {
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
func (ii *IntermediaryInstitution) SwiftFieldTagField() string {
	return ii.alphaField(ii.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ii *IntermediaryInstitution) SwiftLineOneField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ii *IntermediaryInstitution) SwiftLineTwoField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ii *IntermediaryInstitution) SwiftLineThreeField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ii *IntermediaryInstitution) SwiftLineFourField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (ii *IntermediaryInstitution) SwiftLineFiveField() string {
	return ii.alphaField(ii.CoverPayment.SwiftLineFive, 35)
}

// FormatSwiftFieldTag returns CoverPayment.SwiftFieldTag formatted according to the FormatOptions
func (ii *IntermediaryInstitution) FormatSwiftFieldTag(options FormatOptions) string {
	return ii.formatAlphaField(ii.CoverPayment.SwiftFieldTag, 5, options)
}

// FormatSwiftLineOne returns CoverPayment.SwiftLineOne formatted according to the FormatOptions
func (ii *IntermediaryInstitution) FormatSwiftLineOne(options FormatOptions) string {
	return ii.formatAlphaField(ii.CoverPayment.SwiftLineOne, 35, options)
}

// FormatSwiftLineTwo returns CoverPayment.SwiftLineTwo formatted according to the FormatOptions
func (ii *IntermediaryInstitution) FormatSwiftLineTwo(options FormatOptions) string {
	return ii.formatAlphaField(ii.CoverPayment.SwiftLineTwo, 35, options)
}

// FormatSwiftLineThree returns CoverPayment.SwiftLineThree formatted according to the FormatOptions
func (ii *IntermediaryInstitution) FormatSwiftLineThree(options FormatOptions) string {
	return ii.formatAlphaField(ii.CoverPayment.SwiftLineThree, 35, options)
}

// FormatSwiftLineFour returns CoverPayment.SwiftLineFour formatted according to the FormatOptions
func (ii *IntermediaryInstitution) FormatSwiftLineFour(options FormatOptions) string {
	return ii.formatAlphaField(ii.CoverPayment.SwiftLineFour, 35, options)
}

// FormatSwiftLineFive returns CoverPayment.SwiftLineFive formatted according to the FormatOptions
func (ii *IntermediaryInstitution) FormatSwiftLineFive(options FormatOptions) string {
	return ii.formatAlphaField(ii.CoverPayment.SwiftLineFive, 35, options)
}
