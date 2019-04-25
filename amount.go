// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// Amount (up to a penny less than $10 billion) {2000}
type Amount struct {
	// tag
	tag string
	// Amount 12 numeric, right-justified with leading zeros, an implied decimal point and no commas; e.g., $12,345.67 becomes 000001234567 Can be all zeros for subtype 90
	Amount string `json:"amount"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAmount returns a new Amount
func NewAmount() *Amount {
	a := &Amount{
		tag: TagAmount,
	}
	return a
}

// Parse takes the input string and parses the Amount value
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (a *Amount) Parse(record string) {
	a.tag = record[:6]
	a.Amount = a.parseStringField(record[6:18])
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
	// ToDo:  Amount cannot be all zero's except if SubType = 90
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
	return a.numericStringField(a.Amount, 12)
}
