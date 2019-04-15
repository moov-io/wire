// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIBeneficiaryFIAdvice is the financial institution beneficiary financial institution
type FIBeneficiaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFIAdvice returns a new FIBeneficiaryFIAdvice
func NewFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := &FIBeneficiaryFIAdvice{
		tag: TagFIBeneficiaryFIAdvice,
	}
	return fibfia
}

// Parse takes the input string and parses the FIBeneficiaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfia *FIBeneficiaryFIAdvice) Parse(record string) {
	fibfia.tag = record[:6]
	fibfia.Advice.AdviceCode = fibfia.parseStringField(record[6:9])
	fibfia.Advice.LineOne = fibfia.parseStringField(record[9:35])
	fibfia.Advice.LineTwo = fibfia.parseStringField(record[35:68])
	fibfia.Advice.LineThree = fibfia.parseStringField(record[68:101])
	fibfia.Advice.LineFour = fibfia.parseStringField(record[101:134])
	fibfia.Advice.LineFive = fibfia.parseStringField(record[134:167])
	fibfia.Advice.LineSix = fibfia.parseStringField(record[167:200])
}

// String writes FIBeneficiaryFIAdvice
func (fibfia *FIBeneficiaryFIAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(200)
	buf.WriteString(fibfia.tag)
	buf.WriteString(fibfia.AdviceCodeField())
	buf.WriteString(fibfia.LineOneField())
	buf.WriteString(fibfia.LineTwoField())
	buf.WriteString(fibfia.LineThreeField())
	buf.WriteString(fibfia.LineFourField())
	buf.WriteString(fibfia.LineFiveField())
	buf.WriteString(fibfia.LineSixField())
	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfia *FIBeneficiaryFIAdvice) Validate() error {
	if err := fibfia.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (fibfia *FIBeneficiaryFIAdvice) fieldInclusion() error {
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fibfia *FIBeneficiaryFIAdvice) AdviceCodeField() string {
	return fibfia.alphaField(fibfia.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fibfia *FIBeneficiaryFIAdvice) LineOneField() string {
	return fibfia.alphaField(fibfia.Advice.LineOne, 26)
}

// LineTwoField gets a string of the LineTwo field
func (fibfia *FIBeneficiaryFIAdvice) LineTwoField() string {
	return fibfia.alphaField(fibfia.Advice.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fibfia *FIBeneficiaryFIAdvice) LineThreeField() string {
	return fibfia.alphaField(fibfia.Advice.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fibfia *FIBeneficiaryFIAdvice) LineFourField() string {
	return fibfia.alphaField(fibfia.Advice.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fibfia *FIBeneficiaryFIAdvice) LineFiveField() string {
	return fibfia.alphaField(fibfia.Advice.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fibfia *FIBeneficiaryFIAdvice) LineSixField() string {
	return fibfia.alphaField(fibfia.Advice.LineSix, 33)
}
