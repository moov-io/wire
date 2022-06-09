// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// OutputMessageAccountabilityData is the Output Message Accountability Data (OMAD) of the wire
type OutputMessageAccountabilityData struct {
	// tag
	tag string
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
func NewOutputMessageAccountabilityData() *OutputMessageAccountabilityData {
	omad := &OutputMessageAccountabilityData{
		tag: TagOutputMessageAccountabilityData,
	}
	return omad
}

// Parse takes the input string and parses the OutputMessageAccountabilityData values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (omad *OutputMessageAccountabilityData) Parse(record string) error {
	if utf8.RuneCountInString(record) < 14 {
		return NewTagMinLengthErr(14, len(record))
	}

	omad.tag = record[:6]
	length := 6

	value, read, err := omad.parseVariableStringField(record[length:], 8)
	if err != nil {
		return fieldError("OutputCycleDate", err)
	}
	omad.OutputCycleDate = value
	length += read

	value, read, err = omad.parseVariableStringField(record[length:], 8)
	if err != nil {
		return fieldError("OutputDestinationID", err)
	}
	omad.OutputDestinationID = value
	length += read

	if len(record) < length+6 {
		return fieldError("OutputSequenceNumber", ErrValidLengthSize)
	}

	omad.OutputSequenceNumber = record[length : length+6]
	length += 6

	value, read, err = omad.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("OutputDate", err)
	}
	omad.OutputDate = value
	length += read

	value, read, err = omad.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("OutputTime", err)
	}
	omad.OutputTime = value
	length += read

	value, read, err = omad.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("OutputFRBApplicationIdentification", err)
	}
	omad.OutputFRBApplicationIdentification = value
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
	}

	return nil
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
func (omad *OutputMessageAccountabilityData) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(40)

	buf.WriteString(omad.tag)
	buf.WriteString(omad.OutputCycleDateField(options...))
	buf.WriteString(omad.OutputDestinationIDField(options...))
	buf.WriteString(omad.OutputSequenceNumberField())
	buf.WriteString(omad.OutputDateField(options...))
	buf.WriteString(omad.OutputTimeField(options...))
	buf.WriteString(omad.OutputFRBApplicationIdentificationField(options...))

	if omad.parseFirstOption(options) {
		return omad.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
func (omad *OutputMessageAccountabilityData) OutputCycleDateField(options ...bool) string {
	return omad.alphaVariableField(omad.OutputCycleDate, 8, omad.parseFirstOption(options))
}

// OutputDestinationIDField gets a string of the OutputDestinationID field
func (omad *OutputMessageAccountabilityData) OutputDestinationIDField(options ...bool) string {
	return omad.alphaVariableField(omad.OutputDestinationID, 8, omad.parseFirstOption(options))
}

// OutputSequenceNumberField gets a string of the OutputSequenceNumber field
func (omad *OutputMessageAccountabilityData) OutputSequenceNumberField() string {
	return omad.numericStringField(omad.OutputSequenceNumber, 6)
}

// OutputDateField gets a string of the OutputDate field
func (omad *OutputMessageAccountabilityData) OutputDateField(options ...bool) string {
	return omad.alphaVariableField(omad.OutputDate, 4, omad.parseFirstOption(options))
}

// OutputTimeField gets a string of the OutputTime field
func (omad *OutputMessageAccountabilityData) OutputTimeField(options ...bool) string {
	return omad.alphaVariableField(omad.OutputTime, 4, omad.parseFirstOption(options))
}

// OutputFRBApplicationIdentificationField gets a string of the OutputFRBApplicationIdentification field
func (omad *OutputMessageAccountabilityData) OutputFRBApplicationIdentificationField(options ...bool) string {
	return omad.alphaVariableField(omad.OutputFRBApplicationIdentification, 4, omad.parseFirstOption(options))
}
