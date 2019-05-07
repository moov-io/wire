package wire

import (
	"github.com/moov-io/base"
	"strings"
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
func TestInstructedAmountRequired(t *testing.T) {
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
func TestInstructedAmountValid(t *testing.T) {
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

// TestParseInstructedAmountWrongLength parses a wrong InstructedAmount record length
func TestParseInstructedAmountWrongLength(t *testing.T) {
	var line = "{3710}USD4567,89"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	err := r.parseInstructedAmount()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(24, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseInstructedAmountReaderParseError parses a wrong InstructedAmount reader parse error
func TestParseInstructedAmountReaderParseError(t *testing.T) {
	var line = "{3710}USD000000004567Z89"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	err := r.parseInstructedAmount()
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