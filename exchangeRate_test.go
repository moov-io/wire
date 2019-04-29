package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockExchangeRate creates a ExchangeRate
func mockExchangeRate() *ExchangeRate {
	eRate := NewExchangeRate()
	eRate.ExchangeRate = "1,2345"
	return eRate
}

// TestMockExchangeRate validates mockExchangeRate
func TestMockExchangeRate(t *testing.T) {
	eRate := mockExchangeRate()
	if err := eRate.Validate(); err != nil {
		t.Error("mockExchangeRate does not validate and will break other tests")
	}
}

// TestExchangeRate validates ExchangeRate
func TestExchangeRateNumeric(t *testing.T) {
	eRate := mockExchangeRate()
	eRate.ExchangeRate = "1,--0.00"
	if err := eRate.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
