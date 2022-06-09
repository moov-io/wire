// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIBeneficiary is the financial institution beneficiary
type FIBeneficiary struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiary returns a new FIBeneficiary
func NewFIBeneficiary() *FIBeneficiary {
	fib := &FIBeneficiary{
		tag: TagFIBeneficiary,
	}
	return fib
}

// Parse takes the input string and parses the FIBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fib *FIBeneficiary) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	fib.tag = record[:6]
	length := 6

	value, read, err := fib.parseVariableStringField(record[length:], 30)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fib.FIToFI.LineOne = value
	length += read

	value, read, err = fib.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fib.FIToFI.LineTwo = value
	length += read

	value, read, err = fib.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fib.FIToFI.LineThree = value
	length += read

	value, read, err = fib.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fib.FIToFI.LineFour = value
	length += read

	value, read, err = fib.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fib.FIToFI.LineFive = value
	length += read

	value, read, err = fib.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fib.FIToFI.LineSix = value
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (fib *FIBeneficiary) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(fib),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fib.tag = TagFIBeneficiary
	return nil
}

// String writes FIBeneficiary
func (fib *FIBeneficiary) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(201)
	buf.WriteString(fib.tag)

	buf.WriteString(fib.LineOneField(options...))
	buf.WriteString(fib.LineTwoField(options...))
	buf.WriteString(fib.LineThreeField(options...))
	buf.WriteString(fib.LineFourField(options...))
	buf.WriteString(fib.LineFiveField(options...))
	buf.WriteString(fib.LineSixField(options...))

	if fib.parseFirstOption(options) {
		return fib.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fib *FIBeneficiary) Validate() error {
	if fib.tag != TagFIBeneficiary {
		return fieldError("tag", ErrValidTagForType, fib.tag)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fib.FIToFI.LineOne)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fib.FIToFI.LineTwo)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fib.FIToFI.LineThree)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fib.FIToFI.LineFour)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fib.FIToFI.LineFive)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fib.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fib *FIBeneficiary) LineOneField(options ...bool) string {
	return fib.alphaVariableField(fib.FIToFI.LineOne, 30, fib.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (fib *FIBeneficiary) LineTwoField(options ...bool) string {
	return fib.alphaVariableField(fib.FIToFI.LineTwo, 33, fib.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (fib *FIBeneficiary) LineThreeField(options ...bool) string {
	return fib.alphaVariableField(fib.FIToFI.LineThree, 33, fib.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (fib *FIBeneficiary) LineFourField(options ...bool) string {
	return fib.alphaVariableField(fib.FIToFI.LineFour, 33, fib.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (fib *FIBeneficiary) LineFiveField(options ...bool) string {
	return fib.alphaVariableField(fib.FIToFI.LineFive, 33, fib.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (fib *FIBeneficiary) LineSixField(options ...bool) string {
	return fib.alphaVariableField(fib.FIToFI.LineSix, 33, fib.parseFirstOption(options))
}
