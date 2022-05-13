// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &GrossAmountRemittanceDocument{}

// GrossAmountRemittanceDocument is the gross amount remittance document
type GrossAmountRemittanceDocument struct {
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

// NewGrossAmountRemittanceDocument returns a new GrossAmountRemittanceDocument
func NewGrossAmountRemittanceDocument(isVariable bool) *GrossAmountRemittanceDocument {
	gard := &GrossAmountRemittanceDocument{
		tag:              TagGrossAmountRemittanceDocument,
		isVariableLength: isVariable,
	}
	return gard
}

// Parse takes the input string and parses the GrossAmountRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (gard *GrossAmountRemittanceDocument) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 8 {
		return NewTagWrongLengthErr(8, len(record)), 0
	}
	gard.tag = record[:6]

	return nil, 6 + gard.RemittanceAmount.Parse(record[6:])
}

func (gard *GrossAmountRemittanceDocument) UnmarshalJSON(data []byte) error {
	type Alias GrossAmountRemittanceDocument
	aux := struct {
		*Alias
	}{
		(*Alias)(gard),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	gard.tag = TagGrossAmountRemittanceDocument
	return nil
}

// String writes GrossAmountRemittanceDocument
func (gard *GrossAmountRemittanceDocument) String() string {
	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(gard.tag)
	buf.WriteString(gard.RemittanceAmount.String(gard.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on GrossAmountRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (gard *GrossAmountRemittanceDocument) Validate() error {
	if err := gard.fieldInclusion(); err != nil {
		return err
	}
	if gard.tag != TagGrossAmountRemittanceDocument {
		return fieldError("tag", ErrValidTagForType, gard.tag)
	}
	if err := gard.isCurrencyCode(gard.RemittanceAmount.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, gard.RemittanceAmount.CurrencyCode)
	}
	if err := gard.isAmount(gard.RemittanceAmount.Amount); err != nil {
		return fieldError("Amount", err, gard.RemittanceAmount.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (gard *GrossAmountRemittanceDocument) fieldInclusion() error {
	if gard.RemittanceAmount.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	if gard.RemittanceAmount.CurrencyCode == "" {
		return fieldError("CurrencyCode", ErrFieldRequired)
	}
	return nil
}
