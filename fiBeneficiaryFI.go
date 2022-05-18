// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIBeneficiaryFI{}

// FIBeneficiaryFI is the financial institution beneficiary financial institution
type FIBeneficiaryFI struct {
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

// NewFIBeneficiaryFI returns a new FIBeneficiaryFI
func NewFIBeneficiaryFI(isVariable bool) *FIBeneficiaryFI {
	fibfi := &FIBeneficiaryFI{
		tag:              TagFIBeneficiaryFI,
		isVariableLength: isVariable,
	}
	return fibfi
}

// Parse takes the input string and parses the FIBeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfi *FIBeneficiaryFI) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if fibfi.tag, read, err = fibfi.parseTag(record); err != nil {
		return 0, fieldError("FIBeneficiaryFI.Tag", err)
	}
	length += read

	if read, err = fibfi.FIToFI.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (fibfi *FIBeneficiaryFI) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfi.tag = TagFIBeneficiaryFI
	return nil
}

// String writes FIBeneficiaryFI
func (fibfi *FIBeneficiaryFI) String() string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fibfi.tag)
	buf.WriteString(fibfi.FIToFI.String(fibfi.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfi *FIBeneficiaryFI) Validate() error {
	if fibfi.tag != TagFIBeneficiaryFI {
		return fieldError("tag", ErrValidTagForType, fibfi.tag)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fibfi.FIToFI.LineOne)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfi.FIToFI.LineTwo)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fibfi.FIToFI.LineThree)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fibfi.FIToFI.LineFour)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fibfi.FIToFI.LineFive)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fibfi.FIToFI.LineSix)
	}
	return nil
}
