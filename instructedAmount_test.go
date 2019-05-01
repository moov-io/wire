package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockInstructedAmount creates a InstructedAmount
func mockInstructedAmount() *InstructedAmount {
	ia := NewInstructedAmount()
	ia.CurrencyCode = "USD"
	ia.Amount = "4567,89"
	return ia
}

// TestMockInstructedAmount validates mockInstructedAmount
func TestMockInstructedAmount(t *testing.T) {
	ia := mockInstructedAmount()
	if err := ia.Validate(); err != nil {
		t.Error("mockInstructedAmount does not validate and will break other tests")
	}
}

// TestInstructedAmountAmountRequired validates InstructedAmount Amount is required
func TestInstructedAmountAmountRequired(t *testing.T) {
	ia := mockInstructedAmount()
	ia.Amount = ""
	if err := ia.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructedAmountCurrencyCodeRequired validates InstructedAmount CurrencyCode is required
func TestInstructedAmountCurrencyCodeRequired(t *testing.T) {
	ia := mockInstructedAmount()
	ia.CurrencyCode = ""
	if err := ia.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructedAmountAmountValid validates Amount
func TestInstructedAmountAmountValid(t *testing.T) {
	ia := mockInstructedAmount()
	ia.Amount = "X,"
	if err := ia.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructedAmountCurrencyCodeValid validates Amount
func TestInstructedAmountCurrencyCodeValid(t *testing.T) {
	ia := mockInstructedAmount()
	ia.CurrencyCode = "XZP"
	if err := ia.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
