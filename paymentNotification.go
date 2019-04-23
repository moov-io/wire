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
	ContactFaxNumber string `json:"faxNumber,omitempty"`
	// EndToEndIdentification
	EndToEndIdentification string `json:"endToEndIdentification,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewPaymentNotification returns a new PaymentNotification
func NewPaymentNotification() *PaymentNotification {
	pn := &PaymentNotification{
		tag: TagPaymentNotification,
	}
	return pn
}

// Parse takes the input string and parses the PaymentNotification values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pn *PaymentNotification) Parse(record string) {
	pn.tag = record[:6]
	pn.PaymentNotificationIndicator = pn.parseStringField(record[6:7])
	pn.ContactNotificationElectronicAddress = pn.parseStringField(record[7:2055])
	pn.ContactName = pn.parseStringField(record[2055:2195])
	pn.ContactPhoneNumber = pn.parseStringField(record[2195:2230])
	pn.ContactMobileNumber = pn.parseStringField(record[2230:2265])
	pn.ContactFaxNumber = pn.parseStringField(record[2265:2300])
	pn.EndToEndIdentification = pn.parseStringField(record[2300:2335])
}

// String writes PaymentNotification
func (pn *PaymentNotification) String() string {
	var buf strings.Builder
	buf.Grow(2335)
	buf.WriteString(pn.tag)
	buf.WriteString(pn.PaymentNotificationIndicatorField())
	buf.WriteString(pn.ContactNotificationElectronicAddressField())
	buf.WriteString(pn.ContactNameField())
	buf.WriteString(pn.ContactPhoneNumberField())
	buf.WriteString(pn.ContactMobileNumberField())
	buf.WriteString(pn.ContactFaxNumberField())
	buf.WriteString(pn.EndToEndIdentificationField())
	return buf.String()
}

// Validate performs WIRE format rule checks on PaymentNotification and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (pn *PaymentNotification) Validate() error {
	if err := pn.fieldInclusion(); err != nil {
		return err
	}
	if err := pn.isNumeric(pn.PaymentNotificationIndicator); err != nil {
		return fieldError("PaymentNotificationIndicator", err, pn.PaymentNotificationIndicator)
	}
	if err := pn.isAlphanumeric(pn.ContactNotificationElectronicAddress); err != nil {
		return fieldError("ContactNotificationElectronicAddress", err, pn.ContactNotificationElectronicAddress)
	}
	if err := pn.isAlphanumeric(pn.ContactName); err != nil {
		return fieldError("ContactName", err, pn.ContactName)
	}
	if err := pn.isAlphanumeric(pn.ContactPhoneNumber); err != nil {
		return fieldError("ContactPhoneNumber", err, pn.ContactPhoneNumber)
	}
	if err := pn.isAlphanumeric(pn.ContactMobileNumber); err != nil {
		return fieldError("ContactMobileNumber", err, pn.ContactMobileNumber)
	}
	if err := pn.isAlphanumeric(pn.ContactFaxNumber); err != nil {
		return fieldError("FaxNumber", err, pn.ContactFaxNumber)
	}
	if err := pn.isAlphanumeric(pn.EndToEndIdentification); err != nil {
		return fieldError("EndToEndIdentification", err, pn.EndToEndIdentification)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (pn *PaymentNotification) fieldInclusion() error {
	return nil
}

// PaymentNotificationIndicatorField gets a string of PaymentNotificationIndicator field
func (pn *PaymentNotification) PaymentNotificationIndicatorField() string {
	return pn.alphaField(pn.PaymentNotificationIndicator, 1)
}

// ContactNotificationElectronicAddressField gets a string of ContactNotificationElectronicAddress field
func (pn *PaymentNotification) ContactNotificationElectronicAddressField() string {
	return pn.alphaField(pn.ContactNotificationElectronicAddress, 2048)
}

// ContactNameField gets a string of ContactName field
func (pn *PaymentNotification) ContactNameField() string {
	return pn.alphaField(pn.ContactName, 140)
}

// ContactPhoneNumberField gets a string of ContactPhoneNumberField field
func (pn *PaymentNotification) ContactPhoneNumberField() string {
	return pn.alphaField(pn.ContactPhoneNumber, 35)
}

// ContactMobileNumberField gets a string of ContactMobileNumber field
func (pn *PaymentNotification) ContactMobileNumberField() string {
	return pn.alphaField(pn.ContactMobileNumber, 35)
}

// ContactFaxNumberField gets a string of FaxNumber field
func (pn *PaymentNotification) ContactFaxNumberField() string {
	return pn.alphaField(pn.ContactFaxNumber, 35)
}

// EndToEndIdentificationField gets a string of EndToEndIdentification field
func (pn *PaymentNotification) EndToEndIdentificationField() string {
	return pn.alphaField(pn.EndToEndIdentification, 35)
}
