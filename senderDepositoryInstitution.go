// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// SenderDepositoryInstitution {3100}
type SenderDepositoryInstitution struct {
	// tag
	tag string
	// SenderABANumber
	SenderABANumber string `json:"senderABANumber"`
	// SenderShortName
	SenderShortName string `json:"senderShortName,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSenderDepositoryInstitution returns a new SenderDepositoryInstitution
func NewSenderDepositoryInstitution() *SenderDepositoryInstitution {
	sdi := &SenderDepositoryInstitution{
		tag: TagSenderDepositoryInstitution,
	}
	return sdi
}

// Parse takes the input string and parses the SenderDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sdi *SenderDepositoryInstitution) Parse(record string) error {
	if utf8.RuneCountInString(record) < 8 {
		return NewTagMinLengthErr(8, len(record))
	}

	sdi.tag = record[:6]
	length := 6

	value, read, err := sdi.parseVariableStringField(record[length:], 9)
	if err != nil {
		return fieldError("SenderABANumber", err)
	}
	sdi.SenderABANumber = value
	length += read

	value, read, err = sdi.parseVariableStringField(record[length:], 18)
	if err != nil {
		return fieldError("SenderShortName", err)
	}
	fmt.Println("SenderShortName", value)
	sdi.SenderShortName = value
	length += read

	if !sdi.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (sdi *SenderDepositoryInstitution) UnmarshalJSON(data []byte) error {
	type Alias SenderDepositoryInstitution
	aux := struct {
		*Alias
	}{
		(*Alias)(sdi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	sdi.tag = TagSenderDepositoryInstitution
	return nil
}

// String returns a fixed-width SenderDepositoryInstitution record
func (sdi *SenderDepositoryInstitution) String() string {
	return sdi.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a SenderDepositoryInstitution record formatted according to the FormatOptions
func (sdi *SenderDepositoryInstitution) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(39)

	buf.WriteString(sdi.tag)
	buf.WriteString(sdi.FormatSenderABANumber(options))
	buf.WriteString(sdi.FormatSenderShortName(options))

	return buf.String()
}

// Validate performs WIRE format rule checks on SenderDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sdi *SenderDepositoryInstitution) Validate() error {
	if err := sdi.fieldInclusion(); err != nil {
		return err
	}
	if sdi.tag != TagSenderDepositoryInstitution {
		return fieldError("tag", ErrValidTagForType, sdi.tag)
	}
	if err := sdi.isNumeric(sdi.SenderABANumber); err != nil {
		return fieldError("SenderABANumber", err, sdi.SenderABANumber)
	}
	if err := sdi.isAlphanumeric(sdi.SenderShortName); err != nil {
		return fieldError("SenderShortName", err, sdi.SenderShortName)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sdi *SenderDepositoryInstitution) fieldInclusion() error {
	if sdi.SenderABANumber == "" {
		return fieldError("SenderABANumber", ErrFieldRequired, sdi.SenderABANumber)
	}
	return nil
}

// SenderABANumberField gets a string of the SenderABANumber field
func (sdi *SenderDepositoryInstitution) SenderABANumberField() string {
	return sdi.alphaField(sdi.SenderABANumber, 9)
}

// SenderShortNameField gets a string of the SenderShortName field
func (sdi *SenderDepositoryInstitution) SenderShortNameField() string {
	return sdi.alphaField(sdi.SenderShortName, 18)
}

// FormatSenderABANumber returns SenderABANumber formatted according to the FormatOptions
func (sdi *SenderDepositoryInstitution) FormatSenderABANumber(options FormatOptions) string {
	return sdi.formatAlphaField(sdi.SenderABANumber, 9, options)
}

// FormatSenderShortName returns SenderShortName formatted according to the FormatOptions
func (sdi *SenderDepositoryInstitution) FormatSenderShortName(options FormatOptions) string {
	output := sdi.formatAlphaField(sdi.SenderShortName, 18, options)

	//If this element is not present,the delimiter is not permitted
	if output == "*" {
		output = ""
	}

	return output
}
