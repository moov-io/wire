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
}

// String writes SenderDepositoryInstitution
func (sdi *SenderDepositoryInstitution) String() string {
	var buf strings.Builder
	// ToDo: Check on separator
	buf.Grow(33)
	buf.WriteString(sdi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (sdi *SenderDepositoryInstitution) Validate() error {
	if err := sdi.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (sdi *SenderDepositoryInstitution) fieldInclusion() error {
	return nil
}
