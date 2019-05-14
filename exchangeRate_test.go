package wire

import (
	"github.com/moov-io/base"
	"strings"
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

// TestParseExchangeRateWrongLength parses a wrong ExchangeRate record length
func TestParseExchangeRateWrongLength(t *testing.T) {
	var line = "{3720}1,2345"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	err := r.parseExchangeRate()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(18, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseExchangeRateReaderParseError parses a wrong ExchangeRate reader parse error
func TestParseExchangeRateReaderParseError(t *testing.T) {
	var line = "{3720}1,2345Z     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	err := r.parseExchangeRate()
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

// TestExchangeRateTagError validates a ExchangeRate tag
func TestExchangeRateTagError(t *testing.T) {
	eRate := mockCurrencyInstructedAmount()
	eRate.tag = "{9999}"
	if err := eRate.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
