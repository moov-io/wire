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

	if err := fib.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width FIBeneficiary record
func (fib *FIBeneficiary) String() string {
	return fib.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIBeneficiary record formatted according to the FormatOptions
func (fib *FIBeneficiary) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(201)
	buf.WriteString(fib.tag)

	buf.WriteString(fib.FormatLineOne(options))
	buf.WriteString(fib.FormatLineTwo(options))
	buf.WriteString(fib.FormatLineThree(options))
	buf.WriteString(fib.FormatLineFour(options))
	buf.WriteString(fib.FormatLineFive(options))
	buf.WriteString(fib.FormatLineSix(options))

	if options.VariableLengthFields {
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
func (fib *FIBeneficiary) LineOneField() string {
	return fib.alphaField(fib.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (fib *FIBeneficiary) LineTwoField() string {
	return fib.alphaField(fib.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fib *FIBeneficiary) LineThreeField() string {
	return fib.alphaField(fib.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fib *FIBeneficiary) LineFourField() string {
	return fib.alphaField(fib.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fib *FIBeneficiary) LineFiveField() string {
	return fib.alphaField(fib.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fib *FIBeneficiary) LineSixField() string {
	return fib.alphaField(fib.FIToFI.LineSix, 33)
}

// FormatLineOne returns FIToFI.LineOne formatted according to the FormatOptions
func (fib *FIBeneficiary) FormatLineOne(options FormatOptions) string {
	return fib.formatAlphaField(fib.FIToFI.LineOne, 30, options)
}

// FormatLineTwo returns FIToFI.LineTwo formatted according to the FormatOptions
func (fib *FIBeneficiary) FormatLineTwo(options FormatOptions) string {
	return fib.formatAlphaField(fib.FIToFI.LineTwo, 33, options)
}

// FormatLineThree returns FIToFI.LineThree formatted according to the FormatOptions
func (fib *FIBeneficiary) FormatLineThree(options FormatOptions) string {
	return fib.formatAlphaField(fib.FIToFI.LineThree, 33, options)
}

// FormatLineFour returns FIToFI.LineFour formatted according to the FormatOptions
func (fib *FIBeneficiary) FormatLineFour(options FormatOptions) string {
	return fib.formatAlphaField(fib.FIToFI.LineFour, 33, options)
}

// FormatLineFive returns FIToFI.LineFive formatted according to the FormatOptions
func (fib *FIBeneficiary) FormatLineFive(options FormatOptions) string {
	return fib.formatAlphaField(fib.FIToFI.LineFive, 33, options)
}

// FormatLineSix returns FIToFI.LineSix formatted according to the FormatOptions
func (fib *FIBeneficiary) FormatLineSix(options FormatOptions) string {
	return fib.formatAlphaField(fib.FIToFI.LineSix, 33, options)
}
