// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// FIPaymentMethodToBeneficiary is the financial institution payment method to beneficiary
type FIPaymentMethodToBeneficiary struct {
	// tag
	tag string
	// PaymentMethod is payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`
	// Additional is additional information
	AdditionalInformation string `json:"Additional,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIPaymentMethodToBeneficiary returns a new FIPaymentMethodToBeneficiary
func NewFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	pm := &FIPaymentMethodToBeneficiary{
		tag:           TagFIPaymentMethodToBeneficiary,
		PaymentMethod: "CHECK",
	}
	return pm
}

// Parse takes the input string and parses the FIPaymentMethodToBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pm *FIPaymentMethodToBeneficiary) Parse(record string) error {
	if utf8.RuneCountInString(record) != 41 {
		return NewTagWrongLengthErr(41, len(record))
	}
	pm.tag = record[:6]
	pm.PaymentMethod = pm.parseStringField(record[6:11])
	pm.AdditionalInformation = pm.parseStringField(record[11:41])
	return nil
}

// String writes FIPaymentMethodToBeneficiary
func (pm *FIPaymentMethodToBeneficiary) String() string {
	var buf strings.Builder
	buf.Grow(41)
	buf.WriteString(pm.tag)
	buf.WriteString(pm.PaymentMethodField())
	buf.WriteString(pm.AdditionalInformationField())
	return buf.String()
}

// Validate performs WIRE format rule checks on FIPaymentMethodToBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (pm *FIPaymentMethodToBeneficiary) Validate() error {
	if err := pm.fieldInclusion(); err != nil {
		return err
	}
	if pm.tag != TagFIPaymentMethodToBeneficiary {
		return fieldError("tag", ErrValidTagForType, pm.tag)
	}
	if err := pm.isAlphanumeric(pm.AdditionalInformation); err != nil {
		return fieldError("AdditionalInformation", err, pm.AdditionalInformation)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (pm *FIPaymentMethodToBeneficiary) fieldInclusion() error {
	if pm.PaymentMethod != PaymentMethod {
		return fieldError("PaymentMethod", ErrFieldInclusion, pm.PaymentMethod)
	}
	return nil
}

// PaymentMethodField gets a string of the PaymentMethod field
func (pm *FIPaymentMethodToBeneficiary) PaymentMethodField() string {
	return pm.alphaField(pm.PaymentMethod, 5)
}

// AdditionalInformationField gets a string of the AdditionalInformation field
func (pm *FIPaymentMethodToBeneficiary) AdditionalInformationField() string {
	return pm.alphaField(pm.AdditionalInformation, 30)
}
