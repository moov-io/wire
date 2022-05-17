// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &TypeSubType{}

// TypeSubType {1510}
type TypeSubType struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
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
func NewTypeSubType(isVariable bool) *TypeSubType {
	tst := &TypeSubType{
		tag:              TagTypeSubType,
		isVariableLength: isVariable,
	}
	return tst
}

// Parse takes the input string and parses the TypeSubType values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (tst *TypeSubType) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, utf8.RuneCountInString(record))
	}

	tst.tag = tst.parseStringField(record[:6])

	length := 6
	read := 0

	tst.TypeCode, read = tst.parseVariableStringField(record[length:], 2)
	length += read

	tst.SubTypeCode, read = tst.parseVariableStringField(record[length:], 2)
	length += read

	return length, nil
}

func (tst *TypeSubType) UnmarshalJSON(data []byte) error {
	type Alias TypeSubType
	aux := struct {
		*Alias
	}{
		(*Alias)(tst),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	tst.tag = TagTypeSubType
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
	if tst.tag != TagTypeSubType {
		return fieldError("tag", ErrValidTagForType, tst.tag)
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
	return tst.alphaVariableField(tst.TypeCode, 2, tst.isVariableLength)
}

// SubTypeCodeField gets a string of the SubTypeCode field
func (tst *TypeSubType) SubTypeCodeField() string {
	return tst.alphaVariableField(tst.SubTypeCode, 2, tst.isVariableLength)
}
