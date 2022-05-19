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
func (pn *PaymentNotification) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 13 {
		return 0, NewTagWrongLengthErr(13, len(record))
	}

	var err error
	var length, read int

	if pn.tag, read, err = pn.parseTag(record); err != nil {
		return 0, fieldError("PaymentNotification.Tag", err)
	}
	length += read

	pn.PaymentNotificationIndicator = pn.parseStringField(record[length : length+1])
	length += 1

	if pn.ContactNotificationElectronicAddress, read, err = pn.parseVariableStringField(record[length:], 2048); err != nil {
		return 0, fieldError("ContactNotificationElectronicAddress", err)
	}
	length += read

	if pn.ContactName, read, err = pn.parseVariableStringField(record[length:], 140); err != nil {
		return 0, fieldError("ContactName", err)
	}
	length += read

	if pn.ContactPhoneNumber, read, err = pn.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("ContactPhoneNumber", err)
	}
	length += read

	if pn.ContactMobileNumber, read, err = pn.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("ContactMobileNumber", err)
	}
	length += read

	if pn.ContactFaxNumber, read, err = pn.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("ContactFaxNumber", err)
	}
	length += read

	if pn.EndToEndIdentification, read, err = pn.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("EndToEndIdentification", err)
	}
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
func (pn *PaymentNotification) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(2335)

	buf.WriteString(pn.tag)
	buf.WriteString(pn.PaymentNotificationIndicatorField())
	buf.WriteString(pn.ContactNotificationElectronicAddressField(isCompressed))
	buf.WriteString(pn.ContactNameField(isCompressed))
	buf.WriteString(pn.ContactPhoneNumberField(isCompressed))
	buf.WriteString(pn.ContactMobileNumberField(isCompressed))
	buf.WriteString(pn.ContactFaxNumberField(isCompressed))
	buf.WriteString(pn.EndToEndIdentificationField(isCompressed))

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
func (pn *PaymentNotification) ContactNotificationElectronicAddressField(isCompressed bool) string {
	return pn.alphaVariableField(pn.ContactNotificationElectronicAddress, 2048, isCompressed)
}

// ContactNameField gets a string of ContactName field
func (pn *PaymentNotification) ContactNameField(isCompressed bool) string {
	return pn.alphaVariableField(pn.ContactName, 140, isCompressed)
}

// ContactPhoneNumberField gets a string of ContactPhoneNumberField field
func (pn *PaymentNotification) ContactPhoneNumberField(isCompressed bool) string {
	return pn.alphaVariableField(pn.ContactPhoneNumber, 35, isCompressed)
}

// ContactMobileNumberField gets a string of ContactMobileNumber field
func (pn *PaymentNotification) ContactMobileNumberField(isCompressed bool) string {
	return pn.alphaVariableField(pn.ContactMobileNumber, 35, isCompressed)
}

// ContactFaxNumberField gets a string of FaxNumber field
func (pn *PaymentNotification) ContactFaxNumberField(isCompressed bool) string {
	return pn.alphaVariableField(pn.ContactFaxNumber, 35, isCompressed)
}

// EndToEndIdentificationField gets a string of EndToEndIdentification field
func (pn *PaymentNotification) EndToEndIdentificationField(isCompressed bool) string {
	return pn.alphaVariableField(pn.EndToEndIdentification, 35, isCompressed)
}
