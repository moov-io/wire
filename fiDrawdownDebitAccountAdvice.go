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
	length := 9

	value, read, err := debitDDAdvice.parseVariableStringField(record[length:], 26)
	if err != nil {
		return fieldError("LineOne", err)
	}
	debitDDAdvice.Advice.LineOne = value
	length += read

	value, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	debitDDAdvice.Advice.LineTwo = value
	length += read

	value, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	debitDDAdvice.Advice.LineThree = value
	length += read

	value, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	debitDDAdvice.Advice.LineFour = value
	length += read

	value, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	debitDDAdvice.Advice.LineFive = value
	length += read

	value, read, err = debitDDAdvice.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	debitDDAdvice.Advice.LineSix = value
	length += read

	if err := debitDDAdvice.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width FIDrawdownDebitAccountAdvice record
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) String() string {
	return debitDDAdvice.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIDrawdownDebitAccountAdvice record formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(debitDDAdvice.tag)
	buf.WriteString(debitDDAdvice.AdviceCodeField())
	buf.WriteString(debitDDAdvice.FormatLineOne(options))
	buf.WriteString(debitDDAdvice.FormatLineTwo(options))
	buf.WriteString(debitDDAdvice.FormatLineThree(options))
	buf.WriteString(debitDDAdvice.FormatLineFour(options))
	buf.WriteString(debitDDAdvice.FormatLineFive(options))
	buf.WriteString(debitDDAdvice.FormatLineSix(options))

	if options.VariableLengthFields {
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
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineOneField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.LineOne, 26)
}

// LineTwoField gets a string of the LineTwo field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineTwoField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineThreeField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineFourField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineFiveField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) LineSixField() string {
	return debitDDAdvice.alphaField(debitDDAdvice.Advice.LineSix, 33)
}

// FormatLineOne returns Advice.LineOne formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) FormatLineOne(options FormatOptions) string {
	return debitDDAdvice.formatAlphaField(debitDDAdvice.Advice.LineOne, 26, options)
}

// FormatLineTwo returns Advice.LineTwo formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) FormatLineTwo(options FormatOptions) string {
	return debitDDAdvice.formatAlphaField(debitDDAdvice.Advice.LineTwo, 33, options)
}

// FormatLineThree returns Advice.LineThree formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) FormatLineThree(options FormatOptions) string {
	return debitDDAdvice.formatAlphaField(debitDDAdvice.Advice.LineThree, 33, options)
}

// FormatLineFour returns Advice.LineFour formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) FormatLineFour(options FormatOptions) string {
	return debitDDAdvice.formatAlphaField(debitDDAdvice.Advice.LineFour, 33, options)
}

// FormatLineFive returns Advice.LineFive formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) FormatLineFive(options FormatOptions) string {
	return debitDDAdvice.formatAlphaField(debitDDAdvice.Advice.LineFive, 33, options)
}

// FormatLineSix returns Advice.LineSix formatted according to the FormatOptions
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) FormatLineSix(options FormatOptions) string {
	return debitDDAdvice.formatAlphaField(debitDDAdvice.Advice.LineSix, 33, options)
}
