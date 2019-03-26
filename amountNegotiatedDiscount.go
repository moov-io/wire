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
func NewAmountNegotiatedDiscount() AmountNegotiatedDiscount  {
	and := AmountNegotiatedDiscount {
		tag: TagAmountNegotiatedDiscount,
	}
	return and
}

// Parse takes the input string and parses the AmountNegotiatedDiscount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (and *AmountNegotiatedDiscount) Parse(record string) {
}

// String writes AmountNegotiatedDiscount
func (and *AmountNegotiatedDiscount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(22)
	buf.WriteString(and.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on AmountNegotiatedDiscount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (and *AmountNegotiatedDiscount) Validate() error {
	if err := and.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (and *AmountNegotiatedDiscount) fieldInclusion() error {
	return nil
}
