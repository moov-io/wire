// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &Remittance{}

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
func (ri *Remittance) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 11 {
		return 0, NewTagWrongLengthErr(11, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if ri.tag, read, err = ri.parseTag(record); err != nil {
		return 0, fieldError("Remittance.Tag", err)
	}
	length += read

	if read, err = ri.CoverPayment.ParseFour(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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

// String writes Remittance
func (ri *Remittance) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(151)

	buf.WriteString(ri.tag)
	buf.WriteString(ri.CoverPayment.StringFour(isCompressed))

	return buf.String()
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
