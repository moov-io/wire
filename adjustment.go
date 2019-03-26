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
func NewAdjustment() Adjustment  {
	adj := Adjustment {
		tag: TagAdjustment,
	}
	return adj
}

// Parse takes the input string and parses the Adjustment values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (adj *Adjustment) Parse(record string) {
}

// String writes Adjustment
func (adj *Adjustment) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(168)
	buf.WriteString(adj.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on Adjustment and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (adj *Adjustment) Validate() error {
	if err := adj.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (adj *Adjustment) fieldInclusion() error {
	return nil
}
