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

	value, read, err := omad.parseFixedStringField(record[length:], 8)
	if err != nil {
		return fieldError("OutputCycleDate", err)
	}
	omad.OutputCycleDate = value
	length += read

	value, read, err = omad.parseFixedStringField(record[length:], 8)
	if err != nil {
		return fieldError("OutputDestinationID", err)
	}
	omad.OutputDestinationID = value
	length += read

	if len(record) < length+6 {
		return fieldError("OutputSequenceNumber", ErrValidLength)
	}

	omad.OutputSequenceNumber = record[length : length+6]
	length += 6

	value, read, err = omad.parseFixedStringField(record[length:], 4)
	if err != nil {
		return fieldError("OutputDate", err)
	}
	omad.OutputDate = value
	length += read

	value, read, err = omad.parseFixedStringField(record[length:], 4)
	if err != nil {
		return fieldError("OutputTime", err)
	}
	omad.OutputTime = value
	length += read

	value, read, err = omad.parseFixedStringField(record[length:], 4)
	if err != nil {
		return fieldError("OutputFRBApplicationIdentification", err)
	}
	omad.OutputFRBApplicationIdentification = value
	length += read

	if err := omad.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width OutputMessageAccountabilityData record
func (omad *OutputMessageAccountabilityData) String() string {
	return omad.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a OutputMessageAccountabilityData record formatted according to the FormatOptions
func (omad *OutputMessageAccountabilityData) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(40)

	// All fields are fixed fields
	options.VariableLengthFields = false

	buf.WriteString(omad.tag)
	buf.WriteString(omad.OutputCycleDateField())
	buf.WriteString(omad.OutputDestinationIDField())
	buf.WriteString(omad.OutputSequenceNumberField())
	buf.WriteString(omad.OutputDateField())
	buf.WriteString(omad.OutputTimeField())
	buf.WriteString(omad.OutputFRBApplicationIdentificationField())

	if options.VariableLengthFields {
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

// OutputCycleDateField
func (omad *OutputMessageAccountabilityData) OutputCycleDateField() string {
	return omad.parseAlphaField(omad.OutputCycleDate, 8)
}

// FormatOutputCycleDate returns OutputCycleDate formatted according to the FormatOptions
func (omad *OutputMessageAccountabilityData) FormatOutputCycleDate(options FormatOptions) string {
	return omad.formatAlphaField(omad.OutputCycleDate, 8, options)
}

// OutputDestinationIDField
func (omad *OutputMessageAccountabilityData) OutputDestinationIDField() string {
	return omad.parseAlphaField(omad.OutputDestinationID, 8)
}

// FormatOutputDestinationID returns OutputDestinationID formatted according to the FormatOptions
func (omad *OutputMessageAccountabilityData) FormatOutputDestinationID(options FormatOptions) string {
	return omad.formatAlphaField(omad.OutputDestinationID, 8, options)
}

// OutputSequenceNumberField gets a string of the OutputSequenceNumber field
func (omad *OutputMessageAccountabilityData) OutputSequenceNumberField() string {
	return omad.numericStringField(omad.OutputSequenceNumber, 6)
}

// OutputDateField
func (omad *OutputMessageAccountabilityData) OutputDateField() string {
	return omad.parseAlphaField(omad.OutputDate, 4)
}

// FormatOutputDate returns OutputDate formatted according to the FormatOptions
func (omad *OutputMessageAccountabilityData) FormatOutputDate(options FormatOptions) string {
	return omad.formatAlphaField(omad.OutputDate, 4, options)
}

// OutputTimeField
func (omad *OutputMessageAccountabilityData) OutputTimeField() string {
	return omad.parseAlphaField(omad.OutputTime, 4)
}

// FormatOutputTime returns OutputTime formatted according to the FormatOptions
func (omad *OutputMessageAccountabilityData) FormatOutputTime(options FormatOptions) string {
	return omad.formatAlphaField(omad.OutputTime, 4, options)
}

// OutputFRBApplicationIdentificationField
func (omad *OutputMessageAccountabilityData) OutputFRBApplicationIdentificationField() string {
	return omad.parseAlphaField(omad.OutputFRBApplicationIdentification, 4)
}

// FormatOutputFRBApplicationIdentification returns OutputFRBApplicationIdentification formatted according to the FormatOptions
func (omad *OutputMessageAccountabilityData) FormatOutputFRBApplicationIdentification(options FormatOptions) string {
	return omad.formatAlphaField(omad.OutputFRBApplicationIdentification, 4, options)
}
