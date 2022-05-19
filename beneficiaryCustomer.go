// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &BeneficiaryCustomer{}

// BeneficiaryCustomer is the beneficiary customer
type BeneficiaryCustomer struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryCustomer returns a new BeneficiaryCustomer
func NewBeneficiaryCustomer() *BeneficiaryCustomer {
	bc := &BeneficiaryCustomer{
		tag: TagBeneficiaryCustomer,
	}
	return bc
}

// Parse takes the input string and parses the BeneficiaryCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bc *BeneficiaryCustomer) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if bc.tag, read, err = bc.parseTag(record); err != nil {
		return 0, fieldError("BeneficiaryCustomer.Tag", err)
	}
	length += read

	if read, err = bc.CoverPayment.ParseFive(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (bc *BeneficiaryCustomer) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryCustomer
	aux := struct {
		*Alias
	}{
		(*Alias)(bc),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bc.tag = TagBeneficiaryCustomer
	return nil
}

// String writes BeneficiaryCustomer
func (bc *BeneficiaryCustomer) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(bc.tag)
	buf.WriteString(bc.CoverPayment.StringFive(isCompressed))

	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BeneficiaryCustomer) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	if bc.tag != TagBeneficiaryCustomer {
		return fieldError("tag", ErrValidTagForType, bc.tag)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, bc.CoverPayment.SwiftFieldTag)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, bc.CoverPayment.SwiftLineOne)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, bc.CoverPayment.SwiftLineTwo)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, bc.CoverPayment.SwiftLineThree)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, bc.CoverPayment.SwiftLineFour)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, bc.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bc *BeneficiaryCustomer) fieldInclusion() error {
	if bc.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, bc.CoverPayment.SwiftLineSix)
	}
	return nil
}
