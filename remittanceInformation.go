// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// RemittanceInformation is the remittance information
type RemittanceInformation struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceInformation returns a new RemittanceInformation
func NewRemittanceInformation() RemittanceInformation {
	ri := RemittanceInformation{
		tag: TagRemittanceInformation,
	}
	return ri
}

// Parse takes the input string and parses the RemittanceInformation values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ri *RemittanceInformation) Parse(record string) {
}

// String writes RemittanceInformation
func (ri *RemittanceInformation) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(145)
	buf.WriteString(ri.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceInformation and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ri *RemittanceInformation) Validate() error {
	if err := ri.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ri *RemittanceInformation) fieldInclusion() error {
	return nil
}
