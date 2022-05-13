// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// RemittanceAmount is remittance amount
type RemittanceAmount struct {
	// CurrencyCode
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Amount Must contain at least one numeric character and only one decimal period marker (e.g., $1,234.56 should be entered as 1234.56). Can have up to 5 numeric characters following the decimal period marker (e.g., 1234.56789). Amount must be greater than zero (i.e., at least .01).
	Amount string `json:"amount,omitempty"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the RemittanceAmount values
func (r *RemittanceAmount) Parse(record string) int {

	length := 0
	read := 0

	r.CurrencyCode, read = r.parseVariableStringField(record[length:], 3)
	length += read

	r.Amount, read = r.parseVariableStringField(record[length:], 19)
	length += read

	return length
}

// String writes RemittanceAmount
func (r *RemittanceAmount) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(22)

	buf.WriteString(r.CurrencyCodeField(isVariable))
	buf.WriteString(r.AmountField(isVariable))

	return buf.String()
}

// CurrencyCodeField gets a string of the CurrencyCode field
func (r *RemittanceAmount) CurrencyCodeField(isVariable bool) string {
	return r.alphaVariableField(r.CurrencyCode, 3, isVariable)
}

// AmountField gets a string of the Amount field
func (r *RemittanceAmount) AmountField(isVariable bool) string {
	return r.alphaVariableField(r.Amount, 19, isVariable)
}
