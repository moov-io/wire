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
	// is variable length
	isVariableLength bool
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryCustomer returns a new BeneficiaryCustomer
func NewBeneficiaryCustomer(isVariable bool) *BeneficiaryCustomer {
	bc := &BeneficiaryCustomer{
		tag:              TagBeneficiaryCustomer,
		isVariableLength: isVariable,
	}
	return bc
}

// Parse takes the input string and parses the BeneficiaryCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bc *BeneficiaryCustomer) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 12 {
		return NewTagWrongLengthErr(12, len(record)), 0
	}
	bc.tag = record[:6]

	length := 6
	read := 0

	bc.CoverPayment.SwiftFieldTag, read = bc.parseVariableStringField(record[length:], 5)
	length += read

	bc.CoverPayment.SwiftLineOne, read = bc.parseVariableStringField(record[length:], 35)
	length += read

	bc.CoverPayment.SwiftLineTwo, read = bc.parseVariableStringField(record[length:], 35)
	length += read

	bc.CoverPayment.SwiftLineThree, read = bc.parseVariableStringField(record[length:], 35)
	length += read

	bc.CoverPayment.SwiftLineFour, read = bc.parseVariableStringField(record[length:], 35)
	length += read

	bc.CoverPayment.SwiftLineFive, read = bc.parseVariableStringField(record[length:], 35)
	length += read

	return nil, length
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
func (bc *BeneficiaryCustomer) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(bc.tag)
	buf.WriteString(bc.SwiftFieldTagField())
	buf.WriteString(bc.SwiftLineOneField())
	buf.WriteString(bc.SwiftLineTwoField())
	buf.WriteString(bc.SwiftLineThreeField())
	buf.WriteString(bc.SwiftLineFourField())
	buf.WriteString(bc.SwiftLineFiveField())
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

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (bc *BeneficiaryCustomer) SwiftFieldTagField() string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftFieldTag, 5, bc.isVariableLength)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (bc *BeneficiaryCustomer) SwiftLineOneField() string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineOne, 35, bc.isVariableLength)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (bc *BeneficiaryCustomer) SwiftLineTwoField() string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineTwo, 35, bc.isVariableLength)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (bc *BeneficiaryCustomer) SwiftLineThreeField() string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineThree, 35, bc.isVariableLength)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (bc *BeneficiaryCustomer) SwiftLineFourField() string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineFour, 35, bc.isVariableLength)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (bc *BeneficiaryCustomer) SwiftLineFiveField() string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineFive, 35, bc.isVariableLength)
}
