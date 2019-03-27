// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// OrderingCustomer is the ordering customer
type OrderingCustomer struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOrderingCustomer returns a new OrderingCustomer
func NewOrderingCustomer() OrderingCustomer {
	oc := OrderingCustomer{
		tag: TagOrderingCustomer,
	}
	return oc
}

// Parse takes the input string and parses the OrderingCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (oc *OrderingCustomer) Parse(record string) {
}

// String writes OrderingCustomer
func (oc *OrderingCustomer) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(180)
	buf.WriteString(oc.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OrderingCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oc *OrderingCustomer) Validate() error {
	if err := oc.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oc *OrderingCustomer) fieldInclusion() error {
	return nil
}
