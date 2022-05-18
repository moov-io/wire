// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &OriginatorOptionF{}

// OriginatorOptionF is originator option F information
type OriginatorOptionF struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// PartyIdentifier must be one of the following two formats:
	// 1. /Account Number (slash followed by at least one
	// valid non-space character:  e.g., /123456)
	// 2. Unique Identifier/ (4 character code followed by a slash and at least one valid non-space character:
	// e.g., SOSE/123-456-789)
	//
	// ARNU: Alien Registration Number
	// CCPT: Passport Number
	// CUST: Customer Identification Number
	// DRLC: Driver’s License Number
	// EMPL: Employer Number
	// NIDN: National Identify Number
	// SOSE: Social Security Number
	// TXID: Tax Identification Number
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
func NewOriginatorOptionF(isVariable bool) *OriginatorOptionF {
	oof := &OriginatorOptionF{
		tag:              TagOriginatorOptionF,
		isVariableLength: isVariable,
	}
	return oof
}

// Parse takes the input string and parses the OriginatorOptionF values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oof *OriginatorOptionF) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 11 {
		return 0, NewTagWrongLengthErr(11, len(record))
	}

	var err error
	var length, read int

	if oof.tag, read, err = oof.parseTag(record); err != nil {
		return 0, fieldError("OriginatorOptionF.Tag", err)
	}
	length += read

	if oof.PartyIdentifier, read, err = oof.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("PartyIdentifier", err)
	}
	length += read

	if oof.Name, read, err = oof.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("Name", err)
	}
	length += read

	if oof.LineOne, read, err = oof.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineOne", err)
	}
	length += read

	if oof.LineTwo, read, err = oof.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineTwo", err)
	}
	length += read

	if oof.LineThree, read, err = oof.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineThree", err)
	}
	length += read

	return length, nil
}

func (oof *OriginatorOptionF) UnmarshalJSON(data []byte) error {
	type Alias OriginatorOptionF
	aux := struct {
		*Alias
	}{
		(*Alias)(oof),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	oof.tag = TagOriginatorOptionF
	return nil
}

// String writes OriginatorOptionF
func (oof *OriginatorOptionF) String() string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(oof.tag)
	buf.WriteString(oof.PartyIdentifierField())
	buf.WriteString(oof.NameField())
	buf.WriteString(oof.LineOneField())
	buf.WriteString(oof.LineTwoField())
	buf.WriteString(oof.LineThreeField())

	return buf.String()
}

// Validate performs WIRE format rule checks on OriginatorOptionF and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oof *OriginatorOptionF) Validate() error {
	if err := oof.fieldInclusion(); err != nil {
		return err
	}
	if err := oof.validatePartyIdentifier(oof.PartyIdentifier); err != nil {
		return fieldError("PartyIdentifier", err, oof.PartyIdentifier)
	}
	if err := oof.validateOptionFName(oof.Name); err != nil {
		return fieldError("Name", err, oof.Name)
	}
	if err := oof.validateOptionFLine(oof.LineOne); err != nil {
		return fieldError("LineOne", err, oof.LineOne)
	}
	if err := oof.validateOptionFLine(oof.LineTwo); err != nil {
		return fieldError("LineTwo", err, oof.LineTwo)
	}
	if err := oof.validateOptionFLine(oof.LineThree); err != nil {
		return fieldError("LineThree", err, oof.LineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oof *OriginatorOptionF) fieldInclusion() error {
	return nil
}

// PartyIdentifierField gets a string of the PartyIdentifier field
func (oof *OriginatorOptionF) PartyIdentifierField() string {
	return oof.alphaVariableField(oof.PartyIdentifier, 35, oof.isVariableLength)
}

// NameField gets a string of the Name field
func (oof *OriginatorOptionF) NameField() string {
	return oof.alphaVariableField(oof.Name, 35, oof.isVariableLength)
}

// LineOneField gets a string of the LineOne field
func (oof *OriginatorOptionF) LineOneField() string {
	return oof.alphaVariableField(oof.LineOne, 35, oof.isVariableLength)
}

// LineTwoField gets a string of the LineTwo field
func (oof *OriginatorOptionF) LineTwoField() string {
	return oof.alphaVariableField(oof.LineTwo, 35, oof.isVariableLength)
}

// LineThreeField gets a string of the LineThree field
func (oof *OriginatorOptionF) LineThreeField() string {
	return oof.alphaVariableField(oof.LineThree, 35, oof.isVariableLength)
}
