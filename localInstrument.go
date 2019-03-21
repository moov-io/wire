// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// LocalInstrument is the LocalInstrument of the wire
type LocalInstrument struct {
	// tag
	tag string
	// LocalInstrumentCode is local instrument code
	LocalInstrumentCode string `json:"LocalInstrument,omitempty"`
	// ProprietaryCode is proprietary code
	ProprietaryCode string `json:"proprietaryCode,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewLocalInstrument returns a new LocalInstrument
func NewLocalInstrument() LocalInstrument  {
	li := LocalInstrument {
		tag: TagLocalInstrument,
	}
	return li
}

// Parse takes the input string and parses the LocalInstrument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (li *LocalInstrument) Parse(record string) {
}

// String writes LocalInstrument
func (li *LocalInstrument) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(39)
	buf.WriteString(li.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiverDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (li *LocalInstrument) Validate() error {
	if err := li.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (li *LocalInstrument) fieldInclusion() error {
	return nil
}

