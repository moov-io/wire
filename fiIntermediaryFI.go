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
	length := 6

	value, read, err := fiifi.parseVariableStringField(record[length:], 30)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fiifi.FIToFI.LineOne = value
	length += read

	value, read, err = fiifi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fiifi.FIToFI.LineTwo = value
	length += read

	value, read, err = fiifi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fiifi.FIToFI.LineThree = value
	length += read

	value, read, err = fiifi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fiifi.FIToFI.LineFour = value
	length += read

	value, read, err = fiifi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fiifi.FIToFI.LineFive = value
	length += read

	value, read, err = fiifi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fiifi.FIToFI.LineSix = value
	length += read

	if !fiifi.verifyDataWithReadLength(record, length) {
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

// String returns a fixed-width FIIntermediaryFI record
func (fiifi *FIIntermediaryFI) String() string {
	return fiifi.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIIntermediaryFI record formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fiifi.tag)
	buf.WriteString(fiifi.FormatLineOne(options))
	buf.WriteString(fiifi.FormatLineTwo(options))
	buf.WriteString(fiifi.FormatLineThree(options))
	buf.WriteString(fiifi.FormatLineFour(options))
	buf.WriteString(fiifi.FormatLineFive(options))
	buf.WriteString(fiifi.FormatLineSix(options))

	if options.VariableLengthFields {
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
func (fiifi *FIIntermediaryFI) LineOneField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (fiifi *FIIntermediaryFI) LineTwoField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fiifi *FIIntermediaryFI) LineThreeField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fiifi *FIIntermediaryFI) LineFourField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fiifi *FIIntermediaryFI) LineFiveField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fiifi *FIIntermediaryFI) LineSixField() string {
	return fiifi.alphaField(fiifi.FIToFI.LineSix, 33)
}

// FormatLineOne returns FIToFI.LineOne formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) FormatLineOne(options FormatOptions) string {
	return fiifi.formatAlphaField(fiifi.FIToFI.LineOne, 30, options)
}

// FormatLineTwo returns FIToFI.LineTwo formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) FormatLineTwo(options FormatOptions) string {
	return fiifi.formatAlphaField(fiifi.FIToFI.LineTwo, 33, options)
}

// FormatLineThree returns FIToFI.LineThree formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) FormatLineThree(options FormatOptions) string {
	return fiifi.formatAlphaField(fiifi.FIToFI.LineThree, 33, options)
}

// FormatLineFour returns FIToFI.LineFour formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) FormatLineFour(options FormatOptions) string {
	return fiifi.formatAlphaField(fiifi.FIToFI.LineFour, 33, options)
}

// FormatLineFive returns FIToFI.LineFive formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) FormatLineFive(options FormatOptions) string {
	return fiifi.formatAlphaField(fiifi.FIToFI.LineFive, 33, options)
}

// FormatLineSix returns FIToFI.LineSix formatted according to the FormatOptions
func (fiifi *FIIntermediaryFI) FormatLineSix(options FormatOptions) string {
	return fiifi.formatAlphaField(fiifi.FIToFI.LineSix, 33, options)
}
