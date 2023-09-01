// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func NewServiceMessage() *ServiceMessage {
	sm := &ServiceMessage{
		tag: TagServiceMessage,
	}
	return sm
}

// Parse takes the input string and parses the ServiceMessage values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sm *ServiceMessage) Parse(record string) error {
	if utf8.RuneCountInString(record) < 8 {
		return NewTagMinLengthErr(8, len(record))
	}

	sm.tag = record[:6]
	length := 6

	value, read, err := sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineOne", err)
	}
	sm.LineOne = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	sm.LineTwo = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineThree", err)
	}
	sm.LineThree = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineFour", err)
	}
	sm.LineFour = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineFive", err)
	}
	sm.LineFive = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineSix", err)
	}
	sm.LineSix = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineSeven", err)
	}
	sm.LineSeven = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineEight", err)
	}
	sm.LineEight = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineNine", err)
	}
	sm.LineNine = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineTen", err)
	}
	sm.LineTen = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineEleven", err)
	}
	sm.LineEleven = value
	length += read

	value, read, err = sm.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("LineTwelve", err)
	}
	sm.LineTwelve = value
	length += read

	if err := sm.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
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

// String returns a fixed-width ServiceMessage record
func (sm *ServiceMessage) String() string {
	return sm.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a ServiceMessage record formatted according to the FormatOptions
func (sm *ServiceMessage) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(426)

	buf.WriteString(sm.tag)
	buf.WriteString(sm.FormatLineOne(options) + Delimiter)
	buf.WriteString(sm.FormatLineTwo(options) + Delimiter)
	buf.WriteString(sm.FormatLineThree(options) + Delimiter)
	buf.WriteString(sm.FormatLineFour(options) + Delimiter)
	buf.WriteString(sm.FormatLineFive(options) + Delimiter)
	buf.WriteString(sm.FormatLineSix(options) + Delimiter)
	buf.WriteString(sm.FormatLineSeven(options) + Delimiter)
	buf.WriteString(sm.FormatLineEight(options) + Delimiter)
	buf.WriteString(sm.FormatLineNine(options) + Delimiter)
	buf.WriteString(sm.FormatLineTen(options) + Delimiter)
	buf.WriteString(sm.FormatLineEleven(options) + Delimiter)
	buf.WriteString(sm.FormatLineTwelve(options) + Delimiter)

	if options.VariableLengthFields {
		return sm.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
	return sm.alphaField(sm.LineOne, 35)
}

// LineTwoField gets a string of the LineTwo field
func (sm *ServiceMessage) LineTwoField() string {
	return sm.alphaField(sm.LineTwo, 35)
}

// LineThreeField gets a string of the LineThree field
func (sm *ServiceMessage) LineThreeField() string {
	return sm.alphaField(sm.LineThree, 35)
}

// LineFourField gets a string of the LineFour field
func (sm *ServiceMessage) LineFourField() string {
	return sm.alphaField(sm.LineFour, 35)
}

// LineFiveField gets a string of the LineFive field
func (sm *ServiceMessage) LineFiveField() string {
	return sm.alphaField(sm.LineFive, 35)
}

// LineSixField gets a string of the LineSix field
func (sm *ServiceMessage) LineSixField() string {
	return sm.alphaField(sm.LineSix, 35)
}

// LineSevenField gets a string of the LineSeven field
func (sm *ServiceMessage) LineSevenField() string {
	return sm.alphaField(sm.LineSeven, 35)
}

// LineEightField gets a string of the LineEight field
func (sm *ServiceMessage) LineEightField() string {
	return sm.alphaField(sm.LineEight, 35)
}

// LineNineField gets a string of the LineNine field
func (sm *ServiceMessage) LineNineField() string {
	return sm.alphaField(sm.LineNine, 35)
}

// LineTenField gets a string of the LineTen field
func (sm *ServiceMessage) LineTenField() string {
	return sm.alphaField(sm.LineTen, 35)
}

// LineElevenField gets a string of the LineEleven field
func (sm *ServiceMessage) LineElevenField() string {
	return sm.alphaField(sm.LineEleven, 35)
}

// LineTwelveField gets a string of the LineTwelve field
func (sm *ServiceMessage) LineTwelveField() string {
	return sm.alphaField(sm.LineTwelve, 35)
}

// FormatLineOne returns LineOne formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineOne(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineOne, 35, options)
}

// FormatLineTwo returns LineTwo formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineTwo(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineTwo, 35, options)
}

// FormatLineThree returns LineThree formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineThree(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineThree, 35, options)
}

// FormatLineFour returns LineFour formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineFour(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineFour, 35, options)
}

// FormatLineFive returns LineFive formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineFive(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineFive, 35, options)
}

// FormatLineSix returns LineSix formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineSix(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineSix, 35, options)
}

// FormatLineSeven returns LineSeven formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineSeven(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineSeven, 35, options)
}

// FormatLineEight returns LineEight formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineEight(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineEight, 35, options)
}

// FormatLineNine returns LineNine formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineNine(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineNine, 35, options)
}

// FormatLineTen returns LineTen formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineTen(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineTen, 35, options)
}

// FormatLineEleven returns LineEleven formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineEleven(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineEleven, 35, options)
}

// FormatLineTwelve returns LineTwelve formatted according to the FormatOptions
func (sm *ServiceMessage) FormatLineTwelve(options FormatOptions) string {
	return sm.formatAlphaField(sm.LineTwelve, 35, options)
}
