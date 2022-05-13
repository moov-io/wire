// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &AmountNegotiatedDiscount{}

// AmountNegotiatedDiscount is the amount negotiated discount
type AmountNegotiatedDiscount struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// RemittanceAmount is remittance amounts
	RemittanceAmount RemittanceAmount `json:"remittanceAmount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAmountNegotiatedDiscount returns a new AmountNegotiatedDiscount
func NewAmountNegotiatedDiscount(isVariable bool) *AmountNegotiatedDiscount {
	nd := &AmountNegotiatedDiscount{
		tag:              TagAmountNegotiatedDiscount,
		isVariableLength: isVariable,
	}
	return nd
}

// Parse takes the input string and parses the AmountNegotiatedDiscount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (nd *AmountNegotiatedDiscount) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 8 {
		return NewTagWrongLengthErr(8, len(record)), 0
	}
	nd.tag = record[:6]

	return nil, 6 + nd.RemittanceAmount.Parse(record[6:])
}

func (nd *AmountNegotiatedDiscount) UnmarshalJSON(data []byte) error {
	type Alias AmountNegotiatedDiscount
	aux := struct {
		*Alias
	}{
		(*Alias)(nd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	nd.tag = TagAmountNegotiatedDiscount
	return nil
}

// String writes AmountNegotiatedDiscount
func (nd *AmountNegotiatedDiscount) String() string {
	var buf strings.Builder
	buf.Grow(28)
	buf.WriteString(nd.tag)
	buf.WriteString(nd.RemittanceAmount.String(nd.isVariableLength))
	return buf.String()
}

// Validate performs WIRE format rule checks on AmountNegotiatedDiscount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (nd *AmountNegotiatedDiscount) Validate() error {
	if err := nd.fieldInclusion(); err != nil {
		return err
	}
	if nd.tag != TagAmountNegotiatedDiscount {
		return fieldError("tag", ErrValidTagForType, nd.tag)
	}
	if err := nd.isCurrencyCode(nd.RemittanceAmount.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, nd.RemittanceAmount.CurrencyCode)
	}
	if err := nd.isAmount(nd.RemittanceAmount.Amount); err != nil {
		return fieldError("Amount", err, nd.RemittanceAmount.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (nd *AmountNegotiatedDiscount) fieldInclusion() error {
	if nd.RemittanceAmount.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	if nd.RemittanceAmount.CurrencyCode == "" {
		return fieldError("CurrencyCode", ErrFieldRequired)
	}
	return nil
}
