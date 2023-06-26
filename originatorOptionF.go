// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// OriginatorOptionF is originator option F information
type OriginatorOptionF struct {
	// tag
	tag string
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
func NewOriginatorOptionF() *OriginatorOptionF {
	oof := &OriginatorOptionF{
		tag: TagOriginatorOptionF,
	}
	return oof
}

// Parse takes the input string and parses the OriginatorOptionF values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oof *OriginatorOptionF) Parse(record string) error {
	if utf8.RuneCountInString(record) < 13 {
		return NewTagMinLengthErr(13, len(record))
	}

	oof.tag = oof.parseStringField(record[:6])
	length := 6

	value, read, err := oof.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("PartyIdentifier", err)
	}
	oof.PartyIdentifier = value
	length += read

	value, read, err = oof.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Name", err)
	}
	oof.Name = value
	length += read

	value, read, err = oof.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineOne", err)
	}
	oof.LineOne = value
	length += read

	value, read, err = oof.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	oof.LineTwo = value
	length += read

	value, read, err = oof.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineThree", err)
	}
	oof.LineThree = value
	length += read

	if err := oof.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width OriginatorOptionF record
func (oof *OriginatorOptionF) String() string {
	return oof.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a OriginatorOptionF record formatted according to the FormatOptions
func (oof *OriginatorOptionF) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(oof.tag)
	buf.WriteString(oof.FormatPartyIdentifier(options) + Delimiter)
	buf.WriteString(oof.FormatName(options) + Delimiter)
	buf.WriteString(oof.FormatLineOne(options) + Delimiter)
	buf.WriteString(oof.FormatLineTwo(options) + Delimiter)
	buf.WriteString(oof.FormatLineThree(options) + Delimiter)

	if options.VariableLengthFields {
		return oof.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// FormatPartyIdentifier returns PartyIdentifier formatted according to the FormatOptions
func (oof *OriginatorOptionF) FormatPartyIdentifier(options FormatOptions) string {
	return oof.formatAlphaField(oof.PartyIdentifier, 35, options)
}

// FormatName returns Name formatted according to the FormatOptions
func (oof *OriginatorOptionF) FormatName(options FormatOptions) string {
	return oof.formatAlphaField(oof.Name, 35, options)
}

// FormatLineOne returns LineOne formatted according to the FormatOptions
func (oof *OriginatorOptionF) FormatLineOne(options FormatOptions) string {
	return oof.formatAlphaField(oof.LineOne, 35, options)
}

// FormatLineTwo returns LineTwo formatted according to the FormatOptions
func (oof *OriginatorOptionF) FormatLineTwo(options FormatOptions) string {
	return oof.formatAlphaField(oof.LineTwo, 35, options)
}

// FormatLineThree returns LineThree formatted according to the FormatOptions
func (oof *OriginatorOptionF) FormatLineThree(options FormatOptions) string {
	return oof.formatAlphaField(oof.LineThree, 35, options)
}
