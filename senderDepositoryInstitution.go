// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &SenderDepositoryInstitution{}

// SenderDepositoryInstitution {3100}
type SenderDepositoryInstitution struct {
	// tag
	tag string
	// SenderABANumber
	SenderABANumber string `json:"senderABANumber"`
	// SenderShortName
	SenderShortName string `json:"senderShortName"`

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
func (sdi *SenderDepositoryInstitution) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if sdi.tag, read, err = sdi.parseTag(record); err != nil {
		return 0, fieldError("SenderDepositoryInstitution.Tag", err)
	}
	length += read

	if sdi.SenderABANumber, read, err = sdi.parseVariableStringField(record[length:], 9); err != nil {
		return 0, fieldError("SenderABANumber", err)
	}
	length += read

	if sdi.SenderShortName, read, err = sdi.parseVariableStringField(record[length:], 18); err != nil {
		return 0, fieldError("SenderShortName", err)
	}
	length += read

	return length, nil
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

// String writes SenderDepositoryInstitution
func (sdi *SenderDepositoryInstitution) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(39)

	buf.WriteString(sdi.tag)
	buf.WriteString(sdi.SenderABANumberField(isCompressed))
	buf.WriteString(sdi.SenderShortNameField(isCompressed))

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
	if sdi.SenderShortName == "" {
		return fieldError("SenderShortName", ErrFieldRequired, sdi.SenderShortName)
	}
	return nil
}

// SenderABANumberField gets a string of the SenderABANumber field
func (sdi *SenderDepositoryInstitution) SenderABANumberField(isCompressed bool) string {
	return sdi.alphaVariableField(sdi.SenderABANumber, 9, isCompressed)
}

// SenderShortNameField gets a string of the SenderShortName field
func (sdi *SenderDepositoryInstitution) SenderShortNameField(isCompressed bool) string {
	return sdi.alphaVariableField(sdi.SenderShortName, 18, isCompressed)
}
