package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockAmount creates an a Amount
func mockAmount() *Amount {
	a := NewAmount()
	a.Amount = "000001234567"
	return a
}

// TestMockAmount validates mockAmount
func TestMockAmount(t *testing.T) {
	a := mockAmount()
	if err := a.Validate(); err != nil {
		t.Error("mockAmount does not validate and will break other tests")
	}
}

// TestAmountValid validates Amount
func TestAmountValid(t *testing.T) {
	a := mockAmount()
	a.Amount = "X,"
	if err := a.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAmountRequired validates Amount is required
func TestAmountRequired(t *testing.T) {
	a := mockAmount()
	a.Amount = ""
	if err := a.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
