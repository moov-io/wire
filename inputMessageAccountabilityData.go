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
func NewInputMessageAccountabilityData() *InputMessageAccountabilityData {
	imad := &InputMessageAccountabilityData{
		tag: TagInputMessageAccountabilityData,
	}
	return imad
}

// Parse takes the input string and parses the InputMessageAccountabilityData values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (imad *InputMessageAccountabilityData) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 9 {
		return 0, NewTagWrongLengthErr(9, len(record))
	}

	var err error
	var length, read int

	if imad.tag, read, err = imad.parseTag(record); err != nil {
		return 0, fieldError("InputMessageAccountabilityData.Tag", err)
	}
	length += read

	if imad.InputCycleDate, read, err = imad.parseVariableStringField(record[length:], 8); err != nil {
		return 0, fieldError("InputCycleDate", err)
	}
	length += read

	if imad.InputSource, read, err = imad.parseVariableStringField(record[length:], 8); err != nil {
		return 0, fieldError("InputSource", err)
	}
	length += read

	if imad.InputSequenceNumber, read, err = imad.parseVariableStringField(record[length:], 6); err != nil {
		return 0, fieldError("InputSequenceNumber", err)
	}
	length += read

	return length, nil
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
func (imad *InputMessageAccountabilityData) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(imad.tag)
	buf.WriteString(imad.InputCycleDateField(isCompressed))
	buf.WriteString(imad.InputSourceField(isCompressed))
	buf.WriteString(imad.InputSequenceNumberField(isCompressed))

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
func (imad *InputMessageAccountabilityData) InputCycleDateField(isCompressed bool) string {
	return imad.alphaVariableField(imad.InputCycleDate, 8, isCompressed)
}

// InputSourceField gets a string of the InputSource field
func (imad *InputMessageAccountabilityData) InputSourceField(isCompressed bool) string {
	return imad.alphaVariableField(imad.InputSource, 8, isCompressed)
}

// InputSequenceNumberField gets a string of the InputSequenceNumber field
func (imad *InputMessageAccountabilityData) InputSequenceNumberField(isCompressed bool) string {
	return imad.alphaVariableField(imad.InputSequenceNumber, 6, isCompressed)
}
