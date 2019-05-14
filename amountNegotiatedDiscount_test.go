package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// AmountNegotiatedDiscount creates a AmountNegotiatedDiscount
func mockAmountNegotiatedDiscount() *AmountNegotiatedDiscount {
	nd := NewAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = "USD"
	nd.RemittanceAmount.Amount = "1234.56"
	return nd
}

// TestMockAmountNegotiatedDiscount validates mockAmountNegotiatedDiscount
func TestMockAmountNegotiatedDiscount(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	if err := nd.Validate(); err != nil {
		t.Error("mockAmountNegotiatedDiscount does not validate and will break other tests")
	}
}

// TestAmountNegotiatedDiscountAmountValid validates AmountNegotiatedDiscount Amount
func TestAmountNegotiatedDiscountValid(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.Amount = "X,"
	if err := nd.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAmountNegotiatedDiscountCurrencyCodeValid validates AmountNegotiatedDiscount CurrencyCode
func TestAmountNegotiatedDiscountCurrencyCodeValid(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = "XZP"
	if err := nd.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
		}
	}
}

// TestAmountNegotiatedDiscountAmountRequired validates AmountNegotiatedDiscount Amount is required
func TestAmountNegotiatedDiscountRequired(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.Amount = ""
	if err := nd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAmountNegotiatedDiscountCurrencyCodeRequired validates AmountNegotiatedDiscount CurrencyCode is required
func TestAmountNegotiatedDiscountCurrencyCodeRequired(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = ""
	if err := nd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAmountNegotiatedDiscountWrongLength parses a wrong AmountNegotiatedDiscount record length
func TestParseAmountNegotiatedDiscountWrongLength(t *testing.T) {
	var line = "{8550}USD1234.56          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	nd := mockAmountNegotiatedDiscount()
	fwm.SetAmountNegotiatedDiscount(nd)
	err := r.parseAmountNegotiatedDiscount()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(28, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAmountNegotiatedDiscountReaderParseError parses a wrong AmountNegotiatedDiscount reader parse error
func TestParseAmountNegotiatedDiscountReaderParseError(t *testing.T) {
	var line = "{8550}USD1234.56Z           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	nd := mockAmountNegotiatedDiscount()
	fwm.SetAmountNegotiatedDiscount(nd)
	err := r.parseAmountNegotiatedDiscount()
	if err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAmountNegotiatedDiscountTagError validates AmountNegotiatedDiscount tag
func TestAmountNegotiatedDiscountTagError(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.tag = "{9999}"
	if err := nd.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
