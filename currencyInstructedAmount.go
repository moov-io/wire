// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// CurrencyInstructedAmount is the currency instructed amount
type CurrencyInstructedAmount struct {
	// tag
	tag string
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag"`
	// Amount is the instructed amount
	// Amount  Must begin with at least one numeric character (0-9) and contain only one decimal comma marker
	// (e.g., $1,234.56 should be entered as 1234,56 and $0.99 should be entered as
	Amount string `json:"amount"`
	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewCurrencyInstructedAmount returns a new CurrencyInstructedAmount
func NewCurrencyInstructedAmount() CurrencyInstructedAmount {
	cia := CurrencyInstructedAmount{
		tag: TagCurrencyInstructedAmount,
	}
	return cia
}

// Parse takes the input string and parses the CurrencyInstructedAmount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (cia *CurrencyInstructedAmount) Parse(record string) {
	cia.tag = record[:6]
	cia.SwiftFieldTag = cia.parseStringField(record[6:11])
	cia.Amount = cia.parseStringField(record[11:29])
}

// String writes CurrencyInstructedAmount
func (cia *CurrencyInstructedAmount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(29)
	buf.WriteString(cia.tag)
	buf.WriteString(cia.SwiftFieldTagField())
	buf.WriteString(cia.AmountField())
	return buf.String()
}

// Validate performs WIRE format rule checks on CurrencyInstructedAmount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (cia *CurrencyInstructedAmount) Validate() error {
	if err := cia.fieldInclusion(); err != nil {
		return err
	}
	if err := cia.isAlphanumeric(cia.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, cia.SwiftFieldTag)
	}
	if err := cia.isAmount(cia.Amount); err != nil {
		return fieldError("Amount", err, cia.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (cia *CurrencyInstructedAmount) fieldInclusion() error {
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (cia *CurrencyInstructedAmount) SwiftFieldTagField() string {
	return cia.alphaField(cia.SwiftFieldTag, 5)
}

// AmountField gets a string of the AmountTag field
func (cia *CurrencyInstructedAmount) AmountField() string {
	return cia.alphaField(cia.Amount, 18)
}
