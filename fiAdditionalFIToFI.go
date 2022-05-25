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
	// AdditionalFiToFi is additional financial institution to financial institution information
	AdditionalFIToFI AdditionalFIToFI `json:"additionalFiToFi,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIAdditionalFIToFI returns a new FIAdditionalFIToFI
func NewFIAdditionalFIToFI() *FIAdditionalFIToFI {
	fifi := &FIAdditionalFIToFI{
		tag: TagFIAdditionalFIToFI,
	}
	return fifi
}

// Parse takes the input string and parses the FIAdditionalFIToFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fifi *FIAdditionalFIToFI) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 6 {
		return 0, NewTagWrongLengthErr(6, len(record))
	}

	var err error
	var length, read int

	if fifi.tag, read, err = fifi.parseTag(record); err != nil {
		return 0, fieldError("FIAdditionalFIToFI.Tag", err)
	}
	length += read

	if read, err = fifi.AdditionalFIToFI.Parse(record[length:]); err != nil {
		return length, nil
	}
	length += read

	return length, nil
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
func (fifi *FIAdditionalFIToFI) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(216)

	buf.WriteString(fifi.tag)
	buf.WriteString(fifi.AdditionalFIToFI.String(isCompressed))

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
