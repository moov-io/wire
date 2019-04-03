// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// ServiceMessage is the ServiceMessage of the wire
type ServiceMessage struct {
	// tag
	tag string
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
func NewServiceMessage() ServiceMessage {
	sm := ServiceMessage{
		tag: TagServiceMessage,
	}
	return sm
}

// Parse takes the input string and parses the ServiceMessage values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sm *ServiceMessage) Parse(record string) {
	sm.tag = record[:6]
	sm.LineOne = sm.parseStringField(record[6:41])
	sm.LineTwo = sm.parseStringField(record[41:76])
	sm.LineThree = sm.parseStringField(record[76:111])
	sm.LineFour = sm.parseStringField(record[111:146])
	sm.LineFive = sm.parseStringField(record[146:181])
	sm.LineSix = sm.parseStringField(record[181:216])
	sm.LineSeven = sm.parseStringField(record[216:251])
	sm.LineEight = sm.parseStringField(record[251:286])
	sm.LineNine = sm.parseStringField(record[286:321])
	sm.LineTen = sm.parseStringField(record[321:356])
	sm.LineEleven = sm.parseStringField(record[356:391])
	sm.LineTwelve = sm.parseStringField(record[391:426])
}

// String writes ServiceMessage
func (sm *ServiceMessage) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(426)
	buf.WriteString(sm.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ServiceMessage and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sm *ServiceMessage) Validate() error {
	if err := sm.fieldInclusion(); err != nil {
		return err
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
	return nil
}
