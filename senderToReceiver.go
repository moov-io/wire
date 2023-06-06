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

	if err := str.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width SenderToReceiver record
func (str *SenderToReceiver) String() string {
	return str.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a SenderToReceiver record formatted according to the FormatOptions
func (str *SenderToReceiver) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(221)

	buf.WriteString(str.tag)
	buf.WriteString(str.FormatSwiftFieldTag(options))
	buf.WriteString(str.FormatSwiftLineOne(options))
	buf.WriteString(str.FormatSwiftLineTwo(options))
	buf.WriteString(str.FormatSwiftLineThree(options))
	buf.WriteString(str.FormatSwiftLineFour(options))
	buf.WriteString(str.FormatSwiftLineFive(options))
	buf.WriteString(str.FormatSwiftLineSix(options))

	if options.VariableLengthFields {
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

// FormatSwiftFieldTag returns CoverPayment.SwiftFieldTag formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftFieldTag(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftFieldTag, 5, options)
}

// FormatSwiftLineOne returns CoverPayment.SwiftLineOne formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftLineOne(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftLineOne, 35, options)
}

// FormatSwiftLineTwo returns CoverPayment.SwiftLineTwo formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftLineTwo(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftLineTwo, 35, options)
}

// FormatSwiftLineThree returns CoverPayment.SwiftLineThree formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftLineThree(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftLineThree, 35, options)
}

// FormatSwiftLineFour returns CoverPayment.SwiftLineFour formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftLineFour(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftLineFour, 35, options)
}

// FormatSwiftLineFive returns CoverPayment.SwiftLineFive formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftLineFive(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftLineFive, 35, options)
}

// FormatSwiftLineSix returns CoverPayment.SwiftLineSix formatted according to the FormatOptions
func (str *SenderToReceiver) FormatSwiftLineSix(options FormatOptions) string {
	return str.formatAlphaField(str.CoverPayment.SwiftLineSix, 35, options)
}
