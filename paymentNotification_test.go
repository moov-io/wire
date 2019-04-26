package wire

import (
	"github.com/moov-io/base"
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
