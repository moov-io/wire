package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// MockAdjustment creates a Adjustment
func mockAdjustment() *Adjustment {
	adj := NewAdjustment()
	adj.AdjustmentReasonCode = PricingError
	adj.CreditDebitIndicator = CreditIndicator
	adj.RemittanceAmount.CurrencyCode = "USD"
	adj.RemittanceAmount.Amount = "1234.56"
	adj.AdditionalInfo = " Adjustment Additional Information"
	return adj
}

// TestMockAdjustment validates mockAdjustment
func TestMockAdjustment(t *testing.T) {
	adj := mockAdjustment()
	if err := adj.Validate(); err != nil {
		t.Error("mockAdjustment does not validate and will break other tests")
	}
}

// TestAdjustmentReasonCodeValid validates Adjustment AdjustmentReasonCode
func TestAdjustmentReasonCodeValid(t *testing.T) {
	adj := mockAdjustment()
	adj.AdjustmentReasonCode = "ZZ"
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrAdjustmentReasonCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestCreditDebitIndicatorValid validates Adjustment CreditDebitIndicator
func TestCreditDebitIndicatorValid(t *testing.T) {
	adj := mockAdjustment()
	adj.CreditDebitIndicator = "ZZZZ"
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrCreditDebitIndicator) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAdjustmentAmountValid validates Adjustment Amount
func TestAdjustmentAmountValid(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.Amount = "X,"
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAdjustmentCurrencyCodeValid validates Adjustment CurrencyCode
func TestAdjustmentCurrencyCodeValid(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.CurrencyCode = "XZP"
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAdjustmentReasonCodeRequired validates Adjustment AdjustmentReasonCode is required
func TestAdjustmentReasonCodeRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.AdjustmentReasonCode = ""
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestCreditDebitIndicatorRequired validates Adjustment CreditDebitIndicator is required
func TestCreditDebitIndicatorRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.CreditDebitIndicator = ""
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAdjustmentAmountRequired validates Adjustment Amount is required
func TestAdjustmentAmountRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.Amount = ""
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAdjustmentCurrencyCodeRequired validates Adjustment CurrencyCode is required
func TestAdjustmentCurrencyCodeRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.CurrencyCode = ""
	if err := adj.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAdjustmentWrongLength parses a wrong Adjustment record length
func TestParseAdjustmentWrongLength(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56Z             Adjustment Additional Information                                                                                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	adj := mockAdjustment()
	fwm.SetAdjustment(adj)
	err := r.parseAdjustment()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(174, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAdjustmentReaderParseError parses a wrong Adjustment reader parse error
func TestParseAdjustmentReaderParseError(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56Z             Adjustment Additional Information                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	adj := mockAdjustment()
	fwm.SetAdjustment(adj)
	err := r.parseAdjustment()
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
