// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// BeneficiaryCustomer is the beneficiary customer
type BeneficiaryCustomer struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryCustomer returns a new BeneficiaryCustomer
func NewBeneficiaryCustomer() BeneficiaryCustomer  {
	bc := BeneficiaryCustomer {
		tag: TagBeneficiaryCustomer,
	}
	return bc
}

// Parse takes the input string and parses the BeneficiaryCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bc *BeneficiaryCustomer) Parse(record string) {
}

// String writes BeneficiaryCustomer
func (bc *BeneficiaryCustomer) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(180)
	buf.WriteString(bc.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BeneficiaryCustomer) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bc *BeneficiaryCustomer) fieldInclusion() error {
	return nil
}
