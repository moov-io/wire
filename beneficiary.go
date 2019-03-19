// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

type Beneficiary struct {
	// tag
	tag string
	// Personal
	Personal Personal `json:"personal,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiary returns a new Beneficiary
func NewBeneficiary() Beneficiary  {
	b := Beneficiary {
		tag: TagBeneficiary,
	}
	return b
}

// Parse takes the input string and parses the Beneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (b *Beneficiary) Parse(record string) {
}

// String writes Beneficiary
func (b *Beneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(b.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiverDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (b *Beneficiary) Validate() error {
	if err := b.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (b *Beneficiary) fieldInclusion() error {
	return nil
}
