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
	// RemittanceAmount is remittance amounts
	RemittanceAmount RemittanceAmount `json:"remittanceAmount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewGrossAmountRemittanceDocument returns a new GrossAmountRemittanceDocument
func NewGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	gard := &GrossAmountRemittanceDocument{
		tag: TagGrossAmountRemittanceDocument,
	}
	return gard
}

// Parse takes the input string and parses the GrossAmountRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (gard *GrossAmountRemittanceDocument) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, len(record))
	}

	var err error
	var length, read int

	if gard.tag, read, err = gard.parseTag(record); err != nil {
		return 0, fieldError("GrossAmountRemittanceDocument.Tag", err)
	}
	length += read

	if read, err = gard.RemittanceAmount.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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
func (gard *GrossAmountRemittanceDocument) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(gard.tag)
	buf.WriteString(gard.RemittanceAmount.String(isCompressed))

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
