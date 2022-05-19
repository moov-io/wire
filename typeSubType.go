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
func (tst *TypeSubType) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if tst.tag, read, err = tst.parseTag(record); err != nil {
		return 0, fieldError("TypeSubType.Tag", err)
	}
	length += read

	if tst.TypeCode, read, err = tst.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("TypeCode", err)
	}
	length += read

	if tst.SubTypeCode, read, err = tst.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("SubTypeCode", err)
	}
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
func (tst *TypeSubType) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(10)

	buf.WriteString(tst.tag)
	buf.WriteString(tst.TypeCodeField(isCompressed))
	buf.WriteString(tst.SubTypeCodeField(isCompressed))

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
func (tst *TypeSubType) TypeCodeField(isCompressed bool) string {
	return tst.alphaVariableField(tst.TypeCode, 2, isCompressed)
}

// SubTypeCodeField gets a string of the SubTypeCode field
func (tst *TypeSubType) SubTypeCodeField(isCompressed bool) string {
	return tst.alphaVariableField(tst.SubTypeCode, 2, isCompressed)
}
