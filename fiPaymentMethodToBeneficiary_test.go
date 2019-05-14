package wire

import (
	"github.com/moov-io/base"
	"strings"
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

// TestParseFIPaymentMethodToBeneficiaryWrongLength parses a wrong FIPaymentMethodToBeneficiary record length
func TestParseFIPaymentMethodToBeneficiaryWrongLength(t *testing.T) {
	var line = "{6420}CHECKAdditional Information      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.SetFIPaymentMethodToBeneficiary(pm)
	err := r.parseFIPaymentMethodToBeneficiary()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(41, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIPaymentMethodToBeneficiaryReaderParseError parses a wrong FIPaymentMethodToBeneficiary reader parse error
func TestParseFIPaymentMethodToBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6420}CHECK®dditional Information        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.SetFIPaymentMethodToBeneficiary(pm)
	err := r.parseFIPaymentMethodToBeneficiary()
	if err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIPaymentMethodToBeneficiaryTagError validates a FIPaymentMethodToBeneficiary tag
func TestFIPaymentMethodToBeneficiaryTagError(t *testing.T) {
	pm := mockFIPaymentMethodToBeneficiary()
	pm.tag = "{9999}"
	if err := pm.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
