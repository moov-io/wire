// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// UnstructuredAddenda is the unstructured addenda information
type UnstructuredAddenda struct {
	// tag
	tag string
	// AddendaLength  Addenda Length must be numeric, padded with leading zeros if less than four characters and must equal length of content in Addenda Information (e.g., if content of Addenda Information is 987 characters, Addenda Length must be 0987).
	AddendaLength string `json:"addendaLength,omitempty"`
	// Addenda
	Addenda string `json:"addenda,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewUnstructuredAddenda returns a new UnstructuredAddenda
func NewUnstructuredAddenda() *UnstructuredAddenda {
	ua := &UnstructuredAddenda{
		tag: TagUnstructuredAddenda,
	}
	return ua
}

// Parse takes the input string and parses the UnstructuredAddenda values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ua *UnstructuredAddenda) Parse(record string) error {
	// First check ua.tag and ua.AddendaLength
	if utf8.RuneCountInString(record) < 10 {
		return NewTagWrongLengthErr(10, utf8.RuneCountInString(record))
	}
	ua.tag = record[:6]
	ua.AddendaLength = record[6:10]
	al := ua.parseNumField(ua.AddendaLength)
	// check RuneCount for entire record
	if utf8.RuneCountInString(record) != 10+al {
		return NewTagWrongLengthErr(10+al, utf8.RuneCountInString(record))
	}
	ua.Addenda = ua.parseStringField(record[10 : 10+al])
	return nil
}

func (ua *UnstructuredAddenda) UnmarshalJSON(data []byte) error {
	type Alias UnstructuredAddenda
	aux := struct {
		*Alias
	}{
		(*Alias)(ua),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ua.tag = TagUnstructuredAddenda
	return nil
}

// String writes UnstructuredAddenda
func (ua *UnstructuredAddenda) String() string {
	var buf strings.Builder
	buf.Grow(10)
	buf.WriteString(ua.tag)
	buf.WriteString(ua.AddendaLengthField())

	if size := ua.parseNumField(ua.AddendaLength); validSizeInt(size) {
		buf.Grow(size)
	}

	buf.WriteString(ua.AddendaField())
	return buf.String()
}

// Validate performs WIRE format rule checks on UnstructuredAddenda and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// AddendaLength must be numeric, padded with leading zeros if less than four characters and must equal
//
//	length of content in Addenda Information (e.g., if content of Addenda Information is 987 characters,
//	Addenda Length must be 0987).
func (ua *UnstructuredAddenda) Validate() error {
	if err := ua.fieldInclusion(); err != nil {
		return err
	}
	if ua.tag != TagUnstructuredAddenda {
		return fieldError("tag", ErrValidTagForType, ua.tag)
	}
	if err := ua.isNumeric(ua.AddendaLength); err != nil {
		return fieldError("AddendaLength", err, ua.AddendaLength)
	}
	if err := ua.isAlphanumeric(ua.Addenda); err != nil {
		return fieldError("Addenda", err, ua.Addenda)
	}

	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ua *UnstructuredAddenda) fieldInclusion() error {
	// If UnstructuredAddenda is defined, AddendaLength is required, however it could be "0000"), but
	// I'm not sure of the point
	if ua.AddendaLength == "" {
		return fieldError("AddendaLength", ErrFieldRequired)
	}
	return nil
}

// AddendaLengthField gets a string of the AddendaLength field
func (ua *UnstructuredAddenda) AddendaLengthField() string {
	return ua.alphaField(ua.AddendaLength, 4)
}

// AddendaField gets a string of the Addenda field
func (ua *UnstructuredAddenda) AddendaField() string {
	max := ua.parseNumField(ua.AddendaLength)
	if max < 0 || !validSizeInt(max) {
		return ""
	}
	return ua.alphaField(ua.Addenda, uint(max))
}
