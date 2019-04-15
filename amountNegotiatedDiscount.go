// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (nd *AmountNegotiatedDiscount) Parse(record string) {
	nd.tag = record[:6]
	nd.RemittanceAmount.CurrencyCode = nd.parseStringField(record[6:9])
	nd.RemittanceAmount.Amount = nd.parseStringField(record[9:28])
}

// String writes AmountNegotiatedDiscount
func (nd *AmountNegotiatedDiscount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(28)
	buf.WriteString(nd.tag)
	buf.WriteString(nd.CurrencyCodeField())
	buf.WriteString(nd.AmountField())
	return buf.String()
}

// Validate performs WIRE format rule checks on AmountNegotiatedDiscount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (nd *AmountNegotiatedDiscount) Validate() error {
	if err := nd.fieldInclusion(); err != nil {
		return err
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
	return nil
}

// CurrencyCodeField gets a string of the CurrencyCode field
func (nd *AmountNegotiatedDiscount) CurrencyCodeField() string {
	return nd.alphaField(nd.RemittanceAmount.CurrencyCode, 3)
}

// AmountField gets a string of the Amount field
func (nd *AmountNegotiatedDiscount) AmountField() string {
	return nd.alphaField(nd.RemittanceAmount.Amount, 19)
}
