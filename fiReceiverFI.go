// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIReceiverFI is the financial institution receiver financial institution
type FIReceiverFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FiToFi `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIReceiverFI returns a new FIReceiverFI
func NewFIReceiverFI() FIReceiverFI {
	rfi := FIReceiverFI{
		tag: TagFIReceiverFI,
	}
	return rfi
}

// Parse takes the input string and parses the FIReceiverFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rfi *FIReceiverFI) Parse(record string) {
}

// String writes FIReceiverFI
func (rfi *FIReceiverFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(195)
	buf.WriteString(rfi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIReceiverFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rfi *FIReceiverFI) Validate() error {
	if err := rfi.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rfi *FIReceiverFI) fieldInclusion() error {
	return nil
}
