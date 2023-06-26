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

	if err := bc.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width BeneficiaryCustomer record
func (bc *BeneficiaryCustomer) String() string {
	return bc.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a BeneficiaryCustomer record formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(bc.tag)
	buf.WriteString(bc.FormatSwiftFieldTag(options) + Delimiter)
	buf.WriteString(bc.FormatSwiftLineOne(options) + Delimiter)
	buf.WriteString(bc.FormatSwiftLineTwo(options) + Delimiter)
	buf.WriteString(bc.FormatSwiftLineThree(options) + Delimiter)
	buf.WriteString(bc.FormatSwiftLineFour(options) + Delimiter)
	buf.WriteString(bc.FormatSwiftLineFive(options) + Delimiter)

	if options.VariableLengthFields {
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
func (bc *BeneficiaryCustomer) SwiftFieldTagField() string {
	return bc.alphaField(bc.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (bc *BeneficiaryCustomer) SwiftLineOneField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (bc *BeneficiaryCustomer) SwiftLineTwoField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (bc *BeneficiaryCustomer) SwiftLineThreeField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (bc *BeneficiaryCustomer) SwiftLineFourField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (bc *BeneficiaryCustomer) SwiftLineFiveField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineFive, 35)
}

// FormatSwiftFieldTag returns CoverPayment.SwiftFieldTag formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) FormatSwiftFieldTag(options FormatOptions) string {
	return bc.formatAlphaField(bc.CoverPayment.SwiftFieldTag, 5, options)
}

// FormatSwiftLineOne returns CoverPayment.SwiftLineOne formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) FormatSwiftLineOne(options FormatOptions) string {
	return bc.formatAlphaField(bc.CoverPayment.SwiftLineOne, 35, options)
}

// FormatSwiftLineTwo returns CoverPayment.SwiftLineTwo formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) FormatSwiftLineTwo(options FormatOptions) string {
	return bc.formatAlphaField(bc.CoverPayment.SwiftLineTwo, 35, options)
}

// FormatSwiftLineThree returns CoverPayment.SwiftLineThree formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) FormatSwiftLineThree(options FormatOptions) string {
	return bc.formatAlphaField(bc.CoverPayment.SwiftLineThree, 35, options)
}

// FormatSwiftLineFour returns CoverPayment.SwiftLineFour formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) FormatSwiftLineFour(options FormatOptions) string {
	return bc.formatAlphaField(bc.CoverPayment.SwiftLineFour, 35, options)
}

// FormatSwiftLineFive returns CoverPayment.SwiftLineFive formatted according to the FormatOptions
func (bc *BeneficiaryCustomer) FormatSwiftLineFive(options FormatOptions) string {
	return bc.formatAlphaField(bc.CoverPayment.SwiftLineFive, 35, options)
}
