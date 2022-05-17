// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIBeneficiary{}

// FIBeneficiary is the financial institution beneficiary
type FIBeneficiary struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiary returns a new FIBeneficiary
func NewFIBeneficiary(isVariable bool) *FIBeneficiary {
	fib := &FIBeneficiary{
		tag:              TagFIBeneficiary,
		isVariableLength: isVariable,
	}
	return fib
}

// Parse takes the input string and parses the FIBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fib *FIBeneficiary) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}
	fib.tag = record[:6]

	return 6 + fib.FIToFI.Parse(record[6:]), nil
}

func (fib *FIBeneficiary) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(fib),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fib.tag = TagFIBeneficiary
	return nil
}

// String writes FIBeneficiary
func (fib *FIBeneficiary) String() string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fib.tag)
	buf.WriteString(fib.FIToFI.String(fib.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fib *FIBeneficiary) Validate() error {
	if fib.tag != TagFIBeneficiary {
		return fieldError("tag", ErrValidTagForType, fib.tag)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fib.FIToFI.LineOne)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fib.FIToFI.LineTwo)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fib.FIToFI.LineThree)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fib.FIToFI.LineFour)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fib.FIToFI.LineFive)
	}
	if err := fib.isAlphanumeric(fib.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fib.FIToFI.LineSix)
	}
	return nil
}
