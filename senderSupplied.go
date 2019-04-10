// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// SenderSupplied {1500}
type SenderSupplied struct {
	// tag
	tag string
	// FormatVersion 30
	FormatVersion string `json:"formatVersion"`
	// UserRequestCorrelation
	UserRequestCorrelation string `json:"userRequestCorrelation"`
	// TestProductionCode T: Test P: Production
	TestProductionCode string `json:"testProductionCode"`
	// MessageDuplicationCode '': Original Message P: Resend
	MessageDuplicationCode string `json:"messageDuplicationCode"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSenderSupplied returns a new SenderSupplied
func NewSenderSupplied() *SenderSupplied {
	ss := &SenderSupplied{
		tag:                    TagSenderSupplied,
		FormatVersion:          FormatVersion,
		TestProductionCode:     EnvironmentTest,
		MessageDuplicationCode: MessageDuplicationOriginal,
	}
	return ss
}

// Parse takes the input string and parses the SenderSupplied values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ss *SenderSupplied) Parse(record string) {
	ss.tag = ss.parseStringField(record[:6])
	ss.FormatVersion = ss.parseStringField(record[6:8])
	ss.UserRequestCorrelation = ss.parseStringField(record[8:16])
	ss.TestProductionCode = ss.parseStringField(record[16:17])
	ss.MessageDuplicationCode = ss.parseStringField(record[17:18])
}

// String writes SenderSupplied
func (ss *SenderSupplied) String() string {
	var buf strings.Builder
	buf.Grow(18)
	buf.WriteString(ss.tag)
	buf.WriteString(ss.FormatVersionField())
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderSupplied and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ss *SenderSupplied) Validate() error {
	if err := ss.fieldInclusion(); err != nil {
		return err
	}
	if ss.FormatVersion != FormatVersion {
		return fieldError("FormatVersion", ErrFormatVersion, ss.FormatVersion)
	}
	if err := ss.isAlphanumeric(ss.UserRequestCorrelation); err != nil {
		return fieldError("UserRequestCorrelation", err, ss.UserRequestCorrelation)
	}
	if err := ss.isTestProductionCode(ss.TestProductionCode); err != nil {
		return fieldError("TestProductionCode", err, ss.TestProductionCode)
	}
	if err := ss.isMessageDuplicationCode(ss.MessageDuplicationCode); err != nil {
		return fieldError("MessageDuplicationCode", err, ss.MessageDuplicationCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ss *SenderSupplied) fieldInclusion() error {
	if ss.UserRequestCorrelation == "" {
		return fieldError("UserRequestCorrelation", ErrFieldRequired, ss.UserRequestCorrelation)
	}
	return nil
}

// FormatVersionField gets a string of the FormatVersion field
func (ss *SenderSupplied) FormatVersionField() string {
	return ss.alphaField(ss.FormatVersion, 2)
}

// UserRequestCorrelationField gets a string of the UserRequestCorrelation field
func (ss *SenderSupplied) UserRequestCorrelationField() string {
	return ss.alphaField(ss.UserRequestCorrelation, 8)
}

// TestProductionCodeField gets a string of the TestProductionCoden field
func (ss *SenderSupplied) TestProductionCodeField() string {
	return ss.alphaField(ss.TestProductionCode, 1)
}

// MessageDuplicationCodeField gets a string of the MessageDuplicationCode field
func (ss *SenderSupplied) MessageDuplicationCodeField() string {
	return ss.alphaField(ss.MessageDuplicationCode, 1)
}
