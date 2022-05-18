// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &OriginatorFI{}

// OriginatorFI is the originator Financial Institution
type OriginatorFI struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginatorFI returns a new OriginatorFI
func NewOriginatorFI(isVariable bool) *OriginatorFI {
	ofi := &OriginatorFI{
		tag:              TagOriginatorFI,
		isVariableLength: isVariable,
	}
	return ofi
}

// Parse takes the input string and parses the OriginatorFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ofi *OriginatorFI) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if ofi.tag, read, err = ofi.parseTag(record); err != nil {
		return 0, fieldError("OriginatorFI.Tag", err)
	}
	length += read

	if read, err = ofi.FinancialInstitution.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (ofi *OriginatorFI) UnmarshalJSON(data []byte) error {
	type Alias OriginatorFI
	aux := struct {
		*Alias
	}{
		(*Alias)(ofi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ofi.tag = TagOriginatorFI
	return nil
}

// String writes OriginatorFI
func (ofi *OriginatorFI) String() string {
	var buf strings.Builder
	buf.Grow(181)

	buf.WriteString(ofi.tag)
	buf.WriteString(ofi.FinancialInstitution.String(ofi.isVariableLength))

	return buf.String()
}

// Validate performs WIRE format rule checks on OriginatorFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (ofi *OriginatorFI) Validate() error {
	if err := ofi.fieldInclusion(); err != nil {
		return err
	}
	if ofi.tag != TagOriginatorFI {
		return fieldError("tag", ErrValidTagForType, ofi.tag)
	}
	if err := ofi.isIdentificationCode(ofi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, ofi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch ofi.FinancialInstitution.IdentificationCode {
	case
		"B", "C", "D", "F", "U":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, ofi.FinancialInstitution.IdentificationCode)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, ofi.FinancialInstitution.Identifier)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, ofi.FinancialInstitution.Name)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ofi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ofi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := ofi.isAlphanumeric(ofi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ofi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ofi *OriginatorFI) fieldInclusion() error {
	if ofi.FinancialInstitution.IdentificationCode != "" && ofi.FinancialInstitution.Identifier == "" {
		return fieldError("Identifier", ErrFieldRequired)
	}
	if ofi.FinancialInstitution.IdentificationCode == "" && ofi.FinancialInstitution.Identifier != "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	return nil
}
