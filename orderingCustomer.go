// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// OrderingCustomer is the ordering customer
type OrderingCustomer struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingCustomer returns a new OrderingCustomer
func NewOrderingCustomer() *OrderingCustomer {
	oc := &OrderingCustomer{
		tag: TagOrderingCustomer,
	}
	return oc
}

// Parse takes the input string and parses the OrderingCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oc *OrderingCustomer) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	oc.tag = record[:6]
	length := 6

	value, read, err := oc.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	oc.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = oc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	oc.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = oc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	oc.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = oc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	oc.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = oc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	oc.CoverPayment.SwiftLineFour = value
	length += read

	value, read, err = oc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFive", err)
	}
	oc.CoverPayment.SwiftLineFive = value
	length += read

	if !oc.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (oc *OrderingCustomer) UnmarshalJSON(data []byte) error {
	type Alias OrderingCustomer
	aux := struct {
		*Alias
	}{
		(*Alias)(oc),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	oc.tag = TagOrderingCustomer
	return nil
}

// String writes OrderingCustomer
func (oc *OrderingCustomer) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(oc.tag)

	buf.WriteString(oc.SwiftFieldTagField(options...))
	buf.WriteString(oc.SwiftLineOneField(options...))
	buf.WriteString(oc.SwiftLineTwoField(options...))
	buf.WriteString(oc.SwiftLineThreeField(options...))
	buf.WriteString(oc.SwiftLineFourField(options...))
	buf.WriteString(oc.SwiftLineFiveField(options...))

	if oc.parseFirstOption(options) {
		return oc.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on OrderingCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oc *OrderingCustomer) Validate() error {
	if err := oc.fieldInclusion(); err != nil {
		return err
	}
	if oc.tag != TagOrderingCustomer {
		return fieldError("tag", ErrValidTagForType, oc.tag)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, oc.CoverPayment.SwiftFieldTag)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, oc.CoverPayment.SwiftLineOne)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, oc.CoverPayment.SwiftLineTwo)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, oc.CoverPayment.SwiftLineThree)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, oc.CoverPayment.SwiftLineFour)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, oc.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oc *OrderingCustomer) fieldInclusion() error {
	if oc.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, oc.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (oc *OrderingCustomer) SwiftFieldTagField(options ...bool) string {
	return oc.alphaVariableField(oc.CoverPayment.SwiftFieldTag, 5, oc.parseFirstOption(options))
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (oc *OrderingCustomer) SwiftLineOneField(options ...bool) string {
	return oc.alphaVariableField(oc.CoverPayment.SwiftLineOne, 35, oc.parseFirstOption(options))
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (oc *OrderingCustomer) SwiftLineTwoField(options ...bool) string {
	return oc.alphaVariableField(oc.CoverPayment.SwiftLineTwo, 35, oc.parseFirstOption(options))
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (oc *OrderingCustomer) SwiftLineThreeField(options ...bool) string {
	return oc.alphaVariableField(oc.CoverPayment.SwiftLineThree, 35, oc.parseFirstOption(options))
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (oc *OrderingCustomer) SwiftLineFourField(options ...bool) string {
	return oc.alphaVariableField(oc.CoverPayment.SwiftLineFour, 35, oc.parseFirstOption(options))
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (oc *OrderingCustomer) SwiftLineFiveField(options ...bool) string {
	return oc.alphaVariableField(oc.CoverPayment.SwiftLineFive, 35, oc.parseFirstOption(options))
}
