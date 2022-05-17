// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &PrimaryRemittanceDocument{}

// PrimaryRemittanceDocument is primary remittance document
type PrimaryRemittanceDocument struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
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
func NewPrimaryRemittanceDocument(isVariable bool) *PrimaryRemittanceDocument {
	prd := &PrimaryRemittanceDocument{
		tag:              TagPrimaryRemittanceDocument,
		isVariableLength: isVariable,
	}
	return prd
}

// Parse takes the input string and parses the PrimaryRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (prd *PrimaryRemittanceDocument) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 10 {
		return 0, NewTagWrongLengthErr(10, len(record))
	}

	prd.tag = record[:6]

	length := 6
	read := 0

	prd.DocumentTypeCode, read = prd.parseVariableStringField(record[length:], 4)
	length += read

	prd.ProprietaryDocumentTypeCode, read = prd.parseVariableStringField(record[length:], 35)
	length += read

	prd.DocumentIdentificationNumber, read = prd.parseVariableStringField(record[length:], 35)
	length += read

	prd.Issuer, read = prd.parseVariableStringField(record[length:], 35)
	length += read

	return length, nil
}

func (prd *PrimaryRemittanceDocument) UnmarshalJSON(data []byte) error {
	type Alias PrimaryRemittanceDocument
	aux := struct {
		*Alias
	}{
		(*Alias)(prd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	prd.tag = TagPrimaryRemittanceDocument
	return nil
}

// String writes PrimaryRemittanceDocument
func (prd *PrimaryRemittanceDocument) String() string {
	var buf strings.Builder
	buf.Grow(115)

	buf.WriteString(prd.tag)
	buf.WriteString(prd.DocumentTypeCodeField())
	buf.WriteString(prd.ProprietaryDocumentTypeCodeField())
	buf.WriteString(prd.DocumentIdentificationNumberField())
	buf.WriteString(prd.IssuerField())

	return buf.String()
}

// Validate performs WIRE format rule checks on PrimaryRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// Document Type Code and Document Identification Number are mandatory for each set of remittance data.
// Proprietary Document Type Code is mandatory for Document Type Code PROP; otherwise not permitted.
func (prd *PrimaryRemittanceDocument) Validate() error {
	if err := prd.fieldInclusion(); err != nil {
		return err
	}
	if prd.tag != TagPrimaryRemittanceDocument {
		return fieldError("tag", ErrValidTagForType, prd.tag)
	}
	if err := prd.isDocumentTypeCode(prd.DocumentTypeCode); err != nil {
		return fieldError("DocumentTypeCode", err, prd.DocumentTypeCode)
	}
	if err := prd.isAlphanumeric(prd.ProprietaryDocumentTypeCode); err != nil {
		return fieldError("ProprietaryDocumentTypeCode", err, prd.ProprietaryDocumentTypeCode)
	}
	if err := prd.isAlphanumeric(prd.DocumentIdentificationNumber); err != nil {
		return fieldError("DocumentIdentificationNumber", err, prd.DocumentIdentificationNumber)
	}
	if err := prd.isAlphanumeric(prd.Issuer); err != nil {
		return fieldError("Issuer", err, prd.Issuer)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (prd *PrimaryRemittanceDocument) fieldInclusion() error {
	if prd.DocumentIdentificationNumber == "" {
		return fieldError("DocumentIdentificationNumber", ErrFieldRequired)
	}
	switch prd.DocumentTypeCode {
	case ProprietaryDocumentType:
		if prd.ProprietaryDocumentTypeCode == "" {
			return fieldError("ProprietaryDocumentTypeCode", ErrFieldRequired)
		}
	default:
		if strings.TrimSpace(prd.ProprietaryDocumentTypeCode) != "" {
			return fieldError("ProprietaryDocumentTypeCode", ErrInvalidProperty, prd.ProprietaryDocumentTypeCode)
		}
	}
	return nil
}

// DocumentTypeCodeField gets a string of the DocumentTypeCode field
func (prd *PrimaryRemittanceDocument) DocumentTypeCodeField() string {
	return prd.alphaVariableField(prd.DocumentTypeCode, 4, prd.isVariableLength)
}

// ProprietaryDocumentTypeCodeField gets a string of the ProprietaryDocumentTypeCode field
func (prd *PrimaryRemittanceDocument) ProprietaryDocumentTypeCodeField() string {
	return prd.alphaVariableField(prd.ProprietaryDocumentTypeCode, 35, prd.isVariableLength)
}

// DocumentIdentificationNumberField gets a string of the DocumentIdentificationNumber field
func (prd *PrimaryRemittanceDocument) DocumentIdentificationNumberField() string {
	return prd.alphaVariableField(prd.DocumentIdentificationNumber, 35, prd.isVariableLength)
}

// IssuerField gets a string of the Issuer field
func (prd *PrimaryRemittanceDocument) IssuerField() string {
	return prd.alphaVariableField(prd.Issuer, 35, prd.isVariableLength)
}
