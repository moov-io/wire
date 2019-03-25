// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIBeneficiaryFIAdvice is the financial institution beneficiary financial institution
type FIBeneficiaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFIAdvice returns a new FIBeneficiaryFIAdvice
func NewFIBeneficiaryFIAdvice() FIBeneficiaryFIAdvice {
	ifia := FIBeneficiaryFIAdvice{
		tag: TagFIBeneficiaryFIAdvice,
	}
	return ifia
}

// Parse takes the input string and parses the FIBeneficiaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ifia *FIBeneficiaryFIAdvice) Parse(record string) {
}

// String writes FIBeneficiaryFIAdvice
func (ifia *FIBeneficiaryFIAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(194)
	buf.WriteString(ifia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ifia *FIBeneficiaryFIAdvice) Validate() error {
	if err := ifia.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ifia *FIBeneficiaryFIAdvice) fieldInclusion() error {
	return nil
}
