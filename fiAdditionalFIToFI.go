// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIAdditionalFIToFI{}

// FIAdditionalFIToFI is the financial institution beneficiary financial institution
type FIAdditionalFIToFI struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// AdditionalFiToFi is additional financial institution to financial institution information
	AdditionalFIToFI AdditionalFIToFI `json:"additionalFiToFi,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIAdditionalFIToFI returns a new FIAdditionalFIToFI
func NewFIAdditionalFIToFI(isVariable bool) *FIAdditionalFIToFI {
	fifi := &FIAdditionalFIToFI{
		tag:              TagFIAdditionalFIToFI,
		isVariableLength: isVariable,
	}
	return fifi
}

// Parse takes the input string and parses the FIAdditionalFIToFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fifi *FIAdditionalFIToFI) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 12 {
		return NewTagWrongLengthErr(12, len(record)), 0
	}
	fifi.tag = record[:6]

	return nil, 6 + fifi.AdditionalFIToFI.Parse(record[6:])
}

func (fifi *FIAdditionalFIToFI) UnmarshalJSON(data []byte) error {
	type Alias FIAdditionalFIToFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fifi.tag = TagFIAdditionalFIToFI
	return nil
}

// String writes FIAdditionalFIToFI
func (fifi *FIAdditionalFIToFI) String() string {
	var buf strings.Builder
	buf.Grow(216)

	buf.WriteString(fifi.tag)
	buf.WriteString(fifi.AdditionalFIToFI.String(fifi.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on FIAdditionalFIToFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fifi *FIAdditionalFIToFI) Validate() error {
	if fifi.tag != TagFIAdditionalFIToFI {
		return fieldError("tag", ErrValidTagForType, fifi.tag)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fifi.AdditionalFIToFI.LineOne)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fifi.AdditionalFIToFI.LineTwo)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fifi.AdditionalFIToFI.LineThree)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fifi.AdditionalFIToFI.LineFour)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fifi.AdditionalFIToFI.LineFive)
	}
	if err := fifi.isAlphanumeric(fifi.AdditionalFIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fifi.AdditionalFIToFI.LineSix)
	}
	return nil
}
