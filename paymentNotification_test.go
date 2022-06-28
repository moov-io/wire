package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockPaymentNotification creates a PaymentNotification
func mockPaymentNotification() *PaymentNotification {
	pn := NewPaymentNotification()
	pn.PaymentNotificationIndicator = "1"
	pn.ContactNotificationElectronicAddress = "http://moov.io"
	pn.ContactName = "Contact Name"
	pn.ContactPhoneNumber = "5555551212"
	pn.ContactMobileNumber = "5551231212"
	pn.ContactFaxNumber = "5554561212"
	pn.EndToEndIdentification = "End To End Identification"
	return pn
}

// TestMockPaymentNotification validates mockPaymentNotification
func TestMockPaymentNotification(t *testing.T) {
	pn := mockPaymentNotification()

	require.NoError(t, pn.Validate(), "mockPaymentNotification does not validate and will break other tests")
}

// TestPaymentNotificationIndicatorNumeric validates PaymentNotificationIndicator is numeric
func TestPaymentNotificationIndicatorNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.PaymentNotificationIndicator = "Z"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("PaymentNotificationIndicator", ErrNonNumeric, pn.PaymentNotificationIndicator).Error())
}

// TestPaymentNotificationContactNotificationElectronicAddressAlphaNumeric validates PaymentNotification ContactNotificationElectronicAddress is alphanumeric
func TestContactNotificationElectronicAddressAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactNotificationElectronicAddress = "®"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("ContactNotificationElectronicAddress", ErrNonAlphanumeric, pn.ContactNotificationElectronicAddress).Error())
}

// TestPaymentNotificationContactNameAlphaNumeric validates PaymentNotification ContactName is alphanumeric
func TestContactNameAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactName = "®"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("ContactName", ErrNonAlphanumeric, pn.ContactName).Error())
}

// TestPaymentNotificationContactPhoneNumberAlphaNumeric validates PaymentNotification ContactPhoneNumber is alphanumeric
func TestContactPhoneNumberAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactPhoneNumber = "®"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("ContactPhoneNumber", ErrNonAlphanumeric, pn.ContactPhoneNumber).Error())
}

// TestPaymentNotificationContactMobileNumberAlphaNumeric validates PaymentNotification ContactMobileNumber is alphanumeric
func TestContactMobileNumberAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactMobileNumber = "®"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("ContactMobileNumber", ErrNonAlphanumeric, pn.ContactMobileNumber).Error())
}

// TestPaymentNotificationContactFaxNumberAlphaNumeric validates PaymentNotification ContactFaxNumber is alphanumeric
func TestContactContactFaxNumberNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactFaxNumber = "®"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("FaxNumber", ErrNonAlphanumeric, pn.ContactFaxNumber).Error())
}

// TestPaymentNotificationEndToEndIdentificationAlphaNumeric validates PaymentNotification EndToEndIdentification is alphanumeric
func TestContactEndToEndIdentificationNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.EndToEndIdentification = "®"

	err := pn.Validate()

	require.EqualError(t, err, fieldError("EndToEndIdentification", ErrNonAlphanumeric, pn.EndToEndIdentification).Error())
}

// TestParsePaymentNotificationWrongLength parses a wrong PaymentNotification record length
func TestParsePaymentNotificationWrongLength(t *testing.T) {
	var line = "{3620}1http://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  Contact Name                                                                                                                                5555551212                         5551231212                         5554561212                         End To End Identification"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePaymentNotification()

	require.EqualError(t, err, r.parseError(fieldError("EndToEndIdentification", ErrValidLength)).Error())
}

// TestParsePaymentNotificationReaderParseError parses a wrong PaymentNotification reader parse error
func TestParsePaymentNotificationReaderParseError(t *testing.T) {
	var line = "{3620}Zhttp://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            Contact Name                                                                                                                                5555551212                         5551231212                         5554561212                         End To End Identification"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePaymentNotification()

	require.EqualError(t, err, r.parseError(fieldError("PaymentNotificationIndicator", ErrNonNumeric, "Z")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("PaymentNotificationIndicator", ErrNonNumeric, "Z")).Error())
}

// TestPaymentNotificationTagError validates a PaymentNotification tag
func TestPaymentNotificationTagError(t *testing.T) {
	pn := mockPaymentNotification()
	pn.tag = "{9999}"

	require.EqualError(t, pn.Validate(), fieldError("tag", ErrValidTagForType, pn.tag).Error())
}

// TestStringPaymentNotificationVariableLength parses using variable length
func TestStringPaymentNotificationVariableLength(t *testing.T) {
	var line = "{3620}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePaymentNotification()
	require.Nil(t, err)

	line = "{3620}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePaymentNotification()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3620}*********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePaymentNotification()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3620}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePaymentNotification()
	require.Equal(t, err, nil)
}

// TestStringPaymentNotificationOptions validates Format() formatted according to the FormatOptions
func TestStringPaymentNotificationOptions(t *testing.T) {
	var line = "{3620}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePaymentNotification()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.PaymentNotification
	require.Equal(t, record.String(), "{3620}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{3620}*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
