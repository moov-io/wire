// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// BeneficiaryReference is a reference for the beneficiary
type BeneficiaryReference struct {
	// tag
	tag string
	// BeneficiaryReference
	BeneficiaryReference string `json:"beneficiaryReference,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryReference returns a new BeneficiaryReference
func NewBeneficiaryReference() BeneficiaryReference  {
	br := BeneficiaryReference {
		tag: TagBeneficiaryReference,
	}
	return br
}

// Parse takes the input string and parses the BeneficiaryReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (br *BeneficiaryReference) Parse(record string) {
}

// String writes BeneficiaryReference
func (br *BeneficiaryReference) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(16)
	buf.WriteString(br.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryReference and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (br *BeneficiaryReference) Validate() error {
	if err := br.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (br *BeneficiaryReference) fieldInclusion() error {
	return nil
}

