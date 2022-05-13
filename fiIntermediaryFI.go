// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIIntermediaryFI{}

// FIIntermediaryFI is the financial institution intermediary financial institution
type FIIntermediaryFI struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFI returns a new FIIntermediaryFI
func NewFIIntermediaryFI(isVariable bool) *FIIntermediaryFI {
	fiifi := &FIIntermediaryFI{
		tag:              TagFIIntermediaryFI,
		isVariableLength: isVariable,
	}
	return fiifi
}

// Parse takes the input string and parses the FIIntermediaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifi *FIIntermediaryFI) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 12 {
		return NewTagWrongLengthErr(12, len(record)), 0
	}
	fiifi.tag = record[:6]

	return nil, 6 + fiifi.FIToFI.Parse(record[6:])
}

func (fiifi *FIIntermediaryFI) UnmarshalJSON(data []byte) error {
	type Alias FIIntermediaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fiifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fiifi.tag = TagFIIntermediaryFI
	return nil
}

// String writes FIIntermediaryFI
func (fiifi *FIIntermediaryFI) String() string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fiifi.tag)
	buf.WriteString(fiifi.FIToFI.String(fiifi.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiifi *FIIntermediaryFI) Validate() error {
	if fiifi.tag != TagFIIntermediaryFI {
		return fieldError("tag", ErrValidTagForType, fiifi.tag)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fiifi.FIToFI.LineOne)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiifi.FIToFI.LineTwo)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fiifi.FIToFI.LineThree)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fiifi.FIToFI.LineFour)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fiifi.FIToFI.LineFive)
	}
	if err := fiifi.isAlphanumeric(fiifi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fiifi.FIToFI.LineSix)
	}
	return nil
}
