// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// Personal is personal demographic information
type Personal struct {
	// IdentificationCode:  * `1` - Passport Number * `2` - Tax Identification Number * `3` - Driver’s License Number * `4` - Alien Registration Number * `5` - Corporate Identification * `9` - Other Identification
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name string `json:"name"`
	// Address
	Address Address `json:"address"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the Personal values
func (p *Personal) Parse(record string) (length int, err error) {

	var read int

	p.IdentificationCode = p.parseStringField(record[length : length+1])
	length += 1

	if p.Identifier, read, err = p.parseVariableStringField(record[length:], 34); err != nil {
		return 0, fieldError("Identifier", err)
	}
	length += read

	if p.Name, read, err = p.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("Name", err)
	}
	length += read

	if read, err = p.Address.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return
}

// String writes Personal
func (p *Personal) String(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(175)

	buf.WriteString(p.IdentificationCodeField())
	buf.WriteString(p.IdentifierField(isVariable))
	buf.WriteString(p.NameField(isVariable))
	buf.WriteString(p.Address.String(isVariable))

	return buf.String()
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (p *Personal) IdentificationCodeField() string {
	return p.alphaField(p.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (p *Personal) IdentifierField(isVariable bool) string {
	return p.alphaVariableField(p.Identifier, 34, isVariable)
}

// NameField gets a string of the Name field
func (p *Personal) NameField(isVariable bool) string {
	return p.alphaVariableField(p.Name, 35, isVariable)
}
