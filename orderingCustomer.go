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
	oc.tag = record[:6]
	oc.CoverPayment.SwiftFieldTag = oc.parseStringField(record[6:11])
	oc.CoverPayment.SwiftLineOne = oc.parseStringField(record[11:46])
	oc.CoverPayment.SwiftLineTwo = oc.parseStringField(record[46:81])
	oc.CoverPayment.SwiftLineThree = oc.parseStringField(record[81:116])
	oc.CoverPayment.SwiftLineFour = oc.parseStringField(record[116:151])
	oc.CoverPayment.SwiftLineFive = oc.parseStringField(record[151:186])

}

// String writes OrderingCustomer
func (oc *OrderingCustomer) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(186)
	buf.WriteString(oc.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on OrderingCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (oc *OrderingCustomer) Validate() error {
	if err := oc.fieldInclusion(); err != nil {
		return err
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, oc.CoverPayment.SwiftFieldTag)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, oc.CoverPayment.SwiftLineOne)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, oc.CoverPayment.SwiftLineTwo)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, oc.CoverPayment.SwiftLineThree)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, oc.CoverPayment.SwiftLineFour)
	}
	if err := oc.isAlphanumeric(oc.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, oc.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (oc *OrderingCustomer) fieldInclusion() error {
	return nil
}
