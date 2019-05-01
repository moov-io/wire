package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// GrossAmountRemittanceDocument creates a GrossAmountRemittanceDocument
func mockGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	gard := NewGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = "USD"
	gard.RemittanceAmount.Amount = "1234.56"
	return gard
}

// TestMockGrossAmountRemittanceDocument validates mockGrossAmountRemittanceDocument
func TestMockGrossAmountRemittanceDocument(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	if err := gard.Validate(); err != nil {
		t.Error("mockGrossAmountRemittanceDocument does not validate and will break other tests")
	}
}

// TestGrossAmountRemittanceAmountRequired validates GrossAmountRemittance Amount is required
func TestGrossAmountRemittanceAmountRequired(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.Amount = ""
	if err := gard.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestGrossAmountRemittanceCurrencyCodeRequired validates GrossAmountRemittance CurrencyCode is required
func TestGrossAmountRemittanceCurrencyCodeRequired(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = ""
	if err := gard.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestGrossAmountRemittanceAmountValid validates Amount
func TestGrossAmountRemittanceAmountValid(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.Amount = "X,"
	if err := gard.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestGrossAmountRemittanceCurrencyCodeValid validates Amount
func TestGrossAmountRemittanceCurrencyCodeValid(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = "XZP"
	if err := gard.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
