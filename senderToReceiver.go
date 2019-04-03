// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func NewSenderToReceiver() SenderToReceiver {
	sr := SenderToReceiver{
		tag: TagSenderToReceiver,
	}
	return sr
}

// Parse takes the input string and parses the SenderToReceiver values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sr *SenderToReceiver) Parse(record string) {
	sr.tag = record[:6]
	sr.CoverPayment.SwiftFieldTag = sr.parseStringField(record[6:11])
	sr.CoverPayment.SwiftLineOne = sr.parseStringField(record[11:46])
	sr.CoverPayment.SwiftLineTwo = sr.parseStringField(record[46:81])
	sr.CoverPayment.SwiftLineThree = sr.parseStringField(record[81:116])
	sr.CoverPayment.SwiftLineFour = sr.parseStringField(record[116:151])
	sr.CoverPayment.SwiftLineFive = sr.parseStringField(record[151:186])
	sr.CoverPayment.SwiftLineSix = sr.parseStringField(record[186:221])
}

// String writes SenderToReceiver
func (sr *SenderToReceiver) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(221)
	buf.WriteString(sr.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderToReceiver and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sr *SenderToReceiver) Validate() error {
	if err := sr.fieldInclusion(); err != nil {
		return err
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, sr.CoverPayment.SwiftFieldTag)
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, sr.CoverPayment.SwiftLineOne)
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, sr.CoverPayment.SwiftLineTwo)
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, sr.CoverPayment.SwiftLineThree)
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, sr.CoverPayment.SwiftLineFour)
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, sr.CoverPayment.SwiftLineFive)
	}
	if err := sr.isAlphanumeric(sr.CoverPayment.SwiftLineSix); err != nil {
		return fieldError("SwiftLineSix", err, sr.CoverPayment.SwiftLineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sr *SenderToReceiver) fieldInclusion() error {
	return nil
}
