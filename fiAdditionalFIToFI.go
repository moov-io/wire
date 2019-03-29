// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIAdditionalFIToFI is the financial institution beneficiary financial institution
type FIAdditionalFIToFI struct {
	// tag
	tag string
	// AdditionalFiToFi is additional financial institution to financial institution information
	AdditionalFiToFi AdditionalFiToFi `json:"additionalFiToFi,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIAdditionalFIToFI returns a new FIAdditionalFIToFI
func NewFIAdditionalFIToFI() FIAdditionalFIToFI {
	additional := FIAdditionalFIToFI{
		tag: TagFIAdditionalFIToFI,
	}
	return additional
}

// Parse takes the input string and parses the FIAdditionalFIToFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (additional *FIAdditionalFIToFI) Parse(record string) {
}

// String writes FIAdditionalFIToFI
func (additional *FIAdditionalFIToFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(195)
	buf.WriteString(additional.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIAdditionalFIToFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (additional *FIAdditionalFIToFI) Validate() error {
	if err := additional.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (additional *FIAdditionalFIToFI) fieldInclusion() error {
	return nil
}
