// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIBeneficiaryFIAdvice is the financial institution beneficiary financial institution
type FIBeneficiaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFIAdvice returns a new FIBeneficiaryFIAdvice
func NewFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := &FIBeneficiaryFIAdvice{
		tag: TagFIBeneficiaryFIAdvice,
	}
	return fibfia
}

// Parse takes the input string and parses the FIBeneficiaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfia *FIBeneficiaryFIAdvice) Parse(record string) error {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagMinLengthErr(9, len(record))
	}

	fibfia.tag = record[:6]
	fibfia.Advice.AdviceCode = fibfia.parseStringField(record[6:9])

	var err error
	length := 9
	read := 0

	if fibfia.Advice.LineOne, read, err = fibfia.parseVariableStringField(record[length:], 26); err != nil {
		return fieldError("LineOne", err)
	}
	length += read

	if fibfia.Advice.LineTwo, read, err = fibfia.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineTwo", err)
	}
	length += read

	if fibfia.Advice.LineThree, read, err = fibfia.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineThree", err)
	}
	length += read

	if fibfia.Advice.LineFour, read, err = fibfia.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineFour", err)
	}
	length += read

	if fibfia.Advice.LineFive, read, err = fibfia.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineFive", err)
	}
	length += read

	if fibfia.Advice.LineSix, read, err = fibfia.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineSix", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (fibfia *FIBeneficiaryFIAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFIAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfia.tag = TagFIBeneficiaryFIAdvice
	return nil
}

// String writes FIBeneficiaryFIAdvice
func (fibfia *FIBeneficiaryFIAdvice) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fibfia.tag)
	buf.WriteString(fibfia.AdviceCodeField())
	buf.WriteString(fibfia.LineOneField(options...))
	buf.WriteString(fibfia.LineTwoField(options...))
	buf.WriteString(fibfia.LineThreeField(options...))
	buf.WriteString(fibfia.LineFourField(options...))
	buf.WriteString(fibfia.LineFiveField(options...))
	buf.WriteString(fibfia.LineSixField(options...))

	if fibfia.parseFirstOption(options) {
		return fibfia.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIBeneficiaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfia *FIBeneficiaryFIAdvice) Validate() error {
	if fibfia.tag != TagFIBeneficiaryFIAdvice {
		return fieldError("tag", ErrValidTagForType, fibfia.tag)
	}
	if err := fibfia.isAdviceCode(fibfia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fibfia.Advice.AdviceCode)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fibfia.Advice.LineOne)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfia.Advice.LineTwo)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fibfia.Advice.LineThree)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fibfia.Advice.LineFour)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fibfia.Advice.LineFive)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fibfia.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fibfia *FIBeneficiaryFIAdvice) AdviceCodeField() string {
	return fibfia.alphaField(fibfia.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fibfia *FIBeneficiaryFIAdvice) LineOneField(options ...bool) string {
	return fibfia.alphaVariableField(fibfia.Advice.LineOne, 26, fibfia.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (fibfia *FIBeneficiaryFIAdvice) LineTwoField(options ...bool) string {
	return fibfia.alphaVariableField(fibfia.Advice.LineTwo, 33, fibfia.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (fibfia *FIBeneficiaryFIAdvice) LineThreeField(options ...bool) string {
	return fibfia.alphaVariableField(fibfia.Advice.LineThree, 33, fibfia.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (fibfia *FIBeneficiaryFIAdvice) LineFourField(options ...bool) string {
	return fibfia.alphaVariableField(fibfia.Advice.LineFour, 33, fibfia.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (fibfia *FIBeneficiaryFIAdvice) LineFiveField(options ...bool) string {
	return fibfia.alphaVariableField(fibfia.Advice.LineFive, 33, fibfia.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (fibfia *FIBeneficiaryFIAdvice) LineSixField(options ...bool) string {
	return fibfia.alphaVariableField(fibfia.Advice.LineSix, 33, fibfia.parseFirstOption(options))
}
