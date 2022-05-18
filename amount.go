// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &Amount{}

// Amount (up to a penny less than $10 billion) {2000}
type Amount struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Amount 12 numeric, right-justified with leading zeros, an implied decimal point and no commas; e.g., $12,345.67 becomes 000001234567 Can be all zeros for subtype 90
	Amount string `json:"amount"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAmount returns a new Amount
func NewAmount(isVariable bool) *Amount {
	a := &Amount{
		tag:              TagAmount,
		isVariableLength: isVariable,
	}
	return a
}

// Parse takes the input string and parses the Amount value
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (a *Amount) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 7 {
		return 0, NewTagWrongLengthErr(7, len(record))
	}

	var err error
	var length, read int

	if a.tag, read, err = a.parseTag(record); err != nil {
		return 0, fieldError("Amount.Tag", err)
	}
	length += read

	if a.Amount, read, err = a.parseVariableStringField(record[length:], 12); err != nil {
		return 0, fieldError("Amount", err)
	}
	length += read

	return length, nil
}

func (a *Amount) UnmarshalJSON(data []byte) error {
	type Alias Amount
	aux := struct {
		*Alias
	}{
		(*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	a.tag = TagAmount
	return nil
}

// String writes Amount
func (a *Amount) String() string {
	var buf strings.Builder
	buf.Grow(18)
	buf.WriteString(a.tag)
	buf.WriteString(a.AmountField())
	return buf.String()
}

// Validate performs WIRE format rule checks on Amount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (a *Amount) Validate() error {
	if err := a.fieldInclusion(); err != nil {
		return err
	}
	if a.tag != TagAmount {
		return fieldError("tag", ErrValidTagForType, a.tag)
	}
	if err := a.isAmountImplied(a.Amount); err != nil {
		return fieldError("Amount", err, a.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (a *Amount) fieldInclusion() error {
	if a.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	return nil
}

// AmountField gets a string of entry addenda batch count zero padded
func (a *Amount) AmountField() string {
	return a.alphaVariableField(a.Amount, 12, a.isVariableLength)
}
