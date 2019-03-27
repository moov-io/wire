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
func NewSenderSupplied() SenderSupplied {
	ss := SenderSupplied{
		tag: TagSenderSupplied,
		FormatVersion: FormatVersion,
		TestProductionCode: EnvironmentTest,
		MessageDuplicationCode: MessageDuplicationOriginal,
	}
	return ss
}

// Parse takes the input string and parses the SenderSupplied values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ss *SenderSupplied) Parse(record string) {
}

// String writes SenderSupplied
func (ss *SenderSupplied) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(18)
	buf.WriteString(ss.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SenderSupplied and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ss *SenderSupplied) Validate() error {
	if err := ss.fieldInclusion(); err != nil {
		return err
	}
/*		if ssi.FormatVersion != FormatVersion {
		return fieldError("FormatVersion", NewErrFormatVersion(30), ssi.FormatVersion)
	}*/
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ss *SenderSupplied) fieldInclusion() error {
	return nil
}
