// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// OriginatorFI is the originator Financial Institution
type OriginatorFI struct {
	// tag
	tag string
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginatorFI returns a new OriginatorFI
func NewOriginatorFI() OriginatorFI  {
	ofi := OriginatorFI {
		tag: TagOriginatorFI,
	}
	return ofi
}

// Parse takes the input string and parses the OriginatorFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ofi *OriginatorFI) Parse(record string) {
}

// String writes OriginatorFI
func (ofi *OriginatorFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(ofi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OriginatorFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ofi *OriginatorFI) Validate() error {
	if err := ofi.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ofi *OriginatorFI) fieldInclusion() error {
	return nil
}
