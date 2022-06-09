// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
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
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	str.tag = record[:6]
	length := 6

	value, read, err := str.parseVariableStringField(record[length:], 5)
	if err != nil {
		return fieldError("SwiftFieldTag", err)
	}
	str.CoverPayment.SwiftFieldTag = value
	length += read

	value, read, err = str.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineOne", err)
	}
	str.CoverPayment.SwiftLineOne = value
	length += read

	value, read, err = str.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineTwo", err)
	}
	str.CoverPayment.SwiftLineTwo = value
	length += read

	value, read, err = str.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineThree", err)
	}
	str.CoverPayment.SwiftLineThree = value
	length += read

	value, read, err = str.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFour", err)
	}
	str.CoverPayment.SwiftLineFour = value
	length += read

	value, read, err = str.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineFive", err)
	}
	str.CoverPayment.SwiftLineFive = value
	length += read

	value, read, err = str.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("SwiftLineSix", err)
	}
	str.CoverPayment.SwiftLineSix = value
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (str *SenderToReceiver) UnmarshalJSON(data []byte) error {
	type Alias SenderToReceiver
	aux := struct {
		*Alias
	}{
		(*Alias)(str),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	str.tag = TagSenderToReceiver
	return nil
}

// String writes SenderToReceiver
func (str *SenderToReceiver) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(221)

	buf.WriteString(str.tag)
	buf.WriteString(str.SwiftFieldTagField(options...))
	buf.WriteString(str.SwiftLineOneField(options...))
	buf.WriteString(str.SwiftLineTwoField(options...))
	buf.WriteString(str.SwiftLineThreeField(options...))
	buf.WriteString(str.SwiftLineFourField(options...))
	buf.WriteString(str.SwiftLineFiveField(options...))
	buf.WriteString(str.SwiftLineSixField(options...))

	if str.parseFirstOption(options) {
		return str.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
func (str *SenderToReceiver) SwiftFieldTagField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftFieldTag, 5, str.parseFirstOption(options))
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (str *SenderToReceiver) SwiftLineOneField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftLineOne, 35, str.parseFirstOption(options))
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (str *SenderToReceiver) SwiftLineTwoField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftLineTwo, 35, str.parseFirstOption(options))
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (str *SenderToReceiver) SwiftLineThreeField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftLineThree, 35, str.parseFirstOption(options))
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (str *SenderToReceiver) SwiftLineFourField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftLineFour, 35, str.parseFirstOption(options))
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (str *SenderToReceiver) SwiftLineFiveField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftLineFive, 35, str.parseFirstOption(options))
}

// SwiftLineSixField gets a string of the SwiftLineSix field
func (str *SenderToReceiver) SwiftLineSixField(options ...bool) string {
	return str.alphaVariableField(str.CoverPayment.SwiftLineSix, 35, str.parseFirstOption(options))
}
