// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &IntermediaryInstitution{}

// IntermediaryInstitution is the intermediary institution
type IntermediaryInstitution struct {
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

// NewIntermediaryInstitution returns a new IntermediaryInstitution
func NewIntermediaryInstitution(isVariable bool) *IntermediaryInstitution {
	ii := &IntermediaryInstitution{
		tag:              TagIntermediaryInstitution,
		isVariableLength: isVariable,
	}
	return ii
}

// Parse takes the input string and parses the IntermediaryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ii *IntermediaryInstitution) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if ii.tag, read, err = ii.parseTag(record); err != nil {
		return 0, fieldError("IntermediaryInstitution.Tag", err)
	}
	length += read

	if read, err = ii.CoverPayment.ParseFive(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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
func (ii *IntermediaryInstitution) String() string {
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(ii.tag)
	buf.WriteString(ii.CoverPayment.StringFive(ii.isVariableLength))

	return buf.String()
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
