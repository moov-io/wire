package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockAccountCreditedDrawdown creates a AccountCreditedDrawdown
func mockAccountCreditedDrawdown() *AccountCreditedDrawdown {
	creditDD := NewAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = "123456789"
	return creditDD
}

// TestMockAccountCreditedDrawdown validates mockAccountCreditedDrawdown
func TestMockAccountCreditedDrawdown(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	if err := creditDD.Validate(); err != nil {
		t.Error("mockAccountCreditedDrawdown does not validate and will break other tests")
	}
}

// TestAccountCreditedDrawDownNumberAlphaNumeric validates AccountCreditedDrawdown is alphanumeric
func TestDrawdownCreditAccountNumberAlphaNumeric(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = "®"
	if err := creditDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAccountCreditedDrawdownNumberRequired validates AccountCreditedDrawdown is required
func TestDrawdownCreditAccountNumberRequired(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = ""
	if err := creditDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAccountCreditedDrawdownWrongLength parses a wrong AccountCreditedDrawdown record length
func TestParseAccountCreditedDrawdownWrongLength(t *testing.T) {
	var line = "{5400}12345678"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	crediDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(crediDD)
	err := r.parseAccountCreditedDrawdown()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(15, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAccountCreditedDrawdownReaderParseError parses a wrong AccountCreditedDrawdown reader parse error
func TestParseAccountCreditedDrawdownReaderParseError(t *testing.T) {
	var line = "{5400}12345678Z"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	crediDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(crediDD)
	err := r.parseAccountCreditedDrawdown()
	if err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAccountCreditedDrawdownTagError validates AccountCreditedDrawdown tag
func TestAccountCreditedDrawdownTagError(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.tag = "{9999}"
	if err := creditDD.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
