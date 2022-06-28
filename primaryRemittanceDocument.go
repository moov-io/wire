// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func NewPrimaryRemittanceDocument() *PrimaryRemittanceDocument {
	prd := &PrimaryRemittanceDocument{
		tag: TagPrimaryRemittanceDocument,
	}
	return prd
}

// Parse takes the input string and parses the PrimaryRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (prd *PrimaryRemittanceDocument) Parse(record string) error {
	if utf8.RuneCountInString(record) < 13 {
		return NewTagMinLengthErr(13, len(record))
	}

	prd.tag = record[:6]
	prd.DocumentTypeCode = record[6:10]
	length := 10

	value, read, err := prd.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ProprietaryDocumentTypeCode", err)
	}
	prd.ProprietaryDocumentTypeCode = value
	length += read

	value, read, err = prd.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("DocumentIdentificationNumber", err)
	}
	prd.DocumentIdentificationNumber = value
	length += read

	value, read, err = prd.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Issuer", err)
	}
	prd.Issuer = value
	length += read

	if !prd.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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

// String returns a fixed-width PrimaryRemittanceDocument record
func (prd *PrimaryRemittanceDocument) String() string {
	return prd.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a PrimaryRemittanceDocument record formatted according to the FormatOptions
func (prd *PrimaryRemittanceDocument) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(115)

	buf.WriteString(prd.tag)
	buf.WriteString(prd.DocumentTypeCodeField())
	buf.WriteString(prd.FormatProprietaryDocumentTypeCode(options))
	buf.WriteString(prd.FormatDocumentIdentificationNumber(options))
	buf.WriteString(prd.FormatIssuer(options))

	if options.VariableLengthFields {
		return prd.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
	return prd.alphaField(prd.DocumentTypeCode, 4)
}

// ProprietaryDocumentTypeCodeField gets a string of the ProprietaryDocumentTypeCode field
func (prd *PrimaryRemittanceDocument) ProprietaryDocumentTypeCodeField() string {
	return prd.alphaField(prd.ProprietaryDocumentTypeCode, 35)
}

// DocumentIdentificationNumberField gets a string of the DocumentIdentificationNumber field
func (prd *PrimaryRemittanceDocument) DocumentIdentificationNumberField() string {
	return prd.alphaField(prd.DocumentIdentificationNumber, 35)
}

// IssuerField gets a string of the Issuer field
func (prd *PrimaryRemittanceDocument) IssuerField() string {
	return prd.alphaField(prd.Issuer, 35)
}

// FormatProprietaryDocumentTypeCode returns ProprietaryDocumentTypeCode formatted according to the FormatOptions
func (prd *PrimaryRemittanceDocument) FormatProprietaryDocumentTypeCode(options FormatOptions) string {
	return prd.formatAlphaField(prd.ProprietaryDocumentTypeCode, 35, options)
}

// FormatDocumentIdentificationNumber returns DocumentIdentificationNumber formatted according to the FormatOptions
func (prd *PrimaryRemittanceDocument) FormatDocumentIdentificationNumber(options FormatOptions) string {
	return prd.formatAlphaField(prd.DocumentIdentificationNumber, 35, options)
}

// FormatIssuer returns Issuer formatted according to the FormatOptions
func (prd *PrimaryRemittanceDocument) FormatIssuer(options FormatOptions) string {
	return prd.formatAlphaField(prd.Issuer, 35, options)
}
