// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (bc *BeneficiaryCustomer) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	bc.tag = record[:6]
	length := 6

	value, read, err := bc.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	bc.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = bc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	bc.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = bc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	bc.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = bc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	bc.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = bc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	bc.CoverPayment.SwiftLineFour = value
	length += read

	value, read, err = bc.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFive", err)
	}
	bc.CoverPayment.SwiftLineFive = value
	length += read

	if !bc.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(bc.tag)
	buf.WriteString(bc.SwiftFieldTagField(options...))
	buf.WriteString(bc.SwiftLineOneField(options...))
	buf.WriteString(bc.SwiftLineTwoField(options...))
	buf.WriteString(bc.SwiftLineThreeField(options...))
	buf.WriteString(bc.SwiftLineFourField(options...))
	buf.WriteString(bc.SwiftLineFiveField(options...))

	if bc.parseFirstOption(options) {
		return bc.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
func (bc *BeneficiaryCustomer) SwiftFieldTagField(options ...bool) string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftFieldTag, 5, bc.parseFirstOption(options))
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (bc *BeneficiaryCustomer) SwiftLineOneField(options ...bool) string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineOne, 35, bc.parseFirstOption(options))
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (bc *BeneficiaryCustomer) SwiftLineTwoField(options ...bool) string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineTwo, 35, bc.parseFirstOption(options))
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (bc *BeneficiaryCustomer) SwiftLineThreeField(options ...bool) string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineThree, 35, bc.parseFirstOption(options))
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (bc *BeneficiaryCustomer) SwiftLineFourField(options ...bool) string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineFour, 35, bc.parseFirstOption(options))
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (bc *BeneficiaryCustomer) SwiftLineFiveField(options ...bool) string {
	return bc.alphaVariableField(bc.CoverPayment.SwiftLineFive, 35, bc.parseFirstOption(options))
}
