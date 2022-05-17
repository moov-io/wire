// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIBeneficiaryAdvice{}

// FIBeneficiaryAdvice is the financial institution beneficiary advice
type FIBeneficiaryAdvice struct {
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

// NewFIBeneficiaryAdvice returns a new FIBeneficiaryAdvice
func NewFIBeneficiaryAdvice(isVariable bool) *FIBeneficiaryAdvice {
	fiba := &FIBeneficiaryAdvice{
		tag:              TagFIBeneficiaryAdvice,
		isVariableLength: isVariable,
	}
	return fiba
}

// Parse takes the input string and parses the FIBeneficiaryAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiba *FIBeneficiaryAdvice) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 13 {
		return 0, NewTagWrongLengthErr(13, len(record))
	}
	fiba.tag = record[:6]

	return 6 + fiba.Advice.Parse(record[6:]), nil
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
func (fiba *FIBeneficiaryAdvice) String() string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fiba.tag)
	buf.WriteString(fiba.Advice.String(fiba.isVariableLength))

	return buf.String()
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
