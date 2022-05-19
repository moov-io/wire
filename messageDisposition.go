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
func NewMessageDisposition() *MessageDisposition {
	md := &MessageDisposition{
		tag:                    TagMessageDisposition,
		FormatVersion:          FormatVersion,
		TestProductionCode:     EnvironmentProduction,
		MessageDuplicationCode: MessageDuplicationOriginal,
	}
	return md
}

// Parse takes the input string and parses the MessageDisposition values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (md *MessageDisposition) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 10 {
		return 0, NewTagWrongLengthErr(10, len(record))
	}

	var err error
	var length, read int

	if md.tag, read, err = md.parseTag(record); err != nil {
		return 0, fieldError("MessageDisposition.Tag", err)
	}
	length += read

	if md.FormatVersion, read, err = md.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("FormatVersion", err)
	}
	length += read

	if md.TestProductionCode, read, err = md.parseVariableStringField(record[length:], 1); err != nil {
		return 0, fieldError("TestProductionCode", err)
	}
	length += read

	if md.MessageDuplicationCode, read, err = md.parseVariableStringField(record[length:], 1); err != nil {
		return 0, fieldError("MessageDuplicationCode", err)
	}
	length += read

	if md.MessageStatusIndicator, read, err = md.parseVariableStringField(record[length:], 1); err != nil {
		return 0, fieldError("MessageStatusIndicator", err)
	}
	length += read

	return length, nil
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
func (md *MessageDisposition) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(11)

	buf.WriteString(md.tag)
	buf.WriteString(md.MessageDispositionFormatVersionField(isCompressed))
	buf.WriteString(md.MessageDispositionTestProductionCodeField(isCompressed))
	buf.WriteString(md.MessageDispositionMessageDuplicationCodeField(isCompressed))
	buf.WriteString(md.MessageDispositionMessageStatusIndicatorField(isCompressed))

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
func (md *MessageDisposition) MessageDispositionFormatVersionField(isCompressed bool) string {
	return md.alphaVariableField(md.FormatVersion, 2, isCompressed)
}

// MessageDispositionTestProductionCodeField gets a string of the TestProductionCoden field
func (md *MessageDisposition) MessageDispositionTestProductionCodeField(isCompressed bool) string {
	return md.alphaVariableField(md.TestProductionCode, 1, isCompressed)
}

// MessageDispositionMessageDuplicationCodeField gets a string of the MessageDuplicationCode field
func (md *MessageDisposition) MessageDispositionMessageDuplicationCodeField(isCompressed bool) string {
	return md.alphaVariableField(md.MessageDuplicationCode, 1, isCompressed)
}

// MessageDispositionMessageStatusIndicatorField gets a string of the MessageDuplicationCode field
func (md *MessageDisposition) MessageDispositionMessageStatusIndicatorField(isCompressed bool) string {
	return md.alphaVariableField(md.MessageStatusIndicator, 1, isCompressed)
}
