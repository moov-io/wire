// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// CoverPayment is cover payment data
type CoverPayment struct {
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag,omitempty"`
	// SwiftLineOne
	SwiftLineOne string `json:"swiftLineOne,omitempty"`
	// SwiftLineTwo
	SwiftLineTwo string `json:"swiftLineTwo,omitempty"`
	// SwiftLineThree
	SwiftLineThree string `json:"swiftLineThree,omitempty"`
	// SwiftLineFour
	SwiftLineFour string `json:"swiftLineFour,omitempty"`
	// SwiftLineFive
	SwiftLineFive string `json:"swiftLineFive,omitempty"`
	// SwiftLineSix
	SwiftLineSix string `json:"swiftLineSix,omitempty"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the CoverPayment values
func (c *CoverPayment) ParseFour(record string) (length int, err error) {

	var read int

	if c.SwiftFieldTag, read, err = c.parseVariableStringField(record[length:], 5); err != nil {
		return 0, fieldError("SwiftFieldTag", err)
	}
	length += read

	if c.SwiftLineOne, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineOne", err)
	}
	length += read

	if c.SwiftLineTwo, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineTwo", err)
	}
	length += read

	if c.SwiftLineThree, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineThree", err)
	}
	length += read

	if c.SwiftLineFour, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineFour", err)
	}
	length += read

	return
}

// Parse takes the input string and parses the CoverPayment values
func (c *CoverPayment) ParseFive(record string) (length int, err error) {

	var read int

	if c.SwiftFieldTag, read, err = c.parseVariableStringField(record[length:], 5); err != nil {
		return 0, fieldError("SwiftFieldTag", err)
	}
	length += read

	if c.SwiftLineOne, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineOne", err)
	}
	length += read

	if c.SwiftLineTwo, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineTwo", err)
	}
	length += read

	if c.SwiftLineThree, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineThree", err)
	}
	length += read

	if c.SwiftLineFour, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineFour", err)
	}
	length += read

	if c.SwiftLineFive, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineFive", err)
	}
	length += read

	return
}

// Parse takes the input string and parses the CoverPayment values
func (c *CoverPayment) ParseSix(record string) (length int, err error) {

	var read int

	if c.SwiftFieldTag, read, err = c.parseVariableStringField(record[length:], 5); err != nil {
		return 0, fieldError("SwiftFieldTag", err)
	}
	length += read

	if c.SwiftLineOne, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineOne", err)
	}
	length += read

	if c.SwiftLineTwo, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineTwo", err)
	}
	length += read

	if c.SwiftLineThree, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineThree", err)
	}
	length += read

	if c.SwiftLineFour, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineFour", err)
	}
	length += read

	if c.SwiftLineFive, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineFive", err)
	}
	length += read

	if c.SwiftLineSix, read, err = c.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("SwiftLineSix", err)
	}
	length += read

	return
}

// String writes BeneficiaryCustomer
func (c *CoverPayment) StringFour(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(180)

	buf.WriteString(c.SwiftFieldTagField(isVariable))
	buf.WriteString(c.SwiftLineOneField(isVariable))
	buf.WriteString(c.SwiftLineTwoField(isVariable))
	buf.WriteString(c.SwiftLineThreeField(isVariable))
	buf.WriteString(c.SwiftLineFourField(isVariable))

	return buf.String()
}

// String writes BeneficiaryCustomer
func (c *CoverPayment) StringFive(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(180)

	buf.WriteString(c.SwiftFieldTagField(isVariable))
	buf.WriteString(c.SwiftLineOneField(isVariable))
	buf.WriteString(c.SwiftLineTwoField(isVariable))
	buf.WriteString(c.SwiftLineThreeField(isVariable))
	buf.WriteString(c.SwiftLineFourField(isVariable))
	buf.WriteString(c.SwiftLineFiveField(isVariable))

	return buf.String()
}

// String writes BeneficiaryCustomer
func (c *CoverPayment) StringSix(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(180)

	buf.WriteString(c.SwiftFieldTagField(isVariable))
	buf.WriteString(c.SwiftLineOneField(isVariable))
	buf.WriteString(c.SwiftLineTwoField(isVariable))
	buf.WriteString(c.SwiftLineThreeField(isVariable))
	buf.WriteString(c.SwiftLineFourField(isVariable))
	buf.WriteString(c.SwiftLineFiveField(isVariable))
	buf.WriteString(c.SwiftLineSixField(isVariable))

	return buf.String()
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (c *CoverPayment) SwiftFieldTagField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftFieldTag, 5, isVariable)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (c *CoverPayment) SwiftLineOneField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftLineOne, 35, isVariable)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (c *CoverPayment) SwiftLineTwoField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftLineTwo, 35, isVariable)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (c *CoverPayment) SwiftLineThreeField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftLineThree, 35, isVariable)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (c *CoverPayment) SwiftLineFourField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftLineFour, 35, isVariable)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (c *CoverPayment) SwiftLineFiveField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftLineFive, 35, isVariable)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (c *CoverPayment) SwiftLineSixField(isVariable bool) string {
	return c.alphaVariableField(c.SwiftLineSix, 35, isVariable)
}
