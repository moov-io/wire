// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIDrawdownDebitAccountAdvice is the financial institution drawdown debit account advice
type FIDrawdownDebitAccountAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIDrawdownDebitAccountAdvice returns a new FIDrawdownDebitAccountAdvice
func NewFIDrawdownDebitAccountAdvice() FIDrawdownDebitAccountAdvice {
	ba := FIDrawdownDebitAccountAdvice{
		tag: TagFIDrawdownDebitAccountAdvice,
	}
	return ba
}

// Parse takes the input string and parses the FIDrawdownDebitAccountAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ba *FIDrawdownDebitAccountAdvice) Parse(record string) {
}

// String writes FIDrawdownDebitAccountAdvice
func (ba *FIDrawdownDebitAccountAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(194)
	buf.WriteString(ba.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ba *FIDrawdownDebitAccountAdvice) Validate() error {
	if err := ba.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ba *FIDrawdownDebitAccountAdvice) fieldInclusion() error {
	return nil
}
