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

	if len(record) != length {
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

// String writes FIBeneficiaryAdvice
func (fiba *FIBeneficiaryAdvice) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fiba.tag)
	buf.WriteString(fiba.AdviceCodeField())
	buf.WriteString(fiba.LineOneField(options...))
	buf.WriteString(fiba.LineTwoField(options...))
	buf.WriteString(fiba.LineThreeField(options...))
	buf.WriteString(fiba.LineFourField(options...))
	buf.WriteString(fiba.LineFiveField(options...))
	buf.WriteString(fiba.LineSixField(options...))

	if fiba.parseFirstOption(options) {
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
func (fiba *FIBeneficiaryAdvice) LineOneField(options ...bool) string {
	return fiba.alphaVariableField(fiba.Advice.LineOne, 26, fiba.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (fiba *FIBeneficiaryAdvice) LineTwoField(options ...bool) string {
	return fiba.alphaVariableField(fiba.Advice.LineTwo, 33, fiba.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (fiba *FIBeneficiaryAdvice) LineThreeField(options ...bool) string {
	return fiba.alphaVariableField(fiba.Advice.LineThree, 33, fiba.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (fiba *FIBeneficiaryAdvice) LineFourField(options ...bool) string {
	return fiba.alphaVariableField(fiba.Advice.LineFour, 33, fiba.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (fiba *FIBeneficiaryAdvice) LineFiveField(options ...bool) string {
	return fiba.alphaVariableField(fiba.Advice.LineFive, 33, fiba.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (fiba *FIBeneficiaryAdvice) LineSixField(options ...bool) string {
	return fiba.alphaVariableField(fiba.Advice.LineSix, 33, fiba.parseFirstOption(options))
}
