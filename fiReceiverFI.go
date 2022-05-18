// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIReceiverFI{}

// FIReceiverFI is the financial institution receiver financial institution
type FIReceiverFI struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// FIToFI is financial institution to financial institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIReceiverFI returns a new FIReceiverFI
func NewFIReceiverFI(isVariable bool) *FIReceiverFI {
	firfi := &FIReceiverFI{
		tag:              TagFIReceiverFI,
		isVariableLength: isVariable,
	}
	return firfi
}

// Parse takes the input string and parses the FIReceiverFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (firfi *FIReceiverFI) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if firfi.tag, read, err = firfi.parseTag(record); err != nil {
		return 0, fieldError("FIReceiverFI.Tag", err)
	}
	length += read

	if read, err = firfi.FIToFI.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (firfi *FIReceiverFI) UnmarshalJSON(data []byte) error {
	type Alias FIReceiverFI
	aux := struct {
		*Alias
	}{
		(*Alias)(firfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	firfi.tag = TagFIReceiverFI
	return nil
}

// String writes FIReceiverFI
func (firfi *FIReceiverFI) String() string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(firfi.tag)
	buf.WriteString(firfi.FIToFI.String(firfi.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on FIReceiverFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (firfi *FIReceiverFI) Validate() error {
	if firfi.tag != TagFIReceiverFI {
		return fieldError("tag", ErrValidTagForType, firfi.tag)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, firfi.FIToFI.LineOne)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, firfi.FIToFI.LineTwo)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, firfi.FIToFI.LineThree)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, firfi.FIToFI.LineFour)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, firfi.FIToFI.LineFive)
	}
	if err := firfi.isAlphanumeric(firfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, firfi.FIToFI.LineSix)
	}
	return nil
}
