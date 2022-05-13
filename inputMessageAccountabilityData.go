// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &InputMessageAccountabilityData{}

// InputMessageAccountabilityData (IMAD) {1520}
type InputMessageAccountabilityData struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// InputCycleDate CCYYMMDD
	InputCycleDate string `json:"inputCycleDate"`
	// InputSource
	InputSource string `json:"inputSource"`
	// InputSequenceNumber
	InputSequenceNumber string `json:"inputSequenceNumber"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInputMessageAccountabilityData returns a new InputMessageAccountabilityData
func NewInputMessageAccountabilityData(isVariable bool) *InputMessageAccountabilityData {
	imad := &InputMessageAccountabilityData{
		tag:              TagInputMessageAccountabilityData,
		isVariableLength: isVariable,
	}
	return imad
}

// Parse takes the input string and parses the InputMessageAccountabilityData values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (imad *InputMessageAccountabilityData) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagWrongLengthErr(9, len(record)), 0
	}

	imad.tag = record[:6]

	length := 6
	read := 0

	imad.InputCycleDate, read = imad.parseVariableStringField(record[length:], 8)
	length += read

	imad.InputSource, read = imad.parseVariableStringField(record[length:], 8)
	length += read

	imad.InputSequenceNumber, read = imad.parseVariableStringField(record[length:], 6)
	length += read

	return nil, length
}

func (imad *InputMessageAccountabilityData) UnmarshalJSON(data []byte) error {
	type Alias InputMessageAccountabilityData
	aux := struct {
		*Alias
	}{
		(*Alias)(imad),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	imad.tag = TagInputMessageAccountabilityData
	return nil
}

// String writes InputMessageAccountabilityData
func (imad *InputMessageAccountabilityData) String() string {
	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(imad.tag)
	buf.WriteString(imad.InputCycleDateField())
	buf.WriteString(imad.InputSourceField())
	buf.WriteString(imad.InputSequenceNumberField())

	return buf.String()
}

// Validate performs WIRE format rule checks on InputMessageAccountabilityData and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (imad *InputMessageAccountabilityData) Validate() error {
	if err := imad.fieldInclusion(); err != nil {
		return err
	}
	if imad.tag != TagInputMessageAccountabilityData {
		return fieldError("tag", ErrValidTagForType, imad.tag)
	}
	if err := imad.validateDate(imad.InputCycleDate); err != nil {
		return fieldError("InputCycleDate", err, imad.InputCycleDate)
	}
	if err := imad.isAlphanumeric(imad.InputSource); err != nil {
		return fieldError("InputSource", err, imad.InputSource)
	}
	if err := imad.isNumeric(imad.InputSequenceNumber); err != nil {
		return fieldError("InputSequenceNumber", err, imad.InputSequenceNumber)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (imad *InputMessageAccountabilityData) fieldInclusion() error {
	if imad.InputCycleDate == "" {
		return fieldError("InputCycleDate", ErrFieldRequired, imad.InputCycleDate)
	}
	if imad.InputSource == "" {
		return fieldError("InputSource", ErrFieldRequired, imad.InputSource)
	}
	if imad.InputSequenceNumber == "" {
		return fieldError("InputSequenceNumber", ErrFieldRequired, imad.InputSequenceNumber)
	}
	return nil
}

// InputCycleDateField gets a string of the InputCycleDate field
func (imad *InputMessageAccountabilityData) InputCycleDateField() string {
	return imad.alphaVariableField(imad.InputCycleDate, 8, imad.isVariableLength)
}

// InputSourceField gets a string of the InputSource field
func (imad *InputMessageAccountabilityData) InputSourceField() string {
	return imad.alphaVariableField(imad.InputSource, 8, imad.isVariableLength)
}

// InputSequenceNumberField gets a string of the InputSequenceNumber field
func (imad *InputMessageAccountabilityData) InputSequenceNumberField() string {
	return imad.alphaVariableField(imad.InputSequenceNumber, 6, imad.isVariableLength)
}
