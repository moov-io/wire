// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &OrderingCustomer{}

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
func (oc *OrderingCustomer) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if oc.tag, read, err = oc.parseTag(record); err != nil {
		return 0, fieldError("OrderingCustomer.Tag", err)
	}
	length += read

	if read, err = oc.CoverPayment.ParseFive(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(oc.tag)
	buf.WriteString(oc.CoverPayment.StringFive(isCompressed))

	return buf.String()
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
