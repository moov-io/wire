// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// Charges is the Charges of the wire
type Charges struct {
	// tag
	tag string
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
func NewCharges() Charges {
	c := Charges{
		tag: TagCharges,
	}
	return c
}

// Parse takes the input string and parses the Charges values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (c *Charges) Parse(record string) {
	c.tag = record[:6]
	c.ChargeDetails = c.parseStringField(record[6:8])
	c.SendersChargesOne = c.parseStringField(record[8:23])
	c.SendersChargesTwo = c.parseStringField(record[23:38])
	c.SendersChargesThree = c.parseStringField(record[38:53])
	c.SendersChargesFour = c.parseStringField(record[53:68])
}

// String writes Charges
func (c *Charges) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(68)
	buf.WriteString(c.tag)
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

	// ToDo: specific validation for Charges 1-4 fields
	if err := c.isAlphanumeric(c.SendersChargesOne); err != nil {
		return fieldError("SendersChargesOne", err, c.SendersChargesOne)
	}
	if err := c.isAlphanumeric(c.SendersChargesTwo); err != nil {
		return fieldError("SendersChargesTwo", err, c.SendersChargesTwo)
	}
	if err := c.isAlphanumeric(c.SendersChargesThree); err != nil {
		return fieldError("SendersChargesThree", err, c.SendersChargesThree)
	}
	if err := c.isAlphanumeric(c.SendersChargesFour); err != nil {
		return fieldError("SendersChargesFour", err, c.SendersChargesFour)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (c *Charges) fieldInclusion() error {
	return nil
}
