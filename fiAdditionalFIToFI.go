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

	length := 6
	read := 0

	fifi.AdditionalFIToFI.LineOne, read = fifi.parseVariableStringField(record[length:], 35)
	length += read

	fifi.AdditionalFIToFI.LineTwo, read = fifi.parseVariableStringField(record[length:], 35)
	length += read

	fifi.AdditionalFIToFI.LineThree, read = fifi.parseVariableStringField(record[length:], 35)
	length += read

	fifi.AdditionalFIToFI.LineFour, read = fifi.parseVariableStringField(record[length:], 35)
	length += read

	fifi.AdditionalFIToFI.LineFive, read = fifi.parseVariableStringField(record[length:], 35)
	length += read

	fifi.AdditionalFIToFI.LineSix, read = fifi.parseVariableStringField(record[length:], 35)
	length += read

	return nil, length
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
	buf.WriteString(fifi.LineOneField())
	buf.WriteString(fifi.LineTwoField())
	buf.WriteString(fifi.LineThreeField())
	buf.WriteString(fifi.LineFourField())
	buf.WriteString(fifi.LineFiveField())
	buf.WriteString(fifi.LineSixField())
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

// LineOneField gets a string of the LineOne field
func (fifi *FIAdditionalFIToFI) LineOneField() string {
	return fifi.alphaVariableField(fifi.AdditionalFIToFI.LineOne, 35, fifi.isVariableLength)
}

// LineTwoField gets a string of the LineTwo field
func (fifi *FIAdditionalFIToFI) LineTwoField() string {
	return fifi.alphaVariableField(fifi.AdditionalFIToFI.LineTwo, 35, fifi.isVariableLength)
}

// LineThreeField gets a string of the LineThree field
func (fifi *FIAdditionalFIToFI) LineThreeField() string {
	return fifi.alphaVariableField(fifi.AdditionalFIToFI.LineThree, 35, fifi.isVariableLength)
}

// LineFourField gets a string of the LineFour field
func (fifi *FIAdditionalFIToFI) LineFourField() string {
	return fifi.alphaVariableField(fifi.AdditionalFIToFI.LineFour, 35, fifi.isVariableLength)
}

// LineFiveField gets a string of the LineFive field
func (fifi *FIAdditionalFIToFI) LineFiveField() string {
	return fifi.alphaVariableField(fifi.AdditionalFIToFI.LineFive, 35, fifi.isVariableLength)
}

// LineSixField gets a string of the LineSix field
func (fifi *FIAdditionalFIToFI) LineSixField() string {
	return fifi.alphaVariableField(fifi.AdditionalFIToFI.LineSix, 35, fifi.isVariableLength)
}
