// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// BeneficiaryIntermediaryFI {4000}
type BeneficiaryIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"FinancialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryIntermediaryFI returns a new BeneficiaryIntermediaryFI
func NewBeneficiaryIntermediaryFI() BeneficiaryIntermediaryFI {
	bifi := BeneficiaryIntermediaryFI{
		tag: TagBeneficiaryIntermediaryFI,
	}
	return bifi
}

// Parse takes the input string and parses the ReceiverDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bifi *BeneficiaryIntermediaryFI) Parse(record string) {
}

// String writes BeneficiaryIntermediaryFI
func (bifi *BeneficiaryIntermediaryFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(bifi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bifi *BeneficiaryIntermediaryFI) Validate() error {
	if err := bifi.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bifi *BeneficiaryIntermediaryFI) fieldInclusion() error {
	return nil
}
