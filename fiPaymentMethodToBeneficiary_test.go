package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockFIPaymentMethodToBeneficiary creates a FIPaymentMethodToBeneficiary
func mockFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	return pm
}

// TestMockFIPaymentMethodToBeneficiary validates mockFIPaymentMethodToBeneficiary
func TestMockFIPaymentMethodToBeneficiary(t *testing.T) {
	pm := mockFIPaymentMethodToBeneficiary()
	if err := pm.Validate(); err != nil {
		t.Error("mockFIPaymentMethodToBeneficiary does not validate and will break other tests")
	}
}

// TestPaymentMethodValid validates FIPaymentMethodToBeneficiary PaymentMethod
func TestPaymentMethodValid(t *testing.T) {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = ""
	if err := pm.Validate(); err != nil {
		if !base.Match(err, ErrFieldInclusion) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAdditionalInformationAlphaNumeric validates FIPaymentMethodToBeneficiary AdditionalInformation is alphanumeric
func TestAdditionalInformationAlphaNumeric(t *testing.T) {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.AdditionalInformation = "®"
	if err := pm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}