// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (omad *OutputMessageAccountabilityData) Parse(record string) {
	omad.tag = record[:6]
	omad.OutputCycleDate = omad.parseStringField(record[6:14])
	omad.OutputDestinationID = omad.parseStringField(record[14:22])
	omad.OutputSequenceNumber = omad.parseStringField(record[22:28])
	omad.OutputDate = omad.parseStringField(record[28:32])
	omad.OutputTime = omad.parseStringField(record[32:36])
	omad.OutputFRBApplicationIdentification = omad.parseStringField(record[36:40])
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
	return omad.alphaField(omad.OutputCycleDate, 8)
}

// OutputDestinationIDField gets a string of the OutputDestinationID field
func (omad *OutputMessageAccountabilityData) OutputDestinationIDField() string {
	return omad.alphaField(omad.OutputDestinationID, 8)
}

// OutputSequenceNumberField gets a string of the OutputSequenceNumber field
func (omad *OutputMessageAccountabilityData) OutputSequenceNumberField() string {
	return omad.numericStringField(omad.OutputSequenceNumber, 6)
}

// OutputDateField gets a string of the OutputDate field
func (omad *OutputMessageAccountabilityData) OutputDateField() string {
	return omad.alphaField(omad.OutputDate, 4)
}

// OutputTimeField gets a string of the OutputTime field
func (omad *OutputMessageAccountabilityData) OutputTimeField() string {
	return omad.alphaField(omad.OutputTime, 4)
}

// OutputFRBApplicationIdentificationField gets a string of the OutputFRBApplicationIdentification field
func (omad *OutputMessageAccountabilityData) OutputFRBApplicationIdentificationField() string {
	return omad.alphaField(omad.OutputFRBApplicationIdentification, 4)
}
