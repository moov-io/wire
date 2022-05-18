// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &DateRemittanceDocument{}

// DateRemittanceDocument is the date of remittance document
type DateRemittanceDocument struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// DateRemittanceDocument CCYYMMDD
	DateRemittanceDocument string `json:"dateRemittanceDocument,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewDateRemittanceDocument returns a new DateRemittanceDocument
func NewDateRemittanceDocument(isVariable bool) *DateRemittanceDocument {
	drd := &DateRemittanceDocument{
		tag:              TagDateRemittanceDocument,
		isVariableLength: isVariable,
	}
	return drd
}

// Parse takes the input string and parses the DateRemittanceDocument values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (drd *DateRemittanceDocument) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 7 {
		return 0, NewTagWrongLengthErr(7, len(record))
	}

	var err error
	var length, read int

	if drd.tag, read, err = drd.parseTag(record); err != nil {
		return 0, fieldError("DateRemittanceDocument.Tag", err)
	}
	length += read

	if drd.DateRemittanceDocument, read, err = drd.parseVariableStringField(record[length:], 8); err != nil {
		fieldError("DateRemittanceDocument", err)
	}
	length += read

	return length, nil
}

func (drd *DateRemittanceDocument) UnmarshalJSON(data []byte) error {
	type Alias DateRemittanceDocument
	aux := struct {
		*Alias
	}{
		(*Alias)(drd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	drd.tag = TagDateRemittanceDocument
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
	return drd.alphaVariableField(drd.DateRemittanceDocument, 8, drd.isVariableLength)
}
