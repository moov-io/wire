// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIIntermediaryFIAdvice is the financial institution intermediary financial institution
type FIIntermediaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFIAdvice returns a new FIIntermediaryFIAdvice
func NewFIIntermediaryFIAdvice() *FIIntermediaryFIAdvice {
	fiifia := &FIIntermediaryFIAdvice{
		tag: TagFIIntermediaryFIAdvice,
	}
	return fiifia
}

// Parse takes the input string and parses the FIIntermediaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifia *FIIntermediaryFIAdvice) Parse(record string) error {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagMinLengthErr(9, len(record))
	}

	fiifia.tag = record[:6]
	fiifia.Advice.AdviceCode = fiifia.parseStringField(record[6:9])
	length := 9

	value, read, err := fiifia.parseVariableStringField(record[length:], 26)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fiifia.Advice.LineOne = value
	length += read

	value, read, err = fiifia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fiifia.Advice.LineTwo = value
	length += read

	value, read, err = fiifia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fiifia.Advice.LineThree = value
	length += read

	value, read, err = fiifia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fiifia.Advice.LineFour = value
	length += read

	value, read, err = fiifia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fiifia.Advice.LineFive = value
	length += read

	value, read, err = fiifia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fiifia.Advice.LineSix = value
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (fiifia *FIIntermediaryFIAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIIntermediaryFIAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fiifia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fiifia.tag = TagFIIntermediaryFIAdvice
	return nil
}

// String writes FIIntermediaryFIAdvice
func (fiifia *FIIntermediaryFIAdvice) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fiifia.tag)
	buf.WriteString(fiifia.AdviceCodeField())
	buf.WriteString(fiifia.LineOneField(options...))
	buf.WriteString(fiifia.LineTwoField(options...))
	buf.WriteString(fiifia.LineThreeField(options...))
	buf.WriteString(fiifia.LineFourField(options...))
	buf.WriteString(fiifia.LineFiveField(options...))
	buf.WriteString(fiifia.LineSixField(options...))

	if fiifia.parseFirstOption(options) {
		return fiifia.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIIntermediaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiifia *FIIntermediaryFIAdvice) Validate() error {
	if fiifia.tag != TagFIIntermediaryFIAdvice {
		return fieldError("tag", ErrValidTagForType, fiifia.tag)
	}
	if err := fiifia.isAdviceCode(fiifia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fiifia.Advice.AdviceCode)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fiifia.Advice.LineOne)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiifia.Advice.LineTwo)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fiifia.Advice.LineThree)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fiifia.Advice.LineFour)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fiifia.Advice.LineFive)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fiifia.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fiifia *FIIntermediaryFIAdvice) AdviceCodeField() string {
	return fiifia.alphaField(fiifia.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fiifia *FIIntermediaryFIAdvice) LineOneField(options ...bool) string {
	return fiifia.alphaVariableField(fiifia.Advice.LineOne, 26, fiifia.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (fiifia *FIIntermediaryFIAdvice) LineTwoField(options ...bool) string {
	return fiifia.alphaVariableField(fiifia.Advice.LineTwo, 33, fiifia.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (fiifia *FIIntermediaryFIAdvice) LineThreeField(options ...bool) string {
	return fiifia.alphaVariableField(fiifia.Advice.LineThree, 33, fiifia.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (fiifia *FIIntermediaryFIAdvice) LineFourField(options ...bool) string {
	return fiifia.alphaVariableField(fiifia.Advice.LineFour, 33, fiifia.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (fiifia *FIIntermediaryFIAdvice) LineFiveField(options ...bool) string {
	return fiifia.alphaVariableField(fiifia.Advice.LineFive, 33, fiifia.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (fiifia *FIIntermediaryFIAdvice) LineSixField(options ...bool) string {
	return fiifia.alphaVariableField(fiifia.Advice.LineSix, 33, fiifia.parseFirstOption(options))
}
