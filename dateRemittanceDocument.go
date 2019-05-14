// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// DateRemittanceDocument is the date of remittance document
type DateRemittanceDocument struct {
	// tag
	tag string
	// DateRemittanceDocument CCYYMMDD
	DateRemittanceDocument string `json:"dateRemittanceDocument,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewDateRemittanceDocument returns a new DateRemittanceDocument
func NewDateRemittanceDocument() *DateRemittanceDocument {
	drd := &DateRemittanceDocument{
		tag: TagDateRemittanceDocument,
	}
	return drd
}

// Parse takes the input string and parses the DateRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (drd *DateRemittanceDocument) Parse(record string) error {
	if utf8.RuneCountInString(record) != 14 {
		return NewTagWrongLengthErr(14, len(record))
	}
	drd.tag = record[:6]
	drd.DateRemittanceDocument = drd.parseStringField(record[6:14])
	return nil
}

// String writes DateRemittanceDocument
func (drd *DateRemittanceDocument) String() string {
	var buf strings.Builder
	buf.Grow(14)
	buf.WriteString(drd.tag)
	buf.WriteString(drd.DateRemittanceDocumentField())
	return buf.String()
}

// Validate performs WIRE format rule checks on DateRemittanceDocument and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (drd *DateRemittanceDocument) Validate() error {
	if err := drd.fieldInclusion(); err != nil {
		return err
	}
	if drd.tag != TagDateRemittanceDocument {
		return fieldError("tag", ErrValidTagForType, drd.tag)
	}
	if err := drd.validateDate(drd.DateRemittanceDocument); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (drd *DateRemittanceDocument) fieldInclusion() error {
	if drd.DateRemittanceDocument == "" {
		return fieldError("DateRemittanceDocument", ErrFieldRequired)
	}
	return nil
}

// DateRemittanceDocumentField gets a string of the DateRemittanceDocument field
func (drd *DateRemittanceDocument) DateRemittanceDocumentField() string {
	return drd.alphaField(drd.DateRemittanceDocument, 8)
}
