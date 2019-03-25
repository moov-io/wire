// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// OriginatorOptionF is originator option F information
type OriginatorOptionF struct {
	// tag
	tag string
	// PartyIdentifier  Must be one of the following two formats: 1. /Account Number (slash followed by at least one
	// valid non-space character:  e.g., /123456)  2. Unique Identifier/ (4 character code followed by a slash and at
	// least one valid non-space character:
	// e.g., SOSE/123-456-789)
	// ARNU: Alien Registration Number
	// CCPT: Passport Number
	// CUST: Customer Identification Number
	// DRLC: Driver’s License Number
	// EMPL: Employer Number NIDN: National Identify Number
	// SOSE: Social Security Number TXID: Tax Identification Number
	PartyIdentifier string `json:"partyIdentifier,omitempty"`
	// Name  Format:  Must begin with Line Code 1 followed by a slash and at least one valid non-space character:
	// e.g., 1/SMITH JOHN.
	Name string `json:"name,omitempty"`
	// LineOne
	// Format: Must begin with one of the following Line Codes followed by a slash and at least one
	// valid non-space character.
	// 1 Name
	// 2 Address
	// 3 Country and Town
	// 4 Date of Birth
	// 5 Place of Birth
	// 6 Customer Identification Number
	// 7 National Identity Number
	// 8 Additional Information
	// For example:
	// 2/123 MAIN STREET
	// 3/US/NEW YORK, NY 10000
	// 7/111-22-3456
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	// Format: Must begin with one of the following Line Codes followed by a slash and at least one
	// valid non-space character.
	// 1 Name
	// 2 Address
	// 3 Country and Town
	// 4 Date of Birth
	// 5 Place of Birth
	// 6 Customer Identification Number
	// 7 National Identity Number
	// 8 Additional Information
	// For example:
	// 2/123 MAIN STREET
	// 3/US/NEW YORK, NY 10000
	// 7/111-22-3456
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	// Format: Must begin with one of the following Line Codes followed by a slash and at least one
	// valid non-space character.
	// 1 Name
	// 2 Address
	// 3 Country and Town
	// 4 Date of Birth
	// 5 Place of Birth
	// 6 Customer Identification Number
	// 7 National Identity Number
	// 8 Additional Information
	// For example:
	// 2/123 MAIN STREET
	// 3/US/NEW YORK, NY 10000
	// 7/111-22-3456
	LineThree string `json:"lineThree,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginatorOptionF returns a new OriginatorOptionF
func NewOriginatorOptionF() OriginatorOptionF  {
	oof := OriginatorOptionF {
		tag: TagOriginatorOptionF,
	}
	return oof
}

// Parse takes the input string and parses the OriginatorOptionF values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oof *OriginatorOptionF) Parse(record string) {
}

// String writes OriginatorOptionF
func (oof *OriginatorOptionF) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(oof.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OriginatorOptionF and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oof *OriginatorOptionF) Validate() error {
	if err := oof.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oof *OriginatorOptionF) fieldInclusion() error {
	return nil
}


