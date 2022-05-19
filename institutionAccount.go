// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &InstitutionAccount{}

// InstitutionAccount is the institution account
type InstitutionAccount struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInstitutionAccount returns a new InstitutionAccount
func NewInstitutionAccount() *InstitutionAccount {
	iAccount := &InstitutionAccount{
		tag: TagInstitutionAccount,
	}
	return iAccount
}

// Parse takes the input string and parses the InstitutionAccount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (iAccount *InstitutionAccount) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if iAccount.tag, read, err = iAccount.parseTag(record); err != nil {
		return 0, fieldError("InstitutionAccount.Tag", err)
	}
	length += read

	if read, err = iAccount.CoverPayment.ParseFive(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (iAccount *InstitutionAccount) UnmarshalJSON(data []byte) error {
	type Alias InstitutionAccount
	aux := struct {
		*Alias
	}{
		(*Alias)(iAccount),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	iAccount.tag = TagInstitutionAccount
	return nil
}

// String writes InstitutionAccount
func (iAccount *InstitutionAccount) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(iAccount.tag)
	buf.WriteString(iAccount.CoverPayment.StringFive(isCompressed))

	return buf.String()
}

// Validate performs WIRE format rule checks on InstitutionAccount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (iAccount *InstitutionAccount) Validate() error {
	if err := iAccount.fieldInclusion(); err != nil {
		return err
	}
	if iAccount.tag != TagInstitutionAccount {
		return fieldError("tag", ErrValidTagForType, iAccount.tag)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, iAccount.CoverPayment.SwiftFieldTag)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, iAccount.CoverPayment.SwiftLineOne)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, iAccount.CoverPayment.SwiftLineTwo)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, iAccount.CoverPayment.SwiftLineThree)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, iAccount.CoverPayment.SwiftLineFour)
	}
	if err := iAccount.isAlphanumeric(iAccount.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, iAccount.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (iAccount *InstitutionAccount) fieldInclusion() error {
	if iAccount.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, iAccount.CoverPayment.SwiftLineSix)
	}
	return nil
}
