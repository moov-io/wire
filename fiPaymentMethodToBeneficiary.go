// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIPaymentMethodToBeneficiary is the financial institution payment method to beneficiary
type FIPaymentMethodToBeneficiary struct {
	// tag
	tag string
	// PaymentMethod is payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`
	// Additional is additional information
	Additional string `json:"Additional,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIPaymentMethodToBeneficiary returns a new FIPaymentMethodToBeneficiary
func NewFIPaymentMethodToBeneficiary() FIPaymentMethodToBeneficiary {
	b := FIPaymentMethodToBeneficiary{
		tag: TagFIPaymentMethodToBeneficiary,
	}
	return b
}

// Parse takes the input string and parses the FIPaymentMethodToBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (b *FIPaymentMethodToBeneficiary) Parse(record string) {
}

// String writes FIPaymentMethodToBeneficiary
func (b *FIPaymentMethodToBeneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(35)
	buf.WriteString(b.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (b *FIPaymentMethodToBeneficiary) Validate() error {
	if err := b.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (b *FIPaymentMethodToBeneficiary) fieldInclusion() error {
	return nil
}
