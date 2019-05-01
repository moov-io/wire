package wire

import (
	"github.com/moov-io/base"
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
