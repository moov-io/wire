// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
func NewSenderDepositoryInstitution() SenderDepositoryInstitution {
	sdi := SenderDepositoryInstitution{
		tag: TagSenderDepositoryInstitution,
	}
	return sdi
}

// Parse takes the input string and parses the SenderDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (sdi *SenderDepositoryInstitution) Parse(record string) {
	sdi.tag = record[:6]
	sdi.SenderABANumber = sdi.parseStringField(record[6:15])
	sdi.SenderShortName = sdi.parseStringField(record[15:33])
}

// String writes SenderDepositoryInstitution
func (sdi *SenderDepositoryInstitution) String() string {
	var buf strings.Builder
	buf.Grow(39)
	buf.WriteString(sdi.tag)
	buf.WriteString(sdi.SenderABANumberField())
	buf.WriteString(sdi.SenderShortNameField())
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sdi *SenderDepositoryInstitution) Validate() error {
	if err := sdi.fieldInclusion(); err != nil {
		return err
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
func (sdi *SenderDepositoryInstitution) SenderABANumberField() string {
	return sdi.alphaField(sdi.SenderABANumber, 9)
}

// SenderShortNameField gets a string of the SenderShortName field
func (sdi *SenderDepositoryInstitution) SenderShortNameField() string {
	return sdi.alphaField(sdi.SenderShortName, 18)
}
