// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// OrderingInstitution is the ordering institution
type OrderingInstitution struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingInstitution returns a new OrderingInstitution
func NewOrderingInstitution() OrderingInstitution  {
	oi := OrderingInstitution {
		tag: TagOrderingInstitution,
	}
	return oi
}

// Parse takes the input string and parses the OrderingInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oi *OrderingInstitution) Parse(record string) {
}

// String writes OrderingInstitution
func (oi *OrderingInstitution) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(180)
	buf.WriteString(oi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OrderingInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oi *OrderingInstitution) Validate() error {
	if err := oi.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oi *OrderingInstitution) fieldInclusion() error {
	return nil
}
