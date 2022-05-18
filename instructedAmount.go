// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &InstructedAmount{}

// InstructedAmount is the InstructedAmount of the wire
type InstructedAmount struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// CurrencyCode
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Amount  Must begin with at least one numeric character (0-9) and contain only one decimal comma marker
	// (e.g., $1,234.56 should be entered as 1234,56 and $0.99 should be entered as
	Amount string `json:"amount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInstructedAmount returns a new InstructedAmount
func NewInstructedAmount(isVariable bool) *InstructedAmount {
	ia := &InstructedAmount{
		tag:              TagInstructedAmount,
		isVariableLength: isVariable,
	}
	return ia
}

// Parse takes the input string and parses the InstructedAmount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ia *InstructedAmount) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, len(record))
	}

	var err error
	var length, read int

	if ia.tag, read, err = ia.parseTag(record); err != nil {
		return 0, fieldError("InstructedAmount.Tag", err)
	}
	length += read

	if ia.CurrencyCode, read, err = ia.parseVariableStringField(record[length:], 3); err != nil {
		return 0, fieldError("CurrencyCode", err)
	}
	length += read

	if ia.Amount, read, err = ia.parseVariableStringField(record[length:], 15); err != nil {
		return 0, fieldError("Amount", err)
	}
	length += read

	return length, nil
}

func (ia *InstructedAmount) UnmarshalJSON(data []byte) error {
	type Alias InstructedAmount
	aux := struct {
		*Alias
	}{
		(*Alias)(ia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ia.tag = TagInstructedAmount
	return nil
}

// String writes InstructedAmount
func (ia *InstructedAmount) String() string {
	var buf strings.Builder
	buf.Grow(24)

	buf.WriteString(ia.tag)
	buf.WriteString(ia.CurrencyCodeField())
	buf.WriteString(ia.AmountField())

	return buf.String()
}

// Validate performs WIRE format rule checks on InstructedAmount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ia *InstructedAmount) Validate() error {
	if err := ia.fieldInclusion(); err != nil {
		return err
	}
	if ia.tag != TagInstructedAmount {
		return fieldError("tag", ErrValidTagForType, ia.tag)
	}
	if err := ia.isCurrencyCode(ia.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, ia.CurrencyCode)
	}
	if err := ia.isAmount(ia.Amount); err != nil {
		return fieldError("Amount", err, ia.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ia *InstructedAmount) fieldInclusion() error {
	if ia.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	if ia.CurrencyCode == "" {
		return fieldError("CurrencyCode", ErrFieldRequired)

	}
	return nil
}

// CurrencyCodeField gets a string of the CurrencyCode field
func (ia *InstructedAmount) CurrencyCodeField() string {
	return ia.alphaVariableField(ia.CurrencyCode, 3, ia.isVariableLength)
}

// AmountField gets a string of the Amount field
func (ia *InstructedAmount) AmountField() string {
	return ia.alphaVariableField(ia.Amount, 15, ia.isVariableLength)
}
