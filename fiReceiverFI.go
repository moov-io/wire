// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIReceiverFI is the financial institution receiver financial institution
type FIReceiverFI struct {
	// tag
	tag string
	// FIToFI is financial institution to financial institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIReceiverFI returns a new FIReceiverFI
func NewFIReceiverFI() FIReceiverFI {
	rfi := FIReceiverFI{
		tag: TagFIReceiverFI,
	}
	return rfi
}

// Parse takes the input string and parses the FIReceiverFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rfi *FIReceiverFI) Parse(record string) {
	rfi.tag = record[:6]
	rfi.FIToFI.LineOne = rfi.parseStringField(record[6:36])
	rfi.FIToFI.LineTwo = rfi.parseStringField(record[36:69])
	rfi.FIToFI.LineThree = rfi.parseStringField(record[69:104])
	rfi.FIToFI.LineFour = rfi.parseStringField(record[104:139])
	rfi.FIToFI.LineFive = rfi.parseStringField(record[139:174])
	rfi.FIToFI.LineSix = rfi.parseStringField(record[174:209])
}

// String writes FIReceiverFI
func (rfi *FIReceiverFI) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(209)
	buf.WriteString(rfi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIReceiverFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rfi *FIReceiverFI) Validate() error {
	if err := rfi.fieldInclusion(); err != nil {
		return err
	}
	if err:= rfi.isAlphanumeric(rfi.FIToFI.LineOne); err!= nil {
		return fieldError("LineOne", err, rfi.FIToFI.LineOne)
	}
	if err:= rfi.isAlphanumeric(rfi.FIToFI.LineTwo); err!= nil {
		return fieldError("LineTwo", err, rfi.FIToFI.LineTwo)
	}
	if err:= rfi.isAlphanumeric(rfi.FIToFI.LineThree); err!= nil {
		return fieldError("LineThree", err, rfi.FIToFI.LineThree)
	}
	if err:= rfi.isAlphanumeric(rfi.FIToFI.LineFour); err!= nil {
		return fieldError("LineFour", err, rfi.FIToFI.LineFour)
	}
	if err:= rfi.isAlphanumeric(rfi.FIToFI.LineFive); err!= nil {
		return fieldError("LineFive", err, rfi.FIToFI.LineFive)
	}
	if err:= rfi.isAlphanumeric(rfi.FIToFI.LineSix); err!= nil {
		return fieldError("LineSix", err, rfi.FIToFI.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rfi *FIReceiverFI) fieldInclusion() error {
	return nil
}
