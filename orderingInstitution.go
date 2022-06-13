// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// OrderingInstitution is the ordering institution
type OrderingInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingInstitution returns a new OrderingInstitution
func NewOrderingInstitution() *OrderingInstitution {
	oi := &OrderingInstitution{
		tag: TagOrderingInstitution,
	}
	return oi
}

// Parse takes the input string and parses the OrderingInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oi *OrderingInstitution) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	oi.tag = record[:6]
	length := 6

	value, read, err := oi.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	oi.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = oi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	oi.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = oi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	oi.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = oi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	oi.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = oi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	oi.CoverPayment.SwiftLineFour = value
	length += read

	value, read, err = oi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFive", err)
	}
	oi.CoverPayment.SwiftLineFive = value
	length += read

	if !oi.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (oi *OrderingInstitution) UnmarshalJSON(data []byte) error {
	type Alias OrderingInstitution
	aux := struct {
		*Alias
	}{
		(*Alias)(oi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	oi.tag = TagOrderingInstitution
	return nil
}

// String writes OrderingInstitution
func (oi *OrderingInstitution) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(oi.tag)

	buf.WriteString(oi.SwiftFieldTagField(options...))
	buf.WriteString(oi.SwiftLineOneField(options...))
	buf.WriteString(oi.SwiftLineTwoField(options...))
	buf.WriteString(oi.SwiftLineThreeField(options...))
	buf.WriteString(oi.SwiftLineFourField(options...))
	buf.WriteString(oi.SwiftLineFiveField(options...))

	if oi.parseFirstOption(options) {
		return oi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on OrderingInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oi *OrderingInstitution) Validate() error {
	if err := oi.fieldInclusion(); err != nil {
		return err
	}
	if oi.tag != TagOrderingInstitution {
		return fieldError("tag", ErrValidTagForType, oi.tag)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, oi.CoverPayment.SwiftFieldTag)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, oi.CoverPayment.SwiftLineOne)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, oi.CoverPayment.SwiftLineTwo)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, oi.CoverPayment.SwiftLineThree)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, oi.CoverPayment.SwiftLineFour)
	}
	if err := oi.isAlphanumeric(oi.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, oi.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oi *OrderingInstitution) fieldInclusion() error {
	if oi.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, oi.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (oi *OrderingInstitution) SwiftFieldTagField(options ...bool) string {
	return oi.alphaVariableField(oi.CoverPayment.SwiftFieldTag, 5, oi.parseFirstOption(options))
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (oi *OrderingInstitution) SwiftLineOneField(options ...bool) string {
	return oi.alphaVariableField(oi.CoverPayment.SwiftLineOne, 35, oi.parseFirstOption(options))
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (oi *OrderingInstitution) SwiftLineTwoField(options ...bool) string {
	return oi.alphaVariableField(oi.CoverPayment.SwiftLineTwo, 35, oi.parseFirstOption(options))
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (oi *OrderingInstitution) SwiftLineThreeField(options ...bool) string {
	return oi.alphaVariableField(oi.CoverPayment.SwiftLineThree, 35, oi.parseFirstOption(options))
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (oi *OrderingInstitution) SwiftLineFourField(options ...bool) string {
	return oi.alphaVariableField(oi.CoverPayment.SwiftLineFour, 35, oi.parseFirstOption(options))
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (oi *OrderingInstitution) SwiftLineFiveField(options ...bool) string {
	return oi.alphaVariableField(oi.CoverPayment.SwiftLineFive, 35, oi.parseFirstOption(options))
}
