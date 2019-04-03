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
	ifia.tag = record[:6]
	ifia.Advice.AdviceCode = ifia.parseStringField(record[6:9])
	ifia.Advice.LineOne = ifia.parseStringField(record[9:35])
	ifia.Advice.LineTwo = ifia.parseStringField(record[35:68])
	ifia.Advice.LineThree = ifia.parseStringField(record[68:101])
	ifia.Advice.LineFour = ifia.parseStringField(record[101:134])
	ifia.Advice.LineFive = ifia.parseStringField(record[134:167])
	ifia.Advice.LineSix = ifia.parseStringField(record[167:200])
}

// String writes FIIntermediaryFIAdvice
func (ifia *FIIntermediaryFIAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(200)
	buf.WriteString(ifia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ifia *FIIntermediaryFIAdvice) Validate() error {
	if err := ifia.fieldInclusion(); err != nil {
		return err
	}
	if err := ifia.isAdviceCode(ifia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, ifia.Advice.AdviceCode)
	}
	if err:= ifia.isAlphanumeric(ifia.Advice.LineOne); err!= nil {
		return fieldError("LineOne", err, ifia.Advice.LineOne)
	}
	if err:= ifia.isAlphanumeric(ifia.Advice.LineTwo); err!= nil {
		return fieldError("LineTwo", err, ifia.Advice.LineTwo)
	}
	if err:= ifia.isAlphanumeric(ifia.Advice.LineThree); err!= nil {
		return fieldError("LineThree", err, ifia.Advice.LineThree)
	}
	if err:= ifia.isAlphanumeric(ifia.Advice.LineFour); err!= nil {
		return fieldError("LineFour", err, ifia.Advice.LineFour)
	}
	if err:= ifia.isAlphanumeric(ifia.Advice.LineFive); err!= nil {
		return fieldError("LineFive", err, ifia.Advice.LineFive)
	}
	if err:= ifia.isAlphanumeric(ifia.Advice.LineSix); err!= nil {
		return fieldError("LineSix", err, ifia.Advice.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ifia *FIIntermediaryFIAdvice) fieldInclusion() error {
	return nil
}
