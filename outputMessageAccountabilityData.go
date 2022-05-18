// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &OutputMessageAccountabilityData{}

// OutputMessageAccountabilityData is the Output Message Accountability Data (OMAD) of the wire
type OutputMessageAccountabilityData struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// OutputCycleDate (CCYYMMDD)
	OutputCycleDate string `json:"outputCycleDate,omitempty"`
	// OutputDestinationID
	OutputDestinationID string `json:"outputDestinationID,omitempty"`
	// OutputOutputSequenceNumber
	OutputSequenceNumber string `json:"outputSequenceNumber,omitempty"`
	// OutputDate is the output date
	OutputDate string `json:"outputDate,omitempty"`
	// OutputTime is OutputTime
	OutputTime string `json:"outputTime,omitempty"`
	// OutputFRBApplicationIdentification
	OutputFRBApplicationIdentification string `json:"outputFRBApplicationIdentification,omitempty"`

	// validator is composed for data validation
	// validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewOutputMessageAccountabilityData returns a new OutputMessageAccountabilityData
func NewOutputMessageAccountabilityData(isVariable bool) *OutputMessageAccountabilityData {
	omad := &OutputMessageAccountabilityData{
		tag:              TagOutputMessageAccountabilityData,
		isVariableLength: isVariable,
	}
	return omad
}

// Parse takes the input string and parses the OutputMessageAccountabilityData values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (omad *OutputMessageAccountabilityData) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 12 {
		return 0, NewTagWrongLengthErr(12, len(record))
	}

	var err error
	var length, read int

	if omad.tag, read, err = omad.parseTag(record); err != nil {
		return 0, fieldError("OutputMessageAccountabilityData.Tag", err)
	}
	length += read

	if omad.OutputCycleDate, read, err = omad.parseVariableStringField(record[length:], 8); err != nil {
		return 0, fieldError("OutputCycleDate", err)
	}
	length += read

	if omad.OutputDestinationID, read, err = omad.parseVariableStringField(record[length:], 8); err != nil {
		return 0, fieldError("OutputDestinationID", err)
	}
	length += read

	if omad.OutputSequenceNumber, read, err = omad.parseVariableStringField(record[length:], 6); err != nil {
		return 0, fieldError("OutputSequenceNumber", err)
	}
	length += read

	if omad.OutputDate, read, err = omad.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("OutputDate", err)
	}
	length += read

	if omad.OutputTime, read, err = omad.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("OutputTime", err)
	}
	length += read

	if omad.OutputFRBApplicationIdentification, read, err = omad.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("OutputFRBApplicationIdentification", err)
	}
	length += read

	return length, nil
}

func (omad *OutputMessageAccountabilityData) UnmarshalJSON(data []byte) error {
	type Alias OutputMessageAccountabilityData
	aux := struct {
		*Alias
	}{
		(*Alias)(omad),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	omad.tag = TagOutputMessageAccountabilityData
	return nil
}

// String writes OutputMessageAccountabilityData
func (omad *OutputMessageAccountabilityData) String() string {
	var buf strings.Builder
	buf.Grow(40)

	buf.WriteString(omad.tag)
	buf.WriteString(omad.OutputCycleDateField())
	buf.WriteString(omad.OutputDestinationIDField())
	buf.WriteString(omad.OutputSequenceNumberField())
	buf.WriteString(omad.OutputDateField())
	buf.WriteString(omad.OutputTimeField())
	buf.WriteString(omad.OutputFRBApplicationIdentificationField())

	return buf.String()
}

// Validate performs WIRE format rule checks on OutputMessageAccountabilityData and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (omad *OutputMessageAccountabilityData) Validate() error {
	// Currently no validation as the FED is responsible for the values
	if omad.tag != TagOutputMessageAccountabilityData {
		return fieldError("tag", ErrValidTagForType, omad.tag)
	}
	return nil
}

// OutputCycleDateField gets a string of the OutputCycleDate field
func (omad *OutputMessageAccountabilityData) OutputCycleDateField() string {
	return omad.alphaVariableField(omad.OutputCycleDate, 8, omad.isVariableLength)
}

// OutputDestinationIDField gets a string of the OutputDestinationID field
func (omad *OutputMessageAccountabilityData) OutputDestinationIDField() string {
	return omad.alphaField(omad.OutputDestinationID, 8)
	return omad.alphaVariableField(omad.OutputDestinationID, 8, omad.isVariableLength)
}

// OutputSequenceNumberField gets a string of the OutputSequenceNumber field
func (omad *OutputMessageAccountabilityData) OutputSequenceNumberField() string {
	return omad.alphaVariableField(omad.OutputSequenceNumber, 6, omad.isVariableLength)
}

// OutputDateField gets a string of the OutputDate field
func (omad *OutputMessageAccountabilityData) OutputDateField() string {
	return omad.alphaVariableField(omad.OutputDate, 4, omad.isVariableLength)
}

// OutputTimeField gets a string of the OutputTime field
func (omad *OutputMessageAccountabilityData) OutputTimeField() string {
	return omad.alphaVariableField(omad.OutputTime, 4, omad.isVariableLength)
}

// OutputFRBApplicationIdentificationField gets a string of the OutputFRBApplicationIdentification field
func (omad *OutputMessageAccountabilityData) OutputFRBApplicationIdentificationField() string {
	return omad.alphaVariableField(omad.OutputFRBApplicationIdentification, 4, omad.isVariableLength)
}
