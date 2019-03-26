// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func NewGrossAmountRemittanceDocument() GrossAmountRemittanceDocument  {
	gard := GrossAmountRemittanceDocument {
		tag: TagGrossAmountRemittanceDocument,
	}
	return gard
}

// Parse takes the input string and parses the GrossAmountRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (gard *GrossAmountRemittanceDocument) Parse(record string) {
}

// String writes GrossAmountRemittanceDocument
func (gard *GrossAmountRemittanceDocument) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(22)
	buf.WriteString(gard.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on GrossAmountRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (gard *GrossAmountRemittanceDocument) Validate() error {
	if err := gard.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (gard *GrossAmountRemittanceDocument) fieldInclusion() error {
	return nil
}
