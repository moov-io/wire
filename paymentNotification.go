// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &PaymentNotification{}

// PaymentNotification is the PaymentNotification of the wire
type PaymentNotification struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
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
func NewPaymentNotification(isVariable bool) *PaymentNotification {
	pn := &PaymentNotification{
		tag:              TagPaymentNotification,
		isVariableLength: isVariable,
	}
	return pn
}

// Parse takes the input string and parses the PaymentNotification values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (pn *PaymentNotification) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 13 {
		return 0, NewTagWrongLengthErr(13, len(record))
	}

	pn.tag = record[:6]
	pn.PaymentNotificationIndicator = pn.parseStringField(record[6:7])

	length := 7
	read := 0

	pn.ContactNotificationElectronicAddress, read = pn.parseVariableStringField(record[length:], 2048)
	length += read

	pn.ContactName, read = pn.parseVariableStringField(record[length:], 140)
	length += read

	pn.ContactPhoneNumber, read = pn.parseVariableStringField(record[length:], 35)
	length += read

	pn.ContactMobileNumber, read = pn.parseVariableStringField(record[length:], 35)
	length += read

	pn.ContactFaxNumber, read = pn.parseVariableStringField(record[length:], 35)
	length += read

	pn.EndToEndIdentification, read = pn.parseVariableStringField(record[length:], 35)
	length += read

	return length, nil
}

func (pn *PaymentNotification) UnmarshalJSON(data []byte) error {
	type Alias PaymentNotification
	aux := struct {
		*Alias
	}{
		(*Alias)(pn),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	pn.tag = TagPaymentNotification
	return nil
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
	if pn.tag != TagPaymentNotification {
		return fieldError("tag", ErrValidTagForType, pn.tag)
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

// PaymentNotificationIndicatorField gets a string of PaymentNotificationIndicator field
func (pn *PaymentNotification) PaymentNotificationIndicatorField() string {
	return pn.alphaField(pn.PaymentNotificationIndicator, 1)
}

// ContactNotificationElectronicAddressField gets a string of ContactNotificationElectronicAddress field
func (pn *PaymentNotification) ContactNotificationElectronicAddressField() string {
	return pn.alphaVariableField(pn.ContactNotificationElectronicAddress, 2048, pn.isVariableLength)
}

// ContactNameField gets a string of ContactName field
func (pn *PaymentNotification) ContactNameField() string {
	return pn.alphaVariableField(pn.ContactName, 140, pn.isVariableLength)
}

// ContactPhoneNumberField gets a string of ContactPhoneNumberField field
func (pn *PaymentNotification) ContactPhoneNumberField() string {
	return pn.alphaVariableField(pn.ContactPhoneNumber, 35, pn.isVariableLength)
}

// ContactMobileNumberField gets a string of ContactMobileNumber field
func (pn *PaymentNotification) ContactMobileNumberField() string {
	return pn.alphaVariableField(pn.ContactMobileNumber, 35, pn.isVariableLength)
}

// ContactFaxNumberField gets a string of FaxNumber field
func (pn *PaymentNotification) ContactFaxNumberField() string {
	return pn.alphaVariableField(pn.ContactFaxNumber, 35, pn.isVariableLength)
}

// EndToEndIdentificationField gets a string of EndToEndIdentification field
func (pn *PaymentNotification) EndToEndIdentificationField() string {
	return pn.alphaVariableField(pn.EndToEndIdentification, 35, pn.isVariableLength)
}
