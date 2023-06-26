// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (srd *SecondaryRemittanceDocument) Parse(record string) error {
	if utf8.RuneCountInString(record) < 13 {
		return NewTagMinLengthErr(13, len(record))
	}

	srd.tag = record[:6]
	length := 6

	value, read, err := srd.parseFixedStringField(record[length:], 4)
	if err != nil {
		return fieldError("DocumentTypeCode", err)
	}
	srd.DocumentTypeCode = value
	length += read

	value, read, err = srd.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ProprietaryDocumentTypeCode", err)
	}
	srd.ProprietaryDocumentTypeCode = value
	length += read

	value, read, err = srd.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("DocumentIdentificationNumber", err)
	}
	srd.DocumentIdentificationNumber = value
	length += read

	value, read, err = srd.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("Issuer", err)
	}
	srd.Issuer = value
	length += read

	if err := srd.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
}

func (srd *SecondaryRemittanceDocument) UnmarshalJSON(data []byte) error {
	type Alias SecondaryRemittanceDocument
	aux := struct {
		*Alias
	}{
		(*Alias)(srd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	srd.tag = TagSecondaryRemittanceDocument
	return nil
}

// String returns a fixed-width SecondaryRemittanceDocument record
func (srd *SecondaryRemittanceDocument) String() string {
	return srd.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a SecondaryRemittanceDocument record formatted according to the FormatOptions
func (srd *SecondaryRemittanceDocument) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(115)

	buf.WriteString(srd.tag)
	buf.WriteString(srd.DocumentTypeCodeField())
	buf.WriteString(srd.FormatProprietaryDocumentTypeCode(options) + Delimiter)
	buf.WriteString(srd.FormatDocumentIdentificationNumber(options) + Delimiter)
	buf.WriteString(srd.FormatIssuer(options) + Delimiter)

	if options.VariableLengthFields {
		return srd.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on SecondaryRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// * Document Type Code and Document Identification Number are mandatory.
// * Proprietary Document Type Code is mandatory for Document Type Code PROP; otherwise not permitted.
func (srd *SecondaryRemittanceDocument) Validate() error {
	if err := srd.fieldInclusion(); err != nil {
		return err
	}
	if srd.tag != TagSecondaryRemittanceDocument {
		return fieldError("tag", ErrValidTagForType, srd.tag)
	}
	if err := srd.isDocumentTypeCode(srd.DocumentTypeCode); err != nil {
		return fieldError("DocumentTypeCode", err, srd.DocumentTypeCode)
	}
	if err := srd.isAlphanumeric(srd.ProprietaryDocumentTypeCode); err != nil {
		return fieldError("ProprietaryDocumentTypeCode", err, srd.ProprietaryDocumentTypeCode)
	}
	if err := srd.isAlphanumeric(srd.DocumentIdentificationNumber); err != nil {
		return fieldError("DocumentIdentificationNumber", err, srd.DocumentIdentificationNumber)
	}
	if err := srd.isAlphanumeric(srd.Issuer); err != nil {
		return fieldError("Issuer", err, srd.Issuer)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (srd *SecondaryRemittanceDocument) fieldInclusion() error {
	if srd.DocumentIdentificationNumber == "" {
		return fieldError("DocumentIdentificationNumber", ErrFieldRequired)
	}
	switch srd.DocumentTypeCode {
	case ProprietaryDocumentType:
		if srd.ProprietaryDocumentTypeCode == "" {
			return fieldError("ProprietaryDocumentTypeCode", ErrFieldRequired)
		}
	default:
		if srd.ProprietaryDocumentTypeCode != "" {
			return fieldError("ProprietaryDocumentTypeCode", ErrInvalidProperty, srd.ProprietaryDocumentTypeCode)
		}
	}
	return nil
}

// DocumentTypeCodeField
func (srd *SecondaryRemittanceDocument) DocumentTypeCodeField() string {
	return srd.alphaField(srd.DocumentTypeCode, 4)
}

// FormatDocumentTypeCode returns DocumentTypeCode formatted according to the FormatOptions
func (srd *SecondaryRemittanceDocument) FormatDocumentTypeCode(options FormatOptions) string {
	return srd.formatAlphaField(srd.DocumentTypeCode, 4, options)
}

// FormatProprietaryDocumentTypeCode returns ProprietaryDocumentTypeCode formatted according to the FormatOptions
func (srd *SecondaryRemittanceDocument) FormatProprietaryDocumentTypeCode(options FormatOptions) string {
	return srd.formatAlphaField(srd.ProprietaryDocumentTypeCode, 35, options)
}

// FormatDocumentIdentificationNumber returns DocumentIdentificationNumber formatted according to the FormatOptions
func (srd *SecondaryRemittanceDocument) FormatDocumentIdentificationNumber(options FormatOptions) string {
	return srd.formatAlphaField(srd.DocumentIdentificationNumber, 35, options)
}

// FormatIssuer returns Issuer formatted according to the FormatOptions
func (srd *SecondaryRemittanceDocument) FormatIssuer(options FormatOptions) string {
	return srd.formatAlphaField(srd.Issuer, 35, options)
}
