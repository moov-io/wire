// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// Address is 3 lines of address information
type Address struct {
	// AddressLineOne
	AddressLineOne string `json:"addressLineOne,omitempty"`
	// AddressLineTwo
	AddressLineTwo string `json:"addressLineTwo,omitempty"`
	// AddressLineThree
	AddressLineThree string `json:"addressLineThree,omitempty"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the Address values
func (a *Address) Parse(record string) (length int, err error) {

	var read int

	if a.AddressLineOne, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("AddressLineOne", err)
	}
	length += read

	if a.AddressLineTwo, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("AddressLineTwo", err)
	}
	length += read

	if a.AddressLineThree, read, err = a.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("AddressLineThree", err)
	}
	length += read

	return
}

// String writes BeneficiaryCustomer
func (a *Address) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(105)

	buf.WriteString(a.AddressLineOneField(isVariable))
	buf.WriteString(a.AddressLineTwoField(isVariable))
	buf.WriteString(a.AddressLineThreeField(isVariable))

	return buf.String()
}

// AddressLineOneField gets a string of AddressLineOne field
func (a *Address) AddressLineOneField(isVariable bool) string {
	return a.alphaVariableField(a.AddressLineOne, 35, isVariable)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (a *Address) AddressLineTwoField(isVariable bool) string {
	return a.alphaVariableField(a.AddressLineTwo, 35, isVariable)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (a *Address) AddressLineThreeField(isVariable bool) string {
	return a.alphaVariableField(a.AddressLineThree, 35, isVariable)
}
