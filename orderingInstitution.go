// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &OrderingInstitution{}

// OrderingInstitution is the ordering institution
type OrderingInstitution struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingInstitution returns a new OrderingInstitution
func NewOrderingInstitution(isVariable bool) *OrderingInstitution {
	oi := &OrderingInstitution{
		tag:              TagOrderingInstitution,
		isVariableLength: isVariable,
	}
	return oi
}

// Parse takes the input string and parses the OrderingInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oi *OrderingInstitution) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if oi.tag, read, err = oi.parseTag(record); err != nil {
		return 0, fieldError("OrderingInstitution.Tag", err)
	}
	length += read

	if read, err = oi.CoverPayment.ParseFive(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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
func (oi *OrderingInstitution) String() string {
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(oi.tag)
	buf.WriteString(oi.CoverPayment.StringFive(oi.isVariableLength))

	return buf.String()
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
