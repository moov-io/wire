// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIReceiverFI is the financial institution receiver financial institution
type FIReceiverFI struct {
	// tag
	tag string
	// FIToFI is financial institution to financial institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIReceiverFI returns a new FIReceiverFI
func NewFIReceiverFI() *FIReceiverFI {
	firfi := &FIReceiverFI{
		tag: TagFIReceiverFI,
	}
	return firfi
}

// Parse takes the input string and parses the FIReceiverFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (firfi *FIReceiverFI) Parse(record string) error {
	if utf8.RuneCountInString(record) != 201 {
		return NewTagWrongLengthErr(201, len(record))
	}
	firfi.tag = record[:6]
	firfi.FIToFI.LineOne = firfi.parseStringField(record[6:36])
	firfi.FIToFI.LineTwo = firfi.parseStringField(record[36:69])
	firfi.FIToFI.LineThree = firfi.parseStringField(record[69:102])
	firfi.FIToFI.LineFour = firfi.parseStringField(record[102:135])
	firfi.FIToFI.LineFive = firfi.parseStringField(record[135:168])
	firfi.FIToFI.LineSix = firfi.parseStringField(record[174:201])
	return nil
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
	buf.WriteString(firfi.LineOneField())
	buf.WriteString(firfi.LineTwoField())
	buf.WriteString(firfi.LineThreeField())
	buf.WriteString(firfi.LineFourField())
	buf.WriteString(firfi.LineFiveField())
	buf.WriteString(firfi.LineSixField())
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

// LineOneField gets a string of the LineOne field
func (firfi *FIReceiverFI) LineOneField() string {
	return firfi.alphaField(firfi.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (firfi *FIReceiverFI) LineTwoField() string {
	return firfi.alphaField(firfi.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (firfi *FIReceiverFI) LineThreeField() string {
	return firfi.alphaField(firfi.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (firfi *FIReceiverFI) LineFourField() string {
	return firfi.alphaField(firfi.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (firfi *FIReceiverFI) LineFiveField() string {
	return firfi.alphaField(firfi.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (firfi *FIReceiverFI) LineSixField() string {
	return firfi.alphaField(firfi.FIToFI.LineSix, 33)
}
