package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

//  CurrencyInstructedAmount creates a CurrencyInstructedAmount
func mockCurrencyInstructedAmount() *CurrencyInstructedAmount {
	cia := NewCurrencyInstructedAmount()
	cia.SwiftFieldTag = "Swift Field Tag"
	cia.Amount = "1500,49"
	return cia
}

// TestMockCurrencyInstructedAmount validates mockCurrencyInstructedAmount
func TestMockCurrencyInstructedAmount(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	if err := cia.Validate(); err != nil {
		t.Error("mockCurrencyInstructedAmount does not validate and will break other tests")
	}
}

// TestCurrencyInstructedAmountSwiftFieldTagAlphaNumeric validates CurrencyInstructedAmount SwiftFieldTag is alphanumeric
func TestCurrencyInstructedAmountSwiftFieldTagAlphaNumeric(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.SwiftFieldTag = "®"
	if err := cia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestCurrencyInstructedAmountValid validates CurrencyInstructedAmount InstructedAmount is valid
func TestCurrencyInstructedAmountValid(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.Amount = "1-0"
	if err := cia.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseCurrencyInstructedAmountWrongLength parses a wrong CurrencyInstructedAmount record length
func TestParseCurrencyInstructedAmountWrongLength(t *testing.T) {
	var line = "{7033}Swift000000000001500,4"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	err := r.parseCurrencyInstructedAmount()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(41, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseCurrencyInstructedAmountReaderParseError parses a wrong CurrencyInstructedAmount reader parse error
func TestParseCurrencyInstructedAmountReaderParseError(t *testing.T) {
	var line = "{7033}Swift00000000Z001500,49"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	err := r.parseCurrencyInstructedAmount()
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

// TestCurrencyInstructedAmountTagError validates a CurrencyInstructedAmount tag
func TestCurrencyInstructedAmountTagError(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.tag = "{9999}"
	if err := cia.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
