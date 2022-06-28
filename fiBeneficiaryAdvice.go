// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIBeneficiaryAdvice is the financial institution beneficiary advice
type FIBeneficiaryAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryAdvice returns a new FIBeneficiaryAdvice
func NewFIBeneficiaryAdvice() *FIBeneficiaryAdvice {
	fiba := &FIBeneficiaryAdvice{
		tag: TagFIBeneficiaryAdvice,
	}
	return fiba
}

// Parse takes the input string and parses the FIBeneficiaryAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiba *FIBeneficiaryAdvice) Parse(record string) error {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagMinLengthErr(9, len(record))
	}

	fiba.tag = record[:6]
	fiba.Advice.AdviceCode = fiba.parseStringField(record[6:9])
	length := 9

	value, read, err := fiba.parseVariableStringField(record[length:], 26)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fiba.Advice.LineOne = value
	length += read

	value, read, err = fiba.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fiba.Advice.LineTwo = value
	length += read

	value, read, err = fiba.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fiba.Advice.LineThree = value
	length += read

	value, read, err = fiba.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fiba.Advice.LineFour = value
	length += read

	value, read, err = fiba.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fiba.Advice.LineFive = value
	length += read

	value, read, err = fiba.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fiba.Advice.LineSix = value
	length += read

	if !fiba.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (fiba *FIBeneficiaryAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fiba),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fiba.tag = TagFIBeneficiaryAdvice
	return nil
}

// String returns a fixed-width FIBeneficiaryAdvice record
func (fiba *FIBeneficiaryAdvice) String() string {
	return fiba.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIBeneficiaryAdvice record formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fiba.tag)
	buf.WriteString(fiba.AdviceCodeField())
	buf.WriteString(fiba.FormatLineOne(options))
	buf.WriteString(fiba.FormatLineTwo(options))
	buf.WriteString(fiba.FormatLineThree(options))
	buf.WriteString(fiba.FormatLineFour(options))
	buf.WriteString(fiba.FormatLineFive(options))
	buf.WriteString(fiba.FormatLineSix(options))

	if options.VariableLengthFields {
		return fiba.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIBeneficiaryAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiba *FIBeneficiaryAdvice) Validate() error {
	if fiba.tag != TagFIBeneficiaryAdvice {
		return fieldError("tag", ErrValidTagForType, fiba.tag)
	}
	if err := fiba.isAdviceCode(fiba.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fiba.Advice.AdviceCode)
	}
	if err := fiba.isAlphanumeric(fiba.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fiba.Advice.LineOne)
	}
	if err := fiba.isAlphanumeric(fiba.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiba.Advice.LineTwo)
	}
	if err := fiba.isAlphanumeric(fiba.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fiba.Advice.LineThree)
	}
	if err := fiba.isAlphanumeric(fiba.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fiba.Advice.LineFour)
	}
	if err := fiba.isAlphanumeric(fiba.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fiba.Advice.LineFive)
	}
	if err := fiba.isAlphanumeric(fiba.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fiba.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fiba *FIBeneficiaryAdvice) AdviceCodeField() string {
	return fiba.alphaField(fiba.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fiba *FIBeneficiaryAdvice) LineOneField() string {
	return fiba.alphaField(fiba.Advice.LineOne, 26)
}

// LineTwoField gets a string of the LineTwo field
func (fiba *FIBeneficiaryAdvice) LineTwoField() string {
	return fiba.alphaField(fiba.Advice.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fiba *FIBeneficiaryAdvice) LineThreeField() string {
	return fiba.alphaField(fiba.Advice.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fiba *FIBeneficiaryAdvice) LineFourField() string {
	return fiba.alphaField(fiba.Advice.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fiba *FIBeneficiaryAdvice) LineFiveField() string {
	return fiba.alphaField(fiba.Advice.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fiba *FIBeneficiaryAdvice) LineSixField() string {
	return fiba.alphaField(fiba.Advice.LineSix, 33)
}

// FormatLineOne returns Advice.LineOne formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) FormatLineOne(options FormatOptions) string {
	return fiba.formatAlphaField(fiba.Advice.LineOne, 26, options)
}

// FormatLineTwo returns Advice.LineTwo formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) FormatLineTwo(options FormatOptions) string {
	return fiba.formatAlphaField(fiba.Advice.LineTwo, 33, options)
}

// FormatLineThree returns Advice.LineThree formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) FormatLineThree(options FormatOptions) string {
	return fiba.formatAlphaField(fiba.Advice.LineThree, 33, options)
}

// FormatLineFour returns Advice.LineFour formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) FormatLineFour(options FormatOptions) string {
	return fiba.formatAlphaField(fiba.Advice.LineFour, 33, options)
}

// FormatLineFive returns Advice.LineFive formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) FormatLineFive(options FormatOptions) string {
	return fiba.formatAlphaField(fiba.Advice.LineFive, 33, options)
}

// FormatLineSix returns Advice.LineSix formatted according to the FormatOptions
func (fiba *FIBeneficiaryAdvice) FormatLineSix(options FormatOptions) string {
	return fiba.formatAlphaField(fiba.Advice.LineSix, 33, options)
}
