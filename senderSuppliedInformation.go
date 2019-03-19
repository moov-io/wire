// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// SenderSuppliedInformation {1500}
type SenderSuppliedInformation struct {
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

// NewSenderSuppliedInformation returns a new SenderSuppliedInformation
func NewSenderSuppliedInformation() SenderSuppliedInformation {
	ssi := SenderSuppliedInformation{
		tag: TagSenderSuppliedInformation,
		FormatVersion: FormatVersion,
		TestProductionCode: EnvironmentTest,
		MessageDuplicationCode: MessageDuplicationOriginal,
	}
	return ssi
}

// Parse takes the input string and parses the SenderSuppliedInformation values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ssi *SenderSuppliedInformation) Parse(record string) {
}

// String writes SenderSuppliedInformation
func (ssi *SenderSuppliedInformation) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(18)
	buf.WriteString(ssi.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderSuppliedInformation and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ssi *SenderSuppliedInformation) Validate() error {
	if err := ssi.fieldInclusion(); err != nil {
		return err
	}
/*		if ssi.FormatVersion != FormatVersion {
		return fieldError("FormatVersion", NewErrFormatVersion(30), ssi.FormatVersion)
	}*/
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ssi *SenderSuppliedInformation) fieldInclusion() error {
	return nil
}
