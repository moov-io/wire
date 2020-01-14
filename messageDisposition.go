// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func (md *MessageDisposition) Parse(record string) {
	md.tag = record[:6]
	md.FormatVersion = md.parseStringField(record[6:8])
	md.TestProductionCode = md.parseStringField(record[8:9])
	md.MessageDuplicationCode = md.parseStringField(record[9:10])
	md.MessageStatusIndicator = md.parseStringField(record[10:11])
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
	return md.alphaField(md.FormatVersion, 2)
}

// MessageDispositionTestProductionCodeField gets a string of the TestProductionCoden field
func (md *MessageDisposition) MessageDispositionTestProductionCodeField() string {
	return md.alphaField(md.TestProductionCode, 1)
}

// MessageDispositionMessageDuplicationCodeField gets a string of the MessageDuplicationCode field
func (md *MessageDisposition) MessageDispositionMessageDuplicationCodeField() string {
	return md.alphaField(md.MessageDuplicationCode, 1)
}

// MessageDispositionMessageStatusIndicatorField gets a string of the MessageDuplicationCode field
func (md *MessageDisposition) MessageDispositionMessageStatusIndicatorField() string {
	return md.alphaField(md.MessageStatusIndicator, 1)
}
