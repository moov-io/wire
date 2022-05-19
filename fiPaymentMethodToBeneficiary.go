// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIPaymentMethodToBeneficiary{}

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
func (pm *FIPaymentMethodToBeneficiary) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, len(record))
	}

	var err error
	var length, read int

	if pm.tag, read, err = pm.parseTag(record); err != nil {
		return 0, fieldError("FIPaymentMethodToBeneficiary.Tag", err)
	}
	length += read

	if pm.PaymentMethod, read, err = pm.parseVariableStringField(record[length:], 5); err != nil {
		return 0, fieldError("PaymentMethod", err)
	}
	length += read

	if pm.AdditionalInformation, read, err = pm.parseVariableStringField(record[length:], 30); err != nil {
		return 0, fieldError("AdditionalInformation", err)
	}
	length += read

	return length, nil
}

func (pm *FIPaymentMethodToBeneficiary) UnmarshalJSON(data []byte) error {
	type Alias FIPaymentMethodToBeneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(pm),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	pm.tag = TagFIPaymentMethodToBeneficiary
	return nil
}

// String writes FIPaymentMethodToBeneficiary
func (pm *FIPaymentMethodToBeneficiary) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(41)

	buf.WriteString(pm.tag)
	buf.WriteString(pm.PaymentMethodField(isCompressed))
	buf.WriteString(pm.AdditionalInformationField(isCompressed))

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
func (pm *FIPaymentMethodToBeneficiary) PaymentMethodField(isCompressed bool) string {
	return pm.alphaVariableField(pm.PaymentMethod, 5, isCompressed)
}

// AdditionalInformationField gets a string of the AdditionalInformation field
func (pm *FIPaymentMethodToBeneficiary) AdditionalInformationField(isCompressed bool) string {
	return pm.alphaVariableField(pm.AdditionalInformation, 30, isCompressed)
}
