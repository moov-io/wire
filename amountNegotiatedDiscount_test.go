package wire

import (
	"github.com/moov-io/base"
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