// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// BeneficiaryReference is a reference for the beneficiary
type BeneficiaryReference struct {
	// tag
	tag string
	// BeneficiaryReference
	BeneficiaryReference string `json:"beneficiaryReference,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryReference returns a new BeneficiaryReference
func NewBeneficiaryReference() *BeneficiaryReference {
	br := &BeneficiaryReference{
		tag: TagBeneficiaryReference,
	}
	return br
}

// Parse takes the input string and parses the BeneficiaryReference values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (br *BeneficiaryReference) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	br.tag = record[:6]

	var err error
	length := 6
	read := 0

	if br.BeneficiaryReference, read, err = br.parseVariableStringField(record[length:], 16); err != nil {
		return fieldError("BeneficiaryReference", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (br *BeneficiaryReference) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryReference
	aux := struct {
		*Alias
	}{
		(*Alias)(br),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	br.tag = TagBeneficiaryReference
	return nil
}

// String writes BeneficiaryReference
func (br *BeneficiaryReference) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(22)

	buf.WriteString(br.tag)
	buf.WriteString(br.BeneficiaryReferenceField(options...))

	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryReference and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (br *BeneficiaryReference) Validate() error {
	if br.tag != TagBeneficiaryReference {
		return fieldError("tag", ErrValidTagForType, br.tag)
	}
	if err := br.isAlphanumeric(br.BeneficiaryReference); err != nil {
		return fieldError("BeneficiaryReference", err, br.BeneficiaryReference)
	}
	return nil
}

// BeneficiaryReferenceField gets a string of the BeneficiaryReference field
func (br *BeneficiaryReference) BeneficiaryReferenceField(options ...bool) string {
	return br.alphaVariableField(br.BeneficiaryReference, 16, br.parseFirstOption(options))
}
