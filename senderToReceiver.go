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
}

// String writes SenderToReceiver
func (sr *SenderToReceiver) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(215)
	buf.WriteString(sr.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderToReceiver and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sr *SenderToReceiver) Validate() error {
	if err := sr.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sr *SenderToReceiver) fieldInclusion() error {
	return nil
}
