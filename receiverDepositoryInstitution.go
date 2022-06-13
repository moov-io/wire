// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// ReceiverDepositoryInstitution {3400}
type ReceiverDepositoryInstitution struct {
	// tag
	tag string
	// ReceiverABANumber
	ReceiverABANumber string `json:"receiverABANumber"`
	// ReceiverShortName
	ReceiverShortName string `json:"receiverShortName"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewReceiverDepositoryInstitution returns a new ReceiverDepositoryInstitution
func NewReceiverDepositoryInstitution() *ReceiverDepositoryInstitution {
	rdi := &ReceiverDepositoryInstitution{
		tag: TagReceiverDepositoryInstitution,
	}
	return rdi
}

// Parse takes the input string and parses the ReceiverDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rdi *ReceiverDepositoryInstitution) Parse(record string) error {
	if utf8.RuneCountInString(record) < 10 {
		return NewTagMinLengthErr(10, len(record))
	}

	rdi.tag = record[:6]
	length := 6

	value, read, err := rdi.parseVariableStringField(record[length:], 9)
	if err != nil {
		return fieldError("ReceiverABANumber", err)
	}
	rdi.ReceiverABANumber = value
	length += read

	value, read, err = rdi.parseVariableStringField(record[length:], 18)
	if err != nil {
		return fieldError("ReceiverShortName", err)
	}
	rdi.ReceiverShortName = value
	length += read

	if !rdi.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (rdi *ReceiverDepositoryInstitution) UnmarshalJSON(data []byte) error {
	type Alias ReceiverDepositoryInstitution
	aux := struct {
		*Alias
	}{
		(*Alias)(rdi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rdi.tag = TagReceiverDepositoryInstitution
	return nil
}

// String writes ReceiverDepositoryInstitution
func (rdi *ReceiverDepositoryInstitution) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(33)

	buf.WriteString(rdi.tag)
	buf.WriteString(rdi.ReceiverABANumberField(options...))
	buf.WriteString(rdi.ReceiverShortNameField(options...))

	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiverDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rdi *ReceiverDepositoryInstitution) Validate() error {
	if err := rdi.fieldInclusion(); err != nil {
		return err
	}
	if rdi.tag != TagReceiverDepositoryInstitution {
		return fieldError("tag", ErrValidTagForType, rdi.tag)
	}
	if err := rdi.isNumeric(rdi.ReceiverABANumber); err != nil {
		return fieldError("ReceiverABANumber", err, rdi.ReceiverABANumber)
	}
	if err := rdi.isAlphanumeric(rdi.ReceiverShortName); err != nil {
		return fieldError("ReceiverShortName", err, rdi.ReceiverShortName)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rdi *ReceiverDepositoryInstitution) fieldInclusion() error {
	if rdi.ReceiverABANumber == "" {
		return fieldError("ReceiverABANumber", ErrFieldRequired, rdi.ReceiverABANumber)
	}
	if rdi.ReceiverShortName == "" {
		return fieldError("ReceiverShortName", ErrFieldRequired, rdi.ReceiverShortName)
	}
	return nil
}

// ReceiverABANumberField gets a string of the ReceiverABANumber field
func (rdi *ReceiverDepositoryInstitution) ReceiverABANumberField(options ...bool) string {
	return rdi.alphaVariableField(rdi.ReceiverABANumber, 9, rdi.parseFirstOption(options))
}

// ReceiverShortNameField gets a string of the ReceiverShortName field
func (rdi *ReceiverDepositoryInstitution) ReceiverShortNameField(options ...bool) string {
	return rdi.alphaVariableField(rdi.ReceiverShortName, 18, rdi.parseFirstOption(options))
}
