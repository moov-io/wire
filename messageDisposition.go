// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (md *MessageDisposition) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	md.tag = record[:6]
	length := 6

	value, read, err := md.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("FormatVersion", err)
	}
	md.FormatVersion = value
	length += read

	value, read, err = md.parseVariableStringField(record[length:], 1)
	if err != nil {
		return fieldError("TestProductionCode", err)
	}
	md.TestProductionCode = value
	length += read

	value, read, err = md.parseVariableStringField(record[length:], 1)
	if err != nil {
		return fieldError("MessageDuplicationCode", err)
	}
	md.MessageDuplicationCode = value
	length += read

	value, read, err = md.parseVariableStringField(record[length:], 1)
	if err != nil {
		return fieldError("MessageStatusIndicator", err)
	}
	md.MessageStatusIndicator = value
	length += read

	if !md.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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

// String returns a fixed-width MessageDisposition record
func (md *MessageDisposition) String() string {
	return md.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a MessageDisposition record formatted according to the FormatOptions
func (md *MessageDisposition) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(11)

	buf.WriteString(md.tag)
	buf.WriteString(md.FormatMessageDispositionFormatVersion(options))
	buf.WriteString(md.FormatMessageDispositionTestProductionCode(options))
	buf.WriteString(md.FormatMessageDispositionMessageDuplicationCode(options))
	buf.WriteString(md.FormatMessageDispositionMessageStatusIndicator(options))

	if options.VariableLengthFields {
		return md.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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

// FormatMessageDispositionFormatVersion returns FormatVersion formatted according to the FormatOptions
func (md *MessageDisposition) FormatMessageDispositionFormatVersion(options FormatOptions) string {
	return md.formatAlphaField(md.FormatVersion, 2, options)
}

// FormatMessageDispositionTestProductionCode returns TestProductionCode formatted according to the FormatOptions
func (md *MessageDisposition) FormatMessageDispositionTestProductionCode(options FormatOptions) string {
	return md.formatAlphaField(md.TestProductionCode, 1, options)
}

// FormatMessageDispositionMessageDuplicationCode returns MessageDuplicationCode formatted according to the FormatOptions
func (md *MessageDisposition) FormatMessageDispositionMessageDuplicationCode(options FormatOptions) string {
	return md.formatAlphaField(md.MessageDuplicationCode, 1, options)
}

// FormatMessageDispositionMessageStatusIndicator returns MessageStatusIndicator formatted according to the FormatOptions
func (md *MessageDisposition) FormatMessageDispositionMessageStatusIndicator(options FormatOptions) string {
	return md.formatAlphaField(md.MessageStatusIndicator, 1, options)
}
