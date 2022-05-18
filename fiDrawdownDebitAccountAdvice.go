// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIDrawdownDebitAccountAdvice{}

// FIDrawdownDebitAccountAdvice is the financial institution drawdown debit account advice
type FIDrawdownDebitAccountAdvice struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIDrawdownDebitAccountAdvice returns a new FIDrawdownDebitAccountAdvice
func NewFIDrawdownDebitAccountAdvice(isVariable bool) *FIDrawdownDebitAccountAdvice {
	debitDDAdvice := &FIDrawdownDebitAccountAdvice{
		tag:              TagFIDrawdownDebitAccountAdvice,
		isVariableLength: isVariable,
	}
	return debitDDAdvice
}

// Parse takes the input string and parses the FIDrawdownDebitAccountAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 13 {
		return 0, NewTagWrongLengthErr(13, len(record))
	}

	var err error
	var length, read int

	if debitDDAdvice.tag, read, err = debitDDAdvice.parseTag(record); err != nil {
		return 0, fieldError("FIDrawdownDebitAccountAdvice.Tag", err)
	}
	length += read

	if read, err = debitDDAdvice.Advice.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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
func (debitDDAdvice *FIDrawdownDebitAccountAdvice) String() string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(debitDDAdvice.tag)
	buf.WriteString(debitDDAdvice.Advice.String(debitDDAdvice.isVariableLength))

	return buf.String()
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
