// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &ServiceMessage{}

// ServiceMessage is the ServiceMessage of the wire
type ServiceMessage struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// LineOne
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree
	LineThree string `json:"lineThree,omitempty"`
	// LineFour
	LineFour string `json:"lineFour,omitempty"`
	// LineFive
	LineFive string `json:"lineFive,omitempty"`
	// LineSix
	LineSix string `json:"lineSix,omitempty"`
	// LineSeven
	LineSeven string `json:"lineSeven,omitempty"`
	// LineEight
	LineEight string `json:"lineEight,omitempty"`
	// LineNine
	LineNine string `json:"lineNine,omitempty"`
	// LineTen
	LineTen string `json:"lineTen,omitempty"`
	// LineEleven
	LineEleven string `json:"lineEleven,omitempty"`
	// LineTwelve
	LineTwelve string `json:"lineTwelve,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewServiceMessage returns a new ServiceMessage
func NewServiceMessage(isVariable bool) *ServiceMessage {
	sm := &ServiceMessage{
		tag:              TagServiceMessage,
		isVariableLength: isVariable,
	}
	return sm
}

// Parse takes the input string and parses the ServiceMessage values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sm *ServiceMessage) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 18 {
		return 0, NewTagWrongLengthErr(18, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if sm.tag, read, err = sm.parseTag(record); err != nil {
		return 0, fieldError("ServiceMessage.Tag", err)
	}
	length += read

	if sm.LineOne, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineOne", err)
	}
	length += read

	if sm.LineTwo, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineTwo", err)
	}
	length += read

	if sm.LineThree, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineThree", err)
	}
	length += read

	if sm.LineFour, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineFour", err)
	}
	length += read

	if sm.LineFive, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineFive", err)
	}
	length += read

	if sm.LineSix, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineSix", err)
	}
	length += read

	if sm.LineSeven, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineSeven", err)
	}
	length += read

	if sm.LineEight, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineEight", err)
	}
	length += read

	if sm.LineNine, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineNine", err)
	}
	length += read

	if sm.LineTen, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineTen", err)
	}
	length += read

	if sm.LineEleven, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineEleven", err)
	}
	length += read

	if sm.LineTwelve, read, err = sm.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("LineTwelve", err)
	}
	length += read

	return length, nil
}

func (sm *ServiceMessage) UnmarshalJSON(data []byte) error {
	type Alias ServiceMessage
	aux := struct {
		*Alias
	}{
		(*Alias)(sm),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	sm.tag = TagServiceMessage
	return nil
}

// String writes ServiceMessage
func (sm *ServiceMessage) String() string {
	var buf strings.Builder
	buf.Grow(426)

	buf.WriteString(sm.tag)
	buf.WriteString(sm.LineOneField())
	buf.WriteString(sm.LineTwoField())
	buf.WriteString(sm.LineThreeField())
	buf.WriteString(sm.LineFourField())
	buf.WriteString(sm.LineFiveField())
	buf.WriteString(sm.LineSixField())
	buf.WriteString(sm.LineSevenField())
	buf.WriteString(sm.LineEightField())
	buf.WriteString(sm.LineNineField())
	buf.WriteString(sm.LineTenField())
	buf.WriteString(sm.LineElevenField())
	buf.WriteString(sm.LineTwelveField())

	return buf.String()
}

// Validate performs WIRE format rule checks on ServiceMessage and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sm *ServiceMessage) Validate() error {
	if err := sm.fieldInclusion(); err != nil {
		return err
	}
	if sm.tag != TagServiceMessage {
		return fieldError("tag", ErrValidTagForType, sm.tag)
	}
	if err := sm.isAlphanumeric(sm.LineOne); err != nil {
		return fieldError("LineOne", err, sm.LineOne)
	}
	if err := sm.isAlphanumeric(sm.LineTwo); err != nil {
		return fieldError("LineTwo", err, sm.LineTwo)
	}
	if err := sm.isAlphanumeric(sm.LineThree); err != nil {
		return fieldError("LineThree", err, sm.LineThree)
	}
	if err := sm.isAlphanumeric(sm.LineFour); err != nil {
		return fieldError("LineFour", err, sm.LineFour)
	}
	if err := sm.isAlphanumeric(sm.LineFive); err != nil {
		return fieldError("LineFive", err, sm.LineFive)
	}
	if err := sm.isAlphanumeric(sm.LineSix); err != nil {
		return fieldError("LineSix", err, sm.LineSix)
	}
	if err := sm.isAlphanumeric(sm.LineSeven); err != nil {
		return fieldError("LineSeven", err, sm.LineSeven)
	}
	if err := sm.isAlphanumeric(sm.LineEight); err != nil {
		return fieldError("LineEight", err, sm.LineEight)
	}
	if err := sm.isAlphanumeric(sm.LineNine); err != nil {
		return fieldError("LineNine", err, sm.LineNine)
	}
	if err := sm.isAlphanumeric(sm.LineTen); err != nil {
		return fieldError("LineTen", err, sm.LineTen)
	}
	if err := sm.isAlphanumeric(sm.LineEleven); err != nil {
		return fieldError("LineEleven", err, sm.LineEleven)
	}
	if err := sm.isAlphanumeric(sm.LineTwelve); err != nil {
		return fieldError("LineTwelve", err, sm.LineTwelve)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sm *ServiceMessage) fieldInclusion() error {
	// If ServiceMessage is defined, LineOne is required
	if sm.LineOne == "" {
		return fieldError("LineOne", ErrFieldRequired)
	}

	return nil
}

// LineOneField gets a string of the LineOne field
func (sm *ServiceMessage) LineOneField() string {
	return sm.alphaVariableField(sm.LineOne, 35, sm.isVariableLength)
}

// LineTwoField gets a string of the LineTwo field
func (sm *ServiceMessage) LineTwoField() string {
	return sm.alphaVariableField(sm.LineTwo, 35, sm.isVariableLength)
}

// LineThreeField gets a string of the LineThree field
func (sm *ServiceMessage) LineThreeField() string {
	return sm.alphaVariableField(sm.LineThree, 35, sm.isVariableLength)
}

// LineFourField gets a string of the LineFour field
func (sm *ServiceMessage) LineFourField() string {
	return sm.alphaVariableField(sm.LineFour, 35, sm.isVariableLength)
}

// LineFiveField gets a string of the LineFive field
func (sm *ServiceMessage) LineFiveField() string {
	return sm.alphaVariableField(sm.LineFive, 35, sm.isVariableLength)
}

// LineSixField gets a string of the LineSix field
func (sm *ServiceMessage) LineSixField() string {
	return sm.alphaVariableField(sm.LineSix, 35, sm.isVariableLength)
}

// LineSevenField gets a string of the LineSeven field
func (sm *ServiceMessage) LineSevenField() string {
	return sm.alphaVariableField(sm.LineSeven, 35, sm.isVariableLength)
}

// LineEightField gets a string of the LineEight field
func (sm *ServiceMessage) LineEightField() string {
	return sm.alphaVariableField(sm.LineEight, 35, sm.isVariableLength)
}

// LineNineField gets a string of the LineNine field
func (sm *ServiceMessage) LineNineField() string {
	return sm.alphaVariableField(sm.LineNine, 35, sm.isVariableLength)
}

// LineTenField gets a string of the LineTen field
func (sm *ServiceMessage) LineTenField() string {
	return sm.alphaVariableField(sm.LineTen, 35, sm.isVariableLength)
}

// LineElevenField gets a string of the LineEleven field
func (sm *ServiceMessage) LineElevenField() string {
	return sm.alphaVariableField(sm.LineEleven, 35, sm.isVariableLength)
}

// LineTwelveField gets a string of the LineTwelve field
func (sm *ServiceMessage) LineTwelveField() string {
	return sm.alphaVariableField(sm.LineTwelve, 35, sm.isVariableLength)
}
