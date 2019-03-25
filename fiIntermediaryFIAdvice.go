// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIIntermediaryFIAdvice is the financial institution intermediary financial institution
type FIIntermediaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFIAdvice returns a new FIIntermediaryFIAdvice
func NewFIIntermediaryFIAdvice() FIIntermediaryFIAdvice {
	ifia := FIIntermediaryFIAdvice{
		tag: TagFIIntermediaryFIAdvice,
	}
	return ifia
}

// Parse takes the input string and parses the FIIntermediaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ifia *FIIntermediaryFIAdvice) Parse(record string) {
}

// String writes FIIntermediaryFIAdvice
func (ifia *FIIntermediaryFIAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(194)
	buf.WriteString(ifia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ifia *FIIntermediaryFIAdvice) Validate() error {
	if err := ifia.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ifia *FIIntermediaryFIAdvice) fieldInclusion() error {
	return nil
}



