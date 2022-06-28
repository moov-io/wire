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

// String returns a fixed-width OrderingCustomer record
func (oc *OrderingCustomer) String() string {
	return oc.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a OrderingCustomer record formatted according to the FormatOptions
func (oc *OrderingCustomer) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(oc.tag)

	buf.WriteString(oc.FormatSwiftFieldTag(options))
	buf.WriteString(oc.FormatSwiftLineOne(options))
	buf.WriteString(oc.FormatSwiftLineTwo(options))
	buf.WriteString(oc.FormatSwiftLineThree(options))
	buf.WriteString(oc.FormatSwiftLineFour(options))
	buf.WriteString(oc.FormatSwiftLineFive(options))

	if options.VariableLengthFields {
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
func (oc *OrderingCustomer) SwiftFieldTagField() string {
	return oc.alphaField(oc.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (oc *OrderingCustomer) SwiftLineOneField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (oc *OrderingCustomer) SwiftLineTwoField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (oc *OrderingCustomer) SwiftLineThreeField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (oc *OrderingCustomer) SwiftLineFourField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (oc *OrderingCustomer) SwiftLineFiveField() string {
	return oc.alphaField(oc.CoverPayment.SwiftLineFive, 35)
}

// FormatSwiftFieldTag returns CoverPayment.SwiftFieldTag formatted according to the FormatOptions
func (oc *OrderingCustomer) FormatSwiftFieldTag(options FormatOptions) string {
	return oc.formatAlphaField(oc.CoverPayment.SwiftFieldTag, 5, options)
}

// FormatSwiftLineOne returns CoverPayment.SwiftLineOne formatted according to the FormatOptions
func (oc *OrderingCustomer) FormatSwiftLineOne(options FormatOptions) string {
	return oc.formatAlphaField(oc.CoverPayment.SwiftLineOne, 35, options)
}

// FormatSwiftLineTwo returns CoverPayment.SwiftLineTwo formatted according to the FormatOptions
func (oc *OrderingCustomer) FormatSwiftLineTwo(options FormatOptions) string {
	return oc.formatAlphaField(oc.CoverPayment.SwiftLineTwo, 35, options)
}

// FormatSwiftLineThree returns CoverPayment.SwiftLineThree formatted according to the FormatOptions
func (oc *OrderingCustomer) FormatSwiftLineThree(options FormatOptions) string {
	return oc.formatAlphaField(oc.CoverPayment.SwiftLineThree, 35, options)
}

// FormatSwiftLineFour returns CoverPayment.SwiftLineFour formatted according to the FormatOptions
func (oc *OrderingCustomer) FormatSwiftLineFour(options FormatOptions) string {
	return oc.formatAlphaField(oc.CoverPayment.SwiftLineFour, 35, options)
}

// FormatSwiftLineFive returns CoverPayment.SwiftLineFive formatted according to the FormatOptions
func (oc *OrderingCustomer) FormatSwiftLineFive(options FormatOptions) string {
	return oc.formatAlphaField(oc.CoverPayment.SwiftLineFive, 35, options)
}
