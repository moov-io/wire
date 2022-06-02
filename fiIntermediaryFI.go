// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIIntermediaryFI is the financial institution intermediary financial institution
type FIIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFI returns a new FIIntermediaryFI
func NewFIIntermediaryFI() *FIIntermediaryFI {
	fiifi := &FIIntermediaryFI{
		tag: TagFIIntermediaryFI,
	}
	return fiifi
}

// Parse takes the input string and parses the FIIntermediaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifi *FIIntermediaryFI) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	fiifi.tag = record[:6]

	var err error
	length := 6
	read := 0

	if fiifi.FIToFI.LineOne, read, err = fiifi.parseVariableStringField(record[length:], 30); err != nil {
		return fieldError("LineOne", err)
	}
	length += read

	if fiifi.FIToFI.LineTwo, read, err = fiifi.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineTwo", err)
	}
	length += read

	if fiifi.FIToFI.LineThree, read, err = fiifi.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineThree", err)
	}
	length += read

	if fiifi.FIToFI.LineFour, read, err = fiifi.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineFour", err)
	}
	length += read

	if fiifi.FIToFI.LineFive, read, err = fiifi.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineFive", err)
	}
	length += read

	if fiifi.FIToFI.LineSix, read, err = fiifi.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineSix", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
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
func (fiifi *FIIntermediaryFI) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fiifi.tag)
	buf.WriteString(fiifi.LineOneField(options...))
	buf.WriteString(fiifi.LineTwoField(options...))
	buf.WriteString(fiifi.LineThreeField(options...))
	buf.WriteString(fiifi.LineFourField(options...))
	buf.WriteString(fiifi.LineFiveField(options...))
	buf.WriteString(fiifi.LineSixField(options...))

	if fiifi.parseFirstOption(options) {
		return fiifi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// LineOneField gets a string of the LineOne field
func (fiifi *FIIntermediaryFI) LineOneField(options ...bool) string {
	return fiifi.alphaVariableField(fiifi.FIToFI.LineOne, 30, fiifi.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (fiifi *FIIntermediaryFI) LineTwoField(options ...bool) string {
	return fiifi.alphaVariableField(fiifi.FIToFI.LineTwo, 33, fiifi.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (fiifi *FIIntermediaryFI) LineThreeField(options ...bool) string {
	return fiifi.alphaVariableField(fiifi.FIToFI.LineThree, 33, fiifi.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (fiifi *FIIntermediaryFI) LineFourField(options ...bool) string {
	return fiifi.alphaVariableField(fiifi.FIToFI.LineFour, 33, fiifi.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (fiifi *FIIntermediaryFI) LineFiveField(options ...bool) string {
	return fiifi.alphaVariableField(fiifi.FIToFI.LineFive, 33, fiifi.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (fiifi *FIIntermediaryFI) LineSixField(options ...bool) string {
	return fiifi.alphaVariableField(fiifi.FIToFI.LineSix, 33, fiifi.parseFirstOption(options))
}
