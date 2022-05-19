// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &SenderToReceiver{}

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
func (str *SenderToReceiver) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 13 {
		return 0, NewTagWrongLengthErr(13, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if str.tag, read, err = str.parseTag(record); err != nil {
		return 0, fieldError("SenderToReceiver.Tag", err)
	}
	length += read

	if read, err = str.CoverPayment.ParseSix(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(221)

	buf.WriteString(str.tag)
	buf.WriteString(str.CoverPayment.StringSix(isCompressed))

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
