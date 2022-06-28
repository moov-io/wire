// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIBeneficiaryFI is the financial institution beneficiary financial institution
type FIBeneficiaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFI returns a new FIBeneficiaryFI
func NewFIBeneficiaryFI() *FIBeneficiaryFI {
	fibfi := &FIBeneficiaryFI{
		tag: TagFIBeneficiaryFI,
	}
	return fibfi
}

// Parse takes the input string and parses the FIBeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfi *FIBeneficiaryFI) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	fibfi.tag = record[:6]
	length := 6

	value, read, err := fibfi.parseVariableStringField(record[length:], 30)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fibfi.FIToFI.LineOne = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fibfi.FIToFI.LineTwo = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fibfi.FIToFI.LineThree = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fibfi.FIToFI.LineFour = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fibfi.FIToFI.LineFive = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fibfi.FIToFI.LineSix = value
	length += read

	if !fibfi.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (fibfi *FIBeneficiaryFI) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfi.tag = TagFIBeneficiaryFI
	return nil
}

// String returns a fixed-width FIBeneficiaryFI record
func (fibfi *FIBeneficiaryFI) String() string {
	return fibfi.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIBeneficiaryFI record formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fibfi.tag)
	buf.WriteString(fibfi.FormatLineOne(options))
	buf.WriteString(fibfi.FormatLineTwo(options))
	buf.WriteString(fibfi.FormatLineThree(options))
	buf.WriteString(fibfi.FormatLineFour(options))
	buf.WriteString(fibfi.FormatLineFive(options))
	buf.WriteString(fibfi.FormatLineSix(options))

	if options.VariableLengthFields {
		return fibfi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIBeneficiaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfi *FIBeneficiaryFI) Validate() error {
	if fibfi.tag != TagFIBeneficiaryFI {
		return fieldError("tag", ErrValidTagForType, fibfi.tag)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fibfi.FIToFI.LineOne)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfi.FIToFI.LineTwo)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fibfi.FIToFI.LineThree)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fibfi.FIToFI.LineFour)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fibfi.FIToFI.LineFive)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fibfi.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fibfi *FIBeneficiaryFI) LineOneField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineOne, 30)
}

// LineTwoField gets a string of the LineTwo field
func (fibfi *FIBeneficiaryFI) LineTwoField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fibfi *FIBeneficiaryFI) LineThreeField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fibfi *FIBeneficiaryFI) LineFourField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fibfi *FIBeneficiaryFI) LineFiveField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fibfi *FIBeneficiaryFI) LineSixField() string {
	return fibfi.alphaField(fibfi.FIToFI.LineSix, 33)
}

// FormatLineOne returns FIToFI.LineOne formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) FormatLineOne(options FormatOptions) string {
	return fibfi.formatAlphaField(fibfi.FIToFI.LineOne, 30, options)
}

// FormatLineTwo returns FIToFI.LineTwo formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) FormatLineTwo(options FormatOptions) string {
	return fibfi.formatAlphaField(fibfi.FIToFI.LineTwo, 33, options)
}

// FormatLineThree FIToFI.LineThree LineOne formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) FormatLineThree(options FormatOptions) string {
	return fibfi.formatAlphaField(fibfi.FIToFI.LineThree, 33, options)
}

// FormatLineFour returns FIToFI.LineFour formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) FormatLineFour(options FormatOptions) string {
	return fibfi.formatAlphaField(fibfi.FIToFI.LineFour, 33, options)
}

// FormatLineFive returns FIToFI.LineFive formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) FormatLineFive(options FormatOptions) string {
	return fibfi.formatAlphaField(fibfi.FIToFI.LineFive, 33, options)
}

// FormatLineSix returns FIToFI.LineSix formatted according to the FormatOptions
func (fibfi *FIBeneficiaryFI) FormatLineSix(options FormatOptions) string {
	return fibfi.formatAlphaField(fibfi.FIToFI.LineSix, 33, options)
}
