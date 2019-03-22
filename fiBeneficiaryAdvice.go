// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIBeneficiaryAdvice is the financial institution beneficiary advice
type FIBeneficiaryAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryAdvice returns a new FIBeneficiaryAdvice
func NewFIBeneficiaryAdvice() FIBeneficiaryAdvice {
	ba := FIBeneficiaryAdvice{
		tag: TagFIBeneficiaryAdvice,
	}
	return ba
}

// Parse takes the input string and parses the FIBeneficiaryAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ba *FIBeneficiaryAdvice) Parse(record string) {
}

// String writes FIBeneficiaryAdvice
func (ba *FIBeneficiaryAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(194)
	buf.WriteString(ba.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ba *FIBeneficiaryAdvice) Validate() error {
	if err := ba.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ba *FIBeneficiaryAdvice) fieldInclusion() error {
	return nil
}

