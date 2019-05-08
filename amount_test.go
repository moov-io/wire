package wire

import (
	"github.com/moov-io/base"
	"strings"
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

// TestParseAmountWrongLength parses a wrong Amount record length
func TestParseAmountWrongLength(t *testing.T) {
	var line = "{2000}00"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	amt := mockAmount()
	fwm.SetAmount(amt)
	err := r.parseAmount()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(18, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAmountReaderParseError parses a wrong Amount reader parse error
func TestParseAmountReaderParseError(t *testing.T) {
	var line = "{2000}00000Z030022"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	amt := mockAmount()
	fwm.SetAmount(amt)
	err := r.parseAmount()
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
