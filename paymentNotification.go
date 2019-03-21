// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// PaymentNotification is the PaymentNotification of the wire
type PaymentNotification struct {
	// tag
	tag string
	// PaymentNotificationIndicator
	// * `0 - 6` - Reserved for market practice conventions.
	// * `7 - 9` - Reserved for bilateral agreements between Fedwire senders and receivers.
	PaymentNotificationIndicator string `json:"paymentNotificationIndicator,omitempty"`
	// ContactNotificationElectronicAddress
	ContactNotificationElectronicAddress string `json:"contactNotificationElectronicAddress,omitempty"`
	// ContactName
	ContactName string `json:"contactName,omitempty"`
	// ContactPhoneNumber
	ContactPhoneNumber string `json:"contactPhoneNumber,omitempty"`
	// ContactMobileNumber
	ContactMobileNumber string `json:"contactMobileNumber,omitempty"`
	// FaxNumber
	FaxNumber string `json:"faxNumber,omitempty"`
	// EndToEndIdentification
	EndToEndIdentification string `json:"endToEndIdentification,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewPaymentNotification returns a new PaymentNotification
func NewPaymentNotification() PaymentNotification  {
	pn := PaymentNotification {
		tag: TagPaymentNotification,
	}
	return pn
}

// Parse takes the input string and parses the PaymentNotification values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pn *PaymentNotification) Parse(record string) {
}

// String writes PaymentNotification
func (pn *PaymentNotification) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(2338)
	buf.WriteString(pn.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiverDepositoryInstitution and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (pn *PaymentNotification) Validate() error {
	if err := pn.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (pn *PaymentNotification) fieldInclusion() error {
	return nil
}

