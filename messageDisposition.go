// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &MessageDisposition{}

// MessageDisposition is the message disposition of the wire
type MessageDisposition struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// FormatVersion 30
	FormatVersion string `json:"formatVersion,omitempty"`
	// TestTestProductionCode identifies if test or production
	TestProductionCode string `json:"testProductionCode,omitempty"`
	// MessageDuplicationCode  * ` ` - Original Message * `R` - Retrieval of an original message * `P` - Resend
	MessageDuplicationCode string `json:"messageDuplicationCode,omitempty"`
	// MessageStatusIndicator
	MessageStatusIndicator string `json:"messageStatusIndicator,omitempty"`

	// validator is composed for data validation
	// validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewMessageDisposition returns a new MessageDisposition
func NewMessageDisposition(isVariable bool) *MessageDisposition {
	md := &MessageDisposition{
		tag:                    TagMessageDisposition,
		FormatVersion:          FormatVersion,
		TestProductionCode:     EnvironmentProduction,
		MessageDuplicationCode: MessageDuplicationOriginal,
		isVariableLength:       isVariable,
	}
	return md
}

// Parse takes the input string and parses the MessageDisposition values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (md *MessageDisposition) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 10 {
		return NewTagWrongLengthErr(10, len(record)), 0
	}

	md.tag = record[:6]

	length := 6
	read := 0

	md.FormatVersion, read = md.parseVariableStringField(record[length:], 2)
	length += read

	md.TestProductionCode, read = md.parseVariableStringField(record[length:], 1)
	length += read

	md.MessageDuplicationCode, read = md.parseVariableStringField(record[length:], 1)
	length += read

	md.MessageStatusIndicator, read = md.parseVariableStringField(record[length:], 1)
	length += read

	return nil, length + 3
}

func (md *MessageDisposition) UnmarshalJSON(data []byte) error {
	type Alias MessageDisposition
	aux := struct {
		*Alias
	}{
		(*Alias)(md),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	md.tag = TagMessageDisposition
	return nil
}

// String writes MessageDisposition
func (md *MessageDisposition) String() string {
	var buf strings.Builder
	buf.Grow(11)

	buf.WriteString(md.tag)
	buf.WriteString(md.MessageDispositionFormatVersionField())
	buf.WriteString(md.MessageDispositionTestProductionCodeField())
	buf.WriteString(md.MessageDispositionMessageDuplicationCodeField())
	buf.WriteString(md.MessageDispositionMessageStatusIndicatorField())

	return buf.String()
}

// Validate performs WIRE format rule checks on MessageDisposition and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (md *MessageDisposition) Validate() error {
	// Currently no validation as the FED is responsible for the values
	if md.tag != TagMessageDisposition {
		return fieldError("tag", ErrValidTagForType, md.tag)
	}
	return nil
}

// MessageDispositionFormatVersionField gets a string of the FormatVersion field
func (md *MessageDisposition) MessageDispositionFormatVersionField() string {
	return md.alphaVariableField(md.FormatVersion, 2, md.isVariableLength)
}

// MessageDispositionTestProductionCodeField gets a string of the TestProductionCoden field
func (md *MessageDisposition) MessageDispositionTestProductionCodeField() string {
	return md.alphaVariableField(md.TestProductionCode, 1, md.isVariableLength)
}

// MessageDispositionMessageDuplicationCodeField gets a string of the MessageDuplicationCode field
func (md *MessageDisposition) MessageDispositionMessageDuplicationCodeField() string {
	return md.alphaVariableField(md.MessageDuplicationCode, 1, md.isVariableLength)
}

// MessageDispositionMessageStatusIndicatorField gets a string of the MessageDuplicationCode field
func (md *MessageDisposition) MessageDispositionMessageStatusIndicatorField() string {
	return md.alphaVariableField(md.MessageStatusIndicator, 1, md.isVariableLength)
}
