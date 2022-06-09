// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// Adjustment is adjustment
type Adjustment struct {
	// tag
	tag string
	// Adjustment  * `01` - Pricing Error * `03` - Extension Error * `04` - Item Not Accepted (Damaged) * `05` - Item Not Accepted (Quality) * `06` - Quantity Contested 07   Incorrect Product * `11` - Returns (Damaged) * `12` - Returns (Quality) * `59` - Item Not Received * `75` - Total Order Not Received * `81` - Credit as Agreed * `CM` - Covered by Credit Memo
	AdjustmentReasonCode string `json:"adjustmentReasonCode,omitempty"`
	// CreditDebitIndicator  * `CRDT` - Credit * `DBIT` - Debit
	CreditDebitIndicator string `json:"creditDebitIndicator,omitempty"`
	// RemittanceAmount is remittance amounts
	RemittanceAmount RemittanceAmount `json:"remittanceAmount,omitempty"`
	// AdditionalInfo is additional information
	AdditionalInfo string `json:"additionalInfo,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAdjustment returns a new Adjustment
func NewAdjustment() *Adjustment {
	adj := &Adjustment{
		tag: TagAdjustment,
	}
	return adj
}

// Parse takes the input string and parses the Adjustment values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (adj *Adjustment) Parse(record string) error {
	if utf8.RuneCountInString(record) < 10 {
		return NewTagMinLengthErr(10, len(record))
	}

	adj.tag = record[:6]
	length := 6

	value, read, err := adj.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("AdjustmentReasonCode", err)
	}
	adj.AdjustmentReasonCode = value
	length += read

	value, read, err = adj.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("CreditDebitIndicator", err)
	}
	adj.CreditDebitIndicator = value
	length += read

	value, read, err = adj.parseVariableStringField(record[length:], 3)
	if err != nil {
		return fieldError("CurrencyCode", err)
	}
	adj.RemittanceAmount.CurrencyCode = value
	length += read

	value, read, err = adj.parseVariableStringField(record[length:], 19)
	if err != nil {
		return fieldError("CurrencyCode", err)
	}
	adj.RemittanceAmount.Amount = value
	length += read

	value, read, err = adj.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("AdditionalInfo", err)
	}
	adj.AdditionalInfo = value
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (adj *Adjustment) UnmarshalJSON(data []byte) error {
	type Alias Adjustment
	aux := struct {
		*Alias
	}{
		(*Alias)(adj),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	adj.tag = TagAdjustment
	return nil
}

// String writes Adjustment
func (adj *Adjustment) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(168)

	buf.WriteString(adj.tag)
	buf.WriteString(adj.AdjustmentReasonCodeField(options...))
	buf.WriteString(adj.CreditDebitIndicatorField(options...))
	buf.WriteString(adj.CurrencyCodeField(options...))
	buf.WriteString(adj.AmountField(options...))
	buf.WriteString(adj.AdditionalInfoField(options...))

	if adj.parseFirstOption(options) {
		return adj.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on Adjustment and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// Adjustment Reason, Credit Debit Indicator, Currency Code and Amount are mandatory.
func (adj *Adjustment) Validate() error {
	if err := adj.fieldInclusion(); err != nil {
		return err
	}
	if adj.tag != TagAdjustment {
		return fieldError("tag", ErrValidTagForType, adj.tag)
	}
	if err := adj.isAdjustmentReasonCode(adj.AdjustmentReasonCode); err != nil {
		return fieldError("AdjustmentReasonCode", err, adj.AdjustmentReasonCode)
	}
	if err := adj.isCreditDebitIndicator(adj.CreditDebitIndicator); err != nil {
		return fieldError("CreditDebitIndicator", err, adj.CreditDebitIndicator)
	}
	if err := adj.isCurrencyCode(adj.RemittanceAmount.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, adj.RemittanceAmount.CurrencyCode)
	}
	if err := adj.isAmount(adj.RemittanceAmount.Amount); err != nil {
		return fieldError("Amount", err, adj.RemittanceAmount.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (adj *Adjustment) fieldInclusion() error {
	if adj.AdjustmentReasonCode == "" {
		return fieldError("AdjustmentReasonCode", ErrFieldRequired)
	}
	if adj.CreditDebitIndicator == "" {
		return fieldError("CreditDebitIndicator", ErrFieldRequired)
	}
	if adj.RemittanceAmount.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	if adj.RemittanceAmount.CurrencyCode == "" {
		return fieldError("CurrencyCode", ErrFieldRequired)
	}
	return nil
}

// AdjustmentReasonCodeField gets a string of the AdjustmentReasonCode field
func (adj *Adjustment) AdjustmentReasonCodeField(options ...bool) string {
	return adj.alphaVariableField(adj.AdjustmentReasonCode, 2, adj.parseFirstOption(options))
}

// CreditDebitIndicatorField gets a string of the CreditDebitIndicator field
func (adj *Adjustment) CreditDebitIndicatorField(options ...bool) string {
	return adj.alphaVariableField(adj.CreditDebitIndicator, 4, adj.parseFirstOption(options))
}

// CurrencyCodeField gets a string of the CurrencyCode field
func (adj *Adjustment) CurrencyCodeField(options ...bool) string {
	return adj.alphaVariableField(adj.RemittanceAmount.CurrencyCode, 3, adj.parseFirstOption(options))
}

// AmountField gets a string of the Amount field
func (adj *Adjustment) AmountField(options ...bool) string {
	return adj.alphaVariableField(adj.RemittanceAmount.Amount, 19, adj.parseFirstOption(options))
}

// AdditionalInfoField gets a string of the AdditionalInfo field
func (adj *Adjustment) AdditionalInfoField(options ...bool) string {
	return adj.alphaVariableField(adj.AdditionalInfo, 140, adj.parseFirstOption(options))
}
