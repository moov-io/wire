// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// SecondaryRemittanceDocument is the date of remittance document
type SecondaryRemittanceDocument struct {
	// tag
	tag string
	// DocumentTypeCode  * `AROI` - Accounts Receivable Open Item * `DISP` - Dispatch Advice * `FXDR` - Foreign Exchange Deal Reference * `PROP` - Proprietary Document Type PUOR Purchase Order * `RADM` - Remittance Advice Message * `RPIN` - Related Payment Instruction * `SCOR1` - Structured Communication Reference VCHR Voucher
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
func NewSecondaryRemittanceDocument() *SecondaryRemittanceDocument {
	srd := &SecondaryRemittanceDocument{
		tag: TagSecondaryRemittanceDocument,
	}
	return srd
}

// Parse takes the input string and parses the SecondaryRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (srd *SecondaryRemittanceDocument) Parse(record string) {
	srd.tag = record[:6]
	srd.DocumentTypeCode = srd.parseStringField(record[6:10])
	srd.ProprietaryDocumentTypeCode = srd.parseStringField(record[10:45])
	srd.DocumentIdentificationNumber = srd.parseStringField(record[45:80])
	srd.Issuer = srd.parseStringField(record[80:115])
}

// String writes SecondaryRemittanceDocument
func (srd *SecondaryRemittanceDocument) String() string {
	var buf strings.Builder
	buf.Grow(115)
	buf.WriteString(srd.tag)
	buf.WriteString(srd.DocumentTypeCodeField())
	buf.WriteString(srd.ProprietaryDocumentTypeCodeField())
	buf.WriteString(srd.DocumentIdentificationNumberField())
	buf.WriteString(srd.IssuerField())
	return buf.String()
}

// Validate performs WIRE format rule checks on SecondaryRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (srd *SecondaryRemittanceDocument) Validate() error {
	if err := srd.fieldInclusion(); err != nil {
		return err
	}
	if err := srd.isDocumentTypeCode(srd.DocumentTypeCode); err != nil {
		return fieldError("DocumentTypeCode", err, srd.DocumentTypeCode)
	}
	if err := srd.isAlphanumeric(srd.ProprietaryDocumentTypeCode); err != nil {
		return fieldError("ProprietaryDocumentTypeCode", err, srd.ProprietaryDocumentTypeCode)
	}
	if err := srd.isAlphanumeric(srd.DocumentIdentificationNumber); err != nil {
		return fieldError("DocumentIdentificationNumber", err, srd.Issuer)
	}
	if err := srd.isAlphanumeric(srd.Issuer); err != nil {
		return fieldError("Issuer", err, srd.Issuer)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (srd *SecondaryRemittanceDocument) fieldInclusion() error {
	return nil
}

// DocumentTypeCodeField gets a string of the DocumentTypeCode field
func (srd *SecondaryRemittanceDocument) DocumentTypeCodeField() string {
	return srd.alphaField(srd.DocumentTypeCode, 4)
}

// ProprietaryDocumentTypeCodeField gets a string of the ProprietaryDocumentTypeCode field
func (srd *SecondaryRemittanceDocument) ProprietaryDocumentTypeCodeField() string {
	return srd.alphaField(srd.ProprietaryDocumentTypeCode, 35)
}

// DocumentIdentificationNumberField gets a string of the DocumentIdentificationNumber field
func (srd *SecondaryRemittanceDocument) DocumentIdentificationNumberField() string {
	return srd.alphaField(srd.DocumentIdentificationNumber, 35)
}

// IssuerField gets a string of the Issuer field
func (srd *SecondaryRemittanceDocument) IssuerField() string {
	return srd.alphaField(srd.Issuer, 35)
}
