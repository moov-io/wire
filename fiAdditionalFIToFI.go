// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (fifi *FIAdditionalFIToFI) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	fifi.tag = record[:6]
	length := 6

	value, read, err := fifi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fifi.AdditionalFIToFI.LineOne = value
	length += read

	value, read, err = fifi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fifi.AdditionalFIToFI.LineTwo = value
	length += read

	value, read, err = fifi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fifi.AdditionalFIToFI.LineThree = value
	length += read

	value, read, err = fifi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fifi.AdditionalFIToFI.LineFour = value
	length += read

	value, read, err = fifi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fifi.AdditionalFIToFI.LineFive = value
	length += read

	value, read, err = fifi.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fifi.AdditionalFIToFI.LineSix = value
	length += read

	if err := fifi.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width FIAdditionalFIToFI record
func (fifi *FIAdditionalFIToFI) String() string {
	return fifi.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIAdditionalFIToFI record formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(216)

	buf.WriteString(fifi.tag)
	buf.WriteString(fifi.FormatLineOne(options))
	buf.WriteString(fifi.FormatLineTwo(options))
	buf.WriteString(fifi.FormatLineThree(options))
	buf.WriteString(fifi.FormatLineFour(options))
	buf.WriteString(fifi.FormatLineFive(options))
	buf.WriteString(fifi.FormatLineSix(options))

	if options.VariableLengthFields {
		return fifi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
	return fifi.alphaField(fifi.AdditionalFIToFI.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (fifi *FIAdditionalFIToFI) LineTwoField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (fifi *FIAdditionalFIToFI) LineThreeField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (fifi *FIAdditionalFIToFI) LineFourField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (fifi *FIAdditionalFIToFI) LineFiveField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (fifi *FIAdditionalFIToFI) LineSixField() string {
	return fifi.alphaField(fifi.AdditionalFIToFI.LineSix, 35)
}

// FormatLineOne returns AdditionalFIToFI.LineOne formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) FormatLineOne(options FormatOptions) string {
	return fifi.formatAlphaField(fifi.AdditionalFIToFI.LineOne, 35, options)
}

// FormatLineTwo returns AdditionalFIToFI.LineTwo formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) FormatLineTwo(options FormatOptions) string {
	return fifi.formatAlphaField(fifi.AdditionalFIToFI.LineTwo, 35, options)
}

// FormatLineThree returns AdditionalFIToFI.LineThree formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) FormatLineThree(options FormatOptions) string {
	return fifi.formatAlphaField(fifi.AdditionalFIToFI.LineThree, 35, options)
}

// FormatLineFour returns AdditionalFIToFI.LineFour formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) FormatLineFour(options FormatOptions) string {
	return fifi.formatAlphaField(fifi.AdditionalFIToFI.LineFour, 35, options)
}

// FormatLineFive returns AdditionalFIToFI.LineFive formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) FormatLineFive(options FormatOptions) string {
	return fifi.formatAlphaField(fifi.AdditionalFIToFI.LineFive, 35, options)
}

// FormatLineSix returns AdditionalFIToFI.LineSix formatted according to the FormatOptions
func (fifi *FIAdditionalFIToFI) FormatLineSix(options FormatOptions) string {
	return fifi.formatAlphaField(fifi.AdditionalFIToFI.LineSix, 35, options)
}
