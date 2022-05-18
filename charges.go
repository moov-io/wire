// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &Charges{}

// Charges is the Charges of the wire
type Charges struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// ChargeDetails * `B` - Beneficiary * `S` - Shared
	ChargeDetails string `json:"chargeDetails,omitempty"`
	// SendersChargesOne  The first three characters must contain an alpha currency code (e.g., USD).  The remaining
	// characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma
	// marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.
	SendersChargesOne string `json:"sendersChargesOne,omitempty"`
	// SendersChargesTwo  The first three characters must contain an alpha currency code (e.g., USD).  The remaining
	// characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma
	// marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.
	SendersChargesTwo string `json:"sendersChargesTwo,omitempty"`
	// SendersChargesThree  The first three characters must contain an alpha currency code (e.g., USD).  The remaining
	// characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma
	// marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.
	SendersChargesThree string `json:"sendersChargesThree,omitempty"`
	// SendersChargesFour  The first three characters must contain an alpha currency code (e.g., USD).  The remaining
	// characters for the amount must begin with at least one numeric character (0-9) and only one decimal comma
	// marker.  $1,234.56 should be entered as USD1234,56 and $0.99 should be entered as USD0,99.
	SendersChargesFour string `json:"sendersChargesFour,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewCharges returns a new Charges
func NewCharges(isVariable bool) *Charges {
	c := &Charges{
		tag:              TagCharges,
		isVariableLength: isVariable,
	}
	return c
}

// Parse takes the input string and parses the Charges values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (c *Charges) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 11 {
		return 0, NewTagWrongLengthErr(11, len(record))
	}

	var err error
	var length, read int

	if c.tag, read, err = c.parseTag(record); err != nil {
		return 0, fieldError("Charges.Tag", err)
	}
	length += read

	c.ChargeDetails = c.parseStringField(record[length : length+1])
	length += 1

	if c.SendersChargesOne, read, err = c.parseVariableStringField(record[length:], 15); err != nil {
		fieldError("SendersChargesOne", err)
	}
	length += read

	if c.SendersChargesTwo, read, err = c.parseVariableStringField(record[length:], 15); err != nil {
		fieldError("SendersChargesTwo", err)
	}
	length += read

	if c.SendersChargesThree, read, err = c.parseVariableStringField(record[length:], 15); err != nil {
		fieldError("SendersChargesThree", err)
	}
	length += read

	if c.SendersChargesFour, read, err = c.parseVariableStringField(record[length:], 15); err != nil {
		fieldError("SendersChargesFour", err)
	}
	length += read

	return length, nil
}

func (c *Charges) UnmarshalJSON(data []byte) error {
	type Alias Charges
	aux := struct {
		*Alias
	}{
		(*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.tag = TagCharges
	return nil
}

// String writes Charges
func (c *Charges) String() string {
	var buf strings.Builder
	buf.Grow(67)
	buf.WriteString(c.tag)
	buf.WriteString(c.ChargeDetailsField())
	buf.WriteString(c.SendersChargesOneField())
	buf.WriteString(c.SendersChargesTwoField())
	buf.WriteString(c.SendersChargesThreeField())
	buf.WriteString(c.SendersChargesFourField())
	return buf.String()
}

// Validate performs WIRE format rule checks on Charges and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (c *Charges) Validate() error {
	if err := c.fieldInclusion(); err != nil {
		return err
	}
	if err := c.isChargeDetails(c.ChargeDetails); err != nil {
		return fieldError("ChargeDetails", ErrChargeDetails, c.ChargeDetails)
	}
	if err := c.isAlphanumeric(c.SendersChargesOne); err != nil {
		return fieldError("SendersChargesOne", err, c.SendersChargesOne)
	}
	/*	if err := c.validateCharges(c.SendersChargesOne); err != nil {
		return fieldError("SendersChargesOne", err, c.SendersChargesOne)
	}*/
	if err := c.isAlphanumeric(c.SendersChargesTwo); err != nil {
		return fieldError("SendersChargesTwo", err, c.SendersChargesTwo)
	}
	/*	if err := c.validateCharges(c.SendersChargesTwo); err != nil {
		return fieldError("SendersChargesTwo", err, c.SendersChargesTwo)
	}*/
	if err := c.isAlphanumeric(c.SendersChargesThree); err != nil {
		return fieldError("SendersChargesThree", err, c.SendersChargesThree)
	}
	/*	if err := c.validateCharges(c.SendersChargesThree); err != nil {
		return fieldError("SendersChargesThree", err, c.SendersChargesThree)
	}*/
	if err := c.isAlphanumeric(c.SendersChargesFour); err != nil {
		return fieldError("SendersChargesFour", err, c.SendersChargesFour)
	}
	/*	if err := c.validateCharges(c.SendersChargesFour); err != nil {
		return fieldError("SendersChargesFour", err, c.SendersChargesFour)
	}*/
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (c *Charges) fieldInclusion() error {
	return nil
}

// ChargeDetailsField gets a string of the ChargeDetails field
func (c *Charges) ChargeDetailsField() string {
	return c.alphaField(c.ChargeDetails, 1)
}

// SendersChargesOneField gets a string of the SendersChargesOne field
func (c *Charges) SendersChargesOneField() string {
	return c.alphaVariableField(c.SendersChargesOne, 15, c.isVariableLength)
}

// SendersChargesTwoField gets a string of the SendersChargesTwo field
func (c *Charges) SendersChargesTwoField() string {
	return c.alphaVariableField(c.SendersChargesTwo, 15, c.isVariableLength)
}

// SendersChargesThreeField gets a string of the SendersChargesThree field
func (c *Charges) SendersChargesThreeField() string {
	return c.alphaVariableField(c.SendersChargesThree, 15, c.isVariableLength)
}

// SendersChargesFourField gets a string of the SendersChargesFour field
func (c *Charges) SendersChargesFourField() string {
	return c.alphaVariableField(c.SendersChargesFour, 15, c.isVariableLength)
}
