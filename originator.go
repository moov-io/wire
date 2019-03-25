// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// Originator is the originator of the wire
type Originator struct {
	// tag
	tag string
	// Personal
	Personal Personal `json:"personal,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOriginator returns a new Originator
func NewOriginator() Originator  {
	o := Originator {
		tag: TagOriginator,
	}
	return o
}

// Parse takes the input string and parses the Originator values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (o *Originator) Parse(record string) {
}

// String writes Originator
func (o *Originator) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(o.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on Originator and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (o *Originator) Validate() error {
	if err := o.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (o *Originator) fieldInclusion() error {
	return nil
}

