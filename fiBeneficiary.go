// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIBeneficiary is the financial institution beneficiary
type FIBeneficiary struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FiToFi `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiary returns a new FIBeneficiary
func NewFIBeneficiary() FIBeneficiary {
	b := FIBeneficiary{
		tag: TagFIBeneficiary,
	}
	return b
}

// Parse takes the input string and parses the FIBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (b *FIBeneficiary) Parse(record string) {
}

// String writes FIBeneficiary
func (b *FIBeneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(195)
	buf.WriteString(b.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (b *FIBeneficiary) Validate() error {
	if err := b.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (b *FIBeneficiary) fieldInclusion() error {
	return nil
}
