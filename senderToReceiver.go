// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// SenderToReceiver is the remittance information
type SenderToReceiver struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSenderToReceiver returns a new SenderToReceiver
func NewSenderToReceiver() *SenderToReceiver {
	str := &SenderToReceiver{
		tag: TagSenderToReceiver,
	}
	return str
}

// Parse takes the input string and parses the SenderToReceiver values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (str *SenderToReceiver) Parse(record string) error {
	if utf8.RuneCountInString(record) != 221 {
		return NewTagWrongLengthErr(221, utf8.RuneCountInString(record))
	}
	str.tag = record[:6]
	str.CoverPayment.SwiftFieldTag = str.parseStringField(record[6:11])
	str.CoverPayment.SwiftLineOne = str.parseStringField(record[11:46])
	str.CoverPayment.SwiftLineTwo = str.parseStringField(record[46:81])
	str.CoverPayment.SwiftLineThree = str.parseStringField(record[81:116])
	str.CoverPayment.SwiftLineFour = str.parseStringField(record[116:151])
	str.CoverPayment.SwiftLineFive = str.parseStringField(record[151:186])
	str.CoverPayment.SwiftLineSix = str.parseStringField(record[186:221])
	return nil
}

// String writes SenderToReceiver
func (str *SenderToReceiver) String() string {
	var buf strings.Builder
	buf.Grow(221)
	buf.WriteString(str.tag)
	buf.WriteString(str.SwiftFieldTagField())
	buf.WriteString(str.SwiftLineOneField())
	buf.WriteString(str.SwiftLineTwoField())
	buf.WriteString(str.SwiftLineThreeField())
	buf.WriteString(str.SwiftLineFourField())
	buf.WriteString(str.SwiftLineFiveField())
	buf.WriteString(str.SwiftLineSixField())
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderToReceiver and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (str *SenderToReceiver) Validate() error {
	if str.tag != TagSenderToReceiver {
		return fieldError("tag", ErrValidTagForType, str.tag)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, str.CoverPayment.SwiftFieldTag)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, str.CoverPayment.SwiftLineOne)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, str.CoverPayment.SwiftLineTwo)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, str.CoverPayment.SwiftLineThree)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, str.CoverPayment.SwiftLineFour)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, str.CoverPayment.SwiftLineFive)
	}
	if err := str.isAlphanumeric(str.CoverPayment.SwiftLineSix); err != nil {
		return fieldError("SwiftLineSix", err, str.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (str *SenderToReceiver) SwiftFieldTagField() string {
	return str.alphaField(str.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (str *SenderToReceiver) SwiftLineOneField() string {
	return str.alphaField(str.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (str *SenderToReceiver) SwiftLineTwoField() string {
	return str.alphaField(str.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (str *SenderToReceiver) SwiftLineThreeField() string {
	return str.alphaField(str.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (str *SenderToReceiver) SwiftLineFourField() string {
	return str.alphaField(str.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (str *SenderToReceiver) SwiftLineFiveField() string {
	return str.alphaField(str.CoverPayment.SwiftLineFive, 35)
}

// SwiftLineSixField gets a string of the SwiftLineSix field
func (str *SenderToReceiver) SwiftLineSixField() string {
	return str.alphaField(str.CoverPayment.SwiftLineSix, 35)
}
