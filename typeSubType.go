// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// TypeSubType {1510}
type TypeSubType struct {
	// tag
	tag string
	// TypeCode
	TypeCode string `json:"typeCode"`
	// SubTypeCode
	SubTypeCode string `json:"subTypeCode"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewTypeSubType returns a new TypeSubType
func NewTypeSubType() *TypeSubType {
	tst := &TypeSubType{
		tag: TagTypeSubType,
	}
	return tst
}

// Parse takes the input string and parses the TypeSubType values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (tst *TypeSubType) Parse(record string) error {
	if utf8.RuneCountInString(record) != 10 {
		return NewTagWrongLengthErr(10, len(record))
	}
	tst.tag = tst.parseStringField(record[:6])
	tst.TypeCode = tst.parseStringField(record[6:8])
	tst.SubTypeCode = tst.parseStringField(record[8:10])
	return nil
}

// String writes TypeSubType
func (tst *TypeSubType) String() string {
	var buf strings.Builder
	buf.Grow(10)
	buf.WriteString(tst.tag)
	buf.WriteString(tst.TypeCodeField())
	buf.WriteString(tst.SubTypeCodeField())
	return buf.String()
}

// Validate performs WIRE format rule checks on TypeSubType and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (tst *TypeSubType) Validate() error {
	if err := tst.fieldInclusion(); err != nil {
		return err
	}
	if err := tst.isTypeCode(tst.TypeCode); err != nil {
		return fieldError("TypeCode", err, tst.TypeCode)
	}
	if err := tst.isSubTypeCode(tst.SubTypeCode); err != nil {
		return fieldError("SubTypeCode", err, tst.SubTypeCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (tst *TypeSubType) fieldInclusion() error {
	if tst.TypeCode == "" {
		return fieldError("TypeCode", ErrFieldRequired)
	}
	if tst.SubTypeCode == "" {
		return fieldError("SubTypeCode", ErrFieldRequired)
	}
	return nil
}

// TypeCodeField gets a string of the TypeCode field
func (tst *TypeSubType) TypeCodeField() string {
	return tst.alphaField(tst.TypeCode, 2)
}

// SubTypeCodeField gets a string of the SubTypeCode field
func (tst *TypeSubType) SubTypeCodeField() string {
	return tst.alphaField(tst.SubTypeCode, 2)
}
