// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (iAccount *InstitutionAccount) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	iAccount.tag = record[:6]
	length := 6

	value, read, err := iAccount.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	iAccount.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = iAccount.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	iAccount.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = iAccount.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	iAccount.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = iAccount.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	iAccount.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = iAccount.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	iAccount.CoverPayment.SwiftLineFour = value
	length += read

	value, read, err = iAccount.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFive", err)
	}
	iAccount.CoverPayment.SwiftLineFive = value
	length += read

	if err := iAccount.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width InstitutionAccount record
func (iAccount *InstitutionAccount) String() string {
	return iAccount.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a InstitutionAccount record formatted according to the FormatOptions
func (iAccount *InstitutionAccount) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(186)

	buf.WriteString(iAccount.tag)
	buf.WriteString(iAccount.FormatSwiftFieldTag(options))
	buf.WriteString(iAccount.FormatSwiftLineOne(options))
	buf.WriteString(iAccount.FormatSwiftLineTwo(options))
	buf.WriteString(iAccount.FormatSwiftLineThree(options))
	buf.WriteString(iAccount.FormatSwiftLineFour(options))
	buf.WriteString(iAccount.FormatSwiftLineFive(options))

	if options.VariableLengthFields {
		return iAccount.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (iAccount *InstitutionAccount) SwiftFieldTagField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (iAccount *InstitutionAccount) SwiftLineOneField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (iAccount *InstitutionAccount) SwiftLineTwoField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (iAccount *InstitutionAccount) SwiftLineThreeField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (iAccount *InstitutionAccount) SwiftLineFourField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (iAccount *InstitutionAccount) SwiftLineFiveField() string {
	return iAccount.alphaField(iAccount.CoverPayment.SwiftLineFive, 35)
}

// FormatSwiftFieldTag returns CoverPayment.SwiftFieldTag formatted according to the FormatOptions
func (iAccount *InstitutionAccount) FormatSwiftFieldTag(options FormatOptions) string {
	return iAccount.formatAlphaField(iAccount.CoverPayment.SwiftFieldTag, 5, options)
}

// FormatSwiftLineOne returns CoverPayment.SwiftLineOne formatted according to the FormatOptions
func (iAccount *InstitutionAccount) FormatSwiftLineOne(options FormatOptions) string {
	return iAccount.formatAlphaField(iAccount.CoverPayment.SwiftLineOne, 35, options)
}

// FormatSwiftLineTwo returns CoverPayment.SwiftLineTwo formatted according to the FormatOptions
func (iAccount *InstitutionAccount) FormatSwiftLineTwo(options FormatOptions) string {
	return iAccount.formatAlphaField(iAccount.CoverPayment.SwiftLineTwo, 35, options)
}

// FormatSwiftLineThree returns CoverPayment.SwiftLineThree formatted according to the FormatOptions
func (iAccount *InstitutionAccount) FormatSwiftLineThree(options FormatOptions) string {
	return iAccount.formatAlphaField(iAccount.CoverPayment.SwiftLineThree, 35, options)
}

// FormatSwiftLineFour returns CoverPayment.SwiftLineFour formatted according to the FormatOptions
func (iAccount *InstitutionAccount) FormatSwiftLineFour(options FormatOptions) string {
	return iAccount.formatAlphaField(iAccount.CoverPayment.SwiftLineFour, 35, options)
}

// FormatSwiftLineFive returns CoverPayment.SwiftLineFive formatted according to the FormatOptions
func (iAccount *InstitutionAccount) FormatSwiftLineFive(options FormatOptions) string {
	return iAccount.formatAlphaField(iAccount.CoverPayment.SwiftLineFive, 35, options)
}
