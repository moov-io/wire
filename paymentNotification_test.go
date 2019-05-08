package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
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
	if err := pn.Validate(); err != nil {
		t.Error("mockPaymentNotification does not validate and will break other tests")
	}
}

// TestPaymentNotificationIndicatorNumeric validates PaymentNotificationIndicator is numeric
func TestPaymentNotificationIndicatorNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.PaymentNotificationIndicator = "Z"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestPaymentNotificationContactNotificationElectronicAddressAlphaNumeric validates PaymentNotification ContactNotificationElectronicAddress is alphanumeric
func TestContactNotificationElectronicAddressAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactNotificationElectronicAddress = "®"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestPaymentNotificationContactNameAlphaNumeric validates PaymentNotification ContactName is alphanumeric
func TestContactNameAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactName = "®"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestPaymentNotificationContactPhoneNumberAlphaNumeric validates PaymentNotification ContactPhoneNumber is alphanumeric
func TestContactPhoneNumberAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactPhoneNumber = "®"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestPaymentNotificationContactMobileNumberAlphaNumeric validates PaymentNotification ContactMobileNumber is alphanumeric
func TestContactMobileNumberAlphaNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactMobileNumber = "®"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestPaymentNotificationContactFaxNumberAlphaNumeric validates PaymentNotification ContactFaxNumber is alphanumeric
func TestContactContactFaxNumberNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.ContactFaxNumber = "®"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestPaymentNotificationEndToEndIdentificationAlphaNumeric validates PaymentNotification EndToEndIdentification is alphanumeric
func TestContactEndToEndIdentificationNumeric(t *testing.T) {
	pn := mockPaymentNotification()
	pn.EndToEndIdentification = "®"
	if err := pn.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParsePaymentNotificationWrongLength parses a wrong PaymentNotification record length
func TestParsePaymentNotificationWrongLength(t *testing.T) {
	var line = "{3620}1http://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  Contact Name                                                                                                                                5555551212                         5551231212                         5554561212                         End To End Identification"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	pn := mockPaymentNotification()
	fwm.SetPaymentNotification(pn)
	err := r.parsePaymentNotification()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(2335, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParsePaymentNotificationReaderParseError parses a wrong PaymentNotification reader parse error
func TestParsePaymentNotificationReaderParseError(t *testing.T) {
	var line = "{3620}Zhttp://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            Contact Name                                                                                                                                5555551212                         5551231212                         5554561212                         End To End Identification"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	pn := mockPaymentNotification()
	fwm.SetPaymentNotification(pn)
	err := r.parsePaymentNotification()
	if err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
