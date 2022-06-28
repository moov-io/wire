// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

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
func (pn *PaymentNotification) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	pn.tag = record[:6]
	length := 6

	value, read, err := pn.parseVariableStringField(record[length:], 1)
	if err != nil {
		return fieldError("PaymentNotificationIndicator", err)
	}
	pn.PaymentNotificationIndicator = value
	length += read

	value, read, err = pn.parseVariableStringField(record[length:], 2048)
	if err != nil {
		return fieldError("ContactNotificationElectronicAddress", err)
	}
	pn.ContactNotificationElectronicAddress = value
	length += read

	value, read, err = pn.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("ContactName", err)
	}
	pn.ContactName = value
	length += read

	value, read, err = pn.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactPhoneNumber", err)
	}
	pn.ContactPhoneNumber = value
	length += read

	value, read, err = pn.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactMobileNumber", err)
	}
	pn.ContactMobileNumber = value
	length += read

	value, read, err = pn.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactFaxNumber", err)
	}
	pn.ContactFaxNumber = value
	length += read

	value, read, err = pn.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("EndToEndIdentification", err)
	}
	pn.EndToEndIdentification = value
	length += read

	if !pn.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
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

// String returns a fixed-width PaymentNotification record
func (pn *PaymentNotification) String() string {
	return pn.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a PaymentNotification record formatted according to the FormatOptions
func (pn *PaymentNotification) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(2335)

	buf.WriteString(pn.tag)
	buf.WriteString(pn.FormatPaymentNotificationIndicator(options))
	buf.WriteString(pn.FormatContactNotificationElectronicAddress(options))
	buf.WriteString(pn.FormatContactName(options))
	buf.WriteString(pn.FormatContactPhoneNumber(options))
	buf.WriteString(pn.FormatContactMobileNumber(options))
	buf.WriteString(pn.FormatContactFaxNumber(options))
	buf.WriteString(pn.FormatEndToEndIdentification(options))

	if options.VariableLengthFields {
		return pn.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
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
func (pn *PaymentNotification) FormatPaymentNotificationIndicator(options FormatOptions) string {
	return pn.formatAlphaField(pn.PaymentNotificationIndicator, 1, options)
}

// ContactNotificationElectronicAddressField gets a string of ContactNotificationElectronicAddress field
func (pn *PaymentNotification) FormatContactNotificationElectronicAddress(options FormatOptions) string {
	return pn.formatAlphaField(pn.ContactNotificationElectronicAddress, 2048, options)
}

// ContactNameField gets a string of ContactName field
func (pn *PaymentNotification) FormatContactName(options FormatOptions) string {
	return pn.formatAlphaField(pn.ContactName, 140, options)
}

// ContactPhoneNumberField gets a string of ContactPhoneNumberField field
func (pn *PaymentNotification) FormatContactPhoneNumber(options FormatOptions) string {
	return pn.formatAlphaField(pn.ContactPhoneNumber, 35, options)
}

// ContactMobileNumberField gets a string of ContactMobileNumber field
func (pn *PaymentNotification) FormatContactMobileNumber(options FormatOptions) string {
	return pn.formatAlphaField(pn.ContactMobileNumber, 35, options)
}

// ContactFaxNumberField gets a string of FaxNumber field
func (pn *PaymentNotification) FormatContactFaxNumber(options FormatOptions) string {
	return pn.formatAlphaField(pn.ContactFaxNumber, 35, options)
}

// EndToEndIdentificationField gets a string of EndToEndIdentification field
func (pn *PaymentNotification) FormatEndToEndIdentification(options FormatOptions) string {
	return pn.formatAlphaField(pn.EndToEndIdentification, 35, options)
}
