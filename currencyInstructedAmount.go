// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// CurrencyInstructedAmount is the currency instructed amount
type CurrencyInstructedAmount struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewCurrencyInstructedAmount returns a new CurrencyInstructedAmount
func NewCurrencyInstructedAmount() CurrencyInstructedAmount  {
	cia := CurrencyInstructedAmount {
		tag: TagCurrencyInstructedAmount,
	}
	return cia
}

// Parse takes the input string and parses the CurrencyInstructedAmount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (cia *CurrencyInstructedAmount) Parse(record string) {
}

// String writes CurrencyInstructedAmount
func (cia *CurrencyInstructedAmount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(23)
	buf.WriteString(cia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on CurrencyInstructedAmount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (cia *CurrencyInstructedAmount) Validate() error {
	if err := cia.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (cia *CurrencyInstructedAmount) fieldInclusion() error {
	return nil
}
