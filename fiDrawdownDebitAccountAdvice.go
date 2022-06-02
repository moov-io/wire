// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIDrawdownDebitAccountAdvice is the financial institution drawdown debit account advice
type FIDrawdownDebitAccountAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIDrawdownDebitAccountAdvice returns a new FIDrawdownDebitAccountAdvice
func NewFIDrawdownDebitAccountAdvice() *FIDrawdownDebitAccountAdvice {
	debitDDAdvice := &FIDrawdownDebitAccountAdvice{
		tag: TagFIDrawdownDebitAccountAdvice,
	}
	return debitDDAdvice
}

// Parse takes the input string and parses the FIDrawdownDebitAccountAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) Parse(record string) error {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagMinLengthErr(9, len(record))
	}

	debitDDAdvice.tag = record[:6]
	debitDDAdvice.Advice.AdviceCode = debitDDAdvice.parseStringField(record[6:9])

	var err error
	length := 9
	read := 0

	if debitDDAdvice.Advice.LineOne, read, err = debitDDAdvice.parseVariableStringField(record[length:], 26); err != nil {
		return fieldError("LineOne", err)
	}
	length += read

	if debitDDAdvice.Advice.LineTwo, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineTwo", err)
	}
	length += read

	if debitDDAdvice.Advice.LineThree, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineThree", err)
	}
	length += read

	if debitDDAdvice.Advice.LineFour, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineFour", err)
	}
	length += read

	if debitDDAdvice.Advice.LineFive, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineFive", err)
	}
	length += read

	if debitDDAdvice.Advice.LineSix, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33); err != nil {
		return fieldError("LineSix", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (debitDDAdvice *FIDrawdownDebitAccountAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIDrawdownDebitAccountAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(debitDDAdvice),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	debitDDAdvice.tag = TagFIDrawdownDebitAccountAdvice
	return nil
}

// String writes FIDrawdownDebitAccountAdvice
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(debitDDAdvice.tag)
	buf.WriteString(debitDDAdvice.AdviceCodeField())
	buf.WriteString(debitDDAdvice.LineOneField(options...))
	buf.WriteString(debitDDAdvice.LineTwoField(options...))
	buf.WriteString(debitDDAdvice.LineThreeField(options...))
	buf.WriteString(debitDDAdvice.LineFourField(options...))
	buf.WriteString(debitDDAdvice.LineFiveField(options...))
	buf.WriteString(debitDDAdvice.LineSixField(options...))

	if debitDDAdvice.parseFirstOption(options) {
		return debitDDAdvice.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIDrawdownDebitAccountAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) Validate() error {
	if debitDDAdvice.tag != TagFIDrawdownDebitAccountAdvice {
		return fieldError("tag", ErrValidTagForType, debitDDAdvice.tag)
	}
	if err := debitDDAdvice.isAdviceCode(debitDDAdvice.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, debitDDAdvice.Advice.AdviceCode)
	}
	if err := debitDDAdvice.isAlphanumeric(debitDDAdvice.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, debitDDAdvice.Advice.LineOne)
	}
	if err := debitDDAdvice.isAlphanumeric(debitDDAdvice.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, debitDDAdvice.Advice.LineTwo)
	}
	if err := debitDDAdvice.isAlphanumeric(debitDDAdvice.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, debitDDAdvice.Advice.LineThree)
	}
	if err := debitDDAdvice.isAlphanumeric(debitDDAdvice.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, debitDDAdvice.Advice.LineFour)
	}
	if err := debitDDAdvice.isAlphanumeric(debitDDAdvice.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, debitDDAdvice.Advice.LineFive)
	}
	if err := debitDDAdvice.isAlphanumeric(debitDDAdvice.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, debitDDAdvice.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) AdviceCodeField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineOneField(options ...bool) string {
	return debitDDAdvice.alphaVariableField(debitDDAdvice.Advice.LineOne, 26, debitDDAdvice.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineTwoField(options ...bool) string {
	return debitDDAdvice.alphaVariableField(debitDDAdvice.Advice.LineTwo, 33, debitDDAdvice.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineThreeField(options ...bool) string {
	return debitDDAdvice.alphaVariableField(debitDDAdvice.Advice.LineThree, 33, debitDDAdvice.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineFourField(options ...bool) string {
	return debitDDAdvice.alphaVariableField(debitDDAdvice.Advice.LineFour, 33, debitDDAdvice.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineFiveField(options ...bool) string {
	return debitDDAdvice.alphaVariableField(debitDDAdvice.Advice.LineFive, 33, debitDDAdvice.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineSixField(options ...bool) string {
	return debitDDAdvice.alphaVariableField(debitDDAdvice.Advice.LineSix, 33, debitDDAdvice.parseFirstOption(options))
}
