// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// Remittance is the remittance information
type Remittance struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittance returns a new Remittance
func NewRemittance() *Remittance {
	ri := &Remittance{
		tag: TagRemittance,
	}
	return ri
}

// Parse takes the input string and parses the Remittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ri *Remittance) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	ri.tag = record[:6]
	length := 6

	value, read, err := ri.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	ri.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = ri.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	ri.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = ri.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	ri.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = ri.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	ri.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = ri.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	ri.CoverPayment.SwiftLineFour = value
	length += read

	if !ri.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (ri *Remittance) UnmarshalJSON(data []byte) error {
	type Alias Remittance
	aux := struct {
		*Alias
	}{
		(*Alias)(ri),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ri.tag = TagRemittance
	return nil
}

// String returns a fixed-width Remittance record
func (ri *Remittance) String() string {
	return ri.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a Remittance record formatted according to the FormatOptions
func (ri *Remittance) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(151)

	buf.WriteString(ri.tag)
	buf.WriteString(ri.FormatSwiftFieldTag(options))
	buf.WriteString(ri.FormatSwiftLineOne(options))
	buf.WriteString(ri.FormatSwiftLineTwo(options))
	buf.WriteString(ri.FormatSwiftLineThree(options))
	buf.WriteString(ri.FormatSwiftLineFour(options))

	if options.VariableLengthFields {
		return ri.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on Remittance and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ri *Remittance) Validate() error {
	if err := ri.fieldInclusion(); err != nil {
		return err
	}
	if ri.tag != TagRemittance {
		return fieldError("tag", ErrValidTagForType, ri.tag)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, ri.CoverPayment.SwiftFieldTag)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, ri.CoverPayment.SwiftLineOne)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, ri.CoverPayment.SwiftLineTwo)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, ri.CoverPayment.SwiftLineThree)
	}
	if err := ri.isAlphanumeric(ri.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, ri.CoverPayment.SwiftLineFour)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ri *Remittance) fieldInclusion() error {
	if ri.CoverPayment.SwiftLineFive != "" {
		return fieldError("SwiftLineFive", ErrInvalidProperty, ri.CoverPayment.SwiftLineFive)
	}
	if ri.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, ri.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (ri *Remittance) SwiftFieldTagField() string {
	return ri.alphaField(ri.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (ri *Remittance) SwiftLineOneField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (ri *Remittance) SwiftLineTwoField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (ri *Remittance) SwiftLineThreeField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (ri *Remittance) SwiftLineFourField() string {
	return ri.alphaField(ri.CoverPayment.SwiftLineFour, 35)
}

// FormatSwiftFieldTag returns CoverPayment.SwiftFieldTag formatted according to the FormatOptions
func (ri *Remittance) FormatSwiftFieldTag(options FormatOptions) string {
	return ri.formatAlphaField(ri.CoverPayment.SwiftFieldTag, 5, options)
}

// FormatSwiftLineOne returns CoverPayment.SwiftLineOne formatted according to the FormatOptions
func (ri *Remittance) FormatSwiftLineOne(options FormatOptions) string {
	return ri.formatAlphaField(ri.CoverPayment.SwiftLineOne, 35, options)
}

// FormatSwiftLineTwo returns CoverPayment.SwiftLineTwo formatted according to the FormatOptions
func (ri *Remittance) FormatSwiftLineTwo(options FormatOptions) string {
	return ri.formatAlphaField(ri.CoverPayment.SwiftLineTwo, 35, options)
}

// FormatSwiftLineThree returns CoverPayment.SwiftLineThree formatted according to the FormatOptions
func (ri *Remittance) FormatSwiftLineThree(options FormatOptions) string {
	return ri.formatAlphaField(ri.CoverPayment.SwiftLineThree, 35, options)
}

// FormatSwiftLineFour returns CoverPayment.SwiftLineFour formatted according to the FormatOptions
func (ri *Remittance) FormatSwiftLineFour(options FormatOptions) string {
	return ri.formatAlphaField(ri.CoverPayment.SwiftLineFour, 35, options)
}
