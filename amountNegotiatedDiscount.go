// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// AmountNegotiatedDiscount is the amount negotiated discount
type AmountNegotiatedDiscount struct {
	// tag
	tag string
	// RemittanceAmount is remittance amounts
	RemittanceAmount RemittanceAmount `json:"remittanceAmount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAmountNegotiatedDiscount returns a new AmountNegotiatedDiscount
func NewAmountNegotiatedDiscount() *AmountNegotiatedDiscount {
	nd := &AmountNegotiatedDiscount{
		tag: TagAmountNegotiatedDiscount,
	}
	return nd
}

// Parse takes the input string and parses the AmountNegotiatedDiscount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (nd *AmountNegotiatedDiscount) Parse(record string) error {
	if utf8.RuneCountInString(record) < 8 {
		return NewTagMinLengthErr(8, len(record))
	}

	nd.tag = record[:6]
	length := 6

	value, read, err := nd.parseVariableStringField(record[length:], 3)
	if err != nil {
		return fieldError("CurrencyCode", err)
	}
	nd.RemittanceAmount.CurrencyCode = value
	length += read

	value, read, err = nd.parseVariableStringField(record[length:], 19)
	if err != nil {
		return fieldError("Amount", err)
	}
	nd.RemittanceAmount.Amount = value
	length += read

	if !nd.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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
func (nd *AmountNegotiatedDiscount) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(nd.tag)
	buf.WriteString(nd.CurrencyCodeField(options...))
	buf.WriteString(nd.AmountField(options...))

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

// CurrencyCodeField gets a string of the CurrencyCode field
func (nd *AmountNegotiatedDiscount) CurrencyCodeField(options ...bool) string {
	return nd.alphaVariableField(nd.RemittanceAmount.CurrencyCode, 3, nd.parseFirstOption(options))
}

// AmountField gets a string of the Amount field
func (nd *AmountNegotiatedDiscount) AmountField(options ...bool) string {
	return nd.alphaVariableField(nd.RemittanceAmount.Amount, 19, nd.parseFirstOption(options))
}
