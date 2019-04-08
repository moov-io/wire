// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func NewReceiverDepositoryInstitution() ReceiverDepositoryInstitution {
	rdi := ReceiverDepositoryInstitution{
		tag: TagReceiverDepositoryInstitution,
	}
	return rdi
}

// Parse takes the input string and parses the ReceiverDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rdi *ReceiverDepositoryInstitution) Parse(record string) {
	rdi.tag = record[:6]
	rdi.ReceiverABANumber = rdi.parseStringField(record[6:15])
	rdi.ReceiverShortName = rdi.parseStringField(record[15:33])
}

// String writes ReceiverDepositoryInstitution
func (rdi *ReceiverDepositoryInstitution) String() string {
	var buf strings.Builder
	buf.Grow(33)
	buf.WriteString(rdi.tag)
	buf.WriteString(rdi.ReceiverABANumber)
	buf.WriteString(rdi.ReceiverShortName)
	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiverDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rdi *ReceiverDepositoryInstitution) Validate() error {
	if err := rdi.fieldInclusion(); err != nil {
		return err
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
		return fieldError("ReceiveABANumber", ErrFieldRequired, rdi.ReceiverABANumber)
	}
	if rdi.ReceiverShortName == "" {
		return fieldError("ReceiverShortName", ErrFieldRequired, rdi.ReceiverShortName)
	}
	return nil
}

// ReceiverABANumberField gets a string of the ReceiverABANumber field
func (rdi *ReceiverDepositoryInstitution) ReceiverABANumberField() string {
	return rdi.alphaField(rdi.ReceiverABANumber, 9)
}

// ReceiverShortNameField gets a string of the ReceiverShortName field
func (rdi *ReceiverDepositoryInstitution) ReceiverShortNameField() string {
	return rdi.alphaField(rdi.ReceiverShortName, 18)
}
