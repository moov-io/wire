// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &Adjustment{}

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
func (adj *Adjustment) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 11 {
		return 0, NewTagWrongLengthErr(11, len(record))
	}

	var err error
	var length, read int

	if adj.tag, read, err = adj.parseTag(record); err != nil {
		return 0, fieldError("Adjustment.Tag", err)
	}
	length += read

	if adj.AdjustmentReasonCode, read, err = adj.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("AdjustmentReasonCode", err)
	}
	length += read

	if adj.CreditDebitIndicator, read, err = adj.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("CreditDebitIndicator", err)
	}
	length += read

	if read, err = adj.RemittanceAmount.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	if adj.AdditionalInfo, read, err = adj.parseVariableStringField(record[length:], 140); err != nil {
		return 0, fieldError("AdditionalInfo", err)
	}
	length += read

	return length, nil
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

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(168)

	buf.WriteString(adj.tag)
	buf.WriteString(adj.AdjustmentReasonCodeField(isCompressed))
	buf.WriteString(adj.CreditDebitIndicatorField(isCompressed))
	buf.WriteString(adj.RemittanceAmount.String(isCompressed))
	buf.WriteString(adj.AdditionalInfoField(isCompressed))

	return buf.String()
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
func (adj *Adjustment) AdjustmentReasonCodeField(isCompressed bool) string {
	return adj.alphaVariableField(adj.AdjustmentReasonCode, 2, isCompressed)
}

// CreditDebitIndicatorField gets a string of the CreditDebitIndicator field
func (adj *Adjustment) CreditDebitIndicatorField(isCompressed bool) string {
	return adj.alphaVariableField(adj.CreditDebitIndicator, 4, isCompressed)
}

// AdditionalInfoField gets a string of the AdditionalInfo field
func (adj *Adjustment) AdditionalInfoField(isCompressed bool) string {
	return adj.alphaVariableField(adj.AdditionalInfo, 140, isCompressed)
}
