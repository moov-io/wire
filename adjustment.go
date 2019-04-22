// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (adj *Adjustment) Parse(record string) {
	adj.tag = record[:6]
	adj.AdjustmentReasonCode = adj.parseStringField(record[6:8])
	adj.CreditDebitIndicator = adj.parseStringField(record[8:12])
	adj.RemittanceAmount.CurrencyCode = adj.parseStringField(record[12:15])
	adj.RemittanceAmount.Amount = adj.parseStringField(record[15:34])
	adj.AdditionalInfo = adj.parseStringField(record[34:174])
}

// String writes Adjustment
func (adj *Adjustment) String() string {
	var buf strings.Builder
	buf.Grow(168)
	buf.WriteString(adj.tag)
	buf.WriteString(adj.AdjustmentReasonCodeField())
	buf.WriteString(adj.CreditDebitIndicatorField())
	buf.WriteString(adj.CurrencyCodeField())
	buf.WriteString(adj.AmountField())
	buf.WriteString(adj.AdditionalInfoField())
	return buf.String()
}

// Validate performs WIRE format rule checks on Adjustment and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (adj *Adjustment) Validate() error {
	if err := adj.fieldInclusion(); err != nil {
		return err
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
	return nil
}

// AdjustmentReasonCodeField gets a string of the AdjustmentReasonCode field
func (adj *Adjustment) AdjustmentReasonCodeField() string {
	return adj.alphaField(adj.AdjustmentReasonCode, 2)
}

// CreditDebitIndicatorField gets a string of the CreditDebitIndicator field
func (adj *Adjustment) CreditDebitIndicatorField() string {
	return adj.alphaField(adj.CreditDebitIndicator, 4)
}

// CurrencyCodeField gets a string of the CurrencyCode field
func (adj *Adjustment) CurrencyCodeField() string {
	return adj.alphaField(adj.RemittanceAmount.CurrencyCode, 3)
}

// AmountField gets a string of the Amount field
func (adj *Adjustment) AmountField() string {
	return adj.alphaField(adj.RemittanceAmount.Amount, 19)
}

// AdditionalInfoField gets a string of the AdditionalInfo field
func (adj *Adjustment) AdditionalInfoField() string {
	return adj.alphaField(adj.AdditionalInfo, 140)
}
