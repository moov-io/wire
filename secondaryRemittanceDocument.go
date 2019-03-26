// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// SecondaryRemittanceDocument is the date of remittance document
type SecondaryRemittanceDocument struct {
	// tag
	tag string
	// SecondaryRemittanceDocument  * `AROI` - Accounts Receivable Open Item * `DISP` - Dispatch Advice * `FXDR` - Foreign Exchange Deal Reference * `PROP` - Proprietary Document Type PUOR Purchase Order * `RADM` - Remittance Advice Message * `RPIN` - Related Payment Instruction * `SCOR1` - Structured Communication Reference VCHR Voucher
	DocumentTypeCode string `json:"documentTypeCode,omitempty"`
	// proprietaryDocumentTypeCode
	ProprietaryDocumentTypeCode string `json:"proprietaryDocumentTypeCode,omitempty"`
	// documentIdentificationNumber
	DocumentIdentificationNumber string `json:"documentIdentificationNumber,omitempty"`
	// Issuer
	Issuer string `json:"issuer,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewSecondaryRemittanceDocument returns a new SecondaryRemittanceDocument
func NewSecondaryRemittanceDocument() SecondaryRemittanceDocument  {
	srd := SecondaryRemittanceDocument {
		tag: TagSecondaryRemittanceDocument,
	}
	return srd
}

// Parse takes the input string and parses the SecondaryRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (srd *SecondaryRemittanceDocument) Parse(record string) {
}

// String writes SecondaryRemittanceDocument
func (srd *SecondaryRemittanceDocument) String() string {
	var buf strings.Builder
	buf.Grow(8)
	buf.WriteString(srd.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on SecondaryRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (srd *SecondaryRemittanceDocument) Validate() error {
	if err := srd.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (srd *SecondaryRemittanceDocument) fieldInclusion() error {
	return nil
}
