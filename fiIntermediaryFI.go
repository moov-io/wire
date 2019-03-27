// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIIntermediaryFI is the financial institution intermediary financial institution
type FIIntermediaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FiToFi `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFI returns a new FIIntermediaryFI
func NewFIIntermediaryFI() FIIntermediaryFI {
	ifi := FIIntermediaryFI{
		tag: TagFIIntermediaryFI,
	}
	return ifi
}

// Parse takes the input string and parses the FIIntermediaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ifi *FIIntermediaryFI) Parse(record string) {
}

// String writes FIIntermediaryFI
func (ifi *FIIntermediaryFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(195)
	buf.WriteString(ifi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ifi *FIIntermediaryFI) Validate() error {
	if err := ifi.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ifi *FIIntermediaryFI) fieldInclusion() error {
	return nil
}
