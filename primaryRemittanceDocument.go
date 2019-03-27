// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// PrimaryRemittanceDocument is primary remittance document
type PrimaryRemittanceDocument struct {
	// tag
	tag string
	// DocumentTypeCode  * `AROI` - Accounts Receivable Open Item * `BOLD` - Bill of Lading Shipping Notice * `CINV` - Commercial Invoice * `CMCN` - Commercial Contract * `CNFA` - Credit Note Related to Financial Adjustment * `CREN` - Credit Note * `DEBN` - Debit Note * `DISP` - Dispatch Advice * `DNFA` - Debit Note Related to Financial Adjustment HIRI Hire Invoice * `MSIN` - Metered Service Invoice * `PROP` - Proprietary Document Type * `PUOR` - Purchase Order * `SBIN` - Self Billed Invoice * `SOAC` - Statement of Account * `TSUT` - Trade Services Utility Transaction VCHR Voucher
	DocumentTypeCode string `json:"documentTypeCode,omitempty"`
	// ProprietaryDocumentTypeCode
	ProprietaryDocumentTypeCode string `json:"proprietaryDocumentTypeCode,omitempty"`
	// DocumentIdentificationNumber
	DocumentIdentificationNumber string `json:"documentIdentificationNumber,omitempty"`
	// Issuer
	Issuer string `json:"issuer,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewPrimaryRemittanceDocument returns a new PrimaryRemittanceDocument
func NewPrimaryRemittanceDocument() PrimaryRemittanceDocument {
	prd := PrimaryRemittanceDocument{
		tag: TagPrimaryRemittanceDocument,
	}
	return prd
}

// Parse takes the input string and parses the PrimaryRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (prd *PrimaryRemittanceDocument) Parse(record string) {
}

// String writes PrimaryRemittanceDocument
func (prd *PrimaryRemittanceDocument) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(109)
	buf.WriteString(prd.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on PrimaryRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (prd *PrimaryRemittanceDocument) Validate() error {
	if err := prd.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (prd *PrimaryRemittanceDocument) fieldInclusion() error {
	return nil
}
