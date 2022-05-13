// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIIntermediaryFI{}

// FIIntermediaryFIAdvice is the financial institution intermediary financial institution
type FIIntermediaryFIAdvice struct {
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

// NewFIIntermediaryFIAdvice returns a new FIIntermediaryFIAdvice
func NewFIIntermediaryFIAdvice(isVariable bool) *FIIntermediaryFIAdvice {
	fiifia := &FIIntermediaryFIAdvice{
		tag:              TagFIIntermediaryFIAdvice,
		isVariableLength: isVariable,
	}
	return fiifia
}

// Parse takes the input string and parses the FIIntermediaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifia *FIIntermediaryFIAdvice) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 13 {
		return NewTagWrongLengthErr(13, len(record)), 0
	}
	fiifia.tag = record[:6]

	return nil, 6 + fiifia.Advice.Parse(record[6:])
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
func (fiifia *FIIntermediaryFIAdvice) String() string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fiifia.tag)
	buf.WriteString(fiifia.Advice.String(fiifia.isVariableLength))

	return buf.String()
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
